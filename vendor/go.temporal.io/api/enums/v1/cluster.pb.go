// The MIT License
//
// Copyright (c) 2022 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/api/enums/v1/cluster.proto

package enums

import (
	fmt "fmt"
	math "math"
	strconv "strconv"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ClusterMemberRole int32

const (
	CLUSTER_MEMBER_ROLE_UNSPECIFIED ClusterMemberRole = 0
	CLUSTER_MEMBER_ROLE_FRONTEND    ClusterMemberRole = 1
	CLUSTER_MEMBER_ROLE_HISTORY     ClusterMemberRole = 2
	CLUSTER_MEMBER_ROLE_MATCHING    ClusterMemberRole = 3
	CLUSTER_MEMBER_ROLE_WORKER      ClusterMemberRole = 4
)

var ClusterMemberRole_name = map[int32]string{
	0: "Unspecified",
	1: "Frontend",
	2: "History",
	3: "Matching",
	4: "Worker",
}

var ClusterMemberRole_value = map[string]int32{
	"Unspecified": 0,
	"Frontend":    1,
	"History":     2,
	"Matching":    3,
	"Worker":      4,
}

func (ClusterMemberRole) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_907d202b9a74ef99, []int{0}
}

func init() {
	proto.RegisterEnum("temporal.api.enums.v1.ClusterMemberRole", ClusterMemberRole_name, ClusterMemberRole_value)
}

func init() {
	proto.RegisterFile("temporal/api/enums/v1/cluster.proto", fileDescriptor_907d202b9a74ef99)
}

var fileDescriptor_907d202b9a74ef99 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0x49, 0xcd, 0x2d,
	0xc8, 0x2f, 0x4a, 0xcc, 0xd1, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcd, 0x2b, 0xcd, 0x2d, 0xd6, 0x2f,
	0x33, 0xd4, 0x4f, 0xce, 0x29, 0x2d, 0x2e, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0x85, 0x29, 0xd2, 0x4b, 0x2c, 0xc8, 0xd4, 0x03, 0x2b, 0xd2, 0x2b, 0x33, 0xd4, 0xda, 0xcb,
	0xc8, 0x25, 0xe8, 0x0c, 0x51, 0xe8, 0x9b, 0x9a, 0x9b, 0x94, 0x5a, 0x14, 0x94, 0x9f, 0x93, 0x2a,
	0xa4, 0xcc, 0x25, 0xef, 0xec, 0x13, 0x1a, 0x1c, 0xe2, 0x1a, 0x14, 0xef, 0xeb, 0xea, 0xeb, 0xe4,
	0x1a, 0x14, 0x1f, 0xe4, 0xef, 0xe3, 0x1a, 0x1f, 0xea, 0x17, 0x1c, 0xe0, 0xea, 0xec, 0xe9, 0xe6,
	0xe9, 0xea, 0x22, 0xc0, 0x20, 0xa4, 0xc0, 0x25, 0x83, 0x4d, 0x91, 0x5b, 0x90, 0xbf, 0x5f, 0x88,
	0xab, 0x9f, 0x8b, 0x00, 0xa3, 0x90, 0x3c, 0x97, 0x34, 0x36, 0x15, 0x1e, 0x9e, 0xc1, 0x21, 0xfe,
	0x41, 0x91, 0x02, 0x4c, 0xb8, 0x8c, 0xf0, 0x75, 0x0c, 0x71, 0xf6, 0xf0, 0xf4, 0x73, 0x17, 0x60,
	0x16, 0x92, 0xe3, 0x92, 0xc2, 0xa6, 0x22, 0xdc, 0x3f, 0xc8, 0xdb, 0x35, 0x48, 0x80, 0xc5, 0x69,
	0x0b, 0xe3, 0x85, 0x87, 0x72, 0x0c, 0x37, 0x1e, 0xca, 0x31, 0x7c, 0x78, 0x28, 0xc7, 0xd8, 0xf0,
	0x48, 0x8e, 0x71, 0xc5, 0x23, 0x39, 0xc6, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c,
	0xf0, 0x48, 0x8e, 0xf1, 0xc5, 0x23, 0x39, 0x86, 0x0f, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63,
	0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x06, 0x2e, 0x89, 0xcc, 0x7c, 0x3d, 0xac, 0x01,
	0xe2, 0xc4, 0x03, 0x0d, 0x8d, 0x00, 0x50, 0xa8, 0x05, 0x30, 0x46, 0x29, 0xa6, 0x23, 0xa9, 0xcc,
	0xcc, 0x47, 0x09, 0x62, 0x6b, 0x30, 0x63, 0x15, 0x93, 0x68, 0x08, 0x4c, 0x81, 0x63, 0x41, 0xa6,
	0x9e, 0x2b, 0xd8, 0xa8, 0x30, 0xc3, 0x57, 0x4c, 0x12, 0x30, 0x71, 0x2b, 0x2b, 0xc7, 0x82, 0x4c,
	0x2b, 0x2b, 0xb0, 0x8c, 0x95, 0x55, 0x98, 0x61, 0x12, 0x1b, 0x38, 0x52, 0x8c, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xf2, 0x4c, 0xd9, 0x3a, 0xbb, 0x01, 0x00, 0x00,
}

func (x ClusterMemberRole) String() string {
	s, ok := ClusterMemberRole_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
