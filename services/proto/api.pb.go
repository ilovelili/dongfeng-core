// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/ilovelili/dongfeng-core/services/proto/api.proto

package dongfeng_svc_core_server

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/ilovelili/dongfeng-protobuf"
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

func init() {
	proto.RegisterFile("github.com/ilovelili/dongfeng-core/services/proto/api.proto", fileDescriptor_704a4f15475a1c86)
}

var fileDescriptor_704a4f15475a1c86 = []byte{
	// 854 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x97, 0x5b, 0x4f, 0xd4, 0x40,
	0x14, 0xc7, 0x25, 0x46, 0x23, 0x03, 0x78, 0x19, 0x5f, 0x14, 0x1f, 0x14, 0x50, 0x40, 0x81, 0x6e,
	0xc2, 0xe5, 0xc9, 0x27, 0x05, 0xb3, 0x92, 0x00, 0x6e, 0xb8, 0x08, 0x46, 0x03, 0xb6, 0xdd, 0xb3,
	0xdd, 0x89, 0xdd, 0x4e, 0xe9, 0xcc, 0x6e, 0xf4, 0x0b, 0x98, 0x98, 0xf8, 0x29, 0xfd, 0x24, 0xa6,
	0xd3, 0x9d, 0xd9, 0xe9, 0x65, 0xda, 0x6e, 0xe2, 0xfa, 0x42, 0xc8, 0xfe, 0x7f, 0xf3, 0xff, 0x9f,
	0x69, 0x4f, 0x4f, 0xa7, 0xe8, 0xb5, 0x47, 0x78, 0xb7, 0xef, 0x58, 0x2e, 0xed, 0x35, 0x88, 0x4f,
	0x07, 0xe0, 0x13, 0x9f, 0x34, 0xda, 0x34, 0xf0, 0x3a, 0x10, 0x78, 0x1b, 0x2e, 0x8d, 0xa0, 0xc1,
	0x20, 0x1a, 0x10, 0x17, 0x58, 0x23, 0x8c, 0x28, 0xa7, 0x0d, 0x3b, 0x24, 0x96, 0xf8, 0x0f, 0x3f,
	0x92, 0x9c, 0xc5, 0x06, 0xae, 0x15, 0xb3, 0x56, 0xcc, 0x42, 0x34, 0xbf, 0x53, 0x6e, 0x2b, 0x96,
	0x3b, 0xfd, 0x8e, 0xfa, 0x25, 0x31, 0xdc, 0xfc, 0xb3, 0x80, 0x6e, 0xda, 0x21, 0xc1, 0x07, 0xe8,
	0xd6, 0x01, 0xf5, 0x48, 0x80, 0x9f, 0x5a, 0x69, 0xc2, 0xe9, 0x77, 0x2c, 0xa1, 0x1c, 0xc3, 0x75,
	0x1f, 0x18, 0x9f, 0x7f, 0x66, 0x06, 0x58, 0x48, 0x03, 0x06, 0x8b, 0x37, 0xf0, 0x05, 0x9a, 0xde,
	0xb3, 0x59, 0xd7, 0xa1, 0x76, 0xd4, 0xc6, 0x4b, 0x05, 0x0b, 0x94, 0x2a, 0x5d, 0x9f, 0x97, 0x43,
	0xca, 0xf9, 0x33, 0x42, 0x67, 0x61, 0xdb, 0xe6, 0x70, 0xc6, 0x20, 0xc2, 0x45, 0xab, 0x46, 0xb2,
	0xf4, 0x7e, 0x51, 0x41, 0x29, 0xf3, 0x73, 0x84, 0x9a, 0xc0, 0x77, 0x7d, 0x9b, 0x31, 0x60, 0x78,
	0xb1, 0x60, 0x99, 0x94, 0xa5, 0xf5, 0x52, 0x29, 0xa3, 0x8c, 0xbf, 0xa2, 0xb9, 0x24, 0x50, 0x7a,
	0x9b, 0x4b, 0x4a, 0xd9, 0x2f, 0x57, 0x61, 0x2a, 0xe1, 0x23, 0x9a, 0x6e, 0x02, 0x6f, 0xf5, 0x43,
	0xe2, 0x1b, 0x2b, 0x17, 0x6a, 0x45, 0xe5, 0x43, 0x46, 0xf9, 0x5e, 0xa2, 0x99, 0x24, 0x50, 0x08,
	0x25, 0x75, 0xa7, 0xcc, 0x97, 0xab, 0x30, 0xe5, 0x7f, 0x85, 0x66, 0x35, 0x81, 0xfd, 0xfb, 0x80,
	0x2f, 0x68, 0xa6, 0x09, 0xfc, 0x14, 0x6c, 0xb7, 0x0b, 0x11, 0x2b, 0xec, 0x98, 0x91, 0x5e, 0xd6,
	0x31, 0x3a, 0xa5, 0xdc, 0xdb, 0xf2, 0xc6, 0x0e, 0x25, 0xbc, 0x62, 0x2c, 0x2c, 0x13, 0xb1, 0x5a,
	0x0d, 0xaa, 0x14, 0x40, 0x77, 0x53, 0x12, 0x9b, 0x58, 0x4c, 0x13, 0xf8, 0x1b, 0xce, 0x21, 0x68,
	0xdb, 0x81, 0x0b, 0xc5, 0x31, 0x29, 0xa4, 0x2c, 0x26, 0x03, 0xaa, 0x98, 0x1e, 0xba, 0x9f, 0x54,
	0x30, 0x52, 0xf1, 0x2b, 0x63, 0x99, 0xf9, 0xac, 0xb5, 0x5a, 0xac, 0x8a, 0x0b, 0xd0, 0x83, 0xac,
	0xca, 0x26, 0x99, 0x77, 0x85, 0x66, 0xe3, 0xe7, 0xa8, 0xfb, 0x83, 0x91, 0x78, 0x3d, 0x36, 0xf4,
	0x92, 0x04, 0xca, 0x3a, 0x3a, 0x85, 0xa9, 0x00, 0x4f, 0x76, 0x83, 0xd4, 0xb0, 0xf9, 0x26, 0x67,
	0x53, 0x5e, 0xd6, 0x20, 0x55, 0x50, 0x17, 0xdd, 0x4b, 0x6b, 0x6c, 0x52, 0x49, 0xbf, 0xa6, 0xd0,
	0x7c, 0xdc, 0x2e, 0x1e, 0xbc, 0x07, 0xe2, 0x75, 0xf9, 0xb9, 0xf8, 0xdb, 0x3a, 0xb4, 0x19, 0x8f,
	0xbb, 0x7d, 0xcb, 0xd0, 0x5d, 0x85, 0xb8, 0x2c, 0x60, 0x7b, 0xbc, 0x45, 0xaa, 0x96, 0xdf, 0x53,
	0xe8, 0x49, 0x9e, 0x3b, 0xd9, 0x93, 0xc5, 0xd4, 0xf3, 0x95, 0xbc, 0xac, 0x66, 0x67, 0xcc, 0x55,
	0xaa, 0x1c, 0x07, 0xcd, 0x35, 0x81, 0xbf, 0x3d, 0xdc, 0x97, 0xf9, 0x86, 0x46, 0x51, 0x84, 0x4c,
	0x5c, 0xa9, 0xe4, 0x54, 0xc6, 0xcf, 0x29, 0xf4, 0xb8, 0x09, 0x3c, 0xa9, 0xe4, 0x94, 0x66, 0xae,
	0xfe, 0x66, 0xb1, 0x51, 0x21, 0x2d, 0xc3, 0xb7, 0xc6, 0x5a, 0x93, 0xed, 0x83, 0x34, 0x36, 0xba,
	0xf4, 0xb5, 0x5c, 0xb3, 0x57, 0x7e, 0x7b, 0xbc, 0x45, 0xaa, 0x96, 0x13, 0x74, 0xa7, 0x09, 0xfc,
	0x10, 0x82, 0x3e, 0xc3, 0x0b, 0xc5, 0x1e, 0xb1, 0x28, 0x63, 0x16, 0xcb, 0x10, 0x65, 0xfa, 0x49,
	0x9c, 0x30, 0x8e, 0xc1, 0x25, 0x21, 0x30, 0x6c, 0x78, 0x07, 0x27, 0x72, 0xd9, 0xc9, 0x48, 0x83,
	0x32, 0xd3, 0x7b, 0x3f, 0xf0, 0x22, 0x68, 0x13, 0x08, 0xb8, 0x71, 0x7a, 0x8f, 0x90, 0x8a, 0xe9,
	0xad, 0x83, 0x2a, 0xe6, 0x1a, 0xe1, 0x94, 0x74, 0x64, 0xf7, 0x80, 0xe1, 0xb5, 0x2a, 0x87, 0x18,
	0x93, 0x71, 0xeb, 0xf5, 0xe0, 0xfc, 0x04, 0xd7, 0x37, 0x67, 0x9e, 0xe0, 0xf9, 0xfd, 0xad, 0xd5,
	0x62, 0xf5, 0xb9, 0x17, 0x4f, 0xde, 0x88, 0xba, 0xfd, 0x08, 0x7a, 0x22, 0xcd, 0x70, 0x85, 0x34,
	0xa6, 0x6c, 0xee, 0x65, 0x49, 0x95, 0x14, 0xca, 0x9d, 0x69, 0x32, 0x36, 0x57, 0x5b, 0x10, 0xb7,
	0x5e, 0x0f, 0xd6, 0xcf, 0xcf, 0x49, 0x35, 0x1d, 0xe2, 0x83, 0xe9, 0x34, 0x34, 0x94, 0x2b, 0x4e,
	0x43, 0x8a, 0x52, 0xe6, 0x4c, 0xb4, 0x60, 0x2b, 0x82, 0x81, 0x0c, 0xb0, 0x4c, 0x4b, 0x61, 0xf0,
	0x21, 0x3a, 0x82, 0xef, 0xd9, 0xa8, 0x46, 0x6d, 0x3e, 0x13, 0xaa, 0x69, 0xff, 0x23, 0xf4, 0x52,
	0x9c, 0x2a, 0x87, 0xbf, 0x9b, 0xdf, 0xf1, 0x43, 0xbd, 0xea, 0x1d, 0xaf, 0x30, 0xfd, 0x5c, 0xb9,
	0x1b, 0x41, 0x72, 0x17, 0xc5, 0x9e, 0x56, 0xca, 0xee, 0xb3, 0xbe, 0x99, 0xd5, 0x6a, 0x30, 0x7f,
	0x7a, 0x9d, 0x74, 0xca, 0x1e, 0xf8, 0x30, 0xe1, 0x94, 0x0b, 0xf1, 0x01, 0xf4, 0xce, 0xa1, 0xf4,
	0x9b, 0x71, 0xb0, 0x26, 0x6a, 0xc5, 0x60, 0x95, 0x50, 0xfe, 0x13, 0x48, 0x28, 0x25, 0x5f, 0x28,
	0x42, 0xaf, 0xfe, 0x42, 0x19, 0x62, 0x5a, 0x03, 0xe3, 0x44, 0x38, 0xa2, 0x9c, 0x74, 0x88, 0x6b,
	0x73, 0x42, 0x03, 0x6c, 0x7e, 0xb0, 0x75, 0x4c, 0xa6, 0x6d, 0xd4, 0xa4, 0x55, 0xe8, 0x00, 0x3d,
	0xcc, 0xeb, 0x0c, 0xd7, 0xf3, 0x51, 0x97, 0xd0, 0xaa, 0x8b, 0xeb, 0x9b, 0x8d, 0x8f, 0x3d, 0xae,
	0x0b, 0x8c, 0x11, 0xc7, 0x87, 0x96, 0xcd, 0xbb, 0x0c, 0x1b, 0xde, 0x08, 0x19, 0xac, 0x6c, 0xb3,
	0x45, 0xb4, 0x0c, 0x75, 0x6e, 0x0b, 0x6c, 0xeb, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xa0,
	0x6d, 0x29, 0x7b, 0x11, 0x00, 0x00,
}
