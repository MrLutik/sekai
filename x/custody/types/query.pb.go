// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kira/custody/query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CustodyByAddressRequest struct {
	Addr github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=addr,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"addr,omitempty" yaml:"addr"`
}

func (m *CustodyByAddressRequest) Reset()         { *m = CustodyByAddressRequest{} }
func (m *CustodyByAddressRequest) String() string { return proto.CompactTextString(m) }
func (*CustodyByAddressRequest) ProtoMessage()    {}
func (*CustodyByAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c8d8c1227063c95, []int{0}
}
func (m *CustodyByAddressRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CustodyByAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CustodyByAddressRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CustodyByAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustodyByAddressRequest.Merge(m, src)
}
func (m *CustodyByAddressRequest) XXX_Size() int {
	return m.Size()
}
func (m *CustodyByAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CustodyByAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CustodyByAddressRequest proto.InternalMessageInfo

func (m *CustodyByAddressRequest) GetAddr() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Addr
	}
	return nil
}

type CustodyByAddressResponse struct {
	CustodySettings *CustodySettings `protobuf:"bytes,1,opt,name=custody_settings,json=custodySettings,proto3" json:"custody_settings,omitempty"`
}

func (m *CustodyByAddressResponse) Reset()         { *m = CustodyByAddressResponse{} }
func (m *CustodyByAddressResponse) String() string { return proto.CompactTextString(m) }
func (*CustodyByAddressResponse) ProtoMessage()    {}
func (*CustodyByAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c8d8c1227063c95, []int{1}
}
func (m *CustodyByAddressResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CustodyByAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CustodyByAddressResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CustodyByAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustodyByAddressResponse.Merge(m, src)
}
func (m *CustodyByAddressResponse) XXX_Size() int {
	return m.Size()
}
func (m *CustodyByAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CustodyByAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CustodyByAddressResponse proto.InternalMessageInfo

func (m *CustodyByAddressResponse) GetCustodySettings() *CustodySettings {
	if m != nil {
		return m.CustodySettings
	}
	return nil
}

type CustodyWhiteListByAddressRequest struct {
	Addr github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=addr,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"addr,omitempty" yaml:"addr"`
}

func (m *CustodyWhiteListByAddressRequest) Reset()         { *m = CustodyWhiteListByAddressRequest{} }
func (m *CustodyWhiteListByAddressRequest) String() string { return proto.CompactTextString(m) }
func (*CustodyWhiteListByAddressRequest) ProtoMessage()    {}
func (*CustodyWhiteListByAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c8d8c1227063c95, []int{2}
}
func (m *CustodyWhiteListByAddressRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CustodyWhiteListByAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CustodyWhiteListByAddressRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CustodyWhiteListByAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustodyWhiteListByAddressRequest.Merge(m, src)
}
func (m *CustodyWhiteListByAddressRequest) XXX_Size() int {
	return m.Size()
}
func (m *CustodyWhiteListByAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CustodyWhiteListByAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CustodyWhiteListByAddressRequest proto.InternalMessageInfo

func (m *CustodyWhiteListByAddressRequest) GetAddr() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Addr
	}
	return nil
}

type CustodyWhiteListByAddressResponse struct {
	CustodyWhiteList *CustodyWhiteList `protobuf:"bytes,1,opt,name=custody_white_list,json=custodyWhiteList,proto3" json:"custody_white_list,omitempty"`
}

