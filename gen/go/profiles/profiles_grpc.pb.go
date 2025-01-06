// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.29.1
// source: profiles.proto

package profilesv1

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

// ProfilesClient is the client API for Profiles service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfilesClient interface {
	CreateProfile(ctx context.Context, in *CreateProfileRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error)
	// ToggleSaveIdea saves the idea to the board if not saved and removes it otherwise
	ToggleSaveIdea(ctx context.Context, in *ToggleSaveRequest, opts ...grpc.CallOption) (*ToggleSaveResponse, error)
	IsIdeaSaved(ctx context.Context, in *IsIdeaSavedRequest, opts ...grpc.CallOption) (*IsIdeaSavedResponse, error)
	GetSavedIdeas(ctx context.Context, in *GetSavedIdeasRequest, opts ...grpc.CallOption) (*GetSavedIdeasResponse, error)
	GetSavedIdeasIds(ctx context.Context, in *GetSavedIdeasIdsRequest, opts ...grpc.CallOption) (*GetSavedIdeasIdsResponse, error)
	MoveIdeasToBoard(ctx context.Context, in *MoveIdeaToBoardRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type profilesClient struct {
	cc grpc.ClientConnInterface
}

func NewProfilesClient(cc grpc.ClientConnInterface) ProfilesClient {
	return &profilesClient{cc}
}

func (c *profilesClient) CreateProfile(ctx context.Context, in *CreateProfileRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/profiles.Profiles/CreateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilesClient) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error) {
	out := new(GetProfileResponse)
	err := c.cc.Invoke(ctx, "/profiles.Profiles/GetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilesClient) ToggleSaveIdea(ctx context.Context, in *ToggleSaveRequest, opts ...grpc.CallOption) (*ToggleSaveResponse, error) {
	out := new(ToggleSaveResponse)
	err := c.cc.Invoke(ctx, "/profiles.Profiles/ToggleSaveIdea", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilesClient) IsIdeaSaved(ctx context.Context, in *IsIdeaSavedRequest, opts ...grpc.CallOption) (*IsIdeaSavedResponse, error) {
	out := new(IsIdeaSavedResponse)
	err := c.cc.Invoke(ctx, "/profiles.Profiles/IsIdeaSaved", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilesClient) GetSavedIdeas(ctx context.Context, in *GetSavedIdeasRequest, opts ...grpc.CallOption) (*GetSavedIdeasResponse, error) {
	out := new(GetSavedIdeasResponse)
	err := c.cc.Invoke(ctx, "/profiles.Profiles/GetSavedIdeas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilesClient) GetSavedIdeasIds(ctx context.Context, in *GetSavedIdeasIdsRequest, opts ...grpc.CallOption) (*GetSavedIdeasIdsResponse, error) {
	out := new(GetSavedIdeasIdsResponse)
	err := c.cc.Invoke(ctx, "/profiles.Profiles/GetSavedIdeasIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilesClient) MoveIdeasToBoard(ctx context.Context, in *MoveIdeaToBoardRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/profiles.Profiles/MoveIdeasToBoard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfilesServer is the server API for Profiles service.
// All implementations must embed UnimplementedProfilesServer
// for forward compatibility
type ProfilesServer interface {
	CreateProfile(context.Context, *CreateProfileRequest) (*emptypb.Empty, error)
	GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error)
	// ToggleSaveIdea saves the idea to the board if not saved and removes it otherwise
	ToggleSaveIdea(context.Context, *ToggleSaveRequest) (*ToggleSaveResponse, error)
	IsIdeaSaved(context.Context, *IsIdeaSavedRequest) (*IsIdeaSavedResponse, error)
	GetSavedIdeas(context.Context, *GetSavedIdeasRequest) (*GetSavedIdeasResponse, error)
	GetSavedIdeasIds(context.Context, *GetSavedIdeasIdsRequest) (*GetSavedIdeasIdsResponse, error)
	MoveIdeasToBoard(context.Context, *MoveIdeaToBoardRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedProfilesServer()
}

// UnimplementedProfilesServer must be embedded to have forward compatible implementations.
type UnimplementedProfilesServer struct {
}

func (UnimplementedProfilesServer) CreateProfile(context.Context, *CreateProfileRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProfile not implemented")
}
func (UnimplementedProfilesServer) GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedProfilesServer) ToggleSaveIdea(context.Context, *ToggleSaveRequest) (*ToggleSaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToggleSaveIdea not implemented")
}
func (UnimplementedProfilesServer) IsIdeaSaved(context.Context, *IsIdeaSavedRequest) (*IsIdeaSavedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsIdeaSaved not implemented")
}
func (UnimplementedProfilesServer) GetSavedIdeas(context.Context, *GetSavedIdeasRequest) (*GetSavedIdeasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSavedIdeas not implemented")
}
func (UnimplementedProfilesServer) GetSavedIdeasIds(context.Context, *GetSavedIdeasIdsRequest) (*GetSavedIdeasIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSavedIdeasIds not implemented")
}
func (UnimplementedProfilesServer) MoveIdeasToBoard(context.Context, *MoveIdeaToBoardRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveIdeasToBoard not implemented")
}
func (UnimplementedProfilesServer) mustEmbedUnimplementedProfilesServer() {}

// UnsafeProfilesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfilesServer will
// result in compilation errors.
type UnsafeProfilesServer interface {
	mustEmbedUnimplementedProfilesServer()
}

func RegisterProfilesServer(s grpc.ServiceRegistrar, srv ProfilesServer) {
	s.RegisterService(&Profiles_ServiceDesc, srv)
}

func _Profiles_CreateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilesServer).CreateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profiles.Profiles/CreateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilesServer).CreateProfile(ctx, req.(*CreateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profiles_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilesServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profiles.Profiles/GetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilesServer).GetProfile(ctx, req.(*GetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profiles_ToggleSaveIdea_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToggleSaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilesServer).ToggleSaveIdea(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profiles.Profiles/ToggleSaveIdea",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilesServer).ToggleSaveIdea(ctx, req.(*ToggleSaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profiles_IsIdeaSaved_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsIdeaSavedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilesServer).IsIdeaSaved(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profiles.Profiles/IsIdeaSaved",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilesServer).IsIdeaSaved(ctx, req.(*IsIdeaSavedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profiles_GetSavedIdeas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSavedIdeasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilesServer).GetSavedIdeas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profiles.Profiles/GetSavedIdeas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilesServer).GetSavedIdeas(ctx, req.(*GetSavedIdeasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profiles_GetSavedIdeasIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSavedIdeasIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilesServer).GetSavedIdeasIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profiles.Profiles/GetSavedIdeasIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilesServer).GetSavedIdeasIds(ctx, req.(*GetSavedIdeasIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profiles_MoveIdeasToBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveIdeaToBoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilesServer).MoveIdeasToBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profiles.Profiles/MoveIdeasToBoard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilesServer).MoveIdeasToBoard(ctx, req.(*MoveIdeaToBoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Profiles_ServiceDesc is the grpc.ServiceDesc for Profiles service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Profiles_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "profiles.Profiles",
	HandlerType: (*ProfilesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProfile",
			Handler:    _Profiles_CreateProfile_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _Profiles_GetProfile_Handler,
		},
		{
			MethodName: "ToggleSaveIdea",
			Handler:    _Profiles_ToggleSaveIdea_Handler,
		},
		{
			MethodName: "IsIdeaSaved",
			Handler:    _Profiles_IsIdeaSaved_Handler,
		},
		{
			MethodName: "GetSavedIdeas",
			Handler:    _Profiles_GetSavedIdeas_Handler,
		},
		{
			MethodName: "GetSavedIdeasIds",
			Handler:    _Profiles_GetSavedIdeasIds_Handler,
		},
		{
			MethodName: "MoveIdeasToBoard",
			Handler:    _Profiles_MoveIdeasToBoard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profiles.proto",
}
