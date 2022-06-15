// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kira/multistaking/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/query"
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

type QueryStakingPoolsRequest struct {
}

func (m *QueryStakingPoolsRequest) Reset()         { *m = QueryStakingPoolsRequest{} }
func (m *QueryStakingPoolsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryStakingPoolsRequest) ProtoMessage()    {}
func (*QueryStakingPoolsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c25b06fd9ed13dd9, []int{0}
}
func (m *QueryStakingPoolsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryStakingPoolsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryStakingPoolsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryStakingPoolsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStakingPoolsRequest.Merge(m, src)
}
func (m *QueryStakingPoolsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryStakingPoolsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStakingPoolsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStakingPoolsRequest proto.InternalMessageInfo

type QueryStakingPoolsResponse struct {
	Pools []StakingPool `protobuf:"bytes,1,rep,name=pools,proto3" json:"pools"`
}

func (m *QueryStakingPoolsResponse) Reset()         { *m = QueryStakingPoolsResponse{} }
func (m *QueryStakingPoolsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryStakingPoolsResponse) ProtoMessage()    {}
func (*QueryStakingPoolsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c25b06fd9ed13dd9, []int{1}
}
func (m *QueryStakingPoolsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryStakingPoolsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryStakingPoolsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryStakingPoolsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStakingPoolsResponse.Merge(m, src)
}
func (m *QueryStakingPoolsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryStakingPoolsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStakingPoolsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStakingPoolsResponse proto.InternalMessageInfo

func (m *QueryStakingPoolsResponse) GetPools() []StakingPool {
	if m != nil {
		return m.Pools
	}
	return nil
}

type QueryOutstandingRewardsRequest struct {
	Delegator string `protobuf:"bytes,1,opt,name=delegator,proto3" json:"delegator,omitempty"`
}

func (m *QueryOutstandingRewardsRequest) Reset()         { *m = QueryOutstandingRewardsRequest{} }
func (m *QueryOutstandingRewardsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryOutstandingRewardsRequest) ProtoMessage()    {}
func (*QueryOutstandingRewardsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c25b06fd9ed13dd9, []int{2}
}
func (m *QueryOutstandingRewardsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryOutstandingRewardsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryOutstandingRewardsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryOutstandingRewardsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryOutstandingRewardsRequest.Merge(m, src)
}
func (m *QueryOutstandingRewardsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryOutstandingRewardsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryOutstandingRewardsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryOutstandingRewardsRequest proto.InternalMessageInfo

func (m *QueryOutstandingRewardsRequest) GetDelegator() string {
	if m != nil {
		return m.Delegator
	}
	return ""
}

type QueryOutstandingRewardsResponse struct {
	Rewards []github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,1,rep,name=rewards,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"rewards"`
}

func (m *QueryOutstandingRewardsResponse) Reset()         { *m = QueryOutstandingRewardsResponse{} }
func (m *QueryOutstandingRewardsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryOutstandingRewardsResponse) ProtoMessage()    {}
func (*QueryOutstandingRewardsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c25b06fd9ed13dd9, []int{3}
}
func (m *QueryOutstandingRewardsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryOutstandingRewardsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryOutstandingRewardsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryOutstandingRewardsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryOutstandingRewardsResponse.Merge(m, src)
}
func (m *QueryOutstandingRewardsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryOutstandingRewardsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryOutstandingRewardsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryOutstandingRewardsResponse proto.InternalMessageInfo

type QueryUndelegationsRequest struct {
}

func (m *QueryUndelegationsRequest) Reset()         { *m = QueryUndelegationsRequest{} }
func (m *QueryUndelegationsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryUndelegationsRequest) ProtoMessage()    {}
func (*QueryUndelegationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c25b06fd9ed13dd9, []int{4}
}
func (m *QueryUndelegationsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryUndelegationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryUndelegationsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryUndelegationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUndelegationsRequest.Merge(m, src)
}
func (m *QueryUndelegationsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryUndelegationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUndelegationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUndelegationsRequest proto.InternalMessageInfo

type QueryUndelegationsResponse struct {
	Undelegations []Undelegation `protobuf:"bytes,1,rep,name=undelegations,proto3" json:"undelegations"`
}

func (m *QueryUndelegationsResponse) Reset()         { *m = QueryUndelegationsResponse{} }
func (m *QueryUndelegationsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryUndelegationsResponse) ProtoMessage()    {}
func (*QueryUndelegationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c25b06fd9ed13dd9, []int{5}
}
func (m *QueryUndelegationsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryUndelegationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryUndelegationsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryUndelegationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUndelegationsResponse.Merge(m, src)
}
func (m *QueryUndelegationsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryUndelegationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUndelegationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUndelegationsResponse proto.InternalMessageInfo

func (m *QueryUndelegationsResponse) GetUndelegations() []Undelegation {
	if m != nil {
		return m.Undelegations
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryStakingPoolsRequest)(nil), "kira.multistaking.QueryStakingPoolsRequest")
	proto.RegisterType((*QueryStakingPoolsResponse)(nil), "kira.multistaking.QueryStakingPoolsResponse")
	proto.RegisterType((*QueryOutstandingRewardsRequest)(nil), "kira.multistaking.QueryOutstandingRewardsRequest")
	proto.RegisterType((*QueryOutstandingRewardsResponse)(nil), "kira.multistaking.QueryOutstandingRewardsResponse")
	proto.RegisterType((*QueryUndelegationsRequest)(nil), "kira.multistaking.QueryUndelegationsRequest")
	proto.RegisterType((*QueryUndelegationsResponse)(nil), "kira.multistaking.QueryUndelegationsResponse")
}

func init() { proto.RegisterFile("kira/multistaking/query.proto", fileDescriptor_c25b06fd9ed13dd9) }

var fileDescriptor_c25b06fd9ed13dd9 = []byte{
	// 516 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xae, 0x19, 0x03, 0xd5, 0xb0, 0x03, 0x16, 0x87, 0x2e, 0x8c, 0x74, 0x8a, 0x90, 0xa8, 0x06,
	0x8b, 0xd7, 0x22, 0x71, 0xe0, 0xc0, 0xa1, 0xbb, 0x80, 0x76, 0x00, 0x82, 0x10, 0x12, 0x97, 0xc9,
	0x59, 0x8d, 0xb1, 0x9a, 0xfa, 0x65, 0xb1, 0x03, 0xec, 0xca, 0x2f, 0x40, 0xe2, 0x88, 0xc4, 0x9f,
	0xe0, 0x4f, 0xec, 0x38, 0x89, 0x0b, 0xe2, 0x30, 0xa1, 0x96, 0x9f, 0xc1, 0x01, 0xc5, 0x71, 0xb5,
	0x54, 0x4d, 0xd0, 0x76, 0x8a, 0xfd, 0xde, 0xfb, 0xde, 0xfb, 0xfc, 0x7d, 0x2f, 0xf8, 0xf6, 0x58,
	0x66, 0x8c, 0x4e, 0xf2, 0xc4, 0x48, 0x6d, 0xd8, 0x58, 0x2a, 0x41, 0x0f, 0x73, 0x9e, 0x1d, 0x85,
	0x69, 0x06, 0x06, 0xc8, 0x8d, 0x22, 0x1d, 0x56, 0xd3, 0xde, 0xd6, 0x01, 0xe8, 0x09, 0x68, 0x1a,
	0x33, 0xcd, 0xcb, 0x5a, 0xfa, 0xbe, 0x1f, 0x73, 0xc3, 0xfa, 0x34, 0x65, 0x42, 0x2a, 0x66, 0x24,
	0xa8, 0x12, 0xee, 0xdd, 0x14, 0x20, 0xc0, 0x1e, 0x69, 0x71, 0x72, 0xd1, 0x75, 0x01, 0x20, 0x12,
	0x4e, 0xed, 0x2d, 0xce, 0xdf, 0x52, 0xa6, 0xdc, 0x3c, 0x6f, 0xc3, 0xa5, 0x58, 0x2a, 0x29, 0x53,
	0x0a, 0x8c, 0xed, 0xa6, 0x5d, 0xf6, 0xce, 0x32, 0xd9, 0xea, 0xa5, 0xac, 0x0a, 0x3c, 0xdc, 0x79,
	0x51, 0xd0, 0x7a, 0x59, 0x46, 0x9f, 0x03, 0x24, 0x3a, 0xe2, 0x87, 0x39, 0xd7, 0x26, 0x78, 0x8d,
	0xd7, 0x6b, 0x72, 0x3a, 0x05, 0xa5, 0x39, 0x79, 0x84, 0x57, 0xd3, 0x22, 0xd0, 0x41, 0x9b, 0x2b,
	0xbd, 0x6b, 0x03, 0x3f, 0x5c, 0x7a, 0x7c, 0x58, 0xc1, 0x0d, 0x2f, 0x1f, 0x9f, 0x76, 0x5b, 0x51,
	0x09, 0x09, 0x1e, 0x63, 0xdf, 0x36, 0x7e, 0x96, 0x1b, 0x6d, 0x98, 0x1a, 0x49, 0x25, 0x22, 0xfe,
	0x81, 0x65, 0xa3, 0xf9, 0x68, 0xb2, 0x81, 0xdb, 0x23, 0x9e, 0x70, 0xc1, 0x0c, 0x64, 0x1d, 0xb4,
	0x89, 0x7a, 0xed, 0xe8, 0x2c, 0x10, 0x24, 0xb8, 0xdb, 0x88, 0x77, 0xf4, 0x9e, 0xe2, 0xab, 0x59,
	0x19, 0xb2, 0x04, 0xdb, 0x43, 0x5a, 0x10, 0xf8, 0x75, 0xda, 0xbd, 0x2b, 0xa4, 0x79, 0x97, 0xc7,
	0xe1, 0x01, 0x4c, 0xa8, 0x33, 0xa7, 0xfc, 0x6c, 0xeb, 0xd1, 0x98, 0x9a, 0xa3, 0x94, 0xeb, 0x70,
	0x17, 0xa4, 0x8a, 0xe6, 0xf8, 0xe0, 0x96, 0x93, 0xe1, 0x95, 0x72, 0x0c, 0x0a, 0x91, 0xe7, 0x1a,
	0x49, 0xec, 0xd5, 0x25, 0x1d, 0x8b, 0x3d, 0xbc, 0x96, 0x57, 0x13, 0x4e, 0xac, 0x6e, 0x8d, 0x58,
	0xd5, 0x06, 0x4e, 0xad, 0x45, 0xec, 0xe0, 0xef, 0x0a, 0x5e, 0xb5, 0xb3, 0xc8, 0x57, 0x84, 0xaf,
	0x57, 0x4d, 0x21, 0xf7, 0x6a, 0x1a, 0x36, 0xd9, 0xea, 0xdd, 0x3f, 0x5f, 0x71, 0xf9, 0x84, 0x60,
	0xe7, 0xd3, 0x8f, 0x3f, 0x5f, 0x2e, 0x6d, 0x91, 0x1e, 0x5d, 0xde, 0xa7, 0xf9, 0x2a, 0xbb, 0xfb,
	0xbe, 0x75, 0x97, 0x7c, 0x47, 0x98, 0x2c, 0x3b, 0x43, 0xfa, 0x4d, 0x63, 0x1b, 0xb7, 0xc0, 0x1b,
	0x5c, 0x04, 0xe2, 0xf8, 0x3e, 0xb4, 0x7c, 0x77, 0x48, 0xf8, 0x1f, 0xbe, 0x70, 0x06, 0xdf, 0x77,
	0x2e, 0x93, 0x6f, 0x08, 0xaf, 0x2d, 0x98, 0x48, 0x1a, 0x75, 0xaa, 0x5b, 0x04, 0x6f, 0xfb, 0x9c,
	0xd5, 0x17, 0x90, 0x75, 0xc1, 0xfe, 0xe1, 0x93, 0xe3, 0xa9, 0x8f, 0x4e, 0xa6, 0x3e, 0xfa, 0x3d,
	0xf5, 0xd1, 0xe7, 0x99, 0xdf, 0x3a, 0x99, 0xf9, 0xad, 0x9f, 0x33, 0xbf, 0xf5, 0x26, 0xac, 0xac,
	0xf4, 0x9e, 0xcc, 0xd8, 0x2e, 0x64, 0x9c, 0x6a, 0x3e, 0x66, 0x92, 0x7e, 0x5c, 0xec, 0x6c, 0xd7,
	0x3b, 0xbe, 0x62, 0x7f, 0xfd, 0x07, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x3f, 0x4e, 0x08,
	0xcf, 0x04, 0x00, 0x00,
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
	StakingPools(ctx context.Context, in *QueryStakingPoolsRequest, opts ...grpc.CallOption) (*QueryStakingPoolsResponse, error)
	OutstandingRewards(ctx context.Context, in *QueryOutstandingRewardsRequest, opts ...grpc.CallOption) (*QueryOutstandingRewardsResponse, error)
	Undelegations(ctx context.Context, in *QueryUndelegationsRequest, opts ...grpc.CallOption) (*QueryUndelegationsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) StakingPools(ctx context.Context, in *QueryStakingPoolsRequest, opts ...grpc.CallOption) (*QueryStakingPoolsResponse, error) {
	out := new(QueryStakingPoolsResponse)
	err := c.cc.Invoke(ctx, "/kira.multistaking.Query/StakingPools", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) OutstandingRewards(ctx context.Context, in *QueryOutstandingRewardsRequest, opts ...grpc.CallOption) (*QueryOutstandingRewardsResponse, error) {
	out := new(QueryOutstandingRewardsResponse)
	err := c.cc.Invoke(ctx, "/kira.multistaking.Query/OutstandingRewards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Undelegations(ctx context.Context, in *QueryUndelegationsRequest, opts ...grpc.CallOption) (*QueryUndelegationsResponse, error) {
	out := new(QueryUndelegationsResponse)
	err := c.cc.Invoke(ctx, "/kira.multistaking.Query/Undelegations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	StakingPools(context.Context, *QueryStakingPoolsRequest) (*QueryStakingPoolsResponse, error)
	OutstandingRewards(context.Context, *QueryOutstandingRewardsRequest) (*QueryOutstandingRewardsResponse, error)
	Undelegations(context.Context, *QueryUndelegationsRequest) (*QueryUndelegationsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) StakingPools(ctx context.Context, req *QueryStakingPoolsRequest) (*QueryStakingPoolsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StakingPools not implemented")
}
func (*UnimplementedQueryServer) OutstandingRewards(ctx context.Context, req *QueryOutstandingRewardsRequest) (*QueryOutstandingRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OutstandingRewards not implemented")
}
func (*UnimplementedQueryServer) Undelegations(ctx context.Context, req *QueryUndelegationsRequest) (*QueryUndelegationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Undelegations not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_StakingPools_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStakingPoolsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).StakingPools(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kira.multistaking.Query/StakingPools",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).StakingPools(ctx, req.(*QueryStakingPoolsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_OutstandingRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryOutstandingRewardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).OutstandingRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kira.multistaking.Query/OutstandingRewards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).OutstandingRewards(ctx, req.(*QueryOutstandingRewardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Undelegations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUndelegationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Undelegations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kira.multistaking.Query/Undelegations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Undelegations(ctx, req.(*QueryUndelegationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kira.multistaking.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StakingPools",
			Handler:    _Query_StakingPools_Handler,
		},
		{
			MethodName: "OutstandingRewards",
			Handler:    _Query_OutstandingRewards_Handler,
		},
		{
			MethodName: "Undelegations",
			Handler:    _Query_Undelegations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kira/multistaking/query.proto",
}

func (m *QueryStakingPoolsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryStakingPoolsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryStakingPoolsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryStakingPoolsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryStakingPoolsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryStakingPoolsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Pools) > 0 {
		for iNdEx := len(m.Pools) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Pools[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryOutstandingRewardsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryOutstandingRewardsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryOutstandingRewardsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Delegator) > 0 {
		i -= len(m.Delegator)
		copy(dAtA[i:], m.Delegator)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Delegator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryOutstandingRewardsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryOutstandingRewardsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryOutstandingRewardsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Rewards) > 0 {
		for iNdEx := len(m.Rewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.Rewards[iNdEx].Size()
				i -= size
				if _, err := m.Rewards[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryUndelegationsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryUndelegationsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryUndelegationsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryUndelegationsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryUndelegationsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryUndelegationsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Undelegations) > 0 {
		for iNdEx := len(m.Undelegations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Undelegations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
func (m *QueryStakingPoolsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryStakingPoolsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Pools) > 0 {
		for _, e := range m.Pools {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryOutstandingRewardsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Delegator)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryOutstandingRewardsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Rewards) > 0 {
		for _, e := range m.Rewards {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryUndelegationsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryUndelegationsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Undelegations) > 0 {
		for _, e := range m.Undelegations {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryStakingPoolsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryStakingPoolsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryStakingPoolsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *QueryStakingPoolsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryStakingPoolsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryStakingPoolsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pools", wireType)
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
			m.Pools = append(m.Pools, StakingPool{})
			if err := m.Pools[len(m.Pools)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryOutstandingRewardsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryOutstandingRewardsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryOutstandingRewardsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Delegator = string(dAtA[iNdEx:postIndex])
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
func (m *QueryOutstandingRewardsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryOutstandingRewardsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryOutstandingRewardsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Coin
			m.Rewards = append(m.Rewards, v)
			if err := m.Rewards[len(m.Rewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryUndelegationsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryUndelegationsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryUndelegationsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *QueryUndelegationsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryUndelegationsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryUndelegationsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Undelegations", wireType)
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
			m.Undelegations = append(m.Undelegations, Undelegation{})
			if err := m.Undelegations[len(m.Undelegations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
