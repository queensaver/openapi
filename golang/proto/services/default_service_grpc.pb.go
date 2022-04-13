// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: services/default_service.proto

package services

import (
	context "context"
	models "github.com/queensaver/openapi/golang/proto/models"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DefaultServiceClient is the client API for DefaultService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DefaultServiceClient interface {
	BboxesGet(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*BboxesGetResponse, error)
	HivesDelete(ctx context.Context, in *HivesDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	HivesGet(ctx context.Context, in *HivesGetRequest, opts ...grpc.CallOption) (*HivesGetResponse, error)
	HivesPost(ctx context.Context, in *HivesPostRequest, opts ...grpc.CallOption) (*models.Hive, error)
	HivesPut(ctx context.Context, in *HivesPutRequest, opts ...grpc.CallOption) (*models.Hive, error)
	LoginPost(ctx context.Context, in *LoginPostRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ScaleGet(ctx context.Context, in *ScaleGetRequest, opts ...grpc.CallOption) (*ScaleGetResponse, error)
	TemperatureGet(ctx context.Context, in *TemperatureGetRequest, opts ...grpc.CallOption) (*TemperatureGetResponse, error)
	TemperaturePost(ctx context.Context, in *TemperaturePostRequest, opts ...grpc.CallOption) (*models.GenericPostResponse, error)
	UserPost(ctx context.Context, in *UserPostRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	VarroaScanGet(ctx context.Context, in *VarroaScanGetRequest, opts ...grpc.CallOption) (*models.VarroaScanResponse, error)
	VarroaScanImagePost(ctx context.Context, in *VarroaScanImagePostRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	VarroaScanPost(ctx context.Context, in *VarroaScanPostRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type defaultServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDefaultServiceClient(cc grpc.ClientConnInterface) DefaultServiceClient {
	return &defaultServiceClient{cc}
}

func (c *defaultServiceClient) BboxesGet(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*BboxesGetResponse, error) {
	out := new(BboxesGetResponse)
	err := c.cc.Invoke(ctx, "/openapi.services.defaultservice.DefaultService/BboxesGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultServiceClient) HivesDelete(ctx context.Context, in *HivesDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/openapi.services.defaultservice.DefaultService/HivesDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultServiceClient) HivesGet(ctx context.Context, in *HivesGetRequest, opts ...grpc.CallOption) (*HivesGetResponse, error) {
	out := new(HivesGetResponse)
	err := c.cc.Invoke(ctx, "/openapi.services.defaultservice.DefaultService/HivesGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultServiceClient) HivesPost(ctx context.Context, in *HivesPostRequest, opts ...grpc.CallOption) (*models.Hive, error) {
	out := new(models.Hive)
	err := c.cc.Invoke(ctx, "/openapi.services.defaultservice.DefaultService/HivesPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultServiceClient) HivesPut(ctx context.Context, in *HivesPutRequest, opts ...grpc.CallOption) (*models.Hive, error) {
	out := new(models.Hive)
	err := c.cc.Invoke(ctx, "/openapi.services.defaultservice.DefaultService/HivesPut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultServiceClient) LoginPost(ctx context.Context, in *LoginPostRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/openapi.services.defaultservice.DefaultService/LoginPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultServiceClient) ScaleGet(ctx context.Context, in *ScaleGetRequest, opts ...grpc.CallOption) (*ScaleGetResponse, error) {
	out := new(ScaleGetResponse)
	err := c.cc.Invoke(ctx, "/openapi.services.defaultservice.DefaultService/ScaleGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultServiceClient) TemperatureGet(ctx context.Context, in *TemperatureGetRequest, opts ...grpc.CallOption) (*TemperatureGetResponse, error) {
	out := new(TemperatureGetResponse)
	err := c.cc.Invoke(ctx, "/openapi.services.defaultservice.DefaultService/TemperatureGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultServiceClient) TemperaturePost(ctx context.Context, in *TemperaturePostRequest, opts ...grpc.CallOption) (*models.GenericPostResponse, error) {
	out := new(models.GenericPostResponse)
	err := c.cc.Invoke(ctx, "/openapi.services.defaultservice.DefaultService/TemperaturePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultServiceClient) UserPost(ctx context.Context, in *UserPostRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/openapi.services.defaultservice.DefaultService/UserPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultServiceClient) VarroaScanGet(ctx context.Context, in *VarroaScanGetRequest, opts ...grpc.CallOption) (*models.VarroaScanResponse, error) {
	out := new(models.VarroaScanResponse)
	err := c.cc.Invoke(ctx, "/openapi.services.defaultservice.DefaultService/VarroaScanGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultServiceClient) VarroaScanImagePost(ctx context.Context, in *VarroaScanImagePostRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/openapi.services.defaultservice.DefaultService/VarroaScanImagePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultServiceClient) VarroaScanPost(ctx context.Context, in *VarroaScanPostRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/openapi.services.defaultservice.DefaultService/VarroaScanPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DefaultServiceServer is the server API for DefaultService service.
// All implementations must embed UnimplementedDefaultServiceServer
// for forward compatibility
type DefaultServiceServer interface {
	BboxesGet(context.Context, *emptypb.Empty) (*BboxesGetResponse, error)
	HivesDelete(context.Context, *HivesDeleteRequest) (*emptypb.Empty, error)
	HivesGet(context.Context, *HivesGetRequest) (*HivesGetResponse, error)
	HivesPost(context.Context, *HivesPostRequest) (*models.Hive, error)
	HivesPut(context.Context, *HivesPutRequest) (*models.Hive, error)
	LoginPost(context.Context, *LoginPostRequest) (*emptypb.Empty, error)
	ScaleGet(context.Context, *ScaleGetRequest) (*ScaleGetResponse, error)
	TemperatureGet(context.Context, *TemperatureGetRequest) (*TemperatureGetResponse, error)
	TemperaturePost(context.Context, *TemperaturePostRequest) (*models.GenericPostResponse, error)
	UserPost(context.Context, *UserPostRequest) (*emptypb.Empty, error)
	VarroaScanGet(context.Context, *VarroaScanGetRequest) (*models.VarroaScanResponse, error)
	VarroaScanImagePost(context.Context, *VarroaScanImagePostRequest) (*emptypb.Empty, error)
	VarroaScanPost(context.Context, *VarroaScanPostRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedDefaultServiceServer()
}

// UnimplementedDefaultServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDefaultServiceServer struct {
}

func (UnimplementedDefaultServiceServer) BboxesGet(context.Context, *emptypb.Empty) (*BboxesGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BboxesGet not implemented")
}
func (UnimplementedDefaultServiceServer) HivesDelete(context.Context, *HivesDeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HivesDelete not implemented")
}
func (UnimplementedDefaultServiceServer) HivesGet(context.Context, *HivesGetRequest) (*HivesGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HivesGet not implemented")
}
func (UnimplementedDefaultServiceServer) HivesPost(context.Context, *HivesPostRequest) (*models.Hive, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HivesPost not implemented")
}
func (UnimplementedDefaultServiceServer) HivesPut(context.Context, *HivesPutRequest) (*models.Hive, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HivesPut not implemented")
}
func (UnimplementedDefaultServiceServer) LoginPost(context.Context, *LoginPostRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginPost not implemented")
}
func (UnimplementedDefaultServiceServer) ScaleGet(context.Context, *ScaleGetRequest) (*ScaleGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScaleGet not implemented")
}
func (UnimplementedDefaultServiceServer) TemperatureGet(context.Context, *TemperatureGetRequest) (*TemperatureGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TemperatureGet not implemented")
}
func (UnimplementedDefaultServiceServer) TemperaturePost(context.Context, *TemperaturePostRequest) (*models.GenericPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TemperaturePost not implemented")
}
func (UnimplementedDefaultServiceServer) UserPost(context.Context, *UserPostRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserPost not implemented")
}
func (UnimplementedDefaultServiceServer) VarroaScanGet(context.Context, *VarroaScanGetRequest) (*models.VarroaScanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VarroaScanGet not implemented")
}
func (UnimplementedDefaultServiceServer) VarroaScanImagePost(context.Context, *VarroaScanImagePostRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VarroaScanImagePost not implemented")
}
func (UnimplementedDefaultServiceServer) VarroaScanPost(context.Context, *VarroaScanPostRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VarroaScanPost not implemented")
}
func (UnimplementedDefaultServiceServer) mustEmbedUnimplementedDefaultServiceServer() {}

// UnsafeDefaultServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DefaultServiceServer will
// result in compilation errors.
type UnsafeDefaultServiceServer interface {
	mustEmbedUnimplementedDefaultServiceServer()
}

func RegisterDefaultServiceServer(s grpc.ServiceRegistrar, srv DefaultServiceServer) {
	s.RegisterService(&DefaultService_ServiceDesc, srv)
}

func _DefaultService_BboxesGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).BboxesGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openapi.services.defaultservice.DefaultService/BboxesGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).BboxesGet(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DefaultService_HivesDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HivesDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).HivesDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openapi.services.defaultservice.DefaultService/HivesDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).HivesDelete(ctx, req.(*HivesDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DefaultService_HivesGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HivesGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).HivesGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openapi.services.defaultservice.DefaultService/HivesGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).HivesGet(ctx, req.(*HivesGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DefaultService_HivesPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HivesPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).HivesPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openapi.services.defaultservice.DefaultService/HivesPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).HivesPost(ctx, req.(*HivesPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DefaultService_HivesPut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HivesPutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).HivesPut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openapi.services.defaultservice.DefaultService/HivesPut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).HivesPut(ctx, req.(*HivesPutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DefaultService_LoginPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).LoginPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openapi.services.defaultservice.DefaultService/LoginPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).LoginPost(ctx, req.(*LoginPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DefaultService_ScaleGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScaleGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).ScaleGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openapi.services.defaultservice.DefaultService/ScaleGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).ScaleGet(ctx, req.(*ScaleGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DefaultService_TemperatureGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TemperatureGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).TemperatureGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openapi.services.defaultservice.DefaultService/TemperatureGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).TemperatureGet(ctx, req.(*TemperatureGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DefaultService_TemperaturePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TemperaturePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).TemperaturePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openapi.services.defaultservice.DefaultService/TemperaturePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).TemperaturePost(ctx, req.(*TemperaturePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DefaultService_UserPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).UserPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openapi.services.defaultservice.DefaultService/UserPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).UserPost(ctx, req.(*UserPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DefaultService_VarroaScanGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VarroaScanGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).VarroaScanGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openapi.services.defaultservice.DefaultService/VarroaScanGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).VarroaScanGet(ctx, req.(*VarroaScanGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DefaultService_VarroaScanImagePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VarroaScanImagePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).VarroaScanImagePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openapi.services.defaultservice.DefaultService/VarroaScanImagePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).VarroaScanImagePost(ctx, req.(*VarroaScanImagePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DefaultService_VarroaScanPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VarroaScanPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).VarroaScanPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openapi.services.defaultservice.DefaultService/VarroaScanPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).VarroaScanPost(ctx, req.(*VarroaScanPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DefaultService_ServiceDesc is the grpc.ServiceDesc for DefaultService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DefaultService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openapi.services.defaultservice.DefaultService",
	HandlerType: (*DefaultServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BboxesGet",
			Handler:    _DefaultService_BboxesGet_Handler,
		},
		{
			MethodName: "HivesDelete",
			Handler:    _DefaultService_HivesDelete_Handler,
		},
		{
			MethodName: "HivesGet",
			Handler:    _DefaultService_HivesGet_Handler,
		},
		{
			MethodName: "HivesPost",
			Handler:    _DefaultService_HivesPost_Handler,
		},
		{
			MethodName: "HivesPut",
			Handler:    _DefaultService_HivesPut_Handler,
		},
		{
			MethodName: "LoginPost",
			Handler:    _DefaultService_LoginPost_Handler,
		},
		{
			MethodName: "ScaleGet",
			Handler:    _DefaultService_ScaleGet_Handler,
		},
		{
			MethodName: "TemperatureGet",
			Handler:    _DefaultService_TemperatureGet_Handler,
		},
		{
			MethodName: "TemperaturePost",
			Handler:    _DefaultService_TemperaturePost_Handler,
		},
		{
			MethodName: "UserPost",
			Handler:    _DefaultService_UserPost_Handler,
		},
		{
			MethodName: "VarroaScanGet",
			Handler:    _DefaultService_VarroaScanGet_Handler,
		},
		{
			MethodName: "VarroaScanImagePost",
			Handler:    _DefaultService_VarroaScanImagePost_Handler,
		},
		{
			MethodName: "VarroaScanPost",
			Handler:    _DefaultService_VarroaScanPost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/default_service.proto",
}