func (m *CustodyWhiteListByAddressResponse) Reset()         { *m = CustodyWhiteListByAddressResponse{} }
func (m *CustodyWhiteListByAddressResponse) String() string { return proto.CompactTextString(m) }
func (*CustodyWhiteListByAddressResponse) ProtoMessage()    {}
func (*CustodyWhiteListByAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c8d8c1227063c95, []int{3}
}
func (m *CustodyWhiteListByAddressResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CustodyWhiteListByAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CustodyWhiteListByAddressResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CustodyWhiteListByAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustodyWhiteListByAddressResponse.Merge(m, src)
}
func (m *CustodyWhiteListByAddressResponse) XXX_Size() int {
	return m.Size()
}
func (m *CustodyWhiteListByAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CustodyWhiteListByAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CustodyWhiteListByAddressResponse proto.InternalMessageInfo

func (m *CustodyWhiteListByAddressResponse) GetCustodyWhiteList() *CustodyWhiteList {
	if m != nil {
		return m.CustodyWhiteList
	}
	return nil
}

func init() {
	proto.RegisterType((*CustodyByAddressRequest)(nil), "kira.custody.CustodyByAddressRequest")
	proto.RegisterType((*CustodyByAddressResponse)(nil), "kira.custody.CustodyByAddressResponse")
	proto.RegisterType((*CustodyWhiteListByAddressRequest)(nil), "kira.custody.CustodyWhiteListByAddressRequest")
	proto.RegisterType((*CustodyWhiteListByAddressResponse)(nil), "kira.custody.CustodyWhiteListByAddressResponse")
}

func init() { proto.RegisterFile("kira/custody/query.proto", fileDescriptor_6c8d8c1227063c95) }

var fileDescriptor_6c8d8c1227063c95 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0xcd, 0xce, 0xd2, 0x40,
	0x14, 0xa5, 0x44, 0x5d, 0x0c, 0x24, 0x92, 0x89, 0x89, 0xd8, 0x68, 0xc1, 0x26, 0x88, 0x2c, 0xe8,
	0x24, 0xb8, 0x73, 0x25, 0xb0, 0x31, 0x91, 0x8d, 0x68, 0x62, 0xe2, 0x86, 0x0c, 0xed, 0xa4, 0x4c,
	0x80, 0x4e, 0xe9, 0x9d, 0x46, 0x1a, 0xe3, 0xc6, 0x27, 0x20, 0xf1, 0x39, 0x7c, 0x0f, 0xe3, 0x8a,
	0xc4, 0x8d, 0x2b, 0x62, 0xc0, 0x27, 0x70, 0xe9, 0xca, 0xb4, 0x1d, 0x2a, 0x3f, 0x15, 0xbf, 0xd5,
	0xb7, 0x1a, 0x72, 0xcf, 0xc9, 0xb9, 0xe7, 0x1c, 0x6e, 0x51, 0x75, 0xca, 0x03, 0x4a, 0xec, 0x10,
	0xa4, 0x70, 0x22, 0xb2, 0x08, 0x59, 0x10, 0x59, 0x7e, 0x20, 0xa4, 0xc0, 0xe5, 0x18, 0xb1, 0x14,
	0xa2, 0xdf, 0x71, 0x85, 0x2b, 0x12, 0x80, 0xc4, 0xbf, 0x52, 0x8e, 0x7e, 0xdf, 0x15, 0xc2, 0x9d,
	0x31, 0x42, 0x7d, 0x4e, 0xa8, 0xe7, 0x09, 0x49, 0x25, 0x17, 0x1e, 0x28, 0x54, 0x3f, 0xd2, 0x56,
	0x6f, 0x8a, 0x99, 0x02, 0xdd, 0xed, 0xa7, 0x83, 0x5e, 0xd4, 0x75, 0x9c, 0x80, 0x01, 0x0c, 0xd9,
	0x22, 0x64, 0x20, 0xf1, 0x6b, 0x74, 0x83, 0x3a, 0x4e, 0x50, 0xd5, 0xea, 0xda, 0xe3, 0x72, 0xef,
	0xd9, 0xaf, 0x4d, 0xad, 0x14, 0xd1, 0xf9, 0xec, 0xa9, 0x19, 0x4f, 0xcd, 0xdf, 0x9b, 0x5a, 0xdb,
	0xe5, 0x72, 0x12, 0x8e, 0x2d, 0x5b, 0xcc, 0x89, 0x2d, 0x60, 0x2e, 0x40, 0x3d, 0x6d, 0x70, 0xa6,
	0x44, 0x46, 0x3e, 0x03, 0xab, 0x6b, 0xdb, 0x7b, 0xd9, 0x44, 0xcd, 0x74, 0x50, 0xf5, 0x7c, 0x21,
	0xf8, 0xc2, 0x03, 0x86, 0x9f, 0xa3, 0x8a, 0x72, 0x37, 0x02, 0x26, 0x25, 0xf7, 0x5c, 0x48, 0xb6,
	0x97, 0x3a, 0x0f, 0xac, 0xc3, 0x16, 0x2c, 0xa5, 0xf0, 0x4a, 0x91, 0x86, 0xb7, 0xed, 0xe3, 0x81,
	0xb9, 0x44, 0x75, 0xc5, 0x79, 0x33, 0xe1, 0x92, 0x0d, 0x38, 0xc8, 0x6b, 0xca, 0xb7, 0x40, 0x0f,
	0x2f, 0x6c, 0x56, 0x41, 0x07, 0x08, 0xef, 0x83, 0xbe, 0x8b, 0x59, 0xa3, 0x19, 0x07, 0xa9, 0xa2,
	0x1a, 0xb9, 0x51, 0x33, 0xb1, 0xe1, 0xbe, 0xa2, 0x6c, 0xd2, 0xf9, 0x5a, 0x44, 0x37, 0x5f, 0xc6,
	0x17, 0x83, 0x57, 0x1a, 0xaa, 0x9c, 0xb6, 0x8b, 0x1b, 0xb9, 0x82, 0xa7, 0x75, 0xe8, 0x8f, 0xfe,
	0x47, 0x4b, 0xbd, 0x9b, 0xed, 0x8f, 0xdf, 0x7e, 0x7e, 0x2a, 0x36, 0x71, 0x83, 0xe4, 0x9d, 0x55,
	0xf6, 0xc7, 0x91, 0xf7, 0x71, 0x1d, 0x1f, 0xf0, 0x67, 0x0d, 0xdd, 0xfb, 0x67, 0x21, 0xd8, 0xba,
	0x1c, 0xf6, 0xcc, 0x24, 0xb9, 0x32, 0x5f, 0xb9, 0x25, 0x89, 0xdb, 0x16, 0x6e, 0xe6, 0xbb, 0xfd,
	0xdb, 0xbe, 0xf2, 0xdb, 0xeb, 0x7f, 0xd9, 0x1a, 0xda, 0x7a, 0x6b, 0x68, 0x3f, 0xb6, 0x86, 0xb6,
	0xda, 0x19, 0x85, 0xf5, 0xce, 0x28, 0x7c, 0xdf, 0x19, 0x85, 0xb7, 0xad, 0x83, 0x73, 0x78, 0xc1,
	0x03, 0xda, 0x17, 0x01, 0x23, 0xc0, 0xa6, 0x94, 0x93, 0x65, 0x26, 0x9c, 0x5c, 0xc5, 0xf8, 0x56,
	0xf2, 0x71, 0x3d, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0x58, 0x66, 0x35, 0x35, 0xd6, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	CustodyByAddress(ctx context.Context, in *CustodyByAddressRequest, opts ...grpc.CallOption) (*CustodyByAddressResponse, error)
	CustodyWhiteListByAddress(ctx context.Context, in *CustodyWhiteListByAddressRequest, opts ...grpc.CallOption) (*CustodyWhiteListByAddressResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) CustodyByAddress(ctx context.Context, in *CustodyByAddressRequest, opts ...grpc.CallOption) (*CustodyByAddressResponse, error) {
	out := new(CustodyByAddressResponse)
	err := c.cc.Invoke(ctx, "/kira.custody.Query/CustodyByAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CustodyWhiteListByAddress(ctx context.Context, in *CustodyWhiteListByAddressRequest, opts ...grpc.CallOption) (*CustodyWhiteListByAddressResponse, error) {
	out := new(CustodyWhiteListByAddressResponse)
	err := c.cc.Invoke(ctx, "/kira.custody.Query/CustodyWhiteListByAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	CustodyByAddress(context.Context, *CustodyByAddressRequest) (*CustodyByAddressResponse, error)
	CustodyWhiteListByAddress(context.Context, *CustodyWhiteListByAddressRequest) (*CustodyWhiteListByAddressResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) CustodyByAddress(ctx context.Context, req *CustodyByAddressRequest) (*CustodyByAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CustodyByAddress not implemented")
}
func (*UnimplementedQueryServer) CustodyWhiteListByAddress(ctx context.Context, req *CustodyWhiteListByAddressRequest) (*CustodyWhiteListByAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CustodyWhiteListByAddress not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_CustodyByAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustodyByAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CustodyByAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kira.custody.Query/CustodyByAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CustodyByAddress(ctx, req.(*CustodyByAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CustodyWhiteListByAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustodyWhiteListByAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CustodyWhiteListByAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kira.custody.Query/CustodyWhiteListByAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CustodyWhiteListByAddress(ctx, req.(*CustodyWhiteListByAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kira.custody.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CustodyByAddress",
			Handler:    _Query_CustodyByAddress_Handler,
		},
		{
			MethodName: "CustodyWhiteListByAddress",
			Handler:    _Query_CustodyWhiteListByAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kira/custody/query.proto",
}

func (m *CustodyByAddressRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CustodyByAddressRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CustodyByAddressRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Addr) > 0 {
		i -= len(m.Addr)
		copy(dAtA[i:], m.Addr)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Addr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CustodyByAddressResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CustodyByAddressResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CustodyByAddressResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CustodySettings != nil {
		{
			size, err := m.CustodySettings.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CustodyWhiteListByAddressRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CustodyWhiteListByAddressRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CustodyWhiteListByAddressRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Addr) > 0 {
		i -= len(m.Addr)
		copy(dAtA[i:], m.Addr)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Addr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CustodyWhiteListByAddressResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CustodyWhiteListByAddressResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CustodyWhiteListByAddressResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CustodyWhiteList != nil {
		{
			size, err := m.CustodyWhiteList.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CustodyByAddressRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Addr)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *CustodyByAddressResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CustodySettings != nil {
		l = m.CustodySettings.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *CustodyWhiteListByAddressRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Addr)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *CustodyWhiteListByAddressResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CustodyWhiteList != nil {
		l = m.CustodyWhiteList.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CustodyByAddressRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: CustodyByAddressRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CustodyByAddressRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addr = append(m.Addr[:0], dAtA[iNdEx:postIndex]...)
			if m.Addr == nil {
				m.Addr = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *CustodyByAddressResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: CustodyByAddressResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CustodyByAddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CustodySettings", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CustodySettings == nil {
				m.CustodySettings = &CustodySettings{}
			}
			if err := m.CustodySettings.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *CustodyWhiteListByAddressRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: CustodyWhiteListByAddressRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CustodyWhiteListByAddressRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addr = append(m.Addr[:0], dAtA[iNdEx:postIndex]...)
			if m.Addr == nil {
				m.Addr = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *CustodyWhiteListByAddressResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: CustodyWhiteListByAddressResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CustodyWhiteListByAddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CustodyWhiteList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CustodyWhiteList == nil {
				m.CustodyWhiteList = &CustodyWhiteList{}
			}
			if err := m.CustodyWhiteList.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
