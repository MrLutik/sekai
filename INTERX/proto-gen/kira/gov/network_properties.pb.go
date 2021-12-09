// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kira/gov/network_properties.proto

package gov

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type NetworkProperty int32

const (
	NetworkProperty_MIN_TX_FEE                     NetworkProperty = 0
	NetworkProperty_MAX_TX_FEE                     NetworkProperty = 1
	NetworkProperty_VOTE_QUORUM                    NetworkProperty = 2
	NetworkProperty_DEFAULT_PROPOSAL_END_TIME      NetworkProperty = 3
	NetworkProperty_MINIMUM_PROPOSAL_END_TIME      NetworkProperty = 4
	NetworkProperty_PROPOSAL_ENACTMENT_TIME        NetworkProperty = 5
	NetworkProperty_MIN_PROPOSAL_END_BLOCKS        NetworkProperty = 6
	NetworkProperty_MIN_PROPOSAL_ENACTMENT_BLOCKS  NetworkProperty = 7
	NetworkProperty_ENABLE_FOREIGN_FEE_PAYMENTS    NetworkProperty = 8
	NetworkProperty_MISCHANCE_RANK_DECREASE_AMOUNT NetworkProperty = 9
	NetworkProperty_MAX_MISCHANCE                  NetworkProperty = 10
	NetworkProperty_MISCHANCE_CONFIDENCE           NetworkProperty = 11
	NetworkProperty_INACTIVE_RANK_DECREASE_PERCENT NetworkProperty = 12
	NetworkProperty_POOR_NETWORK_MAX_BANK_SEND     NetworkProperty = 13
	NetworkProperty_MIN_VALIDATORS                 NetworkProperty = 14
	NetworkProperty_JAIL_MAX_TIME                  NetworkProperty = 15
	NetworkProperty_ENABLE_TOKEN_WHITELIST         NetworkProperty = 16
	NetworkProperty_ENABLE_TOKEN_BLACKLIST         NetworkProperty = 17
	NetworkProperty_MIN_IDENTITY_APPROVAL_TIP      NetworkProperty = 18
	NetworkProperty_UNIQUE_IDENTITY_KEYS           NetworkProperty = 19
)

var NetworkProperty_name = map[int32]string{
	0:  "MIN_TX_FEE",
	1:  "MAX_TX_FEE",
	2:  "VOTE_QUORUM",
	3:  "DEFAULT_PROPOSAL_END_TIME",
	4:  "MINIMUM_PROPOSAL_END_TIME",
	5:  "PROPOSAL_ENACTMENT_TIME",
	6:  "MIN_PROPOSAL_END_BLOCKS",
	7:  "MIN_PROPOSAL_ENACTMENT_BLOCKS",
	8:  "ENABLE_FOREIGN_FEE_PAYMENTS",
	9:  "MISCHANCE_RANK_DECREASE_AMOUNT",
	10: "MAX_MISCHANCE",
	11: "MISCHANCE_CONFIDENCE",
	12: "INACTIVE_RANK_DECREASE_PERCENT",
	13: "POOR_NETWORK_MAX_BANK_SEND",
	14: "MIN_VALIDATORS",
	15: "JAIL_MAX_TIME",
	16: "ENABLE_TOKEN_WHITELIST",
	17: "ENABLE_TOKEN_BLACKLIST",
	18: "MIN_IDENTITY_APPROVAL_TIP",
	19: "UNIQUE_IDENTITY_KEYS",
}

var NetworkProperty_value = map[string]int32{
	"MIN_TX_FEE":                     0,
	"MAX_TX_FEE":                     1,
	"VOTE_QUORUM":                    2,
	"DEFAULT_PROPOSAL_END_TIME":      3,
	"MINIMUM_PROPOSAL_END_TIME":      4,
	"PROPOSAL_ENACTMENT_TIME":        5,
	"MIN_PROPOSAL_END_BLOCKS":        6,
	"MIN_PROPOSAL_ENACTMENT_BLOCKS":  7,
	"ENABLE_FOREIGN_FEE_PAYMENTS":    8,
	"MISCHANCE_RANK_DECREASE_AMOUNT": 9,
	"MAX_MISCHANCE":                  10,
	"MISCHANCE_CONFIDENCE":           11,
	"INACTIVE_RANK_DECREASE_PERCENT": 12,
	"POOR_NETWORK_MAX_BANK_SEND":     13,
	"MIN_VALIDATORS":                 14,
	"JAIL_MAX_TIME":                  15,
	"ENABLE_TOKEN_WHITELIST":         16,
	"ENABLE_TOKEN_BLACKLIST":         17,
	"MIN_IDENTITY_APPROVAL_TIP":      18,
	"UNIQUE_IDENTITY_KEYS":           19,
}

