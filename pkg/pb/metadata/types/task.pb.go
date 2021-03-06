// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metadata/types/task.proto

package pbtypes

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type SubTaskAction int32

const (
	SubTaskAction_NULL                      SubTaskAction = 0
	SubTaskAction_StartConfd                SubTaskAction = 1
	SubTaskAction_RegisterMetadata          SubTaskAction = 2
	SubTaskAction_DeregisterMetadata        SubTaskAction = 3
	SubTaskAction_RegisterCmd               SubTaskAction = 4
	SubTaskAction_DeregisterCmd             SubTaskAction = 5
	SubTaskAction_GetTaskStatus             SubTaskAction = 6
	SubTaskAction_StopConfd                 SubTaskAction = 7
	SubTaskAction_RegisterMetadataMapping   SubTaskAction = 8
	SubTaskAction_DeregisterMetadataMapping SubTaskAction = 9
)

var SubTaskAction_name = map[int32]string{
	0: "NULL",
	1: "StartConfd",
	2: "RegisterMetadata",
	3: "DeregisterMetadata",
	4: "RegisterCmd",
	5: "DeregisterCmd",
	6: "GetTaskStatus",
	7: "StopConfd",
	8: "RegisterMetadataMapping",
	9: "DeregisterMetadataMapping",
}

var SubTaskAction_value = map[string]int32{
	"NULL":                      0,
	"StartConfd":                1,
	"RegisterMetadata":          2,
	"DeregisterMetadata":        3,
	"RegisterCmd":               4,
	"DeregisterCmd":             5,
	"GetTaskStatus":             6,
	"StopConfd":                 7,
	"RegisterMetadataMapping":   8,
	"DeregisterMetadataMapping": 9,
}

func (x SubTaskAction) String() string {
	return proto.EnumName(SubTaskAction_name, int32(x))
}

func (SubTaskAction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_48a76b9d476194b0, []int{0}
}

type SubTaskId struct {
	TaskId               string   `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubTaskId) Reset()         { *m = SubTaskId{} }
func (m *SubTaskId) String() string { return proto.CompactTextString(m) }
func (*SubTaskId) ProtoMessage()    {}
func (*SubTaskId) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a76b9d476194b0, []int{0}
}

func (m *SubTaskId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubTaskId.Unmarshal(m, b)
}
func (m *SubTaskId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubTaskId.Marshal(b, m, deterministic)
}
func (m *SubTaskId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubTaskId.Merge(m, src)
}
func (m *SubTaskId) XXX_Size() int {
	return xxx_messageInfo_SubTaskId.Size(m)
}
func (m *SubTaskId) XXX_DiscardUnknown() {
	xxx_messageInfo_SubTaskId.DiscardUnknown(m)
}

var xxx_messageInfo_SubTaskId proto.InternalMessageInfo

func (m *SubTaskId) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

type SubTaskMessage struct {
	Action               string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action"`
	TaskId               string   `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id"`
	Directive            string   `protobuf:"bytes,3,opt,name=directive,proto3" json:"directive"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubTaskMessage) Reset()         { *m = SubTaskMessage{} }
func (m *SubTaskMessage) String() string { return proto.CompactTextString(m) }
func (*SubTaskMessage) ProtoMessage()    {}
func (*SubTaskMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a76b9d476194b0, []int{1}
}

func (m *SubTaskMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubTaskMessage.Unmarshal(m, b)
}
func (m *SubTaskMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubTaskMessage.Marshal(b, m, deterministic)
}
func (m *SubTaskMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubTaskMessage.Merge(m, src)
}
func (m *SubTaskMessage) XXX_Size() int {
	return xxx_messageInfo_SubTaskMessage.Size(m)
}
func (m *SubTaskMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SubTaskMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SubTaskMessage proto.InternalMessageInfo

func (m *SubTaskMessage) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *SubTaskMessage) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *SubTaskMessage) GetDirective() string {
	if m != nil {
		return m.Directive
	}
	return ""
}

type SubTaskStatus struct {
	TaskId               string   `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubTaskStatus) Reset()         { *m = SubTaskStatus{} }
