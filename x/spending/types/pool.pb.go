// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kira/spending/pool.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type ClaimInfo struct {
	Account   string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	PoolName  string `protobuf:"bytes,2,opt,name=pool_name,json=poolName,proto3" json:"pool_name,omitempty"`
	LastClaim uint64 `protobuf:"varint,3,opt,name=last_claim,json=lastClaim,proto3" json:"last_claim,omitempty"`
}

func (m *ClaimInfo) Reset()         { *m = ClaimInfo{} }
func (m *ClaimInfo) String() string { return proto.CompactTextString(m) }
func (*ClaimInfo) ProtoMessage()    {}
func (*ClaimInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_6027931ab19c9a21, []int{0}
}
func (m *ClaimInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimInfo.Merge(m, src)
}
func (m *ClaimInfo) XXX_Size() int {
	return m.Size()
}
func (m *ClaimInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimInfo proto.InternalMessageInfo

func (m *ClaimInfo) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *ClaimInfo) GetPoolName() string {
	if m != nil {
		return m.PoolName
	}
	return ""
}

func (m *ClaimInfo) GetLastClaim() uint64 {
	if m != nil {
		return m.LastClaim
	}
	return 0
}

type PermInfo struct {
	OwnerRoles    []uint64 `protobuf:"varint,1,rep,packed,name=owner_roles,json=ownerRoles,proto3" json:"owner_roles,omitempty"`
	OwnerAccounts []string `protobuf:"bytes,2,rep,name=owner_accounts,json=ownerAccounts,proto3" json:"owner_accounts,omitempty"`
}

func (m *PermInfo) Reset()         { *m = PermInfo{} }
func (m *PermInfo) String() string { return proto.CompactTextString(m) }
func (*PermInfo) ProtoMessage()    {}
func (*PermInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_6027931ab19c9a21, []int{1}
}
func (m *PermInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PermInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PermInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PermInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PermInfo.Merge(m, src)
}
func (m *PermInfo) XXX_Size() int {
	return m.Size()
}
func (m *PermInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PermInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PermInfo proto.InternalMessageInfo

func (m *PermInfo) GetOwnerRoles() []uint64 {
	if m != nil {
		return m.OwnerRoles
	}
	return nil
}

func (m *PermInfo) GetOwnerAccounts() []string {
	if m != nil {
		return m.OwnerAccounts
	}
	return nil
}

type SpendingPool struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// claim-start & claim-end - defines the exact time period (Unix timestamps) between which tokens can be claimed from the pool, allowing for a precise funds distribution.
	ClaimStart uint64 `protobuf:"varint,2,opt,name=claim_start,json=claimStart,proto3" json:"claim_start,omitempty"`
	ClaimEnd   uint64 `protobuf:"varint,3,opt,name=claim_end,json=claimEnd,proto3" json:"claim_end,omitempty"`
	Token      string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	// rate of distribution in the smallest token denomination per 1 second (this value can be a float number, smaller than actual denomination)
	Rate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=rate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"rate" yaml:"rate"`
	// pool specific % of owner accounts that must vote YES or NO for any of the pool proposals to be valid.
	VoteQuorum uint64 `protobuf:"varint,6,opt,name=vote_quorum,json=voteQuorum,proto3" json:"vote_quorum,omitempty"`
	// period of time in seconds that any of the pool proposals must last before passing or being rejected
	VotePeriod uint64 `protobuf:"varint,7,opt,name=vote_period,json=votePeriod,proto3" json:"vote_period,omitempty"`
	// period of time that must pass before any of the pool proposal is enacted
	VoteEnactment uint64 `protobuf:"varint,8,opt,name=vote_enactment,json=voteEnactment,proto3" json:"vote_enactment,omitempty"`
	// defines a list of accounts/roles controlling the spending pool via “governance-like” proposals
	Owners *PermInfo `protobuf:"bytes,9,opt,name=owners,proto3" json:"owners,omitempty"`
	// defines set of accounts/roles to which funds can be distributed
	Beneficiaries *PermInfo                              `protobuf:"bytes,10,opt,name=beneficiaries,proto3" json:"beneficiaries,omitempty"`
	Balance       github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,11,opt,name=balance,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"balance"`
}

func (m *SpendingPool) Reset()         { *m = SpendingPool{} }
func (m *SpendingPool) String() string { return proto.CompactTextString(m) }
func (*SpendingPool) ProtoMessage()    {}
func (*SpendingPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_6027931ab19c9a21, []int{2}
}
func (m *SpendingPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SpendingPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SpendingPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SpendingPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpendingPool.Merge(m, src)
}
func (m *SpendingPool) XXX_Size() int {
	return m.Size()
}
func (m *SpendingPool) XXX_DiscardUnknown() {
	xxx_messageInfo_SpendingPool.DiscardUnknown(m)
}

