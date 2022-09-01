//
//Queensaver API
//
//Queensaver API to send in sensor data and retrieve it. It's also used for user management.
//
//The version of the OpenAPI document: 0.0.1
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: models/post_stands_response.proto

package models

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PostStandsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// HTTP response code. Used for internal purposes, will be sent out at the API.
	HttpResponseCode int32  `protobuf:"varint,1,opt,name=httpResponseCode,proto3" json:"httpResponseCode,omitempty"`
	Stand            *Stand `protobuf:"bytes,109757398,opt,name=stand,proto3" json:"stand,omitempty"`
}

func (x *PostStandsResponse) Reset() {
	*x = PostStandsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_post_stands_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostStandsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostStandsResponse) ProtoMessage() {}

func (x *PostStandsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_models_post_stands_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostStandsResponse.ProtoReflect.Descriptor instead.
func (*PostStandsResponse) Descriptor() ([]byte, []int) {
	return file_models_post_stands_response_proto_rawDescGZIP(), []int{0}
}

func (x *PostStandsResponse) GetHttpResponseCode() int32 {
	if x != nil {
		return x.HttpResponseCode
	}
	return 0
}

func (x *PostStandsResponse) GetStand() *Stand {
	if x != nil {
		return x.Stand
	}
	return nil
}

var File_models_post_stands_response_proto protoreflect.FileDescriptor

var file_models_post_stands_response_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x74,
	0x61, 0x6e, 0x64, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x1a, 0x12, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x69, 0x0a, 0x12, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x10, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x18, 0xd6, 0x87, 0xab, 0x34,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x53,
	0x74, 0x61, 0x6e, 0x64, 0x52, 0x05, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x50, 0x00, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_post_stands_response_proto_rawDescOnce sync.Once
	file_models_post_stands_response_proto_rawDescData = file_models_post_stands_response_proto_rawDesc
)

func file_models_post_stands_response_proto_rawDescGZIP() []byte {
	file_models_post_stands_response_proto_rawDescOnce.Do(func() {
		file_models_post_stands_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_post_stands_response_proto_rawDescData)
	})
	return file_models_post_stands_response_proto_rawDescData
}

var file_models_post_stands_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_post_stands_response_proto_goTypes = []interface{}{
	(*PostStandsResponse)(nil), // 0: openapi.PostStandsResponse
	(*Stand)(nil),              // 1: openapi.Stand
}
var file_models_post_stands_response_proto_depIdxs = []int32{
	1, // 0: openapi.PostStandsResponse.stand:type_name -> openapi.Stand
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_models_post_stands_response_proto_init() }
func file_models_post_stands_response_proto_init() {
	if File_models_post_stands_response_proto != nil {
		return
	}
	file_models_stand_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_post_stands_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostStandsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_models_post_stands_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_post_stands_response_proto_goTypes,
		DependencyIndexes: file_models_post_stands_response_proto_depIdxs,
		MessageInfos:      file_models_post_stands_response_proto_msgTypes,
	}.Build()
	File_models_post_stands_response_proto = out.File
	file_models_post_stands_response_proto_rawDesc = nil
	file_models_post_stands_response_proto_goTypes = nil
	file_models_post_stands_response_proto_depIdxs = nil
}
