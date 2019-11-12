// Code generated by protoc-gen-go. DO NOT EDIT.
// source: jobs.proto

package jobs

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

type CreateJobRequest struct {
	Job                  *types.Job `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateJobRequest) Reset()         { *m = CreateJobRequest{} }
func (m *CreateJobRequest) String() string { return proto.CompactTextString(m) }
func (*CreateJobRequest) ProtoMessage()    {}
func (*CreateJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_45644d4410f878a7, []int{0}
}

func (m *CreateJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateJobRequest.Unmarshal(m, b)
}
func (m *CreateJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateJobRequest.Marshal(b, m, deterministic)
}
func (m *CreateJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateJobRequest.Merge(m, src)
}
func (m *CreateJobRequest) XXX_Size() int {
	return xxx_messageInfo_CreateJobRequest.Size(m)
}
func (m *CreateJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateJobRequest proto.InternalMessageInfo

func (m *CreateJobRequest) GetJob() *types.Job {
	if m != nil {
		return m.Job
	}
	return nil
}

type CreateJobResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateJobResponse) Reset()         { *m = CreateJobResponse{} }
func (m *CreateJobResponse) String() string { return proto.CompactTextString(m) }
func (*CreateJobResponse) ProtoMessage()    {}
func (*CreateJobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_45644d4410f878a7, []int{1}
}

func (m *CreateJobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateJobResponse.Unmarshal(m, b)
}
func (m *CreateJobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateJobResponse.Marshal(b, m, deterministic)
}
func (m *CreateJobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateJobResponse.Merge(m, src)
}
func (m *CreateJobResponse) XXX_Size() int {
	return xxx_messageInfo_CreateJobResponse.Size(m)
}
func (m *CreateJobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateJobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateJobResponse proto.InternalMessageInfo

func (m *CreateJobResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CancelJobRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelJobRequest) Reset()         { *m = CancelJobRequest{} }
func (m *CancelJobRequest) String() string { return proto.CompactTextString(m) }
func (*CancelJobRequest) ProtoMessage()    {}
func (*CancelJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_45644d4410f878a7, []int{2}
}

func (m *CancelJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelJobRequest.Unmarshal(m, b)
}
func (m *CancelJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelJobRequest.Marshal(b, m, deterministic)
}
func (m *CancelJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelJobRequest.Merge(m, src)
}
func (m *CancelJobRequest) XXX_Size() int {
	return xxx_messageInfo_CancelJobRequest.Size(m)
}
func (m *CancelJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CancelJobRequest proto.InternalMessageInfo

func (m *CancelJobRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CancelJobResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelJobResponse) Reset()         { *m = CancelJobResponse{} }
func (m *CancelJobResponse) String() string { return proto.CompactTextString(m) }
func (*CancelJobResponse) ProtoMessage()    {}
func (*CancelJobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_45644d4410f878a7, []int{3}
}

func (m *CancelJobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelJobResponse.Unmarshal(m, b)
}
func (m *CancelJobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelJobResponse.Marshal(b, m, deterministic)
}
func (m *CancelJobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelJobResponse.Merge(m, src)
}
func (m *CancelJobResponse) XXX_Size() int {
	return xxx_messageInfo_CancelJobResponse.Size(m)
}
func (m *CancelJobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelJobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CancelJobResponse proto.InternalMessageInfo

type WorkerRequest struct {
	// Types that are valid to be assigned to Command:
	//	*WorkerRequest_WorkerStarted
	//	*WorkerRequest_TaskCompleted
	Command              isWorkerRequest_Command `protobuf_oneof:"command"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *WorkerRequest) Reset()         { *m = WorkerRequest{} }
func (m *WorkerRequest) String() string { return proto.CompactTextString(m) }
func (*WorkerRequest) ProtoMessage()    {}
func (*WorkerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_45644d4410f878a7, []int{4}
}

