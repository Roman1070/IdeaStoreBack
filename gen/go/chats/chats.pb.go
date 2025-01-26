// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v5.29.1
// source: chats.proto

package chatsv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SendMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *MessageData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return file_chats_proto_rawDescGZIP(), []int{0}
}

func (x *SendMessageRequest) GetData() *MessageData {
	if x != nil {
		return x.Data
	}
	return nil
}

type SendMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SendMessageResponse) Reset() {
	*x = SendMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageResponse) ProtoMessage() {}

func (x *SendMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageResponse.ProtoReflect.Descriptor instead.
func (*SendMessageResponse) Descriptor() ([]byte, []int) {
	return file_chats_proto_rawDescGZIP(), []int{1}
}

func (x *SendMessageResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstId  int64 `protobuf:"varint,1,opt,name=first_id,json=firstId,proto3" json:"first_id,omitempty"`
	SecondId int64 `protobuf:"varint,2,opt,name=second_id,json=secondId,proto3" json:"second_id,omitempty"`
}

func (x *GetMessagesRequest) Reset() {
	*x = GetMessagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessagesRequest) ProtoMessage() {}

func (x *GetMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessagesRequest.ProtoReflect.Descriptor instead.
func (*GetMessagesRequest) Descriptor() ([]byte, []int) {
	return file_chats_proto_rawDescGZIP(), []int{2}
}

func (x *GetMessagesRequest) GetFirstId() int64 {
	if x != nil {
		return x.FirstId
	}
	return 0
}

func (x *GetMessagesRequest) GetSecondId() int64 {
	if x != nil {
		return x.SecondId
	}
	return 0
}

type MessageData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SenderId           int64  `protobuf:"varint,2,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	RecieverId         int64  `protobuf:"varint,3,opt,name=reciever_id,json=recieverId,proto3" json:"reciever_id,omitempty"`
	FileName           string `protobuf:"bytes,4,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Text               string `protobuf:"bytes,5,opt,name=text,proto3" json:"text,omitempty"`
	SendingDate        string `protobuf:"bytes,6,opt,name=sending_date,json=sendingDate,proto3" json:"sending_date,omitempty"`
	CheckChatExistance bool   `protobuf:"varint,7,opt,name=check_chat_existance,json=checkChatExistance,proto3" json:"check_chat_existance,omitempty"`
}

func (x *MessageData) Reset() {
	*x = MessageData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chats_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageData) ProtoMessage() {}

func (x *MessageData) ProtoReflect() protoreflect.Message {
	mi := &file_chats_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageData.ProtoReflect.Descriptor instead.
func (*MessageData) Descriptor() ([]byte, []int) {
	return file_chats_proto_rawDescGZIP(), []int{3}
}

func (x *MessageData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MessageData) GetSenderId() int64 {
	if x != nil {
		return x.SenderId
	}
	return 0
}

func (x *MessageData) GetRecieverId() int64 {
	if x != nil {
		return x.RecieverId
	}
	return 0
}

func (x *MessageData) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *MessageData) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *MessageData) GetSendingDate() string {
	if x != nil {
		return x.SendingDate
	}
	return ""
}

func (x *MessageData) GetCheckChatExistance() bool {
	if x != nil {
		return x.CheckChatExistance
	}
	return false
}

type GetMessagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*MessageData `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *GetMessagesResponse) Reset() {
	*x = GetMessagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chats_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessagesResponse) ProtoMessage() {}

func (x *GetMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chats_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessagesResponse.ProtoReflect.Descriptor instead.
func (*GetMessagesResponse) Descriptor() ([]byte, []int) {
	return file_chats_proto_rawDescGZIP(), []int{4}
}

func (x *GetMessagesResponse) GetMessages() []*MessageData {
	if x != nil {
		return x.Messages
	}
	return nil
}

type CreateChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstId  int64 `protobuf:"varint,1,opt,name=first_id,json=firstId,proto3" json:"first_id,omitempty"`
	SecondId int64 `protobuf:"varint,2,opt,name=second_id,json=secondId,proto3" json:"second_id,omitempty"`
}

func (x *CreateChatRequest) Reset() {
	*x = CreateChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chats_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChatRequest) ProtoMessage() {}

func (x *CreateChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chats_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChatRequest.ProtoReflect.Descriptor instead.
func (*CreateChatRequest) Descriptor() ([]byte, []int) {
	return file_chats_proto_rawDescGZIP(), []int{5}
}

func (x *CreateChatRequest) GetFirstId() int64 {
	if x != nil {
		return x.FirstId
	}
	return 0
}

func (x *CreateChatRequest) GetSecondId() int64 {
	if x != nil {
		return x.SecondId
	}
	return 0
}

type GetUsersChatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUsersChatsRequest) Reset() {
	*x = GetUsersChatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chats_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersChatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersChatsRequest) ProtoMessage() {}

func (x *GetUsersChatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chats_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersChatsRequest.ProtoReflect.Descriptor instead.
func (*GetUsersChatsRequest) Descriptor() ([]byte, []int) {
	return file_chats_proto_rawDescGZIP(), []int{6}
}

func (x *GetUsersChatsRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetUsersChatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chats []*ChatUserData `protobuf:"bytes,1,rep,name=chats,proto3" json:"chats,omitempty"`
}

func (x *GetUsersChatsResponse) Reset() {
	*x = GetUsersChatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chats_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersChatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersChatsResponse) ProtoMessage() {}

func (x *GetUsersChatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chats_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersChatsResponse.ProtoReflect.Descriptor instead.
func (*GetUsersChatsResponse) Descriptor() ([]byte, []int) {
	return file_chats_proto_rawDescGZIP(), []int{7}
}

func (x *GetUsersChatsResponse) GetChats() []*ChatUserData {
	if x != nil {
		return x.Chats
	}
	return nil
}

type ChatUserData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Avatar string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (x *ChatUserData) Reset() {
	*x = ChatUserData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chats_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatUserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatUserData) ProtoMessage() {}

func (x *ChatUserData) ProtoReflect() protoreflect.Message {
	mi := &file_chats_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatUserData.ProtoReflect.Descriptor instead.
func (*ChatUserData) Descriptor() ([]byte, []int) {
	return file_chats_proto_rawDescGZIP(), []int{8}
}

func (x *ChatUserData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ChatUserData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChatUserData) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

type DeleteChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId int64 `protobuf:"varint,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
}

