// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: ticket.proto

package build

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

type TicketChannelType int32

const (
	TicketChannelType_MAIL TicketChannelType = 0
)

// Enum value maps for TicketChannelType.
var (
	TicketChannelType_name = map[int32]string{
		0: "MAIL",
	}
	TicketChannelType_value = map[string]int32{
		"MAIL": 0,
	}
)

func (x TicketChannelType) Enum() *TicketChannelType {
	p := new(TicketChannelType)
	*p = x
	return p
}

func (x TicketChannelType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TicketChannelType) Descriptor() protoreflect.EnumDescriptor {
	return file_ticket_proto_enumTypes[0].Descriptor()
}

func (TicketChannelType) Type() protoreflect.EnumType {
	return &file_ticket_proto_enumTypes[0]
}

func (x TicketChannelType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TicketChannelType.Descriptor instead.
func (TicketChannelType) EnumDescriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{0}
}

type TicketState int32

const (
	TicketState_CREATED    TicketState = 0
	TicketState_DELETED    TicketState = 1
	TicketState_IN_PROCESS TicketState = 2
	TicketState_OBSERVE    TicketState = 3
	TicketState_END        TicketState = 4
)

// Enum value maps for TicketState.
var (
	TicketState_name = map[int32]string{
		0: "CREATED",
		1: "DELETED",
		2: "IN_PROCESS",
		3: "OBSERVE",
		4: "END",
	}
	TicketState_value = map[string]int32{
		"CREATED":    0,
		"DELETED":    1,
		"IN_PROCESS": 2,
		"OBSERVE":    3,
		"END":        4,
	}
)

func (x TicketState) Enum() *TicketState {
	p := new(TicketState)
	*p = x
	return p
}

func (x TicketState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TicketState) Descriptor() protoreflect.EnumDescriptor {
	return file_ticket_proto_enumTypes[1].Descriptor()
}

func (TicketState) Type() protoreflect.EnumType {
	return &file_ticket_proto_enumTypes[1]
}

func (x TicketState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TicketState.Descriptor instead.
func (TicketState) EnumDescriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{1}
}

type CreateTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WriteUId    string            `protobuf:"bytes,1,opt,name=writeUId,proto3" json:"writeUId,omitempty"`
	TicketId    string            `protobuf:"bytes,2,opt,name=ticketId,proto3" json:"ticketId,omitempty"`
	ChannelType TicketChannelType `protobuf:"varint,3,opt,name=channelType,proto3,enum=TicketChannelType" json:"channelType,omitempty"`
	Requirement string            `protobuf:"bytes,4,opt,name=requirement,proto3" json:"requirement,omitempty"`
	Because     string            `protobuf:"bytes,5,opt,name=because,proto3" json:"because,omitempty"`
	State       *TicketState      `protobuf:"varint,6,opt,name=state,proto3,enum=TicketState,oneof" json:"state,omitempty"`
}

func (x *CreateTicketRequest) Reset() {
	*x = CreateTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTicketRequest) ProtoMessage() {}

func (x *CreateTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTicketRequest.ProtoReflect.Descriptor instead.
func (*CreateTicketRequest) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTicketRequest) GetWriteUId() string {
	if x != nil {
		return x.WriteUId
	}
	return ""
}

func (x *CreateTicketRequest) GetTicketId() string {
	if x != nil {
		return x.TicketId
	}
	return ""
}

func (x *CreateTicketRequest) GetChannelType() TicketChannelType {
	if x != nil {
		return x.ChannelType
	}
	return TicketChannelType_MAIL
}

func (x *CreateTicketRequest) GetRequirement() string {
	if x != nil {
		return x.Requirement
	}
	return ""
}

func (x *CreateTicketRequest) GetBecause() string {
	if x != nil {
		return x.Because
	}
	return ""
}

func (x *CreateTicketRequest) GetState() TicketState {
	if x != nil && x.State != nil {
		return *x.State
	}
	return TicketState_CREATED
}

type CreateTicketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok string `protobuf:"bytes,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *CreateTicketResponse) Reset() {
	*x = CreateTicketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTicketResponse) ProtoMessage() {}

func (x *CreateTicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTicketResponse.ProtoReflect.Descriptor instead.
func (*CreateTicketResponse) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTicketResponse) GetOk() string {
	if x != nil {
		return x.Ok
	}
	return ""
}

var File_ticket_proto protoreflect.FileDescriptor

var file_ticket_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf2,
	0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x72, 0x69, 0x74, 0x65, 0x55,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x72, 0x69, 0x74, 0x65, 0x55,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x34,
	0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x65, 0x63, 0x61, 0x75, 0x73,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x65, 0x63, 0x61, 0x75, 0x73, 0x65,
	0x12, 0x27, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0c, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x22, 0x26, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x6b, 0x2a, 0x1d, 0x0a, 0x11, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x00, 0x2a, 0x4d, 0x0a, 0x0b, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45,
	0x41, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45,
	0x44, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53,
	0x53, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x42, 0x53, 0x45, 0x52, 0x56, 0x45, 0x10, 0x03,
	0x12, 0x07, 0x0a, 0x03, 0x45, 0x4e, 0x44, 0x10, 0x04, 0x32, 0x45, 0x0a, 0x06, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x3b, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ticket_proto_rawDescOnce sync.Once
	file_ticket_proto_rawDescData = file_ticket_proto_rawDesc
)

func file_ticket_proto_rawDescGZIP() []byte {
	file_ticket_proto_rawDescOnce.Do(func() {
		file_ticket_proto_rawDescData = protoimpl.X.CompressGZIP(file_ticket_proto_rawDescData)
	})
	return file_ticket_proto_rawDescData
}

var file_ticket_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_ticket_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ticket_proto_goTypes = []any{
	(TicketChannelType)(0),       // 0: TicketChannelType
	(TicketState)(0),             // 1: TicketState
	(*CreateTicketRequest)(nil),  // 2: CreateTicketRequest
	(*CreateTicketResponse)(nil), // 3: CreateTicketResponse
}
var file_ticket_proto_depIdxs = []int32{
	0, // 0: CreateTicketRequest.channelType:type_name -> TicketChannelType
	1, // 1: CreateTicketRequest.state:type_name -> TicketState
	2, // 2: Ticket.CreateTicket:input_type -> CreateTicketRequest
	3, // 3: Ticket.CreateTicket:output_type -> CreateTicketResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ticket_proto_init() }
func file_ticket_proto_init() {
	if File_ticket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ticket_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateTicketRequest); i {
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
		file_ticket_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateTicketResponse); i {
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
	file_ticket_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ticket_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ticket_proto_goTypes,
		DependencyIndexes: file_ticket_proto_depIdxs,
		EnumInfos:         file_ticket_proto_enumTypes,
		MessageInfos:      file_ticket_proto_msgTypes,
	}.Build()
	File_ticket_proto = out.File
	file_ticket_proto_rawDesc = nil
	file_ticket_proto_goTypes = nil
	file_ticket_proto_depIdxs = nil
}
