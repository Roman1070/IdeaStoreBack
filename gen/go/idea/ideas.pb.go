// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v5.29.1
// source: ideas.proto

package ideasv1

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

type GetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetAllRequest) Reset() {
	*x = GetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ideas_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRequest) ProtoMessage() {}

func (x *GetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ideas_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRequest.ProtoReflect.Descriptor instead.
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return file_ideas_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetIdeasRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GetIdeasRequest) Reset() {
	*x = GetIdeasRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ideas_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIdeasRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIdeasRequest) ProtoMessage() {}

func (x *GetIdeasRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ideas_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIdeasRequest.ProtoReflect.Descriptor instead.
func (*GetIdeasRequest) Descriptor() ([]byte, []int) {
	return file_ideas_proto_rawDescGZIP(), []int{1}
}

func (x *GetIdeasRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type GetIdeasResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ideas []*IdeaData `protobuf:"bytes,1,rep,name=ideas,proto3" json:"ideas,omitempty"`
}

func (x *GetIdeasResponse) Reset() {
	*x = GetIdeasResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ideas_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIdeasResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIdeasResponse) ProtoMessage() {}

func (x *GetIdeasResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ideas_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIdeasResponse.ProtoReflect.Descriptor instead.
func (*GetIdeasResponse) Descriptor() ([]byte, []int) {
	return file_ideas_proto_rawDescGZIP(), []int{2}
}

func (x *GetIdeasResponse) GetIdeas() []*IdeaData {
	if x != nil {
		return x.Ideas
	}
	return nil
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Link        string `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`
	Tags        string `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`
	Image       string `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	UserId      int64  `protobuf:"varint,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ideas_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ideas_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_ideas_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateRequest) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *CreateRequest) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *CreateRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *CreateRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdeaId int64 `protobuf:"varint,1,opt,name=idea_id,json=ideaId,proto3" json:"idea_id,omitempty"` // ID of the created idea.
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ideas_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ideas_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_ideas_proto_rawDescGZIP(), []int{4}
}

func (x *CreateResponse) GetIdeaId() int64 {
	if x != nil {
		return x.IdeaId
	}
	return 0
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdeaId int64 `protobuf:"varint,1,opt,name=idea_id,json=ideaId,proto3" json:"idea_id,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ideas_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ideas_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_ideas_proto_rawDescGZIP(), []int{5}
}

func (x *GetRequest) GetIdeaId() int64 {
	if x != nil {
		return x.IdeaId
	}
	return 0
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Link        string `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`
	Tags        string `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`
	Image       string `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	UserId      int64  `protobuf:"varint,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Likes       int32  `protobuf:"varint,7,opt,name=likes,proto3" json:"likes,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ideas_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ideas_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_ideas_proto_rawDescGZIP(), []int{6}
}

func (x *GetResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetResponse) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *GetResponse) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *GetResponse) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *GetResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetResponse) GetLikes() int32 {
	if x != nil {
		return x.Likes
	}
	return 0
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdeaId int64 `protobuf:"varint,1,opt,name=idea_id,json=ideaId,proto3" json:"idea_id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ideas_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ideas_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_ideas_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteRequest) GetIdeaId() int64 {
	if x != nil {
		return x.IdeaId
	}
	return 0
}

type IdeaData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdeaId      int64  `protobuf:"varint,1,opt,name=idea_id,json=ideaId,proto3" json:"idea_id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Link        string `protobuf:"bytes,4,opt,name=link,proto3" json:"link,omitempty"`
	Tags        string `protobuf:"bytes,5,opt,name=tags,proto3" json:"tags,omitempty"`
	Image       string `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty"`
	Saved       bool   `protobuf:"varint,7,opt,name=saved,proto3" json:"saved,omitempty"`
}

func (x *IdeaData) Reset() {
	*x = IdeaData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ideas_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdeaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdeaData) ProtoMessage() {}

func (x *IdeaData) ProtoReflect() protoreflect.Message {
	mi := &file_ideas_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdeaData.ProtoReflect.Descriptor instead.
func (*IdeaData) Descriptor() ([]byte, []int) {
	return file_ideas_proto_rawDescGZIP(), []int{8}
}

func (x *IdeaData) GetIdeaId() int64 {
	if x != nil {
		return x.IdeaId
	}
	return 0
}

func (x *IdeaData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *IdeaData) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *IdeaData) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *IdeaData) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *IdeaData) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *IdeaData) GetSaved() bool {
	if x != nil {
		return x.Saved
	}
	return false
}

type GetAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ideas []*IdeaData `protobuf:"bytes,1,rep,name=ideas,proto3" json:"ideas,omitempty"`
}

func (x *GetAllResponse) Reset() {
	*x = GetAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ideas_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllResponse) ProtoMessage() {}