func (m *SubTaskStatus) String() string { return proto.CompactTextString(m) }
func (*SubTaskStatus) ProtoMessage()    {}
func (*SubTaskStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a76b9d476194b0, []int{2}
}

func (m *SubTaskStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubTaskStatus.Unmarshal(m, b)
}
func (m *SubTaskStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubTaskStatus.Marshal(b, m, deterministic)
}
func (m *SubTaskStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubTaskStatus.Merge(m, src)
}
func (m *SubTaskStatus) XXX_Size() int {
	return xxx_messageInfo_SubTaskStatus.Size(m)
}
func (m *SubTaskStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_SubTaskStatus.DiscardUnknown(m)
}

var xxx_messageInfo_SubTaskStatus proto.InternalMessageInfo

func (m *SubTaskStatus) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *SubTaskStatus) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

//
//{
//"action": "StartConfd",
//"taskId": "t-abcdefgh",
//"directive": {"frontgateId": "cl-abcdefgh", "droneIp": "192.168.0.1", "timeout": 600}
//}
type SubTask_StartConfd struct {
	Action               string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action"`
	TaskId               string   `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id"`
	FrontgateId          string   `protobuf:"bytes,3,opt,name=frontgate_id,json=frontgateId,proto3" json:"frontgate_id"`
	DroneIp              string   `protobuf:"bytes,4,opt,name=drone_ip,json=droneIp,proto3" json:"drone_ip"`
	Timeout              int32    `protobuf:"varint,5,opt,name=timeout,proto3" json:"timeout"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubTask_StartConfd) Reset()         { *m = SubTask_StartConfd{} }
func (m *SubTask_StartConfd) String() string { return proto.CompactTextString(m) }
func (*SubTask_StartConfd) ProtoMessage()    {}
func (*SubTask_StartConfd) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a76b9d476194b0, []int{3}
}

func (m *SubTask_StartConfd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubTask_StartConfd.Unmarshal(m, b)
}
func (m *SubTask_StartConfd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubTask_StartConfd.Marshal(b, m, deterministic)
}
func (m *SubTask_StartConfd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubTask_StartConfd.Merge(m, src)
}
func (m *SubTask_StartConfd) XXX_Size() int {
	return xxx_messageInfo_SubTask_StartConfd.Size(m)
}
func (m *SubTask_StartConfd) XXX_DiscardUnknown() {
	xxx_messageInfo_SubTask_StartConfd.DiscardUnknown(m)
}

var xxx_messageInfo_SubTask_StartConfd proto.InternalMessageInfo

func (m *SubTask_StartConfd) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *SubTask_StartConfd) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *SubTask_StartConfd) GetFrontgateId() string {
	if m != nil {
		return m.FrontgateId
	}
	return ""
}

func (m *SubTask_StartConfd) GetDroneIp() string {
	if m != nil {
		return m.DroneIp
	}
	return ""
}

func (m *SubTask_StartConfd) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

//
//{
//"action": "StopConfd",
//"taskId": "t-abcdefgh",
//"directive": {"frontgateId": "cl-abcdefgh", "droneIp": "192.168.0.1", "timeout": 600}
//}
type SubTask_StopConfd struct {
	Action               string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action"`
	TaskId               string   `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id"`
	FrontgateId          string   `protobuf:"bytes,3,opt,name=frontgate_id,json=frontgateId,proto3" json:"frontgate_id"`
	DroneIp              string   `protobuf:"bytes,4,opt,name=drone_ip,json=droneIp,proto3" json:"drone_ip"`
	Timeout              int32    `protobuf:"varint,5,opt,name=timeout,proto3" json:"timeout"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubTask_StopConfd) Reset()         { *m = SubTask_StopConfd{} }
func (m *SubTask_StopConfd) String() string { return proto.CompactTextString(m) }
func (*SubTask_StopConfd) ProtoMessage()    {}
func (*SubTask_StopConfd) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a76b9d476194b0, []int{4}
}

func (m *SubTask_StopConfd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubTask_StopConfd.Unmarshal(m, b)
}
func (m *SubTask_StopConfd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubTask_StopConfd.Marshal(b, m, deterministic)
}
func (m *SubTask_StopConfd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubTask_StopConfd.Merge(m, src)
}
func (m *SubTask_StopConfd) XXX_Size() int {
	return xxx_messageInfo_SubTask_StopConfd.Size(m)
}
func (m *SubTask_StopConfd) XXX_DiscardUnknown() {
	xxx_messageInfo_SubTask_StopConfd.DiscardUnknown(m)
}

var xxx_messageInfo_SubTask_StopConfd proto.InternalMessageInfo

func (m *SubTask_StopConfd) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *SubTask_StopConfd) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *SubTask_StopConfd) GetFrontgateId() string {
	if m != nil {
		return m.FrontgateId
	}
	return ""
}

func (m *SubTask_StopConfd) GetDroneIp() string {
	if m != nil {
		return m.DroneIp
	}
	return ""
}

func (m *SubTask_StopConfd) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

//
//{
//"action": "RegisterMetadata",
//"taskId": "t-abcdefgh",
//"directive": {"frontgateId": "cl-abcdefgh", "cnodes": "{\"key\", \"value\"}", "timeout": 600, "retry": 5}
//}
type SubTask_RegisterMetadata struct {
	Action               string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action"`
	TaskId               string   `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id"`
	FrontgateId          string   `protobuf:"bytes,3,opt,name=frontgate_id,json=frontgateId,proto3" json:"frontgate_id"`
	Cnodes               string   `protobuf:"bytes,4,opt,name=cnodes,proto3" json:"cnodes"`
	Timeout              int32    `protobuf:"varint,5,opt,name=timeout,proto3" json:"timeout"`
	Retry                int32    `protobuf:"varint,6,opt,name=retry,proto3" json:"retry"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubTask_RegisterMetadata) Reset()         { *m = SubTask_RegisterMetadata{} }
func (m *SubTask_RegisterMetadata) String() string { return proto.CompactTextString(m) }
func (*SubTask_RegisterMetadata) ProtoMessage()    {}
func (*SubTask_RegisterMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a76b9d476194b0, []int{5}
}

func (m *SubTask_RegisterMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubTask_RegisterMetadata.Unmarshal(m, b)
}
func (m *SubTask_RegisterMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubTask_RegisterMetadata.Marshal(b, m, deterministic)
}
func (m *SubTask_RegisterMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubTask_RegisterMetadata.Merge(m, src)
}
func (m *SubTask_RegisterMetadata) XXX_Size() int {
	return xxx_messageInfo_SubTask_RegisterMetadata.Size(m)
}
func (m *SubTask_RegisterMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_SubTask_RegisterMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_SubTask_RegisterMetadata proto.InternalMessageInfo

func (m *SubTask_RegisterMetadata) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *SubTask_RegisterMetadata) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *SubTask_RegisterMetadata) GetFrontgateId() string {
	if m != nil {
		return m.FrontgateId
	}
	return ""
}

func (m *SubTask_RegisterMetadata) GetCnodes() string {
	if m != nil {
		return m.Cnodes
	}
	return ""
}

func (m *SubTask_RegisterMetadata) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *SubTask_RegisterMetadata) GetRetry() int32 {
	if m != nil {
		return m.Retry
	}
	return 0
}

//
//{
//"action": "DeregisterMetadata",
//"taskId": "t-abcdefgh",
//"directive": {"frontgateId": "cl-abcdefgh", "droneIp": "192.168.0.1", cnodes": "{\"key\", \"value\"}", "timeout": 600, "retry": 5}
//}
type SubTask_DeregisterMetadata struct {
	Action               string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action"`
	TaskId               string   `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id"`
	FrontgateId          string   `protobuf:"bytes,3,opt,name=frontgate_id,json=frontgateId,proto3" json:"frontgate_id"`
	DroneIp              string   `protobuf:"bytes,4,opt,name=drone_ip,json=droneIp,proto3" json:"drone_ip"`
	Cnodes               string   `protobuf:"bytes,5,opt,name=cnodes,proto3" json:"cnodes"`
	Timeout              int32    `protobuf:"varint,6,opt,name=timeout,proto3" json:"timeout"`
	Retry                int32    `protobuf:"varint,7,opt,name=retry,proto3" json:"retry"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubTask_DeregisterMetadata) Reset()         { *m = SubTask_DeregisterMetadata{} }
func (m *SubTask_DeregisterMetadata) String() string { return proto.CompactTextString(m) }
func (*SubTask_DeregisterMetadata) ProtoMessage()    {}
func (*SubTask_DeregisterMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a76b9d476194b0, []int{6}
}

func (m *SubTask_DeregisterMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubTask_DeregisterMetadata.Unmarshal(m, b)
}
func (m *SubTask_DeregisterMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubTask_DeregisterMetadata.Marshal(b, m, deterministic)
}
func (m *SubTask_DeregisterMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubTask_DeregisterMetadata.Merge(m, src)
}
func (m *SubTask_DeregisterMetadata) XXX_Size() int {
	return xxx_messageInfo_SubTask_DeregisterMetadata.Size(m)
}
func (m *SubTask_DeregisterMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_SubTask_DeregisterMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_SubTask_DeregisterMetadata proto.InternalMessageInfo

func (m *SubTask_DeregisterMetadata) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *SubTask_DeregisterMetadata) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *SubTask_DeregisterMetadata) GetFrontgateId() string {
	if m != nil {
		return m.FrontgateId
	}
	return ""
}

func (m *SubTask_DeregisterMetadata) GetDroneIp() string {
	if m != nil {
		return m.DroneIp
	}
	return ""
}

func (m *SubTask_DeregisterMetadata) GetCnodes() string {
	if m != nil {
		return m.Cnodes
	}
	return ""
}

func (m *SubTask_DeregisterMetadata) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *SubTask_DeregisterMetadata) GetRetry() int32 {
	if m != nil {
		return m.Retry
	}
	return 0
}

//
//{
//"action": "RegisterCmd",
//"taskId": "t-abcdefgh",
//"directive": {"frontgateId": "cl-abcdefgh", "droneIp": "192.168.0.1", cnodes": "{\"key\", \"value\"}", "timeout": 600, "retry": 5}
//}
type SubTask_RegisterCmd struct {
	Action               string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action"`
	TaskId               string   `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id"`
	FrontgateId          string   `protobuf:"bytes,3,opt,name=frontgate_id,json=frontgateId,proto3" json:"frontgate_id"`
	DroneIp              string   `protobuf:"bytes,4,opt,name=drone_ip,json=droneIp,proto3" json:"drone_ip"`
	Cnodes               string   `protobuf:"bytes,5,opt,name=cnodes,proto3" json:"cnodes"`
	Timeout              int32    `protobuf:"varint,6,opt,name=timeout,proto3" json:"timeout"`
	Retry                int32    `protobuf:"varint,7,opt,name=retry,proto3" json:"retry"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubTask_RegisterCmd) Reset()         { *m = SubTask_RegisterCmd{} }
func (m *SubTask_RegisterCmd) String() string { return proto.CompactTextString(m) }
func (*SubTask_RegisterCmd) ProtoMessage()    {}
func (*SubTask_RegisterCmd) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a76b9d476194b0, []int{7}
}

func (m *SubTask_RegisterCmd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubTask_RegisterCmd.Unmarshal(m, b)
}
func (m *SubTask_RegisterCmd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubTask_RegisterCmd.Marshal(b, m, deterministic)
}
func (m *SubTask_RegisterCmd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubTask_RegisterCmd.Merge(m, src)
}
func (m *SubTask_RegisterCmd) XXX_Size() int {
	return xxx_messageInfo_SubTask_RegisterCmd.Size(m)
}
func (m *SubTask_RegisterCmd) XXX_DiscardUnknown() {
	xxx_messageInfo_SubTask_RegisterCmd.DiscardUnknown(m)
}

var xxx_messageInfo_SubTask_RegisterCmd proto.InternalMessageInfo

func (m *SubTask_RegisterCmd) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *SubTask_RegisterCmd) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *SubTask_RegisterCmd) GetFrontgateId() string {
	if m != nil {
		return m.FrontgateId
	}
	return ""
}

func (m *SubTask_RegisterCmd) GetDroneIp() string {
	if m != nil {
		return m.DroneIp
	}
	return ""
}

func (m *SubTask_RegisterCmd) GetCnodes() string {
	if m != nil {
		return m.Cnodes
	}
	return ""
}

func (m *SubTask_RegisterCmd) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *SubTask_RegisterCmd) GetRetry() int32 {
	if m != nil {
		return m.Retry
	}
	return 0
}

//
//{
//"action": "DeregisterCmd",
//"taskId": "t-abcdefgh",
//"directive": {"frontgateId": "cl-abcdefgh", "droneIp": "192.168.0.1", cnodes": "{\"key\", \"value\"}", "timeout": 600, "retry": 5}
//}
type SubTask_DeregisterCmd struct {
	Action               string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action"`
	TaskId               string   `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id"`
	FrontgateId          string   `protobuf:"bytes,3,opt,name=frontgate_id,json=frontgateId,proto3" json:"frontgate_id"`
	DroneIp              string   `protobuf:"bytes,4,opt,name=drone_ip,json=droneIp,proto3" json:"drone_ip"`
	Cnodes               string   `protobuf:"bytes,5,opt,name=cnodes,proto3" json:"cnodes"`
	Timeout              int32    `protobuf:"varint,6,opt,name=timeout,proto3" json:"timeout"`
	Retry                int32    `protobuf:"varint,7,opt,name=retry,proto3" json:"retry"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubTask_DeregisterCmd) Reset()         { *m = SubTask_DeregisterCmd{} }
func (m *SubTask_DeregisterCmd) String() string { return proto.CompactTextString(m) }
func (*SubTask_DeregisterCmd) ProtoMessage()    {}
func (*SubTask_DeregisterCmd) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a76b9d476194b0, []int{8}
}

func (m *SubTask_DeregisterCmd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubTask_DeregisterCmd.Unmarshal(m, b)
}
func (m *SubTask_DeregisterCmd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubTask_DeregisterCmd.Marshal(b, m, deterministic)
}
func (m *SubTask_DeregisterCmd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubTask_DeregisterCmd.Merge(m, src)
}
func (m *SubTask_DeregisterCmd) XXX_Size() int {
	return xxx_messageInfo_SubTask_DeregisterCmd.Size(m)
}
func (m *SubTask_DeregisterCmd) XXX_DiscardUnknown() {
	xxx_messageInfo_SubTask_DeregisterCmd.DiscardUnknown(m)
}

var xxx_messageInfo_SubTask_DeregisterCmd proto.InternalMessageInfo

func (m *SubTask_DeregisterCmd) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *SubTask_DeregisterCmd) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *SubTask_DeregisterCmd) GetFrontgateId() string {
	if m != nil {
		return m.FrontgateId
	}
	return ""
}

func (m *SubTask_DeregisterCmd) GetDroneIp() string {
	if m != nil {
		return m.DroneIp
	}
	return ""
}

func (m *SubTask_DeregisterCmd) GetCnodes() string {
	if m != nil {
		return m.Cnodes
	}
	return ""
}

func (m *SubTask_DeregisterCmd) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *SubTask_DeregisterCmd) GetRetry() int32 {
	if m != nil {
		return m.Retry
	}
	return 0
}

//
//{
//"action": "GetTaskStatus",
//"taskId": "t-abcdefgh"
//}
type SubTask_GetTaskStatus struct {
	Action               string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action"`
	TaskId               string   `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubTask_GetTaskStatus) Reset()         { *m = SubTask_GetTaskStatus{} }
func (m *SubTask_GetTaskStatus) String() string { return proto.CompactTextString(m) }
func (*SubTask_GetTaskStatus) ProtoMessage()    {}
func (*SubTask_GetTaskStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a76b9d476194b0, []int{9}
}

func (m *SubTask_GetTaskStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubTask_GetTaskStatus.Unmarshal(m, b)
}
func (m *SubTask_GetTaskStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubTask_GetTaskStatus.Marshal(b, m, deterministic)
}
func (m *SubTask_GetTaskStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubTask_GetTaskStatus.Merge(m, src)
}
func (m *SubTask_GetTaskStatus) XXX_Size() int {
	return xxx_messageInfo_SubTask_GetTaskStatus.Size(m)
}
func (m *SubTask_GetTaskStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_SubTask_GetTaskStatus.DiscardUnknown(m)
}

var xxx_messageInfo_SubTask_GetTaskStatus proto.InternalMessageInfo

func (m *SubTask_GetTaskStatus) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *SubTask_GetTaskStatus) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func init() {
	proto.RegisterEnum("metadata.types.SubTaskAction", SubTaskAction_name, SubTaskAction_value)
	proto.RegisterType((*SubTaskId)(nil), "metadata.types.SubTaskId")
	proto.RegisterType((*SubTaskMessage)(nil), "metadata.types.SubTaskMessage")
	proto.RegisterType((*SubTaskStatus)(nil), "metadata.types.SubTaskStatus")
	proto.RegisterType((*SubTask_StartConfd)(nil), "metadata.types.SubTask_StartConfd")
	proto.RegisterType((*SubTask_StopConfd)(nil), "metadata.types.SubTask_StopConfd")
	proto.RegisterType((*SubTask_RegisterMetadata)(nil), "metadata.types.SubTask_RegisterMetadata")
	proto.RegisterType((*SubTask_DeregisterMetadata)(nil), "metadata.types.SubTask_DeregisterMetadata")
	proto.RegisterType((*SubTask_RegisterCmd)(nil), "metadata.types.SubTask_RegisterCmd")
	proto.RegisterType((*SubTask_DeregisterCmd)(nil), "metadata.types.SubTask_DeregisterCmd")
	proto.RegisterType((*SubTask_GetTaskStatus)(nil), "metadata.types.SubTask_GetTaskStatus")
}

func init() { proto.RegisterFile("metadata/types/task.proto", fileDescriptor_48a76b9d476194b0) }

var fileDescriptor_48a76b9d476194b0 = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x66, 0xdb, 0xd8, 0x8e, 0xa7, 0x24, 0x6c, 0x97, 0x92, 0x3a, 0xfc, 0x48, 0xc5, 0xe2, 0x50,
	0x71, 0x88, 0x0f, 0x48, 0x08, 0x89, 0x0b, 0x50, 0x24, 0x88, 0xd4, 0x70, 0x48, 0xe0, 0xc2, 0x25,
	0xda, 0xc4, 0x5b, 0x6b, 0x15, 0xc5, 0xbb, 0x5a, 0x4f, 0x10, 0x7d, 0x17, 0x78, 0x0d, 0x9e, 0x02,
	0x71, 0xe1, 0xc4, 0xdb, 0xa0, 0xac, 0xed, 0xd8, 0x49, 0x14, 0xa4, 0x1e, 0x40, 0xf4, 0x64, 0x7d,
	0xf3, 0x8d, 0x66, 0xbe, 0x6f, 0x77, 0x3c, 0x0b, 0xdd, 0xb9, 0x40, 0x1e, 0x73, 0xe4, 0x11, 0x5e,
	0x6a, 0x91, 0x45, 0xc8, 0xb3, 0x59, 0x4f, 0x1b, 0x85, 0x8a, 0xb5, 0x4b, 0xaa, 0x67, 0xa9, 0xf0,
	0x11, 0xf8, 0xa3, 0xc5, 0xe4, 0x3d, 0xcf, 0x66, 0xfd, 0x98, 0x1d, 0x83, 0xb7, 0x4c, 0x1d, 0xcb,
	0x38, 0x20, 0x27, 0xe4, 0xd4, 0x1f, 0xba, 0x68, 0x89, 0x70, 0x0c, 0xed, 0x22, 0x6b, 0x20, 0xb2,
	0x8c, 0x27, 0x82, 0x75, 0xc0, 0xe5, 0x53, 0x94, 0x2a, 0x2d, 0x33, 0x73, 0x54, 0x2f, 0xb1, 0x57,
	0x2f, 0xc1, 0xee, 0x83, 0x1f, 0x4b, 0x23, 0xa6, 0x28, 0x3f, 0x89, 0x60, 0xdf, 0x52, 0x55, 0x20,
	0x7c, 0x01, 0xad, 0xa2, 0xc1, 0x08, 0x39, 0x2e, 0xb2, 0x9d, 0x52, 0x96, 0x8d, 0x33, 0x9b, 0x52,
	0xd6, 0xcf, 0x51, 0xf8, 0x95, 0x00, 0x2b, 0x4a, 0x8c, 0x47, 0xc8, 0x0d, 0x9e, 0xa9, 0xf4, 0x22,
	0xbe, 0xba, 0xce, 0x87, 0x70, 0xf3, 0xc2, 0xa8, 0x14, 0x13, 0x8e, 0x62, 0xc9, 0xe6, 0x52, 0x0f,
	0x56, 0xb1, 0x7e, 0xcc, 0xba, 0xd0, 0x8c, 0x8d, 0x4a, 0xc5, 0x58, 0xea, 0xa0, 0x61, 0x69, 0xcf,
	0xe2, 0xbe, 0x66, 0x01, 0x78, 0x28, 0xe7, 0x42, 0x2d, 0x30, 0x70, 0x4e, 0xc8, 0xa9, 0x33, 0x2c,
	0x61, 0xf8, 0x85, 0xc0, 0x61, 0xa5, 0x4f, 0xe9, 0xff, 0x4c, 0xde, 0x37, 0x02, 0x41, 0x29, 0x6f,
	0x28, 0x12, 0x99, 0xa1, 0x30, 0x83, 0x62, 0x54, 0xfe, 0x8a, 0xca, 0x0e, 0xb8, 0xd3, 0x54, 0xc5,
	0x22, 0x2b, 0x34, 0x16, 0x68, 0xb7, 0x44, 0x76, 0x04, 0x8e, 0x11, 0x68, 0x2e, 0x03, 0xd7, 0xc6,
	0x73, 0x10, 0xfe, 0x24, 0x70, 0xb7, 0x14, 0xfe, 0x5a, 0x98, 0x7f, 0x21, 0xfd, 0x0f, 0x07, 0x5c,
	0xb9, 0x72, 0x76, 0xb9, 0x72, 0x77, 0xb8, 0xf2, 0xea, 0xae, 0xbe, 0x13, 0xb8, 0xbd, 0x79, 0x1d,
	0x67, 0xf3, 0xf8, 0xba, 0xda, 0xf9, 0x41, 0xe0, 0xce, 0xf6, 0x25, 0x5d, 0x63, 0x43, 0x6f, 0x2b,
	0x3f, 0x6f, 0x04, 0xd6, 0xf6, 0xd6, 0x55, 0xfd, 0x3c, 0xfe, 0x45, 0x56, 0xab, 0xef, 0x65, 0x9e,
	0xda, 0x84, 0xc6, 0xbb, 0x0f, 0xe7, 0xe7, 0xf4, 0x06, 0x6b, 0x03, 0x54, 0xab, 0x8c, 0x12, 0x76,
	0x04, 0x74, 0xf3, 0xdf, 0xa4, 0x7b, 0xac, 0x03, 0x6c, 0x7b, 0xf0, 0xe9, 0x3e, 0xbb, 0x05, 0x07,
	0xb5, 0xd1, 0xa1, 0x0d, 0x76, 0x08, 0xad, 0xb5, 0xc3, 0xa7, 0xce, 0x32, 0xb4, 0xa6, 0x9f, 0xba,
	0xac, 0x05, 0xfe, 0x6a, 0x3f, 0x51, 0x8f, 0xdd, 0x83, 0xe3, 0xcd, 0x9e, 0x03, 0xae, 0xb5, 0x4c,
	0x13, 0xda, 0x64, 0x0f, 0xa0, 0xbb, 0xdd, 0xba, 0xa4, 0xfd, 0x57, 0xcf, 0x3e, 0x3e, 0x55, 0x5a,
	0xa4, 0x5a, 0xa2, 0x91, 0x9f, 0x7b, 0x52, 0x45, 0x15, 0x8a, 0xf4, 0x2c, 0x89, 0xf4, 0x24, 0x5a,
	0x7f, 0xa9, 0x9e, 0xeb, 0x89, 0xfd, 0x4e, 0x5c, 0xfb, 0x5a, 0x3d, 0xf9, 0x1d, 0x00, 0x00, 0xff,
	0xff, 0x66, 0xfc, 0x71, 0x49, 0xca, 0x06, 0x00, 0x00,
}
