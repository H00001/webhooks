// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/paid_organic_search_term_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A paid organic search term view providing a view of search stats across
// ads and organic listings aggregated by search term at the ad group level.
type PaidOrganicSearchTermView struct {
	// The resource name of the search term view.
	// Search term view resource names have the form:
	//
	// `customers/{customer_id}/paidOrganicSearchTermViews/{campaign_id}~
	// {ad_group_id}~{URL-base64 search term}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The search term.
	SearchTerm           *wrappers.StringValue `protobuf:"bytes,2,opt,name=search_term,json=searchTerm,proto3" json:"search_term,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PaidOrganicSearchTermView) Reset()         { *m = PaidOrganicSearchTermView{} }
func (m *PaidOrganicSearchTermView) String() string { return proto.CompactTextString(m) }
func (*PaidOrganicSearchTermView) ProtoMessage()    {}
func (*PaidOrganicSearchTermView) Descriptor() ([]byte, []int) {
	return fileDescriptor_9878fc316cec3005, []int{0}
}

func (m *PaidOrganicSearchTermView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaidOrganicSearchTermView.Unmarshal(m, b)
}
func (m *PaidOrganicSearchTermView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaidOrganicSearchTermView.Marshal(b, m, deterministic)
}
func (m *PaidOrganicSearchTermView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaidOrganicSearchTermView.Merge(m, src)
}
func (m *PaidOrganicSearchTermView) XXX_Size() int {
	return xxx_messageInfo_PaidOrganicSearchTermView.Size(m)
}
func (m *PaidOrganicSearchTermView) XXX_DiscardUnknown() {
	xxx_messageInfo_PaidOrganicSearchTermView.DiscardUnknown(m)
}

var xxx_messageInfo_PaidOrganicSearchTermView proto.InternalMessageInfo

func (m *PaidOrganicSearchTermView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *PaidOrganicSearchTermView) GetSearchTerm() *wrappers.StringValue {
	if m != nil {
		return m.SearchTerm
	}
	return nil
}

func init() {
	proto.RegisterType((*PaidOrganicSearchTermView)(nil), "google.ads.googleads.v3.resources.PaidOrganicSearchTermView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/paid_organic_search_term_view.proto", fileDescriptor_9878fc316cec3005)
}

var fileDescriptor_9878fc316cec3005 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x41, 0x6b, 0xfa, 0x30,
	0x1c, 0xa5, 0xfd, 0xc3, 0x1f, 0xfe, 0xf5, 0xbf, 0x8b, 0x27, 0x15, 0x11, 0xdd, 0x10, 0x3c, 0x25,
	0x60, 0x6f, 0x19, 0x3b, 0x54, 0x18, 0xc2, 0x0e, 0x53, 0x74, 0xf4, 0x30, 0x0a, 0x25, 0xb6, 0x59,
	0x16, 0xb0, 0x49, 0x49, 0x5a, 0x3d, 0x88, 0xc7, 0x1d, 0x86, 0xdf, 0x62, 0xc7, 0x7d, 0x94, 0x7d,
	0x14, 0x3f, 0xc5, 0xb0, 0x6d, 0xa2, 0x87, 0xe9, 0x6e, 0xaf, 0x7d, 0xef, 0xf7, 0x7e, 0xef, 0x25,
	0x71, 0xee, 0xa9, 0x10, 0x74, 0x49, 0x20, 0x8e, 0x15, 0x2c, 0xe1, 0x01, 0xad, 0x5c, 0x28, 0x89,
	0x12, 0xb9, 0x8c, 0x88, 0x82, 0x29, 0x66, 0x71, 0x28, 0x24, 0xc5, 0x9c, 0x45, 0xa1, 0x22, 0x58,
	0x46, 0xaf, 0x61, 0x46, 0x64, 0x12, 0xae, 0x18, 0x59, 0x83, 0x54, 0x8a, 0x4c, 0xd4, 0x7b, 0xe5,
	0x2c, 0xc0, 0xb1, 0x02, 0xc6, 0x06, 0xac, 0x5c, 0x60, 0x6c, 0x5a, 0x4d, 0xbd, 0x29, 0x65, 0xc6,
	0xbc, 0x9c, 0x6e, 0x75, 0x2a, 0xaa, 0xf8, 0x5a, 0xe4, 0x2f, 0x70, 0x2d, 0x71, 0x9a, 0x12, 0xa9,
	0x2a, 0xbe, 0x7d, 0x32, 0x8a, 0x39, 0x17, 0x19, 0xce, 0x98, 0xe0, 0x15, 0x7b, 0xbd, 0xb3, 0x9d,
	0xe6, 0x14, 0xb3, 0x78, 0x52, 0x46, 0x9c, 0x17, 0x09, 0x9f, 0x88, 0x4c, 0x7c, 0x46, 0xd6, 0xf5,
	0x1b, 0xe7, 0x4a, 0x6f, 0x0b, 0x39, 0x4e, 0x48, 0xc3, 0xea, 0x5a, 0x83, 0x7f, 0xb3, 0xff, 0xfa,
	0xe7, 0x23, 0x4e, 0x48, 0xfd, 0xce, 0xa9, 0x9d, 0x14, 0x6b, 0xd8, 0x5d, 0x6b, 0x50, 0x1b, 0xb6,
	0xab, 0x26, 0x40, 0xc7, 0x02, 0xf3, 0x4c, 0x32, 0x4e, 0x7d, 0xbc, 0xcc, 0xc9, 0xcc, 0x51, 0x66,
	0x0f, 0x7a, 0xb7, 0xf6, 0xde, 0x9b, 0xe5, 0x0c, 0x8f, 0xcd, 0x2b, 0x94, 0x32, 0x05, 0x22, 0x91,
	0xc0, 0xf3, 0xf1, 0x26, 0x51, 0xae, 0x32, 0x91, 0x10, 0xa9, 0xe0, 0x46, 0xc3, 0x6d, 0x71, 0xe4,
	0x3f, 0xea, 0x15, 0xdc, 0x5c, 0xbc, 0x8e, 0xed, 0x68, 0x67, 0x3b, 0xfd, 0x48, 0x24, 0xe0, 0xd7,
	0x0b, 0x19, 0x75, 0xce, 0xa6, 0x9a, 0x1e, 0x0a, 0x4f, 0xad, 0xe7, 0x87, 0xca, 0x84, 0x8a, 0x25,
	0xe6, 0x14, 0x08, 0x49, 0x21, 0x25, 0xbc, 0x38, 0x0e, 0x78, 0x6c, 0x78, 0xe1, 0xed, 0xdc, 0x1a,
	0xf4, 0x61, 0xff, 0x19, 0x7b, 0xde, 0xa7, 0xdd, 0x1b, 0x97, 0x96, 0x5e, 0xac, 0x40, 0x09, 0x0f,
	0xc8, 0x77, 0xc1, 0x4c, 0x2b, 0xbf, 0xb4, 0x26, 0xf0, 0x62, 0x15, 0x18, 0x4d, 0xe0, 0xbb, 0x81,
	0xd1, 0xec, 0xed, 0x7e, 0x49, 0x20, 0xe4, 0xc5, 0x0a, 0x21, 0xa3, 0x42, 0xc8, 0x77, 0x11, 0x32,
	0xba, 0xc5, 0xdf, 0x22, 0xac, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x3e, 0x5c, 0xf8, 0xe7,
	0x02, 0x00, 0x00,
}
