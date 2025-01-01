// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v5.29.1
// source: profiles.proto

package profilesv1

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

type GetSavedIdeasIdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetSavedIdeasIdsRequest) Reset() {
	*x = GetSavedIdeasIdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiles_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSavedIdeasIdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSavedIdeasIdsRequest) ProtoMessage() {}

func (x *GetSavedIdeasIdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profiles_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSavedIdeasIdsRequest.ProtoReflect.Descriptor instead.
func (*GetSavedIdeasIdsRequest) Descriptor() ([]byte, []int) {
	return file_profiles_proto_rawDescGZIP(), []int{0}
}

func (x *GetSavedIdeasIdsRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetSavedIdeasIdsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdeasIds []int64 `protobuf:"varint,1,rep,packed,name=ideas_ids,json=ideasIds,proto3" json:"ideas_ids,omitempty"`
}

func (x *GetSavedIdeasIdsResponse) Reset() {
	*x = GetSavedIdeasIdsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiles_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSavedIdeasIdsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSavedIdeasIdsResponse) ProtoMessage() {}

func (x *GetSavedIdeasIdsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_profiles_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSavedIdeasIdsResponse.ProtoReflect.Descriptor instead.
func (*GetSavedIdeasIdsResponse) Descriptor() ([]byte, []int) {
	return file_profiles_proto_rawDescGZIP(), []int{1}
}

func (x *GetSavedIdeasIdsResponse) GetIdeasIds() []int64 {
	if x != nil {
		return x.IdeasIds
	}
	return nil
}

type ToggleSaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	IdeaId  int64 `protobuf:"varint,2,opt,name=idea_id,json=ideaId,proto3" json:"idea_id,omitempty"`
	BoardId int64 `protobuf:"varint,3,opt,name=board_id,json=boardId,proto3" json:"board_id,omitempty"`
}

func (x *ToggleSaveRequest) Reset() {
	*x = ToggleSaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiles_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToggleSaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToggleSaveRequest) ProtoMessage() {}

func (x *ToggleSaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profiles_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToggleSaveRequest.ProtoReflect.Descriptor instead.
func (*ToggleSaveRequest) Descriptor() ([]byte, []int) {
	return file_profiles_proto_rawDescGZIP(), []int{2}
}

func (x *ToggleSaveRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ToggleSaveRequest) GetIdeaId() int64 {
	if x != nil {
		return x.IdeaId
	}
	return 0
}

func (x *ToggleSaveRequest) GetBoardId() int64 {
	if x != nil {
		return x.BoardId
	}
	return 0
}

type ToggleSaveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NowSaved bool `protobuf:"varint,1,opt,name=now_saved,json=nowSaved,proto3" json:"now_saved,omitempty"`
}

func (x *ToggleSaveResponse) Reset() {
	*x = ToggleSaveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiles_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToggleSaveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToggleSaveResponse) ProtoMessage() {}

func (x *ToggleSaveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_profiles_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToggleSaveResponse.ProtoReflect.Descriptor instead.
func (*ToggleSaveResponse) Descriptor() ([]byte, []int) {
	return file_profiles_proto_rawDescGZIP(), []int{3}
}

func (x *ToggleSaveResponse) GetNowSaved() bool {
	if x != nil {
		return x.NowSaved
	}
	return false
}

type CreateProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateProfileRequest) Reset() {
	*x = CreateProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiles_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProfileRequest) ProtoMessage() {}

func (x *CreateProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profiles_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProfileRequest.ProtoReflect.Descriptor instead.
func (*CreateProfileRequest) Descriptor() ([]byte, []int) {
	return file_profiles_proto_rawDescGZIP(), []int{4}
}

func (x *CreateProfileRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateProfileRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateProfileRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ProfileData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email       string  `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	AvatarImage string  `protobuf:"bytes,3,opt,name=avatarImage,proto3" json:"avatarImage,omitempty"`
	Name        string  `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Link        string  `protobuf:"bytes,6,opt,name=link,proto3" json:"link,omitempty"`
	Boards      []int64 `protobuf:"varint,7,rep,packed,name=boards,proto3" json:"boards,omitempty"`
	SavedIdeas  []int64 `protobuf:"varint,8,rep,packed,name=saved_ideas,json=savedIdeas,proto3" json:"saved_ideas,omitempty"`
}

func (x *ProfileData) Reset() {
	*x = ProfileData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiles_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileData) ProtoMessage() {}

func (x *ProfileData) ProtoReflect() protoreflect.Message {
	mi := &file_profiles_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileData.ProtoReflect.Descriptor instead.
func (*ProfileData) Descriptor() ([]byte, []int) {
	return file_profiles_proto_rawDescGZIP(), []int{5}
}

func (x *ProfileData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProfileData) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ProfileData) GetAvatarImage() string {
	if x != nil {
		return x.AvatarImage
	}
	return ""
}

func (x *ProfileData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProfileData) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProfileData) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *ProfileData) GetBoards() []int64 {
	if x != nil {
		return x.Boards
	}
	return nil
}

