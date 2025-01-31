// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: api/v1/togo_user_session_service.proto

package togo_user_session_v1

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

type RegisterOrLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegisterOrLoginRequest) Reset() {
	*x = RegisterOrLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_togo_user_session_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterOrLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterOrLoginRequest) ProtoMessage() {}

func (x *RegisterOrLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_togo_user_session_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterOrLoginRequest.ProtoReflect.Descriptor instead.
func (*RegisterOrLoginRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_togo_user_session_service_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterOrLoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterOrLoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterOrLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RegisterOrLoginResponse) Reset() {
	*x = RegisterOrLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_togo_user_session_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterOrLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterOrLoginResponse) ProtoMessage() {}

func (x *RegisterOrLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_togo_user_session_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterOrLoginResponse.ProtoReflect.Descriptor instead.
func (*RegisterOrLoginResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_togo_user_session_service_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterOrLoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type VerifyTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *VerifyTokenRequest) Reset() {
	*x = VerifyTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_togo_user_session_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyTokenRequest) ProtoMessage() {}

func (x *VerifyTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_togo_user_session_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyTokenRequest.ProtoReflect.Descriptor instead.
func (*VerifyTokenRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_togo_user_session_service_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type VerifyTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *VerifyTokenResponse) Reset() {
	*x = VerifyTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_togo_user_session_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyTokenResponse) ProtoMessage() {}

func (x *VerifyTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_togo_user_session_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyTokenResponse.ProtoReflect.Descriptor instead.
func (*VerifyTokenResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_togo_user_session_service_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyTokenResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *VerifyTokenResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

var File_api_v1_togo_user_session_service_proto protoreflect.FileDescriptor

var file_api_v1_togo_user_session_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x67, 0x6f, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x74, 0x6f, 0x67, 0x6f, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0x50,
	0x0a, 0x16, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x72, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x2f, 0x0a, 0x17, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x72, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x2a, 0x0a, 0x12, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x4a, 0x0a,
	0x13, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xf0, 0x01, 0x0a, 0x16, 0x54, 0x6f,
	0x67, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x70, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x4f, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x2c, 0x2e, 0x74, 0x6f, 0x67, 0x6f, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x74, 0x6f, 0x67, 0x6f, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x28, 0x2e, 0x74, 0x6f, 0x67, 0x6f, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x29, 0x2e, 0x74, 0x6f, 0x67, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a,
	0x2f, 0x74, 0x6f, 0x67, 0x6f, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2d, 0x76, 0x31, 0x3b, 0x74, 0x6f, 0x67, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_v1_togo_user_session_service_proto_rawDescOnce sync.Once
	file_api_v1_togo_user_session_service_proto_rawDescData = file_api_v1_togo_user_session_service_proto_rawDesc
)

func file_api_v1_togo_user_session_service_proto_rawDescGZIP() []byte {
	file_api_v1_togo_user_session_service_proto_rawDescOnce.Do(func() {
		file_api_v1_togo_user_session_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_togo_user_session_service_proto_rawDescData)
	})
	return file_api_v1_togo_user_session_service_proto_rawDescData
}

var file_api_v1_togo_user_session_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_v1_togo_user_session_service_proto_goTypes = []interface{}{
	(*RegisterOrLoginRequest)(nil),  // 0: togo.user_session.v1.RegisterOrLoginRequest
	(*RegisterOrLoginResponse)(nil), // 1: togo.user_session.v1.RegisterOrLoginResponse
	(*VerifyTokenRequest)(nil),      // 2: togo.user_session.v1.VerifyTokenRequest
	(*VerifyTokenResponse)(nil),     // 3: togo.user_session.v1.VerifyTokenResponse
}
var file_api_v1_togo_user_session_service_proto_depIdxs = []int32{
	0, // 0: togo.user_session.v1.TogoUserSessionService.RegisterOrLogin:input_type -> togo.user_session.v1.RegisterOrLoginRequest
	2, // 1: togo.user_session.v1.TogoUserSessionService.VerifyToken:input_type -> togo.user_session.v1.VerifyTokenRequest
	1, // 2: togo.user_session.v1.TogoUserSessionService.RegisterOrLogin:output_type -> togo.user_session.v1.RegisterOrLoginResponse
	3, // 3: togo.user_session.v1.TogoUserSessionService.VerifyToken:output_type -> togo.user_session.v1.VerifyTokenResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1_togo_user_session_service_proto_init() }
func file_api_v1_togo_user_session_service_proto_init() {
	if File_api_v1_togo_user_session_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_togo_user_session_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterOrLoginRequest); i {
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
		file_api_v1_togo_user_session_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterOrLoginResponse); i {
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
		file_api_v1_togo_user_session_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyTokenRequest); i {
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
		file_api_v1_togo_user_session_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyTokenResponse); i {
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
			RawDescriptor: file_api_v1_togo_user_session_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_togo_user_session_service_proto_goTypes,
		DependencyIndexes: file_api_v1_togo_user_session_service_proto_depIdxs,
		MessageInfos:      file_api_v1_togo_user_session_service_proto_msgTypes,
	}.Build()
	File_api_v1_togo_user_session_service_proto = out.File
	file_api_v1_togo_user_session_service_proto_rawDesc = nil
	file_api_v1_togo_user_session_service_proto_goTypes = nil
	file_api_v1_togo_user_session_service_proto_depIdxs = nil
}
