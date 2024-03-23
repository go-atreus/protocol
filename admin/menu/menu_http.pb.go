// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// source: admin/menu/menu.proto

package menu

import (
	context "context"
	gin "github.com/gin-gonic/gin"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = gin.Version

const OperationMenugetMenu = "/Atreus.menu.Menu/getMenu"
const OperationMenugetMenuList = "/Atreus.menu.Menu/getMenuList"

type MenuHTTPServer interface {
	// GetMenu获取菜单树
	GetMenu(context.Context, *GetMenuReq) (*GetMenuResp, error)
	// GetMenuList分页获取基础menu列表
	GetMenuList(context.Context, *GetMenuReq) (*GetMenuResp, error)
}

func RegisterMenuHTTPServer(s *gin.Engine, handleFunc gin.HandlerFunc, srv MenuHTTPServer) {
	r := s.Group("/")
	r.POST("/menu/getMenu", _Menu_GetMenu0_HTTP_Handler(srv))
	r.POST("/menu/getMenuList", _Menu_GetMenuList0_HTTP_Handler(srv))
}

func _Menu_GetMenu0_HTTP_Handler(srv MenuHTTPServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var in GetMenuReq
		if err := ctx.Bind(&in); err != nil {
			return
		}
		if err := ctx.BindQuery(&in); err != nil {
			return
		}
		out, err := srv.GetMenu(ctx, &in)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		reply := out
		ctx.JSON(200, reply)
		return
	}
}

func _Menu_GetMenuList0_HTTP_Handler(srv MenuHTTPServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var in GetMenuReq
		if err := ctx.Bind(&in); err != nil {
			return
		}
		if err := ctx.BindQuery(&in); err != nil {
			return
		}
		out, err := srv.GetMenuList(ctx, &in)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		reply := out
		ctx.JSON(200, reply)
		return
	}
}
