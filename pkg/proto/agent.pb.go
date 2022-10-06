// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: agent.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{0}
}

type InitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url        string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	AuthKey    string `protobuf:"bytes,2,opt,name=authKey,proto3" json:"authKey,omitempty"`
	SignKey    string `protobuf:"bytes,3,opt,name=signKey,proto3" json:"signKey,omitempty"`
	TlsEnabled bool   `protobuf:"varint,4,opt,name=tlsEnabled,proto3" json:"tlsEnabled,omitempty"`
}

func (x *InitRequest) Reset() {
	*x = InitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitRequest) ProtoMessage() {}

func (x *InitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitRequest.ProtoReflect.Descriptor instead.
func (*InitRequest) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{1}
}

func (x *InitRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *InitRequest) GetAuthKey() string {
	if x != nil {
		return x.AuthKey
	}
	return ""
}

func (x *InitRequest) GetSignKey() string {
	if x != nil {
		return x.SignKey
	}
	return ""
}

func (x *InitRequest) GetTlsEnabled() bool {
	if x != nil {
		return x.TlsEnabled
	}
	return false
}

type CreatEnvRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventTag    string      `protobuf:"bytes,1,opt,name=eventTag,proto3" json:"eventTag,omitempty"`
	EnvType     int32       `protobuf:"varint,2,opt,name=envType,proto3" json:"envType,omitempty"`
	Vms         []*VmConfig `protobuf:"bytes,3,rep,name=vms,proto3" json:"vms,omitempty"`
	InitialLabs int32       `protobuf:"varint,4,opt,name=initialLabs,proto3" json:"initialLabs,omitempty"`
	Exercises   []string    `protobuf:"bytes,5,rep,name=exercises,proto3" json:"exercises,omitempty"`
}

func (x *CreatEnvRequest) Reset() {
	*x = CreatEnvRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatEnvRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatEnvRequest) ProtoMessage() {}

func (x *CreatEnvRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatEnvRequest.ProtoReflect.Descriptor instead.
func (*CreatEnvRequest) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{2}
}

func (x *CreatEnvRequest) GetEventTag() string {
	if x != nil {
		return x.EventTag
	}
	return ""
}

func (x *CreatEnvRequest) GetEnvType() int32 {
	if x != nil {
		return x.EnvType
	}
	return 0
}

func (x *CreatEnvRequest) GetVms() []*VmConfig {
	if x != nil {
		return x.Vms
	}
	return nil
}

func (x *CreatEnvRequest) GetInitialLabs() int32 {
	if x != nil {
		return x.InitialLabs
	}
	return 0
}

func (x *CreatEnvRequest) GetExercises() []string {
	if x != nil {
		return x.Exercises
	}
	return nil
}

type CreateLabRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventTag string `protobuf:"bytes,1,opt,name=eventTag,proto3" json:"eventTag,omitempty"`
	IsVPN    bool   `protobuf:"varint,2,opt,name=isVPN,proto3" json:"isVPN,omitempty"`
}

func (x *CreateLabRequest) Reset() {
	*x = CreateLabRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLabRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLabRequest) ProtoMessage() {}

func (x *CreateLabRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLabRequest.ProtoReflect.Descriptor instead.
func (*CreateLabRequest) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{3}
}

func (x *CreateLabRequest) GetEventTag() string {
	if x != nil {
		return x.EventTag
	}
	return ""
}

func (x *CreateLabRequest) GetIsVPN() bool {
	if x != nil {
		return x.IsVPN
	}
	return false
}

type VmConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image    string  `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	MemoryMB uint32  `protobuf:"varint,2,opt,name=memoryMB,proto3" json:"memoryMB,omitempty"`
	Cpu      float64 `protobuf:"fixed64,3,opt,name=cpu,proto3" json:"cpu,omitempty"`
}

func (x *VmConfig) Reset() {
	*x = VmConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VmConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VmConfig) ProtoMessage() {}

func (x *VmConfig) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VmConfig.ProtoReflect.Descriptor instead.
func (*VmConfig) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{4}
}

func (x *VmConfig) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *VmConfig) GetMemoryMB() uint32 {
	if x != nil {
		return x.MemoryMB
	}
	return 0
}

func (x *VmConfig) GetCpu() float64 {
	if x != nil {
		return x.Cpu
	}
	return 0
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{5}
}

func (x *StatusResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Lab struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag       string      `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	EventTag  string      `protobuf:"bytes,2,opt,name=eventTag,proto3" json:"eventTag,omitempty"`
	Exercises []*Exercise `protobuf:"bytes,3,rep,name=exercises,proto3" json:"exercises,omitempty"`
	GuacPort  uint32      `protobuf:"varint,4,opt,name=guacPort,proto3" json:"guacPort,omitempty"`
	IsVPN     bool        `protobuf:"varint,5,opt,name=isVPN,proto3" json:"isVPN,omitempty"`
}

func (x *Lab) Reset() {
	*x = Lab{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lab) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lab) ProtoMessage() {}

func (x *Lab) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lab.ProtoReflect.Descriptor instead.
func (*Lab) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{6}
}

func (x *Lab) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Lab) GetEventTag() string {
	if x != nil {
		return x.EventTag
	}
	return ""
}

func (x *Lab) GetExercises() []*Exercise {
	if x != nil {
		return x.Exercises
	}
	return nil
}

