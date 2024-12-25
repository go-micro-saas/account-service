// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.6
// source: api/account-service/v1/services/user_auth.service.v1.proto

package servicev1

import (
	context "context"
	resources "github.com/go-micro-saas/account-service/api/account-service/v1/resources"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SrvUserAuthV1_Ping_FullMethodName                 = "/saas.api.account.servicev1.SrvUserAuthV1/Ping"
	SrvUserAuthV1_SendPhoneSignupCode_FullMethodName  = "/saas.api.account.servicev1.SrvUserAuthV1/SendPhoneSignupCode"
	SrvUserAuthV1_SendEmailSignupCode_FullMethodName  = "/saas.api.account.servicev1.SrvUserAuthV1/SendEmailSignupCode"
	SrvUserAuthV1_SignupByEmail_FullMethodName        = "/saas.api.account.servicev1.SrvUserAuthV1/SignupByEmail"
	SrvUserAuthV1_SignupByPhone_FullMethodName        = "/saas.api.account.servicev1.SrvUserAuthV1/SignupByPhone"
	SrvUserAuthV1_LoginOrSignupByPhone_FullMethodName = "/saas.api.account.servicev1.SrvUserAuthV1/LoginOrSignupByPhone"
	SrvUserAuthV1_LoginOrSignupByEmail_FullMethodName = "/saas.api.account.servicev1.SrvUserAuthV1/LoginOrSignupByEmail"
	SrvUserAuthV1_RefreshToken_FullMethodName         = "/saas.api.account.servicev1.SrvUserAuthV1/RefreshToken"
	SrvUserAuthV1_LoginByEmail_FullMethodName         = "/saas.api.account.servicev1.SrvUserAuthV1/LoginByEmail"
	SrvUserAuthV1_LoginByPhone_FullMethodName         = "/saas.api.account.servicev1.SrvUserAuthV1/LoginByPhone"
	SrvUserAuthV1_LoginByOpenApi_FullMethodName       = "/saas.api.account.servicev1.SrvUserAuthV1/LoginByOpenApi"
)

// SrvUserAuthV1Client is the client API for SrvUserAuthV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SrvUserAuthV1Client interface {
	// 身份验证-Ping测试
	Ping(ctx context.Context, in *resources.PingReq, opts ...grpc.CallOption) (*resources.PingResp, error)
	// 身份验证-发送手机注册码
	SendPhoneSignupCode(ctx context.Context, in *resources.SendPhoneSignupCodeReq, opts ...grpc.CallOption) (*resources.SendSignupCodeResp, error)
	// 身份验证-发送邮箱注册码
	SendEmailSignupCode(ctx context.Context, in *resources.SendEmailSignupCodeReq, opts ...grpc.CallOption) (*resources.SendSignupCodeResp, error)
	// 身份验证-Email注册
	SignupByEmail(ctx context.Context, in *resources.SignupByEmailReq, opts ...grpc.CallOption) (*resources.LoginResp, error)
	// 身份验证-手机注册
	SignupByPhone(ctx context.Context, in *resources.SignupByPhoneReq, opts ...grpc.CallOption) (*resources.LoginResp, error)
	// 身份验证-手机登陆(自动注册)
	LoginOrSignupByPhone(ctx context.Context, in *resources.LoginOrSignupByPhoneReq, opts ...grpc.CallOption) (*resources.LoginResp, error)
	// 身份验证-邮箱登陆(自动注册)
	LoginOrSignupByEmail(ctx context.Context, in *resources.LoginOrSignupByEmailReq, opts ...grpc.CallOption) (*resources.LoginResp, error)
	// 身份验证-刷新Token
	RefreshToken(ctx context.Context, in *resources.RefreshTokenReq, opts ...grpc.CallOption) (*resources.LoginResp, error)
	// 身份验证-Email登录
	LoginByEmail(ctx context.Context, in *resources.LoginByEmailReq, opts ...grpc.CallOption) (*resources.LoginResp, error)
	// 身份验证-手机登录
	LoginByPhone(ctx context.Context, in *resources.LoginByPhoneReq, opts ...grpc.CallOption) (*resources.LoginResp, error)
	// 身份验证-开发平台登录
	LoginByOpenApi(ctx context.Context, in *resources.LoginByOpenApiReq, opts ...grpc.CallOption) (*resources.LoginResp, error)
}

