/*
Sniperkit-Bot
- Status: analyzed
*/

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/saver.proto

/*
Package tensorflow is a generated protocol buffer package.

It is generated from these files:
	tensorflow/core/protobuf/saver.proto

It has these top-level messages:
	SaverDef
*/
package tfpb

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

// A version number that identifies a different on-disk checkpoint format.
// Usually, each subclass of BaseSaverBuilder works with a particular
// version/format.  However, it is possible that the same builder may be
// upgraded to support a newer checkpoint format in the future.
type SaverDef_CheckpointFormatVersion int32

const (
	// Internal legacy format.
	SaverDef_LEGACY SaverDef_CheckpointFormatVersion = 0
	// Deprecated format: tf.Saver() which works with tensorflow::table::Table.
	SaverDef_V1 SaverDef_CheckpointFormatVersion = 1
	// Current format: more efficient.
	SaverDef_V2 SaverDef_CheckpointFormatVersion = 2
)

var SaverDef_CheckpointFormatVersion_name = map[int32]string{
	0: "LEGACY",
	1: "V1",
	2: "V2",
}
var SaverDef_CheckpointFormatVersion_value = map[string]int32{
	"LEGACY": 0,
	"V1":     1,
	"V2":     2,
}

func (x SaverDef_CheckpointFormatVersion) String() string {
	return proto.EnumName(SaverDef_CheckpointFormatVersion_name, int32(x))
}
func (SaverDef_CheckpointFormatVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{0, 0}
}

// Protocol buffer representing the configuration of a Saver.
type SaverDef struct {
	// The name of the tensor in which to specify the filename when saving or
	// restoring a model checkpoint.
	FilenameTensorName string `protobuf:"bytes,1,opt,name=filename_tensor_name,json=filenameTensorName" json:"filename_tensor_name,omitempty"`
	// The operation to run when saving a model checkpoint.
	SaveTensorName string `protobuf:"bytes,2,opt,name=save_tensor_name,json=saveTensorName" json:"save_tensor_name,omitempty"`
	// The operation to run when restoring a model checkpoint.
	RestoreOpName string `protobuf:"bytes,3,opt,name=restore_op_name,json=restoreOpName" json:"restore_op_name,omitempty"`
	// Maximum number of checkpoints to keep.  If 0, no checkpoints are deleted.
	MaxToKeep int32 `protobuf:"varint,4,opt,name=max_to_keep,json=maxToKeep" json:"max_to_keep,omitempty"`
	// Shard the save files, one per device that has Variable nodes.
	Sharded bool `protobuf:"varint,5,opt,name=sharded" json:"sharded,omitempty"`
	// How often to keep an additional checkpoint. If not specified, only the last
	// "max_to_keep" checkpoints are kept; if specified, in addition to keeping
	// the last "max_to_keep" checkpoints, an additional checkpoint will be kept
	// for every n hours of training.
	KeepCheckpointEveryNHours float32                          `protobuf:"fixed32,6,opt,name=keep_checkpoint_every_n_hours,json=keepCheckpointEveryNHours" json:"keep_checkpoint_every_n_hours,omitempty"`
	Version                   SaverDef_CheckpointFormatVersion `protobuf:"varint,7,opt,name=version,enum=tensorflow.SaverDef_CheckpointFormatVersion" json:"version,omitempty"`
}

func (m *SaverDef) Reset()                    { *m = SaverDef{} }
func (m *SaverDef) String() string            { return proto.CompactTextString(m) }
func (*SaverDef) ProtoMessage()               {}
func (*SaverDef) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *SaverDef) GetFilenameTensorName() string {
	if m != nil {
		return m.FilenameTensorName
	}
	return ""
}

func (m *SaverDef) GetSaveTensorName() string {
	if m != nil {
		return m.SaveTensorName
	}
	return ""
}

func (m *SaverDef) GetRestoreOpName() string {
	if m != nil {
		return m.RestoreOpName
	}
	return ""
}

func (m *SaverDef) GetMaxToKeep() int32 {
	if m != nil {
		return m.MaxToKeep
	}
	return 0
}

func (m *SaverDef) GetSharded() bool {
	if m != nil {
		return m.Sharded
	}
	return false
}

func (m *SaverDef) GetKeepCheckpointEveryNHours() float32 {
	if m != nil {
		return m.KeepCheckpointEveryNHours
	}
	return 0
}

func (m *SaverDef) GetVersion() SaverDef_CheckpointFormatVersion {
	if m != nil {
		return m.Version
	}
	return SaverDef_LEGACY
}

func init() {
	proto.RegisterType((*SaverDef)(nil), "tensorflow.SaverDef")
	proto.RegisterEnum("tensorflow.SaverDef_CheckpointFormatVersion", SaverDef_CheckpointFormatVersion_name, SaverDef_CheckpointFormatVersion_value)
}

