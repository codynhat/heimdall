// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/ad_group_extension_setting_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

// Request message for
// [AdGroupExtensionSettingService.GetAdGroupExtensionSetting][google.ads.googleads.v1.services.AdGroupExtensionSettingService.GetAdGroupExtensionSetting].
type GetAdGroupExtensionSettingRequest struct {
	// The resource name of the ad group extension setting to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAdGroupExtensionSettingRequest) Reset()         { *m = GetAdGroupExtensionSettingRequest{} }
func (m *GetAdGroupExtensionSettingRequest) String() string { return proto.CompactTextString(m) }
func (*GetAdGroupExtensionSettingRequest) ProtoMessage()    {}
func (*GetAdGroupExtensionSettingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_054674c18d01b7a6, []int{0}
}

func (m *GetAdGroupExtensionSettingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAdGroupExtensionSettingRequest.Unmarshal(m, b)
}
func (m *GetAdGroupExtensionSettingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAdGroupExtensionSettingRequest.Marshal(b, m, deterministic)
}
func (m *GetAdGroupExtensionSettingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAdGroupExtensionSettingRequest.Merge(m, src)
}
func (m *GetAdGroupExtensionSettingRequest) XXX_Size() int {
	return xxx_messageInfo_GetAdGroupExtensionSettingRequest.Size(m)
}
func (m *GetAdGroupExtensionSettingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAdGroupExtensionSettingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAdGroupExtensionSettingRequest proto.InternalMessageInfo

func (m *GetAdGroupExtensionSettingRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for
// [AdGroupExtensionSettingService.MutateAdGroupExtensionSettings][google.ads.googleads.v1.services.AdGroupExtensionSettingService.MutateAdGroupExtensionSettings].
type MutateAdGroupExtensionSettingsRequest struct {
	// The ID of the customer whose ad group extension settings are being
	// modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual ad group extension
	// settings.
	Operations []*AdGroupExtensionSettingOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAdGroupExtensionSettingsRequest) Reset()         { *m = MutateAdGroupExtensionSettingsRequest{} }
func (m *MutateAdGroupExtensionSettingsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupExtensionSettingsRequest) ProtoMessage()    {}
func (*MutateAdGroupExtensionSettingsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_054674c18d01b7a6, []int{1}
}

func (m *MutateAdGroupExtensionSettingsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupExtensionSettingsRequest.Unmarshal(m, b)
}
func (m *MutateAdGroupExtensionSettingsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupExtensionSettingsRequest.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupExtensionSettingsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupExtensionSettingsRequest.Merge(m, src)
}
func (m *MutateAdGroupExtensionSettingsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupExtensionSettingsRequest.Size(m)
}
func (m *MutateAdGroupExtensionSettingsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupExtensionSettingsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupExtensionSettingsRequest proto.InternalMessageInfo

func (m *MutateAdGroupExtensionSettingsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateAdGroupExtensionSettingsRequest) GetOperations() []*AdGroupExtensionSettingOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateAdGroupExtensionSettingsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateAdGroupExtensionSettingsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on an ad group extension setting.
type AdGroupExtensionSettingOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*AdGroupExtensionSettingOperation_Create
	//	*AdGroupExtensionSettingOperation_Update
	//	*AdGroupExtensionSettingOperation_Remove
	Operation            isAdGroupExtensionSettingOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *AdGroupExtensionSettingOperation) Reset()         { *m = AdGroupExtensionSettingOperation{} }
func (m *AdGroupExtensionSettingOperation) String() string { return proto.CompactTextString(m) }
func (*AdGroupExtensionSettingOperation) ProtoMessage()    {}
func (*AdGroupExtensionSettingOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_054674c18d01b7a6, []int{2}
}

func (m *AdGroupExtensionSettingOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupExtensionSettingOperation.Unmarshal(m, b)
}
func (m *AdGroupExtensionSettingOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupExtensionSettingOperation.Marshal(b, m, deterministic)
}
func (m *AdGroupExtensionSettingOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupExtensionSettingOperation.Merge(m, src)
}
func (m *AdGroupExtensionSettingOperation) XXX_Size() int {
	return xxx_messageInfo_AdGroupExtensionSettingOperation.Size(m)
}
func (m *AdGroupExtensionSettingOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupExtensionSettingOperation.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupExtensionSettingOperation proto.InternalMessageInfo

func (m *AdGroupExtensionSettingOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isAdGroupExtensionSettingOperation_Operation interface {
	isAdGroupExtensionSettingOperation_Operation()
}

type AdGroupExtensionSettingOperation_Create struct {
	Create *resources.AdGroupExtensionSetting `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type AdGroupExtensionSettingOperation_Update struct {
	Update *resources.AdGroupExtensionSetting `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type AdGroupExtensionSettingOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*AdGroupExtensionSettingOperation_Create) isAdGroupExtensionSettingOperation_Operation() {}

func (*AdGroupExtensionSettingOperation_Update) isAdGroupExtensionSettingOperation_Operation() {}

func (*AdGroupExtensionSettingOperation_Remove) isAdGroupExtensionSettingOperation_Operation() {}

func (m *AdGroupExtensionSettingOperation) GetOperation() isAdGroupExtensionSettingOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *AdGroupExtensionSettingOperation) GetCreate() *resources.AdGroupExtensionSetting {
	if x, ok := m.GetOperation().(*AdGroupExtensionSettingOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *AdGroupExtensionSettingOperation) GetUpdate() *resources.AdGroupExtensionSetting {
	if x, ok := m.GetOperation().(*AdGroupExtensionSettingOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *AdGroupExtensionSettingOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*AdGroupExtensionSettingOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AdGroupExtensionSettingOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AdGroupExtensionSettingOperation_Create)(nil),
		(*AdGroupExtensionSettingOperation_Update)(nil),
		(*AdGroupExtensionSettingOperation_Remove)(nil),
	}
}

// Response message for an ad group extension setting mutate.
type MutateAdGroupExtensionSettingsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateAdGroupExtensionSettingResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *MutateAdGroupExtensionSettingsResponse) Reset() {
	*m = MutateAdGroupExtensionSettingsResponse{}
}
func (m *MutateAdGroupExtensionSettingsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupExtensionSettingsResponse) ProtoMessage()    {}
func (*MutateAdGroupExtensionSettingsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_054674c18d01b7a6, []int{3}
}

func (m *MutateAdGroupExtensionSettingsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupExtensionSettingsResponse.Unmarshal(m, b)
}
func (m *MutateAdGroupExtensionSettingsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupExtensionSettingsResponse.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupExtensionSettingsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupExtensionSettingsResponse.Merge(m, src)
}
func (m *MutateAdGroupExtensionSettingsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupExtensionSettingsResponse.Size(m)
}
func (m *MutateAdGroupExtensionSettingsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupExtensionSettingsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupExtensionSettingsResponse proto.InternalMessageInfo

func (m *MutateAdGroupExtensionSettingsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateAdGroupExtensionSettingsResponse) GetResults() []*MutateAdGroupExtensionSettingResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the ad group extension setting mutate.
type MutateAdGroupExtensionSettingResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAdGroupExtensionSettingResult) Reset()         { *m = MutateAdGroupExtensionSettingResult{} }
func (m *MutateAdGroupExtensionSettingResult) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupExtensionSettingResult) ProtoMessage()    {}
func (*MutateAdGroupExtensionSettingResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_054674c18d01b7a6, []int{4}
}

func (m *MutateAdGroupExtensionSettingResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupExtensionSettingResult.Unmarshal(m, b)
}
func (m *MutateAdGroupExtensionSettingResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupExtensionSettingResult.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupExtensionSettingResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupExtensionSettingResult.Merge(m, src)
}
func (m *MutateAdGroupExtensionSettingResult) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupExtensionSettingResult.Size(m)
}
func (m *MutateAdGroupExtensionSettingResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupExtensionSettingResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupExtensionSettingResult proto.InternalMessageInfo

func (m *MutateAdGroupExtensionSettingResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAdGroupExtensionSettingRequest)(nil), "google.ads.googleads.v1.services.GetAdGroupExtensionSettingRequest")
	proto.RegisterType((*MutateAdGroupExtensionSettingsRequest)(nil), "google.ads.googleads.v1.services.MutateAdGroupExtensionSettingsRequest")
	proto.RegisterType((*AdGroupExtensionSettingOperation)(nil), "google.ads.googleads.v1.services.AdGroupExtensionSettingOperation")
	proto.RegisterType((*MutateAdGroupExtensionSettingsResponse)(nil), "google.ads.googleads.v1.services.MutateAdGroupExtensionSettingsResponse")
	proto.RegisterType((*MutateAdGroupExtensionSettingResult)(nil), "google.ads.googleads.v1.services.MutateAdGroupExtensionSettingResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/ad_group_extension_setting_service.proto", fileDescriptor_054674c18d01b7a6)
}

