// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/ad_group_ad.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import common "google.golang.org/genproto/googleapis/ads/googleads/v0/common"
import enums "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// An ad group ad.
type AdGroupAd struct {
	// The resource name of the ad.
	// Ad group ad resource names have the form:
	//
	// `customers/{customer_id}/adGroupAds/{ad_group_id}_{ad_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The status of the ad.
	Status enums.AdGroupAdStatusEnum_AdGroupAdStatus `protobuf:"varint,3,opt,name=status,proto3,enum=google.ads.googleads.v0.enums.AdGroupAdStatusEnum_AdGroupAdStatus" json:"status,omitempty"`
	// The ad group to which the ad belongs.
	AdGroup *wrappers.StringValue `protobuf:"bytes,4,opt,name=ad_group,json=adGroup,proto3" json:"ad_group,omitempty"`
	// The ad.
	Ad *Ad `protobuf:"bytes,5,opt,name=ad,proto3" json:"ad,omitempty"`
	// Policy information for the ad.
	PolicySummary        *AdGroupAdPolicySummary `protobuf:"bytes,6,opt,name=policy_summary,json=policySummary,proto3" json:"policy_summary,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *AdGroupAd) Reset()         { *m = AdGroupAd{} }
func (m *AdGroupAd) String() string { return proto.CompactTextString(m) }
func (*AdGroupAd) ProtoMessage()    {}
func (*AdGroupAd) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_ad_066bc07cfa249cf3, []int{0}
}
func (m *AdGroupAd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupAd.Unmarshal(m, b)
}
func (m *AdGroupAd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupAd.Marshal(b, m, deterministic)
}
func (dst *AdGroupAd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupAd.Merge(dst, src)
}
func (m *AdGroupAd) XXX_Size() int {
	return xxx_messageInfo_AdGroupAd.Size(m)
}
func (m *AdGroupAd) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupAd.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupAd proto.InternalMessageInfo

func (m *AdGroupAd) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AdGroupAd) GetStatus() enums.AdGroupAdStatusEnum_AdGroupAdStatus {
	if m != nil {
		return m.Status
	}
	return enums.AdGroupAdStatusEnum_UNSPECIFIED
}

func (m *AdGroupAd) GetAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.AdGroup
	}
	return nil
}

func (m *AdGroupAd) GetAd() *Ad {
	if m != nil {
		return m.Ad
	}
	return nil
}

func (m *AdGroupAd) GetPolicySummary() *AdGroupAdPolicySummary {
	if m != nil {
		return m.PolicySummary
	}
	return nil
}

// Contains policy information for an ad.
type AdGroupAdPolicySummary struct {
	// The list of policy findings for this ad.
	PolicyTopicEntries []*common.PolicyTopicEntry `protobuf:"bytes,1,rep,name=policy_topic_entries,json=policyTopicEntries,proto3" json:"policy_topic_entries,omitempty"`
	// Where in the review process this ad is.
	ReviewStatus enums.PolicyReviewStatusEnum_PolicyReviewStatus `protobuf:"varint,2,opt,name=review_status,json=reviewStatus,proto3,enum=google.ads.googleads.v0.enums.PolicyReviewStatusEnum_PolicyReviewStatus" json:"review_status,omitempty"`
	// The overall approval status of this ad, calculated based on the status of
	// its individual policy topic entries.
	ApprovalStatus       enums.PolicyApprovalStatusEnum_PolicyApprovalStatus `protobuf:"varint,3,opt,name=approval_status,json=approvalStatus,proto3,enum=google.ads.googleads.v0.enums.PolicyApprovalStatusEnum_PolicyApprovalStatus" json:"approval_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                            `json:"-"`
	XXX_unrecognized     []byte                                              `json:"-"`
	XXX_sizecache        int32                                               `json:"-"`
}

func (m *AdGroupAdPolicySummary) Reset()         { *m = AdGroupAdPolicySummary{} }
func (m *AdGroupAdPolicySummary) String() string { return proto.CompactTextString(m) }
func (*AdGroupAdPolicySummary) ProtoMessage()    {}
func (*AdGroupAdPolicySummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_ad_066bc07cfa249cf3, []int{1}
}
func (m *AdGroupAdPolicySummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupAdPolicySummary.Unmarshal(m, b)
}
func (m *AdGroupAdPolicySummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupAdPolicySummary.Marshal(b, m, deterministic)
}
func (dst *AdGroupAdPolicySummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupAdPolicySummary.Merge(dst, src)
}
func (m *AdGroupAdPolicySummary) XXX_Size() int {
	return xxx_messageInfo_AdGroupAdPolicySummary.Size(m)
}
func (m *AdGroupAdPolicySummary) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupAdPolicySummary.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupAdPolicySummary proto.InternalMessageInfo

