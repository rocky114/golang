// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package student

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

// StudentManagerClient is the client API for StudentManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudentManagerClient interface {
	Echo(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error)
}

type studentManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentManagerClient(cc grpc.ClientConnInterface) StudentManagerClient {
	return &studentManagerClient{cc}
}

func (c *studentManagerClient) Echo(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error) {
	out := new(StringMessage)
	err := c.cc.Invoke(ctx, "/student.StudentManager/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentManagerServer is the server API for StudentManager service.
// All implementations must embed UnimplementedStudentManagerServer
// for forward compatibility
type StudentManagerServer interface {
	Echo(context.Context, *StringMessage) (*StringMessage, error)
	mustEmbedUnimplementedStudentManagerServer()
}

// UnimplementedStudentManagerServer must be embedded to have forward compatible implementations.
type UnimplementedStudentManagerServer struct {
}

func (UnimplementedStudentManagerServer) Echo(context.Context, *StringMessage) (*StringMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedStudentManagerServer) mustEmbedUnimplementedStudentManagerServer() {}

// UnsafeStudentManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudentManagerServer will
// result in compilation errors.
type UnsafeStudentManagerServer interface {
	mustEmbedUnimplementedStudentManagerServer()
}

func RegisterStudentManagerServer(s grpc.ServiceRegistrar, srv StudentManagerServer) {
	s.RegisterService(&StudentManager_ServiceDesc, srv)
}

func _StudentManager_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentManagerServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentManager/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentManagerServer).Echo(ctx, req.(*StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// StudentManager_ServiceDesc is the grpc.ServiceDesc for StudentManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StudentManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "student.StudentManager",
	HandlerType: (*StudentManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _StudentManager_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "student.proto",
}
