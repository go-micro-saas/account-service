// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v3.21.6
// source: api/account-service/v1/services/user_auth.service.v1.proto

package servicev1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	resources "github.com/go-micro-saas/account-service/api/account-service/v1/resources"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationSrvUserAuthV1LoginByEmail = "/saas.api.account.servicev1.SrvUserAuthV1/LoginByEmail"
const OperationSrvUserAuthV1LoginByPhone = "/saas.api.account.servicev1.SrvUserAuthV1/LoginByPhone"
const OperationSrvUserAuthV1OpenApiLogin = "/saas.api.account.servicev1.SrvUserAuthV1/OpenApiLogin"
const OperationSrvUserAuthV1Ping = "/saas.api.account.servicev1.SrvUserAuthV1/Ping"
const OperationSrvUserAuthV1RefreshToken = "/saas.api.account.servicev1.SrvUserAuthV1/RefreshToken"

type SrvUserAuthV1HTTPServer interface {
	// LoginByEmail 身份验证-Email登录
	LoginByEmail(context.Context, *resources.LoginByEmailReq) (*resources.LoginResp, error)
	// LoginByPhone 身份验证-手机登录
	LoginByPhone(context.Context, *resources.LoginByPhoneReq) (*resources.LoginResp, error)
	// OpenApiLogin 身份验证-开发平台登录
	OpenApiLogin(context.Context, *resources.OpenApiLoginReq) (*resources.LoginResp, error)
	// Ping 身份验证-Ping测试
	Ping(context.Context, *resources.PingReq) (*resources.PingResp, error)
	// RefreshToken 身份验证-刷新Token
	RefreshToken(context.Context, *resources.RefreshTokenReq) (*resources.LoginResp, error)
}

func RegisterSrvUserAuthV1HTTPServer(s *http.Server, srv SrvUserAuthV1HTTPServer) {
	r := s.Route("/")
	r.GET("/api/v1/user/auth/ping", _SrvUserAuthV1_Ping0_HTTP_Handler(srv))
	r.POST("/api/v1/user/auth/refresh-token", _SrvUserAuthV1_RefreshToken0_HTTP_Handler(srv))
	r.POST("/api/v1/user/auth/login-by-email", _SrvUserAuthV1_LoginByEmail0_HTTP_Handler(srv))
	r.POST("/api/v1/user/auth/login-by-phone", _SrvUserAuthV1_LoginByPhone0_HTTP_Handler(srv))
	r.POST("/api/v1/user/open-api/login", _SrvUserAuthV1_OpenApiLogin0_HTTP_Handler(srv))
}

func _SrvUserAuthV1_Ping0_HTTP_Handler(srv SrvUserAuthV1HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in resources.PingReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSrvUserAuthV1Ping)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Ping(ctx, req.(*resources.PingReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*resources.PingResp)
		return ctx.Result(200, reply)
	}
}

func _SrvUserAuthV1_RefreshToken0_HTTP_Handler(srv SrvUserAuthV1HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in resources.RefreshTokenReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSrvUserAuthV1RefreshToken)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RefreshToken(ctx, req.(*resources.RefreshTokenReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*resources.LoginResp)
		return ctx.Result(200, reply)
	}
}

func _SrvUserAuthV1_LoginByEmail0_HTTP_Handler(srv SrvUserAuthV1HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in resources.LoginByEmailReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSrvUserAuthV1LoginByEmail)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LoginByEmail(ctx, req.(*resources.LoginByEmailReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*resources.LoginResp)
		return ctx.Result(200, reply)
	}
}

func _SrvUserAuthV1_LoginByPhone0_HTTP_Handler(srv SrvUserAuthV1HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in resources.LoginByPhoneReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSrvUserAuthV1LoginByPhone)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LoginByPhone(ctx, req.(*resources.LoginByPhoneReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*resources.LoginResp)
		return ctx.Result(200, reply)
	}
}

func _SrvUserAuthV1_OpenApiLogin0_HTTP_Handler(srv SrvUserAuthV1HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in resources.OpenApiLoginReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSrvUserAuthV1OpenApiLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.OpenApiLogin(ctx, req.(*resources.OpenApiLoginReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*resources.LoginResp)
		return ctx.Result(200, reply)
	}
}

type SrvUserAuthV1HTTPClient interface {
	LoginByEmail(ctx context.Context, req *resources.LoginByEmailReq, opts ...http.CallOption) (rsp *resources.LoginResp, err error)
	LoginByPhone(ctx context.Context, req *resources.LoginByPhoneReq, opts ...http.CallOption) (rsp *resources.LoginResp, err error)
	OpenApiLogin(ctx context.Context, req *resources.OpenApiLoginReq, opts ...http.CallOption) (rsp *resources.LoginResp, err error)
	Ping(ctx context.Context, req *resources.PingReq, opts ...http.CallOption) (rsp *resources.PingResp, err error)
	RefreshToken(ctx context.Context, req *resources.RefreshTokenReq, opts ...http.CallOption) (rsp *resources.LoginResp, err error)
}

type SrvUserAuthV1HTTPClientImpl struct {
	cc *http.Client
}

func NewSrvUserAuthV1HTTPClient(client *http.Client) SrvUserAuthV1HTTPClient {
	return &SrvUserAuthV1HTTPClientImpl{client}
}

func (c *SrvUserAuthV1HTTPClientImpl) LoginByEmail(ctx context.Context, in *resources.LoginByEmailReq, opts ...http.CallOption) (*resources.LoginResp, error) {
	var out resources.LoginResp
	pattern := "/api/v1/user/auth/login-by-email"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSrvUserAuthV1LoginByEmail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SrvUserAuthV1HTTPClientImpl) LoginByPhone(ctx context.Context, in *resources.LoginByPhoneReq, opts ...http.CallOption) (*resources.LoginResp, error) {
	var out resources.LoginResp
	pattern := "/api/v1/user/auth/login-by-phone"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSrvUserAuthV1LoginByPhone))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SrvUserAuthV1HTTPClientImpl) OpenApiLogin(ctx context.Context, in *resources.OpenApiLoginReq, opts ...http.CallOption) (*resources.LoginResp, error) {
	var out resources.LoginResp
	pattern := "/api/v1/user/open-api/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSrvUserAuthV1OpenApiLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SrvUserAuthV1HTTPClientImpl) Ping(ctx context.Context, in *resources.PingReq, opts ...http.CallOption) (*resources.PingResp, error) {
	var out resources.PingResp
	pattern := "/api/v1/user/auth/ping"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSrvUserAuthV1Ping))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SrvUserAuthV1HTTPClientImpl) RefreshToken(ctx context.Context, in *resources.RefreshTokenReq, opts ...http.CallOption) (*resources.LoginResp, error) {
	var out resources.LoginResp
	pattern := "/api/v1/user/auth/refresh-token"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSrvUserAuthV1RefreshToken))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