func (x *Lab) GetGuacPort() uint32 {
	if x != nil {
		return x.GuacPort
	}
	return 0
}

func (x *Lab) GetIsVPN() bool {
	if x != nil {
		return x.IsVPN
	}
	return false
}

type Exercise struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag    string   `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Flags  []string `protobuf:"bytes,2,rep,name=flags,proto3" json:"flags,omitempty"`
	Errors []string `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *Exercise) Reset() {
	*x = Exercise{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Exercise) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Exercise) ProtoMessage() {}

func (x *Exercise) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Exercise.ProtoReflect.Descriptor instead.
func (*Exercise) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{7}
}

func (x *Exercise) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Exercise) GetFlags() []string {
	if x != nil {
		return x.Flags
	}
	return nil
}

func (x *Exercise) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_agent_proto protoreflect.FileDescriptor

var file_agent_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x73, 0x0a,
	0x0b, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x69, 0x67, 0x6e,
	0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x69, 0x67, 0x6e, 0x4b,
	0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6c, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x74, 0x6c, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x22, 0xaa, 0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x45, 0x6e, 0x76, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x61, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x76, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x6e, 0x76, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x03,
	0x76, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x56, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x03, 0x76, 0x6d, 0x73, 0x12,
	0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x4c, 0x61, 0x62, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x4c, 0x61, 0x62,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x22,
	0x44, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x67, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x73, 0x56, 0x50, 0x4e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x69, 0x73, 0x56, 0x50, 0x4e, 0x22, 0x4e, 0x0a, 0x08, 0x56, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x4d, 0x42, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x4d, 0x42, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x03, 0x63, 0x70, 0x75, 0x22, 0x2a, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x94, 0x01, 0x0a, 0x03, 0x4c, 0x61, 0x62, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x67, 0x12, 0x2d, 0x0a, 0x09, 0x65, 0x78, 0x65, 0x72, 0x63,
	0x69, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2e, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x09, 0x65, 0x78, 0x65,
	0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x75, 0x61, 0x63, 0x50, 0x6f,
	0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x67, 0x75, 0x61, 0x63, 0x50, 0x6f,
	0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x56, 0x50, 0x4e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x69, 0x73, 0x56, 0x50, 0x4e, 0x22, 0x4a, 0x0a, 0x08, 0x45, 0x78, 0x65, 0x72,
	0x63, 0x69, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x32, 0xf2, 0x01, 0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x33,
	0x0a, 0x04, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x12, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x49,
	0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x45, 0x6e, 0x76, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0f, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x46, 0x6f, 0x72, 0x45, 0x6e, 0x76, 0x12, 0x17, 0x2e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x29,
	0x0a, 0x09, 0x4c, 0x61, 0x62, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x0c, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x4c, 0x61, 0x62, 0x22, 0x00, 0x30, 0x01, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x61, 0x75, 0x2d, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2d, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x68, 0x61, 0x61,
	0x75, 0x6b, 0x69, 0x6e, 0x73, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_agent_proto_rawDescOnce sync.Once
	file_agent_proto_rawDescData = file_agent_proto_rawDesc
)

func file_agent_proto_rawDescGZIP() []byte {
	file_agent_proto_rawDescOnce.Do(func() {
		file_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_agent_proto_rawDescData)
	})
	return file_agent_proto_rawDescData
}

var file_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_agent_proto_goTypes = []interface{}{
	(*Empty)(nil),            // 0: agent.Empty
	(*InitRequest)(nil),      // 1: agent.InitRequest
	(*CreatEnvRequest)(nil),  // 2: agent.CreatEnvRequest
	(*CreateLabRequest)(nil), // 3: agent.CreateLabRequest
	(*VmConfig)(nil),         // 4: agent.VmConfig
	(*StatusResponse)(nil),   // 5: agent.StatusResponse
	(*Lab)(nil),              // 6: agent.Lab
	(*Exercise)(nil),         // 7: agent.Exercise
}
var file_agent_proto_depIdxs = []int32{
	4, // 0: agent.CreatEnvRequest.vms:type_name -> agent.VmConfig
	7, // 1: agent.Lab.exercises:type_name -> agent.Exercise
	1, // 2: agent.Agent.Init:input_type -> agent.InitRequest
	2, // 3: agent.Agent.CreateEnvironment:input_type -> agent.CreatEnvRequest
	3, // 4: agent.Agent.CreateLabForEnv:input_type -> agent.CreateLabRequest
	0, // 5: agent.Agent.LabStream:input_type -> agent.Empty
	5, // 6: agent.Agent.Init:output_type -> agent.StatusResponse
	5, // 7: agent.Agent.CreateEnvironment:output_type -> agent.StatusResponse
	5, // 8: agent.Agent.CreateLabForEnv:output_type -> agent.StatusResponse
	6, // 9: agent.Agent.LabStream:output_type -> agent.Lab
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_agent_proto_init() }
func file_agent_proto_init() {
	if File_agent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_agent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatEnvRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_agent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLabRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_agent_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VmConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_agent_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_agent_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lab); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_agent_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Exercise); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_agent_proto_goTypes,
		DependencyIndexes: file_agent_proto_depIdxs,
		MessageInfos:      file_agent_proto_msgTypes,
	}.Build()
	File_agent_proto = out.File
	file_agent_proto_rawDesc = nil
	file_agent_proto_goTypes = nil
	file_agent_proto_depIdxs = nil
}