func (x *GetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ideas_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllResponse.ProtoReflect.Descriptor instead.
func (*GetAllResponse) Descriptor() ([]byte, []int) {
	return file_ideas_proto_rawDescGZIP(), []int{9}
}

func (x *GetAllResponse) GetIdeas() []*IdeaData {
	if x != nil {
		return x.Ideas
	}
	return nil
}

var File_ideas_proto protoreflect.FileDescriptor

var file_ideas_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x69, 0x64, 0x65, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x69,
	0x64, 0x65, 0x61, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x28, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x23, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x49, 0x64, 0x65, 0x61, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x22, 0x39, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x49, 0x64, 0x65, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x69, 0x64, 0x65, 0x61, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x69, 0x64, 0x65, 0x61, 0x73, 0x2e, 0x49, 0x64, 0x65, 0x61,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x69, 0x64, 0x65, 0x61, 0x73, 0x22, 0x9c, 0x01, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x29, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x69, 0x64, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69,
	0x64, 0x65, 0x61, 0x49, 0x64, 0x22, 0x25, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x64, 0x65, 0x61, 0x49, 0x64, 0x22, 0xb0, 0x01, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6b,
	0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x22,
	0x28, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x69, 0x64, 0x65, 0x61, 0x49, 0x64, 0x22, 0xad, 0x01, 0x0a, 0x08, 0x49, 0x64,
	0x65, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x65, 0x61, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x64, 0x65, 0x61, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x61, 0x76, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x73, 0x61, 0x76, 0x65, 0x64, 0x22, 0x37, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x69,
	0x64, 0x65, 0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x69, 0x64, 0x65,
	0x61, 0x73, 0x2e, 0x49, 0x64, 0x65, 0x61, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x69, 0x64, 0x65,
	0x61, 0x73, 0x32, 0xab, 0x02, 0x0a, 0x05, 0x49, 0x64, 0x65, 0x61, 0x73, 0x12, 0x39, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x64, 0x65, 0x61, 0x12, 0x14, 0x2e, 0x69, 0x64, 0x65,
	0x61, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x69, 0x64, 0x65, 0x61, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x64,
	0x65, 0x61, 0x12, 0x11, 0x2e, 0x69, 0x64, 0x65, 0x61, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x69, 0x64, 0x65, 0x61, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x49, 0x64, 0x65, 0x61, 0x12, 0x14, 0x2e, 0x69, 0x64, 0x65, 0x61, 0x73, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x49, 0x64, 0x65, 0x61, 0x73, 0x12, 0x14, 0x2e, 0x69, 0x64, 0x65, 0x61, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x69,
	0x64, 0x65, 0x61, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x64, 0x65, 0x61, 0x73, 0x12,
	0x16, 0x2e, 0x69, 0x64, 0x65, 0x61, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x64, 0x65, 0x61, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x69, 0x64, 0x65, 0x61, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x49, 0x64, 0x65, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x1d, 0x5a, 0x1b, 0x69, 0x64, 0x65, 0x61, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x69,
	0x64, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x3b, 0x69, 0x64, 0x65, 0x61, 0x73, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ideas_proto_rawDescOnce sync.Once
	file_ideas_proto_rawDescData = file_ideas_proto_rawDesc
)

func file_ideas_proto_rawDescGZIP() []byte {
	file_ideas_proto_rawDescOnce.Do(func() {
		file_ideas_proto_rawDescData = protoimpl.X.CompressGZIP(file_ideas_proto_rawDescData)
	})
	return file_ideas_proto_rawDescData
}

var file_ideas_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_ideas_proto_goTypes = []interface{}{
	(*GetAllRequest)(nil),    // 0: ideas.GetAllRequest
	(*GetIdeasRequest)(nil),  // 1: ideas.GetIdeasRequest
	(*GetIdeasResponse)(nil), // 2: ideas.GetIdeasResponse
	(*CreateRequest)(nil),    // 3: ideas.CreateRequest
	(*CreateResponse)(nil),   // 4: ideas.CreateResponse
	(*GetRequest)(nil),       // 5: ideas.GetRequest
	(*GetResponse)(nil),      // 6: ideas.GetResponse
	(*DeleteRequest)(nil),    // 7: ideas.DeleteRequest
	(*IdeaData)(nil),         // 8: ideas.IdeaData
	(*GetAllResponse)(nil),   // 9: ideas.GetAllResponse
	(*emptypb.Empty)(nil),    // 10: google.protobuf.Empty
}
var file_ideas_proto_depIdxs = []int32{
	8,  // 0: ideas.GetIdeasResponse.ideas:type_name -> ideas.IdeaData
	8,  // 1: ideas.GetAllResponse.ideas:type_name -> ideas.IdeaData
	3,  // 2: ideas.Ideas.CreateIdea:input_type -> ideas.CreateRequest
	5,  // 3: ideas.Ideas.GetIdea:input_type -> ideas.GetRequest
	7,  // 4: ideas.Ideas.DeleteIdea:input_type -> ideas.DeleteRequest
	0,  // 5: ideas.Ideas.GetAllIdeas:input_type -> ideas.GetAllRequest
	1,  // 6: ideas.Ideas.GetIdeas:input_type -> ideas.GetIdeasRequest
	4,  // 7: ideas.Ideas.CreateIdea:output_type -> ideas.CreateResponse
	6,  // 8: ideas.Ideas.GetIdea:output_type -> ideas.GetResponse
	10, // 9: ideas.Ideas.DeleteIdea:output_type -> google.protobuf.Empty
	9,  // 10: ideas.Ideas.GetAllIdeas:output_type -> ideas.GetAllResponse
	2,  // 11: ideas.Ideas.GetIdeas:output_type -> ideas.GetIdeasResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_ideas_proto_init() }
func file_ideas_proto_init() {
	if File_ideas_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ideas_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRequest); i {
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
		file_ideas_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIdeasRequest); i {
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
		file_ideas_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIdeasResponse); i {
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
		file_ideas_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_ideas_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_ideas_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_ideas_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_ideas_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_ideas_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdeaData); i {
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
		file_ideas_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllResponse); i {
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
			RawDescriptor: file_ideas_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ideas_proto_goTypes,
		DependencyIndexes: file_ideas_proto_depIdxs,
		MessageInfos:      file_ideas_proto_msgTypes,
	}.Build()
	File_ideas_proto = out.File
	file_ideas_proto_rawDesc = nil
	file_ideas_proto_goTypes = nil
	file_ideas_proto_depIdxs = nil
}
