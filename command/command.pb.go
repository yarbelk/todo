// Code generated by protoc-gen-go.
// source: github.com/yarbelk/todo/command/command.proto
// DO NOT EDIT!

/*
Package command is a generated protocol buffer package.

It is generated from these files:
	github.com/yarbelk/todo/command/command.proto

It has these top-level messages:
	TaskCreate
	TaskCreated
*/
package command

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type TaskCreate struct {
	Description string `protobuf:"bytes,1,opt,name=description" json:"description,omitempty"`
}

func (m *TaskCreate) Reset()                    { *m = TaskCreate{} }
func (m *TaskCreate) String() string            { return proto.CompactTextString(m) }
func (*TaskCreate) ProtoMessage()               {}
func (*TaskCreate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type TaskCreated struct {
	Id          string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
}

func (m *TaskCreated) Reset()                    { *m = TaskCreated{} }
func (m *TaskCreated) String() string            { return proto.CompactTextString(m) }
func (*TaskCreated) ProtoMessage()               {}
func (*TaskCreated) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*TaskCreate)(nil), "TaskCreate")
	proto.RegisterType((*TaskCreated)(nil), "TaskCreated")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Commander service

type CommanderClient interface {
	NewTask(ctx context.Context, in *TaskCreate, opts ...client.CallOption) (*TaskCreated, error)
}

type commanderClient struct {
	c           client.Client
	serviceName string
}

func NewCommanderClient(serviceName string, c client.Client) CommanderClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "commander"
	}
	return &commanderClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *commanderClient) NewTask(ctx context.Context, in *TaskCreate, opts ...client.CallOption) (*TaskCreated, error) {
	req := c.c.NewRequest(c.serviceName, "Commander.NewTask", in)
	out := new(TaskCreated)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Commander service

type CommanderHandler interface {
	NewTask(context.Context, *TaskCreate, *TaskCreated) error
}

func RegisterCommanderHandler(s server.Server, hdlr CommanderHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Commander{hdlr}, opts...))
}

type Commander struct {
	CommanderHandler
}

func (h *Commander) NewTask(ctx context.Context, in *TaskCreate, out *TaskCreated) error {
	return h.CommanderHandler.NewTask(ctx, in, out)
}

func init() { proto.RegisterFile("github.com/yarbelk/todo/command/command.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xaf, 0x4c, 0x2c, 0x4a, 0x4a, 0xcd, 0xc9, 0xd6, 0x2f,
	0xc9, 0x4f, 0xc9, 0xd7, 0x4f, 0xce, 0xcf, 0xcd, 0x4d, 0xcc, 0x4b, 0x81, 0xd1, 0x7a, 0x05, 0x45,
	0xf9, 0x25, 0xf9, 0x4a, 0x7a, 0x5c, 0x5c, 0x21, 0x89, 0xc5, 0xd9, 0xce, 0x45, 0xa9, 0x89, 0x25,
	0xa9, 0x42, 0x0a, 0x5c, 0xdc, 0x29, 0xa9, 0xc5, 0xc9, 0x45, 0x99, 0x05, 0x25, 0x99, 0xf9, 0x79,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0xc8, 0x42, 0x4a, 0xf6, 0x5c, 0xdc, 0x08, 0xf5, 0x29,
	0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x4c, 0x60, 0x75, 0x4c, 0x99, 0x29, 0xe8, 0x06, 0x30,
	0x63, 0x18, 0x60, 0x64, 0xcc, 0xc5, 0xe9, 0x0c, 0x71, 0x41, 0x6a, 0x91, 0x90, 0x1a, 0x17, 0xbb,
	0x5f, 0x6a, 0x39, 0xc8, 0x40, 0x21, 0x6e, 0x3d, 0x84, 0xb9, 0x52, 0x3c, 0x48, 0x9c, 0x14, 0x25,
	0x86, 0x24, 0x36, 0xb0, 0x63, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x9c, 0xe2, 0x5b,
	0xdd, 0x00, 0x00, 0x00,
}
