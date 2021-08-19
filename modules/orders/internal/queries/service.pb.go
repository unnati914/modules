// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/modules/orders/internal/queries/service.proto

package queries

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	order "github.com/persistenceOne/persistenceSDK/modules/orders/internal/queries/order"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("persistence_sdk/modules/orders/internal/queries/service.proto", fileDescriptor_7ac98c7eb4014811)
}

var fileDescriptor_7ac98c7eb4014811 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x2d, 0x48, 0x2d, 0x2a,
	0xce, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x4e, 0x8d, 0x2f, 0x4e, 0xc9, 0xd6, 0xcf, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0xcf, 0x2f, 0x4a, 0x49, 0x2d, 0x2a, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d,
	0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x2c, 0x4d, 0x2d, 0xca, 0x4c, 0x2d, 0xd6, 0x2f, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xd2, 0x47, 0xd3, 0xae, 0x07, 0xd5,
	0xae, 0x07, 0xd1, 0xae, 0x07, 0xd3, 0xae, 0x07, 0xd5, 0x2e, 0xe5, 0x48, 0xaa, 0x7d, 0x60, 0x71,
	0x30, 0xaf, 0x12, 0x62, 0xa7, 0x94, 0x4c, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x7e, 0x62, 0x41,
	0xa6, 0x7e, 0x62, 0x5e, 0x5e, 0x7e, 0x49, 0x62, 0x49, 0x66, 0x7e, 0x5e, 0x31, 0x44, 0xd6, 0xe8,
	0x11, 0x23, 0x17, 0x6b, 0x20, 0x48, 0xb5, 0xd0, 0x0d, 0x46, 0x2e, 0x56, 0x7f, 0x90, 0x6e, 0xa1,
	0x60, 0x3d, 0x74, 0x67, 0x92, 0xe8, 0x6c, 0x88, 0xb8, 0x1e, 0xd8, 0xdc, 0xa0, 0xd4, 0xc2, 0xd2,
	0xd4, 0xe2, 0x12, 0xa9, 0x10, 0xea, 0x1a, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0xaa, 0xa4, 0xd2,
	0x74, 0xf9, 0xc9, 0x64, 0x26, 0x39, 0x21, 0x19, 0xe4, 0x90, 0x0d, 0x76, 0xf1, 0x86, 0x85, 0x0f,
	0x98, 0x72, 0x4a, 0x3a, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18,
	0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x8f, 0xf4,
	0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x64, 0x13, 0xfc, 0xf3, 0x52, 0xd1, 0x0d,
	0x24, 0x10, 0xf0, 0x49, 0x6c, 0xe0, 0xf0, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x2f,
	0x04, 0x6b, 0x22, 0x02, 0x00, 0x00,
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
	Order(ctx context.Context, in *order.QueryRequest, opts ...grpc.CallOption) (*order.QueryResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Order(ctx context.Context, in *order.QueryRequest, opts ...grpc.CallOption) (*order.QueryResponse, error) {
	out := new(order.QueryResponse)
	err := c.cc.Invoke(ctx, "/persistence_sdk.modules.orders.internal.queries.Query/Order", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	Order(context.Context, *order.QueryRequest) (*order.QueryResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Order(ctx context.Context, req *order.QueryRequest) (*order.QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Order not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Order_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(order.QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Order(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/persistence_sdk.modules.orders.internal.queries.Query/Order",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Order(ctx, req.(*order.QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "persistence_sdk.modules.orders.internal.queries.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Order",
			Handler:    _Query_Order_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "persistence_sdk/modules/orders/internal/queries/service.proto",
}
