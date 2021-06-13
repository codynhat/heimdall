// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/websecurityscanner/v1alpha/scan_run.proto

package websecurityscanner

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Types of ScanRun execution state.
type ScanRun_ExecutionState int32

const (
	// Represents an invalid state caused by internal server error. This value
	// should never be returned.
	ScanRun_EXECUTION_STATE_UNSPECIFIED ScanRun_ExecutionState = 0
	// The scan is waiting in the queue.
	ScanRun_QUEUED ScanRun_ExecutionState = 1
	// The scan is in progress.
	ScanRun_SCANNING ScanRun_ExecutionState = 2
	// The scan is either finished or stopped by user.
	ScanRun_FINISHED ScanRun_ExecutionState = 3
)

var ScanRun_ExecutionState_name = map[int32]string{
	0: "EXECUTION_STATE_UNSPECIFIED",
	1: "QUEUED",
	2: "SCANNING",
	3: "FINISHED",
}

var ScanRun_ExecutionState_value = map[string]int32{
	"EXECUTION_STATE_UNSPECIFIED": 0,
	"QUEUED":                      1,
	"SCANNING":                    2,
	"FINISHED":                    3,
}

func (x ScanRun_ExecutionState) String() string {
	return proto.EnumName(ScanRun_ExecutionState_name, int32(x))
}

func (ScanRun_ExecutionState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d1e91fc2897e59cf, []int{0, 0}
}

// Types of ScanRun result state.
type ScanRun_ResultState int32

const (
	// Default value. This value is returned when the ScanRun is not yet
	// finished.
	ScanRun_RESULT_STATE_UNSPECIFIED ScanRun_ResultState = 0
	// The scan finished without errors.
	ScanRun_SUCCESS ScanRun_ResultState = 1
	// The scan finished with errors.
	ScanRun_ERROR ScanRun_ResultState = 2
	// The scan was terminated by user.
	ScanRun_KILLED ScanRun_ResultState = 3
)

var ScanRun_ResultState_name = map[int32]string{
	0: "RESULT_STATE_UNSPECIFIED",
	1: "SUCCESS",
	2: "ERROR",
	3: "KILLED",
}

var ScanRun_ResultState_value = map[string]int32{
	"RESULT_STATE_UNSPECIFIED": 0,
	"SUCCESS":                  1,
	"ERROR":                    2,
	"KILLED":                   3,
}

func (x ScanRun_ResultState) String() string {
	return proto.EnumName(ScanRun_ResultState_name, int32(x))
}

func (ScanRun_ResultState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d1e91fc2897e59cf, []int{0, 1}
}