var xxx_messageInfo_SpendingPool proto.InternalMessageInfo

func (m *SpendingPool) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SpendingPool) GetClaimStart() uint64 {
	if m != nil {
		return m.ClaimStart
	}
	return 0
}

func (m *SpendingPool) GetClaimEnd() uint64 {
	if m != nil {
		return m.ClaimEnd
	}
	return 0
}

func (m *SpendingPool) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *SpendingPool) GetVoteQuorum() uint64 {
	if m != nil {
		return m.VoteQuorum
	}
	return 0
}

func (m *SpendingPool) GetVotePeriod() uint64 {
	if m != nil {
		return m.VotePeriod
	}
	return 0
}

func (m *SpendingPool) GetVoteEnactment() uint64 {
	if m != nil {
		return m.VoteEnactment
	}
	return 0
}

func (m *SpendingPool) GetOwners() *PermInfo {
	if m != nil {
		return m.Owners
	}
	return nil
}

func (m *SpendingPool) GetBeneficiaries() *PermInfo {
	if m != nil {
		return m.Beneficiaries
	}
	return nil
}

func init() {
	proto.RegisterType((*ClaimInfo)(nil), "kira.spending.ClaimInfo")
	proto.RegisterType((*PermInfo)(nil), "kira.spending.PermInfo")
	proto.RegisterType((*SpendingPool)(nil), "kira.spending.SpendingPool")
}

func init() { proto.RegisterFile("kira/spending/pool.proto", fileDescriptor_6027931ab19c9a21) }

