// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/shared_set_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v3/resources"
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

// Request message for [SharedSetService.GetSharedSet][google.ads.googleads.v3.services.SharedSetService.GetSharedSet].
type GetSharedSetRequest struct {
	// Required. The resource name of the shared set to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSharedSetRequest) Reset()         { *m = GetSharedSetRequest{} }
func (m *GetSharedSetRequest) String() string { return proto.CompactTextString(m) }
func (*GetSharedSetRequest) ProtoMessage()    {}
func (*GetSharedSetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a612c7f11f5fd3bc, []int{0}
}

func (m *GetSharedSetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSharedSetRequest.Unmarshal(m, b)
}
func (m *GetSharedSetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSharedSetRequest.Marshal(b, m, deterministic)
}
func (m *GetSharedSetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSharedSetRequest.Merge(m, src)
}
func (m *GetSharedSetRequest) XXX_Size() int {
	return xxx_messageInfo_GetSharedSetRequest.Size(m)
}
func (m *GetSharedSetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSharedSetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSharedSetRequest proto.InternalMessageInfo

func (m *GetSharedSetRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [SharedSetService.MutateSharedSets][google.ads.googleads.v3.services.SharedSetService.MutateSharedSets].
type MutateSharedSetsRequest struct {
	// Required. The ID of the customer whose shared sets are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual shared sets.
	Operations []*SharedSetOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
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

func (m *MutateSharedSetsRequest) Reset()         { *m = MutateSharedSetsRequest{} }
func (m *MutateSharedSetsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateSharedSetsRequest) ProtoMessage()    {}
func (*MutateSharedSetsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a612c7f11f5fd3bc, []int{1}
}

func (m *MutateSharedSetsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateSharedSetsRequest.Unmarshal(m, b)
}
func (m *MutateSharedSetsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateSharedSetsRequest.Marshal(b, m, deterministic)
}
func (m *MutateSharedSetsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateSharedSetsRequest.Merge(m, src)
}
func (m *MutateSharedSetsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateSharedSetsRequest.Size(m)
}
func (m *MutateSharedSetsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateSharedSetsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateSharedSetsRequest proto.InternalMessageInfo

func (m *MutateSharedSetsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateSharedSetsRequest) GetOperations() []*SharedSetOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateSharedSetsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateSharedSetsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on an shared set.
type SharedSetOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*SharedSetOperation_Create
	//	*SharedSetOperation_Update
	//	*SharedSetOperation_Remove
	Operation            isSharedSetOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *SharedSetOperation) Reset()         { *m = SharedSetOperation{} }
func (m *SharedSetOperation) String() string { return proto.CompactTextString(m) }
func (*SharedSetOperation) ProtoMessage()    {}
func (*SharedSetOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_a612c7f11f5fd3bc, []int{2}
}

func (m *SharedSetOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SharedSetOperation.Unmarshal(m, b)
}
func (m *SharedSetOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SharedSetOperation.Marshal(b, m, deterministic)
}
func (m *SharedSetOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SharedSetOperation.Merge(m, src)
}
func (m *SharedSetOperation) XXX_Size() int {
	return xxx_messageInfo_SharedSetOperation.Size(m)
}
func (m *SharedSetOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_SharedSetOperation.DiscardUnknown(m)
}

var xxx_messageInfo_SharedSetOperation proto.InternalMessageInfo

func (m *SharedSetOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isSharedSetOperation_Operation interface {
	isSharedSetOperation_Operation()
}

type SharedSetOperation_Create struct {
	Create *resources.SharedSet `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type SharedSetOperation_Update struct {
	Update *resources.SharedSet `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type SharedSetOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*SharedSetOperation_Create) isSharedSetOperation_Operation() {}

func (*SharedSetOperation_Update) isSharedSetOperation_Operation() {}

func (*SharedSetOperation_Remove) isSharedSetOperation_Operation() {}

func (m *SharedSetOperation) GetOperation() isSharedSetOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *SharedSetOperation) GetCreate() *resources.SharedSet {
	if x, ok := m.GetOperation().(*SharedSetOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *SharedSetOperation) GetUpdate() *resources.SharedSet {
	if x, ok := m.GetOperation().(*SharedSetOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *SharedSetOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*SharedSetOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SharedSetOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SharedSetOperation_Create)(nil),
		(*SharedSetOperation_Update)(nil),
		(*SharedSetOperation_Remove)(nil),
	}
}

// Response message for a shared set mutate.
type MutateSharedSetsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateSharedSetResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *MutateSharedSetsResponse) Reset()         { *m = MutateSharedSetsResponse{} }
func (m *MutateSharedSetsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateSharedSetsResponse) ProtoMessage()    {}
func (*MutateSharedSetsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a612c7f11f5fd3bc, []int{3}
}

func (m *MutateSharedSetsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateSharedSetsResponse.Unmarshal(m, b)
}
func (m *MutateSharedSetsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateSharedSetsResponse.Marshal(b, m, deterministic)
}
func (m *MutateSharedSetsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateSharedSetsResponse.Merge(m, src)
}
func (m *MutateSharedSetsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateSharedSetsResponse.Size(m)
}
func (m *MutateSharedSetsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateSharedSetsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateSharedSetsResponse proto.InternalMessageInfo

func (m *MutateSharedSetsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateSharedSetsResponse) GetResults() []*MutateSharedSetResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the shared set mutate.
type MutateSharedSetResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateSharedSetResult) Reset()         { *m = MutateSharedSetResult{} }
func (m *MutateSharedSetResult) String() string { return proto.CompactTextString(m) }
func (*MutateSharedSetResult) ProtoMessage()    {}
func (*MutateSharedSetResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_a612c7f11f5fd3bc, []int{4}
}

func (m *MutateSharedSetResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateSharedSetResult.Unmarshal(m, b)
}
func (m *MutateSharedSetResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateSharedSetResult.Marshal(b, m, deterministic)
}
func (m *MutateSharedSetResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateSharedSetResult.Merge(m, src)
}
func (m *MutateSharedSetResult) XXX_Size() int {
	return xxx_messageInfo_MutateSharedSetResult.Size(m)
}
func (m *MutateSharedSetResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateSharedSetResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateSharedSetResult proto.InternalMessageInfo

func (m *MutateSharedSetResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetSharedSetRequest)(nil), "google.ads.googleads.v3.services.GetSharedSetRequest")
	proto.RegisterType((*MutateSharedSetsRequest)(nil), "google.ads.googleads.v3.services.MutateSharedSetsRequest")
	proto.RegisterType((*SharedSetOperation)(nil), "google.ads.googleads.v3.services.SharedSetOperation")
	proto.RegisterType((*MutateSharedSetsResponse)(nil), "google.ads.googleads.v3.services.MutateSharedSetsResponse")
	proto.RegisterType((*MutateSharedSetResult)(nil), "google.ads.googleads.v3.services.MutateSharedSetResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/shared_set_service.proto", fileDescriptor_a612c7f11f5fd3bc)
}

var fileDescriptor_a612c7f11f5fd3bc = []byte{
	// 757 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4d, 0x6b, 0xdb, 0x48,
	0x18, 0x5e, 0xc9, 0x4b, 0x76, 0x33, 0x4e, 0x76, 0xc3, 0x84, 0x6c, 0x84, 0x77, 0x61, 0x8d, 0x36,
	0xb0, 0xc6, 0x04, 0x09, 0xec, 0x94, 0x12, 0xa5, 0xa1, 0xc8, 0xb4, 0x49, 0x7a, 0x48, 0x93, 0xca,
	0x10, 0x68, 0x31, 0x88, 0x89, 0x34, 0x71, 0x44, 0x24, 0x8d, 0x3a, 0x33, 0x32, 0x84, 0x90, 0x4b,
	0x4f, 0xbd, 0xf7, 0x1f, 0xf4, 0xd8, 0x7b, 0x0f, 0xfd, 0x0b, 0xa1, 0xb7, 0xde, 0x72, 0x28, 0x3d,
	0xf4, 0x50, 0xfa, 0x1b, 0x7a, 0x28, 0xd2, 0x68, 0xe4, 0x8f, 0x24, 0x98, 0xe4, 0xf6, 0x7a, 0xde,
	0xe7, 0x79, 0xde, 0x6f, 0x19, 0xac, 0xf7, 0x09, 0xe9, 0x87, 0xd8, 0x44, 0x3e, 0x33, 0x85, 0x99,
	0x59, 0x83, 0xb6, 0xc9, 0x30, 0x1d, 0x04, 0x1e, 0x66, 0x26, 0x3b, 0x46, 0x14, 0xfb, 0x2e, 0xc3,
	0xdc, 0x2d, 0xde, 0x8c, 0x84, 0x12, 0x4e, 0x60, 0x5d, 0xe0, 0x0d, 0xe4, 0x33, 0xa3, 0xa4, 0x1a,
	0x83, 0xb6, 0x21, 0xa9, 0xb5, 0xd6, 0x4d, 0xe2, 0x14, 0x33, 0x92, 0xd2, 0x71, 0x75, 0xa1, 0x5a,
	0xfb, 0x47, 0x72, 0x92, 0xc0, 0x44, 0x71, 0x4c, 0x38, 0xe2, 0x01, 0x89, 0x59, 0xe1, 0x5d, 0x1e,
	0xf1, 0x7a, 0x61, 0x80, 0x63, 0x49, 0xfb, 0x77, 0xc4, 0x71, 0x14, 0xe0, 0xd0, 0x77, 0x0f, 0xf1,
	0x31, 0x1a, 0x04, 0x84, 0x16, 0x80, 0x22, 0x5b, 0x33, 0xff, 0x75, 0x98, 0x1e, 0x15, 0xa8, 0x08,
	0xb1, 0x93, 0x09, 0x6d, 0x9a, 0x78, 0x26, 0xe3, 0x88, 0xa7, 0x45, 0x50, 0xfd, 0x21, 0x58, 0xdc,
	0xc6, 0xbc, 0x9b, 0x67, 0xda, 0xc5, 0xdc, 0xc1, 0x2f, 0x53, 0xcc, 0x38, 0x6c, 0x80, 0x79, 0x59,
	0x87, 0x1b, 0xa3, 0x08, 0x6b, 0x4a, 0x5d, 0x69, 0xcc, 0x76, 0x2a, 0x5f, 0x6c, 0xd5, 0x99, 0x93,
	0x9e, 0xa7, 0x28, 0xc2, 0xfa, 0x37, 0x05, 0x2c, 0xef, 0xa6, 0x1c, 0x71, 0x5c, 0x8a, 0x30, 0xa9,
	0xb2, 0x02, 0xaa, 0x5e, 0xca, 0x38, 0x89, 0x30, 0x75, 0x03, 0x7f, 0x54, 0x03, 0xc8, 0xf7, 0x27,
	0x3e, 0x7c, 0x0e, 0x00, 0x49, 0x30, 0x15, 0xbd, 0xd0, 0xd4, 0x7a, 0xa5, 0x51, 0x6d, 0xad, 0x19,
	0xd3, 0x06, 0x60, 0x94, 0xe1, 0xf6, 0x24, 0xb9, 0x90, 0x1e, 0x8a, 0xc1, 0xff, 0xc1, 0x9f, 0x09,
	0xa2, 0x3c, 0x40, 0xa1, 0x7b, 0x84, 0x82, 0x30, 0xa5, 0x58, 0xab, 0xd4, 0x95, 0xc6, 0xef, 0xce,
	0x1f, 0xc5, 0xf3, 0x96, 0x78, 0x85, 0xff, 0x81, 0xf9, 0x01, 0x0a, 0x03, 0x1f, 0x71, 0xec, 0x92,
	0x38, 0x3c, 0xd5, 0x7e, 0xcd, 0x61, 0x73, 0xf2, 0x71, 0x2f, 0x0e, 0x4f, 0xf5, 0xd7, 0x2a, 0x80,
	0x57, 0xa3, 0xc2, 0x0d, 0x50, 0x4d, 0x93, 0x9c, 0x99, 0x35, 0x3c, 0x67, 0x56, 0x5b, 0x35, 0x59,
	0x80, 0x9c, 0x89, 0xb1, 0x95, 0xcd, 0x64, 0x17, 0xb1, 0x13, 0x07, 0x08, 0x78, 0x66, 0xc3, 0x2d,
	0x30, 0xe3, 0x51, 0x8c, 0xb8, 0xe8, 0x70, 0xb5, 0xb5, 0x7a, 0x63, 0xe1, 0xe5, 0x5e, 0x0d, 0x2b,
	0xdf, 0xf9, 0xc5, 0x29, 0xd8, 0x99, 0x8e, 0x50, 0xd5, 0xd4, 0xbb, 0xe9, 0x08, 0x36, 0xd4, 0xc0,
	0x0c, 0xc5, 0x11, 0x19, 0x88, 0x46, 0xcd, 0x66, 0x1e, 0xf1, 0xbb, 0x53, 0x05, 0xb3, 0x65, 0x67,
	0xf5, 0xf7, 0x0a, 0xd0, 0xae, 0x4e, 0x9d, 0x25, 0x24, 0x66, 0x59, 0x2e, 0x4b, 0x13, 0x5d, 0x77,
	0x31, 0xa5, 0x84, 0xe6, 0x92, 0xd5, 0x16, 0x94, 0xa9, 0xd1, 0xc4, 0x33, 0xba, 0xf9, 0x32, 0x3a,
	0x8b, 0xe3, 0xf3, 0x78, 0x9c, 0xc1, 0xe1, 0x33, 0xf0, 0x1b, 0xc5, 0x2c, 0x0d, 0xb9, 0xdc, 0x8a,
	0xfb, 0xd3, 0xb7, 0x62, 0x22, 0x29, 0x27, 0xe7, 0x3b, 0x52, 0x47, 0x7f, 0x00, 0x96, 0xae, 0x45,
	0x64, 0x0b, 0x70, 0xcd, 0xc2, 0x8f, 0xef, 0x7a, 0xeb, 0x63, 0x05, 0x2c, 0x94, 0xc4, 0xae, 0x08,
	0x09, 0x3f, 0x28, 0x60, 0x6e, 0xf4, 0x84, 0xe0, 0xbd, 0xe9, 0x59, 0x5e, 0x73, 0x72, 0xb5, 0x5b,
	0x4d, 0x4c, 0x7f, 0x74, 0x69, 0x8f, 0x27, 0xfc, 0xea, 0xd3, 0xd7, 0x37, 0xaa, 0x01, 0x57, 0xb3,
	0x4f, 0xd0, 0xd9, 0x98, 0x67, 0x53, 0xde, 0x1a, 0x33, 0x9b, 0xc5, 0x37, 0x29, 0x1b, 0x97, 0xd9,
	0x3c, 0x87, 0x9f, 0x15, 0xb0, 0x30, 0x39, 0x46, 0xb8, 0x7e, 0xeb, 0x2e, 0xcb, 0x83, 0xaf, 0x59,
	0x77, 0xa1, 0x8a, 0xad, 0xd1, 0xbb, 0x97, 0xf6, 0x5f, 0x23, 0x5f, 0x8b, 0xd5, 0xe1, 0x19, 0xe7,
	0xa5, 0xad, 0xe9, 0x66, 0x56, 0xda, 0xb0, 0x96, 0xb3, 0x11, 0xf0, 0x66, 0xf3, 0x7c, 0xa4, 0x32,
	0x2b, 0xca, 0x63, 0x58, 0x4a, 0xb3, 0xf6, 0xf7, 0x85, 0xad, 0x0d, 0xf3, 0x28, 0xac, 0x24, 0x60,
	0x86, 0x47, 0xa2, 0xce, 0x0f, 0x05, 0xac, 0x78, 0x24, 0x9a, 0x9a, 0x73, 0x67, 0x69, 0x72, 0xe8,
	0xfb, 0xd9, 0x51, 0xef, 0x2b, 0x2f, 0x76, 0x0a, 0x6a, 0x9f, 0x84, 0x28, 0xee, 0x1b, 0x84, 0xf6,
	0xcd, 0x3e, 0x8e, 0xf3, 0x93, 0x37, 0x87, 0xc1, 0x6e, 0xfe, 0x03, 0xda, 0x90, 0xc6, 0x5b, 0xb5,
	0xb2, 0x6d, 0xdb, 0xef, 0xd4, 0xfa, 0xb6, 0x10, 0xb4, 0x7d, 0x66, 0x08, 0x33, 0xb3, 0x0e, 0xda,
	0x46, 0x11, 0x98, 0x5d, 0x48, 0x48, 0xcf, 0xf6, 0x59, 0xaf, 0x84, 0xf4, 0x0e, 0xda, 0x3d, 0x09,
	0xf9, 0xae, 0xae, 0x88, 0x77, 0xcb, 0xb2, 0x7d, 0x66, 0x59, 0x25, 0xc8, 0xb2, 0x0e, 0xda, 0x96,
	0x25, 0x61, 0x87, 0x33, 0x79, 0x9e, 0xed, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x4c, 0xb1,
	0x43, 0x27, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SharedSetServiceClient is the client API for SharedSetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SharedSetServiceClient interface {
	// Returns the requested shared set in full detail.
	GetSharedSet(ctx context.Context, in *GetSharedSetRequest, opts ...grpc.CallOption) (*resources.SharedSet, error)
	// Creates, updates, or removes shared sets. Operation statuses are returned.
	MutateSharedSets(ctx context.Context, in *MutateSharedSetsRequest, opts ...grpc.CallOption) (*MutateSharedSetsResponse, error)
}

type sharedSetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSharedSetServiceClient(cc grpc.ClientConnInterface) SharedSetServiceClient {
	return &sharedSetServiceClient{cc}
}

func (c *sharedSetServiceClient) GetSharedSet(ctx context.Context, in *GetSharedSetRequest, opts ...grpc.CallOption) (*resources.SharedSet, error) {
	out := new(resources.SharedSet)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.SharedSetService/GetSharedSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sharedSetServiceClient) MutateSharedSets(ctx context.Context, in *MutateSharedSetsRequest, opts ...grpc.CallOption) (*MutateSharedSetsResponse, error) {
	out := new(MutateSharedSetsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.SharedSetService/MutateSharedSets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SharedSetServiceServer is the server API for SharedSetService service.
type SharedSetServiceServer interface {
	// Returns the requested shared set in full detail.
	GetSharedSet(context.Context, *GetSharedSetRequest) (*resources.SharedSet, error)
	// Creates, updates, or removes shared sets. Operation statuses are returned.
	MutateSharedSets(context.Context, *MutateSharedSetsRequest) (*MutateSharedSetsResponse, error)
}

// UnimplementedSharedSetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSharedSetServiceServer struct {
}

func (*UnimplementedSharedSetServiceServer) GetSharedSet(ctx context.Context, req *GetSharedSetRequest) (*resources.SharedSet, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetSharedSet not implemented")
}
func (*UnimplementedSharedSetServiceServer) MutateSharedSets(ctx context.Context, req *MutateSharedSetsRequest) (*MutateSharedSetsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateSharedSets not implemented")
}

func RegisterSharedSetServiceServer(s *grpc.Server, srv SharedSetServiceServer) {
	s.RegisterService(&_SharedSetService_serviceDesc, srv)
}

func _SharedSetService_GetSharedSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSharedSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharedSetServiceServer).GetSharedSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.SharedSetService/GetSharedSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharedSetServiceServer).GetSharedSet(ctx, req.(*GetSharedSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SharedSetService_MutateSharedSets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateSharedSetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharedSetServiceServer).MutateSharedSets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.SharedSetService/MutateSharedSets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharedSetServiceServer).MutateSharedSets(ctx, req.(*MutateSharedSetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SharedSetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.SharedSetService",
	HandlerType: (*SharedSetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSharedSet",
			Handler:    _SharedSetService_GetSharedSet_Handler,
		},
		{
			MethodName: "MutateSharedSets",
			Handler:    _SharedSetService_MutateSharedSets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/shared_set_service.proto",
}
