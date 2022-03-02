// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kira/spending/proposal.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// proposal-spending-pool-update - a function to create a proposal allowing
// modification of the existing spending pool, adding owners, beneficiaries,
// or otherwise editing any of the existing properties.
type UpdateSpendingPoolProposal struct {
	Name          string                                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ClaimStart    uint64                                 `protobuf:"varint,2,opt,name=claim_start,json=claimStart,proto3" json:"claim_start,omitempty"`
	ClaimEnd      uint64                                 `protobuf:"varint,3,opt,name=claim_end,json=claimEnd,proto3" json:"claim_end,omitempty"`
	Token         string                                 `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	Rate          github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=rate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"rate" yaml:"rate"`
	VoteQuorum    uint64                                 `protobuf:"varint,6,opt,name=vote_quorum,json=voteQuorum,proto3" json:"vote_quorum,omitempty"`
	VotePeriod    uint64                                 `protobuf:"varint,7,opt,name=vote_period,json=votePeriod,proto3" json:"vote_period,omitempty"`
	VoteEnactment uint64                                 `protobuf:"varint,8,opt,name=vote_enactment,json=voteEnactment,proto3" json:"vote_enactment,omitempty"`
	Owners        PermInfo                               `protobuf:"bytes,9,opt,name=owners,proto3" json:"owners"`
	Beneficiaries PermInfo                               `protobuf:"bytes,10,opt,name=beneficiaries,proto3" json:"beneficiaries"`
}

func (m *UpdateSpendingPoolProposal) Reset()         { *m = UpdateSpendingPoolProposal{} }
func (m *UpdateSpendingPoolProposal) String() string { return proto.CompactTextString(m) }
func (*UpdateSpendingPoolProposal) ProtoMessage()    {}
func (*UpdateSpendingPoolProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_e006ef21562b5bc9, []int{0}
}
func (m *UpdateSpendingPoolProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateSpendingPoolProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateSpendingPoolProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateSpendingPoolProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateSpendingPoolProposal.Merge(m, src)
}
func (m *UpdateSpendingPoolProposal) XXX_Size() int {
	return m.Size()
}
func (m *UpdateSpendingPoolProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateSpendingPoolProposal.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateSpendingPoolProposal proto.InternalMessageInfo

func (m *UpdateSpendingPoolProposal) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateSpendingPoolProposal) GetClaimStart() uint64 {
	if m != nil {
		return m.ClaimStart
	}
	return 0
}

func (m *UpdateSpendingPoolProposal) GetClaimEnd() uint64 {
	if m != nil {
		return m.ClaimEnd
	}
	return 0
}

func (m *UpdateSpendingPoolProposal) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *UpdateSpendingPoolProposal) GetVoteQuorum() uint64 {
	if m != nil {
		return m.VoteQuorum
	}
	return 0
}

func (m *UpdateSpendingPoolProposal) GetVotePeriod() uint64 {
	if m != nil {
		return m.VotePeriod
	}
	return 0
}

func (m *UpdateSpendingPoolProposal) GetVoteEnactment() uint64 {
	if m != nil {
		return m.VoteEnactment
	}
	return 0
}

func (m *UpdateSpendingPoolProposal) GetOwners() PermInfo {
	if m != nil {
		return m.Owners
	}
	return PermInfo{}
}

func (m *UpdateSpendingPoolProposal) GetBeneficiaries() PermInfo {
	if m != nil {
		return m.Beneficiaries
	}
	return PermInfo{}
}

// SpendingPoolDistributionProposal - force distribution of tokens to all
// beneficiaries registered in the claims array (this function should be
// automatically triggered before upgrades are executed)
type SpendingPoolDistributionProposal struct {
	PoolName string `protobuf:"bytes,1,opt,name=pool_name,json=poolName,proto3" json:"pool_name,omitempty"`
}