func (m *AdGroupAdPolicySummary) GetPolicyTopicEntries() []*common.PolicyTopicEntry {
	if m != nil {
		return m.PolicyTopicEntries
	}
	return nil
}

func (m *AdGroupAdPolicySummary) GetReviewStatus() enums.PolicyReviewStatusEnum_PolicyReviewStatus {
	if m != nil {
		return m.ReviewStatus
	}
	return enums.PolicyReviewStatusEnum_UNSPECIFIED
}

func (m *AdGroupAdPolicySummary) GetApprovalStatus() enums.PolicyApprovalStatusEnum_PolicyApprovalStatus {
	if m != nil {
		return m.ApprovalStatus
	}
	return enums.PolicyApprovalStatusEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*AdGroupAd)(nil), "google.ads.googleads.v0.resources.AdGroupAd")
	proto.RegisterType((*AdGroupAdPolicySummary)(nil), "google.ads.googleads.v0.resources.AdGroupAdPolicySummary")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/ad_group_ad.proto", fileDescriptor_ad_group_ad_066bc07cfa249cf3)
}

var fileDescriptor_ad_group_ad_066bc07cfa249cf3 = []byte{
	// 544 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xdf, 0x8a, 0xd3, 0x4e,
	0x14, 0x26, 0xe9, 0xef, 0x57, 0xdd, 0x59, 0xb7, 0x42, 0x10, 0x09, 0x8b, 0x48, 0x77, 0xa5, 0x50,
	0x14, 0x26, 0xa1, 0x8b, 0xff, 0xe2, 0x55, 0x0a, 0x4b, 0x45, 0x44, 0x96, 0x54, 0x7a, 0xb1, 0x14,
	0xe2, 0xb4, 0x33, 0x86, 0x40, 0x26, 0x33, 0xcc, 0x24, 0x5d, 0xfa, 0x3a, 0x5e, 0xfa, 0x28, 0x3e,
	0x84, 0x17, 0x5e, 0xfa, 0x0a, 0xde, 0x48, 0x66, 0x26, 0xd9, 0x66, 0xd7, 0x6e, 0x7b, 0x77, 0xce,
	0xc9, 0xf7, 0x7d, 0xe7, 0x9c, 0x39, 0x5f, 0xc0, 0x59, 0xc2, 0x58, 0x92, 0x11, 0x0f, 0x61, 0xe9,
	0xe9, 0xb0, 0x8a, 0x56, 0xbe, 0x27, 0x88, 0x64, 0xa5, 0x58, 0x12, 0xe9, 0x21, 0x1c, 0x27, 0x82,
	0x95, 0x3c, 0x46, 0x18, 0x72, 0xc1, 0x0a, 0xe6, 0x9c, 0x68, 0x24, 0x44, 0x58, 0xc2, 0x86, 0x04,
	0x57, 0x3e, 0x6c, 0x48, 0xc7, 0x2f, 0xb6, 0xe9, 0x2e, 0x19, 0xa5, 0x2c, 0xf7, 0x38, 0xcb, 0xd2,
	0xe5, 0x5a, 0xeb, 0x1d, 0xbf, 0xda, 0x06, 0x26, 0x79, 0x49, 0x5b, 0x03, 0xc4, 0xb2, 0x40, 0x45,
	0x29, 0x0d, 0x2f, 0xb8, 0x9b, 0xa7, 0x7b, 0xc4, 0x88, 0x73, 0xc1, 0x56, 0x28, 0x6b, 0x73, 0xdf,
	0xec, 0xc5, 0x15, 0x64, 0x95, 0x92, 0xab, 0x36, 0xf3, 0xf9, 0x3e, 0x4f, 0x66, 0xb0, 0x4f, 0x0d,
	0x56, 0x65, 0x8b, 0xf2, 0xab, 0x77, 0x25, 0x10, 0xe7, 0x44, 0x18, 0xad, 0xd3, 0x9f, 0x36, 0x38,
	0x08, 0xf1, 0xa4, 0xda, 0x2e, 0xc4, 0xce, 0x33, 0x70, 0x54, 0x6b, 0xc4, 0x39, 0xa2, 0xc4, 0xb5,
	0xfa, 0xd6, 0xf0, 0x20, 0x7a, 0x50, 0x17, 0x3f, 0x21, 0x4a, 0x9c, 0x4b, 0xd0, 0xd5, 0xe3, 0xb8,
	0x9d, 0xbe, 0x35, 0xec, 0x8d, 0xc6, 0x70, 0xdb, 0x35, 0xd4, 0x26, 0xb0, 0x91, 0x9f, 0x2a, 0xd6,
	0x79, 0x5e, 0xd2, 0x9b, 0xb5, 0xc8, 0x28, 0x3a, 0xaf, 0xc1, 0xfd, 0xfa, 0xb1, 0xdd, 0xff, 0xfa,
	0xd6, 0xf0, 0x70, 0xf4, 0xa4, 0x56, 0xaf, 0x37, 0x80, 0xd3, 0x42, 0xa4, 0x79, 0x32, 0x43, 0x59,
	0x49, 0xa2, 0x7b, 0x48, 0x0b, 0x39, 0x2f, 0x81, 0x8d, 0xb0, 0xfb, 0xbf, 0xa2, 0x0c, 0xe0, 0x4e,
	0x7b, 0xc0, 0x10, 0x47, 0x36, 0xc2, 0xce, 0x17, 0xd0, 0x33, 0x0f, 0x2d, 0x4b, 0x4a, 0x91, 0x58,
	0xbb, 0x5d, 0x25, 0xf1, 0x76, 0x2f, 0x09, 0xb3, 0xc3, 0x85, 0x52, 0x98, 0x6a, 0x81, 0xe8, 0x88,
	0x6f, 0xa6, 0xa7, 0xbf, 0x6c, 0xf0, 0xf8, 0xdf, 0x48, 0x67, 0x01, 0x1e, 0x99, 0xe6, 0x05, 0xe3,
	0xe9, 0x32, 0x26, 0x79, 0x21, 0x52, 0x22, 0x5d, 0xab, 0xdf, 0x19, 0x1e, 0x8e, 0xfc, 0xad, 0x23,
	0x68, 0x07, 0x43, 0x2d, 0xf6, 0xb9, 0xa2, 0x9e, 0xe7, 0x85, 0x58, 0x47, 0x0e, 0x6f, 0x57, 0x52,
	0x22, 0x1d, 0x5a, 0x5d, 0x74, 0xc3, 0x42, 0xae, 0xad, 0x6e, 0xf6, 0x7e, 0xc7, 0xcd, 0xb4, 0x76,
	0xa4, 0x98, 0x1b, 0x67, 0xbb, 0x5d, 0xae, 0xbc, 0x71, 0x9d, 0x39, 0x25, 0x78, 0x78, 0xc3, 0xed,
	0xc6, 0x24, 0x1f, 0xf7, 0x6a, 0x18, 0x1a, 0xee, 0xad, 0x96, 0xed, 0x0f, 0x51, 0x0f, 0xb5, 0xf2,
	0xf1, 0x1f, 0x0b, 0x0c, 0x96, 0x8c, 0xee, 0x3e, 0xda, 0xb8, 0x77, 0x7d, 0x8b, 0xca, 0x4f, 0x17,
	0xd6, 0xe5, 0x07, 0x43, 0x4a, 0x58, 0x86, 0xf2, 0x04, 0x32, 0x91, 0x78, 0x09, 0xc9, 0x95, 0xdb,
	0xea, 0xbf, 0x8b, 0xa7, 0xf2, 0x8e, 0x9f, 0xed, 0x5d, 0x13, 0x7d, 0xb3, 0x3b, 0x93, 0x30, 0xfc,
	0x6e, 0x9f, 0x4c, 0xb4, 0x64, 0x88, 0x25, 0xd4, 0x61, 0x15, 0xcd, 0x7c, 0x18, 0xd5, 0xc8, 0x1f,
	0x35, 0x66, 0x1e, 0x62, 0x39, 0x6f, 0x30, 0xf3, 0x99, 0x3f, 0x6f, 0x30, 0xbf, 0xed, 0x81, 0xfe,
	0x10, 0x04, 0x21, 0x96, 0x41, 0xd0, 0xa0, 0x82, 0x60, 0xe6, 0x07, 0x41, 0x83, 0x5b, 0x74, 0xd5,
	0xb0, 0x67, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x16, 0x68, 0x0e, 0xf9, 0x4b, 0x05, 0x00, 0x00,
}