func init() { proto.RegisterFile("tensorflow/core/protobuf/saver.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xd1, 0x41, 0x4f, 0xea, 0x40,
	0x10, 0x07, 0xf0, 0xb7, 0xe5, 0x51, 0x60, 0xc8, 0xe3, 0x35, 0xab, 0x89, 0xf5, 0xa0, 0x69, 0x88,
	0xd1, 0x1e, 0x4c, 0x51, 0x8c, 0x77, 0x05, 0x41, 0x13, 0x0d, 0x92, 0x4a, 0x48, 0x3c, 0x6d, 0x0a,
	0x4c, 0xa5, 0x81, 0x76, 0x9b, 0xed, 0x82, 0xf8, 0x11, 0xfc, 0xc6, 0x1e, 0xcd, 0x2e, 0x56, 0xf0,
	0xe0, 0xa9, 0xdd, 0x9d, 0xdf, 0x7f, 0x27, 0x99, 0x81, 0x23, 0x89, 0x49, 0xc6, 0x45, 0x38, 0xe7,
	0xaf, 0x8d, 0x31, 0x17, 0xd8, 0x48, 0x05, 0x97, 0x7c, 0xb4, 0x08, 0x1b, 0x59, 0xb0, 0x44, 0xe1,
	0xe9, 0x23, 0x85, 0x8d, 0xaa, 0xbf, 0x17, 0xa0, 0xfc, 0xa4, 0x6a, 0x37, 0x18, 0xd2, 0x33, 0xd8,
	0x0d, 0xa3, 0x39, 0x26, 0x41, 0x8c, 0x6c, 0x6d, 0x98, 0xfa, 0xb7, 0x89, 0x43, 0xdc, 0x8a, 0x4f,
	0xf3, 0xda, 0x40, 0x97, 0x7a, 0x41, 0x8c, 0xd4, 0x05, 0x4b, 0xbd, 0xfc, 0x43, 0x1b, 0x5a, 0xd7,
	0xd4, 0xfd, 0x96, 0x3c, 0x86, 0xff, 0x02, 0x33, 0xc9, 0x05, 0x32, 0x9e, 0xae, 0x61, 0x41, 0xc3,
	0x7f, 0x5f, 0xd7, 0x8f, 0xa9, 0x76, 0x87, 0x50, 0x8d, 0x83, 0x15, 0x93, 0x9c, 0xcd, 0x10, 0x53,
	0xfb, 0xaf, 0x43, 0xdc, 0xa2, 0x5f, 0x89, 0x83, 0xd5, 0x80, 0xdf, 0x23, 0xa6, 0xd4, 0x86, 0x52,
	0x36, 0x0d, 0xc4, 0x04, 0x27, 0x76, 0xd1, 0x21, 0x6e, 0xd9, 0xcf, 0x8f, 0xf4, 0x0a, 0x0e, 0x54,
	0x84, 0x8d, 0xa7, 0x38, 0x9e, 0xa5, 0x3c, 0x4a, 0x24, 0xc3, 0x25, 0x8a, 0x37, 0x96, 0xb0, 0x29,
	0x5f, 0x88, 0xcc, 0x36, 0x1d, 0xe2, 0x1a, 0xfe, 0xbe, 0x42, 0xed, 0x6f, 0xd3, 0x51, 0xa4, 0x77,
	0xa7, 0x00, 0xed, 0x42, 0x69, 0x89, 0x22, 0x8b, 0x78, 0x62, 0x97, 0x1c, 0xe2, 0xd6, 0x9a, 0xa7,
	0xde, 0x66, 0x54, 0x5e, 0x3e, 0x26, 0x6f, 0x13, 0xee, 0x72, 0x11, 0x07, 0x72, 0xb8, 0xce, 0xf8,
	0x79, 0xb8, 0x7e, 0x09, 0x7b, 0xbf, 0x18, 0x0a, 0x60, 0x3e, 0x74, 0x6e, 0xaf, 0xdb, 0xcf, 0xd6,
	0x1f, 0x6a, 0x82, 0x31, 0x3c, 0xb7, 0x88, 0xfe, 0x36, 0x2d, 0xa3, 0x75, 0x02, 0x3b, 0x5c, 0xbc,
	0x6c, 0xb7, 0x5c, 0xc8, 0x68, 0xde, 0xaa, 0xea, 0xc6, 0x7d, 0xb5, 0xba, 0xac, 0x4f, 0x3e, 0x08,
	0x19, 0x99, 0x7a, 0x8f, 0x17, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xa1, 0x8e, 0x12, 0xef,
	0x01, 0x00, 0x00,
}
