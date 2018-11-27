// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb.proto

It has these top-level messages:
	Metadata
	BurndownSparseMatrixRow
	BurndownSparseMatrix
	BurndownAnalysisResults
	CompressedSparseRowMatrix
	Couples
	TouchedFiles
	CouplesAnalysisResults
	UASTChange
	UASTChangesSaverResults
	ShotnessRecord
	ShotnessAnalysisResults
	FileHistory
	FileHistoryResultMessage
	DevDay
	DayDevs
	DevsAnalysisResults
	Sentiment
	CommentSentimentResults
	AnalysisResults
*/
package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Metadata struct {
	// this format is versioned
	Version int32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// git hash of the revision from which Hercules is built
	Hash string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	// repository's name
	Repository string `protobuf:"bytes,3,opt,name=repository,proto3" json:"repository,omitempty"`
	// UNIX timestamp of the first analysed commit
	BeginUnixTime int64 `protobuf:"varint,4,opt,name=begin_unix_time,json=beginUnixTime,proto3" json:"begin_unix_time,omitempty"`
	// UNIX timestamp of the last analysed commit
	EndUnixTime int64 `protobuf:"varint,5,opt,name=end_unix_time,json=endUnixTime,proto3" json:"end_unix_time,omitempty"`
	// number of processed commits
	Commits int32 `protobuf:"varint,6,opt,name=commits,proto3" json:"commits,omitempty"`
	// duration of the analysis in milliseconds
	RunTime int64 `protobuf:"varint,7,opt,name=run_time,json=runTime,proto3" json:"run_time,omitempty"`
	// time taken by each pipeline item in seconds
	RunTimePerItem map[string]float64 `protobuf:"bytes,8,rep,name=run_time_per_item,json=runTimePerItem" json:"run_time_per_item,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{0} }

func (m *Metadata) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Metadata) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Metadata) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *Metadata) GetBeginUnixTime() int64 {
	if m != nil {
		return m.BeginUnixTime
	}
	return 0
}

func (m *Metadata) GetEndUnixTime() int64 {
	if m != nil {
		return m.EndUnixTime
	}
	return 0
}

func (m *Metadata) GetCommits() int32 {
	if m != nil {
		return m.Commits
	}
	return 0
}

func (m *Metadata) GetRunTime() int64 {
	if m != nil {
		return m.RunTime
	}
	return 0
}

func (m *Metadata) GetRunTimePerItem() map[string]float64 {
	if m != nil {
		return m.RunTimePerItem
	}
	return nil
}

type BurndownSparseMatrixRow struct {
	// the first `len(column)` elements are stored,
	// the rest `number_of_columns - len(column)` values are zeros
	Columns []uint32 `protobuf:"varint,1,rep,packed,name=columns" json:"columns,omitempty"`
}

func (m *BurndownSparseMatrixRow) Reset()                    { *m = BurndownSparseMatrixRow{} }
func (m *BurndownSparseMatrixRow) String() string            { return proto.CompactTextString(m) }
func (*BurndownSparseMatrixRow) ProtoMessage()               {}
func (*BurndownSparseMatrixRow) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{1} }

func (m *BurndownSparseMatrixRow) GetColumns() []uint32 {
	if m != nil {
		return m.Columns
	}
	return nil
}

type BurndownSparseMatrix struct {
	Name            string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	NumberOfRows    int32  `protobuf:"varint,2,opt,name=number_of_rows,json=numberOfRows,proto3" json:"number_of_rows,omitempty"`
	NumberOfColumns int32  `protobuf:"varint,3,opt,name=number_of_columns,json=numberOfColumns,proto3" json:"number_of_columns,omitempty"`
	// `len(row)` matches `number_of_rows`
	Rows []*BurndownSparseMatrixRow `protobuf:"bytes,4,rep,name=rows" json:"rows,omitempty"`
}

func (m *BurndownSparseMatrix) Reset()                    { *m = BurndownSparseMatrix{} }
func (m *BurndownSparseMatrix) String() string            { return proto.CompactTextString(m) }
func (*BurndownSparseMatrix) ProtoMessage()               {}
func (*BurndownSparseMatrix) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{2} }

func (m *BurndownSparseMatrix) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BurndownSparseMatrix) GetNumberOfRows() int32 {
	if m != nil {
		return m.NumberOfRows
	}
	return 0
}

func (m *BurndownSparseMatrix) GetNumberOfColumns() int32 {
	if m != nil {
		return m.NumberOfColumns
	}
	return 0
}

func (m *BurndownSparseMatrix) GetRows() []*BurndownSparseMatrixRow {
	if m != nil {
		return m.Rows
	}
	return nil
}

type BurndownAnalysisResults struct {
	// how many days are in each band [burndown_project, burndown_file, burndown_developer]
	Granularity int32 `protobuf:"varint,1,opt,name=granularity,proto3" json:"granularity,omitempty"`
	// how frequently we measure the state of each band [burndown_project, burndown_file, burndown_developer]
	Sampling int32 `protobuf:"varint,2,opt,name=sampling,proto3" json:"sampling,omitempty"`
	// always exists
	Project *BurndownSparseMatrix `protobuf:"bytes,3,opt,name=project" json:"project,omitempty"`
	// this is included if `-burndown-files` was specified
	Files []*BurndownSparseMatrix `protobuf:"bytes,4,rep,name=files" json:"files,omitempty"`
	// these two are included if `-burndown-people` was specified
	People []*BurndownSparseMatrix `protobuf:"bytes,5,rep,name=people" json:"people,omitempty"`
	// rows and cols order correspond to `burndown_developer`
	PeopleInteraction *CompressedSparseRowMatrix `protobuf:"bytes,6,opt,name=people_interaction,json=peopleInteraction" json:"people_interaction,omitempty"`
}

func (m *BurndownAnalysisResults) Reset()                    { *m = BurndownAnalysisResults{} }
func (m *BurndownAnalysisResults) String() string            { return proto.CompactTextString(m) }
func (*BurndownAnalysisResults) ProtoMessage()               {}
func (*BurndownAnalysisResults) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{3} }

func (m *BurndownAnalysisResults) GetGranularity() int32 {
	if m != nil {
		return m.Granularity
	}
	return 0
}

func (m *BurndownAnalysisResults) GetSampling() int32 {
	if m != nil {
		return m.Sampling
	}
	return 0
}

func (m *BurndownAnalysisResults) GetProject() *BurndownSparseMatrix {
	if m != nil {
		return m.Project
	}
	return nil
}

func (m *BurndownAnalysisResults) GetFiles() []*BurndownSparseMatrix {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *BurndownAnalysisResults) GetPeople() []*BurndownSparseMatrix {
	if m != nil {
		return m.People
	}
	return nil
}

func (m *BurndownAnalysisResults) GetPeopleInteraction() *CompressedSparseRowMatrix {
	if m != nil {
		return m.PeopleInteraction
	}
	return nil
}

type CompressedSparseRowMatrix struct {
	NumberOfRows    int32 `protobuf:"varint,1,opt,name=number_of_rows,json=numberOfRows,proto3" json:"number_of_rows,omitempty"`
	NumberOfColumns int32 `protobuf:"varint,2,opt,name=number_of_columns,json=numberOfColumns,proto3" json:"number_of_columns,omitempty"`
	// https://en.wikipedia.org/wiki/Sparse_matrix#Compressed_sparse_row_.28CSR.2C_CRS_or_Yale_format.29
	Data    []int64 `protobuf:"varint,3,rep,packed,name=data" json:"data,omitempty"`
	Indices []int32 `protobuf:"varint,4,rep,packed,name=indices" json:"indices,omitempty"`
	Indptr  []int64 `protobuf:"varint,5,rep,packed,name=indptr" json:"indptr,omitempty"`
}

func (m *CompressedSparseRowMatrix) Reset()                    { *m = CompressedSparseRowMatrix{} }
func (m *CompressedSparseRowMatrix) String() string            { return proto.CompactTextString(m) }
func (*CompressedSparseRowMatrix) ProtoMessage()               {}
func (*CompressedSparseRowMatrix) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{4} }

func (m *CompressedSparseRowMatrix) GetNumberOfRows() int32 {
	if m != nil {
		return m.NumberOfRows
	}
	return 0
}

func (m *CompressedSparseRowMatrix) GetNumberOfColumns() int32 {
	if m != nil {
		return m.NumberOfColumns
	}
	return 0
}

func (m *CompressedSparseRowMatrix) GetData() []int64 {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *CompressedSparseRowMatrix) GetIndices() []int32 {
	if m != nil {
		return m.Indices
	}
	return nil
}

func (m *CompressedSparseRowMatrix) GetIndptr() []int64 {
	if m != nil {
		return m.Indptr
	}
	return nil
}

type Couples struct {
	// name of each `matrix`'s row and column
	Index []string `protobuf:"bytes,1,rep,name=index" json:"index,omitempty"`
	// is always square
	Matrix *CompressedSparseRowMatrix `protobuf:"bytes,2,opt,name=matrix" json:"matrix,omitempty"`
}

func (m *Couples) Reset()                    { *m = Couples{} }
func (m *Couples) String() string            { return proto.CompactTextString(m) }
func (*Couples) ProtoMessage()               {}
func (*Couples) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{5} }

func (m *Couples) GetIndex() []string {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *Couples) GetMatrix() *CompressedSparseRowMatrix {
	if m != nil {
		return m.Matrix
	}
	return nil
}

type TouchedFiles struct {
	Files []int32 `protobuf:"varint,1,rep,packed,name=files" json:"files,omitempty"`
}

func (m *TouchedFiles) Reset()                    { *m = TouchedFiles{} }
func (m *TouchedFiles) String() string            { return proto.CompactTextString(m) }
func (*TouchedFiles) ProtoMessage()               {}
func (*TouchedFiles) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{6} }

func (m *TouchedFiles) GetFiles() []int32 {
	if m != nil {
		return m.Files
	}
	return nil
}

type CouplesAnalysisResults struct {
	FileCouples   *Couples `protobuf:"bytes,6,opt,name=file_couples,json=fileCouples" json:"file_couples,omitempty"`
	PeopleCouples *Couples `protobuf:"bytes,7,opt,name=people_couples,json=peopleCouples" json:"people_couples,omitempty"`
	// order corresponds to `people_couples::index`
	PeopleFiles []*TouchedFiles `protobuf:"bytes,8,rep,name=people_files,json=peopleFiles" json:"people_files,omitempty"`
}

func (m *CouplesAnalysisResults) Reset()                    { *m = CouplesAnalysisResults{} }
func (m *CouplesAnalysisResults) String() string            { return proto.CompactTextString(m) }
func (*CouplesAnalysisResults) ProtoMessage()               {}
func (*CouplesAnalysisResults) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{7} }

func (m *CouplesAnalysisResults) GetFileCouples() *Couples {
	if m != nil {
		return m.FileCouples
	}
	return nil
}

func (m *CouplesAnalysisResults) GetPeopleCouples() *Couples {
	if m != nil {
		return m.PeopleCouples
	}
	return nil
}

func (m *CouplesAnalysisResults) GetPeopleFiles() []*TouchedFiles {
	if m != nil {
		return m.PeopleFiles
	}
	return nil
}

type UASTChange struct {
	FileName   string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	SrcBefore  string `protobuf:"bytes,2,opt,name=src_before,json=srcBefore,proto3" json:"src_before,omitempty"`
	SrcAfter   string `protobuf:"bytes,3,opt,name=src_after,json=srcAfter,proto3" json:"src_after,omitempty"`
	UastBefore string `protobuf:"bytes,4,opt,name=uast_before,json=uastBefore,proto3" json:"uast_before,omitempty"`
	UastAfter  string `protobuf:"bytes,5,opt,name=uast_after,json=uastAfter,proto3" json:"uast_after,omitempty"`
}

func (m *UASTChange) Reset()                    { *m = UASTChange{} }
func (m *UASTChange) String() string            { return proto.CompactTextString(m) }
func (*UASTChange) ProtoMessage()               {}
func (*UASTChange) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{8} }

func (m *UASTChange) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *UASTChange) GetSrcBefore() string {
	if m != nil {
		return m.SrcBefore
	}
	return ""
}

func (m *UASTChange) GetSrcAfter() string {
	if m != nil {
		return m.SrcAfter
	}
	return ""
}

func (m *UASTChange) GetUastBefore() string {
	if m != nil {
		return m.UastBefore
	}
	return ""
}

func (m *UASTChange) GetUastAfter() string {
	if m != nil {
		return m.UastAfter
	}
	return ""
}

type UASTChangesSaverResults struct {
	Changes []*UASTChange `protobuf:"bytes,1,rep,name=changes" json:"changes,omitempty"`
}

func (m *UASTChangesSaverResults) Reset()                    { *m = UASTChangesSaverResults{} }
func (m *UASTChangesSaverResults) String() string            { return proto.CompactTextString(m) }
func (*UASTChangesSaverResults) ProtoMessage()               {}
func (*UASTChangesSaverResults) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{9} }

func (m *UASTChangesSaverResults) GetChanges() []*UASTChange {
	if m != nil {
		return m.Changes
	}
	return nil
}

type ShotnessRecord struct {
	InternalRole string          `protobuf:"bytes,1,opt,name=internal_role,json=internalRole,proto3" json:"internal_role,omitempty"`
	Roles        []int32         `protobuf:"varint,2,rep,packed,name=roles" json:"roles,omitempty"`
	Name         string          `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	File         string          `protobuf:"bytes,4,opt,name=file,proto3" json:"file,omitempty"`
	Counters     map[int32]int32 `protobuf:"bytes,5,rep,name=counters" json:"counters,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (m *ShotnessRecord) Reset()                    { *m = ShotnessRecord{} }
func (m *ShotnessRecord) String() string            { return proto.CompactTextString(m) }
func (*ShotnessRecord) ProtoMessage()               {}
func (*ShotnessRecord) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{10} }

func (m *ShotnessRecord) GetInternalRole() string {
	if m != nil {
		return m.InternalRole
	}
	return ""
}

func (m *ShotnessRecord) GetRoles() []int32 {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *ShotnessRecord) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ShotnessRecord) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *ShotnessRecord) GetCounters() map[int32]int32 {
	if m != nil {
		return m.Counters
	}
	return nil
}

type ShotnessAnalysisResults struct {
	Records []*ShotnessRecord `protobuf:"bytes,1,rep,name=records" json:"records,omitempty"`
}

func (m *ShotnessAnalysisResults) Reset()                    { *m = ShotnessAnalysisResults{} }
func (m *ShotnessAnalysisResults) String() string            { return proto.CompactTextString(m) }
func (*ShotnessAnalysisResults) ProtoMessage()               {}
func (*ShotnessAnalysisResults) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{11} }

func (m *ShotnessAnalysisResults) GetRecords() []*ShotnessRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

type FileHistory struct {
	Commits []string `protobuf:"bytes,1,rep,name=commits" json:"commits,omitempty"`
}

func (m *FileHistory) Reset()                    { *m = FileHistory{} }
func (m *FileHistory) String() string            { return proto.CompactTextString(m) }
func (*FileHistory) ProtoMessage()               {}
func (*FileHistory) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{12} }

func (m *FileHistory) GetCommits() []string {
	if m != nil {
		return m.Commits
	}
	return nil
}

type FileHistoryResultMessage struct {
	Files map[string]*FileHistory `protobuf:"bytes,1,rep,name=files" json:"files,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *FileHistoryResultMessage) Reset()                    { *m = FileHistoryResultMessage{} }
func (m *FileHistoryResultMessage) String() string            { return proto.CompactTextString(m) }
func (*FileHistoryResultMessage) ProtoMessage()               {}
func (*FileHistoryResultMessage) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{13} }

