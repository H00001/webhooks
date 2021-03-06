// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/servicedirectory/v1beta1/service.proto

package servicedirectory

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// An individual service. A service contains a name and optional metadata.
// A service must exist before
// [endpoints][google.cloud.servicedirectory.v1beta1.Endpoint] can be
// added to it.
type Service struct {
	// Immutable. The resource name for the service in the format
	// 'projects/*/locations/*/namespaces/*/services/*'.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. Metadata for the service. This data can be consumed by service
	// clients.  The entire metadata dictionary may contain up to 2000 characters,
	// spread across all key-value pairs. Metadata that goes beyond any these
	// limits will be rejected.
	Metadata map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Output only. Endpoints associated with this service. Returned on LookupService.Resolve.
	// Control plane clients should use RegistrationService.ListEndpoints.
	Endpoints            []*Endpoint `protobuf:"bytes,3,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_2464dc41fdae2e63, []int{0}
}

func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Service) GetEndpoints() []*Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func init() {
	proto.RegisterType((*Service)(nil), "google.cloud.servicedirectory.v1beta1.Service")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.servicedirectory.v1beta1.Service.MetadataEntry")
}

func init() {
	proto.RegisterFile("google/cloud/servicedirectory/v1beta1/service.proto", fileDescriptor_2464dc41fdae2e63)
}

var fileDescriptor_2464dc41fdae2e63 = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x8b, 0xd4, 0x30,
	0x14, 0xc7, 0xc9, 0xd4, 0x55, 0x37, 0x2a, 0x48, 0x11, 0xac, 0x45, 0x70, 0x11, 0xc4, 0xf5, 0x92,
	0xb0, 0xae, 0x07, 0xd9, 0xf5, 0xb2, 0x03, 0x7b, 0x14, 0xc6, 0x8a, 0x17, 0x41, 0x24, 0x93, 0x3e,
	0x6b, 0xb4, 0xcd, 0x2b, 0x49, 0xa6, 0x30, 0xd4, 0xe2, 0x1f, 0x21, 0xfe, 0x7f, 0x9e, 0xfd, 0x0b,
	0x3c, 0x4a, 0x9b, 0x74, 0x66, 0x1c, 0x0f, 0xdb, 0xdb, 0xb7, 0xef, 0xf5, 0xfb, 0x79, 0xbf, 0x42,
	0x4f, 0x0b, 0xc4, 0xa2, 0x04, 0x2e, 0x4b, 0x5c, 0xe5, 0xdc, 0x82, 0x69, 0x94, 0x84, 0x5c, 0x19,
	0x90, 0x0e, 0xcd, 0x9a, 0x37, 0x27, 0x4b, 0x70, 0xe2, 0x64, 0x4c, 0xb0, 0xda, 0xa0, 0xc3, 0xf8,
	0x89, 0x37, 0xb1, 0xc1, 0xc4, 0xf6, 0x4d, 0x2c, 0x98, 0xd2, 0x47, 0x81, 0x2d, 0x6a, 0xc5, 0x3f,
	0x29, 0x28, 0xf3, 0x8f, 0x4b, 0xf8, 0x2c, 0x1a, 0x85, 0xc6, 0x73, 0xd2, 0x07, 0x3b, 0x3f, 0x18,
	0xb0, 0xb8, 0x32, 0x63, 0x89, 0xf4, 0xc5, 0xb4, 0xbe, 0x40, 0xe7, 0x35, 0x2a, 0xed, 0x82, 0xeb,
	0xe1, 0x0e, 0x50, 0x68, 0x8d, 0x4e, 0x38, 0x85, 0xda, 0xfa, 0xec, 0xe3, 0x1f, 0x11, 0xbd, 0xf1,
	0xd6, 0x93, 0xe2, 0xfb, 0xf4, 0x9a, 0x16, 0x15, 0x24, 0xe4, 0x88, 0x1c, 0x1f, 0xce, 0xa3, 0x5f,
	0x17, 0x07, 0xd9, 0x10, 0x88, 0x3f, 0xd0, 0x9b, 0x15, 0x38, 0x91, 0x0b, 0x27, 0x92, 0xd9, 0x51,
	0x74, 0x7c, 0xeb, 0xf9, 0x2b, 0x36, 0x69, 0x5c, 0x16, 0xd0, 0xec, 0x75, 0xb0, 0x5f, 0x6a, 0x67,
	0xd6, 0x3d, 0x9a, 0x64, 0x1b, 0x64, 0x9c, 0xd1, 0xc3, 0xb1, 0x67, 0x9b, 0x44, 0x03, 0x9f, 0x4f,
	0xe4, 0x5f, 0x06, 0x5f, 0x8f, 0x8c, 0xb2, 0x2d, 0x26, 0x3d, 0xa7, 0x77, 0xfe, 0xa9, 0x19, 0xdf,
	0xa5, 0xd1, 0x57, 0x58, 0xfb, 0xd9, 0xb2, 0x5e, 0xc6, 0xf7, 0xe8, 0x41, 0x23, 0xca, 0x15, 0x24,
	0xb3, 0x21, 0xe6, 0x3f, 0xce, 0x66, 0x2f, 0xc9, 0xd9, 0xf7, 0xdf, 0x17, 0xdf, 0xe8, 0xd3, 0xff,
	0x8a, 0xfa, 0x96, 0x44, 0xad, 0x2c, 0x93, 0x58, 0xf1, 0x71, 0x6d, 0x6f, 0x6a, 0x83, 0x5f, 0x40,
	0x3a, 0xcb, 0xdb, 0xa0, 0x3a, 0x5e, 0xa2, 0xf4, 0x6b, 0xe6, 0xed, 0x28, 0x3b, 0xde, 0xef, 0xd2,
	0xd6, 0x42, 0x82, 0xe5, 0xed, 0x46, 0x77, 0xe3, 0x21, 0x2d, 0x6f, 0x83, 0xea, 0xe6, 0x3f, 0x09,
	0x7d, 0x26, 0xb1, 0x9a, 0xb6, 0x84, 0xf9, 0xed, 0xd0, 0xc9, 0xa2, 0xbf, 0xe8, 0x82, 0xbc, 0x7f,
	0x17, 0x6c, 0x05, 0x96, 0x42, 0x17, 0x0c, 0x4d, 0xc1, 0x0b, 0xd0, 0xc3, 0xbd, 0xf9, 0x76, 0x86,
	0x2b, 0x9e, 0xd1, 0xf9, 0x7e, 0xe2, 0x0f, 0x21, 0xcb, 0xeb, 0x03, 0xe4, 0xf4, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xe5, 0xbb, 0x8d, 0x0d, 0x22, 0x03, 0x00, 0x00,
}