var fileDescriptor_6027931ab19c9a21 = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x6e, 0xd3, 0x4c,
	0x14, 0xc5, 0xe3, 0xc6, 0x6d, 0xe2, 0xc9, 0x97, 0x6f, 0x31, 0xaa, 0xc4, 0xa8, 0x08, 0x3b, 0x8a,
	0x04, 0x8a, 0x90, 0xb0, 0x25, 0xd8, 0x55, 0xea, 0x82, 0xfe, 0x91, 0xa8, 0x90, 0x50, 0xea, 0x6e,
	0x10, 0x9b, 0x68, 0xe2, 0xdc, 0x98, 0x51, 0xec, 0x19, 0x33, 0x33, 0x01, 0xfa, 0x16, 0x3c, 0x02,
	0x8f, 0xd3, 0x65, 0x97, 0x88, 0x45, 0x84, 0x92, 0x0d, 0x0b, 0x56, 0x3c, 0x01, 0x9a, 0x6b, 0x87,
	0xd2, 0x0d, 0x12, 0x2b, 0xcf, 0xfc, 0xee, 0x99, 0x73, 0x8f, 0xaf, 0xc7, 0x84, 0x2d, 0x84, 0xe6,
	0x89, 0xa9, 0x40, 0xce, 0x84, 0xcc, 0x93, 0x4a, 0xa9, 0x22, 0xae, 0xb4, 0xb2, 0x8a, 0xf6, 0x5d,
	0x25, 0xde, 0x56, 0x0e, 0xa2, 0x5c, 0xa9, 0xbc, 0x80, 0x04, 0x8b, 0xd3, 0xe5, 0x3c, 0xb1, 0xa2,
	0x04, 0x63, 0x79, 0x59, 0xd5, 0xfa, 0x83, 0xfd, 0x5c, 0xe5, 0x0a, 0x97, 0x89, 0x5b, 0xd5, 0x74,
	0xc8, 0x49, 0x70, 0x52, 0x70, 0x51, 0x9e, 0xcb, 0xb9, 0xa2, 0x8c, 0x74, 0x78, 0x96, 0xa9, 0xa5,
	0xb4, 0xcc, 0x1b, 0x78, 0xa3, 0x20, 0xdd, 0x6e, 0xe9, 0x7d, 0x12, 0xb8, 0xd6, 0x13, 0xc9, 0x4b,
	0x60, 0x3b, 0x58, 0xeb, 0x3a, 0xf0, 0x8a, 0x97, 0x40, 0x1f, 0x10, 0x52, 0x70, 0x63, 0x27, 0x99,
	0x33, 0x62, 0xed, 0x81, 0x37, 0xf2, 0xd3, 0xc0, 0x11, 0x74, 0x1e, 0xbe, 0x26, 0xdd, 0x31, 0xe8,
	0xba, 0x43, 0x44, 0x7a, 0xea, 0x83, 0x04, 0x3d, 0xd1, 0xaa, 0x00, 0xc3, 0xbc, 0x41, 0x7b, 0xe4,
	0xa7, 0x04, 0x51, 0xea, 0x08, 0x7d, 0x48, 0xfe, 0xaf, 0x05, 0x4d, 0x67, 0xc3, 0x76, 0x06, 0xed,
	0x51, 0x90, 0xf6, 0x91, 0x3e, 0x6f, 0xe0, 0xa1, 0xff, 0xfd, 0x73, 0xe4, 0x0d, 0x7f, 0xb4, 0xc9,
	0x7f, 0x97, 0xcd, 0x00, 0xc6, 0x4a, 0x15, 0x94, 0x12, 0x1f, 0x13, 0xd6, 0xe9, 0x71, 0xed, 0x5a,
	0x62, 0xb0, 0x89, 0xb1, 0x5c, 0x5b, 0x0c, 0xef, 0xa7, 0x04, 0xd1, 0xa5, 0x23, 0xee, 0xdd, 0x6a,
	0x01, 0xc8, 0x59, 0x93, 0xbe, 0x8b, 0xe0, 0x4c, 0xce, 0xe8, 0x3e, 0xd9, 0xb5, 0x6a, 0x01, 0x92,
	0xf9, 0x68, 0x59, 0x6f, 0xe8, 0x05, 0xf1, 0x35, 0xb7, 0xc0, 0x76, 0x1d, 0x3c, 0x3e, 0xba, 0x5e,
	0x45, 0xad, 0xaf, 0xab, 0xe8, 0x51, 0x2e, 0xec, 0xdb, 0xe5, 0x34, 0xce, 0x54, 0x99, 0x64, 0xca,
	0x94, 0xca, 0x34, 0x8f, 0x27, 0x66, 0xb6, 0x48, 0xec, 0x55, 0x05, 0x26, 0x3e, 0x85, 0xec, 0xe7,
	0x2a, 0xea, 0x5d, 0xf1, 0xb2, 0x38, 0x1c, 0x3a, 0x8f, 0x61, 0x8a, 0x56, 0x2e, 0xe6, 0x7b, 0x65,
	0x61, 0xf2, 0x6e, 0xa9, 0xf4, 0xb2, 0x64, 0x7b, 0x75, 0x4c, 0x87, 0x2e, 0x90, 0xfc, 0x16, 0x54,
	0xa0, 0x85, 0x9a, 0xb1, 0xce, 0xad, 0x60, 0x8c, 0xc4, 0x8d, 0x0e, 0x05, 0x20, 0x79, 0x66, 0x4b,
	0x90, 0x96, 0x75, 0x51, 0xd3, 0x77, 0xf4, 0x6c, 0x0b, 0x69, 0x42, 0xf6, 0x70, 0x96, 0x86, 0x05,
	0x03, 0x6f, 0xd4, 0x7b, 0x7a, 0x2f, 0xbe, 0x73, 0x91, 0xe2, 0xed, 0xb7, 0x4a, 0x1b, 0x19, 0x3d,
	0x22, 0xfd, 0x29, 0x48, 0x98, 0x8b, 0x4c, 0x70, 0x2d, 0xc0, 0x30, 0xf2, 0xf7, 0x73, 0x77, 0xd5,
	0xf4, 0x05, 0xe9, 0x4c, 0x79, 0xc1, 0x65, 0x06, 0xac, 0x87, 0xe3, 0x8a, 0xff, 0x61, 0x5c, 0xe7,
	0xd2, 0xa6, 0xdb, 0xe3, 0xc7, 0xa7, 0xd7, 0xeb, 0xd0, 0xbb, 0x59, 0x87, 0xde, 0xb7, 0x75, 0xe8,
	0x7d, 0xda, 0x84, 0xad, 0x9b, 0x4d, 0xd8, 0xfa, 0xb2, 0x09, 0x5b, 0x6f, 0x1e, 0xff, 0x61, 0xf5,
	0x52, 0x68, 0x7e, 0xa2, 0x34, 0x24, 0x06, 0x16, 0x5c, 0x24, 0x1f, 0x6f, 0x7f, 0x1e, 0xb4, 0x9c,
	0xee, 0xe1, 0xc5, 0x7f, 0xf6, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xad, 0x66, 0xa7, 0x3b, 0x5a, 0x03,
	0x00, 0x00,
}