func (m *FileHistoryResultMessage) GetFiles() map[string]*FileHistory {
	if m != nil {
		return m.Files
	}
	return nil
}

type DevDay struct {
	Commits int32 `protobuf:"varint,1,opt,name=commits,proto3" json:"commits,omitempty"`
	Added   int32 `protobuf:"varint,2,opt,name=added,proto3" json:"added,omitempty"`
	Removed int32 `protobuf:"varint,3,opt,name=removed,proto3" json:"removed,omitempty"`
	Changed int32 `protobuf:"varint,4,opt,name=changed,proto3" json:"changed,omitempty"`
}

func (m *DevDay) Reset()                    { *m = DevDay{} }
func (m *DevDay) String() string            { return proto.CompactTextString(m) }
func (*DevDay) ProtoMessage()               {}
func (*DevDay) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{14} }

func (m *DevDay) GetCommits() int32 {
	if m != nil {
		return m.Commits
	}
	return 0
}

func (m *DevDay) GetAdded() int32 {
	if m != nil {
		return m.Added
	}
	return 0
}

func (m *DevDay) GetRemoved() int32 {
	if m != nil {
		return m.Removed
	}
	return 0
}

func (m *DevDay) GetChanged() int32 {
	if m != nil {
		return m.Changed
	}
	return 0
}

