// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axelar/multisig/v1beta1/service.proto

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

func init() {
	proto.RegisterFile("axelar/multisig/v1beta1/service.proto", fileDescriptor_2f253d13b0297bdf)
}
func init() {
	golang_proto.RegisterFile("axelar/multisig/v1beta1/service.proto", fileDescriptor_2f253d13b0297bdf)
}

var fileDescriptor_2f253d13b0297bdf = []byte{
	// 634 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x7b, 0xa0, 0x16, 0xf5, 0x5a, 0x84, 0x74, 0x42, 0x42, 0x2a, 0xc5, 0x6d, 0xdd, 0x34,
	0xa1, 0x69, 0xea, 0x23, 0x01, 0x16, 0x46, 0xd4, 0x25, 0x8a, 0xa0, 0xa5, 0xde, 0x58, 0x22, 0x3b,
	0x9c, 0xdc, 0x53, 0x13, 0x9f, 0xeb, 0x3b, 0x87, 0x44, 0x88, 0xa5, 0x0b, 0x12, 0x13, 0x82, 0x01,
	0x24, 0x06, 0x24, 0x76, 0xfe, 0x07, 0x46, 0xc6, 0x0a, 0x16, 0x46, 0x94, 0xf0, 0x87, 0xa0, 0x3b,
	0x9f, 0x83, 0x4b, 0xe5, 0x1f, 0x6c, 0x89, 0xf2, 0xb9, 0x7b, 0x9f, 0xbc, 0xf7, 0x7d, 0x36, 0xdc,
	0x72, 0x46, 0xa4, 0xef, 0x84, 0x78, 0x10, 0xf5, 0x05, 0xe5, 0xd4, 0xc3, 0xc3, 0xa6, 0x4b, 0x84,
	0xd3, 0xc4, 0x9c, 0x84, 0x43, 0xda, 0x23, 0x56, 0x10, 0x32, 0xc1, 0xd0, 0x8d, 0x18, 0xb3, 0x12,
	0xcc, 0xd2, 0xd8, 0xca, 0x75, 0x8f, 0x79, 0x4c, 0x31, 0x58, 0x7e, 0x8a, 0xf1, 0x95, 0x55, 0x8f,
	0x31, 0xaf, 0x4f, 0xb0, 0x13, 0x50, 0xec, 0xf8, 0x3e, 0x13, 0x8e, 0xa0, 0xcc, 0xe7, 0xfa, 0xd7,
	0xf5, 0xac, 0x9a, 0x62, 0xa4, 0x89, 0xcd, 0x2c, 0xe2, 0x24, 0x22, 0xe1, 0x38, 0x86, 0x5a, 0x5f,
	0xae, 0x40, 0xf8, 0x88, 0x7b, 0x76, 0x2c, 0x8a, 0xde, 0x02, 0xb8, 0x64, 0x0b, 0x27, 0x14, 0x1d,
	0x32, 0xf6, 0x88, 0x8f, 0x76, 0xac, 0x0c, 0x67, 0x2b, 0x45, 0x1d, 0x92, 0x93, 0x88, 0x70, 0xb1,
	0xd2, 0x28, 0x07, 0xf3, 0x80, 0xf9, 0x9c, 0x98, 0xb7, 0x4f, 0x7f, 0xfc, 0x7e, 0x77, 0xc9, 0x34,
	0x6f, 0xe1, 0x7f, 0x3d, 0xb9, 0xa4, 0xbb, 0xc7, 0x0a, 0x7f, 0x00, 0xea, 0xe8, 0x3d, 0x80, 0xcb,
	0x76, 0xe4, 0x0e, 0xa8, 0x38, 0x88, 0xdc, 0x0e, 0x19, 0xa3, 0x9c, 0x42, 0x29, 0x2c, 0xd1, 0xda,
	0x2d, 0x49, 0x6b, 0xaf, 0xba, 0xf2, 0xaa, 0x98, 0x6b, 0x17, 0xbd, 0x14, 0xde, 0x0d, 0x22, 0x57,
	0xca, 0x49, 0xb3, 0xcf, 0x00, 0x5e, 0x8b, 0x2f, 0xb1, 0xa9, 0xe7, 0x3b, 0x22, 0x0a, 0x09, 0xc2,
	0x05, 0xe5, 0x66, 0x64, 0xe2, 0x77, 0xa7, 0xfc, 0x01, 0xad, 0xd8, 0x50, 0x8a, 0x55, 0x73, 0x23,
	0x4b, 0x91, 0x27, 0x47, 0xa4, 0xe4, 0x6b, 0x00, 0x17, 0x0f, 0x65, 0x7a, 0x88, 0xec, 0xdd, 0x76,
	0x66, 0xb5, 0x19, 0x93, 0x88, 0xd5, 0xcb, 0xa0, 0x5a, 0xa9, 0xaa, 0x94, 0xd6, 0xcd, 0x9b, 0x17,
	0x94, 0x42, 0xc5, 0x26, 0x1d, 0xfb, 0x04, 0xe0, 0x72, 0x1c, 0x84, 0xfd, 0x40, 0xec, 0x47, 0x22,
	0x67, 0x96, 0x69, 0xac, 0x78, 0x96, 0xe7, 0x69, 0x6d, 0xd5, 0x52, 0x56, 0x0d, 0xb3, 0x86, 0xb3,
	0x76, 0x21, 0x4e, 0x59, 0x97, 0x05, 0xa2, 0xcb, 0x22, 0x21, 0x0d, 0x3f, 0x02, 0xb8, 0x34, 0xbb,
	0xac, 0x9d, 0xb7, 0x02, 0x29, 0xaa, 0x78, 0x05, 0xce, 0xc1, 0x5a, 0xaf, 0xa9, 0xf4, 0x76, 0xcc,
	0x6a, 0x19, 0x3d, 0x2a, 0x77, 0xa1, 0xf5, 0x7d, 0x1e, 0x2e, 0x3f, 0x91, 0xfb, 0x9b, 0x6c, 0xec,
	0x2b, 0x00, 0xe7, 0x3b, 0x64, 0xdc, 0xde, 0x43, 0x5b, 0x79, 0xb5, 0xdb, 0x7b, 0x89, 0x62, 0xb5,
	0x08, 0xd3, 0x72, 0x58, 0xc9, 0x6d, 0xa3, 0xdc, 0xde, 0x75, 0xe9, 0x33, 0xfc, 0xa2, 0x77, 0xe4,
	0x50, 0xff, 0x25, 0xfa, 0x00, 0xe0, 0xe2, 0x63, 0x32, 0x12, 0xb1, 0x4d, 0x76, 0xce, 0x66, 0x4c,
	0x71, 0xce, 0x52, 0xa8, 0xb6, 0xba, 0xa7, 0xac, 0x2c, 0xd4, 0xc8, 0xb4, 0xf2, 0xc9, 0x48, 0x3d,
	0x3c, 0xd2, 0x6a, 0x43, 0x78, 0x59, 0x66, 0x7f, 0x33, 0xef, 0xaf, 0x27, 0x36, 0x95, 0x7c, 0x48,
	0x7b, 0x54, 0x94, 0x87, 0x81, 0x56, 0xf3, 0xba, 0x23, 0xd3, 0x7e, 0x35, 0x1e, 0xbc, 0x4d, 0x38,
	0xa7, 0xcc, 0x47, 0x45, 0x01, 0xd6, 0x5c, 0x22, 0x63, 0x95, 0xc5, 0xff, 0x67, 0x68, 0x32, 0x51,
	0x5c, 0xfb, 0x9c, 0x02, 0xb8, 0x70, 0xe0, 0x84, 0xce, 0x80, 0xa3, 0xec, 0x60, 0xc4, 0x40, 0xe2,
	0x54, 0x2b, 0xe4, 0xb4, 0x4c, 0x4d, 0xc9, 0x6c, 0xa0, 0xb5, 0x4c, 0x99, 0x40, 0x1d, 0x78, 0x68,
	0x7f, 0x9b, 0x18, 0xe0, 0x6c, 0x62, 0x80, 0x5f, 0x13, 0x03, 0xbc, 0x99, 0x1a, 0x73, 0x5f, 0xa7,
	0x06, 0x38, 0x9b, 0x1a, 0x73, 0x3f, 0xa7, 0xc6, 0xdc, 0xd3, 0xfb, 0x1e, 0x15, 0x47, 0x91, 0x6b,
	0xf5, 0xd8, 0x40, 0x5f, 0xe4, 0x13, 0xf1, 0x9c, 0x85, 0xc7, 0xfa, 0xdb, 0x6e, 0x8f, 0x85, 0x04,
	0x8f, 0xfe, 0xde, 0x2e, 0xc6, 0x01, 0xe1, 0xee, 0x82, 0x7a, 0xc1, 0xdd, 0xfd, 0x13, 0x00, 0x00,
	0xff, 0xff, 0x21, 0xd3, 0xd0, 0x4a, 0x9d, 0x07, 0x00, 0x00,
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
	StartKeygen(ctx context.Context, in *StartKeygenRequest, opts ...grpc.CallOption) (*StartKeygenResponse, error)
	SubmitPubKey(ctx context.Context, in *SubmitPubKeyRequest, opts ...grpc.CallOption) (*SubmitPubKeyResponse, error)
	SubmitSignature(ctx context.Context, in *SubmitSignatureRequest, opts ...grpc.CallOption) (*SubmitSignatureResponse, error)
	RotateKey(ctx context.Context, in *RotateKeyRequest, opts ...grpc.CallOption) (*RotateKeyResponse, error)
	KeygenOptOut(ctx context.Context, in *KeygenOptOutRequest, opts ...grpc.CallOption) (*KeygenOptOutResponse, error)
	KeygenOptIn(ctx context.Context, in *KeygenOptInRequest, opts ...grpc.CallOption) (*KeygenOptInResponse, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) StartKeygen(ctx context.Context, in *StartKeygenRequest, opts ...grpc.CallOption) (*StartKeygenResponse, error) {
	out := new(StartKeygenResponse)
	err := c.cc.Invoke(ctx, "/axelar.multisig.v1beta1.MsgService/StartKeygen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) SubmitPubKey(ctx context.Context, in *SubmitPubKeyRequest, opts ...grpc.CallOption) (*SubmitPubKeyResponse, error) {
	out := new(SubmitPubKeyResponse)
	err := c.cc.Invoke(ctx, "/axelar.multisig.v1beta1.MsgService/SubmitPubKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) SubmitSignature(ctx context.Context, in *SubmitSignatureRequest, opts ...grpc.CallOption) (*SubmitSignatureResponse, error) {
	out := new(SubmitSignatureResponse)
	err := c.cc.Invoke(ctx, "/axelar.multisig.v1beta1.MsgService/SubmitSignature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) RotateKey(ctx context.Context, in *RotateKeyRequest, opts ...grpc.CallOption) (*RotateKeyResponse, error) {
	out := new(RotateKeyResponse)
	err := c.cc.Invoke(ctx, "/axelar.multisig.v1beta1.MsgService/RotateKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) KeygenOptOut(ctx context.Context, in *KeygenOptOutRequest, opts ...grpc.CallOption) (*KeygenOptOutResponse, error) {
	out := new(KeygenOptOutResponse)
	err := c.cc.Invoke(ctx, "/axelar.multisig.v1beta1.MsgService/KeygenOptOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) KeygenOptIn(ctx context.Context, in *KeygenOptInRequest, opts ...grpc.CallOption) (*KeygenOptInResponse, error) {
	out := new(KeygenOptInResponse)
	err := c.cc.Invoke(ctx, "/axelar.multisig.v1beta1.MsgService/KeygenOptIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	StartKeygen(context.Context, *StartKeygenRequest) (*StartKeygenResponse, error)
	SubmitPubKey(context.Context, *SubmitPubKeyRequest) (*SubmitPubKeyResponse, error)
	SubmitSignature(context.Context, *SubmitSignatureRequest) (*SubmitSignatureResponse, error)
	RotateKey(context.Context, *RotateKeyRequest) (*RotateKeyResponse, error)
	KeygenOptOut(context.Context, *KeygenOptOutRequest) (*KeygenOptOutResponse, error)
	KeygenOptIn(context.Context, *KeygenOptInRequest) (*KeygenOptInResponse, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) StartKeygen(ctx context.Context, req *StartKeygenRequest) (*StartKeygenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartKeygen not implemented")
}
func (*UnimplementedMsgServiceServer) SubmitPubKey(ctx context.Context, req *SubmitPubKeyRequest) (*SubmitPubKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitPubKey not implemented")
}
func (*UnimplementedMsgServiceServer) SubmitSignature(ctx context.Context, req *SubmitSignatureRequest) (*SubmitSignatureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitSignature not implemented")
}
func (*UnimplementedMsgServiceServer) RotateKey(ctx context.Context, req *RotateKeyRequest) (*RotateKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RotateKey not implemented")
}
func (*UnimplementedMsgServiceServer) KeygenOptOut(ctx context.Context, req *KeygenOptOutRequest) (*KeygenOptOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeygenOptOut not implemented")
}
func (*UnimplementedMsgServiceServer) KeygenOptIn(ctx context.Context, req *KeygenOptInRequest) (*KeygenOptInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeygenOptIn not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_StartKeygen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartKeygenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).StartKeygen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.multisig.v1beta1.MsgService/StartKeygen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).StartKeygen(ctx, req.(*StartKeygenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_SubmitPubKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitPubKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).SubmitPubKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.multisig.v1beta1.MsgService/SubmitPubKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).SubmitPubKey(ctx, req.(*SubmitPubKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_SubmitSignature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitSignatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).SubmitSignature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.multisig.v1beta1.MsgService/SubmitSignature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).SubmitSignature(ctx, req.(*SubmitSignatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_RotateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RotateKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).RotateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.multisig.v1beta1.MsgService/RotateKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).RotateKey(ctx, req.(*RotateKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_KeygenOptOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeygenOptOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).KeygenOptOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.multisig.v1beta1.MsgService/KeygenOptOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).KeygenOptOut(ctx, req.(*KeygenOptOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_KeygenOptIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeygenOptInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).KeygenOptIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.multisig.v1beta1.MsgService/KeygenOptIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).KeygenOptIn(ctx, req.(*KeygenOptInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "axelar.multisig.v1beta1.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartKeygen",
			Handler:    _MsgService_StartKeygen_Handler,
		},
		{
			MethodName: "SubmitPubKey",
			Handler:    _MsgService_SubmitPubKey_Handler,
		},
		{
			MethodName: "SubmitSignature",
			Handler:    _MsgService_SubmitSignature_Handler,
		},
		{
			MethodName: "RotateKey",
			Handler:    _MsgService_RotateKey_Handler,
		},
		{
			MethodName: "KeygenOptOut",
			Handler:    _MsgService_KeygenOptOut_Handler,
		},
		{
			MethodName: "KeygenOptIn",
			Handler:    _MsgService_KeygenOptIn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "axelar/multisig/v1beta1/service.proto",
}

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryServiceClient interface {
	// KeyID returns the key ID of a key assigned to a given chain.
	// If no key is assigned, it returns the grpc NOT_FOUND error.
	KeyID(ctx context.Context, in *KeyIDRequest, opts ...grpc.CallOption) (*KeyIDResponse, error)
	// NextKeyID returns the key ID assigned for the next rotation on a given
	// chain. If no key rotation is in progress, it returns the grpc NOT_FOUND
	// error.
	NextKeyID(ctx context.Context, in *NextKeyIDRequest, opts ...grpc.CallOption) (*NextKeyIDResponse, error)
	// Key returns the key corresponding to a given key ID.
	// If no key is found, it returns the grpc NOT_FOUND error.
	Key(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*KeyResponse, error)
	// KeygenSession returns the keygen session info for a given key ID.
	// If no key is found, it returns the grpc NOT_FOUND error.
	KeygenSession(ctx context.Context, in *KeygenSessionRequest, opts ...grpc.CallOption) (*KeygenSessionResponse, error)
	Params(ctx context.Context, in *ParamsRequest, opts ...grpc.CallOption) (*ParamsResponse, error)
}

type queryServiceClient struct {
	cc grpc1.ClientConn
}

func NewQueryServiceClient(cc grpc1.ClientConn) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) KeyID(ctx context.Context, in *KeyIDRequest, opts ...grpc.CallOption) (*KeyIDResponse, error) {
	out := new(KeyIDResponse)
	err := c.cc.Invoke(ctx, "/axelar.multisig.v1beta1.QueryService/KeyID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) NextKeyID(ctx context.Context, in *NextKeyIDRequest, opts ...grpc.CallOption) (*NextKeyIDResponse, error) {
	out := new(NextKeyIDResponse)
	err := c.cc.Invoke(ctx, "/axelar.multisig.v1beta1.QueryService/NextKeyID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) Key(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*KeyResponse, error) {
	out := new(KeyResponse)
	err := c.cc.Invoke(ctx, "/axelar.multisig.v1beta1.QueryService/Key", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) KeygenSession(ctx context.Context, in *KeygenSessionRequest, opts ...grpc.CallOption) (*KeygenSessionResponse, error) {
	out := new(KeygenSessionResponse)
	err := c.cc.Invoke(ctx, "/axelar.multisig.v1beta1.QueryService/KeygenSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) Params(ctx context.Context, in *ParamsRequest, opts ...grpc.CallOption) (*ParamsResponse, error) {
	out := new(ParamsResponse)
	err := c.cc.Invoke(ctx, "/axelar.multisig.v1beta1.QueryService/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServiceServer is the server API for QueryService service.
type QueryServiceServer interface {
	// KeyID returns the key ID of a key assigned to a given chain.
	// If no key is assigned, it returns the grpc NOT_FOUND error.
	KeyID(context.Context, *KeyIDRequest) (*KeyIDResponse, error)
	// NextKeyID returns the key ID assigned for the next rotation on a given
	// chain. If no key rotation is in progress, it returns the grpc NOT_FOUND
	// error.
	NextKeyID(context.Context, *NextKeyIDRequest) (*NextKeyIDResponse, error)
	// Key returns the key corresponding to a given key ID.
	// If no key is found, it returns the grpc NOT_FOUND error.
	Key(context.Context, *KeyRequest) (*KeyResponse, error)
	// KeygenSession returns the keygen session info for a given key ID.
	// If no key is found, it returns the grpc NOT_FOUND error.
	KeygenSession(context.Context, *KeygenSessionRequest) (*KeygenSessionResponse, error)
	Params(context.Context, *ParamsRequest) (*ParamsResponse, error)
}

// UnimplementedQueryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (*UnimplementedQueryServiceServer) KeyID(ctx context.Context, req *KeyIDRequest) (*KeyIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeyID not implemented")
}
func (*UnimplementedQueryServiceServer) NextKeyID(ctx context.Context, req *NextKeyIDRequest) (*NextKeyIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NextKeyID not implemented")
}
func (*UnimplementedQueryServiceServer) Key(ctx context.Context, req *KeyRequest) (*KeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Key not implemented")
}
func (*UnimplementedQueryServiceServer) KeygenSession(ctx context.Context, req *KeygenSessionRequest) (*KeygenSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeygenSession not implemented")
}
func (*UnimplementedQueryServiceServer) Params(ctx context.Context, req *ParamsRequest) (*ParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}

func RegisterQueryServiceServer(s grpc1.Server, srv QueryServiceServer) {
	s.RegisterService(&_QueryService_serviceDesc, srv)
}

func _QueryService_KeyID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).KeyID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.multisig.v1beta1.QueryService/KeyID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).KeyID(ctx, req.(*KeyIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_NextKeyID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NextKeyIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).NextKeyID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.multisig.v1beta1.QueryService/NextKeyID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).NextKeyID(ctx, req.(*NextKeyIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_Key_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).Key(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.multisig.v1beta1.QueryService/Key",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).Key(ctx, req.(*KeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_KeygenSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeygenSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).KeygenSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.multisig.v1beta1.QueryService/KeygenSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).KeygenSession(ctx, req.(*KeygenSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
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
		FullMethod: "/axelar.multisig.v1beta1.QueryService/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).Params(ctx, req.(*ParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "axelar.multisig.v1beta1.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "KeyID",
			Handler:    _QueryService_KeyID_Handler,
		},
		{
			MethodName: "NextKeyID",
			Handler:    _QueryService_NextKeyID_Handler,
		},
		{
			MethodName: "Key",
			Handler:    _QueryService_Key_Handler,
		},
		{
			MethodName: "KeygenSession",
			Handler:    _QueryService_KeygenSession_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _QueryService_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "axelar/multisig/v1beta1/service.proto",
}
