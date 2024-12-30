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

type CreateBoardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateBoardRequest) Reset() {
	*x = CreateBoardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boards_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBoardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBoardRequest) ProtoMessage() {}

func (x *CreateBoardRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateBoardRequest.ProtoReflect.Descriptor instead.
func (*CreateBoardRequest) Descriptor() ([]byte, []int) {
	return file_boards_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBoardRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
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
		mi := &file_boards_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBoardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBoardResponse) ProtoMessage() {}

func (x *CreateBoardResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateBoardResponse.ProtoReflect.Descriptor instead.
func (*CreateBoardResponse) Descriptor() ([]byte, []int) {
	return file_boards_proto_rawDescGZIP(), []int{1}
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
		mi := &file_boards_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBoardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBoardRequest) ProtoMessage() {}

func (x *GetBoardRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetBoardRequest.ProtoReflect.Descriptor instead.
func (*GetBoardRequest) Descriptor() ([]byte, []int) {
	return file_boards_proto_rawDescGZIP(), []int{2}
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
		mi := &file_boards_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBoardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBoardResponse) ProtoMessage() {}

func (x *GetBoardResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetBoardResponse.ProtoReflect.Descriptor instead.
func (*GetBoardResponse) Descriptor() ([]byte, []int) {
	return file_boards_proto_rawDescGZIP(), []int{3}
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
		mi := &file_boards_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoardData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoardData) ProtoMessage() {}

func (x *BoardData) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use BoardData.ProtoReflect.Descriptor instead.
func (*BoardData) Descriptor() ([]byte, []int) {
	return file_boards_proto_rawDescGZIP(), []int{4}
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
		mi := &file_boards_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllBoardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllBoardsResponse) ProtoMessage() {}

func (x *GetAllBoardsResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetAllBoardsResponse.ProtoReflect.Descriptor instead.
func (*GetAllBoardsResponse) Descriptor() ([]byte, []int) {
	return file_boards_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllBoardsResponse) GetBoards() []*BoardData {
	if x != nil {
		return x.Boards
	}
	return nil
}

var File_boards_proto protoreflect.FileDescriptor

var file_boards_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x25, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x53, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x6f,
	0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x64, 0x65, 0x61, 0x73, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x03, 0x52, 0x08, 0x69, 0x64, 0x65, 0x61, 0x73, 0x49, 0x64, 0x73, 0x22, 0x4c, 0x0a, 0x09,
	0x42, 0x6f, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x64, 0x65, 0x61, 0x73, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x08, 0x69, 0x64, 0x65, 0x61, 0x73, 0x49, 0x64, 0x73, 0x22, 0x41, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x42, 0x6f, 0x61, 0x72,
	0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x32, 0xd5, 0x01,
	0x0a, 0x06, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x12, 0x46, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x1a, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3d, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x17, 0x2e, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x44, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1f, 0x5a, 0x1d, 0x69, 0x64, 0x65, 0x61, 0x2d, 0x73, 0x74,
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

var file_boards_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_boards_proto_goTypes = []interface{}{
	(*CreateBoardRequest)(nil),   // 0: boards.CreateBoardRequest
	(*CreateBoardResponse)(nil),  // 1: boards.CreateBoardResponse
	(*GetBoardRequest)(nil),      // 2: boards.GetBoardRequest
	(*GetBoardResponse)(nil),     // 3: boards.GetBoardResponse
	(*BoardData)(nil),            // 4: boards.BoardData
	(*GetAllBoardsResponse)(nil), // 5: boards.GetAllBoardsResponse
	(*emptypb.Empty)(nil),        // 6: google.protobuf.Empty
}
var file_boards_proto_depIdxs = []int32{
	4, // 0: boards.GetAllBoardsResponse.boards:type_name -> boards.BoardData
	0, // 1: boards.Boards.CreateBoard:input_type -> boards.CreateBoardRequest
	2, // 2: boards.Boards.GetBoard:input_type -> boards.GetBoardRequest
	6, // 3: boards.Boards.GetAllBoards:input_type -> google.protobuf.Empty
	1, // 4: boards.Boards.CreateBoard:output_type -> boards.CreateBoardResponse
	3, // 5: boards.Boards.GetBoard:output_type -> boards.GetBoardResponse
	5, // 6: boards.Boards.GetAllBoards:output_type -> boards.GetAllBoardsResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_boards_proto_init() }
func file_boards_proto_init() {
	if File_boards_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_boards_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_boards_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_boards_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_boards_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_boards_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_boards_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_boards_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
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