type DayDevs struct {
	Devs map[int32]*DevDay `protobuf:"bytes,1,rep,name=devs" json:"devs,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *DayDevs) Reset()                    { *m = DayDevs{} }
func (m *DayDevs) String() string            { return proto.CompactTextString(m) }
func (*DayDevs) ProtoMessage()               {}
func (*DayDevs) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{15} }

func (m *DayDevs) GetDevs() map[int32]*DevDay {
	if m != nil {
		return m.Devs
	}
	return nil
}

type DevsAnalysisResults struct {
	Days     map[int32]*DayDevs `protobuf:"bytes,1,rep,name=days" json:"days,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	DevIndex []string           `protobuf:"bytes,2,rep,name=dev_index,json=devIndex" json:"dev_index,omitempty"`
}

func (m *DevsAnalysisResults) Reset()                    { *m = DevsAnalysisResults{} }
func (m *DevsAnalysisResults) String() string            { return proto.CompactTextString(m) }
func (*DevsAnalysisResults) ProtoMessage()               {}
func (*DevsAnalysisResults) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{16} }

func (m *DevsAnalysisResults) GetDays() map[int32]*DayDevs {
	if m != nil {
		return m.Days
	}
	return nil
}

func (m *DevsAnalysisResults) GetDevIndex() []string {
	if m != nil {
		return m.DevIndex
	}
	return nil
}

