// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.29.1
// source: boards.proto

package boardsv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BoardsClient is the client API for Boards service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BoardsClient interface {
	CreateBoard(ctx context.Context, in *CreateBoardRequest, opts ...grpc.CallOption) (*CreateBoardResponse, error)
	GetBoard(ctx context.Context, in *GetBoardRequest, opts ...grpc.CallOption) (*GetBoardResponse, error)
	GetAllBoards(ctx context.Context, in *GetAllBoardsRequest, opts ...grpc.CallOption) (*GetAllBoardsResponse, error)
	SetIdeaSaved(ctx context.Context, in *SetIdeaSavedRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetIdeasInBoard(ctx context.Context, in *GetIdeasInBoardRequest, opts ...grpc.CallOption) (*GetIdeasInBoardResponse, error)
}

type boardsClient struct {
	cc grpc.ClientConnInterface
}

func NewBoardsClient(cc grpc.ClientConnInterface) BoardsClient {
	return &boardsClient{cc}
}

func (c *boardsClient) CreateBoard(ctx context.Context, in *CreateBoardRequest, opts ...grpc.CallOption) (*CreateBoardResponse, error) {
	out := new(CreateBoardResponse)
	err := c.cc.Invoke(ctx, "/boards.Boards/CreateBoard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardsClient) GetBoard(ctx context.Context, in *GetBoardRequest, opts ...grpc.CallOption) (*GetBoardResponse, error) {
	out := new(GetBoardResponse)
	err := c.cc.Invoke(ctx, "/boards.Boards/GetBoard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardsClient) GetAllBoards(ctx context.Context, in *GetAllBoardsRequest, opts ...grpc.CallOption) (*GetAllBoardsResponse, error) {
	out := new(GetAllBoardsResponse)
	err := c.cc.Invoke(ctx, "/boards.Boards/GetAllBoards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardsClient) SetIdeaSaved(ctx context.Context, in *SetIdeaSavedRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/boards.Boards/SetIdeaSaved", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardsClient) GetIdeasInBoard(ctx context.Context, in *GetIdeasInBoardRequest, opts ...grpc.CallOption) (*GetIdeasInBoardResponse, error) {
	out := new(GetIdeasInBoardResponse)
	err := c.cc.Invoke(ctx, "/boards.Boards/GetIdeasInBoard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BoardsServer is the server API for Boards service.
// All implementations must embed UnimplementedBoardsServer
// for forward compatibility
type BoardsServer interface {
	CreateBoard(context.Context, *CreateBoardRequest) (*CreateBoardResponse, error)
	GetBoard(context.Context, *GetBoardRequest) (*GetBoardResponse, error)
	GetAllBoards(context.Context, *GetAllBoardsRequest) (*GetAllBoardsResponse, error)
	SetIdeaSaved(context.Context, *SetIdeaSavedRequest) (*emptypb.Empty, error)
	GetIdeasInBoard(context.Context, *GetIdeasInBoardRequest) (*GetIdeasInBoardResponse, error)
	mustEmbedUnimplementedBoardsServer()
}

// UnimplementedBoardsServer must be embedded to have forward compatible implementations.
type UnimplementedBoardsServer struct {
}

func (UnimplementedBoardsServer) CreateBoard(context.Context, *CreateBoardRequest) (*CreateBoardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBoard not implemented")
}
func (UnimplementedBoardsServer) GetBoard(context.Context, *GetBoardRequest) (*GetBoardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBoard not implemented")
}
func (UnimplementedBoardsServer) GetAllBoards(context.Context, *GetAllBoardsRequest) (*GetAllBoardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllBoards not implemented")
}
func (UnimplementedBoardsServer) SetIdeaSaved(context.Context, *SetIdeaSavedRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetIdeaSaved not implemented")
}
func (UnimplementedBoardsServer) GetIdeasInBoard(context.Context, *GetIdeasInBoardRequest) (*GetIdeasInBoardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIdeasInBoard not implemented")
}
func (UnimplementedBoardsServer) mustEmbedUnimplementedBoardsServer() {}

// UnsafeBoardsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BoardsServer will
// result in compilation errors.
type UnsafeBoardsServer interface {
	mustEmbedUnimplementedBoardsServer()
}

func RegisterBoardsServer(s grpc.ServiceRegistrar, srv BoardsServer) {
	s.RegisterService(&Boards_ServiceDesc, srv)
}

func _Boards_CreateBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardsServer).CreateBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/boards.Boards/CreateBoard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardsServer).CreateBoard(ctx, req.(*CreateBoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Boards_GetBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardsServer).GetBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/boards.Boards/GetBoard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardsServer).GetBoard(ctx, req.(*GetBoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Boards_GetAllBoards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllBoardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardsServer).GetAllBoards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/boards.Boards/GetAllBoards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardsServer).GetAllBoards(ctx, req.(*GetAllBoardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Boards_SetIdeaSaved_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetIdeaSavedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardsServer).SetIdeaSaved(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/boards.Boards/SetIdeaSaved",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardsServer).SetIdeaSaved(ctx, req.(*SetIdeaSavedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Boards_GetIdeasInBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIdeasInBoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardsServer).GetIdeasInBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/boards.Boards/GetIdeasInBoard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardsServer).GetIdeasInBoard(ctx, req.(*GetIdeasInBoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Boards_ServiceDesc is the grpc.ServiceDesc for Boards service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Boards_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "boards.Boards",
	HandlerType: (*BoardsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBoard",
			Handler:    _Boards_CreateBoard_Handler,
		},
		{
			MethodName: "GetBoard",
			Handler:    _Boards_GetBoard_Handler,
		},
		{
			MethodName: "GetAllBoards",
			Handler:    _Boards_GetAllBoards_Handler,
		},
		{
			MethodName: "SetIdeaSaved",
			Handler:    _Boards_SetIdeaSaved_Handler,
		},
		{
			MethodName: "GetIdeasInBoard",
			Handler:    _Boards_GetIdeasInBoard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "boards.proto",
}
