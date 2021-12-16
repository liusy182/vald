//
// Copyright (C) 2019-2021 vdaas.org vald team <vald@vdaas.org>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: apis/proto/v1/vald/search.proto

package vald

import (
	payload "github.com/vdaas/vald/apis/grpc/v1/payload"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_apis_proto_v1_vald_search_proto protoreflect.FileDescriptor

var file_apis_proto_v1_vald_search_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f,
	0x76, 0x61, 0x6c, 0x64, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x76, 0x61, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x23, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd0,
	0x09, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x55, 0x0a, 0x06, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x12, 0x1a, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0c, 0x22, 0x07, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x3a, 0x01, 0x2a,
	0x12, 0x5e, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1c,
	0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x2e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0f, 0x22, 0x0a, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x69, 0x64, 0x3a, 0x01, 0x2a,
	0x12, 0x53, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x12, 0x1a, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x59, 0x0a, 0x10, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1c, 0x2e, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01,
	0x12, 0x69, 0x0a, 0x0b, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x1f, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x22, 0x1b,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x72, 0x0a, 0x0f, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x79, 0x49, 0x44, 0x12, 0x21,
	0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x22,
	0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x2f, 0x69, 0x64, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x3a, 0x01, 0x2a, 0x12,
	0x61, 0x0a, 0x0c, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x1a, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12,
	0x22, 0x0d, 0x2f, 0x6c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x3a,
	0x01, 0x2a, 0x12, 0x6a, 0x0a, 0x10, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1c, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x6c, 0x69, 0x6e, 0x65,
	0x61, 0x72, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x69, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x59,
	0x0a, 0x12, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x12, 0x1a, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x5f, 0x0a, 0x16, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42,
	0x79, 0x49, 0x44, 0x12, 0x1c, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x75, 0x0a, 0x11, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x1f, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x22, 0x21,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x6c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x3a, 0x01,
	0x2a, 0x12, 0x7e, 0x0a, 0x15, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x72,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x79, 0x49, 0x44, 0x12, 0x21, 0x2e, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x22, 0x24, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f, 0x6c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x2f, 0x69, 0x64, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x3a, 0x01,
	0x2a, 0x42, 0x53, 0x0a, 0x1a, 0x6f, 0x72, 0x67, 0x2e, 0x76, 0x64, 0x61, 0x61, 0x73, 0x2e, 0x76,
	0x61, 0x6c, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x76, 0x61, 0x6c, 0x64, 0x42,
	0x0a, 0x56, 0x61, 0x6c, 0x64, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x01, 0x5a, 0x27, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x64, 0x61, 0x61, 0x73, 0x2f,
	0x76, 0x61, 0x6c, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76,
	0x31, 0x2f, 0x76, 0x61, 0x6c, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_apis_proto_v1_vald_search_proto_goTypes = []interface{}{
	(*payload.Search_Request)(nil),        // 0: payload.v1.Search.Request
	(*payload.Search_IDRequest)(nil),      // 1: payload.v1.Search.IDRequest
	(*payload.Search_MultiRequest)(nil),   // 2: payload.v1.Search.MultiRequest
	(*payload.Search_MultiIDRequest)(nil), // 3: payload.v1.Search.MultiIDRequest
	(*payload.Search_Response)(nil),       // 4: payload.v1.Search.Response
	(*payload.Search_StreamResponse)(nil), // 5: payload.v1.Search.StreamResponse
	(*payload.Search_Responses)(nil),      // 6: payload.v1.Search.Responses
}
var file_apis_proto_v1_vald_search_proto_depIdxs = []int32{
	0,  // 0: vald.v1.Search.Search:input_type -> payload.v1.Search.Request
	1,  // 1: vald.v1.Search.SearchByID:input_type -> payload.v1.Search.IDRequest
	0,  // 2: vald.v1.Search.StreamSearch:input_type -> payload.v1.Search.Request
	1,  // 3: vald.v1.Search.StreamSearchByID:input_type -> payload.v1.Search.IDRequest
	2,  // 4: vald.v1.Search.MultiSearch:input_type -> payload.v1.Search.MultiRequest
	3,  // 5: vald.v1.Search.MultiSearchByID:input_type -> payload.v1.Search.MultiIDRequest
	0,  // 6: vald.v1.Search.LinearSearch:input_type -> payload.v1.Search.Request
	1,  // 7: vald.v1.Search.LinearSearchByID:input_type -> payload.v1.Search.IDRequest
	0,  // 8: vald.v1.Search.StreamLinearSearch:input_type -> payload.v1.Search.Request
	1,  // 9: vald.v1.Search.StreamLinearSearchByID:input_type -> payload.v1.Search.IDRequest
	2,  // 10: vald.v1.Search.MultiLinearSearch:input_type -> payload.v1.Search.MultiRequest
	3,  // 11: vald.v1.Search.MultiLinearSearchByID:input_type -> payload.v1.Search.MultiIDRequest
	4,  // 12: vald.v1.Search.Search:output_type -> payload.v1.Search.Response
	4,  // 13: vald.v1.Search.SearchByID:output_type -> payload.v1.Search.Response
	5,  // 14: vald.v1.Search.StreamSearch:output_type -> payload.v1.Search.StreamResponse
	5,  // 15: vald.v1.Search.StreamSearchByID:output_type -> payload.v1.Search.StreamResponse
	6,  // 16: vald.v1.Search.MultiSearch:output_type -> payload.v1.Search.Responses
	6,  // 17: vald.v1.Search.MultiSearchByID:output_type -> payload.v1.Search.Responses
	4,  // 18: vald.v1.Search.LinearSearch:output_type -> payload.v1.Search.Response
	4,  // 19: vald.v1.Search.LinearSearchByID:output_type -> payload.v1.Search.Response
	5,  // 20: vald.v1.Search.StreamLinearSearch:output_type -> payload.v1.Search.StreamResponse
	5,  // 21: vald.v1.Search.StreamLinearSearchByID:output_type -> payload.v1.Search.StreamResponse
	6,  // 22: vald.v1.Search.MultiLinearSearch:output_type -> payload.v1.Search.Responses
	6,  // 23: vald.v1.Search.MultiLinearSearchByID:output_type -> payload.v1.Search.Responses
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_apis_proto_v1_vald_search_proto_init() }
func file_apis_proto_v1_vald_search_proto_init() {
	if File_apis_proto_v1_vald_search_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apis_proto_v1_vald_search_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apis_proto_v1_vald_search_proto_goTypes,
		DependencyIndexes: file_apis_proto_v1_vald_search_proto_depIdxs,
	}.Build()
	File_apis_proto_v1_vald_search_proto = out.File
	file_apis_proto_v1_vald_search_proto_rawDesc = nil
	file_apis_proto_v1_vald_search_proto_goTypes = nil
	file_apis_proto_v1_vald_search_proto_depIdxs = nil
}