func (m *SpendingPoolDistributionProposal) Reset()         { *m = SpendingPoolDistributionProposal{} }
func (m *SpendingPoolDistributionProposal) String() string { return proto.CompactTextString(m) }
func (*SpendingPoolDistributionProposal) ProtoMessage()    {}
func (*SpendingPoolDistributionProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_e006ef21562b5bc9, []int{1}
}
func (m *SpendingPoolDistributionProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SpendingPoolDistributionProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SpendingPoolDistributionProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SpendingPoolDistributionProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpendingPoolDistributionProposal.Merge(m, src)
}
func (m *SpendingPoolDistributionProposal) XXX_Size() int {
	return m.Size()
}
func (m *SpendingPoolDistributionProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_SpendingPoolDistributionProposal.DiscardUnknown(m)
}

var xxx_messageInfo_SpendingPoolDistributionProposal proto.InternalMessageInfo

func (m *SpendingPoolDistributionProposal) GetPoolName() string {
	if m != nil {
		return m.PoolName
	}
	return ""
}

// SpendingPoolWithdrawProposal - proposal allowing withdrawal of funds
// from the pool to one or many specified accounts. Withdrawal should only
// be possible if the receiving account/s are on the list of registered
// beneficiaries.
type SpendingPoolWithdrawProposal struct {
	PoolName      string                                    `protobuf:"bytes,1,opt,name=pool_name,json=poolName,proto3" json:"pool_name,omitempty"`
	Beneficiaries []string                                  `protobuf:"bytes,2,rep,name=beneficiaries,proto3" json:"beneficiaries,omitempty"`
	Amounts       []github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,3,rep,name=amounts,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"amounts"`
}

func (m *SpendingPoolWithdrawProposal) Reset()         { *m = SpendingPoolWithdrawProposal{} }
func (m *SpendingPoolWithdrawProposal) String() string { return proto.CompactTextString(m) }
func (*SpendingPoolWithdrawProposal) ProtoMessage()    {}
func (*SpendingPoolWithdrawProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_e006ef21562b5bc9, []int{2}
}
func (m *SpendingPoolWithdrawProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SpendingPoolWithdrawProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SpendingPoolWithdrawProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SpendingPoolWithdrawProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpendingPoolWithdrawProposal.Merge(m, src)
}
func (m *SpendingPoolWithdrawProposal) XXX_Size() int {
	return m.Size()
}
func (m *SpendingPoolWithdrawProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_SpendingPoolWithdrawProposal.DiscardUnknown(m)
}

var xxx_messageInfo_SpendingPoolWithdrawProposal proto.InternalMessageInfo

func (m *SpendingPoolWithdrawProposal) GetPoolName() string {
	if m != nil {
		return m.PoolName
	}
	return ""
}

func (m *SpendingPoolWithdrawProposal) GetBeneficiaries() []string {
	if m != nil {
		return m.Beneficiaries
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdateSpendingPoolProposal)(nil), "kira.gov.UpdateSpendingPoolProposal")
	proto.RegisterType((*SpendingPoolDistributionProposal)(nil), "kira.gov.SpendingPoolDistributionProposal")
	proto.RegisterType((*SpendingPoolWithdrawProposal)(nil), "kira.gov.SpendingPoolWithdrawProposal")
}

func init() { proto.RegisterFile("kira/spending/proposal.proto", fileDescriptor_e006ef21562b5bc9) }

var fileDescriptor_e006ef21562b5bc9 = []byte{
	// 558 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6a, 0x1b, 0x3d,
	0x14, 0xf5, 0xc4, 0x4e, 0x62, 0x2b, 0xe4, 0xfb, 0x40, 0x04, 0xaa, 0x26, 0x61, 0xc6, 0x98, 0xfe,
	0x98, 0x42, 0x66, 0xa0, 0xa5, 0x9b, 0x40, 0x37, 0x76, 0xb2, 0x08, 0x85, 0xe2, 0x38, 0x94, 0x42,
	0x37, 0x46, 0x9e, 0x51, 0x26, 0xc2, 0x96, 0xee, 0x54, 0xd2, 0x24, 0xcd, 0x5b, 0xf4, 0x11, 0xfa,
	0x10, 0xdd, 0x76, 0x9f, 0x4d, 0x21, 0x74, 0x55, 0xba, 0x08, 0xc5, 0xde, 0x74, 0xdd, 0x27, 0x28,
	0xd2, 0x8c, 0x13, 0xd7, 0xd9, 0x78, 0x65, 0xe9, 0x9c, 0x73, 0x8f, 0xcf, 0x9d, 0x7b, 0x85, 0x76,
	0x47, 0x5c, 0xd1, 0x48, 0x67, 0x4c, 0x26, 0x5c, 0xa6, 0x51, 0xa6, 0x20, 0x03, 0x4d, 0xc7, 0x61,
	0xa6, 0xc0, 0x00, 0xae, 0x5b, 0x36, 0x4c, 0xe1, 0x7c, 0x7b, 0x2b, 0x85, 0x14, 0x1c, 0x18, 0xd9,
	0x53, 0xc1, 0x6f, 0x07, 0x29, 0x40, 0x3a, 0x66, 0x91, 0xbb, 0x0d, 0xf3, 0xd3, 0xc8, 0x70, 0xc1,
	0xb4, 0xa1, 0x22, 0x2b, 0x05, 0x0f, 0x17, 0x05, 0x54, 0x5e, 0xce, 0xa8, 0x18, 0xb4, 0x00, 0x3d,
	0x28, 0x4c, 0x8b, 0x4b, 0x49, 0x91, 0x85, 0x50, 0x00, 0x65, 0xa0, 0xd6, 0xb7, 0x2a, 0xda, 0x7e,
	0x9b, 0x25, 0xd4, 0xb0, 0x93, 0x92, 0xed, 0x01, 0x8c, 0x7b, 0x65, 0x6a, 0x8c, 0x51, 0x4d, 0x52,
	0xc1, 0x88, 0xd7, 0xf4, 0xda, 0x8d, 0xbe, 0x3b, 0xe3, 0x00, 0x6d, 0xc4, 0x63, 0xca, 0xc5, 0x40,
	0x1b, 0xaa, 0x0c, 0x59, 0x69, 0x7a, 0xed, 0x5a, 0x1f, 0x39, 0xe8, 0xc4, 0x22, 0x78, 0x07, 0x35,
	0x0a, 0x01, 0x93, 0x09, 0xa9, 0x3a, 0xba, 0xee, 0x80, 0x43, 0x99, 0xe0, 0x2d, 0xb4, 0x6a, 0x60,
	0xc4, 0x24, 0xa9, 0x39, 0xcb, 0xe2, 0x82, 0x8f, 0x51, 0x4d, 0x51, 0xc3, 0xc8, 0xaa, 0x05, 0x3b,
	0xaf, 0xae, 0x6e, 0x82, 0xca, 0xcf, 0x9b, 0xe0, 0x49, 0xca, 0xcd, 0x59, 0x3e, 0x0c, 0x63, 0x10,
	0x65, 0x3f, 0xe5, 0xcf, 0x9e, 0x4e, 0x46, 0x91, 0xb9, 0xcc, 0x98, 0x0e, 0x0f, 0x58, 0xfc, 0xe7,
	0x26, 0xd8, 0xb8, 0xa4, 0x62, 0xbc, 0xdf, 0xb2, 0x1e, 0xad, 0xbe, 0xb3, 0xb2, 0x31, 0xcf, 0xc1,
	0xb0, 0xc1, 0x87, 0x1c, 0x54, 0x2e, 0xc8, 0x5a, 0x11, 0xd3, 0x42, 0xc7, 0x0e, 0xb9, 0x15, 0x64,
	0x4c, 0x71, 0x48, 0xc8, 0xfa, 0x9d, 0xa0, 0xe7, 0x10, 0xfc, 0x18, 0xfd, 0xe7, 0x04, 0x4c, 0xd2,
	0xd8, 0x08, 0x26, 0x0d, 0xa9, 0x3b, 0xcd, 0xa6, 0x45, 0x0f, 0x67, 0x20, 0x7e, 0x89, 0xd6, 0xe0,
	0x42, 0x32, 0xa5, 0x49, 0xa3, 0xe9, 0xb5, 0x37, 0x9e, 0x3f, 0x08, 0xdd, 0x90, 0x67, 0x5f, 0x3b,
	0xec, 0x31, 0x25, 0x8e, 0xe4, 0x29, 0x74, 0x6a, 0xb6, 0xad, 0x7e, 0x29, 0xc6, 0x5d, 0xb4, 0x39,
	0x64, 0x92, 0x9d, 0xf2, 0x98, 0x53, 0xc5, 0x99, 0x26, 0x68, 0x99, 0xea, 0x7f, 0x6b, 0xf6, 0xff,
	0xff, 0xfd, 0x39, 0xf0, 0xbe, 0x7f, 0xd9, 0x5b, 0xef, 0x82, 0x34, 0x4c, 0x9a, 0x56, 0x0f, 0x35,
	0xe7, 0x07, 0x79, 0xc0, 0xb5, 0x51, 0x7c, 0x98, 0x1b, 0x0e, 0xf2, 0x76, 0xa8, 0x3b, 0xa8, 0x61,
	0x37, 0x60, 0x30, 0x37, 0xd9, 0xba, 0x05, 0xde, 0x50, 0xc1, 0xee, 0x3b, 0x7e, 0xf5, 0xd0, 0xee,
	0xbc, 0xe5, 0x3b, 0x6e, 0xce, 0x12, 0x45, 0x2f, 0x96, 0xb2, 0xc3, 0x8f, 0x16, 0xbb, 0x5c, 0x69,
	0x56, 0xdb, 0x8d, 0x85, 0x36, 0xf0, 0x11, 0x5a, 0xa7, 0x02, 0x72, 0x69, 0x34, 0xa9, 0x5a, 0xbe,
	0x13, 0x95, 0x1b, 0xf0, 0x74, 0x89, 0x0d, 0xe8, 0x02, 0x97, 0xfd, 0x59, 0xfd, 0xbd, 0xfc, 0x9d,
	0x83, 0xab, 0x89, 0xef, 0x5d, 0x4f, 0x7c, 0xef, 0xd7, 0xc4, 0xf7, 0x3e, 0x4d, 0xfd, 0xca, 0xf5,
	0xd4, 0xaf, 0xfc, 0x98, 0xfa, 0x95, 0xf7, 0xcf, 0xe6, 0xcc, 0x5f, 0x73, 0x45, 0xbb, 0xa0, 0x58,
	0xa4, 0xd9, 0x88, 0xf2, 0xe8, 0xe3, 0xdd, 0x63, 0x71, 0x7f, 0x32, 0x5c, 0x73, 0xcf, 0xe5, 0xc5,
	0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x72, 0x28, 0xa5, 0x9a, 0xdf, 0x03, 0x00, 0x00,
}

func (this *UpdateSpendingPoolProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpdateSpendingPoolProposal)
	if !ok {
		that2, ok := that.(UpdateSpendingPoolProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.ClaimStart != that1.ClaimStart {
		return false
	}
	if this.ClaimEnd != that1.ClaimEnd {
		return false
	}
	if this.Token != that1.Token {
		return false
	}
	if !this.Rate.Equal(that1.Rate) {
		return false
	}
	if this.VoteQuorum != that1.VoteQuorum {
		return false
	}
	if this.VotePeriod != that1.VotePeriod {
		return false
	}
	if this.VoteEnactment != that1.VoteEnactment {
		return false
	}
	if !this.Owners.Equal(&that1.Owners) {
		return false
	}
	if !this.Beneficiaries.Equal(&that1.Beneficiaries) {
		return false
	}
	return true
}
func (this *SpendingPoolDistributionProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SpendingPoolDistributionProposal)
	if !ok {
		that2, ok := that.(SpendingPoolDistributionProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.PoolName != that1.PoolName {
		return false
	}
	return true
}
func (this *SpendingPoolWithdrawProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SpendingPoolWithdrawProposal)
	if !ok {
		that2, ok := that.(SpendingPoolWithdrawProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.PoolName != that1.PoolName {
		return false
	}
	if len(this.Beneficiaries) != len(that1.Beneficiaries) {
		return false
	}
	for i := range this.Beneficiaries {
		if this.Beneficiaries[i] != that1.Beneficiaries[i] {
			return false
		}
	}
	if len(this.Amounts) != len(that1.Amounts) {
		return false
	}
	for i := range this.Amounts {
		if !this.Amounts[i].Equal(that1.Amounts[i]) {
			return false
		}
	}
	return true
}
func (m *UpdateSpendingPoolProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateSpendingPoolProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateSpendingPoolProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Beneficiaries.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	{
		size, err := m.Owners.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if m.VoteEnactment != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.VoteEnactment))
		i--
		dAtA[i] = 0x40
	}
	if m.VotePeriod != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.VotePeriod))
		i--
		dAtA[i] = 0x38
	}
	if m.VoteQuorum != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.VoteQuorum))
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.Rate.Size()
		i -= size
		if _, err := m.Rate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x22
	}
	if m.ClaimEnd != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.ClaimEnd))
		i--
		dAtA[i] = 0x18
	}
	if m.ClaimStart != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.ClaimStart))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SpendingPoolDistributionProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SpendingPoolDistributionProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SpendingPoolDistributionProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PoolName) > 0 {
		i -= len(m.PoolName)
		copy(dAtA[i:], m.PoolName)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.PoolName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SpendingPoolWithdrawProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SpendingPoolWithdrawProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SpendingPoolWithdrawProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amounts) > 0 {
		for iNdEx := len(m.Amounts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.Amounts[iNdEx].Size()
				i -= size
				if _, err := m.Amounts[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Beneficiaries) > 0 {
		for iNdEx := len(m.Beneficiaries) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Beneficiaries[iNdEx])
			copy(dAtA[i:], m.Beneficiaries[iNdEx])
			i = encodeVarintProposal(dAtA, i, uint64(len(m.Beneficiaries[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.PoolName) > 0 {
		i -= len(m.PoolName)
		copy(dAtA[i:], m.PoolName)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.PoolName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UpdateSpendingPoolProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.ClaimStart != 0 {
		n += 1 + sovProposal(uint64(m.ClaimStart))
	}
	if m.ClaimEnd != 0 {
		n += 1 + sovProposal(uint64(m.ClaimEnd))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = m.Rate.Size()
	n += 1 + l + sovProposal(uint64(l))
	if m.VoteQuorum != 0 {
		n += 1 + sovProposal(uint64(m.VoteQuorum))
	}
	if m.VotePeriod != 0 {
		n += 1 + sovProposal(uint64(m.VotePeriod))
	}
	if m.VoteEnactment != 0 {
		n += 1 + sovProposal(uint64(m.VoteEnactment))
	}
	l = m.Owners.Size()
	n += 1 + l + sovProposal(uint64(l))
	l = m.Beneficiaries.Size()
	n += 1 + l + sovProposal(uint64(l))
	return n
}

func (m *SpendingPoolDistributionProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PoolName)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	return n
}

func (m *SpendingPoolWithdrawProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PoolName)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if len(m.Beneficiaries) > 0 {
		for _, s := range m.Beneficiaries {
			l = len(s)
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	if len(m.Amounts) > 0 {
		for _, e := range m.Amounts {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UpdateSpendingPoolProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UpdateSpendingPoolProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateSpendingPoolProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimStart", wireType)
			}
			m.ClaimStart = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClaimStart |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimEnd", wireType)
			}
			m.ClaimEnd = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClaimEnd |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Rate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteQuorum", wireType)
			}
			m.VoteQuorum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VoteQuorum |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotePeriod", wireType)
			}
			m.VotePeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotePeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteEnactment", wireType)
			}
			m.VoteEnactment = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VoteEnactment |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owners", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Owners.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Beneficiaries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Beneficiaries.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SpendingPoolDistributionProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SpendingPoolDistributionProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SpendingPoolDistributionProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SpendingPoolWithdrawProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SpendingPoolWithdrawProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SpendingPoolWithdrawProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Beneficiaries", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Beneficiaries = append(m.Beneficiaries, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amounts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Coin
			m.Amounts = append(m.Amounts, v)
			if err := m.Amounts[len(m.Amounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)
