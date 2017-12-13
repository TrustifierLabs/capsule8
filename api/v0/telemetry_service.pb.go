// Code generated by protoc-gen-go. DO NOT EDIT.
// source: capsule8/api/v0/telemetry_service.proto

package capsule8_api_v0

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

// A request message to initiate the streaming of telemetry events
type GetEventsRequest struct {
	// The Subscription message defines which events should be
	// returned in the stream.
	Subscription *Subscription `protobuf:"bytes,1,opt,name=subscription" json:"subscription,omitempty"`
}

func (m *GetEventsRequest) Reset()                    { *m = GetEventsRequest{} }
func (m *GetEventsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetEventsRequest) ProtoMessage()               {}
func (*GetEventsRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *GetEventsRequest) GetSubscription() *Subscription {
	if m != nil {
		return m.Subscription
	}
	return nil
}

// A response message containing telemetry events
type GetEventsResponse struct {
	// Can publish one or more message(s) at a time
	Events []*TelemetryEvent `protobuf:"bytes,1,rep,name=events" json:"events,omitempty"`
}

func (m *GetEventsResponse) Reset()                    { *m = GetEventsResponse{} }
func (m *GetEventsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetEventsResponse) ProtoMessage()               {}
func (*GetEventsResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *GetEventsResponse) GetEvents() []*TelemetryEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

// A telemetry event received from a Sensor or Recorder.
type TelemetryEvent struct {
	// The time that the event was received by the backplane (in micros
	// since Unix epoch)
	PublishTimeMicros int64 `protobuf:"varint,1,opt,name=publish_time_micros,json=publishTimeMicros" json:"publish_time_micros,omitempty"`
	// The actual event observed by the Sensor. For historical
	// event subscriptions, this event may be sent from the
	// Recorder.
	Event *Event `protobuf:"bytes,2,opt,name=event" json:"event,omitempty"`
	// An opaque ack for the event. If present, this ack must be sent to
	// the PubsubService's Acknowledge method or else the TelemetryService
	// will re-transmit the event.
	Ack []byte `protobuf:"bytes,3,opt,name=ack,proto3" json:"ack,omitempty"`
}

func (m *TelemetryEvent) Reset()                    { *m = TelemetryEvent{} }
func (m *TelemetryEvent) String() string            { return proto.CompactTextString(m) }
func (*TelemetryEvent) ProtoMessage()               {}
func (*TelemetryEvent) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *TelemetryEvent) GetPublishTimeMicros() int64 {
	if m != nil {
		return m.PublishTimeMicros
	}
	return 0
}

func (m *TelemetryEvent) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *TelemetryEvent) GetAck() []byte {
	if m != nil {
		return m.Ack
	}
	return nil
}

