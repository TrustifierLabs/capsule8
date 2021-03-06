// Code generated by protoc-gen-go. DO NOT EDIT.
// source: capsule8/api/v0/telemetry_service.proto

package capsule8_api_v0

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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
	Events []*ReceivedTelemetryEvent `protobuf:"bytes,1,rep,name=events" json:"events,omitempty"`
}

func (m *GetEventsResponse) Reset()                    { *m = GetEventsResponse{} }
func (m *GetEventsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetEventsResponse) ProtoMessage()               {}
func (*GetEventsResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *GetEventsResponse) GetEvents() []*ReceivedTelemetryEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

// A telemetry event received from a Sensor or Recorder.
type ReceivedTelemetryEvent struct {
	// The time that the event was received by the backplane (in micros
	// since Unix epoch)
	PublishTimeMicros int64 `protobuf:"varint,1,opt,name=publish_time_micros,json=publishTimeMicros" json:"publish_time_micros,omitempty"`
	// The actual event observed by the Sensor. For historical
	// event subscriptions, this event may be sent from the
	// Recorder.
	Event *TelemetryEvent `protobuf:"bytes,2,opt,name=event" json:"event,omitempty"`
	// An opaque ack for the event. If present, this ack must be sent to
	// the PubsubService's Acknowledge method or else the TelemetryService
	// will re-transmit the event.
	Ack []byte `protobuf:"bytes,3,opt,name=ack,proto3" json:"ack,omitempty"`
}

func (m *ReceivedTelemetryEvent) Reset()                    { *m = ReceivedTelemetryEvent{} }
func (m *ReceivedTelemetryEvent) String() string            { return proto.CompactTextString(m) }
func (*ReceivedTelemetryEvent) ProtoMessage()               {}
func (*ReceivedTelemetryEvent) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *ReceivedTelemetryEvent) GetPublishTimeMicros() int64 {
	if m != nil {
		return m.PublishTimeMicros
	}
	return 0
}

func (m *ReceivedTelemetryEvent) GetEvent() *TelemetryEvent {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *ReceivedTelemetryEvent) GetAck() []byte {
	if m != nil {
		return m.Ack
	}
	return nil
}

func init() {
	proto.RegisterType((*GetEventsRequest)(nil), "capsule8.api.v0.GetEventsRequest")
	proto.RegisterType((*GetEventsResponse)(nil), "capsule8.api.v0.GetEventsResponse")
	proto.RegisterType((*ReceivedTelemetryEvent)(nil), "capsule8.api.v0.ReceivedTelemetryEvent")
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
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xd9, 0x06, 0x0b, 0x6e, 0x0b, 0x6d, 0x57, 0x94, 0x50, 0x14, 0xeb, 0x82, 0xb4, 0x78,
	0x48, 0x4b, 0x45, 0x10, 0x2f, 0xe2, 0x41, 0x3c, 0x79, 0x49, 0xeb, 0xb9, 0xa4, 0x71, 0xa8, 0x4b,
	0x93, 0xec, 0x9a, 0xd9, 0x04, 0x3c, 0x09, 0xbe, 0x81, 0xf8, 0x68, 0xbe, 0x82, 0x0f, 0x22, 0xd9,
	0x6d, 0x4b, 0x4d, 0xd4, 0xdb, 0xc2, 0x7c, 0xf3, 0xed, 0x3f, 0x33, 0xb4, 0x1f, 0x06, 0x0a, 0xb3,
	0x08, 0x2e, 0x87, 0x81, 0x12, 0xc3, 0x7c, 0x34, 0xd4, 0x10, 0x41, 0x0c, 0x3a, 0x7d, 0x99, 0x21,
	0xa4, 0xb9, 0x08, 0xc1, 0x53, 0xa9, 0xd4, 0x92, 0xb5, 0xd6, 0xa0, 0x17, 0x28, 0xe1, 0xe5, 0xa3,
	0x2e, 0x2f, 0x77, 0x62, 0x36, 0xc7, 0x30, 0x15, 0x4a, 0x0b, 0x99, 0xd8, 0xa6, 0xee, 0xe9, 0xdf,
	0x76, 0xc8, 0x21, 0xd1, 0x2b, 0xec, 0x70, 0x21, 0xe5, 0x22, 0x02, 0x03, 0x05, 0x49, 0x22, 0x75,
	0x50, 0x38, 0xd0, 0x56, 0xf9, 0x03, 0x6d, 0xdf, 0x81, 0xbe, 0x2d, 0x78, 0xf4, 0xe1, 0x39, 0x03,
	0xd4, 0xec, 0x86, 0x36, 0xb7, 0xbf, 0x73, 0x49, 0x8f, 0x0c, 0x1a, 0xe3, 0x23, 0xaf, 0x14, 0xd2,
	0x9b, 0x6c, 0x41, 0xfe, 0x8f, 0x16, 0x3e, 0xa5, 0x9d, 0x2d, 0x2d, 0x2a, 0x99, 0x20, 0xb0, 0x6b,
	0x5a, 0x37, 0xc1, 0xd0, 0x25, 0x3d, 0x67, 0xd0, 0x18, 0xf7, 0x2b, 0x46, 0x1f, 0x42, 0x10, 0x39,
	0x3c, 0x4e, 0xd7, 0x93, 0x18, 0x83, 0xbf, 0x6a, 0xe3, 0xef, 0x84, 0x1e, 0xfc, 0x8e, 0x30, 0x8f,
	0xee, 0xa9, 0x6c, 0x1e, 0x09, 0x7c, 0x9a, 0x69, 0x11, 0xc3, 0x2c, 0x16, 0x61, 0x2a, 0xd1, 0x44,
	0x77, 0xfc, 0xce, 0xaa, 0x34, 0x15, 0x31, 0xdc, 0x9b, 0x02, 0xbb, 0xa0, 0x3b, 0x46, 0xea, 0xd6,
	0xcc, 0x70, 0xc7, 0x95, 0x28, 0xa5, 0x08, 0x96, 0x66, 0x6d, 0xea, 0x04, 0xe1, 0xd2, 0x75, 0x7a,
	0x64, 0xd0, 0xf4, 0x8b, 0xe7, 0xf8, 0x95, 0xb6, 0x37, 0xe8, 0xc4, 0x1e, 0x95, 0x2d, 0xe9, 0xee,
	0x66, 0x7a, 0x76, 0x52, 0x51, 0x97, 0x17, 0xde, 0xe5, 0xff, 0x21, 0x76, 0x79, 0x7c, 0xff, 0xed,
	0xf3, 0xeb, 0xa3, 0xd6, 0xe2, 0xb4, 0xb8, 0xb4, 0xdd, 0xc7, 0x15, 0x39, 0x1b, 0x91, 0x79, 0xdd,
	0x1c, 0xf2, 0xfc, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xf7, 0x4e, 0x1e, 0x6d, 0x02, 0x00, 0x00,
}
