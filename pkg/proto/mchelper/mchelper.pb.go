// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mchelper/mchelper.proto

package mchelper

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Mod struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Mod) Reset()         { *m = Mod{} }
func (m *Mod) String() string { return proto.CompactTextString(m) }
func (*Mod) ProtoMessage()    {}
func (*Mod) Descriptor() ([]byte, []int) {
	return fileDescriptor_mchelper_4d5396ef3423d7a5, []int{0}
}
func (m *Mod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mod.Unmarshal(m, b)
}
func (m *Mod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mod.Marshal(b, m, deterministic)
}
func (dst *Mod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mod.Merge(dst, src)
}
func (m *Mod) XXX_Size() int {
	return xxx_messageInfo_Mod.Size(m)
}
func (m *Mod) XXX_DiscardUnknown() {
	xxx_messageInfo_Mod.DiscardUnknown(m)
}

var xxx_messageInfo_Mod proto.InternalMessageInfo

func (m *Mod) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Mod) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type GenericOutput struct {
	Output               string   `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenericOutput) Reset()         { *m = GenericOutput{} }
func (m *GenericOutput) String() string { return proto.CompactTextString(m) }
func (*GenericOutput) ProtoMessage()    {}
func (*GenericOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_mchelper_4d5396ef3423d7a5, []int{1}
}
func (m *GenericOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenericOutput.Unmarshal(m, b)
}
func (m *GenericOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenericOutput.Marshal(b, m, deterministic)
}
func (dst *GenericOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericOutput.Merge(dst, src)
}
func (m *GenericOutput) XXX_Size() int {
	return xxx_messageInfo_GenericOutput.Size(m)
}
func (m *GenericOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericOutput.DiscardUnknown(m)
}

var xxx_messageInfo_GenericOutput proto.InternalMessageInfo

func (m *GenericOutput) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_mchelper_4d5396ef3423d7a5, []int{2}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type ListModsResponse struct {
	Mods                 []*Mod   `protobuf:"bytes,2,rep,name=mods,proto3" json:"mods,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListModsResponse) Reset()         { *m = ListModsResponse{} }
func (m *ListModsResponse) String() string { return proto.CompactTextString(m) }
func (*ListModsResponse) ProtoMessage()    {}
func (*ListModsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_mchelper_4d5396ef3423d7a5, []int{3}
}
func (m *ListModsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListModsResponse.Unmarshal(m, b)
}
func (m *ListModsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListModsResponse.Marshal(b, m, deterministic)
}
func (dst *ListModsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListModsResponse.Merge(dst, src)
}
func (m *ListModsResponse) XXX_Size() int {
	return xxx_messageInfo_ListModsResponse.Size(m)
}
func (m *ListModsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListModsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListModsResponse proto.InternalMessageInfo

func (m *ListModsResponse) GetMods() []*Mod {
	if m != nil {
		return m.Mods
	}
	return nil
}

type SendCommandRequest struct {
	Command              string   `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendCommandRequest) Reset()         { *m = SendCommandRequest{} }
func (m *SendCommandRequest) String() string { return proto.CompactTextString(m) }
func (*SendCommandRequest) ProtoMessage()    {}
func (*SendCommandRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_mchelper_4d5396ef3423d7a5, []int{4}
}
func (m *SendCommandRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendCommandRequest.Unmarshal(m, b)
}
func (m *SendCommandRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendCommandRequest.Marshal(b, m, deterministic)
}
func (dst *SendCommandRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendCommandRequest.Merge(dst, src)
}
func (m *SendCommandRequest) XXX_Size() int {
	return xxx_messageInfo_SendCommandRequest.Size(m)
}
func (m *SendCommandRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendCommandRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendCommandRequest proto.InternalMessageInfo

func (m *SendCommandRequest) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

type SendCommandResponse struct {
	Output               []*GenericOutput `protobuf:"bytes,1,rep,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SendCommandResponse) Reset()         { *m = SendCommandResponse{} }
func (m *SendCommandResponse) String() string { return proto.CompactTextString(m) }
func (*SendCommandResponse) ProtoMessage()    {}
func (*SendCommandResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_mchelper_4d5396ef3423d7a5, []int{5}
}
func (m *SendCommandResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendCommandResponse.Unmarshal(m, b)
}
func (m *SendCommandResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendCommandResponse.Marshal(b, m, deterministic)
}
func (dst *SendCommandResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendCommandResponse.Merge(dst, src)
}
func (m *SendCommandResponse) XXX_Size() int {
	return xxx_messageInfo_SendCommandResponse.Size(m)
}
func (m *SendCommandResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendCommandResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendCommandResponse proto.InternalMessageInfo

func (m *SendCommandResponse) GetOutput() []*GenericOutput {
	if m != nil {
		return m.Output
	}
	return nil
}

func init() {
	proto.RegisterType((*Mod)(nil), "mchelper.mchelper.Mod")
	proto.RegisterType((*GenericOutput)(nil), "mchelper.mchelper.GenericOutput")
	proto.RegisterType((*Empty)(nil), "mchelper.mchelper.Empty")
	proto.RegisterType((*ListModsResponse)(nil), "mchelper.mchelper.ListModsResponse")
	proto.RegisterType((*SendCommandRequest)(nil), "mchelper.mchelper.SendCommandRequest")
	proto.RegisterType((*SendCommandResponse)(nil), "mchelper.mchelper.SendCommandResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OperationsServicesClient is the client API for OperationsServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OperationsServicesClient interface {
	ListMods(ctx context.Context, in *Empty, opts ...grpc.CallOption) (OperationsServices_ListModsClient, error)
	SendCommand(ctx context.Context, in *SendCommandRequest, opts ...grpc.CallOption) (OperationsServices_SendCommandClient, error)
}

type operationsServicesClient struct {
	cc *grpc.ClientConn
}

func NewOperationsServicesClient(cc *grpc.ClientConn) OperationsServicesClient {
	return &operationsServicesClient{cc}
}

func (c *operationsServicesClient) ListMods(ctx context.Context, in *Empty, opts ...grpc.CallOption) (OperationsServices_ListModsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_OperationsServices_serviceDesc.Streams[0], "/mchelper.mchelper.OperationsServices/ListMods", opts...)
	if err != nil {
		return nil, err
	}
	x := &operationsServicesListModsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OperationsServices_ListModsClient interface {
	Recv() (*ListModsResponse, error)
	grpc.ClientStream
}

type operationsServicesListModsClient struct {
	grpc.ClientStream
}

func (x *operationsServicesListModsClient) Recv() (*ListModsResponse, error) {
	m := new(ListModsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *operationsServicesClient) SendCommand(ctx context.Context, in *SendCommandRequest, opts ...grpc.CallOption) (OperationsServices_SendCommandClient, error) {
	stream, err := c.cc.NewStream(ctx, &_OperationsServices_serviceDesc.Streams[1], "/mchelper.mchelper.OperationsServices/SendCommand", opts...)
	if err != nil {
		return nil, err
	}
	x := &operationsServicesSendCommandClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OperationsServices_SendCommandClient interface {
	Recv() (*SendCommandResponse, error)
	grpc.ClientStream
}

type operationsServicesSendCommandClient struct {
	grpc.ClientStream
}

func (x *operationsServicesSendCommandClient) Recv() (*SendCommandResponse, error) {
	m := new(SendCommandResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OperationsServicesServer is the server API for OperationsServices service.
type OperationsServicesServer interface {
	ListMods(*Empty, OperationsServices_ListModsServer) error
	SendCommand(*SendCommandRequest, OperationsServices_SendCommandServer) error
}

func RegisterOperationsServicesServer(s *grpc.Server, srv OperationsServicesServer) {
	s.RegisterService(&_OperationsServices_serviceDesc, srv)
}

func _OperationsServices_ListMods_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OperationsServicesServer).ListMods(m, &operationsServicesListModsServer{stream})
}

type OperationsServices_ListModsServer interface {
	Send(*ListModsResponse) error
	grpc.ServerStream
}

type operationsServicesListModsServer struct {
	grpc.ServerStream
}

func (x *operationsServicesListModsServer) Send(m *ListModsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _OperationsServices_SendCommand_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SendCommandRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OperationsServicesServer).SendCommand(m, &operationsServicesSendCommandServer{stream})
}

type OperationsServices_SendCommandServer interface {
	Send(*SendCommandResponse) error
	grpc.ServerStream
}

type operationsServicesSendCommandServer struct {
	grpc.ServerStream
}

func (x *operationsServicesSendCommandServer) Send(m *SendCommandResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _OperationsServices_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mchelper.mchelper.OperationsServices",
	HandlerType: (*OperationsServicesServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListMods",
			Handler:       _OperationsServices_ListMods_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendCommand",
			Handler:       _OperationsServices_SendCommand_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "mchelper/mchelper.proto",
}

func init() { proto.RegisterFile("mchelper/mchelper.proto", fileDescriptor_mchelper_4d5396ef3423d7a5) }

var fileDescriptor_mchelper_4d5396ef3423d7a5 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x35, 0x6d, 0x6d, 0xeb, 0x14, 0x41, 0x47, 0xa8, 0xa1, 0xa7, 0xb2, 0xa2, 0x16, 0x0f, 0x51,
	0xda, 0x8b, 0x27, 0x0f, 0x8a, 0x78, 0x31, 0x14, 0xd2, 0x9b, 0x27, 0x63, 0x76, 0xc0, 0x80, 0xfb,
	0xe1, 0xee, 0xa6, 0xe0, 0x0f, 0xf4, 0x7f, 0x49, 0xd7, 0xac, 0xb6, 0x36, 0xe0, 0xed, 0xbd, 0xc9,
	0xcb, 0xbc, 0xf7, 0x86, 0x85, 0x63, 0x51, 0xbc, 0xd2, 0x9b, 0x26, 0x73, 0x19, 0x40, 0xa2, 0x8d,
	0x72, 0x0a, 0x0f, 0x7f, 0x78, 0x00, 0x6c, 0x06, 0xed, 0x54, 0x71, 0x44, 0xe8, 0xc8, 0x5c, 0x50,
	0x1c, 0x8d, 0xa3, 0xc9, 0x5e, 0xe6, 0x31, 0xc6, 0xd0, 0x5b, 0x92, 0xb1, 0xa5, 0x92, 0x71, 0xcb,
	0x8f, 0x03, 0x65, 0xe7, 0xb0, 0xff, 0x40, 0x92, 0x4c, 0x59, 0xcc, 0x2b, 0xa7, 0x2b, 0x87, 0x43,
	0xe8, 0x2a, 0x8f, 0xea, 0x05, 0x35, 0x63, 0x3d, 0xd8, 0xbd, 0x17, 0xda, 0x7d, 0xb0, 0x1b, 0x38,
	0x78, 0x2c, 0xad, 0x4b, 0x15, 0xb7, 0x19, 0x59, 0xad, 0xa4, 0x25, 0xbc, 0x80, 0x8e, 0x50, 0xdc,
	0xc6, 0xad, 0x71, 0x7b, 0x32, 0x98, 0x0e, 0x93, 0xad, 0x70, 0x49, 0xaa, 0x78, 0xe6, 0x35, 0x2c,
	0x01, 0x5c, 0x90, 0xe4, 0x77, 0x4a, 0x88, 0x5c, 0xf2, 0x8c, 0xde, 0x2b, 0xb2, 0x6e, 0x95, 0xb0,
	0xf8, 0x9e, 0xd4, 0xbe, 0x81, 0xb2, 0x39, 0x1c, 0x6d, 0xe8, 0x6b, 0xcb, 0xeb, 0xb5, 0x9c, 0x2b,
	0xd3, 0x71, 0x83, 0xe9, 0x46, 0xb3, 0xd0, 0x64, 0xfa, 0x19, 0x01, 0xce, 0x35, 0x99, 0xdc, 0x95,
	0x4a, 0xda, 0x05, 0x99, 0x65, 0x59, 0x90, 0xc5, 0x14, 0xfa, 0xa1, 0x17, 0xc6, 0x0d, 0xcb, 0x7c,
	0xfb, 0xd1, 0x49, 0xc3, 0x97, 0xbf, 0xe7, 0x60, 0x3b, 0x57, 0x11, 0x3e, 0xc3, 0x60, 0x2d, 0x36,
	0x9e, 0x36, 0xfc, 0xb7, 0x7d, 0x86, 0xd1, 0xd9, 0x7f, 0xb2, 0x5f, 0x87, 0x5b, 0x78, 0xea, 0x07,
	0xcd, 0x4b, 0xd7, 0xbf, 0x8a, 0xd9, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x65, 0xf6, 0xce, 0x39,
	0x30, 0x02, 0x00, 0x00,
}