func (x NetworkProperty) String() string {
	return proto.EnumName(NetworkProperty_name, int32(x))
}

func (NetworkProperty) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_98011a6048e5dde3, []int{0}
}

type MsgSetNetworkProperties struct {
	NetworkProperties    *NetworkProperties `protobuf:"bytes,1,opt,name=network_properties,json=networkProperties,proto3" json:"network_properties,omitempty"`
	Proposer             []byte             `protobuf:"bytes,2,opt,name=proposer,proto3" json:"proposer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MsgSetNetworkProperties) Reset()         { *m = MsgSetNetworkProperties{} }
func (m *MsgSetNetworkProperties) String() string { return proto.CompactTextString(m) }
func (*MsgSetNetworkProperties) ProtoMessage()    {}
func (*MsgSetNetworkProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_98011a6048e5dde3, []int{0}
}

func (m *MsgSetNetworkProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgSetNetworkProperties.Unmarshal(m, b)
}
func (m *MsgSetNetworkProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgSetNetworkProperties.Marshal(b, m, deterministic)
}
func (m *MsgSetNetworkProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetNetworkProperties.Merge(m, src)
}
func (m *MsgSetNetworkProperties) XXX_Size() int {
	return xxx_messageInfo_MsgSetNetworkProperties.Size(m)
}
func (m *MsgSetNetworkProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetNetworkProperties.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetNetworkProperties proto.InternalMessageInfo

func (m *MsgSetNetworkProperties) GetNetworkProperties() *NetworkProperties {
	if m != nil {
		return m.NetworkProperties
	}
	return nil
}

func (m *MsgSetNetworkProperties) GetProposer() []byte {
	if m != nil {
		return m.Proposer
	}
	return nil
}

type NetworkPropertyValue struct {
	Value                uint64   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	StrValue             string   `protobuf:"bytes,2,opt,name=str_value,json=strValue,proto3" json:"str_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkPropertyValue) Reset()         { *m = NetworkPropertyValue{} }
func (m *NetworkPropertyValue) String() string { return proto.CompactTextString(m) }
func (*NetworkPropertyValue) ProtoMessage()    {}
func (*NetworkPropertyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_98011a6048e5dde3, []int{1}
}

func (m *NetworkPropertyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkPropertyValue.Unmarshal(m, b)
}
func (m *NetworkPropertyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkPropertyValue.Marshal(b, m, deterministic)
}
func (m *NetworkPropertyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkPropertyValue.Merge(m, src)
}
func (m *NetworkPropertyValue) XXX_Size() int {
	return xxx_messageInfo_NetworkPropertyValue.Size(m)
}
func (m *NetworkPropertyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkPropertyValue.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkPropertyValue proto.InternalMessageInfo

func (m *NetworkPropertyValue) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *NetworkPropertyValue) GetStrValue() string {
	if m != nil {
		return m.StrValue
	}
	return ""
}

type NetworkProperties struct {
	MinTxFee                    uint64   `protobuf:"varint,1,opt,name=min_tx_fee,json=minTxFee,proto3" json:"min_tx_fee,omitempty"`
	MaxTxFee                    uint64   `protobuf:"varint,2,opt,name=max_tx_fee,json=maxTxFee,proto3" json:"max_tx_fee,omitempty"`
	VoteQuorum                  uint64   `protobuf:"varint,3,opt,name=vote_quorum,json=voteQuorum,proto3" json:"vote_quorum,omitempty"`
	DefaultProposalEndTime      uint64   `protobuf:"varint,4,opt,name=default_proposal_end_time,json=defaultProposalEndTime,proto3" json:"default_proposal_end_time,omitempty"`
	MinimumProposalEndTime      uint64   `protobuf:"varint,5,opt,name=minimum_proposal_end_time,json=minimumProposalEndTime,proto3" json:"minimum_proposal_end_time,omitempty"`
	ProposalEnactmentTime       uint64   `protobuf:"varint,6,opt,name=proposal_enactment_time,json=proposalEnactmentTime,proto3" json:"proposal_enactment_time,omitempty"`
	MinProposalEndBlocks        uint64   `protobuf:"varint,7,opt,name=min_proposal_end_blocks,json=minProposalEndBlocks,proto3" json:"min_proposal_end_blocks,omitempty"`
	MinProposalEnactmentBlocks  uint64   `protobuf:"varint,8,opt,name=min_proposal_enactment_blocks,json=minProposalEnactmentBlocks,proto3" json:"min_proposal_enactment_blocks,omitempty"`
	EnableForeignFeePayments    bool     `protobuf:"varint,9,opt,name=enable_foreign_fee_payments,json=enableForeignFeePayments,proto3" json:"enable_foreign_fee_payments,omitempty"`
	MischanceRankDecreaseAmount uint64   `protobuf:"varint,10,opt,name=mischance_rank_decrease_amount,json=mischanceRankDecreaseAmount,proto3" json:"mischance_rank_decrease_amount,omitempty"`
	MaxMischance                uint64   `protobuf:"varint,11,opt,name=max_mischance,json=maxMischance,proto3" json:"max_mischance,omitempty"`
	MischanceConfidence         uint64   `protobuf:"varint,12,opt,name=mischance_confidence,json=mischanceConfidence,proto3" json:"mischance_confidence,omitempty"`
	InactiveRankDecreasePercent uint64   `protobuf:"varint,13,opt,name=inactive_rank_decrease_percent,json=inactiveRankDecreasePercent,proto3" json:"inactive_rank_decrease_percent,omitempty"`
	MinValidators               uint64   `protobuf:"varint,14,opt,name=min_validators,json=minValidators,proto3" json:"min_validators,omitempty"`
	PoorNetworkMaxBankSend      uint64   `protobuf:"varint,15,opt,name=poor_network_max_bank_send,json=poorNetworkMaxBankSend,proto3" json:"poor_network_max_bank_send,omitempty"`
	JailMaxTime                 uint64   `protobuf:"varint,16,opt,name=jail_max_time,json=jailMaxTime,proto3" json:"jail_max_time,omitempty"`
	EnableTokenWhitelist        bool     `protobuf:"varint,17,opt,name=enable_token_whitelist,json=enableTokenWhitelist,proto3" json:"enable_token_whitelist,omitempty"`
	EnableTokenBlacklist        bool     `protobuf:"varint,18,opt,name=enable_token_blacklist,json=enableTokenBlacklist,proto3" json:"enable_token_blacklist,omitempty"`
	MinIdentityApprovalTip      uint64   `protobuf:"varint,19,opt,name=min_identity_approval_tip,json=minIdentityApprovalTip,proto3" json:"min_identity_approval_tip,omitempty"`
	UniqueIdentityKeys          string   `protobuf:"bytes,20,opt,name=unique_identity_keys,json=uniqueIdentityKeys,proto3" json:"unique_identity_keys,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *NetworkProperties) Reset()         { *m = NetworkProperties{} }
func (m *NetworkProperties) String() string { return proto.CompactTextString(m) }
func (*NetworkProperties) ProtoMessage()    {}
func (*NetworkProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_98011a6048e5dde3, []int{2}
}

func (m *NetworkProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkProperties.Unmarshal(m, b)
}
func (m *NetworkProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkProperties.Marshal(b, m, deterministic)
}
func (m *NetworkProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkProperties.Merge(m, src)
}
func (m *NetworkProperties) XXX_Size() int {
	return xxx_messageInfo_NetworkProperties.Size(m)
}
func (m *NetworkProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkProperties.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkProperties proto.InternalMessageInfo

func (m *NetworkProperties) GetMinTxFee() uint64 {
	if m != nil {
		return m.MinTxFee
	}
	return 0
}

func (m *NetworkProperties) GetMaxTxFee() uint64 {
	if m != nil {
		return m.MaxTxFee
	}
	return 0
}

func (m *NetworkProperties) GetVoteQuorum() uint64 {
	if m != nil {
		return m.VoteQuorum
	}
	return 0
}

func (m *NetworkProperties) GetDefaultProposalEndTime() uint64 {
	if m != nil {
		return m.DefaultProposalEndTime
	}
	return 0
}

func (m *NetworkProperties) GetMinimumProposalEndTime() uint64 {
	if m != nil {
		return m.MinimumProposalEndTime
	}
	return 0
}

func (m *NetworkProperties) GetProposalEnactmentTime() uint64 {
	if m != nil {
		return m.ProposalEnactmentTime
	}
	return 0
}

func (m *NetworkProperties) GetMinProposalEndBlocks() uint64 {
	if m != nil {
		return m.MinProposalEndBlocks
	}
	return 0
}

func (m *NetworkProperties) GetMinProposalEnactmentBlocks() uint64 {
	if m != nil {
		return m.MinProposalEnactmentBlocks
	}
	return 0
}

func (m *NetworkProperties) GetEnableForeignFeePayments() bool {
	if m != nil {
		return m.EnableForeignFeePayments
	}
	return false
}

func (m *NetworkProperties) GetMischanceRankDecreaseAmount() uint64 {
	if m != nil {
		return m.MischanceRankDecreaseAmount
	}
	return 0
}

func (m *NetworkProperties) GetMaxMischance() uint64 {
	if m != nil {
		return m.MaxMischance
	}
	return 0
}

func (m *NetworkProperties) GetMischanceConfidence() uint64 {
	if m != nil {
		return m.MischanceConfidence
	}
	return 0
}

func (m *NetworkProperties) GetInactiveRankDecreasePercent() uint64 {
	if m != nil {
		return m.InactiveRankDecreasePercent
	}
	return 0
}

func (m *NetworkProperties) GetMinValidators() uint64 {
	if m != nil {
		return m.MinValidators
	}
	return 0
}

func (m *NetworkProperties) GetPoorNetworkMaxBankSend() uint64 {
	if m != nil {
		return m.PoorNetworkMaxBankSend
	}
	return 0
}

func (m *NetworkProperties) GetJailMaxTime() uint64 {
	if m != nil {
		return m.JailMaxTime
	}
	return 0
}

func (m *NetworkProperties) GetEnableTokenWhitelist() bool {
	if m != nil {
		return m.EnableTokenWhitelist
	}
	return false
}

func (m *NetworkProperties) GetEnableTokenBlacklist() bool {
	if m != nil {
		return m.EnableTokenBlacklist
	}
	return false
}

func (m *NetworkProperties) GetMinIdentityApprovalTip() uint64 {
	if m != nil {
		return m.MinIdentityApprovalTip
	}
	return 0
}

func (m *NetworkProperties) GetUniqueIdentityKeys() string {
	if m != nil {
		return m.UniqueIdentityKeys
	}
	return ""
}

func init() {
	proto.RegisterEnum("kira.gov.NetworkProperty", NetworkProperty_name, NetworkProperty_value)
	proto.RegisterType((*MsgSetNetworkProperties)(nil), "kira.gov.MsgSetNetworkProperties")
	proto.RegisterType((*NetworkPropertyValue)(nil), "kira.gov.NetworkPropertyValue")
	proto.RegisterType((*NetworkProperties)(nil), "kira.gov.NetworkProperties")
}

func init() {
	proto.RegisterFile("kira/gov/network_properties.proto", fileDescriptor_98011a6048e5dde3)
}

var fileDescriptor_98011a6048e5dde3 = []byte{
	// 1213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x96, 0xcb, 0x6e, 0xdb, 0x46,
	0x14, 0x86, 0xa3, 0xc4, 0x49, 0xe4, 0xb1, 0x65, 0xd3, 0xb4, 0x62, 0x33, 0x74, 0x1a, 0xb3, 0x2e,
	0x02, 0x04, 0x05, 0x62, 0x35, 0x4d, 0x53, 0xa0, 0x01, 0xba, 0xa0, 0xa4, 0x51, 0xc3, 0x48, 0xbc,
	0x84, 0xa2, 0x94, 0xa4, 0x9b, 0xc1, 0x58, 0x1a, 0xcb, 0x53, 0x89, 0xa4, 0x42, 0x52, 0x8a, 0xfc,
	0x06, 0x05, 0xdf, 0xa0, 0x0b, 0x02, 0x05, 0xfa, 0x0a, 0x7d, 0x9e, 0xae, 0xfb, 0x0c, 0x5d, 0x15,
	0x33, 0x24, 0x65, 0x5d, 0x6c, 0xaf, 0x2c, 0xcf, 0xff, 0x7f, 0xe7, 0x1c, 0xcd, 0xe5, 0x87, 0xc0,
	0xd7, 0x43, 0x1a, 0xe0, 0xca, 0xc0, 0x9f, 0x56, 0x3c, 0x12, 0x7d, 0xf1, 0x83, 0x21, 0x1a, 0x07,
	0xfe, 0x98, 0x04, 0x11, 0x25, 0xe1, 0xe9, 0x38, 0xf0, 0x23, 0x5f, 0x2c, 0x32, 0xcb, 0xe9, 0xc0,
	0x9f, 0xca, 0xe5, 0x81, 0x3f, 0xf0, 0xf9, 0x62, 0x85, 0x7d, 0x4a, 0xf5, 0x93, 0xbf, 0x0b, 0xe0,
	0x50, 0x0f, 0x07, 0x6d, 0x12, 0x19, 0x69, 0x09, 0x6b, 0x5e, 0x41, 0x7c, 0x07, 0xc4, 0xf5, 0xba,
	0x52, 0x41, 0x29, 0x3c, 0xdf, 0xfa, 0xfe, 0xe8, 0x34, 0x2f, 0x7c, 0xba, 0x06, 0xda, 0x7b, 0xde,
	0x5a, 0x2d, 0x1d, 0x14, 0x59, 0x0d, 0x3f, 0x24, 0x81, 0x74, 0x57, 0x29, 0x3c, 0xdf, 0xae, 0xbe,
	0xfc, 0xef, 0x9f, 0xe3, 0x17, 0x03, 0x1a, 0x5d, 0x4c, 0xce, 0x4e, 0x7b, 0xbe, 0x5b, 0xe9, 0xf9,
	0xa1, 0xeb, 0x87, 0xd9, 0x9f, 0x17, 0x61, 0x7f, 0x58, 0x89, 0x2e, 0xc7, 0x24, 0x3c, 0x55, 0x7b,
	0x3d, 0xb5, 0xdf, 0x0f, 0x48, 0x18, 0xda, 0xf3, 0x12, 0x27, 0x26, 0x28, 0x2f, 0xb7, 0xbd, 0xec,
	0xe2, 0xd1, 0x84, 0x88, 0x65, 0x70, 0x7f, 0xca, 0x3e, 0xf0, 0x29, 0x37, 0xec, 0xf4, 0x1f, 0xf1,
	0x08, 0x6c, 0x86, 0x51, 0x80, 0x52, 0x85, 0x75, 0xdf, 0xb4, 0x8b, 0x61, 0x14, 0x70, 0xe4, 0xcd,
	0xc6, 0xbf, 0x7f, 0x1e, 0x17, 0x4e, 0xfe, 0x28, 0x82, 0xbd, 0xf5, 0x1d, 0x78, 0x02, 0x80, 0x4b,
	0x3d, 0x14, 0xcd, 0xd0, 0x39, 0xc9, 0x6b, 0x16, 0x5d, 0xea, 0x39, 0xb3, 0x06, 0x21, 0x5c, 0xc5,
	0xb3, 0x5c, 0xbd, 0x9b, 0xa9, 0x78, 0x96, 0xaa, 0xc7, 0x60, 0x6b, 0xea, 0x47, 0x04, 0x7d, 0x9e,
	0xf8, 0xc1, 0xc4, 0x95, 0xee, 0x71, 0x19, 0xb0, 0xa5, 0xf7, 0x7c, 0x45, 0xfc, 0x09, 0x3c, 0xee,
	0x93, 0x73, 0x3c, 0x19, 0x45, 0x28, 0xfd, 0x5e, 0x78, 0x84, 0x88, 0xd7, 0x47, 0x11, 0x75, 0x89,
	0xb4, 0xc1, 0xed, 0x07, 0x99, 0xc1, 0xca, 0x74, 0xe8, 0xf5, 0x1d, 0xea, 0x12, 0x86, 0xba, 0xd4,
	0xa3, 0xee, 0xc4, 0xbd, 0x06, 0xbd, 0x9f, 0xa2, 0x99, 0x61, 0x15, 0xfd, 0x11, 0x1c, 0x2e, 0x20,
	0xb8, 0x17, 0xb9, 0xc4, 0x8b, 0x52, 0xf0, 0x01, 0x07, 0x1f, 0x8d, 0xe7, 0x44, 0xa6, 0x72, 0xee,
	0x35, 0x38, 0x64, 0x5b, 0xb1, 0xd4, 0xee, 0x6c, 0xe4, 0xf7, 0x86, 0xa1, 0xf4, 0x90, 0x73, 0x65,
	0x97, 0x7a, 0x0b, 0xcd, 0xaa, 0x5c, 0x13, 0x55, 0xf0, 0xd5, 0x0a, 0x96, 0xb7, 0xcc, 0xe0, 0x22,
	0x87, 0xe5, 0x25, 0x38, 0xb3, 0x64, 0x25, 0x7e, 0x06, 0x47, 0xc4, 0xc3, 0x67, 0x23, 0x82, 0xce,
	0xfd, 0x80, 0xd0, 0x81, 0xc7, 0xb6, 0x1b, 0x8d, 0xf1, 0x25, 0xf3, 0x84, 0xd2, 0xa6, 0x52, 0x78,
	0x5e, 0xb4, 0xa5, 0xd4, 0xd2, 0x48, 0x1d, 0x0d, 0x42, 0xac, 0x4c, 0x17, 0x6b, 0xe0, 0xa9, 0x4b,
	0xc3, 0xde, 0x05, 0xf6, 0x7a, 0x04, 0x05, 0xd8, 0x1b, 0xa2, 0x3e, 0xe9, 0x05, 0x04, 0x87, 0x04,
	0x61, 0xd7, 0x9f, 0x78, 0x91, 0x04, 0xf8, 0x08, 0x47, 0x73, 0x97, 0x8d, 0xbd, 0x61, 0x3d, 0xf3,
	0xa8, 0xdc, 0x22, 0x7e, 0x03, 0x4a, 0xec, 0xa8, 0xe7, 0x16, 0x69, 0x8b, 0x33, 0xdb, 0x2e, 0x9e,
	0xe9, 0xf9, 0x9a, 0xf8, 0x12, 0x94, 0xaf, 0x3a, 0xf5, 0x7c, 0xef, 0x9c, 0xf6, 0x09, 0xf3, 0x6e,
	0x73, 0xef, 0xfe, 0x5c, 0xab, 0xcd, 0x25, 0x36, 0x1c, 0x65, 0x5f, 0x97, 0x4e, 0x57, 0x67, 0x1b,
	0x93, 0xa0, 0x47, 0xbc, 0x48, 0x2a, 0xa5, 0xc3, 0xe5, 0xae, 0xc5, 0xd9, 0xac, 0xd4, 0x22, 0x3e,
	0x03, 0x3b, 0x6c, 0x8f, 0xa7, 0x78, 0x44, 0xfb, 0x38, 0xf2, 0x83, 0x50, 0xda, 0xe1, 0x50, 0xc9,
	0xa5, 0x5e, 0x77, 0xbe, 0x28, 0xbe, 0x01, 0xf2, 0xd8, 0xf7, 0x03, 0x94, 0xbf, 0x69, 0xf6, 0x85,
	0xce, 0x58, 0xcf, 0x90, 0x78, 0x7d, 0x69, 0x37, 0xbd, 0x35, 0xcc, 0x91, 0xbd, 0x03, 0x1d, 0xcf,
	0xaa, 0xd8, 0x1b, 0xb6, 0x89, 0xd7, 0x17, 0x4f, 0x40, 0xe9, 0x37, 0x4c, 0x47, 0x9c, 0xe1, 0x77,
	0x45, 0xe0, 0xf6, 0x2d, 0xb6, 0xa8, 0xe3, 0x19, 0xbf, 0x21, 0x3f, 0x80, 0x83, 0xec, 0x9c, 0x22,
	0x7f, 0x48, 0x3c, 0xf4, 0xe5, 0x82, 0x46, 0x64, 0x44, 0xc3, 0x48, 0xda, 0xe3, 0x47, 0x54, 0x4e,
	0x55, 0x87, 0x89, 0x1f, 0x72, 0x6d, 0x8d, 0x3a, 0x1b, 0xe1, 0xde, 0x90, 0x53, 0xe2, 0x1a, 0x55,
	0xcd, 0xb5, 0xec, 0x01, 0x20, 0xb6, 0x8b, 0x11, 0x8d, 0x2e, 0x11, 0x1e, 0x8f, 0x03, 0x7f, 0x8a,
	0x47, 0x28, 0xa2, 0x63, 0x69, 0x7f, 0xfe, 0x00, 0xb4, 0x4c, 0x57, 0x33, 0xd9, 0xa1, 0x63, 0xf1,
	0x3b, 0x50, 0x9e, 0x78, 0xf4, 0xf3, 0x84, 0x5c, 0xd1, 0x43, 0x72, 0x19, 0x4a, 0x65, 0x9e, 0x0b,
	0x62, 0xaa, 0xe5, 0x60, 0x93, 0x5c, 0x86, 0xdf, 0x26, 0x45, 0xb0, 0xbb, 0x92, 0x36, 0xec, 0xed,
	0xeb, 0x9a, 0x81, 0x9c, 0x8f, 0xa8, 0x01, 0xa1, 0x70, 0x47, 0xde, 0x8e, 0x13, 0xa5, 0xa8, 0x2f,
	0x24, 0x83, 0xae, 0x7e, 0xcc, 0xd5, 0x42, 0xa6, 0x2e, 0x24, 0x43, 0xd7, 0x74, 0x20, 0x7a, 0xdf,
	0x31, 0xed, 0x8e, 0x2e, 0xdc, 0x95, 0x77, 0xe2, 0x44, 0x01, 0xdd, 0xa5, 0x64, 0xa8, 0xc3, 0x86,
	0xda, 0x69, 0x39, 0xc8, 0xb2, 0x4d, 0xcb, 0x6c, 0xab, 0x2d, 0x04, 0x8d, 0x3a, 0x72, 0x34, 0x1d,
	0x0a, 0xf7, 0x64, 0x39, 0x4e, 0x94, 0x83, 0xfa, 0x8d, 0xc9, 0xa0, 0x6b, 0x86, 0xa6, 0x77, 0xf4,
	0x6b, 0xd0, 0x8d, 0x14, 0xd5, 0x6f, 0x4c, 0x86, 0x05, 0x44, 0xad, 0x39, 0x3a, 0x34, 0x9c, 0x14,
	0xbc, 0x2f, 0x3f, 0x8e, 0x13, 0xe5, 0x91, 0x75, 0x53, 0x32, 0xb0, 0xad, 0x58, 0x6a, 0x57, 0x6d,
	0x99, 0xb5, 0x66, 0x5b, 0x78, 0x20, 0x4b, 0x71, 0xa2, 0x94, 0xf5, 0x1b, 0x92, 0x61, 0x05, 0xcb,
	0x5b, 0x66, 0xf0, 0x43, 0xf9, 0x69, 0x9c, 0x28, 0xb2, 0x7e, 0x6b, 0x32, 0x40, 0x43, 0xad, 0xb6,
	0x20, 0x6a, 0x98, 0x36, 0xd4, 0x7e, 0x31, 0xd8, 0x76, 0x23, 0x4b, 0xfd, 0xc4, 0xca, 0xb4, 0x85,
	0xa2, 0xfc, 0x24, 0x4e, 0x14, 0x09, 0xde, 0x92, 0x0c, 0xba, 0xd6, 0xae, 0xbd, 0x55, 0x8d, 0x1a,
	0x44, 0xb6, 0x6a, 0x34, 0x51, 0x1d, 0xd6, 0x6c, 0xa8, 0xb6, 0x21, 0x52, 0x75, 0xb3, 0x63, 0x38,
	0xc2, 0xa6, 0x7c, 0x1c, 0x27, 0xca, 0x91, 0x7e, 0x7b, 0x32, 0xb0, 0xa3, 0x9e, 0x17, 0x12, 0x80,
	0x2c, 0xc4, 0x89, 0xb2, 0xad, 0xaf, 0x24, 0xc3, 0x55, 0xa7, 0x9a, 0x69, 0x34, 0xb4, 0x3a, 0x64,
	0xde, 0x2d, 0xf9, 0x30, 0x4e, 0x94, 0x7d, 0xfd, 0xfa, 0x64, 0xd0, 0xd8, 0x8e, 0x68, 0xdd, 0xd5,
	0xd9, 0x2c, 0x68, 0xd7, 0xa0, 0xe1, 0x08, 0xdb, 0xe9, 0x70, 0xda, 0x2d, 0xc9, 0xf0, 0x06, 0xc8,
	0x96, 0x69, 0xda, 0xc8, 0x80, 0xce, 0x07, 0xd3, 0x6e, 0x22, 0x36, 0x69, 0x95, 0x15, 0x6b, 0x43,
	0xa3, 0x2e, 0x94, 0xd2, 0xeb, 0x60, 0x5d, 0xff, 0xe4, 0x9f, 0x81, 0x1d, 0x76, 0x3e, 0x5d, 0xb5,
	0xa5, 0xd5, 0x55, 0xc7, 0xb4, 0xdb, 0xc2, 0x8e, 0xbc, 0x17, 0x27, 0x4a, 0x49, 0x5f, 0x4a, 0x95,
	0x13, 0x50, 0x7a, 0xa7, 0x6a, 0x2d, 0x5e, 0x9a, 0xdf, 0x95, 0x5d, 0x79, 0x37, 0x4e, 0x94, 0xad,
	0x77, 0xcb, 0xc9, 0x90, 0x9d, 0x93, 0x63, 0x36, 0xa1, 0x81, 0x3e, 0xbc, 0xd5, 0x1c, 0xd8, 0xd2,
	0xda, 0x8e, 0x20, 0xa4, 0x17, 0x04, 0xde, 0x90, 0x0c, 0x4b, 0x54, 0xb5, 0xa5, 0xd6, 0x9a, 0x9c,
	0xda, 0x5b, 0xa3, 0x96, 0x92, 0x81, 0x8d, 0xcd, 0x36, 0xd8, 0xd1, 0x9c, 0x4f, 0x48, 0xb5, 0x2c,
	0xdb, 0xec, 0xaa, 0x2d, 0xe4, 0x68, 0x96, 0x20, 0xce, 0x1f, 0xc0, 0x0d, 0xc9, 0xd0, 0x31, 0xb4,
	0xf7, 0x1d, 0x78, 0x45, 0x37, 0xe1, 0xa7, 0xb6, 0xb0, 0x2f, 0x1f, 0xc4, 0x89, 0x22, 0x76, 0xd6,
	0x92, 0x41, 0xde, 0xf8, 0xfd, 0xaf, 0xa7, 0x77, 0xaa, 0xaf, 0x7f, 0x7d, 0xb5, 0xf0, 0x3b, 0xa6,
	0x49, 0x03, 0x5c, 0xf3, 0x03, 0x52, 0x09, 0xc9, 0x10, 0xd3, 0x8a, 0x66, 0x38, 0xd0, 0xfe, 0x58,
	0xe1, 0xbf, 0xb6, 0x5e, 0x0c, 0x88, 0x57, 0xc9, 0x7f, 0xb3, 0x9d, 0x3d, 0xe0, 0x6b, 0xaf, 0xfe,
	0x0f, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xa5, 0xca, 0xa9, 0xc6, 0x09, 0x00, 0x00,
}
