// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/keyword_plan_campaign.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v3/enums"
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

// A Keyword Plan campaign.
// Max number of keyword plan campaigns per plan allowed: 1.
type KeywordPlanCampaign struct {
	// The resource name of the Keyword Plan campaign.
	// KeywordPlanCampaign resource names have the form:
	//
	// `customers/{customer_id}/keywordPlanCampaigns/{kp_campaign_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The keyword plan this campaign belongs to.
	KeywordPlan *wrappers.StringValue `protobuf:"bytes,2,opt,name=keyword_plan,json=keywordPlan,proto3" json:"keyword_plan,omitempty"`
	// The ID of the Keyword Plan campaign.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the Keyword Plan campaign.
	//
	// This field is required and should not be empty when creating Keyword Plan
	// campaigns.
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The languages targeted for the Keyword Plan campaign.
	// Max allowed: 1.
	LanguageConstants []*wrappers.StringValue `protobuf:"bytes,5,rep,name=language_constants,json=languageConstants,proto3" json:"language_constants,omitempty"`
	// Targeting network.
	//
	// This field is required and should not be empty when creating Keyword Plan
	// campaigns.
	KeywordPlanNetwork enums.KeywordPlanNetworkEnum_KeywordPlanNetwork `protobuf:"varint,6,opt,name=keyword_plan_network,json=keywordPlanNetwork,proto3,enum=google.ads.googleads.v3.enums.KeywordPlanNetworkEnum_KeywordPlanNetwork" json:"keyword_plan_network,omitempty"`
	// A default max cpc bid in micros, and in the account currency, for all ad
	// groups under the campaign.
	//
	// This field is required and should not be empty when creating Keyword Plan
	// campaigns.
	CpcBidMicros *wrappers.Int64Value `protobuf:"bytes,7,opt,name=cpc_bid_micros,json=cpcBidMicros,proto3" json:"cpc_bid_micros,omitempty"`
	// The geo targets.
	// Max number allowed: 20.
	GeoTargets           []*KeywordPlanGeoTarget `protobuf:"bytes,8,rep,name=geo_targets,json=geoTargets,proto3" json:"geo_targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *KeywordPlanCampaign) Reset()         { *m = KeywordPlanCampaign{} }
func (m *KeywordPlanCampaign) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanCampaign) ProtoMessage()    {}
func (*KeywordPlanCampaign) Descriptor() ([]byte, []int) {
	return fileDescriptor_17930fe29b1b9aff, []int{0}
}

func (m *KeywordPlanCampaign) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanCampaign.Unmarshal(m, b)
}
func (m *KeywordPlanCampaign) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanCampaign.Marshal(b, m, deterministic)
}
func (m *KeywordPlanCampaign) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanCampaign.Merge(m, src)
}
func (m *KeywordPlanCampaign) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanCampaign.Size(m)
}
func (m *KeywordPlanCampaign) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanCampaign.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanCampaign proto.InternalMessageInfo

func (m *KeywordPlanCampaign) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *KeywordPlanCampaign) GetKeywordPlan() *wrappers.StringValue {
	if m != nil {
		return m.KeywordPlan
	}
	return nil
}

func (m *KeywordPlanCampaign) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *KeywordPlanCampaign) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *KeywordPlanCampaign) GetLanguageConstants() []*wrappers.StringValue {
	if m != nil {
		return m.LanguageConstants
	}
	return nil
}

func (m *KeywordPlanCampaign) GetKeywordPlanNetwork() enums.KeywordPlanNetworkEnum_KeywordPlanNetwork {
	if m != nil {
		return m.KeywordPlanNetwork
	}
	return enums.KeywordPlanNetworkEnum_UNSPECIFIED
}

func (m *KeywordPlanCampaign) GetCpcBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidMicros
	}
	return nil
}

func (m *KeywordPlanCampaign) GetGeoTargets() []*KeywordPlanGeoTarget {
	if m != nil {
		return m.GeoTargets
	}
	return nil
}

// A geo target.
// Next ID: 3
type KeywordPlanGeoTarget struct {
	// Required. The resource name of the geo target.
	GeoTargetConstant    *wrappers.StringValue `protobuf:"bytes,1,opt,name=geo_target_constant,json=geoTargetConstant,proto3" json:"geo_target_constant,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *KeywordPlanGeoTarget) Reset()         { *m = KeywordPlanGeoTarget{} }
func (m *KeywordPlanGeoTarget) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanGeoTarget) ProtoMessage()    {}
func (*KeywordPlanGeoTarget) Descriptor() ([]byte, []int) {
	return fileDescriptor_17930fe29b1b9aff, []int{1}
}