func (x *ProfileData) GetSavedIdeas() []int64 {
	if x != nil {
		return x.SavedIdeas
	}
	return nil
}

type GetProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetProfileRequest) Reset() {
	*x = GetProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiles_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileRequest) ProtoMessage() {}

func (x *GetProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profiles_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileRequest.ProtoReflect.Descriptor instead.
func (*GetProfileRequest) Descriptor() ([]byte, []int) {
	return file_profiles_proto_rawDescGZIP(), []int{6}
}

func (x *GetProfileRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *ProfileData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetProfileResponse) Reset() {
	*x = GetProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiles_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileResponse) ProtoMessage() {}

func (x *GetProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_profiles_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileResponse.ProtoReflect.Descriptor instead.
func (*GetProfileResponse) Descriptor() ([]byte, []int) {
	return file_profiles_proto_rawDescGZIP(), []int{7}
}

func (x *GetProfileResponse) GetData() *ProfileData {
	if x != nil {
		return x.Data
	}
	return nil
}

type IsIdeaSavedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	IdeaId int64 `protobuf:"varint,2,opt,name=idea_id,json=ideaId,proto3" json:"idea_id,omitempty"`
}

func (x *IsIdeaSavedRequest) Reset() {
	*x = IsIdeaSavedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiles_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsIdeaSavedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsIdeaSavedRequest) ProtoMessage() {}

func (x *IsIdeaSavedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profiles_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsIdeaSavedRequest.ProtoReflect.Descriptor instead.
func (*IsIdeaSavedRequest) Descriptor() ([]byte, []int) {
	return file_profiles_proto_rawDescGZIP(), []int{8}
}

func (x *IsIdeaSavedRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *IsIdeaSavedRequest) GetIdeaId() int64 {
	if x != nil {
		return x.IdeaId
	}
	return 0
}

type IsIdeaSavedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Saved bool `protobuf:"varint,1,opt,name=saved,proto3" json:"saved,omitempty"`
}

func (x *IsIdeaSavedResponse) Reset() {
	*x = IsIdeaSavedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiles_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsIdeaSavedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsIdeaSavedResponse) ProtoMessage() {}

func (x *IsIdeaSavedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_profiles_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsIdeaSavedResponse.ProtoReflect.Descriptor instead.
func (*IsIdeaSavedResponse) Descriptor() ([]byte, []int) {
	return file_profiles_proto_rawDescGZIP(), []int{9}
}

func (x *IsIdeaSavedResponse) GetSaved() bool {
	if x != nil {
		return x.Saved
	}
	return false
}

type GetSavedIdeasRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetSavedIdeasRequest) Reset() {
	*x = GetSavedIdeasRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiles_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSavedIdeasRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSavedIdeasRequest) ProtoMessage() {}

func (x *GetSavedIdeasRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profiles_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSavedIdeasRequest.ProtoReflect.Descriptor instead.
func (*GetSavedIdeasRequest) Descriptor() ([]byte, []int) {
	return file_profiles_proto_rawDescGZIP(), []int{10}
}

func (x *GetSavedIdeasRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetSavedIdeasResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ideas []*IdeaData `protobuf:"bytes,1,rep,name=ideas,proto3" json:"ideas,omitempty"`
}

func (x *GetSavedIdeasResponse) Reset() {
	*x = GetSavedIdeasResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiles_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSavedIdeasResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSavedIdeasResponse) ProtoMessage() {}

func (x *GetSavedIdeasResponse) ProtoReflect() protoreflect.Message {
	mi := &file_profiles_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSavedIdeasResponse.ProtoReflect.Descriptor instead.
func (*GetSavedIdeasResponse) Descriptor() ([]byte, []int) {
	return file_profiles_proto_rawDescGZIP(), []int{11}
}

func (x *GetSavedIdeasResponse) GetIdeas() []*IdeaData {
	if x != nil {
		return x.Ideas
	}
	return nil
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
}

func (x *IdeaData) Reset() {
	*x = IdeaData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiles_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdeaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdeaData) ProtoMessage() {}

func (x *IdeaData) ProtoReflect() protoreflect.Message {
	mi := &file_profiles_proto_msgTypes[12]
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
	return file_profiles_proto_rawDescGZIP(), []int{12}
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

var File_profiles_proto protoreflect.FileDescriptor

var file_profiles_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x53, 0x61,
	0x76, 0x65, 0x64, 0x49, 0x64, 0x65, 0x61, 0x73, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x37, 0x0a, 0x18, 0x47,
	0x65, 0x74, 0x53, 0x61, 0x76, 0x65, 0x64, 0x49, 0x64, 0x65, 0x61, 0x73, 0x49, 0x64, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x64, 0x65, 0x61, 0x73,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08, 0x69, 0x64, 0x65, 0x61,
	0x73, 0x49, 0x64, 0x73, 0x22, 0x60, 0x0a, 0x11, 0x54, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x53, 0x61,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x64, 0x65, 0x61, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x12, 0x54, 0x6f, 0x67, 0x67, 0x6c, 0x65,
	0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6e, 0x6f, 0x77, 0x5f, 0x73, 0x61, 0x76, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x6e, 0x6f, 0x77, 0x53, 0x61, 0x76, 0x65, 0x64, 0x22, 0x50, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xd8, 0x01, 0x0a, 0x0b,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e,
	0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x16, 0x0a,
	0x06, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x61, 0x76, 0x65, 0x64, 0x5f, 0x69,
	0x64, 0x65, 0x61, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x61, 0x76, 0x65,
	0x64, 0x49, 0x64, 0x65, 0x61, 0x73, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x46, 0x0a, 0x12,
	0x49, 0x73, 0x49, 0x64, 0x65, 0x61, 0x53, 0x61, 0x76, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69,
	0x64, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x64,
	0x65, 0x61, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x13, 0x49, 0x73, 0x49, 0x64, 0x65, 0x61, 0x53, 0x61,
	0x76, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x61, 0x76, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x73, 0x61, 0x76, 0x65,
	0x64, 0x22, 0x2f, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x61, 0x76, 0x65, 0x64, 0x49, 0x64, 0x65,
	0x61, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x41, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x61, 0x76, 0x65, 0x64, 0x49, 0x64,
	0x65, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x69,
	0x64, 0x65, 0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x49, 0x64, 0x65, 0x61, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05,
	0x69, 0x64, 0x65, 0x61, 0x73, 0x22, 0x97, 0x01, 0x0a, 0x08, 0x49, 0x64, 0x65, 0x61, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x64, 0x65, 0x61, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x32,
	0xe4, 0x03, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x49, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1e, 0x2e,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4b, 0x0a, 0x0e, 0x54, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x53, 0x61, 0x76, 0x65, 0x49, 0x64,
	0x65, 0x61, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x54, 0x6f,
	0x67, 0x67, 0x6c, 0x65, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x54, 0x6f, 0x67, 0x67, 0x6c,
	0x65, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a,
	0x0b, 0x49, 0x73, 0x49, 0x64, 0x65, 0x61, 0x53, 0x61, 0x76, 0x65, 0x64, 0x12, 0x1c, 0x2e, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x49, 0x73, 0x49, 0x64, 0x65, 0x61, 0x53, 0x61,
	0x76, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x49, 0x73, 0x49, 0x64, 0x65, 0x61, 0x53, 0x61, 0x76, 0x65,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x53, 0x61, 0x76, 0x65, 0x64, 0x49, 0x64, 0x65, 0x61, 0x73, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x61, 0x76, 0x65, 0x64, 0x49, 0x64,
	0x65, 0x61, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x61, 0x76, 0x65, 0x64, 0x49, 0x64,
	0x65, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x53, 0x61, 0x76, 0x65, 0x64, 0x49, 0x64, 0x65, 0x61, 0x73, 0x49, 0x64, 0x73, 0x12,
	0x21, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x61,
	0x76, 0x65, 0x64, 0x49, 0x64, 0x65, 0x61, 0x73, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x61, 0x76, 0x65, 0x64, 0x49, 0x64, 0x65, 0x61, 0x73, 0x49, 0x64, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x23, 0x5a, 0x21, 0x69, 0x64, 0x65, 0x61, 0x2d, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x3b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_profiles_proto_rawDescOnce sync.Once
	file_profiles_proto_rawDescData = file_profiles_proto_rawDesc
)

func file_profiles_proto_rawDescGZIP() []byte {
	file_profiles_proto_rawDescOnce.Do(func() {
		file_profiles_proto_rawDescData = protoimpl.X.CompressGZIP(file_profiles_proto_rawDescData)
	})
	return file_profiles_proto_rawDescData
}

var file_profiles_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_profiles_proto_goTypes = []interface{}{
	(*GetSavedIdeasIdsRequest)(nil),  // 0: profiles.GetSavedIdeasIdsRequest
	(*GetSavedIdeasIdsResponse)(nil), // 1: profiles.GetSavedIdeasIdsResponse
	(*ToggleSaveRequest)(nil),        // 2: profiles.ToggleSaveRequest
	(*ToggleSaveResponse)(nil),       // 3: profiles.ToggleSaveResponse
	(*CreateProfileRequest)(nil),     // 4: profiles.CreateProfileRequest
	(*ProfileData)(nil),              // 5: profiles.ProfileData
	(*GetProfileRequest)(nil),        // 6: profiles.GetProfileRequest
	(*GetProfileResponse)(nil),       // 7: profiles.GetProfileResponse
	(*IsIdeaSavedRequest)(nil),       // 8: profiles.IsIdeaSavedRequest
	(*IsIdeaSavedResponse)(nil),      // 9: profiles.IsIdeaSavedResponse
	(*GetSavedIdeasRequest)(nil),     // 10: profiles.GetSavedIdeasRequest
	(*GetSavedIdeasResponse)(nil),    // 11: profiles.GetSavedIdeasResponse
	(*IdeaData)(nil),                 // 12: profiles.IdeaData
	(*emptypb.Empty)(nil),            // 13: google.protobuf.Empty
}
var file_profiles_proto_depIdxs = []int32{
	5,  // 0: profiles.GetProfileResponse.data:type_name -> profiles.ProfileData
	12, // 1: profiles.GetSavedIdeasResponse.ideas:type_name -> profiles.IdeaData
	4,  // 2: profiles.Profiles.CreateProfile:input_type -> profiles.CreateProfileRequest
	6,  // 3: profiles.Profiles.GetProfile:input_type -> profiles.GetProfileRequest
	2,  // 4: profiles.Profiles.ToggleSaveIdea:input_type -> profiles.ToggleSaveRequest
	8,  // 5: profiles.Profiles.IsIdeaSaved:input_type -> profiles.IsIdeaSavedRequest
	10, // 6: profiles.Profiles.GetSavedIdeas:input_type -> profiles.GetSavedIdeasRequest
	0,  // 7: profiles.Profiles.GetSavedIdeasIds:input_type -> profiles.GetSavedIdeasIdsRequest
	13, // 8: profiles.Profiles.CreateProfile:output_type -> google.protobuf.Empty
	7,  // 9: profiles.Profiles.GetProfile:output_type -> profiles.GetProfileResponse
	3,  // 10: profiles.Profiles.ToggleSaveIdea:output_type -> profiles.ToggleSaveResponse
	9,  // 11: profiles.Profiles.IsIdeaSaved:output_type -> profiles.IsIdeaSavedResponse
	11, // 12: profiles.Profiles.GetSavedIdeas:output_type -> profiles.GetSavedIdeasResponse
	1,  // 13: profiles.Profiles.GetSavedIdeasIds:output_type -> profiles.GetSavedIdeasIdsResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_profiles_proto_init() }
func file_profiles_proto_init() {
	if File_profiles_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_profiles_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSavedIdeasIdsRequest); i {
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
		file_profiles_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSavedIdeasIdsResponse); i {
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
		file_profiles_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToggleSaveRequest); i {
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
		file_profiles_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToggleSaveResponse); i {
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
		file_profiles_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProfileRequest); i {
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
		file_profiles_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileData); i {
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
		file_profiles_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfileRequest); i {
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
		file_profiles_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfileResponse); i {
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
		file_profiles_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsIdeaSavedRequest); i {
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
		file_profiles_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsIdeaSavedResponse); i {
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
		file_profiles_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSavedIdeasRequest); i {
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
		file_profiles_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSavedIdeasResponse); i {
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
		file_profiles_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_profiles_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_profiles_proto_goTypes,
		DependencyIndexes: file_profiles_proto_depIdxs,
		MessageInfos:      file_profiles_proto_msgTypes,
	}.Build()
	File_profiles_proto = out.File
	file_profiles_proto_rawDesc = nil
	file_profiles_proto_goTypes = nil
	file_profiles_proto_depIdxs = nil
}
