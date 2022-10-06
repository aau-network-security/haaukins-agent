// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: agent.proto

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

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentClient interface {
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	CreateEnvironment(ctx context.Context, in *CreatEnvRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	CreateLabForEnv(ctx context.Context, in *CreateLabRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	LabStream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Agent_LabStreamClient, error)
}

type agentClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentClient(cc grpc.ClientConnInterface) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/agent.Agent/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) CreateEnvironment(ctx context.Context, in *CreatEnvRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/agent.Agent/CreateEnvironment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) CreateLabForEnv(ctx context.Context, in *CreateLabRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/agent.Agent/CreateLabForEnv", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) LabStream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Agent_LabStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Agent_ServiceDesc.Streams[0], "/agent.Agent/LabStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentLabStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Agent_LabStreamClient interface {
	Recv() (*Lab, error)
	grpc.ClientStream
}

type agentLabStreamClient struct {
	grpc.ClientStream
}

func (x *agentLabStreamClient) Recv() (*Lab, error) {
	m := new(Lab)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AgentServer is the server API for Agent service.
// All implementations must embed UnimplementedAgentServer
// for forward compatibility
type AgentServer interface {
	Init(context.Context, *InitRequest) (*StatusResponse, error)
	CreateEnvironment(context.Context, *CreatEnvRequest) (*StatusResponse, error)
	CreateLabForEnv(context.Context, *CreateLabRequest) (*StatusResponse, error)
	LabStream(*Empty, Agent_LabStreamServer) error
	mustEmbedUnimplementedAgentServer()
}

// UnimplementedAgentServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (UnimplementedAgentServer) Init(context.Context, *InitRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedAgentServer) CreateEnvironment(context.Context, *CreatEnvRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEnvironment not implemented")
}
func (UnimplementedAgentServer) CreateLabForEnv(context.Context, *CreateLabRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLabForEnv not implemented")
}
func (UnimplementedAgentServer) LabStream(*Empty, Agent_LabStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method LabStream not implemented")
}
func (UnimplementedAgentServer) mustEmbedUnimplementedAgentServer() {}

// UnsafeAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServer will
// result in compilation errors.
type UnsafeAgentServer interface {
	mustEmbedUnimplementedAgentServer()
}

func RegisterAgentServer(s grpc.ServiceRegistrar, srv AgentServer) {
	s.RegisterService(&Agent_ServiceDesc, srv)
}

func _Agent_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.Agent/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_CreateEnvironment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatEnvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).CreateEnvironment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.Agent/CreateEnvironment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).CreateEnvironment(ctx, req.(*CreatEnvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_CreateLabForEnv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLabRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).CreateLabForEnv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.Agent/CreateLabForEnv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).CreateLabForEnv(ctx, req.(*CreateLabRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_LabStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgentServer).LabStream(m, &agentLabStreamServer{stream})
}

type Agent_LabStreamServer interface {
	Send(*Lab) error
	grpc.ServerStream
}

type agentLabStreamServer struct {
	grpc.ServerStream
}

func (x *agentLabStreamServer) Send(m *Lab) error {
	return x.ServerStream.SendMsg(m)
}

// Agent_ServiceDesc is the grpc.ServiceDesc for Agent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Agent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "agent.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _Agent_Init_Handler,
		},
		{
			MethodName: "CreateEnvironment",
			Handler:    _Agent_CreateEnvironment_Handler,
		},
		{
			MethodName: "CreateLabForEnv",
			Handler:    _Agent_CreateLabForEnv_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LabStream",
			Handler:       _Agent_LabStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "agent.proto",
}
