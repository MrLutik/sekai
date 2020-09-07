// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type PermissionsByAddressRequest struct {
	ValAddr github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=val_addr,json=valAddr,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"val_addr,omitempty" yaml:"val_addr"`
}

func (m *PermissionsByAddressRequest) Reset()         { *m = PermissionsByAddressRequest{} }
func (m *PermissionsByAddressRequest) String() string { return proto.CompactTextString(m) }
func (*PermissionsByAddressRequest) ProtoMessage()    {}
func (*PermissionsByAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{0}
}
func (m *PermissionsByAddressRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PermissionsByAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PermissionsByAddressRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PermissionsByAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PermissionsByAddressRequest.Merge(m, src)
}
func (m *PermissionsByAddressRequest) XXX_Size() int {
	return m.Size()
}
func (m *PermissionsByAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PermissionsByAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PermissionsByAddressRequest proto.InternalMessageInfo

func (m *PermissionsByAddressRequest) GetValAddr() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.ValAddr
	}
	return nil
}

type PermissionsResponse struct {
	Permissions *Permissions `protobuf:"bytes,1,opt,name=permissions,proto3" json:"permissions,omitempty"`
}

func (m *PermissionsResponse) Reset()         { *m = PermissionsResponse{} }
func (m *PermissionsResponse) String() string { return proto.CompactTextString(m) }
func (*PermissionsResponse) ProtoMessage()    {}
func (*PermissionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{1}
}
func (m *PermissionsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PermissionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PermissionsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PermissionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PermissionsResponse.Merge(m, src)
}
func (m *PermissionsResponse) XXX_Size() int {
	return m.Size()
}
func (m *PermissionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PermissionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PermissionsResponse proto.InternalMessageInfo

func (m *PermissionsResponse) GetPermissions() *Permissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{2}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return m.Size()
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type RolePermissionsRequest struct {
	Role uint64 `protobuf:"varint,1,opt,name=role,proto3" json:"role,omitempty"`
}

func (m *RolePermissionsRequest) Reset()         { *m = RolePermissionsRequest{} }
func (m *RolePermissionsRequest) String() string { return proto.CompactTextString(m) }
func (*RolePermissionsRequest) ProtoMessage()    {}
func (*RolePermissionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{3}
}
func (m *RolePermissionsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RolePermissionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RolePermissionsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RolePermissionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RolePermissionsRequest.Merge(m, src)
}
func (m *RolePermissionsRequest) XXX_Size() int {
	return m.Size()
}
func (m *RolePermissionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RolePermissionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RolePermissionsRequest proto.InternalMessageInfo

func (m *RolePermissionsRequest) GetRole() uint64 {
	if m != nil {
		return m.Role
	}
	return 0
}

type RolePermissionsResponse struct {
	Permissions *Permissions `protobuf:"bytes,1,opt,name=permissions,proto3" json:"permissions,omitempty"`
}

func (m *RolePermissionsResponse) Reset()         { *m = RolePermissionsResponse{} }
func (m *RolePermissionsResponse) String() string { return proto.CompactTextString(m) }
func (*RolePermissionsResponse) ProtoMessage()    {}
func (*RolePermissionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{4}
}
func (m *RolePermissionsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RolePermissionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RolePermissionsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RolePermissionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RolePermissionsResponse.Merge(m, src)
}
func (m *RolePermissionsResponse) XXX_Size() int {
	return m.Size()
}
func (m *RolePermissionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RolePermissionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RolePermissionsResponse proto.InternalMessageInfo

func (m *RolePermissionsResponse) GetPermissions() *Permissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func init() {
	proto.RegisterType((*PermissionsByAddressRequest)(nil), "kira.gov.PermissionsByAddressRequest")
	proto.RegisterType((*PermissionsResponse)(nil), "kira.gov.PermissionsResponse")
	proto.RegisterType((*Empty)(nil), "kira.gov.Empty")
	proto.RegisterType((*RolePermissionsRequest)(nil), "kira.gov.RolePermissionsRequest")
	proto.RegisterType((*RolePermissionsResponse)(nil), "kira.gov.RolePermissionsResponse")
}

func init() { proto.RegisterFile("query.proto", fileDescriptor_5c6ac9b241082464) }

var fileDescriptor_5c6ac9b241082464 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x3f, 0xef, 0xd2, 0x50,
	0x14, 0x6d, 0xcd, 0x0f, 0xc1, 0x87, 0x09, 0xc9, 0x13, 0xff, 0xa4, 0xc4, 0x82, 0x2f, 0x21, 0x71,
	0x90, 0x36, 0xc1, 0xc1, 0xc4, 0xc5, 0x80, 0xff, 0x06, 0x13, 0x82, 0x9d, 0x8c, 0x83, 0xa4, 0xb4,
	0x37, 0xb5, 0x69, 0xcb, 0x2d, 0xef, 0x3d, 0xaa, 0x5d, 0xdc, 0xdd, 0xfc, 0x58, 0x8e, 0x8c, 0x4e,
	0xc4, 0xc0, 0x37, 0x70, 0x74, 0x32, 0xb4, 0x34, 0xad, 0xd0, 0xb8, 0x38, 0xf5, 0xe4, 0xf4, 0xdc,
	0x73, 0xef, 0x3b, 0xf7, 0x92, 0xf6, 0x7a, 0x03, 0x3c, 0x35, 0x62, 0x8e, 0x12, 0x69, 0x2b, 0xf0,
	0xb9, 0x6d, 0x78, 0x98, 0x68, 0x37, 0x3c, 0x4c, 0x72, 0x52, 0xeb, 0x7a, 0xe8, 0x61, 0x06, 0xcd,
	0x23, 0xca, 0x59, 0xf6, 0x85, 0xf4, 0xe6, 0xc0, 0x23, 0x5f, 0x08, 0x1f, 0x57, 0x62, 0x9a, 0x4e,
	0x5c, 0x97, 0x83, 0x10, 0x16, 0xac, 0x37, 0x20, 0x24, 0x5d, 0x90, 0x56, 0x62, 0x87, 0x0b, 0xdb,
	0x75, 0xf9, 0x3d, 0x75, 0xa0, 0x3e, 0xbc, 0x39, 0x7d, 0xf1, 0x6b, 0xd7, 0xef, 0xa4, 0x76, 0x14,
	0x3e, 0x65, 0xc5, 0x1f, 0xf6, 0x7b, 0xd7, 0x1f, 0x79, 0xbe, 0xfc, 0xb8, 0x59, 0x1a, 0x0e, 0x46,
	0xa6, 0x83, 0x22, 0x42, 0x71, 0xfa, 0x8c, 0x84, 0x1b, 0x98, 0x32, 0x8d, 0x41, 0x18, 0x13, 0xc7,
	0x29, 0xec, 0x9b, 0x89, 0x1d, 0x1e, 0x31, 0x9b, 0x91, 0x5b, 0x95, 0xfe, 0x16, 0x88, 0x18, 0x57,
	0x02, 0xe8, 0x13, 0xd2, 0x8e, 0x4b, 0x3a, 0x6b, 0xdd, 0x1e, 0xdf, 0x36, 0x8a, 0x77, 0x19, 0xd5,
	0x9a, 0xaa, 0x92, 0x35, 0x49, 0xe3, 0x65, 0x14, 0xcb, 0x94, 0x3d, 0x22, 0x77, 0x2c, 0x0c, 0xe1,
	0x2f, 0xf3, 0xfc, 0x4d, 0x94, 0x5c, 0x71, 0x0c, 0x21, 0x33, 0xbd, 0xb2, 0x32, 0xcc, 0x2c, 0x72,
	0xf7, 0x42, 0xfd, 0x9f, 0xa3, 0x8c, 0xbf, 0x5e, 0x23, 0x8d, 0xb7, 0xc7, 0xad, 0xd0, 0x0f, 0xa4,
	0x5b, 0x17, 0x32, 0x1d, 0xd6, 0xba, 0x9c, 0x2f, 0x41, 0xbb, 0x5f, 0xdf, 0xec, 0x34, 0x20, 0x53,
	0xe8, 0x2b, 0xd2, 0x7d, 0x0d, 0x72, 0x06, 0xf2, 0x13, 0xf2, 0x60, 0xce, 0x31, 0x06, 0x2e, 0x7d,
	0x10, 0xb4, 0x53, 0x16, 0x66, 0xa1, 0x68, 0xbd, 0x92, 0xb8, 0x50, 0x33, 0x85, 0xbe, 0x23, 0x9d,
	0xb3, 0x14, 0xe8, 0xa0, 0xac, 0xa8, 0x8f, 0x53, 0x7b, 0xf0, 0x0f, 0x45, 0x31, 0xe1, 0xf4, 0xd9,
	0xf7, 0xbd, 0xae, 0x6e, 0xf7, 0xba, 0xfa, 0x73, 0xaf, 0xab, 0xdf, 0x0e, 0xba, 0xb2, 0x3d, 0xe8,
	0xca, 0x8f, 0x83, 0xae, 0xbc, 0x1f, 0x56, 0x0e, 0xe7, 0x8d, 0xcf, 0xed, 0xe7, 0xc8, 0xc1, 0x14,
	0x10, 0xd8, 0xbe, 0xf9, 0xd9, 0xf4, 0x30, 0xc9, 0x6f, 0x67, 0x79, 0x3d, 0x3b, 0xd7, 0xc7, 0x7f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x71, 0x74, 0x9a, 0xde, 0xe8, 0x02, 0x00, 0x00,
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
	// Validators queries a validator by address.
	PermissionsByAddress(ctx context.Context, in *PermissionsByAddressRequest, opts ...grpc.CallOption) (*PermissionsResponse, error)
	GetNetworkProperties(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NetworkProperties, error)
	// RolePermissions returns the permissions of the roles available in the registry.
	RolePermissions(ctx context.Context, in *RolePermissionsRequest, opts ...grpc.CallOption) (*RolePermissionsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) PermissionsByAddress(ctx context.Context, in *PermissionsByAddressRequest, opts ...grpc.CallOption) (*PermissionsResponse, error) {
	out := new(PermissionsResponse)
	err := c.cc.Invoke(ctx, "/kira.gov.Query/PermissionsByAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetNetworkProperties(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NetworkProperties, error) {
	out := new(NetworkProperties)
	err := c.cc.Invoke(ctx, "/kira.gov.Query/GetNetworkProperties", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) RolePermissions(ctx context.Context, in *RolePermissionsRequest, opts ...grpc.CallOption) (*RolePermissionsResponse, error) {
	out := new(RolePermissionsResponse)
	err := c.cc.Invoke(ctx, "/kira.gov.Query/RolePermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Validators queries a validator by address.
	PermissionsByAddress(context.Context, *PermissionsByAddressRequest) (*PermissionsResponse, error)
	GetNetworkProperties(context.Context, *Empty) (*NetworkProperties, error)
	// RolePermissions returns the permissions of the roles available in the registry.
	RolePermissions(context.Context, *RolePermissionsRequest) (*RolePermissionsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) PermissionsByAddress(ctx context.Context, req *PermissionsByAddressRequest) (*PermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PermissionsByAddress not implemented")
}
func (*UnimplementedQueryServer) GetNetworkProperties(ctx context.Context, req *Empty) (*NetworkProperties, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNetworkProperties not implemented")
}
func (*UnimplementedQueryServer) RolePermissions(ctx context.Context, req *RolePermissionsRequest) (*RolePermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RolePermissions not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_PermissionsByAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionsByAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PermissionsByAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kira.gov.Query/PermissionsByAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PermissionsByAddress(ctx, req.(*PermissionsByAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetNetworkProperties_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetNetworkProperties(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kira.gov.Query/GetNetworkProperties",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetNetworkProperties(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_RolePermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RolePermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).RolePermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kira.gov.Query/RolePermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).RolePermissions(ctx, req.(*RolePermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kira.gov.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PermissionsByAddress",
			Handler:    _Query_PermissionsByAddress_Handler,
		},
		{
			MethodName: "GetNetworkProperties",
			Handler:    _Query_GetNetworkProperties_Handler,
		},
		{
			MethodName: "RolePermissions",
			Handler:    _Query_RolePermissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "query.proto",
}

func (m *PermissionsByAddressRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PermissionsByAddressRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PermissionsByAddressRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValAddr) > 0 {
		i -= len(m.ValAddr)
		copy(dAtA[i:], m.ValAddr)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ValAddr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PermissionsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PermissionsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PermissionsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Permissions != nil {
		{
			size, err := m.Permissions.MarshalToSizedBuffer(dAtA[:i])
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

func (m *Empty) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Empty) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Empty) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *RolePermissionsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RolePermissionsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RolePermissionsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Role != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Role))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RolePermissionsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RolePermissionsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RolePermissionsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Permissions != nil {
		{
			size, err := m.Permissions.MarshalToSizedBuffer(dAtA[:i])
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
func (m *PermissionsByAddressRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValAddr)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *PermissionsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Permissions != nil {
		l = m.Permissions.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *Empty) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *RolePermissionsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Role != 0 {
		n += 1 + sovQuery(uint64(m.Role))
	}
	return n
}

func (m *RolePermissionsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Permissions != nil {
		l = m.Permissions.Size()
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
func (m *PermissionsByAddressRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: PermissionsByAddressRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PermissionsByAddressRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValAddr", wireType)
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
			m.ValAddr = append(m.ValAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.ValAddr == nil {
				m.ValAddr = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
func (m *PermissionsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: PermissionsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PermissionsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
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
			if m.Permissions == nil {
				m.Permissions = &Permissions{}
			}
			if err := m.Permissions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
func (m *Empty) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Empty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Empty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
func (m *RolePermissionsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: RolePermissionsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RolePermissionsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			m.Role = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Role |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
func (m *RolePermissionsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: RolePermissionsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RolePermissionsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
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
			if m.Permissions == nil {
				m.Permissions = &Permissions{}
			}
			if err := m.Permissions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
