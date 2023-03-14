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
	CloseEnvironment(ctx context.Context, in *CloseEnvRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	CreateLabForEnv(ctx context.Context, in *CreateLabRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	CreateVpnConfForLab(ctx context.Context, in *CreateVpnConfRequest, opts ...grpc.CallOption) (*CreateVpnConfResponse, error)
	CloseLab(ctx context.Context, in *CloseLabRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	AddExercisesToEnv(ctx context.Context, in *ExerciseRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	AddExercisesToLab(ctx context.Context, in *ExerciseRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	ResetExerciseInLab(ctx context.Context, in *ExerciseRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	StartExerciseInLab(ctx context.Context, in *ExerciseRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	StopExerciseInLab(ctx context.Context, in *ExerciseRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	MonitorStream(ctx context.Context, opts ...grpc.CallOption) (Agent_MonitorStreamClient, error)
	GetLab(ctx context.Context, in *GetLabRequest, opts ...grpc.CallOption) (*GetLabResponse, error)
	GetHostsInLab(ctx context.Context, in *GetHostsRequest, opts ...grpc.CallOption) (*GetHostsResponse, error)
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

func (c *agentClient) CloseEnvironment(ctx context.Context, in *CloseEnvRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/agent.Agent/CloseEnvironment", in, out, opts...)
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

func (c *agentClient) CreateVpnConfForLab(ctx context.Context, in *CreateVpnConfRequest, opts ...grpc.CallOption) (*CreateVpnConfResponse, error) {
	out := new(CreateVpnConfResponse)
	err := c.cc.Invoke(ctx, "/agent.Agent/CreateVpnConfForLab", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) CloseLab(ctx context.Context, in *CloseLabRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/agent.Agent/CloseLab", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) AddExercisesToEnv(ctx context.Context, in *ExerciseRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/agent.Agent/AddExercisesToEnv", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) AddExercisesToLab(ctx context.Context, in *ExerciseRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/agent.Agent/AddExercisesToLab", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ResetExerciseInLab(ctx context.Context, in *ExerciseRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/agent.Agent/ResetExerciseInLab", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) StartExerciseInLab(ctx context.Context, in *ExerciseRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/agent.Agent/StartExerciseInLab", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) StopExerciseInLab(ctx context.Context, in *ExerciseRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/agent.Agent/StopExerciseInLab", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/agent.Agent/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) MonitorStream(ctx context.Context, opts ...grpc.CallOption) (Agent_MonitorStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Agent_ServiceDesc.Streams[0], "/agent.Agent/MonitorStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentMonitorStreamClient{stream}
	return x, nil
}

type Agent_MonitorStreamClient interface {
	Send(*PingRequest) error
	Recv() (*MonitorResponse, error)
	grpc.ClientStream
}

type agentMonitorStreamClient struct {
	grpc.ClientStream
}

func (x *agentMonitorStreamClient) Send(m *PingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentMonitorStreamClient) Recv() (*MonitorResponse, error) {
	m := new(MonitorResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentClient) GetLab(ctx context.Context, in *GetLabRequest, opts ...grpc.CallOption) (*GetLabResponse, error) {
	out := new(GetLabResponse)
	err := c.cc.Invoke(ctx, "/agent.Agent/GetLab", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetHostsInLab(ctx context.Context, in *GetHostsRequest, opts ...grpc.CallOption) (*GetHostsResponse, error) {
	out := new(GetHostsResponse)
	err := c.cc.Invoke(ctx, "/agent.Agent/GetHostsInLab", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServer is the server API for Agent service.
// All implementations must embed UnimplementedAgentServer
// for forward compatibility
type AgentServer interface {
	Init(context.Context, *InitRequest) (*StatusResponse, error)
	CreateEnvironment(context.Context, *CreatEnvRequest) (*StatusResponse, error)
	CloseEnvironment(context.Context, *CloseEnvRequest) (*StatusResponse, error)
	CreateLabForEnv(context.Context, *CreateLabRequest) (*StatusResponse, error)
	CreateVpnConfForLab(context.Context, *CreateVpnConfRequest) (*CreateVpnConfResponse, error)
	CloseLab(context.Context, *CloseLabRequest) (*StatusResponse, error)
	AddExercisesToEnv(context.Context, *ExerciseRequest) (*StatusResponse, error)
	AddExercisesToLab(context.Context, *ExerciseRequest) (*StatusResponse, error)
	ResetExerciseInLab(context.Context, *ExerciseRequest) (*StatusResponse, error)
	StartExerciseInLab(context.Context, *ExerciseRequest) (*StatusResponse, error)
	StopExerciseInLab(context.Context, *ExerciseRequest) (*StatusResponse, error)
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	MonitorStream(Agent_MonitorStreamServer) error
	GetLab(context.Context, *GetLabRequest) (*GetLabResponse, error)
	GetHostsInLab(context.Context, *GetHostsRequest) (*GetHostsResponse, error)
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
func (UnimplementedAgentServer) CloseEnvironment(context.Context, *CloseEnvRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseEnvironment not implemented")
}
func (UnimplementedAgentServer) CreateLabForEnv(context.Context, *CreateLabRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLabForEnv not implemented")
}
func (UnimplementedAgentServer) CreateVpnConfForLab(context.Context, *CreateVpnConfRequest) (*CreateVpnConfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVpnConfForLab not implemented")
}
func (UnimplementedAgentServer) CloseLab(context.Context, *CloseLabRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseLab not implemented")
}
func (UnimplementedAgentServer) AddExercisesToEnv(context.Context, *ExerciseRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddExercisesToEnv not implemented")
}
func (UnimplementedAgentServer) AddExercisesToLab(context.Context, *ExerciseRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddExercisesToLab not implemented")
}
func (UnimplementedAgentServer) ResetExerciseInLab(context.Context, *ExerciseRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetExerciseInLab not implemented")
}
func (UnimplementedAgentServer) StartExerciseInLab(context.Context, *ExerciseRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartExerciseInLab not implemented")
}
func (UnimplementedAgentServer) StopExerciseInLab(context.Context, *ExerciseRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopExerciseInLab not implemented")
}
func (UnimplementedAgentServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedAgentServer) MonitorStream(Agent_MonitorStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method MonitorStream not implemented")
}
func (UnimplementedAgentServer) GetLab(context.Context, *GetLabRequest) (*GetLabResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLab not implemented")
}
func (UnimplementedAgentServer) GetHostsInLab(context.Context, *GetHostsRequest) (*GetHostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHostsInLab not implemented")
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

func _Agent_CloseEnvironment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseEnvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).CloseEnvironment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.Agent/CloseEnvironment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).CloseEnvironment(ctx, req.(*CloseEnvRequest))
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

func _Agent_CreateVpnConfForLab_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVpnConfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).CreateVpnConfForLab(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.Agent/CreateVpnConfForLab",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).CreateVpnConfForLab(ctx, req.(*CreateVpnConfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_CloseLab_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseLabRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).CloseLab(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.Agent/CloseLab",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).CloseLab(ctx, req.(*CloseLabRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_AddExercisesToEnv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExerciseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).AddExercisesToEnv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.Agent/AddExercisesToEnv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).AddExercisesToEnv(ctx, req.(*ExerciseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_AddExercisesToLab_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExerciseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).AddExercisesToLab(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.Agent/AddExercisesToLab",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).AddExercisesToLab(ctx, req.(*ExerciseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ResetExerciseInLab_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExerciseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ResetExerciseInLab(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.Agent/ResetExerciseInLab",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ResetExerciseInLab(ctx, req.(*ExerciseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_StartExerciseInLab_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExerciseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).StartExerciseInLab(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.Agent/StartExerciseInLab",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).StartExerciseInLab(ctx, req.(*ExerciseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_StopExerciseInLab_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExerciseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).StopExerciseInLab(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.Agent/StopExerciseInLab",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).StopExerciseInLab(ctx, req.(*ExerciseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.Agent/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_MonitorStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentServer).MonitorStream(&agentMonitorStreamServer{stream})
}

type Agent_MonitorStreamServer interface {
	Send(*MonitorResponse) error
	Recv() (*PingRequest, error)
	grpc.ServerStream
}

type agentMonitorStreamServer struct {
	grpc.ServerStream
}

func (x *agentMonitorStreamServer) Send(m *MonitorResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentMonitorStreamServer) Recv() (*PingRequest, error) {
	m := new(PingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Agent_GetLab_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLabRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetLab(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.Agent/GetLab",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetLab(ctx, req.(*GetLabRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetHostsInLab_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetHostsInLab(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.Agent/GetHostsInLab",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetHostsInLab(ctx, req.(*GetHostsRequest))
	}
	return interceptor(ctx, in, info, handler)
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
			MethodName: "CloseEnvironment",
			Handler:    _Agent_CloseEnvironment_Handler,
		},
		{
			MethodName: "CreateLabForEnv",
			Handler:    _Agent_CreateLabForEnv_Handler,
		},
		{
			MethodName: "CreateVpnConfForLab",
			Handler:    _Agent_CreateVpnConfForLab_Handler,
		},
		{
			MethodName: "CloseLab",
			Handler:    _Agent_CloseLab_Handler,
		},
		{
			MethodName: "AddExercisesToEnv",
			Handler:    _Agent_AddExercisesToEnv_Handler,
		},
		{
			MethodName: "AddExercisesToLab",
			Handler:    _Agent_AddExercisesToLab_Handler,
		},
		{
			MethodName: "ResetExerciseInLab",
			Handler:    _Agent_ResetExerciseInLab_Handler,
		},
		{
			MethodName: "StartExerciseInLab",
			Handler:    _Agent_StartExerciseInLab_Handler,
		},
		{
			MethodName: "StopExerciseInLab",
			Handler:    _Agent_StopExerciseInLab_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Agent_Ping_Handler,
		},
		{
			MethodName: "GetLab",
			Handler:    _Agent_GetLab_Handler,
		},
		{
			MethodName: "GetHostsInLab",
			Handler:    _Agent_GetHostsInLab_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MonitorStream",
			Handler:       _Agent_MonitorStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "agent.proto",
}