func init() {
	proto.RegisterType((*GetEventsRequest)(nil), "capsule8.api.v0.GetEventsRequest")
	proto.RegisterType((*GetEventsResponse)(nil), "capsule8.api.v0.GetEventsResponse")
	proto.RegisterType((*TelemetryEvent)(nil), "capsule8.api.v0.TelemetryEvent")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TelemetryService service

type TelemetryServiceClient interface {
	// Opens a new stream of telemetry events
	GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (TelemetryService_GetEventsClient, error)
}

type telemetryServiceClient struct {
	cc *grpc.ClientConn
}

func NewTelemetryServiceClient(cc *grpc.ClientConn) TelemetryServiceClient {
	return &telemetryServiceClient{cc}
}

func (c *telemetryServiceClient) GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (TelemetryService_GetEventsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TelemetryService_serviceDesc.Streams[0], c.cc, "/capsule8.api.v0.TelemetryService/GetEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &telemetryServiceGetEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TelemetryService_GetEventsClient interface {
	Recv() (*GetEventsResponse, error)
	grpc.ClientStream
}

type telemetryServiceGetEventsClient struct {
	grpc.ClientStream
}

func (x *telemetryServiceGetEventsClient) Recv() (*GetEventsResponse, error) {
	m := new(GetEventsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for TelemetryService service

type TelemetryServiceServer interface {
	// Opens a new stream of telemetry events
	GetEvents(*GetEventsRequest, TelemetryService_GetEventsServer) error
}

func RegisterTelemetryServiceServer(s *grpc.Server, srv TelemetryServiceServer) {
	s.RegisterService(&_TelemetryService_serviceDesc, srv)
}

func _TelemetryService_GetEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TelemetryServiceServer).GetEvents(m, &telemetryServiceGetEventsServer{stream})
}

type TelemetryService_GetEventsServer interface {
	Send(*GetEventsResponse) error
	grpc.ServerStream
}

type telemetryServiceGetEventsServer struct {
	grpc.ServerStream
}

func (x *telemetryServiceGetEventsServer) Send(m *GetEventsResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _TelemetryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "capsule8.api.v0.TelemetryService",
	HandlerType: (*TelemetryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetEvents",
			Handler:       _TelemetryService_GetEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "capsule8/api/v0/telemetry_service.proto",
}

func init() { proto.RegisterFile("capsule8/api/v0/telemetry_service.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xcd, 0x4b, 0xc3, 0x30,
	0x14, 0x27, 0x16, 0x07, 0x66, 0x43, 0xb7, 0x08, 0x52, 0x2a, 0x62, 0xed, 0xc5, 0x1e, 0x24, 0x2d,
	0xf5, 0xa0, 0x57, 0x0f, 0xe2, 0x45, 0x2f, 0x5d, 0x3d, 0x97, 0xb6, 0x3c, 0x58, 0xb0, 0x1f, 0xb1,
	0x2f, 0x2d, 0x78, 0xf3, 0x4f, 0x97, 0x25, 0xdb, 0xe8, 0x5a, 0xf0, 0x16, 0xf2, 0xfb, 0x7a, 0xbf,
	0xf7, 0xe8, 0x7d, 0x91, 0x49, 0xec, 0x4a, 0x78, 0x0e, 0x32, 0x29, 0x82, 0x3e, 0x0c, 0x14, 0x94,
	0x50, 0x81, 0x6a, 0x7f, 0x52, 0x84, 0xb6, 0x17, 0x05, 0x70, 0xd9, 0x36, 0xaa, 0x61, 0x17, 0x7b,
	0x22, 0xcf, 0xa4, 0xe0, 0x7d, 0xe8, 0x78, 0x63, 0x25, 0x76, 0x39, 0x16, 0xad, 0x90, 0x4a, 0x34,
	0xb5, 0x11, 0x39, 0xd7, 0x63, 0x0e, 0xf4, 0x50, 0x2b, 0x03, 0x7a, 0x9f, 0x74, 0xf9, 0x06, 0xea,
	0x75, 0xfb, 0x83, 0x31, 0x7c, 0x77, 0x80, 0x8a, 0xbd, 0xd0, 0xc5, 0xd0, 0xc6, 0x26, 0x2e, 0xf1,
	0xe7, 0xd1, 0x0d, 0x1f, 0x85, 0xf3, 0xf5, 0x80, 0x14, 0x1f, 0x49, 0xbc, 0x77, 0xba, 0x1a, 0xd8,
	0xa2, 0x6c, 0x6a, 0x04, 0xf6, 0x44, 0x67, 0x3a, 0x1a, 0x6d, 0xe2, 0x5a, 0xfe, 0x3c, 0xba, 0x9d,
	0x38, 0x26, 0xfb, 0xde, 0x5a, 0x19, 0xef, 0xe8, 0xde, 0x2f, 0xa1, 0xe7, 0xc7, 0x10, 0xe3, 0xf4,
	0x52, 0x76, 0x79, 0x29, 0x70, 0x93, 0x2a, 0x51, 0x41, 0x5a, 0x89, 0xa2, 0x6d, 0x50, 0x8f, 0x6a,
	0xc5, 0xab, 0x1d, 0x94, 0x88, 0x0a, 0x3e, 0x34, 0xc0, 0x1e, 0xe8, 0xa9, 0x36, 0xb3, 0x4f, 0x74,
	0x99, 0xab, 0x49, 0xb4, 0x49, 0x34, 0x24, 0xb6, 0xa4, 0x56, 0x56, 0x7c, 0xd9, 0x96, 0x4b, 0xfc,
	0x45, 0xbc, 0x7d, 0x46, 0x1b, 0xba, 0x3c, 0x4c, 0xb0, 0x36, 0x37, 0x61, 0x09, 0x3d, 0x3b, 0x94,
	0x64, 0x77, 0x13, 0xc7, 0xf1, 0x5e, 0x1d, 0xef, 0x3f, 0x8a, 0xd9, 0x51, 0x48, 0xf2, 0x99, 0x3e,
	0xcc, 0xe3, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x94, 0x22, 0x9b, 0xdc, 0x15, 0x02, 0x00, 0x00,
}
