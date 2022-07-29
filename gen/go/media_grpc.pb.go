// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: media.proto

package protobufpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MediaSrvClient is the client API for MediaSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MediaSrvClient interface {
	GetAllMedia(ctx context.Context, in *GetAllMediaReq, opts ...grpc.CallOption) (*GetAllMediaRes, error)
	GetDetailMedia(ctx context.Context, in *GetDetailMediaReq, opts ...grpc.CallOption) (*GetDetailMediaRes, error)
	CreateMedia(ctx context.Context, in *CreateMediaReq, opts ...grpc.CallOption) (*CreateMediaRes, error)
	CreateMediaAgain(ctx context.Context, in *CreateMediaAgainReq, opts ...grpc.CallOption) (*CreateMediaRes, error)
	CreateCategory(ctx context.Context, in *CreateCategoryReq, opts ...grpc.CallOption) (*CreateCategoryRes, error)
	GetListCategory(ctx context.Context, in *GetListCategoryReq, opts ...grpc.CallOption) (*GetListCategoryRes, error)
	AddSignature(ctx context.Context, in *AddSignatureReq, opts ...grpc.CallOption) (*AddSignatureRes, error)
}

type mediaSrvClient struct {
	cc grpc.ClientConnInterface
}

func NewMediaSrvClient(cc grpc.ClientConnInterface) MediaSrvClient {
	return &mediaSrvClient{cc}
}

