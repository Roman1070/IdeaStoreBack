// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v5.29.1
// source: boards.proto

package boardsv1

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

type DeleteBoardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BoardId int64 `protobuf:"varint,2,opt,name=board_id,json=boardId,proto3" json:"board_id,omitempty"`
}

func (x *DeleteBoardRequest) Reset() {
	*x = DeleteBoardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boards_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBoardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBoardRequest) ProtoMessage() {}

func (x *DeleteBoardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_boards_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBoardRequest.ProtoReflect.Descriptor instead.
func (*DeleteBoardRequest) Descriptor() ([]byte, []int) {
	return file_boards_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteBoardRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DeleteBoardRequest) GetBoardId() int64 {
	if x != nil {
		return x.BoardId
	}
	return 0
}

type GetIdeasInBoardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoardId int64 `protobuf:"varint,1,opt,name=board_id,json=boardId,proto3" json:"board_id,omitempty"`
}

func (x *GetIdeasInBoardRequest) Reset() {
	*x = GetIdeasInBoardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boards_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIdeasInBoardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIdeasInBoardRequest) ProtoMessage() {}

func (x *GetIdeasInBoardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_boards_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIdeasInBoardRequest.ProtoReflect.Descriptor instead.
func (*GetIdeasInBoardRequest) Descriptor() ([]byte, []int) {
	return file_boards_proto_rawDescGZIP(), []int{1}
}

func (x *GetIdeasInBoardRequest) GetBoardId() int64 {
	if x != nil {
		return x.BoardId
	}
	return 0
}

type GetIdeasInBoardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ideas []*IdeaData `protobuf:"bytes,1,rep,name=ideas,proto3" json:"ideas,omitempty"`
}

func (x *GetIdeasInBoardResponse) Reset() {
	*x = GetIdeasInBoardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boards_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIdeasInBoardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIdeasInBoardResponse) ProtoMessage() {}

func (x *GetIdeasInBoardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_boards_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIdeasInBoardResponse.ProtoReflect.Descriptor instead.
func (*GetIdeasInBoardResponse) Descriptor() ([]byte, []int) {
	return file_boards_proto_rawDescGZIP(), []int{2}
}

func (x *GetIdeasInBoardResponse) GetIdeas() []*IdeaData {
	if x != nil {
		return x.Ideas
	}
	return nil
}

type GetAllBoardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetAllBoardsRequest) Reset() {
	*x = GetAllBoardsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boards_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllBoardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllBoardsRequest) ProtoMessage() {}

