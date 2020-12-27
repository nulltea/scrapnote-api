// Code generated by protoc-gen-go. DO NOT EDIT.
// source: record.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Record struct {
	UniqueID             string               `protobuf:"bytes,1,opt,name=uniqueID,proto3" json:"uniqueID,omitempty"`
	Content              string               `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	SourceURL            string               `protobuf:"bytes,3,opt,name=sourceURL,proto3" json:"sourceURL,omitempty"`
	MarkerURL            string               `protobuf:"bytes,4,opt,name=markerURL,proto3" json:"markerURL,omitempty"`
	AddedAt              *timestamp.Timestamp `protobuf:"bytes,5,opt,name=addedAt,proto3" json:"addedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf94fd919e302a1d, []int{0}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetUniqueID() string {
	if m != nil {
		return m.UniqueID
	}
	return ""
}

func (m *Record) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Record) GetSourceURL() string {
	if m != nil {
		return m.SourceURL
	}
	return ""
}

func (m *Record) GetMarkerURL() string {
	if m != nil {
		return m.MarkerURL
	}
	return ""
}

func (m *Record) GetAddedAt() *timestamp.Timestamp {
	if m != nil {
		return m.AddedAt
	}
	return nil
}

type RecordFilter struct {
	RecordID             []string `protobuf:"bytes,1,rep,name=recordID,proto3" json:"recordID,omitempty"`
	TopicID              string   `protobuf:"bytes,2,opt,name=topicID,proto3" json:"topicID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecordFilter) Reset()         { *m = RecordFilter{} }
func (m *RecordFilter) String() string { return proto.CompactTextString(m) }
func (*RecordFilter) ProtoMessage()    {}
func (*RecordFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf94fd919e302a1d, []int{1}
}

func (m *RecordFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordFilter.Unmarshal(m, b)
}
func (m *RecordFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordFilter.Marshal(b, m, deterministic)
}
func (m *RecordFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordFilter.Merge(m, src)
}
func (m *RecordFilter) XXX_Size() int {
	return xxx_messageInfo_RecordFilter.Size(m)
}
func (m *RecordFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordFilter.DiscardUnknown(m)
}

var xxx_messageInfo_RecordFilter proto.InternalMessageInfo

func (m *RecordFilter) GetRecordID() []string {
	if m != nil {
		return m.RecordID
	}
	return nil
}

func (m *RecordFilter) GetTopicID() string {
	if m != nil {
		return m.TopicID
	}
	return ""
}

type RecordResponse struct {
	Records              []*Record `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
	Count                int64     `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RecordResponse) Reset()         { *m = RecordResponse{} }
func (m *RecordResponse) String() string { return proto.CompactTextString(m) }
func (*RecordResponse) ProtoMessage()    {}
func (*RecordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf94fd919e302a1d, []int{2}
}

func (m *RecordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordResponse.Unmarshal(m, b)
}
func (m *RecordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordResponse.Marshal(b, m, deterministic)
}
func (m *RecordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordResponse.Merge(m, src)
}
func (m *RecordResponse) XXX_Size() int {
	return xxx_messageInfo_RecordResponse.Size(m)
}
func (m *RecordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RecordResponse proto.InternalMessageInfo

func (m *RecordResponse) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

func (m *RecordResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*Record)(nil), "proto.Record")
	proto.RegisterType((*RecordFilter)(nil), "proto.RecordFilter")
	proto.RegisterType((*RecordResponse)(nil), "proto.RecordResponse")
}

func init() { proto.RegisterFile("record.proto", fileDescriptor_bf94fd919e302a1d) }

var fileDescriptor_bf94fd919e302a1d = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x51, 0x4b, 0xfb, 0x30,
	0x14, 0xc5, 0xff, 0xfd, 0xd7, 0x6e, 0x2e, 0xdb, 0x7c, 0x88, 0x0a, 0xa5, 0x08, 0x8e, 0xbe, 0xb8,
	0x97, 0xb5, 0xd0, 0xcd, 0x0f, 0xa0, 0x0e, 0x65, 0x20, 0x08, 0x55, 0x5f, 0x7c, 0xeb, 0xd2, 0x6b,
	0x17, 0x5c, 0x9b, 0x9a, 0xa6, 0x82, 0xf8, 0xa5, 0xfc, 0x88, 0x92, 0xdc, 0x66, 0xb2, 0x47, 0x9f,
	0xd2, 0x93, 0x73, 0xee, 0xe1, 0xd7, 0x1b, 0x32, 0x92, 0xc0, 0x84, 0xcc, 0xa3, 0x5a, 0x0a, 0x25,
	0xa8, 0x67, 0x8e, 0xe0, 0xbc, 0x10, 0xa2, 0xd8, 0x42, 0x6c, 0xd4, 0xba, 0x7d, 0x8d, 0x15, 0x2f,
	0xa1, 0x51, 0x59, 0x59, 0x63, 0x2e, 0xfc, 0x76, 0x48, 0x2f, 0x35, 0x83, 0x34, 0x20, 0x87, 0x6d,
	0xc5, 0xdf, 0x5b, 0x58, 0x2d, 0x7d, 0x67, 0xe2, 0x4c, 0x07, 0xe9, 0x4e, 0x53, 0x9f, 0xf4, 0x99,
	0xa8, 0x14, 0x54, 0xca, 0xff, 0x6f, 0x2c, 0x2b, 0xe9, 0x19, 0x19, 0x34, 0xa2, 0x95, 0x0c, 0x9e,
	0xd3, 0x7b, 0xdf, 0x35, 0xde, 0xef, 0x85, 0x76, 0xcb, 0x4c, 0xbe, 0x81, 0xd4, 0xee, 0x01, 0xba,
	0xbb, 0x0b, 0xba, 0x20, 0xfd, 0x2c, 0xcf, 0x21, 0xbf, 0x52, 0xbe, 0x37, 0x71, 0xa6, 0xc3, 0x24,
	0x88, 0x90, 0x37, 0xb2, 0xbc, 0xd1, 0x93, 0xe5, 0x4d, 0x6d, 0x34, 0x5c, 0x92, 0x11, 0x12, 0xdf,
	0xf2, 0xad, 0x02, 0xa9, 0xb9, 0xf1, 0xd7, 0x0d, 0xb7, 0xab, 0xb9, 0xad, 0xd6, 0xdc, 0x4a, 0xd4,
	0x9c, 0xad, 0x96, 0x96, 0xbb, 0x93, 0xe1, 0x03, 0x39, 0xc2, 0x96, 0x14, 0x9a, 0x5a, 0x54, 0x0d,
	0xd0, 0x0b, 0xd2, 0xc7, 0xb9, 0xc6, 0xd4, 0x0c, 0x93, 0x31, 0x62, 0x44, 0x5d, 0xce, 0xba, 0xf4,
	0x84, 0x78, 0x4c, 0xb4, 0xdd, 0x2a, 0xdc, 0x14, 0x45, 0xf2, 0x45, 0xc6, 0x18, 0x7c, 0x04, 0xf9,
	0xc1, 0x19, 0xd0, 0x39, 0x71, 0xef, 0x40, 0xd1, 0xe3, 0xbd, 0x16, 0x64, 0x0e, 0x4e, 0xf7, 0xab,
	0x3b, 0x84, 0xf0, 0x1f, 0xbd, 0x24, 0xde, 0x8d, 0xae, 0xfb, 0xdb, 0xd8, 0xf5, 0xe2, 0x25, 0x29,
	0xb8, 0xda, 0xb4, 0xeb, 0x88, 0x89, 0x52, 0x3f, 0xb2, 0x50, 0x9b, 0xd9, 0x67, 0xdc, 0x30, 0x99,
	0xd5, 0x95, 0x50, 0x30, 0xcb, 0x6a, 0x1e, 0x23, 0x7f, 0xac, 0x3f, 0x71, 0xbf, 0x3d, 0x73, 0xcc,
	0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x24, 0x6f, 0x00, 0x3b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RecordServiceClient is the client API for RecordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RecordServiceClient interface {
	Get(ctx context.Context, in *RecordFilter, opts ...grpc.CallOption) (*RecordResponse, error)
	Count(ctx context.Context, in *RecordFilter, opts ...grpc.CallOption) (*RecordResponse, error)
}

type recordServiceClient struct {
	cc *grpc.ClientConn
}

func NewRecordServiceClient(cc *grpc.ClientConn) RecordServiceClient {
	return &recordServiceClient{cc}
}

func (c *recordServiceClient) Get(ctx context.Context, in *RecordFilter, opts ...grpc.CallOption) (*RecordResponse, error) {
	out := new(RecordResponse)
	err := c.cc.Invoke(ctx, "/proto.RecordService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordServiceClient) Count(ctx context.Context, in *RecordFilter, opts ...grpc.CallOption) (*RecordResponse, error) {
	out := new(RecordResponse)
	err := c.cc.Invoke(ctx, "/proto.RecordService/Count", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecordServiceServer is the server API for RecordService service.
type RecordServiceServer interface {
	Get(context.Context, *RecordFilter) (*RecordResponse, error)
	Count(context.Context, *RecordFilter) (*RecordResponse, error)
}

// UnimplementedRecordServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRecordServiceServer struct {
}

func (*UnimplementedRecordServiceServer) Get(ctx context.Context, req *RecordFilter) (*RecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedRecordServiceServer) Count(ctx context.Context, req *RecordFilter) (*RecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Count not implemented")
}

func RegisterRecordServiceServer(s *grpc.Server, srv RecordServiceServer) {
	s.RegisterService(&_RecordService_serviceDesc, srv)
}

func _RecordService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RecordService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServiceServer).Get(ctx, req.(*RecordFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecordService_Count_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServiceServer).Count(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RecordService/Count",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServiceServer).Count(ctx, req.(*RecordFilter))
	}
	return interceptor(ctx, in, info, handler)
}

var _RecordService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.RecordService",
	HandlerType: (*RecordServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _RecordService_Get_Handler,
		},
		{
			MethodName: "Count",
			Handler:    _RecordService_Count_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "record.proto",
}