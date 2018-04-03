// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event/pb/event.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	event/pb/event.proto

It has these top-level messages:
	Event
	EventQueryRequest
	EventQueryResponse
	EventSubscribeRequest
	EventSubscribeResponse
*/
package pb

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

type Event struct {
	Id          uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Contract    string `protobuf:"bytes,2,opt,name=contract" json:"contract,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Args        string `protobuf:"bytes,4,opt,name=args" json:"args,omitempty"`
	BlockNumber uint64 `protobuf:"varint,5,opt,name=block_number,json=blockNumber" json:"block_number,omitempty"`
	TxHash      string `protobuf:"bytes,6,opt,name=tx_hash,json=txHash" json:"tx_hash,omitempty"`
	BlockHash   string `protobuf:"bytes,7,opt,name=block_hash,json=blockHash" json:"block_hash,omitempty"`
	Removed     bool   `protobuf:"varint,8,opt,name=removed" json:"removed,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Event) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Event) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event) GetArgs() string {
	if m != nil {
		return m.Args
	}
	return ""
}

func (m *Event) GetBlockNumber() uint64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

func (m *Event) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

func (m *Event) GetBlockHash() string {
	if m != nil {
		return m.BlockHash
	}
	return ""
}

func (m *Event) GetRemoved() bool {
	if m != nil {
		return m.Removed
	}
	return false
}

type EventQueryRequest struct {
	TxHash string `protobuf:"bytes,1,opt,name=tx_hash,json=txHash" json:"tx_hash,omitempty"`
}

func (m *EventQueryRequest) Reset()                    { *m = EventQueryRequest{} }
func (m *EventQueryRequest) String() string            { return proto.CompactTextString(m) }
func (*EventQueryRequest) ProtoMessage()               {}
func (*EventQueryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *EventQueryRequest) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

type EventQueryResponse struct {
	TxHash string   `protobuf:"bytes,1,opt,name=tx_hash,json=txHash" json:"tx_hash,omitempty"`
	Events []*Event `protobuf:"bytes,2,rep,name=events" json:"events,omitempty"`
}

func (m *EventQueryResponse) Reset()                    { *m = EventQueryResponse{} }
func (m *EventQueryResponse) String() string            { return proto.CompactTextString(m) }
func (*EventQueryResponse) ProtoMessage()               {}
func (*EventQueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *EventQueryResponse) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

func (m *EventQueryResponse) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type EventSubscribeRequest struct {
	LastId uint64 `protobuf:"varint,1,opt,name=last_id,json=lastId" json:"last_id,omitempty"`
}

func (m *EventSubscribeRequest) Reset()                    { *m = EventSubscribeRequest{} }
func (m *EventSubscribeRequest) String() string            { return proto.CompactTextString(m) }
func (*EventSubscribeRequest) ProtoMessage()               {}
func (*EventSubscribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *EventSubscribeRequest) GetLastId() uint64 {
	if m != nil {
		return m.LastId
	}
	return 0
}

type EventSubscribeResponse struct {
	Event *Event `protobuf:"bytes,1,opt,name=event" json:"event,omitempty"`
}

func (m *EventSubscribeResponse) Reset()                    { *m = EventSubscribeResponse{} }
func (m *EventSubscribeResponse) String() string            { return proto.CompactTextString(m) }
func (*EventSubscribeResponse) ProtoMessage()               {}
func (*EventSubscribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *EventSubscribeResponse) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "pb.Event")
	proto.RegisterType((*EventQueryRequest)(nil), "pb.EventQueryRequest")
	proto.RegisterType((*EventQueryResponse)(nil), "pb.EventQueryResponse")
	proto.RegisterType((*EventSubscribeRequest)(nil), "pb.EventSubscribeRequest")
	proto.RegisterType((*EventSubscribeResponse)(nil), "pb.EventSubscribeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EventService service

type EventServiceClient interface {
	Query(ctx context.Context, in *EventQueryRequest, opts ...grpc.CallOption) (*EventQueryResponse, error)
	Subscribe(ctx context.Context, in *EventSubscribeRequest, opts ...grpc.CallOption) (EventService_SubscribeClient, error)
}

type eventServiceClient struct {
	cc *grpc.ClientConn
}

func NewEventServiceClient(cc *grpc.ClientConn) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) Query(ctx context.Context, in *EventQueryRequest, opts ...grpc.CallOption) (*EventQueryResponse, error) {
	out := new(EventQueryResponse)
	err := grpc.Invoke(ctx, "/pb.EventService/Query", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) Subscribe(ctx context.Context, in *EventSubscribeRequest, opts ...grpc.CallOption) (EventService_SubscribeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_EventService_serviceDesc.Streams[0], c.cc, "/pb.EventService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventService_SubscribeClient interface {
	Recv() (*EventSubscribeResponse, error)
	grpc.ClientStream
}

type eventServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *eventServiceSubscribeClient) Recv() (*EventSubscribeResponse, error) {
	m := new(EventSubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for EventService service

type EventServiceServer interface {
	Query(context.Context, *EventQueryRequest) (*EventQueryResponse, error)
	Subscribe(*EventSubscribeRequest, EventService_SubscribeServer) error
}

func RegisterEventServiceServer(s *grpc.Server, srv EventServiceServer) {
	s.RegisterService(&_EventService_serviceDesc, srv)
}

func _EventService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.EventService/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).Query(ctx, req.(*EventQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EventSubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventServiceServer).Subscribe(m, &eventServiceSubscribeServer{stream})
}

type EventService_SubscribeServer interface {
	Send(*EventSubscribeResponse) error
	grpc.ServerStream
}

type eventServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *eventServiceSubscribeServer) Send(m *EventSubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _EventService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _EventService_Query_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _EventService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "event/pb/event.proto",
}