func (this *PermInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PermInfo)
	if !ok {
		that2, ok := that.(PermInfo)
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
	if len(this.OwnerRoles) != len(that1.OwnerRoles) {
		return false
	}
	for i := range this.OwnerRoles {
		if this.OwnerRoles[i] != that1.OwnerRoles[i] {
			return false
		}
	}
	if len(this.OwnerAccounts) != len(that1.OwnerAccounts) {
		return false
	}
	for i := range this.OwnerAccounts {
		if this.OwnerAccounts[i] != that1.OwnerAccounts[i] {
			return false
		}
	}
	return true
}
func (m *ClaimInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastClaim != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.LastClaim))
		i--
		dAtA[i] = 0x18
	}
	if len(m.PoolName) > 0 {
		i -= len(m.PoolName)
		copy(dAtA[i:], m.PoolName)
		i = encodeVarintPool(dAtA, i, uint64(len(m.PoolName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintPool(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PermInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PermInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PermInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OwnerAccounts) > 0 {
		for iNdEx := len(m.OwnerAccounts) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.OwnerAccounts[iNdEx])
			copy(dAtA[i:], m.OwnerAccounts[iNdEx])
			i = encodeVarintPool(dAtA, i, uint64(len(m.OwnerAccounts[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.OwnerRoles) > 0 {
		dAtA2 := make([]byte, len(m.OwnerRoles)*10)
		var j1 int
		for _, num := range m.OwnerRoles {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintPool(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SpendingPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SpendingPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SpendingPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Balance.Size()
		i -= size
		if _, err := m.Balance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	if m.Beneficiaries != nil {
		{
			size, err := m.Beneficiaries.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPool(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	if m.Owners != nil {
		{
			size, err := m.Owners.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPool(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	if m.VoteEnactment != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.VoteEnactment))
		i--
		dAtA[i] = 0x40
	}
	if m.VotePeriod != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.VotePeriod))
		i--
		dAtA[i] = 0x38
	}
	if m.VoteQuorum != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.VoteQuorum))
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.Rate.Size()
		i -= size
		if _, err := m.Rate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintPool(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x22
	}
	if m.ClaimEnd != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.ClaimEnd))
		i--
		dAtA[i] = 0x18
	}
	if m.ClaimStart != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.ClaimStart))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPool(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClaimInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = len(m.PoolName)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	if m.LastClaim != 0 {
		n += 1 + sovPool(uint64(m.LastClaim))
	}
	return n
}

func (m *PermInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.OwnerRoles) > 0 {
		l = 0
		for _, e := range m.OwnerRoles {
			l += sovPool(uint64(e))
		}
		n += 1 + sovPool(uint64(l)) + l
	}
	if len(m.OwnerAccounts) > 0 {
		for _, s := range m.OwnerAccounts {
			l = len(s)
			n += 1 + l + sovPool(uint64(l))
		}
	}
	return n
}

func (m *SpendingPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	if m.ClaimStart != 0 {
		n += 1 + sovPool(uint64(m.ClaimStart))
	}
	if m.ClaimEnd != 0 {
		n += 1 + sovPool(uint64(m.ClaimEnd))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = m.Rate.Size()
	n += 1 + l + sovPool(uint64(l))
	if m.VoteQuorum != 0 {
		n += 1 + sovPool(uint64(m.VoteQuorum))
	}
	if m.VotePeriod != 0 {
		n += 1 + sovPool(uint64(m.VotePeriod))
	}
	if m.VoteEnactment != 0 {
		n += 1 + sovPool(uint64(m.VoteEnactment))
	}
	if m.Owners != nil {
		l = m.Owners.Size()
		n += 1 + l + sovPool(uint64(l))
	}
	if m.Beneficiaries != nil {
		l = m.Beneficiaries.Size()
		n += 1 + l + sovPool(uint64(l))
	}
	l = m.Balance.Size()
	n += 1 + l + sovPool(uint64(l))
	return n
}

func sovPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPool(x uint64) (n int) {
	return sovPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClaimInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			return fmt.Errorf("proto: ClaimInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastClaim", wireType)
			}
			m.LastClaim = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastClaim |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPool
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
func (m *PermInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			return fmt.Errorf("proto: PermInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PermInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPool
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.OwnerRoles = append(m.OwnerRoles, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPool
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthPool
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthPool
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.OwnerRoles) == 0 {
					m.OwnerRoles = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowPool
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.OwnerRoles = append(m.OwnerRoles, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerRoles", wireType)
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerAccounts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerAccounts = append(m.OwnerAccounts, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPool
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
func (m *SpendingPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			return fmt.Errorf("proto: SpendingPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SpendingPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
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
					return ErrIntOverflowPool
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
					return ErrIntOverflowPool
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
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
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
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
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
					return ErrIntOverflowPool
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
					return ErrIntOverflowPool
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
					return ErrIntOverflowPool
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
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Owners == nil {
				m.Owners = &PermInfo{}
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
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Beneficiaries == nil {
				m.Beneficiaries = &PermInfo{}
			}
			if err := m.Beneficiaries.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Balance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPool
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
func skipPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
				return 0, ErrInvalidLengthPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPool = fmt.Errorf("proto: unexpected end of group")
)
