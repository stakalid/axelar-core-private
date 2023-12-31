// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axelar/tss/v1beta1/service.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/axelarnetwork/axelar-core/x/snapshot/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("axelar/tss/v1beta1/service.proto", fileDescriptor_03d0b5f2955aa837) }
func init() {
	golang_proto.RegisterFile("axelar/tss/v1beta1/service.proto", fileDescriptor_03d0b5f2955aa837)
}

var fileDescriptor_03d0b5f2955aa837 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xbf, 0x4b, 0xfb, 0x50,
	0x10, 0xc0, 0xfb, 0xbe, 0x43, 0xe1, 0x1b, 0x9c, 0x82, 0x2e, 0xb1, 0x3c, 0x6a, 0xd0, 0x45, 0x30,
	0xcf, 0xd6, 0xcd, 0xb1, 0x93, 0x8b, 0xe2, 0x8f, 0xcd, 0xed, 0x52, 0x8e, 0x34, 0xd8, 0xe6, 0xd2,
	0x77, 0xd7, 0xda, 0x3a, 0xea, 0xe2, 0x28, 0xf8, 0x0f, 0x39, 0x3a, 0x16, 0x5c, 0x1c, 0xa5, 0xf1,
	0x0f, 0x91, 0x26, 0x69, 0x10, 0x0c, 0xba, 0x25, 0x7c, 0x3e, 0xdc, 0x7d, 0xde, 0x39, 0x6d, 0x98,
	0xe1, 0x10, 0xac, 0x11, 0x66, 0x33, 0xed, 0x84, 0x28, 0xd0, 0x31, 0x8c, 0x76, 0x1a, 0xf7, 0x31,
	0x48, 0x2d, 0x09, 0xb9, 0x6e, 0x61, 0x04, 0xc2, 0x1c, 0x94, 0x86, 0xb7, 0x19, 0x51, 0x44, 0x39,
	0x36, 0xab, 0xaf, 0xc2, 0xf4, 0x5a, 0x11, 0x51, 0x34, 0x44, 0x03, 0x69, 0x6c, 0x20, 0x49, 0x48,
	0x40, 0x62, 0x4a, 0xb8, 0xa4, 0xeb, 0x4d, 0x9c, 0x40, 0xca, 0x03, 0x92, 0x6a, 0x9d, 0xcc, 0x4a,
	0x63, 0xbb, 0xa6, 0xa5, 0x82, 0xba, 0x06, 0x8e, 0x27, 0x68, 0xe7, 0x05, 0xef, 0x3e, 0x2a, 0xc7,
	0x39, 0xe5, 0xe8, 0xaa, 0x68, 0x77, 0xef, 0x9c, 0xff, 0x27, 0x08, 0x56, 0x7a, 0x08, 0xe2, 0xee,
	0x06, 0x3f, 0xdf, 0x10, 0x54, 0xf8, 0x12, 0xc7, 0x13, 0x64, 0xf1, 0xf6, 0xfe, 0xb0, 0x38, 0xa5,
	0x84, 0xd1, 0x6f, 0xdf, 0xbf, 0x7d, 0x3e, 0xff, 0xf3, 0xfc, 0x2d, 0xf3, 0xad, 0x68, 0xb0, 0xd2,
	0x42, 0x04, 0x39, 0x56, 0xfb, 0xdd, 0x07, 0xe5, 0x6c, 0x5c, 0xac, 0xd2, 0xd6, 0x31, 0xec, 0x34,
	0xcf, 0xc1, 0xc2, 0x88, 0xdd, 0x9d, 0xba, 0x1d, 0x05, 0x5b, 0x67, 0xf8, 0xbf, 0x29, 0x65, 0x83,
	0x9f, 0x37, 0xb4, 0x5c, 0xcf, 0xd4, 0x5c, 0x25, 0xcd, 0xdd, 0xde, 0xd9, 0xeb, 0x52, 0xab, 0xc5,
	0x52, 0xab, 0x8f, 0xa5, 0x56, 0x4f, 0x99, 0x6e, 0xbc, 0x64, 0x5a, 0x2d, 0x32, 0xdd, 0x78, 0xcf,
	0x74, 0xe3, 0xfa, 0x30, 0x8a, 0x65, 0x30, 0x09, 0x83, 0x3e, 0x8d, 0xca, 0x19, 0x09, 0xca, 0x2d,
	0xd9, 0x9b, 0xf2, 0xef, 0xa0, 0x4f, 0x16, 0xcd, 0x2c, 0x1f, 0x2c, 0xf3, 0x14, 0x39, 0x6c, 0xe6,
	0x77, 0x3e, 0xfa, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x8d, 0x0f, 0x24, 0x32, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgServiceClient interface {
	HeartBeat(ctx context.Context, in *HeartBeatRequest, opts ...grpc.CallOption) (*HeartBeatResponse, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) HeartBeat(ctx context.Context, in *HeartBeatRequest, opts ...grpc.CallOption) (*HeartBeatResponse, error) {
	out := new(HeartBeatResponse)
	err := c.cc.Invoke(ctx, "/axelar.tss.v1beta1.MsgService/HeartBeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	HeartBeat(context.Context, *HeartBeatRequest) (*HeartBeatResponse, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) HeartBeat(ctx context.Context, req *HeartBeatRequest) (*HeartBeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeartBeat not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_HeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartBeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).HeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.tss.v1beta1.MsgService/HeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).HeartBeat(ctx, req.(*HeartBeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "axelar.tss.v1beta1.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HeartBeat",
			Handler:    _MsgService_HeartBeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "axelar/tss/v1beta1/service.proto",
}

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryServiceClient interface {
	Params(ctx context.Context, in *ParamsRequest, opts ...grpc.CallOption) (*ParamsResponse, error)
}

type queryServiceClient struct {
	cc grpc1.ClientConn
}

func NewQueryServiceClient(cc grpc1.ClientConn) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) Params(ctx context.Context, in *ParamsRequest, opts ...grpc.CallOption) (*ParamsResponse, error) {
	out := new(ParamsResponse)
	err := c.cc.Invoke(ctx, "/axelar.tss.v1beta1.QueryService/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServiceServer is the server API for QueryService service.
type QueryServiceServer interface {
	Params(context.Context, *ParamsRequest) (*ParamsResponse, error)
}

// UnimplementedQueryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (*UnimplementedQueryServiceServer) Params(ctx context.Context, req *ParamsRequest) (*ParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}

func RegisterQueryServiceServer(s grpc1.Server, srv QueryServiceServer) {
	s.RegisterService(&_QueryService_serviceDesc, srv)
}

func _QueryService_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.tss.v1beta1.QueryService/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).Params(ctx, req.(*ParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "axelar.tss.v1beta1.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _QueryService_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "axelar/tss/v1beta1/service.proto",
}
