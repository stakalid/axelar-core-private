// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nexus/v1beta1/service.proto

package types

import (
	context "context"
	fmt "fmt"
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

func init() { proto.RegisterFile("nexus/v1beta1/service.proto", fileDescriptor_e8a22d972057ace6) }
func init() {
	golang_proto.RegisterFile("nexus/v1beta1/service.proto", fileDescriptor_e8a22d972057ace6)
}

var fileDescriptor_e8a22d972057ace6 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0x4a, 0x03, 0x31,
	0x18, 0xc7, 0x9b, 0x0e, 0x0e, 0x07, 0x2e, 0x87, 0x28, 0x56, 0xc9, 0x50, 0x10, 0xa1, 0xd8, 0xa4,
	0x3d, 0x37, 0x47, 0x75, 0x2d, 0x48, 0xdd, 0xdc, 0x72, 0xe7, 0x47, 0x1a, 0xac, 0xf9, 0xce, 0x24,
	0x57, 0xcf, 0xd5, 0x27, 0x10, 0x7c, 0x0b, 0x47, 0x9f, 0xc0, 0xd1, 0xb1, 0xe0, 0xe2, 0x28, 0x77,
	0xce, 0x3e, 0x83, 0x78, 0x17, 0x87, 0x13, 0x0f, 0xba, 0x25, 0xf9, 0xff, 0xbe, 0xef, 0xff, 0x83,
	0x04, 0x3b, 0x1a, 0xf2, 0xcc, 0xf2, 0xc5, 0x38, 0x06, 0x27, 0xc6, 0xdc, 0x82, 0x59, 0xa8, 0x04,
	0x58, 0x6a, 0xd0, 0x61, 0xb8, 0x5e, 0x85, 0xcc, 0x87, 0xbd, 0x0d, 0x89, 0x12, 0xab, 0x84, 0xff,
	0x9c, 0x6a, 0xa8, 0xb7, 0x2b, 0x11, 0xe5, 0x1c, 0xb8, 0x48, 0x15, 0x17, 0x5a, 0xa3, 0x13, 0x4e,
	0xa1, 0xb6, 0x3e, 0xdd, 0x6c, 0xee, 0x77, 0x79, 0xfd, 0x1e, 0x7d, 0x75, 0x83, 0x60, 0x62, 0xe5,
	0x79, 0xdd, 0x17, 0x3e, 0x91, 0x60, 0x6b, 0x0a, 0x52, 0x59, 0x07, 0xe6, 0x64, 0x26, 0x94, 0x9e,
	0x08, 0xa5, 0x9d, 0x50, 0x1a, 0x4c, 0x38, 0x64, 0x0d, 0x0d, 0xd6, 0xc2, 0x4d, 0xe1, 0x26, 0x03,
	0xeb, 0x7a, 0x6c, 0x55, 0xdc, 0xa6, 0xa8, 0x2d, 0xf4, 0x47, 0xf7, 0x6f, 0x9f, 0x8f, 0xdd, 0x41,
	0x7f, 0x8f, 0x8b, 0x1c, 0xe6, 0xc2, 0xf0, 0xda, 0xd8, 0xfc, 0x3f, 0x76, 0x44, 0x06, 0xe1, 0x33,
	0x09, 0xb6, 0x4f, 0xa1, 0x05, 0x08, 0xf9, 0x9f, 0xfe, 0x56, 0xf2, 0x57, 0x78, 0xb4, 0xfa, 0x80,
	0x57, 0x8e, 0x2a, 0xe5, 0x83, 0xfe, 0x7e, 0x53, 0xf9, 0x12, 0xda, 0xa5, 0x8f, 0xcf, 0x5e, 0x0b,
	0x4a, 0x96, 0x05, 0x25, 0x1f, 0x05, 0x25, 0x0f, 0x25, 0xed, 0xbc, 0x94, 0x94, 0x2c, 0x4b, 0xda,
	0x79, 0x2f, 0x69, 0xe7, 0x22, 0x92, 0xca, 0xcd, 0xb2, 0x98, 0x25, 0x78, 0xed, 0x77, 0x6a, 0x70,
	0xb7, 0x68, 0xae, 0xfc, 0x6d, 0x98, 0xa0, 0x01, 0x9e, 0xfb, 0x22, 0x77, 0x97, 0x82, 0x8d, 0xd7,
	0xaa, 0x9f, 0x3c, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x4f, 0xac, 0x63, 0xbf, 0x43, 0x02, 0x00,
	0x00,
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
	RegisterChainMaintainer(ctx context.Context, in *RegisterChainMaintainerRequest, opts ...grpc.CallOption) (*RegisterChainMaintainerResponse, error)
	DeregisterChainMaintainer(ctx context.Context, in *DeregisterChainMaintainerRequest, opts ...grpc.CallOption) (*DeregisterChainMaintainerResponse, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) RegisterChainMaintainer(ctx context.Context, in *RegisterChainMaintainerRequest, opts ...grpc.CallOption) (*RegisterChainMaintainerResponse, error) {
	out := new(RegisterChainMaintainerResponse)
	err := c.cc.Invoke(ctx, "/nexus.v1beta1.MsgService/RegisterChainMaintainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) DeregisterChainMaintainer(ctx context.Context, in *DeregisterChainMaintainerRequest, opts ...grpc.CallOption) (*DeregisterChainMaintainerResponse, error) {
	out := new(DeregisterChainMaintainerResponse)
	err := c.cc.Invoke(ctx, "/nexus.v1beta1.MsgService/DeregisterChainMaintainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	RegisterChainMaintainer(context.Context, *RegisterChainMaintainerRequest) (*RegisterChainMaintainerResponse, error)
	DeregisterChainMaintainer(context.Context, *DeregisterChainMaintainerRequest) (*DeregisterChainMaintainerResponse, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) RegisterChainMaintainer(ctx context.Context, req *RegisterChainMaintainerRequest) (*RegisterChainMaintainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterChainMaintainer not implemented")
}
func (*UnimplementedMsgServiceServer) DeregisterChainMaintainer(ctx context.Context, req *DeregisterChainMaintainerRequest) (*DeregisterChainMaintainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeregisterChainMaintainer not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_RegisterChainMaintainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterChainMaintainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).RegisterChainMaintainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nexus.v1beta1.MsgService/RegisterChainMaintainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).RegisterChainMaintainer(ctx, req.(*RegisterChainMaintainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_DeregisterChainMaintainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeregisterChainMaintainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).DeregisterChainMaintainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nexus.v1beta1.MsgService/DeregisterChainMaintainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).DeregisterChainMaintainer(ctx, req.(*DeregisterChainMaintainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nexus.v1beta1.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterChainMaintainer",
			Handler:    _MsgService_RegisterChainMaintainer_Handler,
		},
		{
			MethodName: "DeregisterChainMaintainer",
			Handler:    _MsgService_DeregisterChainMaintainer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nexus/v1beta1/service.proto",
}