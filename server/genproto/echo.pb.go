// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/showcase/v1alpha1/echo.proto

package genproto // import "github.com/googleapis/gapic-showcase/server/genproto"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "google.golang.org/genproto/googleapis/longrunning"
import status "google.golang.org/genproto/googleapis/rpc/status"

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

// The request message used for the Echo, Collect and Chat methods. If content
// is set in this message then the request will succeed. If a status is
type EchoRequest struct {
	// Types that are valid to be assigned to Response:
	//	*EchoRequest_Content
	//	*EchoRequest_Error
	Response             isEchoRequest_Response `protobuf_oneof:"response"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *EchoRequest) Reset()         { *m = EchoRequest{} }
func (m *EchoRequest) String() string { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()    {}
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_echo_a185ce30e13d26aa, []int{0}
}
func (m *EchoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoRequest.Unmarshal(m, b)
}
func (m *EchoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoRequest.Marshal(b, m, deterministic)
}
func (dst *EchoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoRequest.Merge(dst, src)
}
func (m *EchoRequest) XXX_Size() int {
	return xxx_messageInfo_EchoRequest.Size(m)
}
func (m *EchoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EchoRequest proto.InternalMessageInfo

type isEchoRequest_Response interface {
	isEchoRequest_Response()
}

type EchoRequest_Content struct {
	Content string `protobuf:"bytes,1,opt,name=content,proto3,oneof"`
}
type EchoRequest_Error struct {
	Error *status.Status `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*EchoRequest_Content) isEchoRequest_Response() {}
func (*EchoRequest_Error) isEchoRequest_Response()   {}

func (m *EchoRequest) GetResponse() isEchoRequest_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *EchoRequest) GetContent() string {
	if x, ok := m.GetResponse().(*EchoRequest_Content); ok {
		return x.Content
	}
	return ""
}

func (m *EchoRequest) GetError() *status.Status {
	if x, ok := m.GetResponse().(*EchoRequest_Error); ok {
		return x.Error
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*EchoRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _EchoRequest_OneofMarshaler, _EchoRequest_OneofUnmarshaler, _EchoRequest_OneofSizer, []interface{}{
		(*EchoRequest_Content)(nil),
		(*EchoRequest_Error)(nil),
	}
}

func _EchoRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*EchoRequest)
	// response
	switch x := m.Response.(type) {
	case *EchoRequest_Content:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Content)
	case *EchoRequest_Error:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("EchoRequest.Response has unexpected type %T", x)
	}
	return nil
}

func _EchoRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*EchoRequest)
	switch tag {
	case 1: // response.content
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Response = &EchoRequest_Content{x}
		return true, err
	case 2: // response.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(status.Status)
		err := b.DecodeMessage(msg)
		m.Response = &EchoRequest_Error{msg}
		return true, err
	default:
		return false, nil
	}
}

func _EchoRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*EchoRequest)
	// response
	switch x := m.Response.(type) {
	case *EchoRequest_Content:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Content)))
		n += len(x.Content)
	case *EchoRequest_Error:
		s := proto.Size(x.Error)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// The response message for the Echo methods.
