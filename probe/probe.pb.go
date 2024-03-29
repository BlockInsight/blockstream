// Code generated by protoc-gen-go. DO NOT EDIT.
// source: probe.proto

package probe

import (
	context "context"
	fmt "fmt"
	types "github.com/blockinsight/blockstream/types"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type JobRequest struct {
	Blockchain           types.Blockchain `protobuf:"varint,1,opt,name=blockchain,proto3,enum=types.Blockchain" json:"blockchain,omitempty"`
	ChainId              string           `protobuf:"bytes,2,opt,name=chainId,proto3" json:"chainId,omitempty"`
	RemoteUrl            string           `protobuf:"bytes,3,opt,name=remoteUrl,proto3" json:"remoteUrl,omitempty"`
	Offset               uint64           `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	Count                int64            `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`
	Backoff              uint64           `protobuf:"varint,6,opt,name=backoff,proto3" json:"backoff,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *JobRequest) Reset()         { *m = JobRequest{} }
func (m *JobRequest) String() string { return proto.CompactTextString(m) }
func (*JobRequest) ProtoMessage()    {}
func (*JobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8cc8551cf1a5c4f, []int{0}
}

func (m *JobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobRequest.Unmarshal(m, b)
}
func (m *JobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobRequest.Marshal(b, m, deterministic)
}
func (m *JobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobRequest.Merge(m, src)
}
func (m *JobRequest) XXX_Size() int {
	return xxx_messageInfo_JobRequest.Size(m)
}
func (m *JobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JobRequest proto.InternalMessageInfo

func (m *JobRequest) GetBlockchain() types.Blockchain {
	if m != nil {
		return m.Blockchain
	}
	return types.Blockchain_BTC
}

func (m *JobRequest) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *JobRequest) GetRemoteUrl() string {
	if m != nil {
		return m.RemoteUrl
	}
	return ""
}

func (m *JobRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *JobRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *JobRequest) GetBackoff() uint64 {
	if m != nil {
		return m.Backoff
	}
	return 0
}

type JobResponse struct {
	Block                []byte   `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobResponse) Reset()         { *m = JobResponse{} }
func (m *JobResponse) String() string { return proto.CompactTextString(m) }
func (*JobResponse) ProtoMessage()    {}
func (*JobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8cc8551cf1a5c4f, []int{1}
}

func (m *JobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobResponse.Unmarshal(m, b)
}
func (m *JobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobResponse.Marshal(b, m, deterministic)
}
func (m *JobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobResponse.Merge(m, src)
}
func (m *JobResponse) XXX_Size() int {
	return xxx_messageInfo_JobResponse.Size(m)
}
func (m *JobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JobResponse proto.InternalMessageInfo

func (m *JobResponse) GetBlock() []byte {
	if m != nil {
		return m.Block
	}
	return nil
}

func init() {
	proto.RegisterType((*JobRequest)(nil), "probe.JobRequest")
	proto.RegisterType((*JobResponse)(nil), "probe.JobResponse")
}

func init() { proto.RegisterFile("probe.proto", fileDescriptor_f8cc8551cf1a5c4f) }

var fileDescriptor_f8cc8551cf1a5c4f = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x31, 0x4f, 0xfb, 0x30,
	0x10, 0xc5, 0xe5, 0x7f, 0x9b, 0xfc, 0x95, 0x0b, 0x42, 0xaa, 0x85, 0x90, 0x55, 0x31, 0x44, 0x65,
	0x09, 0x42, 0x4a, 0xa1, 0xb0, 0xb2, 0x94, 0x89, 0x4e, 0xc8, 0x12, 0x0b, 0x5b, 0x6c, 0x2e, 0x4d,
	0xd4, 0x26, 0x17, 0x6c, 0x67, 0xe0, 0xd3, 0xf1, 0xd5, 0x50, 0x6c, 0x4a, 0xbb, 0xf9, 0xf7, 0xee,
	0xde, 0xf9, 0xee, 0x41, 0xda, 0x1b, 0x52, 0x58, 0xf4, 0x86, 0x1c, 0xf1, 0xc8, 0xc3, 0x3c, 0x75,
	0x5f, 0x3d, 0xda, 0xa0, 0x2d, 0xbe, 0x19, 0xc0, 0x86, 0x94, 0xc4, 0xcf, 0x01, 0xad, 0xe3, 0xf7,
	0x00, 0x6a, 0x4f, 0x7a, 0xa7, 0xeb, 0xb2, 0xe9, 0x04, 0xcb, 0x58, 0x7e, 0xbe, 0x9a, 0x15, 0xc1,
	0xb0, 0xfe, 0x2b, 0xc8, 0x93, 0x26, 0x2e, 0xe0, 0xbf, 0x7f, 0xbc, 0x7c, 0x88, 0x7f, 0x19, 0xcb,
	0x13, 0x79, 0x40, 0x7e, 0x05, 0x89, 0xc1, 0x96, 0x1c, 0xbe, 0x99, 0xbd, 0x98, 0xf8, 0xda, 0x51,
	0xe0, 0x97, 0x10, 0x53, 0x55, 0x59, 0x74, 0x62, 0x9a, 0xb1, 0x7c, 0x2a, 0x7f, 0x89, 0x5f, 0x40,
	0xa4, 0x69, 0xe8, 0x9c, 0x88, 0x32, 0x96, 0x4f, 0x64, 0x80, 0xf1, 0x17, 0x55, 0xea, 0x1d, 0x55,
	0x95, 0x88, 0x7d, 0xfb, 0x01, 0x17, 0xd7, 0x90, 0xfa, 0x03, 0x6c, 0x4f, 0x9d, 0xc5, 0xd1, 0xee,
	0x97, 0xf3, 0xcb, 0x9f, 0xc9, 0x00, 0xab, 0x27, 0x88, 0x5e, 0xc7, 0xe3, 0xf9, 0x23, 0x24, 0xcf,
	0x06, 0x4b, 0x87, 0x1b, 0x52, 0x7c, 0x56, 0x84, 0x78, 0x8e, 0x01, 0xcc, 0xf9, 0xa9, 0x14, 0x46,
	0xde, 0xb1, 0xf5, 0xed, 0xfb, 0xcd, 0xb6, 0x71, 0xf5, 0xa0, 0x0a, 0x4d, 0xed, 0xd2, 0x8f, 0x6c,
	0x3a, 0xdb, 0x6c, 0x6b, 0x17, 0xc0, 0x3a, 0x83, 0x65, 0xbb, 0xf4, 0x56, 0x15, 0xfb, 0x64, 0x1f,
	0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x98, 0x92, 0x0e, 0xa5, 0x7c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProbeClient is the client API for Probe service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProbeClient interface {
	CreateJob(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (Probe_CreateJobClient, error)
}

type probeClient struct {
	cc *grpc.ClientConn
}

func NewProbeClient(cc *grpc.ClientConn) ProbeClient {
	return &probeClient{cc}
}

func (c *probeClient) CreateJob(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (Probe_CreateJobClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Probe_serviceDesc.Streams[0], "/probe.Probe/CreateJob", opts...)
	if err != nil {
		return nil, err
	}
	x := &probeCreateJobClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Probe_CreateJobClient interface {
	Recv() (*JobResponse, error)
	grpc.ClientStream
}

type probeCreateJobClient struct {
	grpc.ClientStream
}

func (x *probeCreateJobClient) Recv() (*JobResponse, error) {
	m := new(JobResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProbeServer is the server API for Probe service.
type ProbeServer interface {
	CreateJob(*JobRequest, Probe_CreateJobServer) error
}

func RegisterProbeServer(s *grpc.Server, srv ProbeServer) {
	s.RegisterService(&_Probe_serviceDesc, srv)
}

func _Probe_CreateJob_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(JobRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProbeServer).CreateJob(m, &probeCreateJobServer{stream})
}

type Probe_CreateJobServer interface {
	Send(*JobResponse) error
	grpc.ServerStream
}

type probeCreateJobServer struct {
	grpc.ServerStream
}

func (x *probeCreateJobServer) Send(m *JobResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Probe_serviceDesc = grpc.ServiceDesc{
	ServiceName: "probe.Probe",
	HandlerType: (*ProbeServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateJob",
			Handler:       _Probe_CreateJob_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "probe.proto",
}
