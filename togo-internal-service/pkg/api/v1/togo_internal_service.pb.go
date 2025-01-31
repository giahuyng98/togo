// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: api/v1/togo_internal_service.proto

package togo_internal_v1

import (
	date "google.golang.org/genproto/googleapis/type/date"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title       string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content     string                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	CreatedTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_togo_internal_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_togo_internal_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_api_v1_togo_internal_service_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Task) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Task) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Task) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Task) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

type ListTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string     `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Date      *date.Date `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	PageSize  int32      `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken string     `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListTaskRequest) Reset() {
	*x = ListTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_togo_internal_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTaskRequest) ProtoMessage() {}

func (x *ListTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_togo_internal_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTaskRequest.ProtoReflect.Descriptor instead.
func (*ListTaskRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_togo_internal_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListTaskRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ListTaskRequest) GetDate() *date.Date {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *ListTaskRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListTaskRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks         []*Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	NextPageToken string  `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListTaskResponse) Reset() {
	*x = ListTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_togo_internal_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTaskResponse) ProtoMessage() {}

func (x *ListTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_togo_internal_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTaskResponse.ProtoReflect.Descriptor instead.
func (*ListTaskResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_togo_internal_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListTaskResponse) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

func (x *ListTaskResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type GetTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTaskRequest) Reset() {
	*x = GetTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_togo_internal_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskRequest) ProtoMessage() {}

func (x *GetTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_togo_internal_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskRequest.ProtoReflect.Descriptor instead.
func (*GetTaskRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_togo_internal_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetTaskRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title       string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content     string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	CreatedTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
}

func (x *CreateTaskRequest) Reset() {
	*x = CreateTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_togo_internal_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskRequest) ProtoMessage() {}

func (x *CreateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_togo_internal_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskRequest.ProtoReflect.Descriptor instead.
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_togo_internal_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateTaskRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateTaskRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateTaskRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateTaskRequest) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

var File_api_v1_togo_internal_service_proto protoreflect.FileDescriptor

var file_api_v1_togo_internal_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x67, 0x6f, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x74, 0x6f, 0x67, 0x6f, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9e, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0x8d, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x68, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x6f, 0x67, 0x6f, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x9b, 0x01, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xfe, 0x01, 0x0a, 0x13, 0x54,
	0x6f, 0x67, 0x6f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x53, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x21,
	0x2e, 0x74, 0x6f, 0x67, 0x6f, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x74, 0x6f, 0x67, 0x6f, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x20, 0x2e, 0x74, 0x6f, 0x67, 0x6f, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x6f, 0x67, 0x6f, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x22, 0x00, 0x12, 0x4b,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x23, 0x2e, 0x74,
	0x6f, 0x67, 0x6f, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x74, 0x6f, 0x67, 0x6f, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x22, 0x00, 0x42, 0x24, 0x5a, 0x22, 0x2f,
	0x74, 0x6f, 0x67, 0x6f, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2d, 0x76, 0x31,
	0x3b, 0x74, 0x6f, 0x67, 0x6f, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_togo_internal_service_proto_rawDescOnce sync.Once
	file_api_v1_togo_internal_service_proto_rawDescData = file_api_v1_togo_internal_service_proto_rawDesc
)

func file_api_v1_togo_internal_service_proto_rawDescGZIP() []byte {
	file_api_v1_togo_internal_service_proto_rawDescOnce.Do(func() {
		file_api_v1_togo_internal_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_togo_internal_service_proto_rawDescData)
	})
	return file_api_v1_togo_internal_service_proto_rawDescData
}

var file_api_v1_togo_internal_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_v1_togo_internal_service_proto_goTypes = []interface{}{
	(*Task)(nil),                  // 0: togo.internal.v1.Task
	(*ListTaskRequest)(nil),       // 1: togo.internal.v1.ListTaskRequest
	(*ListTaskResponse)(nil),      // 2: togo.internal.v1.ListTaskResponse
	(*GetTaskRequest)(nil),        // 3: togo.internal.v1.GetTaskRequest
	(*CreateTaskRequest)(nil),     // 4: togo.internal.v1.CreateTaskRequest
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*date.Date)(nil),             // 6: google.type.Date
}
var file_api_v1_togo_internal_service_proto_depIdxs = []int32{
	5, // 0: togo.internal.v1.Task.created_time:type_name -> google.protobuf.Timestamp
	6, // 1: togo.internal.v1.ListTaskRequest.date:type_name -> google.type.Date
	0, // 2: togo.internal.v1.ListTaskResponse.tasks:type_name -> togo.internal.v1.Task
	5, // 3: togo.internal.v1.CreateTaskRequest.created_time:type_name -> google.protobuf.Timestamp
	1, // 4: togo.internal.v1.TogoInternalService.ListTask:input_type -> togo.internal.v1.ListTaskRequest
	3, // 5: togo.internal.v1.TogoInternalService.GetTask:input_type -> togo.internal.v1.GetTaskRequest
	4, // 6: togo.internal.v1.TogoInternalService.CreateTask:input_type -> togo.internal.v1.CreateTaskRequest
	2, // 7: togo.internal.v1.TogoInternalService.ListTask:output_type -> togo.internal.v1.ListTaskResponse
	0, // 8: togo.internal.v1.TogoInternalService.GetTask:output_type -> togo.internal.v1.Task
	0, // 9: togo.internal.v1.TogoInternalService.CreateTask:output_type -> togo.internal.v1.Task
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_v1_togo_internal_service_proto_init() }
func file_api_v1_togo_internal_service_proto_init() {
	if File_api_v1_togo_internal_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_togo_internal_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_api_v1_togo_internal_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTaskRequest); i {
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
		file_api_v1_togo_internal_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTaskResponse); i {
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
		file_api_v1_togo_internal_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTaskRequest); i {
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
		file_api_v1_togo_internal_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTaskRequest); i {
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
			RawDescriptor: file_api_v1_togo_internal_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_togo_internal_service_proto_goTypes,
		DependencyIndexes: file_api_v1_togo_internal_service_proto_depIdxs,
		MessageInfos:      file_api_v1_togo_internal_service_proto_msgTypes,
	}.Build()
	File_api_v1_togo_internal_service_proto = out.File
	file_api_v1_togo_internal_service_proto_rawDesc = nil
	file_api_v1_togo_internal_service_proto_goTypes = nil
	file_api_v1_togo_internal_service_proto_depIdxs = nil
}
