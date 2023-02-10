// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kira/recovery/recovery.proto

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

type RecoveryRecord struct {
	Address   string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Challenge string `protobuf:"bytes,2,opt,name=challenge,proto3" json:"challenge,omitempty"`
	Nonce     string `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *RecoveryRecord) Reset()         { *m = RecoveryRecord{} }
func (m *RecoveryRecord) String() string { return proto.CompactTextString(m) }
func (*RecoveryRecord) ProtoMessage()    {}
func (*RecoveryRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2745d5958530d7f, []int{0}
}
func (m *RecoveryRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RecoveryRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RecoveryRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RecoveryRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecoveryRecord.Merge(m, src)
}
func (m *RecoveryRecord) XXX_Size() int {
	return m.Size()
}
func (m *RecoveryRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_RecoveryRecord.DiscardUnknown(m)
}

var xxx_messageInfo_RecoveryRecord proto.InternalMessageInfo

func (m *RecoveryRecord) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *RecoveryRecord) GetChallenge() string {
	if m != nil {
		return m.Challenge
	}
	return ""
}

func (m *RecoveryRecord) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

type RecoveryToken struct {
	Address          string                                    `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Token            string                                    `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	RrSupply         github_com_cosmos_cosmos_sdk_types.Int    `protobuf:"bytes,3,opt,name=rr_supply,json=rrSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"rr_supply"`
	UnderlyingTokens []github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,4,rep,name=underlying_tokens,json=underlyingTokens,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"underlying_tokens"`
}

func (m *RecoveryToken) Reset()         { *m = RecoveryToken{} }
func (m *RecoveryToken) String() string { return proto.CompactTextString(m) }
func (*RecoveryToken) ProtoMessage()    {}
func (*RecoveryToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2745d5958530d7f, []int{1}
}
func (m *RecoveryToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RecoveryToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RecoveryToken.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RecoveryToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecoveryToken.Merge(m, src)
}
func (m *RecoveryToken) XXX_Size() int {
	return m.Size()
}
func (m *RecoveryToken) XXX_DiscardUnknown() {
	xxx_messageInfo_RecoveryToken.DiscardUnknown(m)
}

var xxx_messageInfo_RecoveryToken proto.InternalMessageInfo

func (m *RecoveryToken) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *RecoveryToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type Rotation struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Rotated string `protobuf:"bytes,2,opt,name=rotated,proto3" json:"rotated,omitempty"`
}

func (m *Rotation) Reset()         { *m = Rotation{} }
func (m *Rotation) String() string { return proto.CompactTextString(m) }
func (*Rotation) ProtoMessage()    {}
func (*Rotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2745d5958530d7f, []int{2}
}
func (m *Rotation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Rotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Rotation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Rotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rotation.Merge(m, src)
}
func (m *Rotation) XXX_Size() int {
	return m.Size()
}
func (m *Rotation) XXX_DiscardUnknown() {
	xxx_messageInfo_Rotation.DiscardUnknown(m)
}

var xxx_messageInfo_Rotation proto.InternalMessageInfo

func (m *Rotation) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Rotation) GetRotated() string {
	if m != nil {
		return m.Rotated
	}
	return ""
}

func init() {
	proto.RegisterType((*RecoveryRecord)(nil), "kira.recovery.RecoveryRecord")
	proto.RegisterType((*RecoveryToken)(nil), "kira.recovery.RecoveryToken")
	proto.RegisterType((*Rotation)(nil), "kira.recovery.Rotation")
}

func init() { proto.RegisterFile("kira/recovery/recovery.proto", fileDescriptor_c2745d5958530d7f) }

var fileDescriptor_c2745d5958530d7f = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xb1, 0x6e, 0xa3, 0x40,
	0x10, 0x86, 0xe1, 0x7c, 0x3e, 0x9b, 0x95, 0x7c, 0xba, 0x43, 0x2e, 0x90, 0x65, 0x61, 0xcb, 0xc5,
	0x9d, 0x15, 0x29, 0x6c, 0x91, 0x3e, 0x85, 0xdd, 0x24, 0x72, 0x47, 0x52, 0x45, 0x51, 0x2c, 0x0c,
	0x1b, 0xbc, 0x02, 0x76, 0xd0, 0xee, 0x12, 0x85, 0xb7, 0xc8, 0x63, 0xe4, 0x51, 0x5c, 0xba, 0x8c,
	0x52, 0x58, 0x11, 0xee, 0xf2, 0x14, 0x11, 0x0b, 0xc4, 0xa9, 0xac, 0x54, 0xbb, 0xff, 0xfc, 0x3f,
	0xdf, 0xcc, 0x88, 0x45, 0xc3, 0x88, 0x72, 0x0f, 0x73, 0xe2, 0xc3, 0x03, 0xe1, 0xf9, 0xe7, 0xc5,
	0x49, 0x39, 0x48, 0x30, 0x7b, 0xa5, 0xeb, 0x34, 0xc5, 0x41, 0x3f, 0x84, 0x10, 0x94, 0x83, 0xcb,
	0x5b, 0x15, 0x1a, 0x8c, 0x42, 0x80, 0x30, 0x26, 0x58, 0xa9, 0x55, 0x76, 0x8f, 0x25, 0x4d, 0x88,
	0x90, 0x5e, 0x92, 0x56, 0x81, 0xc9, 0x1d, 0xfa, 0xed, 0xd6, 0x88, 0xf2, 0xe4, 0x81, 0x69, 0xa1,
	0x8e, 0x17, 0x04, 0x9c, 0x08, 0x61, 0xe9, 0x63, 0x7d, 0x6a, 0xb8, 0x8d, 0x34, 0x87, 0xc8, 0xf0,
	0xd7, 0x5e, 0x1c, 0x13, 0x16, 0x12, 0xeb, 0x87, 0xf2, 0x0e, 0x05, 0xb3, 0x8f, 0xda, 0x0c, 0x98,
	0x4f, 0xac, 0x96, 0x72, 0x2a, 0x31, 0x79, 0xd7, 0x51, 0xaf, 0x69, 0x70, 0x0d, 0x11, 0x61, 0x47,
	0xf8, 0x7d, 0xd4, 0x96, 0x65, 0xa4, 0x66, 0x57, 0xc2, 0x5c, 0x20, 0x83, 0xf3, 0xa5, 0xc8, 0xd2,
	0x34, 0xce, 0x2b, 0xf6, 0xcc, 0xd9, 0xec, 0x46, 0xda, 0xeb, 0x6e, 0xf4, 0x2f, 0xa4, 0x72, 0x9d,
	0xad, 0x1c, 0x1f, 0x12, 0xec, 0x83, 0x48, 0x40, 0xd4, 0xc7, 0xa9, 0x08, 0x22, 0x2c, 0xf3, 0x94,
	0x08, 0xe7, 0x92, 0x49, 0xb7, 0xcb, 0xf9, 0x95, 0xfa, 0xde, 0xbc, 0x45, 0x7f, 0x33, 0x16, 0x10,
	0x1e, 0xe7, 0x94, 0x85, 0x4b, 0xd5, 0x40, 0x58, 0x3f, 0xc7, 0xad, 0xa9, 0x31, 0xc3, 0x35, 0xf4,
	0xff, 0x37, 0xa0, 0x73, 0xa0, 0xcc, 0xfd, 0x73, 0x20, 0xa9, 0xcd, 0xc4, 0xe4, 0x1c, 0x75, 0x5d,
	0x90, 0x9e, 0xa4, 0x70, 0x6c, 0x4d, 0x0b, 0x75, 0x78, 0x99, 0x22, 0x41, 0xbd, 0x68, 0x23, 0x67,
	0x17, 0xcf, 0x85, 0xad, 0x6f, 0x0a, 0x5b, 0xdf, 0x16, 0xb6, 0xfe, 0x56, 0xd8, 0xfa, 0xd3, 0xde,
	0xd6, 0xb6, 0x7b, 0x5b, 0x7b, 0xd9, 0xdb, 0xda, 0xcd, 0xc9, 0x97, 0xc1, 0x16, 0x94, 0x7b, 0x73,
	0xe0, 0x04, 0x0b, 0x12, 0x79, 0x14, 0x3f, 0x1e, 0x5e, 0x89, 0x1a, 0x70, 0xf5, 0x4b, 0xfd, 0xdd,
	0xb3, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe8, 0xc1, 0xd7, 0x56, 0x43, 0x02, 0x00, 0x00,
}

func (this *RecoveryRecord) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RecoveryRecord)
	if !ok {
		that2, ok := that.(RecoveryRecord)
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
	if this.Address != that1.Address {
		return false
	}
	if this.Challenge != that1.Challenge {
		return false
	}
	if this.Nonce != that1.Nonce {
		return false
	}
	return true
}
func (this *RecoveryToken) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RecoveryToken)
	if !ok {
		that2, ok := that.(RecoveryToken)
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
	if this.Address != that1.Address {
		return false
	}
	if this.Token != that1.Token {
		return false
	}
	if !this.RrSupply.Equal(that1.RrSupply) {
		return false
	}
	if len(this.UnderlyingTokens) != len(that1.UnderlyingTokens) {
		return false
	}
	for i := range this.UnderlyingTokens {
		if !this.UnderlyingTokens[i].Equal(that1.UnderlyingTokens[i]) {
			return false
		}
	}
	return true
}
func (this *Rotation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Rotation)
	if !ok {
		that2, ok := that.(Rotation)
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
	if this.Address != that1.Address {
		return false
	}
	if this.Rotated != that1.Rotated {
		return false
	}
	return true
}
func (m *RecoveryRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RecoveryRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RecoveryRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Nonce) > 0 {
		i -= len(m.Nonce)
		copy(dAtA[i:], m.Nonce)
		i = encodeVarintRecovery(dAtA, i, uint64(len(m.Nonce)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Challenge) > 0 {
		i -= len(m.Challenge)
		copy(dAtA[i:], m.Challenge)
		i = encodeVarintRecovery(dAtA, i, uint64(len(m.Challenge)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintRecovery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RecoveryToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RecoveryToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RecoveryToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.UnderlyingTokens) > 0 {
		for iNdEx := len(m.UnderlyingTokens) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.UnderlyingTokens[iNdEx].Size()
				i -= size
				if _, err := m.UnderlyingTokens[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintRecovery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	{
		size := m.RrSupply.Size()
		i -= size
		if _, err := m.RrSupply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintRecovery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintRecovery(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintRecovery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Rotation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Rotation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Rotation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Rotated) > 0 {
		i -= len(m.Rotated)
		copy(dAtA[i:], m.Rotated)
		i = encodeVarintRecovery(dAtA, i, uint64(len(m.Rotated)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintRecovery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRecovery(dAtA []byte, offset int, v uint64) int {
	offset -= sovRecovery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RecoveryRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovRecovery(uint64(l))
	}
	l = len(m.Challenge)
	if l > 0 {
		n += 1 + l + sovRecovery(uint64(l))
	}
	l = len(m.Nonce)
	if l > 0 {
		n += 1 + l + sovRecovery(uint64(l))
	}
	return n
}

func (m *RecoveryToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovRecovery(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovRecovery(uint64(l))
	}
	l = m.RrSupply.Size()
	n += 1 + l + sovRecovery(uint64(l))
	if len(m.UnderlyingTokens) > 0 {
		for _, e := range m.UnderlyingTokens {
			l = e.Size()
			n += 1 + l + sovRecovery(uint64(l))
		}
	}
	return n
}

func (m *Rotation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovRecovery(uint64(l))
	}
	l = len(m.Rotated)
	if l > 0 {
		n += 1 + l + sovRecovery(uint64(l))
	}
	return n
}

func sovRecovery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRecovery(x uint64) (n int) {
	return sovRecovery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RecoveryRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecovery
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
			return fmt.Errorf("proto: RecoveryRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RecoveryRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecovery
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
				return ErrInvalidLengthRecovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRecovery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Challenge", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecovery
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
				return ErrInvalidLengthRecovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRecovery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Challenge = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecovery
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
				return ErrInvalidLengthRecovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRecovery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nonce = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRecovery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRecovery
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
func (m *RecoveryToken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecovery
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
			return fmt.Errorf("proto: RecoveryToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RecoveryToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecovery
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
				return ErrInvalidLengthRecovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRecovery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecovery
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
				return ErrInvalidLengthRecovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRecovery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RrSupply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecovery
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
				return ErrInvalidLengthRecovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRecovery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RrSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnderlyingTokens", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecovery
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
				return ErrInvalidLengthRecovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRecovery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Coin
			m.UnderlyingTokens = append(m.UnderlyingTokens, v)
			if err := m.UnderlyingTokens[len(m.UnderlyingTokens)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRecovery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRecovery
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
func (m *Rotation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecovery
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
			return fmt.Errorf("proto: Rotation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Rotation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecovery
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
				return ErrInvalidLengthRecovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRecovery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rotated", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecovery
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
				return ErrInvalidLengthRecovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRecovery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rotated = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRecovery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRecovery
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
func skipRecovery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRecovery
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
					return 0, ErrIntOverflowRecovery
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
					return 0, ErrIntOverflowRecovery
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
				return 0, ErrInvalidLengthRecovery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRecovery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRecovery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRecovery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRecovery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRecovery = fmt.Errorf("proto: unexpected end of group")
)