type Sentiment struct {
	Value    float32  `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
	Comments []string `protobuf:"bytes,2,rep,name=comments" json:"comments,omitempty"`
	Commits  []string `protobuf:"bytes,3,rep,name=commits" json:"commits,omitempty"`
}

func (m *Sentiment) Reset()                    { *m = Sentiment{} }
func (m *Sentiment) String() string            { return proto.CompactTextString(m) }
func (*Sentiment) ProtoMessage()               {}
func (*Sentiment) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{17} }

func (m *Sentiment) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Sentiment) GetComments() []string {
	if m != nil {
		return m.Comments
	}
	return nil
}

func (m *Sentiment) GetCommits() []string {
	if m != nil {
		return m.Commits
	}
	return nil
}

type CommentSentimentResults struct {
	SentimentByDay map[int32]*Sentiment `protobuf:"bytes,1,rep,name=sentiment_by_day,json=sentimentByDay" json:"sentiment_by_day,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CommentSentimentResults) Reset()                    { *m = CommentSentimentResults{} }
func (m *CommentSentimentResults) String() string            { return proto.CompactTextString(m) }
func (*CommentSentimentResults) ProtoMessage()               {}
func (*CommentSentimentResults) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{18} }

func (m *CommentSentimentResults) GetSentimentByDay() map[int32]*Sentiment {
	if m != nil {
		return m.SentimentByDay
	}
	return nil
}