func (c *mediaSrvClient) GetAllMedia(ctx context.Context, in *GetAllMediaReq, opts ...grpc.CallOption) (*GetAllMediaRes, error) {
	out := new(GetAllMediaRes)
	err := c.cc.Invoke(ctx, "/protobuf.MediaSrv/GetAllMedia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaSrvClient) GetDetailMedia(ctx context.Context, in *GetDetailMediaReq, opts ...grpc.CallOption) (*GetDetailMediaRes, error) {
	out := new(GetDetailMediaRes)
	err := c.cc.Invoke(ctx, "/protobuf.MediaSrv/GetDetailMedia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaSrvClient) CreateMedia(ctx context.Context, in *CreateMediaReq, opts ...grpc.CallOption) (*CreateMediaRes, error) {
	out := new(CreateMediaRes)
	err := c.cc.Invoke(ctx, "/protobuf.MediaSrv/CreateMedia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaSrvClient) CreateMediaAgain(ctx context.Context, in *CreateMediaAgainReq, opts ...grpc.CallOption) (*CreateMediaRes, error) {
	out := new(CreateMediaRes)
	err := c.cc.Invoke(ctx, "/protobuf.MediaSrv/CreateMediaAgain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaSrvClient) CreateCategory(ctx context.Context, in *CreateCategoryReq, opts ...grpc.CallOption) (*CreateCategoryRes, error) {
	out := new(CreateCategoryRes)
	err := c.cc.Invoke(ctx, "/protobuf.MediaSrv/CreateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaSrvClient) GetListCategory(ctx context.Context, in *GetListCategoryReq, opts ...grpc.CallOption) (*GetListCategoryRes, error) {
	out := new(GetListCategoryRes)
	err := c.cc.Invoke(ctx, "/protobuf.MediaSrv/GetListCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaSrvClient) AddSignature(ctx context.Context, in *AddSignatureReq, opts ...grpc.CallOption) (*AddSignatureRes, error) {
	out := new(AddSignatureRes)
	err := c.cc.Invoke(ctx, "/protobuf.MediaSrv/AddSignature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MediaSrvServer is the server API for MediaSrv service.
// All implementations must embed UnimplementedMediaSrvServer
// for forward compatibility
type MediaSrvServer interface {
	GetAllMedia(context.Context, *GetAllMediaReq) (*GetAllMediaRes, error)
	GetDetailMedia(context.Context, *GetDetailMediaReq) (*GetDetailMediaRes, error)
	CreateMedia(context.Context, *CreateMediaReq) (*CreateMediaRes, error)
	CreateMediaAgain(context.Context, *CreateMediaAgainReq) (*CreateMediaRes, error)
	CreateCategory(context.Context, *CreateCategoryReq) (*CreateCategoryRes, error)
	GetListCategory(context.Context, *GetListCategoryReq) (*GetListCategoryRes, error)
	AddSignature(context.Context, *AddSignatureReq) (*AddSignatureRes, error)
	mustEmbedUnimplementedMediaSrvServer()
}

// UnimplementedMediaSrvServer must be embedded to have forward compatible implementations.
type UnimplementedMediaSrvServer struct {
}

func (UnimplementedMediaSrvServer) GetAllMedia(context.Context, *GetAllMediaReq) (*GetAllMediaRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllMedia not implemented")
}
func (UnimplementedMediaSrvServer) GetDetailMedia(context.Context, *GetDetailMediaReq) (*GetDetailMediaRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetailMedia not implemented")
}
func (UnimplementedMediaSrvServer) CreateMedia(context.Context, *CreateMediaReq) (*CreateMediaRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMedia not implemented")
}
func (UnimplementedMediaSrvServer) CreateMediaAgain(context.Context, *CreateMediaAgainReq) (*CreateMediaRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMediaAgain not implemented")
}
func (UnimplementedMediaSrvServer) CreateCategory(context.Context, *CreateCategoryReq) (*CreateCategoryRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (UnimplementedMediaSrvServer) GetListCategory(context.Context, *GetListCategoryReq) (*GetListCategoryRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListCategory not implemented")
}
func (UnimplementedMediaSrvServer) AddSignature(context.Context, *AddSignatureReq) (*AddSignatureRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSignature not implemented")
}
func (UnimplementedMediaSrvServer) mustEmbedUnimplementedMediaSrvServer() {}

// UnsafeMediaSrvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MediaSrvServer will
// result in compilation errors.
type UnsafeMediaSrvServer interface {
	mustEmbedUnimplementedMediaSrvServer()
}

func RegisterMediaSrvServer(s grpc.ServiceRegistrar, srv MediaSrvServer) {
	s.RegisterService(&MediaSrv_ServiceDesc, srv)
}

func _MediaSrv_GetAllMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllMediaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaSrvServer).GetAllMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MediaSrv/GetAllMedia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaSrvServer).GetAllMedia(ctx, req.(*GetAllMediaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaSrv_GetDetailMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDetailMediaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaSrvServer).GetDetailMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MediaSrv/GetDetailMedia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaSrvServer).GetDetailMedia(ctx, req.(*GetDetailMediaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaSrv_CreateMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMediaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaSrvServer).CreateMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MediaSrv/CreateMedia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaSrvServer).CreateMedia(ctx, req.(*CreateMediaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaSrv_CreateMediaAgain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMediaAgainReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaSrvServer).CreateMediaAgain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MediaSrv/CreateMediaAgain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaSrvServer).CreateMediaAgain(ctx, req.(*CreateMediaAgainReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaSrv_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCategoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaSrvServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MediaSrv/CreateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaSrvServer).CreateCategory(ctx, req.(*CreateCategoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaSrv_GetListCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListCategoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaSrvServer).GetListCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MediaSrv/GetListCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaSrvServer).GetListCategory(ctx, req.(*GetListCategoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaSrv_AddSignature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSignatureReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaSrvServer).AddSignature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MediaSrv/AddSignature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaSrvServer).AddSignature(ctx, req.(*AddSignatureReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MediaSrv_ServiceDesc is the grpc.ServiceDesc for MediaSrv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MediaSrv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.MediaSrv",
	HandlerType: (*MediaSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllMedia",
			Handler:    _MediaSrv_GetAllMedia_Handler,
		},
		{
			MethodName: "GetDetailMedia",
			Handler:    _MediaSrv_GetDetailMedia_Handler,
		},
		{
			MethodName: "CreateMedia",
			Handler:    _MediaSrv_CreateMedia_Handler,
		},
		{
			MethodName: "CreateMediaAgain",
			Handler:    _MediaSrv_CreateMediaAgain_Handler,
		},
		{
			MethodName: "CreateCategory",
			Handler:    _MediaSrv_CreateCategory_Handler,
		},
		{
			MethodName: "GetListCategory",
			Handler:    _MediaSrv_GetListCategory_Handler,
		},
		{
			MethodName: "AddSignature",
			Handler:    _MediaSrv_AddSignature_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "media.proto",
}
