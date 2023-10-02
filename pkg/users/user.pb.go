// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.24.2
// source: users/user.proto

package users

import (
	comm "./comm"
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

type LoginMethod int32

const (
	LoginMethod_WX LoginMethod = 0
	LoginMethod_ID LoginMethod = 1
)

// Enum value maps for LoginMethod.
var (
	LoginMethod_name = map[int32]string{
		0: "WX",
		1: "ID",
	}
	LoginMethod_value = map[string]int32{
		"WX": 0,
		"ID": 1,
	}
)

func (x LoginMethod) Enum() *LoginMethod {
	p := new(LoginMethod)
	*p = x
	return p
}

func (x LoginMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoginMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_users_user_proto_enumTypes[0].Descriptor()
}

func (LoginMethod) Type() protoreflect.EnumType {
	return &file_users_user_proto_enumTypes[0]
}

func (x LoginMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoginMethod.Descriptor instead.
func (LoginMethod) EnumDescriptor() ([]byte, []int) {
	return file_users_user_proto_rawDescGZIP(), []int{0}
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method  LoginMethod `protobuf:"varint,1,opt,name=method,proto3,enum=users.LoginMethod" json:"method,omitempty"`
	Payload string      `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_users_user_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetMethod() LoginMethod {
	if x != nil {
		return x.Method
	}
	return LoginMethod_WX
}

func (x *LoginRequest) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *comm.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Uid    uint32               `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_users_user_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetHeader() *comm.ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *LoginResponse) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

var File_users_user_proto protoreflect.FileDescriptor

var file_users_user_proto_rawDesc = []byte{
	0x0a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x1a, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x2f,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a,
	0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a,
	0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x22, 0x51, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x2a, 0x1d, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x06, 0x0a, 0x02, 0x57, 0x58, 0x10, 0x00, 0x12, 0x06, 0x0a,
	0x02, 0x49, 0x44, 0x10, 0x01, 0x32, 0x3c, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x34, 0x0a,
	0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_users_user_proto_rawDescOnce sync.Once
	file_users_user_proto_rawDescData = file_users_user_proto_rawDesc
)

func file_users_user_proto_rawDescGZIP() []byte {
	file_users_user_proto_rawDescOnce.Do(func() {
		file_users_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_users_user_proto_rawDescData)
	})
	return file_users_user_proto_rawDescData
}

var file_users_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_users_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_users_user_proto_goTypes = []interface{}{
	(LoginMethod)(0),            // 0: users.LoginMethod
	(*LoginRequest)(nil),        // 1: users.LoginRequest
	(*LoginResponse)(nil),       // 2: users.LoginResponse
	(*comm.ResponseHeader)(nil), // 3: global.ResponseHeader
}
var file_users_user_proto_depIdxs = []int32{
	0, // 0: users.LoginRequest.method:type_name -> users.LoginMethod
	3, // 1: users.LoginResponse.header:type_name -> global.ResponseHeader
	1, // 2: users.Auth.login:input_type -> users.LoginRequest
	2, // 3: users.Auth.login:output_type -> users.LoginResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_users_user_proto_init() }
func file_users_user_proto_init() {
	if File_users_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_users_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_users_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
			RawDescriptor: file_users_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_users_user_proto_goTypes,
		DependencyIndexes: file_users_user_proto_depIdxs,
		EnumInfos:         file_users_user_proto_enumTypes,
		MessageInfos:      file_users_user_proto_msgTypes,
	}.Build()
	File_users_user_proto = out.File
	file_users_user_proto_rawDesc = nil
	file_users_user_proto_goTypes = nil
	file_users_user_proto_depIdxs = nil
}