func (m *KeywordPlanGeoTarget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanGeoTarget.Unmarshal(m, b)
}
func (m *KeywordPlanGeoTarget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanGeoTarget.Marshal(b, m, deterministic)
}
func (m *KeywordPlanGeoTarget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanGeoTarget.Merge(m, src)
}
func (m *KeywordPlanGeoTarget) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanGeoTarget.Size(m)
}
func (m *KeywordPlanGeoTarget) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanGeoTarget.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanGeoTarget proto.InternalMessageInfo

func (m *KeywordPlanGeoTarget) GetGeoTargetConstant() *wrappers.StringValue {
	if m != nil {
		return m.GeoTargetConstant
	}
	return nil
}

func init() {
	proto.RegisterType((*KeywordPlanCampaign)(nil), "google.ads.googleads.v3.resources.KeywordPlanCampaign")
	proto.RegisterType((*KeywordPlanGeoTarget)(nil), "google.ads.googleads.v3.resources.KeywordPlanGeoTarget")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/keyword_plan_campaign.proto", fileDescriptor_17930fe29b1b9aff)
}

var fileDescriptor_17930fe29b1b9aff = []byte{
	// 570 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdf, 0x6a, 0x13, 0x41,
	0x14, 0xc6, 0xc9, 0xf6, 0x8f, 0x3a, 0xad, 0x05, 0xa7, 0xbd, 0x58, 0x6b, 0x91, 0xb4, 0x52, 0x08,
	0x28, 0xb3, 0xd2, 0x15, 0x95, 0x15, 0x91, 0x4d, 0x91, 0xa8, 0xd5, 0x12, 0xa2, 0x04, 0x91, 0xc0,
	0x32, 0xd9, 0x19, 0x87, 0x25, 0xd9, 0x99, 0x75, 0x66, 0xb6, 0x41, 0x4b, 0x2f, 0x7d, 0x11, 0x2f,
	0x7d, 0x00, 0x1f, 0xc2, 0x47, 0xe9, 0x53, 0x48, 0x66, 0x77, 0x36, 0xc1, 0x6c, 0x8d, 0x77, 0x27,
	0x73, 0xbe, 0xdf, 0x9c, 0xef, 0x9c, 0x33, 0x59, 0xf0, 0x9c, 0x09, 0xc1, 0xc6, 0xd4, 0xc3, 0x44,
	0x79, 0x45, 0x38, 0x8d, 0xce, 0x7c, 0x4f, 0x52, 0x25, 0x72, 0x19, 0x53, 0xe5, 0x8d, 0xe8, 0xd7,
	0x89, 0x90, 0x24, 0xca, 0xc6, 0x98, 0x47, 0x31, 0x4e, 0x33, 0x9c, 0x30, 0x8e, 0x32, 0x29, 0xb4,
	0x80, 0xfb, 0x05, 0x83, 0x30, 0x51, 0xa8, 0xc2, 0xd1, 0x99, 0x8f, 0x2a, 0x7c, 0xf7, 0xe9, 0x55,
	0x15, 0x28, 0xcf, 0xd3, 0xbf, 0x6e, 0xe7, 0x54, 0x4f, 0x84, 0x1c, 0x15, 0x97, 0xef, 0xde, 0xb6,
	0x64, 0x96, 0x54, 0x76, 0xca, 0xd4, 0xdd, 0x32, 0x65, 0x7e, 0x0d, 0xf3, 0xcf, 0xde, 0x44, 0xe2,
	0x2c, 0xa3, 0x52, 0x95, 0xf9, 0xbd, 0x39, 0x14, 0x73, 0x2e, 0x34, 0xd6, 0x89, 0xe0, 0x65, 0xf6,
	0xe0, 0xd7, 0x1a, 0xd8, 0x3e, 0x29, 0xea, 0x76, 0xc7, 0x98, 0x1f, 0x97, 0x3d, 0xc1, 0x7b, 0xe0,
	0xa6, 0xad, 0x13, 0x71, 0x9c, 0x52, 0xb7, 0xd1, 0x6c, 0xb4, 0x6e, 0xf4, 0x36, 0xed, 0xe1, 0x29,
	0x4e, 0x29, 0x7c, 0x01, 0x36, 0xe7, 0x3d, 0xbb, 0x4e, 0xb3, 0xd1, 0xda, 0x38, 0xda, 0x2b, 0xdb,
	0x47, 0xd6, 0x11, 0x7a, 0xaf, 0x65, 0xc2, 0x59, 0x1f, 0x8f, 0x73, 0xda, 0xdb, 0x18, 0xcd, 0xaa,
	0xc1, 0xfb, 0xc0, 0x49, 0x88, 0xbb, 0x62, 0xb0, 0x3b, 0x0b, 0xd8, 0x6b, 0xae, 0x1f, 0x3f, 0x2a,
	0x28, 0x27, 0x21, 0xf0, 0x21, 0x58, 0x35, 0x4e, 0x56, 0xff, 0xa3, 0x8a, 0x51, 0xc2, 0x13, 0x00,
	0xc7, 0x98, 0xb3, 0x1c, 0x33, 0x1a, 0xc5, 0x82, 0x2b, 0x8d, 0xb9, 0x56, 0xee, 0x5a, 0x73, 0x65,
	0x29, 0x7f, 0xcb, 0x72, 0xc7, 0x16, 0x83, 0xdf, 0xc0, 0x4e, 0xdd, 0x82, 0xdc, 0xf5, 0x66, 0xa3,
	0xb5, 0x75, 0xf4, 0x0a, 0x5d, 0xb5, 0x7e, 0xb3, 0x5b, 0x34, 0x37, 0xe3, 0xd3, 0x02, 0x7c, 0xc9,
	0xf3, 0xb4, 0xe6, 0xb8, 0x07, 0x47, 0x0b, 0x67, 0x30, 0x04, 0x5b, 0x71, 0x16, 0x47, 0xc3, 0x84,
	0x44, 0x69, 0x12, 0x4b, 0xa1, 0xdc, 0x6b, 0xcb, 0x67, 0xb6, 0x19, 0x67, 0x71, 0x3b, 0x21, 0xef,
	0x0c, 0x00, 0x3f, 0x82, 0x0d, 0x46, 0x45, 0xa4, 0xb1, 0x64, 0x54, 0x2b, 0xf7, 0xba, 0x19, 0xc2,
	0x13, 0xb4, 0xf4, 0xd1, 0xce, 0x5b, 0xec, 0x50, 0xf1, 0xc1, 0xf0, 0x3d, 0xc0, 0x6c, 0xa8, 0x02,
	0x7d, 0x19, 0x7e, 0x01, 0x0f, 0x66, 0x74, 0x19, 0x65, 0x89, 0x42, 0xb1, 0x48, 0xbd, 0xba, 0xd7,
	0x15, 0xc6, 0xb9, 0xd2, 0x22, 0xa5, 0x52, 0x79, 0xe7, 0x36, 0xbc, 0xf0, 0x46, 0x8b, 0x4a, 0xe5,
	0x9d, 0xd7, 0xfe, 0xe7, 0x2e, 0x0e, 0x08, 0xd8, 0xa9, 0x73, 0x06, 0xdf, 0x82, 0xed, 0x59, 0x9f,
	0xd5, 0xd6, 0xcd, 0xf3, 0x5d, 0xba, 0xf4, 0xaa, 0x29, 0xbb, 0xf5, 0xf6, 0x77, 0x07, 0x1c, 0xc6,
	0x22, 0x5d, 0x3e, 0xa6, 0xb6, 0x5b, 0xd3, 0x67, 0x77, 0x5a, 0xa5, 0xdb, 0xf8, 0xf4, 0xa6, 0xc4,
	0x99, 0x98, 0x3e, 0x2b, 0x24, 0x24, 0xf3, 0x18, 0xe5, 0xc6, 0x83, 0x37, 0x9b, 0xd6, 0x3f, 0x3e,
	0x3c, 0xcf, 0xaa, 0xe8, 0x87, 0xb3, 0xd2, 0x09, 0xc3, 0x9f, 0xce, 0x7e, 0xa7, 0xb8, 0x32, 0x24,
	0x0a, 0x15, 0xe1, 0x34, 0xea, 0xfb, 0xa8, 0x67, 0x95, 0xbf, 0xad, 0x66, 0x10, 0x12, 0x35, 0xa8,
	0x34, 0x83, 0xbe, 0x3f, 0xa8, 0x34, 0x97, 0xce, 0x61, 0x91, 0x08, 0x82, 0x90, 0xa8, 0x20, 0xa8,
	0x54, 0x41, 0xd0, 0xf7, 0x83, 0xa0, 0xd2, 0x0d, 0xd7, 0x8d, 0x59, 0xff, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xd8, 0xd5, 0x7a, 0xf9, 0x24, 0x05, 0x00, 0x00,
}