func (m *WorkerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkerRequest.Unmarshal(m, b)
}
func (m *WorkerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkerRequest.Marshal(b, m, deterministic)
}
func (m *WorkerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkerRequest.Merge(m, src)
}
func (m *WorkerRequest) XXX_Size() int {
	return xxx_messageInfo_WorkerRequest.Size(m)
}
func (m *WorkerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WorkerRequest proto.InternalMessageInfo

type isWorkerRequest_Command interface {
	isWorkerRequest_Command()
}

type WorkerRequest_WorkerStarted struct {
	WorkerStarted *WorkerStartedCommand `protobuf:"bytes,1,opt,name=workerStarted,proto3,oneof"`
}

type WorkerRequest_TaskCompleted struct {
	TaskCompleted *TaskCompletedCommand `protobuf:"bytes,2,opt,name=taskCompleted,proto3,oneof"`
}

func (*WorkerRequest_WorkerStarted) isWorkerRequest_Command() {}

func (*WorkerRequest_TaskCompleted) isWorkerRequest_Command() {}

func (m *WorkerRequest) GetCommand() isWorkerRequest_Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *WorkerRequest) GetWorkerStarted() *WorkerStartedCommand {
	if x, ok := m.GetCommand().(*WorkerRequest_WorkerStarted); ok {
		return x.WorkerStarted
	}
	return nil
}

func (m *WorkerRequest) GetTaskCompleted() *TaskCompletedCommand {
	if x, ok := m.GetCommand().(*WorkerRequest_TaskCompleted); ok {
		return x.TaskCompleted
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*WorkerRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _WorkerRequest_OneofMarshaler, _WorkerRequest_OneofUnmarshaler, _WorkerRequest_OneofSizer, []interface{}{
		(*WorkerRequest_WorkerStarted)(nil),
		(*WorkerRequest_TaskCompleted)(nil),
	}
}

func _WorkerRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*WorkerRequest)
	// command
	switch x := m.Command.(type) {
	case *WorkerRequest_WorkerStarted:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.WorkerStarted); err != nil {
			return err
		}
	case *WorkerRequest_TaskCompleted:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TaskCompleted); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("WorkerRequest.Command has unexpected type %T", x)
	}
	return nil
}

func _WorkerRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*WorkerRequest)
	switch tag {
	case 1: // command.workerStarted
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(WorkerStartedCommand)
		err := b.DecodeMessage(msg)
		m.Command = &WorkerRequest_WorkerStarted{msg}
		return true, err
	case 2: // command.taskCompleted
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TaskCompletedCommand)
		err := b.DecodeMessage(msg)
		m.Command = &WorkerRequest_TaskCompleted{msg}
		return true, err
	default:
		return false, nil
	}
}