type srvUserAuthV1Client struct {
	cc grpc.ClientConnInterface
}

func NewSrvUserAuthV1Client(cc grpc.ClientConnInterface) SrvUserAuthV1Client {
	return &srvUserAuthV1Client{cc}
}

func (c *srvUserAuthV1Client) Ping(ctx context.Context, in *resources.PingReq, opts ...grpc.CallOption) (*resources.PingResp, error) {
	out := new(resources.PingResp)
	err := c.cc.Invoke(ctx, SrvUserAuthV1_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvUserAuthV1Client) SendPhoneSignupCode(ctx context.Context, in *resources.SendPhoneSignupCodeReq, opts ...grpc.CallOption) (*resources.SendSignupCodeResp, error) {
	out := new(resources.SendSignupCodeResp)
	err := c.cc.Invoke(ctx, SrvUserAuthV1_SendPhoneSignupCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvUserAuthV1Client) SendEmailSignupCode(ctx context.Context, in *resources.SendEmailSignupCodeReq, opts ...grpc.CallOption) (*resources.SendSignupCodeResp, error) {
	out := new(resources.SendSignupCodeResp)
	err := c.cc.Invoke(ctx, SrvUserAuthV1_SendEmailSignupCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvUserAuthV1Client) SignupByEmail(ctx context.Context, in *resources.SignupByEmailReq, opts ...grpc.CallOption) (*resources.LoginResp, error) {
	out := new(resources.LoginResp)
	err := c.cc.Invoke(ctx, SrvUserAuthV1_SignupByEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvUserAuthV1Client) SignupByPhone(ctx context.Context, in *resources.SignupByPhoneReq, opts ...grpc.CallOption) (*resources.LoginResp, error) {
	out := new(resources.LoginResp)
	err := c.cc.Invoke(ctx, SrvUserAuthV1_SignupByPhone_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvUserAuthV1Client) LoginOrSignupByPhone(ctx context.Context, in *resources.LoginOrSignupByPhoneReq, opts ...grpc.CallOption) (*resources.LoginResp, error) {
	out := new(resources.LoginResp)
	err := c.cc.Invoke(ctx, SrvUserAuthV1_LoginOrSignupByPhone_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvUserAuthV1Client) LoginOrSignupByEmail(ctx context.Context, in *resources.LoginOrSignupByEmailReq, opts ...grpc.CallOption) (*resources.LoginResp, error) {
	out := new(resources.LoginResp)
	err := c.cc.Invoke(ctx, SrvUserAuthV1_LoginOrSignupByEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvUserAuthV1Client) RefreshToken(ctx context.Context, in *resources.RefreshTokenReq, opts ...grpc.CallOption) (*resources.LoginResp, error) {
	out := new(resources.LoginResp)
	err := c.cc.Invoke(ctx, SrvUserAuthV1_RefreshToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvUserAuthV1Client) LoginByEmail(ctx context.Context, in *resources.LoginByEmailReq, opts ...grpc.CallOption) (*resources.LoginResp, error) {
	out := new(resources.LoginResp)
	err := c.cc.Invoke(ctx, SrvUserAuthV1_LoginByEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvUserAuthV1Client) LoginByPhone(ctx context.Context, in *resources.LoginByPhoneReq, opts ...grpc.CallOption) (*resources.LoginResp, error) {
	out := new(resources.LoginResp)
	err := c.cc.Invoke(ctx, SrvUserAuthV1_LoginByPhone_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvUserAuthV1Client) LoginByOpenApi(ctx context.Context, in *resources.LoginByOpenApiReq, opts ...grpc.CallOption) (*resources.LoginResp, error) {
	out := new(resources.LoginResp)
	err := c.cc.Invoke(ctx, SrvUserAuthV1_LoginByOpenApi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SrvUserAuthV1Server is the server API for SrvUserAuthV1 service.
// All implementations must embed UnimplementedSrvUserAuthV1Server
// for forward compatibility
type SrvUserAuthV1Server interface {
	// 身份验证-Ping测试
	Ping(context.Context, *resources.PingReq) (*resources.PingResp, error)
	// 身份验证-发送手机注册码
	SendPhoneSignupCode(context.Context, *resources.SendPhoneSignupCodeReq) (*resources.SendSignupCodeResp, error)
	// 身份验证-发送邮箱注册码
	SendEmailSignupCode(context.Context, *resources.SendEmailSignupCodeReq) (*resources.SendSignupCodeResp, error)
	// 身份验证-Email注册
	SignupByEmail(context.Context, *resources.SignupByEmailReq) (*resources.LoginResp, error)
	// 身份验证-手机注册
	SignupByPhone(context.Context, *resources.SignupByPhoneReq) (*resources.LoginResp, error)
	// 身份验证-手机登陆(自动注册)
	LoginOrSignupByPhone(context.Context, *resources.LoginOrSignupByPhoneReq) (*resources.LoginResp, error)
	// 身份验证-邮箱登陆(自动注册)
	LoginOrSignupByEmail(context.Context, *resources.LoginOrSignupByEmailReq) (*resources.LoginResp, error)
	// 身份验证-刷新Token
	RefreshToken(context.Context, *resources.RefreshTokenReq) (*resources.LoginResp, error)
	// 身份验证-Email登录
	LoginByEmail(context.Context, *resources.LoginByEmailReq) (*resources.LoginResp, error)
	// 身份验证-手机登录
	LoginByPhone(context.Context, *resources.LoginByPhoneReq) (*resources.LoginResp, error)
	// 身份验证-开发平台登录
	LoginByOpenApi(context.Context, *resources.LoginByOpenApiReq) (*resources.LoginResp, error)
	mustEmbedUnimplementedSrvUserAuthV1Server()
}

// UnimplementedSrvUserAuthV1Server must be embedded to have forward compatible implementations.
type UnimplementedSrvUserAuthV1Server struct {
}

func (UnimplementedSrvUserAuthV1Server) Ping(context.Context, *resources.PingReq) (*resources.PingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedSrvUserAuthV1Server) SendPhoneSignupCode(context.Context, *resources.SendPhoneSignupCodeReq) (*resources.SendSignupCodeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPhoneSignupCode not implemented")
}
func (UnimplementedSrvUserAuthV1Server) SendEmailSignupCode(context.Context, *resources.SendEmailSignupCodeReq) (*resources.SendSignupCodeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmailSignupCode not implemented")
}
func (UnimplementedSrvUserAuthV1Server) SignupByEmail(context.Context, *resources.SignupByEmailReq) (*resources.LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignupByEmail not implemented")
}
func (UnimplementedSrvUserAuthV1Server) SignupByPhone(context.Context, *resources.SignupByPhoneReq) (*resources.LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignupByPhone not implemented")
}
func (UnimplementedSrvUserAuthV1Server) LoginOrSignupByPhone(context.Context, *resources.LoginOrSignupByPhoneReq) (*resources.LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginOrSignupByPhone not implemented")
}
func (UnimplementedSrvUserAuthV1Server) LoginOrSignupByEmail(context.Context, *resources.LoginOrSignupByEmailReq) (*resources.LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginOrSignupByEmail not implemented")
}
func (UnimplementedSrvUserAuthV1Server) RefreshToken(context.Context, *resources.RefreshTokenReq) (*resources.LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedSrvUserAuthV1Server) LoginByEmail(context.Context, *resources.LoginByEmailReq) (*resources.LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginByEmail not implemented")
}
func (UnimplementedSrvUserAuthV1Server) LoginByPhone(context.Context, *resources.LoginByPhoneReq) (*resources.LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginByPhone not implemented")
}
func (UnimplementedSrvUserAuthV1Server) LoginByOpenApi(context.Context, *resources.LoginByOpenApiReq) (*resources.LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginByOpenApi not implemented")
}
func (UnimplementedSrvUserAuthV1Server) mustEmbedUnimplementedSrvUserAuthV1Server() {}

// UnsafeSrvUserAuthV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SrvUserAuthV1Server will
// result in compilation errors.
type UnsafeSrvUserAuthV1Server interface {
	mustEmbedUnimplementedSrvUserAuthV1Server()
}

func RegisterSrvUserAuthV1Server(s grpc.ServiceRegistrar, srv SrvUserAuthV1Server) {
	s.RegisterService(&SrvUserAuthV1_ServiceDesc, srv)
}

func _SrvUserAuthV1_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvUserAuthV1Server).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvUserAuthV1_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvUserAuthV1Server).Ping(ctx, req.(*resources.PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvUserAuthV1_SendPhoneSignupCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.SendPhoneSignupCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvUserAuthV1Server).SendPhoneSignupCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvUserAuthV1_SendPhoneSignupCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvUserAuthV1Server).SendPhoneSignupCode(ctx, req.(*resources.SendPhoneSignupCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvUserAuthV1_SendEmailSignupCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.SendEmailSignupCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvUserAuthV1Server).SendEmailSignupCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvUserAuthV1_SendEmailSignupCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvUserAuthV1Server).SendEmailSignupCode(ctx, req.(*resources.SendEmailSignupCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvUserAuthV1_SignupByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.SignupByEmailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvUserAuthV1Server).SignupByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvUserAuthV1_SignupByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvUserAuthV1Server).SignupByEmail(ctx, req.(*resources.SignupByEmailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvUserAuthV1_SignupByPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.SignupByPhoneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvUserAuthV1Server).SignupByPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvUserAuthV1_SignupByPhone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvUserAuthV1Server).SignupByPhone(ctx, req.(*resources.SignupByPhoneReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvUserAuthV1_LoginOrSignupByPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.LoginOrSignupByPhoneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvUserAuthV1Server).LoginOrSignupByPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvUserAuthV1_LoginOrSignupByPhone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvUserAuthV1Server).LoginOrSignupByPhone(ctx, req.(*resources.LoginOrSignupByPhoneReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvUserAuthV1_LoginOrSignupByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.LoginOrSignupByEmailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvUserAuthV1Server).LoginOrSignupByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvUserAuthV1_LoginOrSignupByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvUserAuthV1Server).LoginOrSignupByEmail(ctx, req.(*resources.LoginOrSignupByEmailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvUserAuthV1_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.RefreshTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvUserAuthV1Server).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvUserAuthV1_RefreshToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvUserAuthV1Server).RefreshToken(ctx, req.(*resources.RefreshTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvUserAuthV1_LoginByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.LoginByEmailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvUserAuthV1Server).LoginByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvUserAuthV1_LoginByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvUserAuthV1Server).LoginByEmail(ctx, req.(*resources.LoginByEmailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvUserAuthV1_LoginByPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.LoginByPhoneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvUserAuthV1Server).LoginByPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvUserAuthV1_LoginByPhone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvUserAuthV1Server).LoginByPhone(ctx, req.(*resources.LoginByPhoneReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvUserAuthV1_LoginByOpenApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.LoginByOpenApiReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvUserAuthV1Server).LoginByOpenApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvUserAuthV1_LoginByOpenApi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvUserAuthV1Server).LoginByOpenApi(ctx, req.(*resources.LoginByOpenApiReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SrvUserAuthV1_ServiceDesc is the grpc.ServiceDesc for SrvUserAuthV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SrvUserAuthV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "saas.api.account.servicev1.SrvUserAuthV1",
	HandlerType: (*SrvUserAuthV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _SrvUserAuthV1_Ping_Handler,
		},
		{
			MethodName: "SendPhoneSignupCode",
			Handler:    _SrvUserAuthV1_SendPhoneSignupCode_Handler,
		},
		{
			MethodName: "SendEmailSignupCode",
			Handler:    _SrvUserAuthV1_SendEmailSignupCode_Handler,
		},
		{
			MethodName: "SignupByEmail",
			Handler:    _SrvUserAuthV1_SignupByEmail_Handler,
		},
		{
			MethodName: "SignupByPhone",
			Handler:    _SrvUserAuthV1_SignupByPhone_Handler,
		},
		{
			MethodName: "LoginOrSignupByPhone",
			Handler:    _SrvUserAuthV1_LoginOrSignupByPhone_Handler,
		},
		{
			MethodName: "LoginOrSignupByEmail",
			Handler:    _SrvUserAuthV1_LoginOrSignupByEmail_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _SrvUserAuthV1_RefreshToken_Handler,
		},
		{
			MethodName: "LoginByEmail",
			Handler:    _SrvUserAuthV1_LoginByEmail_Handler,
		},
		{
			MethodName: "LoginByPhone",
			Handler:    _SrvUserAuthV1_LoginByPhone_Handler,
		},
		{
			MethodName: "LoginByOpenApi",
			Handler:    _SrvUserAuthV1_LoginByOpenApi_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/account-service/v1/services/user_auth.service.v1.proto",
}
