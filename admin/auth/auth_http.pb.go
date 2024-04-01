// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             (unknown)
// source: admin/auth/auth.proto

package auth

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationAuthforceLogout = "/Atreus.auth.Auth/forceLogout"
const OperationAuthgetUserToken = "/Atreus.auth.Auth/getUserToken"
const OperationAuthparseToken = "/Atreus.auth.Auth/parseToken"
const OperationAuthuserLogin = "/Atreus.auth.Auth/userLogin"
const OperationAuthuserToken = "/Atreus.auth.Auth/userToken"

type AuthHTTPServer interface {
	// ForceLogout强制退出登录
	ForceLogout(context.Context, *ForceLogoutReq) (*ForceLogoutResp, error)
	// GetUserToken 管理员获取用户 token
	GetUserToken(context.Context, *GetUserTokenReq) (*GetUserTokenResp, error)
	// ParseToken解析token
	ParseToken(context.Context, *ParseTokenReq) (*ParseTokenResp, error)
	// UserLoginlogin
	UserLogin(context.Context, *UserLoginReq) (*UserTokenResp, error)
	// UserToken生成token
	UserToken(context.Context, *UserTokenReq) (*UserTokenResp, error)
}

func RegisterAuthHTTPServer(s *http.Server, srv AuthHTTPServer) {
	r := s.Route("/")
	r.POST("/auth/login", _Auth_UserLogin0_HTTP_Handler(srv))
	r.POST("/auth/user_token", _Auth_UserToken0_HTTP_Handler(srv))
	r.POST("/auth/get_user_token", _Auth_GetUserToken0_HTTP_Handler(srv))
	r.POST("/auth/force_logout", _Auth_ForceLogout0_HTTP_Handler(srv))
	r.POST("/auth/parse_token", _Auth_ParseToken0_HTTP_Handler(srv))
}

func _Auth_UserLogin0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserLoginReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthuserLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserLogin(ctx, req.(*UserLoginReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserTokenResp)
		return ctx.Result(200, reply)
	}
}

func _Auth_UserToken0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserTokenReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthuserToken)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserToken(ctx, req.(*UserTokenReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserTokenResp)
		return ctx.Result(200, reply)
	}
}

func _Auth_GetUserToken0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserTokenReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthgetUserToken)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserToken(ctx, req.(*GetUserTokenReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserTokenResp)
		return ctx.Result(200, reply)
	}
}

func _Auth_ForceLogout0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ForceLogoutReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthforceLogout)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ForceLogout(ctx, req.(*ForceLogoutReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ForceLogoutResp)
		return ctx.Result(200, reply)
	}
}

func _Auth_ParseToken0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ParseTokenReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthparseToken)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ParseToken(ctx, req.(*ParseTokenReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ParseTokenResp)
		return ctx.Result(200, reply)
	}
}

type AuthHTTPClient interface {
	ForceLogout(ctx context.Context, req *ForceLogoutReq, opts ...http.CallOption) (rsp *ForceLogoutResp, err error)
	GetUserToken(ctx context.Context, req *GetUserTokenReq, opts ...http.CallOption) (rsp *GetUserTokenResp, err error)
	ParseToken(ctx context.Context, req *ParseTokenReq, opts ...http.CallOption) (rsp *ParseTokenResp, err error)
	UserLogin(ctx context.Context, req *UserLoginReq, opts ...http.CallOption) (rsp *UserTokenResp, err error)
	UserToken(ctx context.Context, req *UserTokenReq, opts ...http.CallOption) (rsp *UserTokenResp, err error)
}

type AuthHTTPClientImpl struct {
	cc *http.Client
}

func NewAuthHTTPClient(client *http.Client) AuthHTTPClient {
	return &AuthHTTPClientImpl{client}
}

func (c *AuthHTTPClientImpl) ForceLogout(ctx context.Context, in *ForceLogoutReq, opts ...http.CallOption) (*ForceLogoutResp, error) {
	var out ForceLogoutResp
	pattern := "/auth/force_logout"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthforceLogout))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AuthHTTPClientImpl) GetUserToken(ctx context.Context, in *GetUserTokenReq, opts ...http.CallOption) (*GetUserTokenResp, error) {
	var out GetUserTokenResp
	pattern := "/auth/get_user_token"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthgetUserToken))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AuthHTTPClientImpl) ParseToken(ctx context.Context, in *ParseTokenReq, opts ...http.CallOption) (*ParseTokenResp, error) {
	var out ParseTokenResp
	pattern := "/auth/parse_token"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthparseToken))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AuthHTTPClientImpl) UserLogin(ctx context.Context, in *UserLoginReq, opts ...http.CallOption) (*UserTokenResp, error) {
	var out UserTokenResp
	pattern := "/auth/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthuserLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AuthHTTPClientImpl) UserToken(ctx context.Context, in *UserTokenReq, opts ...http.CallOption) (*UserTokenResp, error) {
	var out UserTokenResp
	pattern := "/auth/user_token"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthuserToken))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