func _WorkerRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*WorkerRequest)
	// command
	switch x := m.Command.(type) {
	case *WorkerRequest_WorkerStarted:
		s := proto.Size(x.WorkerStarted)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *WorkerRequest_TaskCompleted:
		s := proto.Size(x.TaskCompleted)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type WorkerResponse struct {
	// Types that are valid to be assigned to Command:
	//	*WorkerResponse_CreateTask
	Command              isWorkerResponse_Command `protobuf_oneof:"command"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *WorkerResponse) Reset()         { *m = WorkerResponse{} }
func (m *WorkerResponse) String() string { return proto.CompactTextString(m) }
func (*WorkerResponse) ProtoMessage()    {}
func (*WorkerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_45644d4410f878a7, []int{5}
}

func (m *WorkerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkerResponse.Unmarshal(m, b)
}
func (m *WorkerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkerResponse.Marshal(b, m, deterministic)
}
func (m *WorkerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkerResponse.Merge(m, src)
}
func (m *WorkerResponse) XXX_Size() int {
	return xxx_messageInfo_WorkerResponse.Size(m)
}
func (m *WorkerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WorkerResponse proto.InternalMessageInfo

type isWorkerResponse_Command interface {
	isWorkerResponse_Command()
}

type WorkerResponse_CreateTask struct {
	CreateTask *CreateTaskCommand `protobuf:"bytes,1,opt,name=createTask,proto3,oneof"`
}

func (*WorkerResponse_CreateTask) isWorkerResponse_Command() {}

func (m *WorkerResponse) GetCommand() isWorkerResponse_Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *WorkerResponse) GetCreateTask() *CreateTaskCommand {
	if x, ok := m.GetCommand().(*WorkerResponse_CreateTask); ok {
		return x.CreateTask
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*WorkerResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _WorkerResponse_OneofMarshaler, _WorkerResponse_OneofUnmarshaler, _WorkerResponse_OneofSizer, []interface{}{
		(*WorkerResponse_CreateTask)(nil),
	}
}

func _WorkerResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*WorkerResponse)
	// command
	switch x := m.Command.(type) {
	case *WorkerResponse_CreateTask:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CreateTask); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("WorkerResponse.Command has unexpected type %T", x)
	}
	return nil
}

func _WorkerResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*WorkerResponse)
	switch tag {
	case 1: // command.createTask
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CreateTaskCommand)
		err := b.DecodeMessage(msg)
		m.Command = &WorkerResponse_CreateTask{msg}
		return true, err
	default:
		return false, nil
	}
}

func _WorkerResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*WorkerResponse)
	// command
	switch x := m.Command.(type) {
	case *WorkerResponse_CreateTask:
		s := proto.Size(x.CreateTask)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type CreateTaskCommand struct {
	PrevTask             string      `protobuf:"bytes,2,opt,name=prevTask,proto3" json:"prevTask,omitempty"`
	PrevTaskWorker       string      `protobuf:"bytes,3,opt,name=prevTaskWorker,proto3" json:"prevTaskWorker,omitempty"`
	Task                 *types.Task `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreateTaskCommand) Reset()         { *m = CreateTaskCommand{} }
func (m *CreateTaskCommand) String() string { return proto.CompactTextString(m) }
func (*CreateTaskCommand) ProtoMessage()    {}
func (*CreateTaskCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_45644d4410f878a7, []int{6}
}

func (m *CreateTaskCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskCommand.Unmarshal(m, b)
}
func (m *CreateTaskCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskCommand.Marshal(b, m, deterministic)
}
func (m *CreateTaskCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskCommand.Merge(m, src)
}
func (m *CreateTaskCommand) XXX_Size() int {
	return xxx_messageInfo_CreateTaskCommand.Size(m)
}
func (m *CreateTaskCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskCommand.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskCommand proto.InternalMessageInfo

func (m *CreateTaskCommand) GetPrevTask() string {
	if m != nil {
		return m.PrevTask
	}
	return ""
}

func (m *CreateTaskCommand) GetPrevTaskWorker() string {
	if m != nil {
		return m.PrevTaskWorker
	}
	return ""
}

func (m *CreateTaskCommand) GetTask() *types.Task {
	if m != nil {
		return m.Task
	}
	return nil
}

type TaskCompletedCommand struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskCompletedCommand) Reset()         { *m = TaskCompletedCommand{} }
func (m *TaskCompletedCommand) String() string { return proto.CompactTextString(m) }
func (*TaskCompletedCommand) ProtoMessage()    {}
func (*TaskCompletedCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_45644d4410f878a7, []int{7}
}

func (m *TaskCompletedCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskCompletedCommand.Unmarshal(m, b)
}
func (m *TaskCompletedCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskCompletedCommand.Marshal(b, m, deterministic)
}
func (m *TaskCompletedCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskCompletedCommand.Merge(m, src)
}
func (m *TaskCompletedCommand) XXX_Size() int {
	return xxx_messageInfo_TaskCompletedCommand.Size(m)
}
func (m *TaskCompletedCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskCompletedCommand.DiscardUnknown(m)
}

var xxx_messageInfo_TaskCompletedCommand proto.InternalMessageInfo

func (m *TaskCompletedCommand) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type WorkerStartedCommand struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MaxTasks             string   `protobuf:"bytes,2,opt,name=maxTasks,proto3" json:"maxTasks,omitempty"`
	RunningTasks         string   `protobuf:"bytes,3,opt,name=runningTasks,proto3" json:"runningTasks,omitempty"`
	Outputs              []string `protobuf:"bytes,4,rep,name=outputs,proto3" json:"outputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkerStartedCommand) Reset()         { *m = WorkerStartedCommand{} }
func (m *WorkerStartedCommand) String() string { return proto.CompactTextString(m) }
func (*WorkerStartedCommand) ProtoMessage()    {}
func (*WorkerStartedCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_45644d4410f878a7, []int{8}
}

func (m *WorkerStartedCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkerStartedCommand.Unmarshal(m, b)
}
func (m *WorkerStartedCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkerStartedCommand.Marshal(b, m, deterministic)
}
func (m *WorkerStartedCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkerStartedCommand.Merge(m, src)
}
func (m *WorkerStartedCommand) XXX_Size() int {
	return xxx_messageInfo_WorkerStartedCommand.Size(m)
}
func (m *WorkerStartedCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkerStartedCommand.DiscardUnknown(m)
}

var xxx_messageInfo_WorkerStartedCommand proto.InternalMessageInfo

func (m *WorkerStartedCommand) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *WorkerStartedCommand) GetMaxTasks() string {
	if m != nil {
		return m.MaxTasks
	}
	return ""
}

func (m *WorkerStartedCommand) GetRunningTasks() string {
	if m != nil {
		return m.RunningTasks
	}
	return ""
}

func (m *WorkerStartedCommand) GetOutputs() []string {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateJobRequest)(nil), "jobs.CreateJobRequest")
	proto.RegisterType((*CreateJobResponse)(nil), "jobs.CreateJobResponse")
	proto.RegisterType((*CancelJobRequest)(nil), "jobs.CancelJobRequest")
	proto.RegisterType((*CancelJobResponse)(nil), "jobs.CancelJobResponse")
	proto.RegisterType((*WorkerRequest)(nil), "jobs.WorkerRequest")
	proto.RegisterType((*WorkerResponse)(nil), "jobs.WorkerResponse")
	proto.RegisterType((*CreateTaskCommand)(nil), "jobs.CreateTaskCommand")
	proto.RegisterType((*TaskCompletedCommand)(nil), "jobs.TaskCompletedCommand")
	proto.RegisterType((*WorkerStartedCommand)(nil), "jobs.WorkerStartedCommand")
}

func init() { proto.RegisterFile("jobs.proto", fileDescriptor_45644d4410f878a7) }

var fileDescriptor_45644d4410f878a7 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0xcd, 0x8e, 0xd3, 0x30,
	0x18, 0x24, 0x69, 0x45, 0xc9, 0x57, 0xb6, 0xda, 0xf5, 0x56, 0x6c, 0x14, 0x21, 0xb1, 0x32, 0xd2,
	0xaa, 0xa7, 0x6a, 0xb5, 0x5c, 0x40, 0xe2, 0xd4, 0x70, 0x00, 0x4e, 0x28, 0x20, 0x7a, 0xce, 0x8f,
	0x45, 0xd3, 0x36, 0x76, 0xb0, 0x1d, 0x28, 0x37, 0x9e, 0x85, 0x47, 0xe1, 0xc9, 0x56, 0xb6, 0xe3,
	0xd4, 0x49, 0x7b, 0x8b, 0xe7, 0x9b, 0x19, 0x8d, 0x3f, 0x4f, 0x00, 0xb6, 0x2c, 0x13, 0xcb, 0x9a,
	0x33, 0xc9, 0xd0, 0x58, 0x7d, 0x47, 0x53, 0xf9, 0xa7, 0x26, 0x2d, 0x84, 0xef, 0xe1, 0x32, 0xe6,
	0x24, 0x95, 0xe4, 0x33, 0xcb, 0x12, 0xf2, 0xb3, 0x21, 0x42, 0xa2, 0x97, 0x30, 0xda, 0xb2, 0x2c,
	0xf4, 0x6e, 0xbd, 0xc5, 0xf4, 0x01, 0x96, 0x86, 0xae, 0xe6, 0x0a, 0xc6, 0xaf, 0xe1, 0xca, 0x51,
	0x88, 0x9a, 0x51, 0x41, 0xd0, 0x0c, 0xfc, 0xb2, 0xd0, 0x8a, 0x20, 0xf1, 0xcb, 0x02, 0x63, 0xb8,
	0x8c, 0x53, 0x9a, 0x93, 0xbd, 0x63, 0x3b, 0xe4, 0x5c, 0xc3, 0x95, 0xc3, 0x31, 0x46, 0xf8, 0x9f,
	0x07, 0x17, 0x6b, 0xc6, 0x77, 0x84, 0x5b, 0xd9, 0x0a, 0x2e, 0x7e, 0x6b, 0xe0, 0xab, 0x4c, 0xb9,
	0x24, 0x45, 0x9b, 0x2b, 0x5a, 0xea, 0x8b, 0xad, 0xdd, 0x51, 0xcc, 0xaa, 0x2a, 0xa5, 0xc5, 0xc7,
	0x27, 0x49, 0x5f, 0xa2, 0x3c, 0x64, 0x2a, 0x76, 0x31, 0xab, 0xea, 0x3d, 0x51, 0x1e, 0xbe, 0xeb,
	0xf1, 0xcd, 0x1d, 0x39, 0x1e, 0x3d, 0xc9, 0x2a, 0x80, 0x49, 0x6e, 0x66, 0xf8, 0x3b, 0xcc, 0x6c,
	0xc6, 0xf6, 0xfe, 0xef, 0x00, 0x72, 0xbd, 0x14, 0xe5, 0xd5, 0x26, 0xbc, 0x31, 0xee, 0x71, 0x87,
	0x1f, 0xad, 0x1d, 0xb2, 0xeb, 0x7b, 0xb0, 0xab, 0x75, 0xd8, 0x28, 0x82, 0x67, 0x35, 0x27, 0xbf,
	0xb4, 0xb1, 0xaf, 0x97, 0xd7, 0x9d, 0xd1, 0x1d, 0xcc, 0xec, 0xb7, 0x09, 0x14, 0x8e, 0x34, 0x63,
	0x80, 0xa2, 0x57, 0x30, 0x96, 0xc7, 0x60, 0xd3, 0xf6, 0x49, 0x15, 0x21, 0xd1, 0x03, 0x7c, 0x07,
	0xf3, 0x73, 0x5b, 0x38, 0x79, 0xb3, 0xbf, 0x1e, 0xcc, 0xcf, 0xad, 0x7c, 0x48, 0x54, 0xa9, 0xab,
	0xf4, 0xa0, 0x3c, 0x85, 0x4d, 0x6d, 0xcf, 0x08, 0xc3, 0x73, 0xde, 0x50, 0x5a, 0xd2, 0x1f, 0x66,
	0x6e, 0x32, 0xf7, 0x30, 0x14, 0xc2, 0x84, 0x35, 0xb2, 0x6e, 0xa4, 0x08, 0xc7, 0xb7, 0xa3, 0x45,
	0x90, 0xd8, 0xe3, 0xc3, 0x7f, 0x0f, 0xe0, 0x43, 0x29, 0xea, 0x54, 0xe6, 0x1b, 0xc2, 0xd1, 0x7b,
	0x08, 0xba, 0x3a, 0xa2, 0x17, 0xee, 0xca, 0x8f, 0xd5, 0x8b, 0x6e, 0x4e, 0xf0, 0xf6, 0xdd, 0x94,
	0xda, 0x76, 0xb0, 0x53, 0x0f, 0x8a, 0xdb, 0xa9, 0x87, 0x65, 0x45, 0x6f, 0x61, 0xf2, 0xa5, 0xa1,
	0xf9, 0xe6, 0x13, 0x45, 0xd7, 0x6e, 0x1d, 0xad, 0x70, 0xde, 0x07, 0x8d, 0x6a, 0xe1, 0xdd, 0x7b,
	0xd9, 0x53, 0xfd, 0xf7, 0xbd, 0x79, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x8a, 0x15, 0xfd, 0x9e,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DispatcherClient is the client API for Dispatcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DispatcherClient interface {
	CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*CreateJobResponse, error)
	CancelJob(ctx context.Context, in *CancelJobRequest, opts ...grpc.CallOption) (*CancelJobResponse, error)
	PunchIn(ctx context.Context, opts ...grpc.CallOption) (Dispatcher_PunchInClient, error)
}

type dispatcherClient struct {
	cc *grpc.ClientConn
}

func NewDispatcherClient(cc *grpc.ClientConn) DispatcherClient {
	return &dispatcherClient{cc}
}

func (c *dispatcherClient) CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*CreateJobResponse, error) {
	out := new(CreateJobResponse)
	err := c.cc.Invoke(ctx, "/jobs.Dispatcher/CreateJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherClient) CancelJob(ctx context.Context, in *CancelJobRequest, opts ...grpc.CallOption) (*CancelJobResponse, error) {
	out := new(CancelJobResponse)
	err := c.cc.Invoke(ctx, "/jobs.Dispatcher/CancelJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherClient) PunchIn(ctx context.Context, opts ...grpc.CallOption) (Dispatcher_PunchInClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Dispatcher_serviceDesc.Streams[0], "/jobs.Dispatcher/PunchIn", opts...)
	if err != nil {
		return nil, err
	}
	x := &dispatcherPunchInClient{stream}
	return x, nil
}

type Dispatcher_PunchInClient interface {
	Send(*WorkerRequest) error
	Recv() (*WorkerResponse, error)
	grpc.ClientStream
}

type dispatcherPunchInClient struct {
	grpc.ClientStream
}

func (x *dispatcherPunchInClient) Send(m *WorkerRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dispatcherPunchInClient) Recv() (*WorkerResponse, error) {
	m := new(WorkerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DispatcherServer is the server API for Dispatcher service.
type DispatcherServer interface {
	CreateJob(context.Context, *CreateJobRequest) (*CreateJobResponse, error)
	CancelJob(context.Context, *CancelJobRequest) (*CancelJobResponse, error)
	PunchIn(Dispatcher_PunchInServer) error
}

func RegisterDispatcherServer(s *grpc.Server, srv DispatcherServer) {
	s.RegisterService(&_Dispatcher_serviceDesc, srv)
}

func _Dispatcher_CreateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServer).CreateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jobs.Dispatcher/CreateJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServer).CreateJob(ctx, req.(*CreateJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dispatcher_CancelJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServer).CancelJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jobs.Dispatcher/CancelJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServer).CancelJob(ctx, req.(*CancelJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dispatcher_PunchIn_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DispatcherServer).PunchIn(&dispatcherPunchInServer{stream})
}

type Dispatcher_PunchInServer interface {
	Send(*WorkerResponse) error
	Recv() (*WorkerRequest, error)
	grpc.ServerStream
}

type dispatcherPunchInServer struct {
	grpc.ServerStream
}

func (x *dispatcherPunchInServer) Send(m *WorkerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dispatcherPunchInServer) Recv() (*WorkerRequest, error) {
	m := new(WorkerRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Dispatcher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "jobs.Dispatcher",
	HandlerType: (*DispatcherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJob",
			Handler:    _Dispatcher_CreateJob_Handler,
		},
		{
			MethodName: "CancelJob",
			Handler:    _Dispatcher_CancelJob_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PunchIn",
			Handler:       _Dispatcher_PunchIn_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "jobs.proto",
}