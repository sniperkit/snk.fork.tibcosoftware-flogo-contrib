/*
Sniperkit-Bot
- Status: analyzed
*/

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/node_def.proto

/*
Package tensorflow is a generated protocol buffer package.

It is generated from these files:
	tensorflow/core/framework/node_def.proto

It has these top-level messages:
	NodeDef
*/
package tffw

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NodeDef struct {
	// The name given to this operator. Used for naming inputs,
	// logging, visualization, etc.  Unique within a single GraphDef.
	// Must match the regexp "[A-Za-z0-9.][A-Za-z0-9_./]*".
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The operation name.  There may be custom parameters in attrs.
	// Op names starting with an underscore are reserved for internal use.
	Op string `protobuf:"bytes,2,opt,name=op" json:"op,omitempty"`
	// Each input is "node:src_output" with "node" being a string name and
	// "src_output" indicating which output tensor to use from "node". If
	// "src_output" is 0 the ":0" suffix can be omitted.  Regular inputs
	// may optionally be followed by control inputs that have the format
	// "^node".
	Input []string `protobuf:"bytes,3,rep,name=input" json:"input,omitempty"`
	// A (possibly partial) specification for the device on which this
	// node should be placed.
	// The expected syntax for this string is as follows:
	//
	// DEVICE_SPEC ::= PARTIAL_SPEC
	//
	// PARTIAL_SPEC ::= ("/" CONSTRAINT) *
	// CONSTRAINT ::= ("job:" JOB_NAME)
	//              | ("replica:" [1-9][0-9]*)
	//              | ("task:" [1-9][0-9]*)
	//              | ( ("gpu" | "cpu") ":" ([1-9][0-9]* | "*") )
	//
	// Valid values for this string include:
	// * "/job:worker/replica:0/task:1/gpu:3"  (full specification)
	// * "/job:worker/gpu:3"                   (partial specification)
	// * ""                                    (no specification)
	//
	// If the constraints do not resolve to a single device (or if this
	// field is empty or not present), the runtime will attempt to
	// choose a device automatically.
	Device string `protobuf:"bytes,4,opt,name=device" json:"device,omitempty"`
	// Operation-specific graph-construction-time configuration.
	// Note that this should include all attrs defined in the
	// corresponding OpDef, including those with a value matching
	// the default -- this allows the default to change and makes
	// NodeDefs easier to interpret on their own.  However, if
	// an attr with a default is not specified in this list, the
	// default will be used.
	// The "names" (keys) must match the regexp "[a-z][a-z0-9_]+" (and
	// one of the names from the corresponding OpDef's attr field).
	// The values must have a type matching the corresponding OpDef
	// attr's type field.
	// TODO(josh11b): Add some examples here showing best practices.
	Attr map[string]*AttrValue `protobuf:"bytes,5,rep,name=attr" json:"attr,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *NodeDef) Reset()                    { *m = NodeDef{} }
func (m *NodeDef) String() string            { return proto.CompactTextString(m) }
func (*NodeDef) ProtoMessage()               {}
func (*NodeDef) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *NodeDef) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NodeDef) GetOp() string {
	if m != nil {
		return m.Op
	}
	return ""
}

func (m *NodeDef) GetInput() []string {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *NodeDef) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *NodeDef) GetAttr() map[string]*AttrValue {
	if m != nil {
		return m.Attr
	}
	return nil
}

func init() {
	proto.RegisterType((*NodeDef)(nil), "tensorflow.NodeDef")
}

func init() { proto.RegisterFile("tensorflow/core/framework/node_def.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xd9, 0xa4, 0xa9, 0x64, 0x0a, 0x22, 0x83, 0xca, 0x52, 0x10, 0x82, 0xa7, 0x50, 0x21,
	0xc1, 0x7a, 0x11, 0x6f, 0x16, 0xbd, 0x96, 0x92, 0x83, 0xd7, 0x12, 0x9b, 0x89, 0x94, 0xb6, 0x3b,
	0x61, 0xdc, 0xb6, 0xf4, 0x65, 0x7d, 0x0e, 0x8f, 0xb2, 0x9b, 0xd0, 0xf6, 0xe2, 0x6d, 0x66, 0xf6,
	0x9b, 0x9d, 0x8f, 0x1f, 0x52, 0x4b, 0xe6, 0x9b, 0xa5, 0x5e, 0xf3, 0x3e, 0x5f, 0xb0, 0x50, 0x5e,
	0x4b, 0xb9, 0xa1, 0x3d, 0xcb, 0x2a, 0x37, 0x5c, 0xd1, 0xbc, 0xa2, 0x3a, 0x6b, 0x84, 0x2d, 0x23,
	0x9c, 0xc8, 0xe1, 0xe8, 0xff, 0xad, 0xd2, 0x5a, 0x99, 0xef, 0xca, 0xf5, 0x96, 0xda, 0xbd, 0xfb,
	0x1f, 0x05, 0x17, 0x53, 0xae, 0xe8, 0x8d, 0x6a, 0x44, 0xe8, 0x99, 0x72, 0x43, 0x5a, 0x25, 0x2a,
	0x8d, 0x0b, 0x5f, 0xe3, 0x25, 0x04, 0xdc, 0xe8, 0xc0, 0x4f, 0x02, 0x6e, 0xf0, 0x1a, 0xa2, 0xa5,
	0x69, 0xb6, 0x56, 0x87, 0x49, 0x98, 0xc6, 0x45, 0xdb, 0xe0, 0x2d, 0xf4, 0x2b, 0xda, 0x2d, 0x17,
	0xa4, 0x7b, 0x9e, 0xec, 0x3a, 0x7c, 0x84, 0x9e, 0xbb, 0xa8, 0xa3, 0x24, 0x4c, 0x07, 0xe3, 0xbb,
	0xec, 0x24, 0x96, 0x75, 0x47, 0xb3, 0x57, 0x6b, 0xe5, 0xdd, 0x58, 0x39, 0x14, 0x1e, 0x1d, 0x4e,
	0x21, 0x3e, 0x8e, 0xf0, 0x0a, 0xc2, 0x15, 0x1d, 0x3a, 0x21, 0x57, 0xe2, 0x03, 0x44, 0x5e, 0xdf,
	0x2b, 0x0d, 0xc6, 0x37, 0xe7, 0x5f, 0xba, 0xbd, 0x0f, 0xf7, 0x58, 0xb4, 0xcc, 0x4b, 0xf0, 0xac,
	0x26, 0x23, 0xd0, 0x2c, 0x5f, 0xe7, 0xd8, 0x31, 0x8d, 0x49, 0xec, 0x24, 0x66, 0x2e, 0x87, 0x99,
	0xfa, 0x55, 0xea, 0xb3, 0xef, 0x33, 0x79, 0xfa, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x56, 0x19, 0xbb,
	0x7c, 0x77, 0x01, 0x00, 0x00,
}
