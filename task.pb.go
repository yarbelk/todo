// Code generated by protoc-gen-go.
// source: github.com/yarbelk/todo/task.proto
// DO NOT EDIT!

/*
Package todo is a generated protocol buffer package.

It is generated from these files:
	github.com/yarbelk/todo/task.proto

It has these top-level messages:
	TaskDefinition
	TaskID
	TaskList
	AllTasksParams
*/
package todo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TaskDefinition struct {
	Id          string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Completed   bool   `protobuf:"varint,3,opt,name=completed" json:"completed,omitempty"`
}

func (m *TaskDefinition) Reset()                    { *m = TaskDefinition{} }
func (m *TaskDefinition) String() string            { return proto.CompactTextString(m) }
func (*TaskDefinition) ProtoMessage()               {}
func (*TaskDefinition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type TaskID struct {
	Id string `protobuf:"bytes,4,opt,name=id" json:"id,omitempty"`
}

func (m *TaskID) Reset()                    { *m = TaskID{} }
func (m *TaskID) String() string            { return proto.CompactTextString(m) }
func (*TaskID) ProtoMessage()               {}
func (*TaskID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type TaskList struct {
	Tasks []*TaskDefinition `protobuf:"bytes,5,rep,name=tasks" json:"tasks,omitempty"`
}

func (m *TaskList) Reset()                    { *m = TaskList{} }
func (m *TaskList) String() string            { return proto.CompactTextString(m) }
func (*TaskList) ProtoMessage()               {}
func (*TaskList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TaskList) GetTasks() []*TaskDefinition {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type AllTasksParams struct {
}

func (m *AllTasksParams) Reset()                    { *m = AllTasksParams{} }
func (m *AllTasksParams) String() string            { return proto.CompactTextString(m) }
func (*AllTasksParams) ProtoMessage()               {}
func (*AllTasksParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*TaskDefinition)(nil), "todo.TaskDefinition")
	proto.RegisterType((*TaskID)(nil), "todo.TaskID")
	proto.RegisterType((*TaskList)(nil), "todo.TaskList")
	proto.RegisterType((*AllTasksParams)(nil), "todo.AllTasksParams")
}

func init() { proto.RegisterFile("github.com/yarbelk/todo/task.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x8f, 0x3f, 0x4b, 0xc5, 0x30,
	0x14, 0x47, 0xe9, 0x5f, 0xda, 0x5b, 0x28, 0x12, 0x1c, 0x32, 0x38, 0x84, 0x4c, 0xc5, 0x21, 0x05,
	0x05, 0x77, 0xa1, 0x8b, 0xe0, 0x20, 0xc5, 0x0f, 0x60, 0xda, 0x44, 0xbd, 0x34, 0x6d, 0x4a, 0x12,
	0x07, 0xbf, 0xbd, 0xa4, 0x2a, 0xef, 0xbd, 0xf5, 0x9c, 0xf0, 0xcb, 0xb9, 0xc0, 0x3f, 0x30, 0x7c,
	0x7e, 0x4d, 0x62, 0xb6, 0x6b, 0xff, 0x2d, 0xdd, 0xa4, 0xcd, 0xd2, 0x07, 0xab, 0x6c, 0x1f, 0xa4,
	0x5f, 0xc4, 0xee, 0x6c, 0xb0, 0x24, 0x8f, 0x80, 0xbf, 0x41, 0xfb, 0x2a, 0xfd, 0x32, 0xe8, 0x77,
	0xdc, 0x30, 0xa0, 0xdd, 0x48, 0x0b, 0x29, 0x2a, 0x9a, 0xb0, 0xa4, 0xab, 0xc7, 0x14, 0x15, 0x61,
	0xd0, 0x28, 0xed, 0x67, 0x87, 0x7b, 0xd4, 0x34, 0x3d, 0xc4, 0x39, 0x22, 0x37, 0x50, 0xcf, 0x76,
	0xdd, 0x8d, 0x0e, 0x5a, 0xd1, 0x8c, 0x25, 0x5d, 0x35, 0x9e, 0x00, 0xa7, 0x50, 0xc6, 0x1f, 0x9e,
	0x86, 0xbf, 0xe5, 0xfc, 0x7f, 0x99, 0x3f, 0x40, 0x15, 0xcd, 0x33, 0xfa, 0x40, 0x6e, 0xa1, 0x88,
	0x6d, 0x9e, 0x16, 0x2c, 0xeb, 0x9a, 0xbb, 0x6b, 0x11, 0xeb, 0xc4, 0x65, 0xda, 0xf8, 0xfb, 0x84,
	0x5f, 0x41, 0xfb, 0x68, 0x4c, 0x74, 0xfe, 0x45, 0x3a, 0xb9, 0xfa, 0xa9, 0x3c, 0x4e, 0xba, 0xff,
	0x09, 0x00, 0x00, 0xff, 0xff, 0x21, 0xf8, 0x7d, 0xef, 0xf8, 0x00, 0x00, 0x00,
}