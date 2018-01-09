package hercules

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/utils/merkletrie"
)

// This PipelineItem loads the blobs which correspond to the changed files in a commit.
// It must provide the old and the new objects; "cache" rotates and allows to not load
// the same blobs twice. Outdated objects are removed so "cache" never grows big.
type BlobCache struct {
	// Specifies how to handle the situation when we encounter a git submodule - an object without
	// the blob. If false, we look inside .gitmodules and if don't find, raise an error.
	// If true, we do not look inside .gitmodules and always succeed.
	IgnoreMissingSubmodules bool

	repository *git.Repository
	cache      map[plumbing.Hash]*object.Blob
}

const (
	ConfigBlobCacheIgnoreMissingSubmodules = "BlobCache.IgnoreMissingSubmodules"
	DependencyBlobCache                    = "blob_cache"
)

func (cache *BlobCache) Name() string {
	return "BlobCache"
}

func (cache *BlobCache) Provides() []string {
	arr := [...]string{DependencyBlobCache}
	return arr[:]
}

func (cache *BlobCache) Requires() []string {
	arr := [...]string{DependencyTreeChanges}
	return arr[:]
}

func (cache *BlobCache) ListConfigurationOptions() []ConfigurationOption {
	options := [...]ConfigurationOption{{
		Name: ConfigBlobCacheIgnoreMissingSubmodules,
		Description: "Specifies whether to panic if some submodules do not exist and thus " +
			"the corresponding Git objects cannot be loaded.",
		Flag:    "ignore-missing-submodules",
		Type:    BoolConfigurationOption,
		Default: false}}
	return options[:]
}

func (cache *BlobCache) Configure(facts map[string]interface{}) {
	if val, exists := facts[ConfigBlobCacheIgnoreMissingSubmodules].(bool); exists {
		cache.IgnoreMissingSubmodules = val
	}
}

func (cache *BlobCache) Initialize(repository *git.Repository) {
	cache.repository = repository
	cache.cache = map[plumbing.Hash]*object.Blob{}
}

func (self *BlobCache) Consume(deps map[string]interface{}) (map[string]interface{}, error) {
	commit := deps["commit"].(*object.Commit)
	changes := deps[DependencyTreeChanges].(object.Changes)
	cache := map[plumbing.Hash]*object.Blob{}
	newCache := map[plumbing.Hash]*object.Blob{}
	for _, change := range changes {
		action, err := change.Action()
		if err != nil {
			fmt.Fprintf(os.Stderr, "no action in %s\n", change.To.TreeEntry.Hash)
			return nil, err
		}
		var exists bool
		var blob *object.Blob
		switch action {
		case merkletrie.Insert:
			blob, err = self.getBlob(&change.To, commit.File)
			if err != nil {
				fmt.Fprintf(os.Stderr, "file to %s %s\n", change.To.Name, change.To.TreeEntry.Hash)
			} else {
				cache[change.To.TreeEntry.Hash] = blob
				newCache[change.To.TreeEntry.Hash] = blob
			}
		case merkletrie.Delete:
			cache[change.From.TreeEntry.Hash], exists = self.cache[change.From.TreeEntry.Hash]
			if !exists {
				cache[change.From.TreeEntry.Hash], err = self.getBlob(&change.From, commit.File)
				if err != nil {
					if err.Error() != plumbing.ErrObjectNotFound.Error() {
						fmt.Fprintf(os.Stderr, "file from %s %s\n", change.From.Name,
							change.From.TreeEntry.Hash)
					} else {
						cache[change.From.TreeEntry.Hash], err = createDummyBlob(
							change.From.TreeEntry.Hash)
					}
				}
			}
		case merkletrie.Modify:
			blob, err = self.getBlob(&change.To, commit.File)
			if err != nil {
				fmt.Fprintf(os.Stderr, "file to %s\n", change.To.Name)
			} else {
				cache[change.To.TreeEntry.Hash] = blob
				newCache[change.To.TreeEntry.Hash] = blob
			}
			cache[change.From.TreeEntry.Hash], exists = self.cache[change.From.TreeEntry.Hash]
			if !exists {
				cache[change.From.TreeEntry.Hash], err = self.getBlob(&change.From, commit.File)
				if err != nil {
					fmt.Fprintf(os.Stderr, "file from %s\n", change.From.Name)
				}
			}
		}
		if err != nil {
			return nil, err
		}
	}
	self.cache = newCache
	return map[string]interface{}{DependencyBlobCache: cache}, nil
}

// The definition of a function which loads a git file by the specified path.
// The state can be arbitrary though here it always corresponds to the currently processed
// commit.
type FileGetter func(path string) (*object.File, error)

// Returns the blob which corresponds to the specified ChangeEntry.
func (cache *BlobCache) getBlob(entry *object.ChangeEntry, fileGetter FileGetter) (
	*object.Blob, error) {
	blob, err := cache.repository.BlobObject(entry.TreeEntry.Hash)
	if err != nil {
		if err.Error() != plumbing.ErrObjectNotFound.Error() {
			fmt.Fprintf(os.Stderr, "getBlob(%s)\n", entry.TreeEntry.Hash.String())
			return nil, err
		}
		if entry.TreeEntry.Mode != 0160000 {
			// this is not a submodule
			return nil, err
		} else if cache.IgnoreMissingSubmodules {
			return createDummyBlob(entry.TreeEntry.Hash)
		}
		file, err_modules := fileGetter(".gitmodules")
		if err_modules != nil {
			return nil, err_modules
		}
		contents, err_modules := file.Contents()
		if err_modules != nil {
			return nil, err_modules
		}
		modules := config.NewModules()
		err_modules = modules.Unmarshal([]byte(contents))
		if err_modules != nil {
			return nil, err_modules
		}
		_, exists := modules.Submodules[entry.Name]
		if exists {
			// we found that this is a submodule
			return createDummyBlob(entry.TreeEntry.Hash)
		}
		return nil, err
	}
	return blob, nil
}

func init() {
	Registry.Register(&BlobCache{})
}
