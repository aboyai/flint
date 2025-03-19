// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.1
// source: flint/proto/flint.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskId        string                 `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	DagId         string                 `protobuf:"bytes,2,opt,name=dag_id,json=dagId,proto3" json:"dag_id,omitempty"`
	Payload       string                 `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"` // Task-specific data (e.g., JSON)
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskRequest) Reset() {
	*x = TaskRequest{}
	mi := &file_flint_proto_flint_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRequest) ProtoMessage() {}

func (x *TaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flint_proto_flint_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRequest.ProtoReflect.Descriptor instead.
func (*TaskRequest) Descriptor() ([]byte, []int) {
	return file_flint_proto_flint_proto_rawDescGZIP(), []int{0}
}

func (x *TaskRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *TaskRequest) GetDagId() string {
	if x != nil {
		return x.DagId
	}
	return ""
}

func (x *TaskRequest) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type TaskResult struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskId        string                 `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	DagId         string                 `protobuf:"bytes,2,opt,name=dag_id,json=dagId,proto3" json:"dag_id,omitempty"`
	Status        string                 `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"` // success/failure
	ResultData    string                 `protobuf:"bytes,4,opt,name=result_data,json=resultData,proto3" json:"result_data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskResult) Reset() {
	*x = TaskResult{}
	mi := &file_flint_proto_flint_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResult) ProtoMessage() {}

func (x *TaskResult) ProtoReflect() protoreflect.Message {
	mi := &file_flint_proto_flint_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResult.ProtoReflect.Descriptor instead.
func (*TaskResult) Descriptor() ([]byte, []int) {
	return file_flint_proto_flint_proto_rawDescGZIP(), []int{1}
}

func (x *TaskResult) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *TaskResult) GetDagId() string {
	if x != nil {
		return x.DagId
	}
	return ""
}

func (x *TaskResult) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *TaskResult) GetResultData() string {
	if x != nil {
		return x.ResultData
	}
	return ""
}

type HeartbeatRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WorkerId      string                 `protobuf:"bytes,1,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`
	Capacity      int32                  `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"` // available slots
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HeartbeatRequest) Reset() {
	*x = HeartbeatRequest{}
	mi := &file_flint_proto_flint_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HeartbeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatRequest) ProtoMessage() {}

func (x *HeartbeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flint_proto_flint_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatRequest.ProtoReflect.Descriptor instead.
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return file_flint_proto_flint_proto_rawDescGZIP(), []int{2}
}

func (x *HeartbeatRequest) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

func (x *HeartbeatRequest) GetCapacity() int32 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

type HeartbeatResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Alive         bool                   `protobuf:"varint,1,opt,name=alive,proto3" json:"alive,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HeartbeatResponse) Reset() {
	*x = HeartbeatResponse{}
	mi := &file_flint_proto_flint_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HeartbeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatResponse) ProtoMessage() {}

func (x *HeartbeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flint_proto_flint_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatResponse.ProtoReflect.Descriptor instead.
func (*HeartbeatResponse) Descriptor() ([]byte, []int) {
	return file_flint_proto_flint_proto_rawDescGZIP(), []int{3}
}

func (x *HeartbeatResponse) GetAlive() bool {
	if x != nil {
		return x.Alive
	}
	return false
}

var File_flint_proto_flint_proto protoreflect.FileDescriptor

var file_flint_proto_flint_proto_rawDesc = string([]byte{
	0x0a, 0x17, 0x66, 0x6c, 0x69, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6c,
	0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x66, 0x6c, 0x69, 0x6e, 0x74,
	0x22, 0x57, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x64, 0x61, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x61, 0x67, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x75, 0x0a, 0x0a, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64,
	0x12, 0x15, 0x0a, 0x06, 0x64, 0x61, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x64, 0x61, 0x67, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x4b, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x22, 0x29, 0x0a,
	0x11, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x32, 0x85, 0x01, 0x0a, 0x0d, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x0b, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x12, 0x2e, 0x66, 0x6c, 0x69, 0x6e,
	0x74, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x66, 0x6c, 0x69, 0x6e, 0x74, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x3e, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x17, 0x2e,
	0x66, 0x6c, 0x69, 0x6e, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x66, 0x6c, 0x69, 0x6e, 0x74, 0x2e, 0x48,
	0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x13, 0x5a, 0x11, 0x66, 0x6c, 0x69, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_flint_proto_flint_proto_rawDescOnce sync.Once
	file_flint_proto_flint_proto_rawDescData []byte
)

func file_flint_proto_flint_proto_rawDescGZIP() []byte {
	file_flint_proto_flint_proto_rawDescOnce.Do(func() {
		file_flint_proto_flint_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_flint_proto_flint_proto_rawDesc), len(file_flint_proto_flint_proto_rawDesc)))
	})
	return file_flint_proto_flint_proto_rawDescData
}

var file_flint_proto_flint_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_flint_proto_flint_proto_goTypes = []any{
	(*TaskRequest)(nil),       // 0: flint.TaskRequest
	(*TaskResult)(nil),        // 1: flint.TaskResult
	(*HeartbeatRequest)(nil),  // 2: flint.HeartbeatRequest
	(*HeartbeatResponse)(nil), // 3: flint.HeartbeatResponse
}
var file_flint_proto_flint_proto_depIdxs = []int32{
	0, // 0: flint.WorkerService.ExecuteTask:input_type -> flint.TaskRequest
	2, // 1: flint.WorkerService.Heartbeat:input_type -> flint.HeartbeatRequest
	1, // 2: flint.WorkerService.ExecuteTask:output_type -> flint.TaskResult
	3, // 3: flint.WorkerService.Heartbeat:output_type -> flint.HeartbeatResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_flint_proto_flint_proto_init() }
func file_flint_proto_flint_proto_init() {
	if File_flint_proto_flint_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_flint_proto_flint_proto_rawDesc), len(file_flint_proto_flint_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_flint_proto_flint_proto_goTypes,
		DependencyIndexes: file_flint_proto_flint_proto_depIdxs,
		MessageInfos:      file_flint_proto_flint_proto_msgTypes,
	}.Build()
	File_flint_proto_flint_proto = out.File
	file_flint_proto_flint_proto_goTypes = nil
	file_flint_proto_flint_proto_depIdxs = nil
}