type AnalysisResults struct {
	Header *Metadata `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	// the mapped values are dynamic messages which require the second parsing pass.
	Contents map[string][]byte `protobuf:"bytes,2,rep,name=contents" json:"contents,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *AnalysisResults) Reset()                    { *m = AnalysisResults{} }
func (m *AnalysisResults) String() string            { return proto.CompactTextString(m) }
func (*AnalysisResults) ProtoMessage()               {}
func (*AnalysisResults) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{19} }

func (m *AnalysisResults) GetHeader() *Metadata {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *AnalysisResults) GetContents() map[string][]byte {
	if m != nil {
		return m.Contents
	}
	return nil
}

func init() {
	proto.RegisterType((*Metadata)(nil), "Metadata")
	proto.RegisterType((*BurndownSparseMatrixRow)(nil), "BurndownSparseMatrixRow")
	proto.RegisterType((*BurndownSparseMatrix)(nil), "BurndownSparseMatrix")
	proto.RegisterType((*BurndownAnalysisResults)(nil), "BurndownAnalysisResults")
	proto.RegisterType((*CompressedSparseRowMatrix)(nil), "CompressedSparseRowMatrix")
	proto.RegisterType((*Couples)(nil), "Couples")
	proto.RegisterType((*TouchedFiles)(nil), "TouchedFiles")
	proto.RegisterType((*CouplesAnalysisResults)(nil), "CouplesAnalysisResults")
	proto.RegisterType((*UASTChange)(nil), "UASTChange")
	proto.RegisterType((*UASTChangesSaverResults)(nil), "UASTChangesSaverResults")
	proto.RegisterType((*ShotnessRecord)(nil), "ShotnessRecord")
	proto.RegisterType((*ShotnessAnalysisResults)(nil), "ShotnessAnalysisResults")
	proto.RegisterType((*FileHistory)(nil), "FileHistory")
	proto.RegisterType((*FileHistoryResultMessage)(nil), "FileHistoryResultMessage")
	proto.RegisterType((*DevDay)(nil), "DevDay")
	proto.RegisterType((*DayDevs)(nil), "DayDevs")
	proto.RegisterType((*DevsAnalysisResults)(nil), "DevsAnalysisResults")
	proto.RegisterType((*Sentiment)(nil), "Sentiment")
	proto.RegisterType((*CommentSentimentResults)(nil), "CommentSentimentResults")
	proto.RegisterType((*AnalysisResults)(nil), "AnalysisResults")
}

func init() { proto.RegisterFile("pb.proto", fileDescriptorPb) }

