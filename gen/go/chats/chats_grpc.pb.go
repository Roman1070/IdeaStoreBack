// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.29.1
// source: chats.proto

package chatsv1

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

// ChatsClient is the client API for Chats service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatsClient interface {
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesResponse, error)
	CreateChat(ctx context.Context, in *CreateChatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetUsersChats(ctx context.Context, in *GetUsersChatsRequest, opts ...grpc.CallOption) (*GetUsersChatsResponse, error)
	DeleteChat(ctx context.Context, in *DeleteChatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type chatsClient struct {
	cc grpc.ClientConnInterface
}

func NewChatsClient(cc grpc.ClientConnInterface) ChatsClient {
	return &chatsClient{cc}
}

func (c *chatsClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/chats.Chats/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatsClient) GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesResponse, error) {
	out := new(GetMessagesResponse)
	err := c.cc.Invoke(ctx, "/chats.Chats/GetMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatsClient) CreateChat(ctx context.Context, in *CreateChatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/chats.Chats/CreateChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatsClient) GetUsersChats(ctx context.Context, in *GetUsersChatsRequest, opts ...grpc.CallOption) (*GetUsersChatsResponse, error) {
	out := new(GetUsersChatsResponse)
	err := c.cc.Invoke(ctx, "/chats.Chats/GetUsersChats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatsClient) DeleteChat(ctx context.Context, in *DeleteChatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/chats.Chats/DeleteChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatsServer is the server API for Chats service.
// All implementations must embed UnimplementedChatsServer
// for forward compatibility
type ChatsServer interface {
	SendMessage(context.Context, *SendMessageRequest) (*emptypb.Empty, error)
	GetMessages(context.Context, *GetMessagesRequest) (*GetMessagesResponse, error)
	CreateChat(context.Context, *CreateChatRequest) (*emptypb.Empty, error)
	GetUsersChats(context.Context, *GetUsersChatsRequest) (*GetUsersChatsResponse, error)
	DeleteChat(context.Context, *DeleteChatRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedChatsServer()
}

// UnimplementedChatsServer must be embedded to have forward compatible implementations.
type UnimplementedChatsServer struct {
}

func (UnimplementedChatsServer) SendMessage(context.Context, *SendMessageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatsServer) GetMessages(context.Context, *GetMessagesRequest) (*GetMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}
func (UnimplementedChatsServer) CreateChat(context.Context, *CreateChatRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChat not implemented")
}
func (UnimplementedChatsServer) GetUsersChats(context.Context, *GetUsersChatsRequest) (*GetUsersChatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersChats not implemented")
}
func (UnimplementedChatsServer) DeleteChat(context.Context, *DeleteChatRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChat not implemented")
}
func (UnimplementedChatsServer) mustEmbedUnimplementedChatsServer() {}

// UnsafeChatsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatsServer will
// result in compilation errors.
type UnsafeChatsServer interface {
	mustEmbedUnimplementedChatsServer()
}

func RegisterChatsServer(s grpc.ServiceRegistrar, srv ChatsServer) {
	s.RegisterService(&Chats_ServiceDesc, srv)
}

func _Chats_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatsServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chats.Chats/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatsServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chats_GetMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatsServer).GetMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chats.Chats/GetMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatsServer).GetMessages(ctx, req.(*GetMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chats_CreateChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatsServer).CreateChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chats.Chats/CreateChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatsServer).CreateChat(ctx, req.(*CreateChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chats_GetUsersChats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersChatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatsServer).GetUsersChats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chats.Chats/GetUsersChats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatsServer).GetUsersChats(ctx, req.(*GetUsersChatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chats_DeleteChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatsServer).DeleteChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chats.Chats/DeleteChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatsServer).DeleteChat(ctx, req.(*DeleteChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Chats_ServiceDesc is the grpc.ServiceDesc for Chats service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chats_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chats.Chats",
	HandlerType: (*ChatsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _Chats_SendMessage_Handler,
		},
		{
			MethodName: "GetMessages",
			Handler:    _Chats_GetMessages_Handler,
		},
		{
			MethodName: "CreateChat",
			Handler:    _Chats_CreateChat_Handler,
		},
		{
			MethodName: "GetUsersChats",
			Handler:    _Chats_GetUsersChats_Handler,
		},
		{
			MethodName: "DeleteChat",
			Handler:    _Chats_DeleteChat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chats.proto",
}