var fileDescriptor_054674c18d01b7a6 = []byte{
	// 740 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xc7, 0x7f, 0x76, 0xab, 0xfe, 0xe8, 0xa6, 0x80, 0xb4, 0x08, 0x11, 0x05, 0x54, 0x82, 0x5b,
	0xa0, 0xca, 0xc1, 0x56, 0xc2, 0xcd, 0xa5, 0x87, 0x18, 0xb5, 0x49, 0x91, 0x4a, 0x2b, 0x17, 0xf5,
	0x80, 0x22, 0x59, 0x5b, 0x7b, 0x6b, 0x59, 0xb5, 0xbd, 0x66, 0x77, 0x1d, 0x51, 0x55, 0xbd, 0x70,
	0xe4, 0xca, 0x13, 0xc0, 0x91, 0x23, 0x8f, 0xd1, 0x1b, 0xe2, 0x09, 0x90, 0x38, 0xf1, 0x08, 0x48,
	0x48, 0x68, 0xbd, 0xde, 0xf4, 0x8f, 0xe4, 0xb8, 0x52, 0xb9, 0x8d, 0x67, 0xbf, 0xf9, 0xcc, 0xcc,
	0xce, 0xec, 0x04, 0x6c, 0x86, 0x84, 0x84, 0x31, 0xb6, 0x50, 0xc0, 0x2c, 0x69, 0x0a, 0x6b, 0xdc,
	0xb5, 0x18, 0xa6, 0xe3, 0xc8, 0xc7, 0xcc, 0x42, 0x81, 0x17, 0x52, 0x92, 0x67, 0x1e, 0x7e, 0xc7,
	0x71, 0xca, 0x22, 0x92, 0x7a, 0x0c, 0x73, 0x1e, 0xa5, 0xa1, 0x57, 0x6a, 0xcc, 0x8c, 0x12, 0x4e,
	0x60, 0x5b, 0xfe, 0xde, 0x44, 0x01, 0x33, 0x27, 0x28, 0x73, 0xdc, 0x35, 0x15, 0xaa, 0xe5, 0x54,
	0x05, 0xa3, 0x98, 0x91, 0x9c, 0x4e, 0x8f, 0x26, 0xa3, 0xb4, 0x1e, 0x28, 0x46, 0x16, 0x59, 0x28,
	0x4d, 0x09, 0x47, 0x3c, 0x22, 0x29, 0x2b, 0x4f, 0xcb, 0x1c, 0xac, 0xe2, 0x6b, 0x3f, 0x3f, 0xb0,
	0x0e, 0x22, 0x1c, 0x07, 0x5e, 0x82, 0xd8, 0x61, 0xa9, 0xb8, 0x57, 0x2a, 0x68, 0xe6, 0x5b, 0x8c,
	0x23, 0x9e, 0xb3, 0x4b, 0x07, 0x02, 0xec, 0xc7, 0x11, 0x4e, 0xb9, 0x3c, 0x30, 0x86, 0xe0, 0xd1,
	0x00, 0xf3, 0x7e, 0x30, 0x10, 0x79, 0xad, 0xab, 0xb4, 0x76, 0x65, 0x56, 0x2e, 0x7e, 0x9b, 0x63,
	0xc6, 0xe1, 0x12, 0xb8, 0xa9, 0x8a, 0xf0, 0x52, 0x94, 0xe0, 0xa6, 0xd6, 0xd6, 0x56, 0xe6, 0xdd,
	0x05, 0xe5, 0x7c, 0x85, 0x12, 0x6c, 0xfc, 0xd6, 0xc0, 0xe3, 0xad, 0x9c, 0x23, 0x8e, 0x2b, 0x68,
	0x4c, 0xe1, 0x1e, 0x82, 0x86, 0x9f, 0x33, 0x4e, 0x12, 0x4c, 0xbd, 0x28, 0x28, 0x61, 0x40, 0xb9,
	0x36, 0x03, 0xb8, 0x0f, 0x00, 0xc9, 0x30, 0x95, 0xc5, 0x37, 0xf5, 0xf6, 0xcc, 0x4a, 0xa3, 0xe7,
	0x98, 0x75, 0x1d, 0x30, 0x2b, 0xe2, 0x6e, 0x2b, 0x94, 0x7b, 0x8e, 0x0a, 0x9f, 0x82, 0xdb, 0x19,
	0xa2, 0x3c, 0x42, 0xb1, 0x77, 0x80, 0xa2, 0x38, 0xa7, 0xb8, 0x39, 0xd3, 0xd6, 0x56, 0x6e, 0xb8,
	0xb7, 0x4a, 0xf7, 0x86, 0xf4, 0x8a, 0xe2, 0xc7, 0x28, 0x8e, 0x02, 0xc4, 0xb1, 0x47, 0xd2, 0xf8,
	0xa8, 0x39, 0x5b, 0xc8, 0x16, 0x94, 0x73, 0x3b, 0x8d, 0x8f, 0x8c, 0xaf, 0x3a, 0x68, 0xd7, 0x85,
	0x87, 0xab, 0xa0, 0x91, 0x67, 0x05, 0x47, 0xb4, 0xac, 0xe0, 0x34, 0x7a, 0x2d, 0x55, 0x97, 0xea,
	0xaa, 0xb9, 0x21, 0xba, 0xba, 0x85, 0xd8, 0xa1, 0x0b, 0xa4, 0x5c, 0xd8, 0xf0, 0x35, 0x98, 0xf3,
	0x29, 0x46, 0x5c, 0x5e, 0x7e, 0xa3, 0x67, 0x57, 0xde, 0xc7, 0x64, 0xde, 0xaa, 0x2e, 0x64, 0xf8,
	0x9f, 0x5b, 0xb2, 0x04, 0x55, 0xc6, 0x68, 0xea, 0xff, 0x82, 0x2a, 0x59, 0xb0, 0x09, 0xe6, 0x28,
	0x4e, 0xc8, 0x58, 0x5e, 0xe9, 0xbc, 0x38, 0x91, 0xdf, 0x4e, 0x03, 0xcc, 0x4f, 0x7a, 0x60, 0x7c,
	0xd3, 0xc0, 0x93, 0xba, 0x89, 0x61, 0x19, 0x49, 0x19, 0x86, 0x1b, 0xe0, 0xee, 0xa5, 0x6e, 0x79,
	0x98, 0x52, 0x42, 0x8b, 0x00, 0x8d, 0x1e, 0x54, 0x69, 0xd3, 0xcc, 0x37, 0x77, 0x8b, 0xc1, 0x77,
	0xef, 0x5c, 0xec, 0xe3, 0xba, 0x90, 0x43, 0x0f, 0xfc, 0x4f, 0x31, 0xcb, 0x63, 0xae, 0xc6, 0x6a,
	0xbd, 0x7e, 0xac, 0xa6, 0xa6, 0xe8, 0x16, 0x34, 0x57, 0x51, 0x8d, 0x97, 0x60, 0xe9, 0x0a, 0xfa,
	0x2b, 0xbd, 0xa8, 0xde, 0xa7, 0x59, 0xb0, 0x58, 0x81, 0xd9, 0x95, 0xc9, 0xc1, 0x1f, 0x1a, 0x68,
	0x55, 0xbf, 0x5f, 0xf8, 0xa2, 0xbe, 0xba, 0xda, 0xd7, 0xdf, 0xba, 0xc6, 0x4c, 0x18, 0xce, 0xfb,
	0xef, 0x3f, 0x3f, 0xea, 0xcf, 0xa1, 0x2d, 0x16, 0xe1, 0xf1, 0x85, 0x92, 0xd7, 0xd4, 0x83, 0x67,
	0x56, 0xc7, 0x42, 0x15, 0x03, 0x60, 0x75, 0x4e, 0xe0, 0x1f, 0x0d, 0x2c, 0x4e, 0x1f, 0x13, 0x38,
	0xb8, 0x66, 0x17, 0xd5, 0x6a, 0x6a, 0x0d, 0xaf, 0x0f, 0x92, 0x13, 0x6b, 0x0c, 0x8b, 0xca, 0x1d,
	0x63, 0x4d, 0x54, 0x7e, 0x56, 0xea, 0xf1, 0xb9, 0xcd, 0xb7, 0xd6, 0x39, 0xa9, 0x2c, 0xdc, 0x4e,
	0x8a, 0x30, 0xb6, 0xd6, 0x69, 0xdd, 0x3f, 0xed, 0x37, 0xcf, 0x52, 0x29, 0xad, 0x2c, 0x62, 0xa6,
	0x4f, 0x12, 0xe7, 0x83, 0x0e, 0x96, 0x7d, 0x92, 0xd4, 0xa6, 0xed, 0x2c, 0x4d, 0x9f, 0xa4, 0x1d,
	0xb1, 0x7d, 0x76, 0xb4, 0x37, 0xc3, 0x12, 0x14, 0x92, 0x18, 0xa5, 0xa1, 0x49, 0x68, 0x68, 0x85,
	0x38, 0x2d, 0x76, 0x93, 0x75, 0x16, 0xba, 0xfa, 0x1f, 0x75, 0x55, 0x19, 0x9f, 0xf5, 0x99, 0x41,
	0xbf, 0xff, 0x45, 0x6f, 0x0f, 0x24, 0xb0, 0x1f, 0x30, 0x53, 0x9a, 0xc2, 0xda, 0xeb, 0x9a, 0x65,
	0x60, 0x76, 0xaa, 0x24, 0xa3, 0x7e, 0xc0, 0x46, 0x13, 0xc9, 0x68, 0xaf, 0x3b, 0x52, 0x92, 0x5f,
	0xfa, 0xb2, 0xf4, 0xdb, 0x76, 0x3f, 0x60, 0xb6, 0x3d, 0x11, 0xd9, 0xf6, 0x5e, 0xd7, 0xb6, 0x95,
	0x6c, 0x7f, 0xae, 0xc8, 0xf3, 0xd9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x72, 0x76, 0x17, 0x96,
	0xf8, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdGroupExtensionSettingServiceClient is the client API for AdGroupExtensionSettingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdGroupExtensionSettingServiceClient interface {
	// Returns the requested ad group extension setting in full detail.
	GetAdGroupExtensionSetting(ctx context.Context, in *GetAdGroupExtensionSettingRequest, opts ...grpc.CallOption) (*resources.AdGroupExtensionSetting, error)
	// Creates, updates, or removes ad group extension settings. Operation
	// statuses are returned.
	MutateAdGroupExtensionSettings(ctx context.Context, in *MutateAdGroupExtensionSettingsRequest, opts ...grpc.CallOption) (*MutateAdGroupExtensionSettingsResponse, error)
}

type adGroupExtensionSettingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupExtensionSettingServiceClient(cc grpc.ClientConnInterface) AdGroupExtensionSettingServiceClient {
	return &adGroupExtensionSettingServiceClient{cc}
}

func (c *adGroupExtensionSettingServiceClient) GetAdGroupExtensionSetting(ctx context.Context, in *GetAdGroupExtensionSettingRequest, opts ...grpc.CallOption) (*resources.AdGroupExtensionSetting, error) {
	out := new(resources.AdGroupExtensionSetting)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.AdGroupExtensionSettingService/GetAdGroupExtensionSetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adGroupExtensionSettingServiceClient) MutateAdGroupExtensionSettings(ctx context.Context, in *MutateAdGroupExtensionSettingsRequest, opts ...grpc.CallOption) (*MutateAdGroupExtensionSettingsResponse, error) {
	out := new(MutateAdGroupExtensionSettingsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.AdGroupExtensionSettingService/MutateAdGroupExtensionSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupExtensionSettingServiceServer is the server API for AdGroupExtensionSettingService service.
type AdGroupExtensionSettingServiceServer interface {
	// Returns the requested ad group extension setting in full detail.
	GetAdGroupExtensionSetting(context.Context, *GetAdGroupExtensionSettingRequest) (*resources.AdGroupExtensionSetting, error)
	// Creates, updates, or removes ad group extension settings. Operation
	// statuses are returned.
	MutateAdGroupExtensionSettings(context.Context, *MutateAdGroupExtensionSettingsRequest) (*MutateAdGroupExtensionSettingsResponse, error)
}

// UnimplementedAdGroupExtensionSettingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdGroupExtensionSettingServiceServer struct {
}

func (*UnimplementedAdGroupExtensionSettingServiceServer) GetAdGroupExtensionSetting(ctx context.Context, req *GetAdGroupExtensionSettingRequest) (*resources.AdGroupExtensionSetting, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetAdGroupExtensionSetting not implemented")
}
func (*UnimplementedAdGroupExtensionSettingServiceServer) MutateAdGroupExtensionSettings(ctx context.Context, req *MutateAdGroupExtensionSettingsRequest) (*MutateAdGroupExtensionSettingsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateAdGroupExtensionSettings not implemented")
}

func RegisterAdGroupExtensionSettingServiceServer(s *grpc.Server, srv AdGroupExtensionSettingServiceServer) {
	s.RegisterService(&_AdGroupExtensionSettingService_serviceDesc, srv)
}

func _AdGroupExtensionSettingService_GetAdGroupExtensionSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdGroupExtensionSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupExtensionSettingServiceServer).GetAdGroupExtensionSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.AdGroupExtensionSettingService/GetAdGroupExtensionSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupExtensionSettingServiceServer).GetAdGroupExtensionSetting(ctx, req.(*GetAdGroupExtensionSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdGroupExtensionSettingService_MutateAdGroupExtensionSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupExtensionSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupExtensionSettingServiceServer).MutateAdGroupExtensionSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.AdGroupExtensionSettingService/MutateAdGroupExtensionSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupExtensionSettingServiceServer).MutateAdGroupExtensionSettings(ctx, req.(*MutateAdGroupExtensionSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdGroupExtensionSettingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.AdGroupExtensionSettingService",
	HandlerType: (*AdGroupExtensionSettingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdGroupExtensionSetting",
			Handler:    _AdGroupExtensionSettingService_GetAdGroupExtensionSetting_Handler,
		},
		{
			MethodName: "MutateAdGroupExtensionSettings",
			Handler:    _AdGroupExtensionSettingService_MutateAdGroupExtensionSettings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/ad_group_extension_setting_service.proto",
}
