// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// SvcPodServClient is the client API for SvcPodServ service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SvcPodServClient interface {
	AskForHelp(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type svcPodServClient struct {
	cc grpc.ClientConnInterface
}

func NewSvcPodServClient(cc grpc.ClientConnInterface) SvcPodServClient {
	return &svcPodServClient{cc}
}

func (c *svcPodServClient) AskForHelp(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/svcpodserv.SvcPodServ/AskForHelp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SvcPodServServer is the server API for SvcPodServ service.
// All implementations must embed UnimplementedSvcPodServServer
// for forward compatibility
type SvcPodServServer interface {
	AskForHelp(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedSvcPodServServer()
}

// UnimplementedSvcPodServServer must be embedded to have forward compatible implementations.
type UnimplementedSvcPodServServer struct {
}

func (UnimplementedSvcPodServServer) AskForHelp(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AskForHelp not implemented")
}
func (UnimplementedSvcPodServServer) mustEmbedUnimplementedSvcPodServServer() {}

// UnsafeSvcPodServServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SvcPodServServer will
// result in compilation errors.
type UnsafeSvcPodServServer interface {
	mustEmbedUnimplementedSvcPodServServer()
}

func RegisterSvcPodServServer(s grpc.ServiceRegistrar, srv SvcPodServServer) {
	s.RegisterService(&SvcPodServ_ServiceDesc, srv)
}

func _SvcPodServ_AskForHelp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SvcPodServServer).AskForHelp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/svcpodserv.SvcPodServ/AskForHelp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SvcPodServServer).AskForHelp(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// SvcPodServ_ServiceDesc is the grpc.ServiceDesc for SvcPodServ service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SvcPodServ_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svcpodserv.SvcPodServ",
	HandlerType: (*SvcPodServServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AskForHelp",
			Handler:    _SvcPodServ_AskForHelp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svcpodserv.proto",
}
