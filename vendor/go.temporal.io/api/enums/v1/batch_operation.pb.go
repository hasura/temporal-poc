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
// source: temporal/api/enums/v1/batch_operation.proto

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

type BatchOperationType int32

const (
	BATCH_OPERATION_TYPE_UNSPECIFIED BatchOperationType = 0
	BATCH_OPERATION_TYPE_TERMINATE   BatchOperationType = 1
	BATCH_OPERATION_TYPE_CANCEL      BatchOperationType = 2
	BATCH_OPERATION_TYPE_SIGNAL      BatchOperationType = 3
)

var BatchOperationType_name = map[int32]string{
	0: "Unspecified",
	1: "Terminate",
	2: "Cancel",
	3: "Signal",
}

var BatchOperationType_value = map[string]int32{
	"Unspecified": 0,
	"Terminate":   1,
	"Cancel":      2,
	"Signal":      3,
}

func (BatchOperationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3b55ebc4deb738bf, []int{0}
}

type BatchOperationState int32

const (
	BATCH_OPERATION_STATE_UNSPECIFIED BatchOperationState = 0
	BATCH_OPERATION_STATE_RUNNING     BatchOperationState = 1
	BATCH_OPERATION_STATE_COMPLETED   BatchOperationState = 2
	BATCH_OPERATION_STATE_FAILED      BatchOperationState = 3
)

var BatchOperationState_name = map[int32]string{
	0: "Unspecified",
	1: "Running",
	2: "Completed",
	3: "Failed",
}

var BatchOperationState_value = map[string]int32{
	"Unspecified": 0,
	"Running":     1,
	"Completed":   2,
	"Failed":      3,
}

func (BatchOperationState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3b55ebc4deb738bf, []int{1}
}

func init() {
	proto.RegisterEnum("temporal.api.enums.v1.BatchOperationType", BatchOperationType_name, BatchOperationType_value)
	proto.RegisterEnum("temporal.api.enums.v1.BatchOperationState", BatchOperationState_name, BatchOperationState_value)
}

func init() {
	proto.RegisterFile("temporal/api/enums/v1/batch_operation.proto", fileDescriptor_3b55ebc4deb738bf)
}

var fileDescriptor_3b55ebc4deb738bf = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xd2, 0x4f, 0x4b, 0xf3, 0x30,
	0x1c, 0x07, 0xf0, 0xa6, 0x83, 0xe7, 0x90, 0x53, 0xc9, 0xc3, 0x60, 0xf0, 0x3c, 0x66, 0x9b, 0x7f,
	0x2e, 0x13, 0x5a, 0x8a, 0xb7, 0x7a, 0x4a, 0xbb, 0x6c, 0x16, 0xba, 0xb4, 0xac, 0xd9, 0x40, 0x2f,
	0xa5, 0x93, 0xa2, 0x05, 0xb7, 0x86, 0x59, 0x07, 0xde, 0x7c, 0x09, 0xbe, 0x04, 0x4f, 0x22, 0xbe,
	0x05, 0xdf, 0x80, 0xc7, 0x1d, 0x77, 0x74, 0xdd, 0x45, 0x3c, 0xed, 0x25, 0xc8, 0x0a, 0x05, 0x37,
	0xea, 0x2d, 0x24, 0x9f, 0x6f, 0xf8, 0x85, 0x6f, 0xe0, 0x71, 0x1a, 0x8d, 0x45, 0x32, 0x0d, 0x6f,
	0xb4, 0x50, 0xc4, 0x5a, 0x34, 0xb9, 0x1b, 0xdf, 0x6a, 0x33, 0x5d, 0x1b, 0x85, 0xe9, 0xe5, 0x75,
	0x90, 0x88, 0x68, 0x1a, 0xa6, 0x71, 0x32, 0x51, 0xc5, 0x34, 0x49, 0x13, 0x54, 0x2d, 0xb0, 0x1a,
	0x8a, 0x58, 0xcd, 0xb1, 0x3a, 0xd3, 0x5b, 0x4f, 0x00, 0x22, 0x73, 0x13, 0x70, 0x0b, 0xcf, 0xef,
	0x45, 0x84, 0x0e, 0x61, 0xc3, 0x24, 0xdc, 0x3a, 0x0b, 0x5c, 0x8f, 0xf6, 0x09, 0xb7, 0x5d, 0x16,
	0xf0, 0x73, 0x8f, 0x06, 0x03, 0xe6, 0x7b, 0xd4, 0xb2, 0x3b, 0x36, 0x6d, 0x2b, 0x12, 0xda, 0x87,
	0xb8, 0x54, 0x71, 0xda, 0xef, 0xd9, 0x8c, 0x70, 0xaa, 0x00, 0x54, 0x87, 0xff, 0x4a, 0x8d, 0x45,
	0x98, 0x45, 0x1d, 0x45, 0xfe, 0x15, 0xf8, 0x76, 0x97, 0x11, 0x47, 0xa9, 0xb4, 0x9e, 0x01, 0xfc,
	0xbb, 0x3d, 0xa2, 0x9f, 0x86, 0x69, 0x84, 0x8e, 0x60, 0x73, 0x37, 0xe8, 0x73, 0xc2, 0x77, 0x87,
	0x6c, 0xc2, 0xbd, 0x72, 0xd6, 0x1f, 0x30, 0x66, 0xb3, 0xae, 0x02, 0xd0, 0x01, 0xac, 0x97, 0x13,
	0xcb, 0xed, 0x79, 0x0e, 0xe5, 0xb4, 0xad, 0xc8, 0xa8, 0x01, 0xff, 0x97, 0xa3, 0x0e, 0xb1, 0x1d,
	0xda, 0x56, 0x2a, 0xe6, 0x1b, 0x98, 0x2f, 0xb1, 0xb4, 0x58, 0x62, 0x69, 0xbd, 0xc4, 0xe0, 0x21,
	0xc3, 0xe0, 0x25, 0xc3, 0xe0, 0x3d, 0xc3, 0x60, 0x9e, 0x61, 0xf0, 0x91, 0x61, 0xf0, 0x99, 0x61,
	0x69, 0x9d, 0x61, 0xf0, 0xb8, 0xc2, 0xd2, 0x7c, 0x85, 0xa5, 0xc5, 0x0a, 0x4b, 0xb0, 0x16, 0x27,
	0x6a, 0x69, 0x39, 0xe6, 0xce, 0xb3, 0xbd, 0x4d, 0x91, 0x1e, 0xb8, 0x68, 0x5e, 0xfd, 0x08, 0xc4,
	0xc9, 0x56, 0xfb, 0xa7, 0xf9, 0xe2, 0x55, 0xae, 0xf2, 0x02, 0x10, 0x11, 0xab, 0x34, 0xbf, 0x71,
	0xa8, 0x7f, 0xc9, 0xb5, 0x62, 0xdf, 0x30, 0x88, 0x88, 0x0d, 0x23, 0x3f, 0x31, 0x8c, 0xa1, 0x3e,
	0xfa, 0x93, 0xff, 0x93, 0x93, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xda, 0x31, 0xd0, 0xb2, 0x56,
	0x02, 0x00, 0x00,
}

func (x BatchOperationType) String() string {
	s, ok := BatchOperationType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x BatchOperationState) String() string {
	s, ok := BatchOperationState_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