type EchoResponse struct {
	// The content specified in the request.
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoResponse) Reset()         { *m = EchoResponse{} }
func (m *EchoResponse) String() string { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()    {}
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_echo_a185ce30e13d26aa, []int{1}
}
func (m *EchoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoResponse.Unmarshal(m, b)
}
func (m *EchoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoResponse.Marshal(b, m, deterministic)
}
func (dst *EchoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoResponse.Merge(dst, src)
}
func (m *EchoResponse) XXX_Size() int {
	return xxx_messageInfo_EchoResponse.Size(m)
}
func (m *EchoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EchoResponse proto.InternalMessageInfo

func (m *EchoResponse) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

// The request message for the Expand method.
type ExpandRequest struct {
	// The content that will be split into words and returned on the stream.
	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	// The error that is thrown after all words are sent on the stream.
	Error                *status.Status `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ExpandRequest) Reset()         { *m = ExpandRequest{} }
func (m *ExpandRequest) String() string { return proto.CompactTextString(m) }
func (*ExpandRequest) ProtoMessage()    {}
func (*ExpandRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_echo_a185ce30e13d26aa, []int{2}
}
func (m *ExpandRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExpandRequest.Unmarshal(m, b)
}
func (m *ExpandRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExpandRequest.Marshal(b, m, deterministic)
}
func (dst *ExpandRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExpandRequest.Merge(dst, src)
}
func (m *ExpandRequest) XXX_Size() int {
	return xxx_messageInfo_ExpandRequest.Size(m)
}
func (m *ExpandRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExpandRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExpandRequest proto.InternalMessageInfo

func (m *ExpandRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *ExpandRequest) GetError() *status.Status {
	if m != nil {
		return m.Error
	}
	return nil
}

// The request for Wait method.
type WaitRequest struct {
	// The amount of time to wait before returning a response.
	ResponseDelay *duration.Duration `protobuf:"bytes,1,opt,name=response_delay,json=responseDelay,proto3" json:"response_delay,omitempty"`
	// Types that are valid to be assigned to Response:
	//	*WaitRequest_Error
	//	*WaitRequest_Success
	Response             isWaitRequest_Response `protobuf_oneof:"response"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *WaitRequest) Reset()         { *m = WaitRequest{} }
func (m *WaitRequest) String() string { return proto.CompactTextString(m) }
func (*WaitRequest) ProtoMessage()    {}
func (*WaitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_echo_a185ce30e13d26aa, []int{3}
}
func (m *WaitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WaitRequest.Unmarshal(m, b)
}
func (m *WaitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WaitRequest.Marshal(b, m, deterministic)
}
func (dst *WaitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WaitRequest.Merge(dst, src)
}
func (m *WaitRequest) XXX_Size() int {
	return xxx_messageInfo_WaitRequest.Size(m)
}
func (m *WaitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WaitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WaitRequest proto.InternalMessageInfo

type isWaitRequest_Response interface {
	isWaitRequest_Response()
}

type WaitRequest_Error struct {
	Error *status.Status `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}
type WaitRequest_Success struct {
	Success *WaitResponse `protobuf:"bytes,3,opt,name=success,proto3,oneof"`
}

func (*WaitRequest_Error) isWaitRequest_Response()   {}
func (*WaitRequest_Success) isWaitRequest_Response() {}

func (m *WaitRequest) GetResponse() isWaitRequest_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *WaitRequest) GetResponseDelay() *duration.Duration {
	if m != nil {
		return m.ResponseDelay
	}
	return nil
}

func (m *WaitRequest) GetError() *status.Status {
	if x, ok := m.GetResponse().(*WaitRequest_Error); ok {
		return x.Error
	}
	return nil
}

func (m *WaitRequest) GetSuccess() *WaitResponse {
	if x, ok := m.GetResponse().(*WaitRequest_Success); ok {
		return x.Success
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*WaitRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _WaitRequest_OneofMarshaler, _WaitRequest_OneofUnmarshaler, _WaitRequest_OneofSizer, []interface{}{
		(*WaitRequest_Error)(nil),
		(*WaitRequest_Success)(nil),
	}
}

func _WaitRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*WaitRequest)
	// response
	switch x := m.Response.(type) {
	case *WaitRequest_Error:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case *WaitRequest_Success:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Success); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("WaitRequest.Response has unexpected type %T", x)
	}
	return nil
}

func _WaitRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*WaitRequest)
	switch tag {
	case 2: // response.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(status.Status)
		err := b.DecodeMessage(msg)
		m.Response = &WaitRequest_Error{msg}
		return true, err
	case 3: // response.success
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(WaitResponse)
		err := b.DecodeMessage(msg)
		m.Response = &WaitRequest_Success{msg}
		return true, err
	default:
		return false, nil
	}
}

func _WaitRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*WaitRequest)
	// response
	switch x := m.Response.(type) {
	case *WaitRequest_Error:
		s := proto.Size(x.Error)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *WaitRequest_Success:
		s := proto.Size(x.Success)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// The response for Wait method.
type WaitResponse struct {
	// This content can contain anything, the server will not depend on a value
	// here.
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WaitResponse) Reset()         { *m = WaitResponse{} }
func (m *WaitResponse) String() string { return proto.CompactTextString(m) }
func (*WaitResponse) ProtoMessage()    {}
func (*WaitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_echo_a185ce30e13d26aa, []int{4}
}
func (m *WaitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WaitResponse.Unmarshal(m, b)
}
func (m *WaitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WaitResponse.Marshal(b, m, deterministic)
}
func (dst *WaitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WaitResponse.Merge(dst, src)
}
func (m *WaitResponse) XXX_Size() int {
	return xxx_messageInfo_WaitResponse.Size(m)
}
func (m *WaitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WaitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WaitResponse proto.InternalMessageInfo

func (m *WaitResponse) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

// The request for the Pagination methods.
type PaginationRequest struct {
	// The maximum number that will be returned in the response.
	MaxResponse int32 `protobuf:"varint,1,opt,name=max_response,json=maxResponse,proto3" json:"max_response,omitempty"`
	// The amount of numbers to returned in the response.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The position of the page to be returned. This will be a stringified int
	// that will signifiy where to start the page from. Anything other than
	// a stringified integer within the range of 0 and the max_response will
	// cause an error to be thrown. This value is a string as opposed to a int32
	// to follow general google api practices.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// This field is used to show the client's ability to handle servers that
	// return a page that is larger the specified page size.
	PageSizeOverride     int32    `protobuf:"varint,4,opt,name=page_size_override,json=pageSizeOverride,proto3" json:"page_size_override,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaginationRequest) Reset()         { *m = PaginationRequest{} }
func (m *PaginationRequest) String() string { return proto.CompactTextString(m) }
func (*PaginationRequest) ProtoMessage()    {}
func (*PaginationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_echo_a185ce30e13d26aa, []int{5}
}
func (m *PaginationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaginationRequest.Unmarshal(m, b)
}
func (m *PaginationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaginationRequest.Marshal(b, m, deterministic)
}
func (dst *PaginationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaginationRequest.Merge(dst, src)
}
func (m *PaginationRequest) XXX_Size() int {
	return xxx_messageInfo_PaginationRequest.Size(m)
}
func (m *PaginationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PaginationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PaginationRequest proto.InternalMessageInfo

func (m *PaginationRequest) GetMaxResponse() int32 {
	if m != nil {
		return m.MaxResponse
	}
	return 0
}

func (m *PaginationRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *PaginationRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *PaginationRequest) GetPageSizeOverride() int32 {
	if m != nil {
		return m.PageSizeOverride
	}
	return 0
}

// The response for the Pagination method.
type PaginationResponse struct {
	// An increasing list of responses starting at the value specified by the
	// page token. If the page token is empty, then this list will start at 0.
	Responses []int32 `protobuf:"varint,1,rep,packed,name=responses,proto3" json:"responses,omitempty"`
	// The next integer stringified.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaginationResponse) Reset()         { *m = PaginationResponse{} }
func (m *PaginationResponse) String() string { return proto.CompactTextString(m) }
func (*PaginationResponse) ProtoMessage()    {}
func (*PaginationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_echo_a185ce30e13d26aa, []int{6}
}
func (m *PaginationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaginationResponse.Unmarshal(m, b)
}
func (m *PaginationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaginationResponse.Marshal(b, m, deterministic)
}
func (dst *PaginationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaginationResponse.Merge(dst, src)
}
func (m *PaginationResponse) XXX_Size() int {
	return xxx_messageInfo_PaginationResponse.Size(m)
}
func (m *PaginationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PaginationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PaginationResponse proto.InternalMessageInfo

func (m *PaginationResponse) GetResponses() []int32 {
	if m != nil {
		return m.Responses
	}
	return nil
}

func (m *PaginationResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func init() {
	proto.RegisterType((*EchoRequest)(nil), "google.showcase.v1alpha1.EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "google.showcase.v1alpha1.EchoResponse")
	proto.RegisterType((*ExpandRequest)(nil), "google.showcase.v1alpha1.ExpandRequest")
	proto.RegisterType((*WaitRequest)(nil), "google.showcase.v1alpha1.WaitRequest")
	proto.RegisterType((*WaitResponse)(nil), "google.showcase.v1alpha1.WaitResponse")
	proto.RegisterType((*PaginationRequest)(nil), "google.showcase.v1alpha1.PaginationRequest")
	proto.RegisterType((*PaginationResponse)(nil), "google.showcase.v1alpha1.PaginationResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EchoClient is the client API for Echo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoClient interface {
	// This method simply echos the request. This method is showcases unary rpcs.
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
	// This method split the given content into words and will pass each word back
	// through the stream. This method showcases server-side streaming rpcs.
	Expand(ctx context.Context, in *ExpandRequest, opts ...grpc.CallOption) (Echo_ExpandClient, error)
	// This method will collect the words given to it. When the stream is closed
	// by the client, this method will return the a concatenation of the strings
	// passed to it. This method showcases client-side streaming rpcs.
	Collect(ctx context.Context, opts ...grpc.CallOption) (Echo_CollectClient, error)
	// This method, upon receiving a request on the stream, the same content will
	// be passed  back on the stream. This method showcases bidirectional
	// streaming rpcs.
	Chat(ctx context.Context, opts ...grpc.CallOption) (Echo_ChatClient, error)
	// This method will wait the requested amount of and then return.
	// This method showcases how a client handles a request timing out.
	Wait(ctx context.Context, in *WaitRequest, opts ...grpc.CallOption) (*WaitResponse, error)
	// This method returns a page containing numbers specified in the request.
	// If the last number in the page is below the specified maximum, then the
	// response will include a page token indicating the next page.
	Pagination(ctx context.Context, in *PaginationRequest, opts ...grpc.CallOption) (*PaginationResponse, error)
}

type echoClient struct {
	cc *grpc.ClientConn
}

func NewEchoClient(cc *grpc.ClientConn) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/google.showcase.v1alpha1.Echo/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoClient) Expand(ctx context.Context, in *ExpandRequest, opts ...grpc.CallOption) (Echo_ExpandClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Echo_serviceDesc.Streams[0], "/google.showcase.v1alpha1.Echo/Expand", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoExpandClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Echo_ExpandClient interface {
	Recv() (*EchoResponse, error)
	grpc.ClientStream
}

type echoExpandClient struct {
	grpc.ClientStream
}

func (x *echoExpandClient) Recv() (*EchoResponse, error) {
	m := new(EchoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *echoClient) Collect(ctx context.Context, opts ...grpc.CallOption) (Echo_CollectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Echo_serviceDesc.Streams[1], "/google.showcase.v1alpha1.Echo/Collect", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoCollectClient{stream}
	return x, nil
}

type Echo_CollectClient interface {
	Send(*EchoRequest) error
	CloseAndRecv() (*EchoResponse, error)
	grpc.ClientStream
}

type echoCollectClient struct {
	grpc.ClientStream
}

func (x *echoCollectClient) Send(m *EchoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoCollectClient) CloseAndRecv() (*EchoResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EchoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *echoClient) Chat(ctx context.Context, opts ...grpc.CallOption) (Echo_ChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Echo_serviceDesc.Streams[2], "/google.showcase.v1alpha1.Echo/Chat", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoChatClient{stream}
	return x, nil
}

type Echo_ChatClient interface {
	Send(*EchoRequest) error
	Recv() (*EchoResponse, error)
	grpc.ClientStream
}

type echoChatClient struct {
	grpc.ClientStream
}

func (x *echoChatClient) Send(m *EchoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoChatClient) Recv() (*EchoResponse, error) {
	m := new(EchoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *echoClient) Wait(ctx context.Context, in *WaitRequest, opts ...grpc.CallOption) (*WaitResponse, error) {
	out := new(WaitResponse)
	err := c.cc.Invoke(ctx, "/google.showcase.v1alpha1.Echo/Wait", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoClient) Pagination(ctx context.Context, in *PaginationRequest, opts ...grpc.CallOption) (*PaginationResponse, error) {
	out := new(PaginationResponse)
	err := c.cc.Invoke(ctx, "/google.showcase.v1alpha1.Echo/Pagination", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoServer is the server API for Echo service.
type EchoServer interface {
	// This method simply echos the request. This method is showcases unary rpcs.
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
	// This method split the given content into words and will pass each word back
	// through the stream. This method showcases server-side streaming rpcs.
	Expand(*ExpandRequest, Echo_ExpandServer) error
	// This method will collect the words given to it. When the stream is closed
	// by the client, this method will return the a concatenation of the strings
	// passed to it. This method showcases client-side streaming rpcs.
	Collect(Echo_CollectServer) error
	// This method, upon receiving a request on the stream, the same content will
	// be passed  back on the stream. This method showcases bidirectional
	// streaming rpcs.
	Chat(Echo_ChatServer) error
	// This method will wait the requested amount of and then return.
	// This method showcases how a client handles a request timing out.
	Wait(context.Context, *WaitRequest) (*WaitResponse, error)
	// This method returns a page containing numbers specified in the request.
	// If the last number in the page is below the specified maximum, then the
	// response will include a page token indicating the next page.
	Pagination(context.Context, *PaginationRequest) (*PaginationResponse, error)
}

func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.showcase.v1alpha1.Echo/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Echo_Expand_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExpandRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EchoServer).Expand(m, &echoExpandServer{stream})
}

type Echo_ExpandServer interface {
	Send(*EchoResponse) error
	grpc.ServerStream
}

type echoExpandServer struct {
	grpc.ServerStream
}

func (x *echoExpandServer) Send(m *EchoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Echo_Collect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoServer).Collect(&echoCollectServer{stream})
}

type Echo_CollectServer interface {
	SendAndClose(*EchoResponse) error
	Recv() (*EchoRequest, error)
	grpc.ServerStream
}

type echoCollectServer struct {
	grpc.ServerStream
}

func (x *echoCollectServer) SendAndClose(m *EchoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoCollectServer) Recv() (*EchoRequest, error) {
	m := new(EchoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Echo_Chat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoServer).Chat(&echoChatServer{stream})
}

type Echo_ChatServer interface {
	Send(*EchoResponse) error
	Recv() (*EchoRequest, error)
	grpc.ServerStream
}

type echoChatServer struct {
	grpc.ServerStream
}

func (x *echoChatServer) Send(m *EchoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoChatServer) Recv() (*EchoRequest, error) {
	m := new(EchoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Echo_Wait_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WaitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).Wait(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.showcase.v1alpha1.Echo/Wait",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).Wait(ctx, req.(*WaitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Echo_Pagination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaginationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).Pagination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.showcase.v1alpha1.Echo/Pagination",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).Pagination(ctx, req.(*PaginationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.showcase.v1alpha1.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Echo_Echo_Handler,
		},
		{
			MethodName: "Wait",
			Handler:    _Echo_Wait_Handler,
		},
		{
			MethodName: "Pagination",
			Handler:    _Echo_Pagination_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Expand",
			Handler:       _Echo_Expand_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Collect",
			Handler:       _Echo_Collect_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Chat",
			Handler:       _Echo_Chat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "google/showcase/v1alpha1/echo.proto",
}

func init() {
	proto.RegisterFile("google/showcase/v1alpha1/echo.proto", fileDescriptor_echo_a185ce30e13d26aa)
}

var fileDescriptor_echo_a185ce30e13d26aa = []byte{
	// 725 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4f, 0x4f, 0x13, 0x41,
	0x14, 0x67, 0x80, 0x02, 0x9d, 0x02, 0xc2, 0x18, 0x71, 0x59, 0x90, 0xe0, 0x1a, 0x6b, 0x53, 0x71,
	0x97, 0x7f, 0xd1, 0xd8, 0x5b, 0x0b, 0x24, 0x78, 0x92, 0x14, 0x13, 0x23, 0x97, 0x66, 0xd8, 0x8e,
	0xbb, 0x1b, 0xb7, 0x33, 0xeb, 0xce, 0x2c, 0x54, 0x12, 0x63, 0xe2, 0x27, 0x30, 0xf1, 0x33, 0x78,
	0xf0, 0xea, 0x47, 0xf0, 0xe8, 0xc9, 0xc4, 0x1b, 0x27, 0x0f, 0x7e, 0x10, 0xb3, 0xb3, 0x33, 0xa5,
	0x54, 0x0a, 0x35, 0xe1, 0xb2, 0x9b, 0x7d, 0xf3, 0x7b, 0xbf, 0xdf, 0xef, 0xbd, 0x79, 0x6f, 0xe1,
	0x3d, 0x8f, 0x31, 0x2f, 0x24, 0x0e, 0xf7, 0xd9, 0xb1, 0x8b, 0x39, 0x71, 0x8e, 0xd6, 0x70, 0x18,
	0xf9, 0x78, 0xcd, 0x21, 0xae, 0xcf, 0xec, 0x28, 0x66, 0x82, 0x21, 0x23, 0x03, 0xd9, 0x1a, 0x64,
	0x6b, 0x90, 0xb9, 0xa8, 0xd2, 0x71, 0x14, 0x38, 0x98, 0x52, 0x26, 0xb0, 0x08, 0x18, 0xe5, 0x59,
	0x9e, 0xa9, 0xc9, 0x43, 0x46, 0xbd, 0x38, 0xa1, 0x34, 0xa0, 0x9e, 0xc3, 0x22, 0x12, 0x9f, 0x03,
	0x2d, 0x29, 0x90, 0xfc, 0x3a, 0x4c, 0x5e, 0x3b, 0xcd, 0x24, 0x03, 0xa8, 0xf3, 0xdb, 0xea, 0x3c,
	0x8e, 0x5c, 0x87, 0x0b, 0x2c, 0x12, 0x95, 0x68, 0x61, 0x58, 0xd8, 0x71, 0x7d, 0x56, 0x27, 0x6f,
	0x13, 0xc2, 0x05, 0x32, 0xe1, 0xb8, 0xcb, 0xa8, 0x20, 0x54, 0x18, 0x60, 0x19, 0x94, 0xf2, 0xbb,
	0x43, 0x75, 0x1d, 0x40, 0x65, 0x98, 0x23, 0x71, 0xcc, 0x62, 0x63, 0x78, 0x19, 0x94, 0x0a, 0xeb,
	0xc8, 0x56, 0x05, 0xc5, 0x91, 0x6b, 0xef, 0x4b, 0xce, 0xdd, 0xa1, 0x7a, 0x06, 0xa9, 0x41, 0x38,
	0x11, 0x13, 0x1e, 0x31, 0xca, 0x89, 0x55, 0x82, 0x93, 0x99, 0x44, 0xf6, 0x8d, 0x8c, 0x1e, 0x8d,
	0x8e, 0x82, 0xb5, 0x0f, 0xa7, 0x76, 0xda, 0x11, 0xa6, 0x4d, 0x6d, 0xa7, 0x2f, 0x14, 0x95, 0xae,
	0x34, 0xa3, 0xac, 0x58, 0x3f, 0x01, 0x2c, 0xbc, 0xc4, 0x81, 0xd0, 0x9c, 0x3b, 0x70, 0x5a, 0x5b,
	0x6b, 0x34, 0x49, 0x88, 0xdf, 0x49, 0xea, 0xc2, 0xfa, 0xbc, 0xa6, 0xd0, 0x3d, 0xb4, 0xb7, 0x55,
	0x0f, 0x6b, 0x23, 0xbf, 0xab, 0xa0, 0x3e, 0xa5, 0xb3, 0xb6, 0xd3, 0xa4, 0xff, 0xe9, 0x06, 0xaa,
	0xc1, 0x71, 0x9e, 0xb8, 0x2e, 0xe1, 0xdc, 0x18, 0x91, 0xe8, 0xa2, 0xdd, 0x6f, 0x18, 0xec, 0xcc,
	0x6a, 0xa6, 0x94, 0x76, 0x5f, 0x25, 0xf6, 0x76, 0xb4, 0x1b, 0x76, 0x49, 0x47, 0xbf, 0x00, 0x38,
	0xbb, 0x87, 0xbd, 0x80, 0xca, 0x42, 0x74, 0x0b, 0x8a, 0x70, 0xb2, 0x85, 0xdb, 0x0d, 0xcd, 0x27,
	0x93, 0x72, 0x59, 0x95, 0x85, 0x16, 0x6e, 0x77, 0x78, 0x17, 0x60, 0x3e, 0xc2, 0x1e, 0x69, 0xf0,
	0xe0, 0x84, 0xc8, 0x3a, 0x73, 0xf5, 0x89, 0x34, 0xb0, 0x1f, 0x9c, 0x10, 0x74, 0x07, 0x42, 0x79,
	0x28, 0xd8, 0x1b, 0x42, 0x65, 0x5d, 0xf9, 0xba, 0x84, 0xbf, 0x48, 0x03, 0x68, 0x05, 0xa2, 0x4e,
	0x6e, 0x83, 0x1d, 0x91, 0x38, 0x0e, 0x9a, 0xc4, 0x18, 0x95, 0x24, 0x33, 0x9a, 0xe4, 0xb9, 0x8a,
	0x5b, 0x07, 0x10, 0x75, 0xdb, 0x54, 0xfa, 0x8b, 0x30, 0xaf, 0x3d, 0x72, 0x03, 0x2c, 0x8f, 0x94,
	0x72, 0xf5, 0xb3, 0x00, 0x2a, 0xc2, 0x1b, 0x94, 0xb4, 0x45, 0xa3, 0xcb, 0xc5, 0xb0, 0x74, 0x31,
	0x95, 0x86, 0xf7, 0xb4, 0x93, 0xf5, 0xaf, 0x63, 0x70, 0x34, 0x1d, 0x40, 0x94, 0xa8, 0xf7, 0xfd,
	0xfe, 0xdd, 0xef, 0xda, 0x05, 0xb3, 0x78, 0x15, 0x4c, 0xdd, 0xc6, 0xd2, 0xc7, 0x5f, 0x7f, 0x3e,
	0x0f, 0x1b, 0xd6, 0xcd, 0xf3, 0x6b, 0x5f, 0x91, 0x0f, 0x50, 0x46, 0x9f, 0x00, 0x1c, 0xcb, 0xc6,
	0x1a, 0x3d, 0xb8, 0x84, 0xb2, 0x7b, 0xf0, 0x07, 0xd6, 0xde, 0x38, 0xad, 0xce, 0x74, 0x2e, 0x5f,
	0x4d, 0xa4, 0xb4, 0x63, 0x5a, 0xb7, 0x7a, 0xed, 0x48, 0x81, 0x0a, 0x28, 0xaf, 0x02, 0xf4, 0x1e,
	0x8e, 0x6f, 0xb1, 0x30, 0x24, 0xae, 0xb8, 0xee, 0x66, 0xdc, 0x95, 0xea, 0x0b, 0xd6, 0x5c, 0x8f,
	0xba, 0x9b, 0xc9, 0x55, 0x40, 0xb9, 0x04, 0xd0, 0x2b, 0x38, 0xba, 0xe5, 0xe3, 0xeb, 0xd6, 0x2e,
	0x81, 0x55, 0x90, 0xde, 0x71, 0xba, 0x1a, 0x97, 0x51, 0x77, 0xfd, 0x0c, 0xcc, 0x01, 0x17, 0xb1,
	0xef, 0x1d, 0x1f, 0xe3, 0x20, 0xad, 0x09, 0x7d, 0x03, 0x10, 0x9e, 0x0d, 0x30, 0x7a, 0xd8, 0x9f,
	0xf6, 0x9f, 0x6d, 0x34, 0x57, 0x06, 0x03, 0x2b, 0x27, 0xcf, 0x4e, 0xab, 0xd6, 0xf9, 0xf5, 0xbd,
	0x68, 0xd1, 0xa4, 0xdd, 0x25, 0x6b, 0xbe, 0xc7, 0x6e, 0xd4, 0xe1, 0xab, 0x80, 0xb2, 0x39, 0xfb,
	0xa3, 0x3a, 0x1d, 0x32, 0x17, 0x87, 0x3e, 0xe3, 0xa2, 0xf2, 0x64, 0xf3, 0xf1, 0xd3, 0xda, 0x87,
	0xef, 0x55, 0x1b, 0xad, 0xf8, 0x42, 0x44, 0xbc, 0xe2, 0x38, 0x5e, 0x20, 0xfc, 0xe4, 0xd0, 0x76,
	0x59, 0xcb, 0xc9, 0x2c, 0xe2, 0x28, 0xe0, 0x8e, 0x87, 0xa3, 0xc0, 0x7d, 0xa4, 0xcd, 0xc2, 0x39,
	0x97, 0xb5, 0x2e, 0xa8, 0xe0, 0x60, 0x73, 0x90, 0x6c, 0x87, 0x93, 0xf8, 0x88, 0xc4, 0x8e, 0x47,
	0x68, 0xf6, 0xdb, 0x1d, 0x93, 0xaf, 0x8d, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xdd, 0xaa, 0xc5,
	0x80, 0x53, 0x07, 0x00, 0x00,
}