func (x *GetAllBoardsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_boards_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllBoardsRequest.ProtoReflect.Descriptor instead.
func (*GetAllBoardsRequest) Descriptor() ([]byte, []int) {
	return file_boards_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllBoardsRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type SetIdeaSavedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdeaId  int64 `protobuf:"varint,1,opt,name=idea_id,json=ideaId,proto3" json:"idea_id,omitempty"`
	BoardId int64 `protobuf:"varint,2,opt,name=board_id,json=boardId,proto3" json:"board_id,omitempty"`
	Saved   bool  `protobuf:"varint,3,opt,name=saved,proto3" json:"saved,omitempty"`
}

func (x *SetIdeaSavedRequest) Reset() {
	*x = SetIdeaSavedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boards_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetIdeaSavedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetIdeaSavedRequest) ProtoMessage() {}

func (x *SetIdeaSavedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_boards_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetIdeaSavedRequest.ProtoReflect.Descriptor instead.
func (*SetIdeaSavedRequest) Descriptor() ([]byte, []int) {
	return file_boards_proto_rawDescGZIP(), []int{4}
}

func (x *SetIdeaSavedRequest) GetIdeaId() int64 {
	if x != nil {
		return x.IdeaId
	}
	return 0
}

func (x *SetIdeaSavedRequest) GetBoardId() int64 {
	if x != nil {
		return x.BoardId
	}
	return 0
}

func (x *SetIdeaSavedRequest) GetSaved() bool {
	if x != nil {
		return x.Saved
	}
	return false
}

type CreateBoardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	UserId int64  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CreateBoardRequest) Reset() {
	*x = CreateBoardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boards_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBoardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBoardRequest) ProtoMessage() {}

func (x *CreateBoardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_boards_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBoardRequest.ProtoReflect.Descriptor instead.
func (*CreateBoardRequest) Descriptor() ([]byte, []int) {
	return file_boards_proto_rawDescGZIP(), []int{5}
}

func (x *CreateBoardRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateBoardRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CreateBoardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateBoardResponse) Reset() {
	*x = CreateBoardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boards_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBoardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBoardResponse) ProtoMessage() {}

func (x *CreateBoardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_boards_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBoardResponse.ProtoReflect.Descriptor instead.
func (*CreateBoardResponse) Descriptor() ([]byte, []int) {
	return file_boards_proto_rawDescGZIP(), []int{6}
}

func (x *CreateBoardResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetBoardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBoardRequest) Reset() {
	*x = GetBoardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boards_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBoardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBoardRequest) ProtoMessage() {}

func (x *GetBoardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_boards_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBoardRequest.ProtoReflect.Descriptor instead.
func (*GetBoardRequest) Descriptor() ([]byte, []int) {
	return file_boards_proto_rawDescGZIP(), []int{7}
}

func (x *GetBoardRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetBoardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IdeasIds []int64 `protobuf:"varint,3,rep,packed,name=ideas_ids,json=ideasIds,proto3" json:"ideas_ids,omitempty"`
}

func (x *GetBoardResponse) Reset() {
	*x = GetBoardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boards_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBoardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBoardResponse) ProtoMessage() {}

func (x *GetBoardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_boards_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBoardResponse.ProtoReflect.Descriptor instead.
func (*GetBoardResponse) Descriptor() ([]byte, []int) {
	return file_boards_proto_rawDescGZIP(), []int{8}
}

func (x *GetBoardResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetBoardResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetBoardResponse) GetIdeasIds() []int64 {
	if x != nil {
		return x.IdeasIds
	}
	return nil
}

type BoardData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IdeasIds []int64 `protobuf:"varint,3,rep,packed,name=ideas_ids,json=ideasIds,proto3" json:"ideas_ids,omitempty"`
}

func (x *BoardData) Reset() {
	*x = BoardData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boards_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoardData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoardData) ProtoMessage() {}

func (x *BoardData) ProtoReflect() protoreflect.Message {
	mi := &file_boards_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoardData.ProtoReflect.Descriptor instead.
func (*BoardData) Descriptor() ([]byte, []int) {
	return file_boards_proto_rawDescGZIP(), []int{9}
}

func (x *BoardData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BoardData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BoardData) GetIdeasIds() []int64 {
	if x != nil {
		return x.IdeasIds
	}
	return nil
}

type GetAllBoardsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Boards []*BoardData `protobuf:"bytes,1,rep,name=boards,proto3" json:"boards,omitempty"`
}

func (x *GetAllBoardsResponse) Reset() {
	*x = GetAllBoardsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boards_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllBoardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllBoardsResponse) ProtoMessage() {}

func (x *GetAllBoardsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_boards_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllBoardsResponse.ProtoReflect.Descriptor instead.
func (*GetAllBoardsResponse) Descriptor() ([]byte, []int) {
	return file_boards_proto_rawDescGZIP(), []int{10}
}

func (x *GetAllBoardsResponse) GetBoards() []*BoardData {
	if x != nil {
		return x.Boards
	}
	return nil
}

type IdeaData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdeaId int64  `protobuf:"varint,1,opt,name=idea_id,json=ideaId,proto3" json:"idea_id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Image  string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *IdeaData) Reset() {
	*x = IdeaData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boards_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdeaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdeaData) ProtoMessage() {}

func (x *IdeaData) ProtoReflect() protoreflect.Message {
	mi := &file_boards_proto_msgTypes[11]
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
	return file_boards_proto_rawDescGZIP(), []int{11}
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

func (x *IdeaData) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

var File_boards_proto protoreflect.FileDescriptor

var file_boards_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x64, 0x22, 0x33, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x49, 0x64, 0x65, 0x61, 0x73, 0x49, 0x6e, 0x42, 0x6f, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x49, 0x64, 0x22, 0x41, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x49, 0x64, 0x65, 0x61, 0x73, 0x49, 0x6e,
	0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a,
	0x05, 0x69, 0x64, 0x65, 0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x49, 0x64, 0x65, 0x61, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05,
	0x69, 0x64, 0x65, 0x61, 0x73, 0x22, 0x2e, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42,
	0x6f, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5f, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x49, 0x64, 0x65, 0x61,
	0x53, 0x61, 0x76, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x69, 0x64, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69,
	0x64, 0x65, 0x61, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x61, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x73, 0x61, 0x76, 0x65, 0x64, 0x22, 0x41, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x25, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x53, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69,
	0x64, 0x65, 0x61, 0x73, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08,
	0x69, 0x64, 0x65, 0x61, 0x73, 0x49, 0x64, 0x73, 0x22, 0x4c, 0x0a, 0x09, 0x42, 0x6f, 0x61, 0x72,
	0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x64, 0x65,
	0x61, 0x73, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08, 0x69, 0x64,
	0x65, 0x61, 0x73, 0x49, 0x64, 0x73, 0x22, 0x41, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x42, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29,
	0x0a, 0x06, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x06, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x22, 0x4d, 0x0a, 0x08, 0x49, 0x64, 0x65,
	0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x65, 0x61, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x64, 0x65, 0x61, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x32, 0xba, 0x03, 0x0a, 0x06, 0x42, 0x6f, 0x61,
	0x72, 0x64, 0x73, 0x12, 0x46, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x61,
	0x72, 0x64, 0x12, 0x1a, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f,
	0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x17, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x12, 0x1b, 0x2e, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x49, 0x64, 0x65, 0x61,
	0x53, 0x61, 0x76, 0x65, 0x64, 0x12, 0x1b, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x53,
	0x65, 0x74, 0x49, 0x64, 0x65, 0x61, 0x53, 0x61, 0x76, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x49, 0x64, 0x65, 0x61, 0x73, 0x49, 0x6e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x12,
	0x1e, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x64, 0x65, 0x61,
	0x73, 0x49, 0x6e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x64, 0x65, 0x61,
	0x73, 0x49, 0x6e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x43, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x12,
	0x1a, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x1f, 0x5a, 0x1d, 0x69, 0x64, 0x65, 0x61, 0x2d, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x3b, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_boards_proto_rawDescOnce sync.Once
	file_boards_proto_rawDescData = file_boards_proto_rawDesc
)

func file_boards_proto_rawDescGZIP() []byte {
	file_boards_proto_rawDescOnce.Do(func() {
		file_boards_proto_rawDescData = protoimpl.X.CompressGZIP(file_boards_proto_rawDescData)
	})
	return file_boards_proto_rawDescData
}

var file_boards_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_boards_proto_goTypes = []interface{}{
	(*DeleteBoardRequest)(nil),      // 0: boards.DeleteBoardRequest
	(*GetIdeasInBoardRequest)(nil),  // 1: boards.GetIdeasInBoardRequest
	(*GetIdeasInBoardResponse)(nil), // 2: boards.GetIdeasInBoardResponse
	(*GetAllBoardsRequest)(nil),     // 3: boards.GetAllBoardsRequest
	(*SetIdeaSavedRequest)(nil),     // 4: boards.SetIdeaSavedRequest
	(*CreateBoardRequest)(nil),      // 5: boards.CreateBoardRequest
	(*CreateBoardResponse)(nil),     // 6: boards.CreateBoardResponse
	(*GetBoardRequest)(nil),         // 7: boards.GetBoardRequest
	(*GetBoardResponse)(nil),        // 8: boards.GetBoardResponse
	(*BoardData)(nil),               // 9: boards.BoardData
	(*GetAllBoardsResponse)(nil),    // 10: boards.GetAllBoardsResponse
	(*IdeaData)(nil),                // 11: boards.IdeaData
	(*emptypb.Empty)(nil),           // 12: google.protobuf.Empty
}
var file_boards_proto_depIdxs = []int32{
	11, // 0: boards.GetIdeasInBoardResponse.ideas:type_name -> boards.IdeaData
	9,  // 1: boards.GetAllBoardsResponse.boards:type_name -> boards.BoardData
	5,  // 2: boards.Boards.CreateBoard:input_type -> boards.CreateBoardRequest
	7,  // 3: boards.Boards.GetBoard:input_type -> boards.GetBoardRequest
	3,  // 4: boards.Boards.GetAllBoards:input_type -> boards.GetAllBoardsRequest
	4,  // 5: boards.Boards.SetIdeaSaved:input_type -> boards.SetIdeaSavedRequest
	1,  // 6: boards.Boards.GetIdeasInBoard:input_type -> boards.GetIdeasInBoardRequest
	0,  // 7: boards.Boards.DeleteBoard:input_type -> boards.DeleteBoardRequest
	6,  // 8: boards.Boards.CreateBoard:output_type -> boards.CreateBoardResponse
	8,  // 9: boards.Boards.GetBoard:output_type -> boards.GetBoardResponse
	10, // 10: boards.Boards.GetAllBoards:output_type -> boards.GetAllBoardsResponse
	12, // 11: boards.Boards.SetIdeaSaved:output_type -> google.protobuf.Empty
	2,  // 12: boards.Boards.GetIdeasInBoard:output_type -> boards.GetIdeasInBoardResponse
	12, // 13: boards.Boards.DeleteBoard:output_type -> google.protobuf.Empty
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_boards_proto_init() }
func file_boards_proto_init() {
	if File_boards_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_boards_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBoardRequest); i {
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
		file_boards_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIdeasInBoardRequest); i {
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
		file_boards_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIdeasInBoardResponse); i {
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
		file_boards_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllBoardsRequest); i {
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
		file_boards_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetIdeaSavedRequest); i {
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
		file_boards_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBoardRequest); i {
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
		file_boards_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBoardResponse); i {
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
		file_boards_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBoardRequest); i {
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
		file_boards_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBoardResponse); i {
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
		file_boards_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoardData); i {
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
		file_boards_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllBoardsResponse); i {
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
		file_boards_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_boards_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_boards_proto_goTypes,
		DependencyIndexes: file_boards_proto_depIdxs,
		MessageInfos:      file_boards_proto_msgTypes,
	}.Build()
	File_boards_proto = out.File
	file_boards_proto_rawDesc = nil
	file_boards_proto_goTypes = nil
	file_boards_proto_depIdxs = nil
}
