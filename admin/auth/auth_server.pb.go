// Code generated by protoc-gen-go-server. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v0.0.1
// source: admin/auth/auth.proto

package auth

import (
	context "context"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = grpc.Version

type AuthImpl struct {
	cc AuthClient
}

func NewAuthImpl(conn *grpc.ClientConn) *AuthImpl {
	return &AuthImpl{cc: NewAuthClient(conn)}
}

func (c *AuthImpl) UserToken(ctx context.Context, in *UserTokenReq) (*UserTokenResp, error) {
	return c.cc.UserToken(ctx, in)
}

func (c *AuthImpl) GetUserToken(ctx context.Context, in *GetUserTokenReq) (*GetUserTokenResp, error) {
	return c.cc.GetUserToken(ctx, in)
}

func (c *AuthImpl) ForceLogout(ctx context.Context, in *ForceLogoutReq) (*ForceLogoutResp, error) {
	return c.cc.ForceLogout(ctx, in)
}

func (c *AuthImpl) ParseToken(ctx context.Context, in *ParseTokenReq) (*ParseTokenResp, error) {
	return c.cc.ParseToken(ctx, in)
}