func (x *DeleteChatRequest) Reset() {
	*x = DeleteChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chats_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteChatRequest) ProtoMessage() {}

func (x *DeleteChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chats_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteChatRequest.ProtoReflect.Descriptor instead.
func (*DeleteChatRequest) Descriptor() ([]byte, []int) {
	return file_chats_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteChatRequest) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

var File_chats_proto protoreflect.FileDescriptor

var file_chats_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x68, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63,
	0x68, 0x61, 0x74, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x3c, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x73, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x25, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x49, 0x64, 0x22, 0xe1, 0x01, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x44, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f,
	0x63, 0x68, 0x61, 0x74, 0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x74, 0x45,
	0x78, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x45, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2e, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22,
	0x4b, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x66, 0x69, 0x72, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x42, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x73, 0x2e, 0x43, 0x68,
	0x61, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x63, 0x68, 0x61, 0x74,
	0x73, 0x22, 0x4a, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x2c, 0x0a,
	0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x32, 0xe5, 0x02, 0x0a, 0x05,
	0x43, 0x68, 0x61, 0x74, 0x73, 0x12, 0x46, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x19, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x73, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x73, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x19, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61,
	0x74, 0x12, 0x18, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x43, 0x68, 0x61, 0x74, 0x73, 0x12, 0x1b, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x40, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x12,
	0x18, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68,
	0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x42, 0x1d, 0x5a, 0x1b, 0x69, 0x64, 0x65, 0x61, 0x2d, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x3b, 0x63, 0x68, 0x61, 0x74, 0x73,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chats_proto_rawDescOnce sync.Once
	file_chats_proto_rawDescData = file_chats_proto_rawDesc
)

func file_chats_proto_rawDescGZIP() []byte {
	file_chats_proto_rawDescOnce.Do(func() {
		file_chats_proto_rawDescData = protoimpl.X.CompressGZIP(file_chats_proto_rawDescData)
	})
	return file_chats_proto_rawDescData
}

var file_chats_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_chats_proto_goTypes = []interface{}{
	(*SendMessageRequest)(nil),    // 0: chats.SendMessageRequest
	(*SendMessageResponse)(nil),   // 1: chats.SendMessageResponse
	(*GetMessagesRequest)(nil),    // 2: chats.GetMessagesRequest
	(*MessageData)(nil),           // 3: chats.MessageData
	(*GetMessagesResponse)(nil),   // 4: chats.GetMessagesResponse
	(*CreateChatRequest)(nil),     // 5: chats.CreateChatRequest
	(*GetUsersChatsRequest)(nil),  // 6: chats.GetUsersChatsRequest
	(*GetUsersChatsResponse)(nil), // 7: chats.GetUsersChatsResponse
	(*ChatUserData)(nil),          // 8: chats.ChatUserData
	(*DeleteChatRequest)(nil),     // 9: chats.DeleteChatRequest
	(*emptypb.Empty)(nil),         // 10: google.protobuf.Empty
}
var file_chats_proto_depIdxs = []int32{
	3,  // 0: chats.SendMessageRequest.data:type_name -> chats.MessageData
	3,  // 1: chats.GetMessagesResponse.messages:type_name -> chats.MessageData
	8,  // 2: chats.GetUsersChatsResponse.chats:type_name -> chats.ChatUserData
	0,  // 3: chats.Chats.SendMessage:input_type -> chats.SendMessageRequest
	2,  // 4: chats.Chats.GetMessages:input_type -> chats.GetMessagesRequest
	5,  // 5: chats.Chats.CreateChat:input_type -> chats.CreateChatRequest
	6,  // 6: chats.Chats.GetUsersChats:input_type -> chats.GetUsersChatsRequest
	9,  // 7: chats.Chats.DeleteChat:input_type -> chats.DeleteChatRequest
	1,  // 8: chats.Chats.SendMessage:output_type -> chats.SendMessageResponse
	4,  // 9: chats.Chats.GetMessages:output_type -> chats.GetMessagesResponse
	10, // 10: chats.Chats.CreateChat:output_type -> google.protobuf.Empty
	7,  // 11: chats.Chats.GetUsersChats:output_type -> chats.GetUsersChatsResponse
	10, // 12: chats.Chats.DeleteChat:output_type -> google.protobuf.Empty
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_chats_proto_init() }
func file_chats_proto_init() {
	if File_chats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageRequest); i {
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
		file_chats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageResponse); i {
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
		file_chats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessagesRequest); i {
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
		file_chats_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageData); i {
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
		file_chats_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessagesResponse); i {
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
		file_chats_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChatRequest); i {
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
		file_chats_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersChatsRequest); i {
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
		file_chats_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersChatsResponse); i {
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
		file_chats_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatUserData); i {
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
		file_chats_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteChatRequest); i {
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
			RawDescriptor: file_chats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chats_proto_goTypes,
		DependencyIndexes: file_chats_proto_depIdxs,
		MessageInfos:      file_chats_proto_msgTypes,
	}.Build()
	File_chats_proto = out.File
	file_chats_proto_rawDesc = nil
	file_chats_proto_goTypes = nil
	file_chats_proto_depIdxs = nil
}