func init() { proto.RegisterFile("event/pb/event.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xdf, 0x4e, 0x83, 0x30,
	0x14, 0xc6, 0x2d, 0x1b, 0x6c, 0x9c, 0x2d, 0x26, 0x36, 0x6e, 0x56, 0x12, 0x23, 0xe3, 0x8a, 0x0b,
	0xb3, 0x2d, 0xf3, 0x46, 0x1f, 0x40, 0xa3, 0x37, 0x46, 0xf1, 0x01, 0x16, 0x0a, 0x8d, 0x23, 0x6e,
	0x80, 0x6d, 0x59, 0xe6, 0x5b, 0xf8, 0x68, 0x3e, 0x92, 0xe1, 0xb0, 0x11, 0x24, 0x7a, 0xd7, 0xf3,
	0xfb, 0xce, 0x77, 0xfe, 0xa5, 0x70, 0x2a, 0xb6, 0x22, 0xd5, 0xb3, 0x9c, 0xcf, 0xf0, 0x31, 0xcd,
	0x65, 0xa6, 0x33, 0x6a, 0xe4, 0xdc, 0xfb, 0x26, 0x60, 0xde, 0x95, 0x8c, 0x1e, 0x83, 0x91, 0xc4,
	0x8c, 0xb8, 0xc4, 0xef, 0x06, 0x46, 0x12, 0x53, 0x07, 0xfa, 0x51, 0x96, 0x6a, 0x19, 0x46, 0x9a,
	0x19, 0x2e, 0xf1, 0xed, 0xa0, 0x8e, 0x29, 0x85, 0x6e, 0x1a, 0x6e, 0x04, 0xeb, 0x20, 0xc7, 0x77,
	0xc9, 0x42, 0xf9, 0xa6, 0x58, 0xb7, 0x62, 0xe5, 0x9b, 0x4e, 0x60, 0xc8, 0xd7, 0x59, 0xf4, 0xbe,
	0x4c, 0x8b, 0x0d, 0x17, 0x92, 0x99, 0x58, 0x7d, 0x80, 0xec, 0x09, 0x11, 0x3d, 0x83, 0x9e, 0xde,
	0x2d, 0x57, 0xa1, 0x5a, 0x31, 0x0b, 0x9d, 0x96, 0xde, 0x3d, 0x84, 0x6a, 0x45, 0x2f, 0x00, 0x2a,
	0x2f, 0x6a, 0x3d, 0xd4, 0x6c, 0x24, 0x28, 0x33, 0xe8, 0x49, 0xb1, 0xc9, 0xb6, 0x22, 0x66, 0x7d,
	0x97, 0xf8, 0xfd, 0xe0, 0x10, 0x7a, 0x57, 0x70, 0x82, 0x1b, 0xbd, 0x14, 0x42, 0x7e, 0x06, 0xe2,
	0xa3, 0x10, 0x4a, 0x37, 0xdb, 0x90, 0x66, 0x1b, 0xef, 0x19, 0x68, 0x33, 0x5b, 0xe5, 0x59, 0xaa,
	0xc4, 0xbf, 0xe9, 0x74, 0x02, 0x16, 0x9e, 0x50, 0x31, 0xc3, 0xed, 0xf8, 0x83, 0x85, 0x3d, 0xcd,
	0xf9, 0x14, 0x0b, 0x04, 0x7b, 0xc1, 0x9b, 0xc3, 0x08, 0xc1, 0x6b, 0xc1, 0x55, 0x24, 0x13, 0x2e,
	0x1a, 0x33, 0xac, 0x43, 0xa5, 0x97, 0xf5, 0x99, 0xad, 0x32, 0x7c, 0x8c, 0xbd, 0x5b, 0x18, 0xb7,
	0x1d, 0xfb, 0x39, 0x2e, 0xc1, 0xc4, 0xaa, 0x68, 0xf8, 0xd5, 0xad, 0xe2, 0x8b, 0x2f, 0x02, 0xc3,
	0xca, 0x2b, 0xe4, 0x36, 0x89, 0x04, 0xbd, 0x01, 0x13, 0x57, 0xa1, 0xa3, 0x3a, 0xb7, 0x79, 0x08,
	0x67, 0xdc, 0xc6, 0x55, 0x27, 0xef, 0x88, 0xde, 0x83, 0x5d, 0x0f, 0x40, 0xcf, 0xeb, 0xb4, 0xf6,
	0x1a, 0x8e, 0xf3, 0x97, 0x74, 0xa8, 0x32, 0x27, 0xdc, 0xc2, 0xdf, 0x75, 0xfd, 0x13, 0x00, 0x00,
	0xff, 0xff, 0x9e, 0x9c, 0xfe, 0xb2, 0x75, 0x02, 0x00, 0x00,
}