// A ScanRun is a output-only resource representing an actual run of the scan.
type ScanRun struct {
	// The resource name of the ScanRun. The name follows the format of
	// 'projects/{projectId}/scanConfigs/{scanConfigId}/scanRuns/{scanRunId}'.
	// The ScanRun IDs are generated by the system.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The execution state of the ScanRun.
	ExecutionState ScanRun_ExecutionState `protobuf:"varint,2,opt,name=execution_state,json=executionState,proto3,enum=google.cloud.websecurityscanner.v1alpha.ScanRun_ExecutionState" json:"execution_state,omitempty"`
	// The result state of the ScanRun. This field is only available after the
	// execution state reaches "FINISHED".
	ResultState ScanRun_ResultState `protobuf:"varint,3,opt,name=result_state,json=resultState,proto3,enum=google.cloud.websecurityscanner.v1alpha.ScanRun_ResultState" json:"result_state,omitempty"`
	// The time at which the ScanRun started.
	StartTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// The time at which the ScanRun reached termination state - that the ScanRun
	// is either finished or stopped by user.
	EndTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// The number of URLs crawled during this ScanRun. If the scan is in progress,
	// the value represents the number of URLs crawled up to now.
	UrlsCrawledCount int64 `protobuf:"varint,6,opt,name=urls_crawled_count,json=urlsCrawledCount,proto3" json:"urls_crawled_count,omitempty"`
	// The number of URLs tested during this ScanRun. If the scan is in progress,
	// the value represents the number of URLs tested up to now. The number of
	// URLs tested is usually larger than the number URLS crawled because
	// typically a crawled URL is tested with multiple test payloads.
	UrlsTestedCount int64 `protobuf:"varint,7,opt,name=urls_tested_count,json=urlsTestedCount,proto3" json:"urls_tested_count,omitempty"`
	// Whether the scan run has found any vulnerabilities.
	HasVulnerabilities bool `protobuf:"varint,8,opt,name=has_vulnerabilities,json=hasVulnerabilities,proto3" json:"has_vulnerabilities,omitempty"`
	// The percentage of total completion ranging from 0 to 100.
	// If the scan is in queue, the value is 0.
	// If the scan is running, the value ranges from 0 to 100.
	// If the scan is finished, the value is 100.
	ProgressPercent      int32    `protobuf:"varint,9,opt,name=progress_percent,json=progressPercent,proto3" json:"progress_percent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScanRun) Reset()         { *m = ScanRun{} }
func (m *ScanRun) String() string { return proto.CompactTextString(m) }
func (*ScanRun) ProtoMessage()    {}
func (*ScanRun) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e91fc2897e59cf, []int{0}
}

func (m *ScanRun) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanRun.Unmarshal(m, b)
}
func (m *ScanRun) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanRun.Marshal(b, m, deterministic)
}
func (m *ScanRun) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanRun.Merge(m, src)
}
func (m *ScanRun) XXX_Size() int {
	return xxx_messageInfo_ScanRun.Size(m)
}
func (m *ScanRun) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanRun.DiscardUnknown(m)
}

var xxx_messageInfo_ScanRun proto.InternalMessageInfo

func (m *ScanRun) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ScanRun) GetExecutionState() ScanRun_ExecutionState {
	if m != nil {
		return m.ExecutionState
	}
	return ScanRun_EXECUTION_STATE_UNSPECIFIED
}

func (m *ScanRun) GetResultState() ScanRun_ResultState {
	if m != nil {
		return m.ResultState
	}
	return ScanRun_RESULT_STATE_UNSPECIFIED
}

func (m *ScanRun) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *ScanRun) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *ScanRun) GetUrlsCrawledCount() int64 {
	if m != nil {
		return m.UrlsCrawledCount
	}
	return 0
}

func (m *ScanRun) GetUrlsTestedCount() int64 {
	if m != nil {
		return m.UrlsTestedCount
	}
	return 0
}

func (m *ScanRun) GetHasVulnerabilities() bool {
	if m != nil {
		return m.HasVulnerabilities
	}
	return false
}

func (m *ScanRun) GetProgressPercent() int32 {
	if m != nil {
		return m.ProgressPercent
	}
	return 0
}

func init() {
	proto.RegisterEnum("google.cloud.websecurityscanner.v1alpha.ScanRun_ExecutionState", ScanRun_ExecutionState_name, ScanRun_ExecutionState_value)
	proto.RegisterEnum("google.cloud.websecurityscanner.v1alpha.ScanRun_ResultState", ScanRun_ResultState_name, ScanRun_ResultState_value)
	proto.RegisterType((*ScanRun)(nil), "google.cloud.websecurityscanner.v1alpha.ScanRun")
}

func init() {
	proto.RegisterFile("google/cloud/websecurityscanner/v1alpha/scan_run.proto", fileDescriptor_d1e91fc2897e59cf)
}

var fileDescriptor_d1e91fc2897e59cf = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xed, 0x4e, 0xd4, 0x4c,
	0x18, 0x7d, 0xcb, 0xc7, 0x7e, 0xcc, 0x12, 0xe8, 0x3b, 0xfe, 0xa9, 0x68, 0xc2, 0x86, 0x3f, 0x2e,
	0x6a, 0xda, 0x88, 0xd1, 0xc4, 0x8f, 0x44, 0xa1, 0x14, 0x6d, 0x24, 0xcb, 0x3a, 0xdd, 0x35, 0xe2,
	0x9f, 0x66, 0x76, 0x76, 0xe8, 0xd6, 0xb4, 0x33, 0xcd, 0x7c, 0x80, 0x86, 0x70, 0x1f, 0x5e, 0x85,
	0x17, 0xe6, 0x55, 0x98, 0x4e, 0x5b, 0x91, 0x40, 0x02, 0xfe, 0xeb, 0x79, 0xce, 0x73, 0xce, 0x99,
	0xcc, 0xf3, 0x4c, 0xc1, 0xf3, 0x84, 0xf3, 0x24, 0xa3, 0x1e, 0xc9, 0xb8, 0x9e, 0x79, 0xa7, 0x74,
	0x2a, 0x29, 0xd1, 0x22, 0x55, 0xdf, 0x25, 0xc1, 0x8c, 0x51, 0xe1, 0x9d, 0x3c, 0xc1, 0x59, 0x31,
	0xc7, 0x5e, 0x89, 0x63, 0xa1, 0x99, 0x5b, 0x08, 0xae, 0x38, 0x7c, 0x50, 0xe9, 0x5c, 0xa3, 0x73,
	0xaf, 0xea, 0xdc, 0x5a, 0xb7, 0x7e, 0xb7, 0x0e, 0xc0, 0x45, 0xea, 0x09, 0x2a, 0xb9, 0x16, 0x84,
	0x56, 0x1e, 0xeb, 0x1b, 0x35, 0x65, 0xd0, 0x54, 0x1f, 0x7b, 0x2a, 0xcd, 0xa9, 0x54, 0x38, 0x2f,
	0xaa, 0x86, 0xcd, 0x9f, 0x2d, 0xd0, 0x8e, 0x08, 0x66, 0x48, 0x33, 0x08, 0xc1, 0x12, 0xc3, 0x39,
	0x75, 0xac, 0xbe, 0x35, 0xe8, 0x22, 0xf3, 0x0d, 0xe7, 0x60, 0x8d, 0x7e, 0xa3, 0x44, 0xab, 0x94,
	0xb3, 0x58, 0x2a, 0xac, 0xa8, 0xb3, 0xd0, 0xb7, 0x06, 0xab, 0xdb, 0x6f, 0xdc, 0x5b, 0x1e, 0xcf,
	0xad, 0xed, 0xdd, 0xa0, 0xf1, 0x89, 0x4a, 0x1b, 0xb4, 0x4a, 0x2f, 0x61, 0x18, 0x83, 0x15, 0x41,
	0xa5, 0xce, 0x54, 0x1d, 0xb3, 0x68, 0x62, 0x5e, 0xff, 0x73, 0x0c, 0x32, 0x26, 0x55, 0x46, 0x4f,
	0x5c, 0x00, 0xf8, 0x02, 0x00, 0xa9, 0xb0, 0x50, 0x71, 0x79, 0x07, 0xce, 0x52, 0xdf, 0x1a, 0xf4,
	0xb6, 0xd7, 0x1b, 0xfb, 0xe6, 0x82, 0xdc, 0x71, 0x73, 0x41, 0xa8, 0x6b, 0xba, 0x4b, 0x0c, 0x9f,
	0x81, 0x0e, 0x65, 0xb3, 0x4a, 0xb8, 0x7c, 0xa3, 0xb0, 0x4d, 0xd9, 0xcc, 0xc8, 0x1e, 0x03, 0xa8,
	0x45, 0x26, 0x63, 0x22, 0xf0, 0x69, 0x46, 0x67, 0x31, 0xe1, 0x9a, 0x29, 0xa7, 0xd5, 0xb7, 0x06,
	0x8b, 0xc8, 0x2e, 0x19, 0xbf, 0x22, 0xfc, 0xb2, 0x0e, 0x1f, 0x82, 0xff, 0x4d, 0xb7, 0xa2, 0x52,
	0xfd, 0x69, 0x6e, 0x9b, 0xe6, 0xb5, 0x92, 0x18, 0x9b, 0x7a, 0xd5, 0xeb, 0x81, 0x3b, 0x73, 0x2c,
	0xe3, 0x13, 0x9d, 0x31, 0x2a, 0xf0, 0x34, 0xcd, 0x52, 0x95, 0x52, 0xe9, 0x74, 0xfa, 0xd6, 0xa0,
	0x83, 0xe0, 0x1c, 0xcb, 0x4f, 0x97, 0x19, 0xb8, 0x05, 0xec, 0x42, 0xf0, 0x44, 0x50, 0x29, 0xe3,
	0x82, 0x0a, 0x42, 0x99, 0x72, 0xba, 0x7d, 0x6b, 0xb0, 0x8c, 0xd6, 0x9a, 0xfa, 0xa8, 0x2a, 0x6f,
	0x1e, 0x81, 0xd5, 0xcb, 0xa3, 0x82, 0x1b, 0xe0, 0x5e, 0xf0, 0x39, 0xf0, 0x27, 0xe3, 0xf0, 0x70,
	0x18, 0x47, 0xe3, 0x9d, 0x71, 0x10, 0x4f, 0x86, 0xd1, 0x28, 0xf0, 0xc3, 0xfd, 0x30, 0xd8, 0xb3,
	0xff, 0x83, 0x00, 0xb4, 0x3e, 0x4e, 0x82, 0x49, 0xb0, 0x67, 0x5b, 0x70, 0x05, 0x74, 0x22, 0x7f,
	0x67, 0x38, 0x0c, 0x87, 0xef, 0xec, 0x85, 0x12, 0xed, 0x87, 0xc3, 0x30, 0x7a, 0x1f, 0xec, 0xd9,
	0x8b, 0x9b, 0x87, 0xa0, 0xf7, 0xd7, 0x78, 0xe0, 0x7d, 0xe0, 0xa0, 0x20, 0x9a, 0x1c, 0x8c, 0xaf,
	0x35, 0xed, 0x81, 0x76, 0x34, 0xf1, 0xfd, 0x20, 0x8a, 0x6c, 0x0b, 0x76, 0xc1, 0x72, 0x80, 0xd0,
	0x21, 0xb2, 0x17, 0xca, 0xb0, 0x0f, 0xe1, 0xc1, 0x41, 0x69, 0xf8, 0xb2, 0xf8, 0xb5, 0x93, 0x83,
	0xad, 0x6b, 0xb6, 0xa2, 0x9a, 0x0e, 0x2e, 0x52, 0xe9, 0x12, 0x9e, 0x7b, 0xcd, 0x8a, 0xbf, 0x2d,
	0x04, 0xff, 0x4a, 0x89, 0x92, 0xde, 0x59, 0xfd, 0x75, 0x6e, 0x9e, 0x9d, 0xcf, 0xd9, 0x71, 0x9a,
	0x48, 0xef, 0xcc, 0xbc, 0x41, 0x62, 0x50, 0xc5, 0x20, 0xcd, 0x9a, 0xb2, 0xd0, 0xec, 0x7c, 0xf7,
	0x87, 0x05, 0x1e, 0x11, 0x9e, 0xdf, 0x76, 0x2d, 0x77, 0x57, 0xea, 0xe8, 0x51, 0xb9, 0x27, 0x23,
	0xeb, 0xcb, 0x51, 0x2d, 0x4c, 0x78, 0x86, 0x59, 0xe2, 0x72, 0x91, 0x78, 0x09, 0x65, 0x66, 0x8b,
	0xbc, 0x8b, 0x43, 0xdf, 0xf8, 0xb3, 0x78, 0x75, 0x95, 0x9a, 0xb6, 0x8c, 0xcb, 0xd3, 0xdf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xa3, 0x31, 0xfb, 0x25, 0x71, 0x04, 0x00, 0x00,
}