var fileDescriptorPb = []byte{
	// 1232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0xcd, 0x72, 0x1b, 0x45,
	0x10, 0xae, 0xd5, 0xbf, 0x7a, 0x25, 0x3b, 0x99, 0x98, 0x58, 0x11, 0xe5, 0x20, 0x96, 0x10, 0x0c,
	0x09, 0x1b, 0x4a, 0xb9, 0x40, 0xb8, 0xc4, 0xb1, 0x49, 0xc5, 0x07, 0x03, 0x35, 0x4e, 0xe0, 0xb8,
	0x35, 0xd6, 0xb6, 0xed, 0x05, 0xed, 0xac, 0x6a, 0x66, 0x57, 0xb6, 0x5e, 0x86, 0x1b, 0x55, 0x40,
	0x15, 0xc5, 0x81, 0x17, 0xe0, 0x69, 0xb8, 0xf0, 0x12, 0xd4, 0xfc, 0x49, 0x2b, 0xd5, 0x3a, 0x70,
	0xd9, 0xda, 0xee, 0xfe, 0x7a, 0xe6, 0xeb, 0x9f, 0xe9, 0x19, 0xe8, 0xcc, 0xce, 0xc2, 0x99, 0xc8,
	0xf2, 0x2c, 0xf8, 0xbb, 0x06, 0x9d, 0x13, 0xcc, 0x59, 0xcc, 0x72, 0x46, 0x06, 0xd0, 0x9e, 0xa3,
	0x90, 0x49, 0xc6, 0x07, 0xde, 0xc8, 0xdb, 0x6f, 0x52, 0x27, 0x12, 0x02, 0x8d, 0x4b, 0x26, 0x2f,
	0x07, 0xb5, 0x91, 0xb7, 0xdf, 0xa5, 0xfa, 0x9f, 0xdc, 0x07, 0x10, 0x38, 0xcb, 0x64, 0x92, 0x67,
	0x62, 0x31, 0xa8, 0x6b, 0x4b, 0x49, 0x43, 0x1e, 0xc2, 0xf6, 0x19, 0x5e, 0x24, 0x3c, 0x2a, 0x78,
	0x72, 0x1d, 0xe5, 0x49, 0x8a, 0x83, 0xc6, 0xc8, 0xdb, 0xaf, 0xd3, 0xbe, 0x56, 0xbf, 0xe1, 0xc9,
	0xf5, 0xeb, 0x24, 0x45, 0x12, 0x40, 0x1f, 0x79, 0x5c, 0x42, 0x35, 0x35, 0xca, 0x47, 0x1e, 0x2f,
	0x31, 0x03, 0x68, 0x4f, 0xb2, 0x34, 0x4d, 0x72, 0x39, 0x68, 0x19, 0x66, 0x56, 0x24, 0xf7, 0xa0,
	0x23, 0x0a, 0x6e, 0x1c, 0xdb, 0xda, 0xb1, 0x2d, 0x0a, 0xae, 0x9d, 0x5e, 0xc1, 0x6d, 0x67, 0x8a,
	0x66, 0x28, 0xa2, 0x24, 0xc7, 0x74, 0xd0, 0x19, 0xd5, 0xf7, 0xfd, 0xf1, 0x5e, 0xe8, 0x82, 0x0e,
	0xa9, 0x41, 0x7f, 0x8b, 0xe2, 0x38, 0xc7, 0xf4, 0x2b, 0x9e, 0x8b, 0x05, 0xdd, 0x12, 0x6b, 0xca,
	0xe1, 0x01, 0xdc, 0xa9, 0x80, 0x91, 0x5b, 0x50, 0xff, 0x11, 0x17, 0x3a, 0x57, 0x5d, 0xaa, 0x7e,
	0xc9, 0x0e, 0x34, 0xe7, 0x6c, 0x5a, 0xa0, 0x4e, 0x94, 0x47, 0x8d, 0xf0, 0xac, 0xf6, 0xb9, 0x17,
	0x3c, 0x85, 0xdd, 0x17, 0x85, 0xe0, 0x71, 0x76, 0xc5, 0x4f, 0x67, 0x4c, 0x48, 0x3c, 0x61, 0xb9,
	0x48, 0xae, 0x69, 0x76, 0x65, 0x82, 0x9b, 0x16, 0x29, 0x97, 0x03, 0x6f, 0x54, 0xdf, 0xef, 0x53,
	0x27, 0x06, 0xbf, 0x79, 0xb0, 0x53, 0xe5, 0xa5, 0xea, 0xc1, 0x59, 0x8a, 0x76, 0x6b, 0xfd, 0x4f,
	0x1e, 0xc0, 0x16, 0x2f, 0xd2, 0x33, 0x14, 0x51, 0x76, 0x1e, 0x89, 0xec, 0x4a, 0x6a, 0x12, 0x4d,
	0xda, 0x33, 0xda, 0x6f, 0xce, 0x69, 0x76, 0x25, 0xc9, 0x27, 0x70, 0x7b, 0x85, 0x72, 0xdb, 0xd6,
	0x35, 0x70, 0xdb, 0x01, 0x0f, 0x8d, 0x9a, 0x3c, 0x86, 0x86, 0x5e, 0xa7, 0xa1, 0x73, 0x36, 0x08,
	0x6f, 0x08, 0x80, 0x6a, 0x54, 0xf0, 0x47, 0x6d, 0x15, 0xe2, 0x01, 0x67, 0xd3, 0x85, 0x4c, 0x24,
	0x45, 0x59, 0x4c, 0x73, 0x49, 0x46, 0xe0, 0x5f, 0x08, 0xc6, 0x8b, 0x29, 0x13, 0x49, 0xbe, 0xb0,
	0xdd, 0x55, 0x56, 0x91, 0x21, 0x74, 0x24, 0x4b, 0x67, 0xd3, 0x84, 0x5f, 0x58, 0xde, 0x4b, 0x99,
	0x3c, 0x81, 0xf6, 0x4c, 0x64, 0x3f, 0xe0, 0x24, 0xd7, 0x4c, 0xfd, 0xf1, 0x3b, 0xd5, 0x54, 0x1c,
	0x8a, 0x3c, 0x82, 0xe6, 0x79, 0x32, 0x45, 0xc7, 0xfc, 0x06, 0xb8, 0xc1, 0x90, 0x4f, 0xa1, 0x35,
	0xc3, 0x6c, 0x36, 0x55, 0x8d, 0xf7, 0x16, 0xb4, 0x05, 0x91, 0x63, 0x20, 0xe6, 0x2f, 0x4a, 0x78,
	0x8e, 0x82, 0x4d, 0x72, 0x75, 0x5e, 0x5a, 0x9a, 0xd7, 0x30, 0x3c, 0xcc, 0xd2, 0x99, 0x40, 0x29,
	0x31, 0x36, 0xce, 0x34, 0xbb, 0xb2, 0xfe, 0xb7, 0x8d, 0xd7, 0xf1, 0xca, 0x29, 0xf8, 0xd3, 0x83,
	0x7b, 0x37, 0x3a, 0x54, 0xd4, 0xd3, 0xfb, 0xbf, 0xf5, 0xac, 0x55, 0xd7, 0x93, 0x40, 0x43, 0xb5,
	0xfc, 0xa0, 0x3e, 0xaa, 0xef, 0xd7, 0x69, 0xc3, 0x9d, 0xf9, 0x84, 0xc7, 0xc9, 0xc4, 0x26, 0xab,
	0x49, 0x9d, 0x48, 0xee, 0x42, 0x2b, 0xe1, 0xf1, 0x2c, 0x17, 0x3a, 0x2f, 0x75, 0x6a, 0xa5, 0xe0,
	0x14, 0xda, 0x87, 0x59, 0x31, 0x53, 0xa9, 0xdb, 0x81, 0x66, 0xc2, 0x63, 0xbc, 0xd6, 0x7d, 0xdb,
	0xa5, 0x46, 0x20, 0x63, 0x68, 0xa5, 0x3a, 0x04, 0xcd, 0xe3, 0xed, 0x59, 0xb1, 0xc8, 0xe0, 0x01,
	0xf4, 0x5e, 0x67, 0xc5, 0xe4, 0x12, 0xe3, 0x97, 0x89, 0x5d, 0xd9, 0x54, 0xd0, 0xd3, 0xa4, 0x8c,
	0x10, 0xfc, 0xe2, 0xc1, 0x5d, 0xbb, 0xf7, 0x66, 0x87, 0x3d, 0x82, 0x9e, 0xc2, 0x44, 0x13, 0x63,
	0xb6, 0x05, 0xe9, 0x84, 0x16, 0x4e, 0x7d, 0x65, 0x75, 0xbc, 0x9f, 0xc0, 0x96, 0xad, 0xa1, 0x83,
	0xb7, 0x37, 0xe0, 0x7d, 0x63, 0x77, 0x0e, 0x9f, 0x41, 0xcf, 0x3a, 0x18, 0x56, 0x66, 0x8a, 0xf4,
	0xc3, 0x32, 0x67, 0xea, 0x1b, 0x88, 0x16, 0x82, 0x9f, 0x3d, 0x80, 0x37, 0x07, 0xa7, 0xaf, 0x0f,
	0x2f, 0x19, 0xbf, 0x40, 0xf2, 0x2e, 0x74, 0x35, 0xbd, 0xd2, 0xa9, 0xed, 0x28, 0xc5, 0xd7, 0xea,
	0xe4, 0xee, 0x01, 0x48, 0x31, 0x89, 0xce, 0xf0, 0x3c, 0x13, 0x68, 0x67, 0x6c, 0x57, 0x8a, 0xc9,
	0x0b, 0xad, 0x50, 0xbe, 0xca, 0xcc, 0xce, 0x73, 0x14, 0x76, 0xce, 0x76, 0xa4, 0x98, 0x1c, 0x28,
	0x99, 0xbc, 0x07, 0x7e, 0xc1, 0x64, 0xee, 0x9c, 0x1b, 0x66, 0x0c, 0x2b, 0x95, 0xf5, 0xde, 0x03,
	0x2d, 0x59, 0xf7, 0xa6, 0x59, 0x5c, 0x69, 0xb4, 0x7f, 0xf0, 0x1c, 0x76, 0x57, 0x34, 0xe5, 0x29,
	0x9b, 0xa3, 0x70, 0x29, 0xfd, 0x10, 0xda, 0x13, 0xa3, 0xd6, 0x55, 0xf0, 0xc7, 0x7e, 0xb8, 0x82,
	0x52, 0x67, 0x0b, 0xfe, 0xf1, 0x60, 0xeb, 0xf4, 0x32, 0xcb, 0x39, 0x4a, 0x49, 0x71, 0x92, 0x89,
	0x98, 0x7c, 0x00, 0x7d, 0x7d, 0x38, 0x38, 0x9b, 0x46, 0x22, 0x9b, 0xba, 0x88, 0x7b, 0x4e, 0x49,
	0xb3, 0x29, 0xaa, 0x12, 0x2b, 0x9b, 0xea, 0x56, 0x5d, 0x62, 0x2d, 0x2c, 0x27, 0x5b, 0xbd, 0x34,
	0xd9, 0x08, 0x34, 0x54, 0xae, 0x6c, 0x70, 0xfa, 0x9f, 0x7c, 0x01, 0x9d, 0x49, 0x56, 0xa8, 0xf5,
	0xa4, 0x3d, 0xb7, 0x7b, 0xe1, 0x3a, 0x0b, 0x55, 0x4b, 0x6d, 0x37, 0x33, 0x7d, 0x09, 0x1f, 0x7e,
	0x09, 0xfd, 0x35, 0x53, 0x79, 0x8e, 0x37, 0x2b, 0xe6, 0x78, 0xb3, 0x3c, 0xc7, 0x8f, 0x60, 0xd7,
	0x6d, 0xb3, 0xd9, 0x82, 0x1f, 0x43, 0x5b, 0xe8, 0x9d, 0x5d, 0xbe, 0xb6, 0x37, 0x18, 0x51, 0x67,
	0x0f, 0x3e, 0x02, 0x5f, 0xb5, 0xc9, 0xab, 0x44, 0xea, 0xab, 0xb2, 0x74, 0xbd, 0x99, 0x93, 0xe4,
	0xc4, 0xe0, 0x27, 0x0f, 0x06, 0x25, 0xa4, 0xd9, 0xea, 0x04, 0xa5, 0x64, 0x17, 0x48, 0x9e, 0x95,
	0x0f, 0x89, 0x3f, 0x7e, 0x10, 0xde, 0x84, 0xd4, 0x06, 0x9b, 0x07, 0xe3, 0x32, 0x7c, 0x09, 0xb0,
	0x52, 0x56, 0xdc, 0x64, 0x41, 0x39, 0x03, 0xfe, 0xb8, 0xb7, 0xb6, 0x76, 0x29, 0x1f, 0x53, 0x68,
	0x1d, 0xe1, 0xfc, 0x88, 0x6d, 0x04, 0xb1, 0x76, 0x47, 0xef, 0x40, 0x93, 0xc5, 0x31, 0xc6, 0x2e,
	0x9b, 0x5a, 0x50, 0x78, 0x81, 0x69, 0x36, 0xc7, 0xd8, 0xde, 0x3f, 0x4e, 0xd4, 0x2b, 0xe9, 0xe6,
	0x8a, 0x75, 0xc9, 0x9b, 0xae, 0xd7, 0xe2, 0x40, 0x42, 0xfb, 0x88, 0x2d, 0x8e, 0x70, 0x2e, 0xc9,
	0x43, 0x68, 0xc4, 0x38, 0x77, 0xb1, 0x93, 0xd0, 0xea, 0x43, 0xf5, 0x31, 0x91, 0x6a, 0xfb, 0xf0,
	0x39, 0x74, 0x97, 0xaa, 0x8a, 0x4a, 0xef, 0xad, 0xc7, 0xd9, 0x0e, 0x4d, 0x34, 0xe5, 0x10, 0x7f,
	0xf5, 0xe0, 0x8e, 0x5a, 0x62, 0xb3, 0xde, 0x63, 0x35, 0x4e, 0x17, 0x8e, 0xc1, 0xfd, 0xb0, 0x02,
	0xa3, 0x58, 0x2d, 0xd9, 0xb0, 0x85, 0x54, 0x67, 0x39, 0xc6, 0x79, 0x64, 0xa6, 0x66, 0x4d, 0xd7,
	0xba, 0x13, 0xe3, 0xfc, 0x58, 0xc9, 0xc3, 0x03, 0xe8, 0x2e, 0xf1, 0x15, 0x54, 0xef, 0xaf, 0x53,
	0xed, 0xb8, 0x90, 0xcb, 0x5c, 0xbf, 0x87, 0xee, 0x29, 0x72, 0xf5, 0xe4, 0xe1, 0xf9, 0xaa, 0x8b,
	0xd5, 0x22, 0x35, 0x0b, 0x53, 0x37, 0xad, 0x2a, 0x0c, 0xf2, 0x5c, 0x3a, 0x06, 0x4e, 0x2e, 0xd7,
	0xb0, 0xbe, 0xde, 0x88, 0x7f, 0x79, 0xb0, 0x7b, 0x68, 0x60, 0xcb, 0x0d, 0x5c, 0x22, 0xbe, 0x83,
	0x5b, 0xd2, 0xe9, 0xa2, 0xb3, 0x45, 0x14, 0xb3, 0x85, 0x4d, 0xca, 0xe3, 0xf0, 0x06, 0x9f, 0x70,
	0xa9, 0x78, 0xb1, 0x38, 0x62, 0x0b, 0xfb, 0xec, 0x92, 0x6b, 0xca, 0xe1, 0x09, 0xdc, 0xa9, 0x80,
	0x55, 0x64, 0x66, 0xb4, 0x9e, 0x19, 0x58, 0xad, 0x5e, 0xce, 0xcd, 0xef, 0x1e, 0x6c, 0x6f, 0xd6,
	0xf0, 0x7d, 0x68, 0x5d, 0x22, 0x8b, 0x51, 0xe8, 0xe5, 0xfc, 0x71, 0x77, 0xf9, 0x30, 0xa4, 0xd6,
	0x40, 0x9e, 0xa9, 0x7c, 0xf1, 0x7c, 0x99, 0x2f, 0x55, 0xea, 0xcd, 0x32, 0x1f, 0x5a, 0xc0, 0x72,
	0xd4, 0x18, 0xd1, 0x8c, 0x9a, 0x92, 0xe9, 0xbf, 0x9e, 0x8c, 0xbd, 0x12, 0xdf, 0xb3, 0x96, 0x7e,
	0xa2, 0x3f, 0xfd, 0x37, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x05, 0x8d, 0xb5, 0xae, 0x0b, 0x00, 0x00,
}
