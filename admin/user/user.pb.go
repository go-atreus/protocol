// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: admin/user/user.proto

package user

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Secret     string `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	PlatformID int32  `protobuf:"varint,2,opt,name=platformID,proto3" json:"platformID,omitempty"`
	UserID     string `protobuf:"bytes,3,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *UserInfoReq) Reset() {
	*x = UserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoReq) ProtoMessage() {}

func (x *UserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_admin_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoReq.ProtoReflect.Descriptor instead.
func (*UserInfoReq) Descriptor() ([]byte, []int) {
	return file_admin_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfoReq) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *UserInfoReq) GetPlatformID() int32 {
	if x != nil {
		return x.PlatformID
	}
	return 0
}

func (x *UserInfoReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type UserInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	UserID int32  `protobuf:"varint,4,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *UserInfoResp) Reset() {
	*x = UserInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoResp) ProtoMessage() {}

func (x *UserInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_admin_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoResp.ProtoReflect.Descriptor instead.
func (*UserInfoResp) Descriptor() ([]byte, []int) {
	return file_admin_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserInfoResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UserInfoResp) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type ForceLogoutReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlatformID int32  `protobuf:"varint,1,opt,name=platformID,proto3" json:"platformID,omitempty"`
	UserID     string `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *ForceLogoutReq) Reset() {
	*x = ForceLogoutReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_user_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForceLogoutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForceLogoutReq) ProtoMessage() {}

func (x *ForceLogoutReq) ProtoReflect() protoreflect.Message {
	mi := &file_admin_user_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForceLogoutReq.ProtoReflect.Descriptor instead.
func (*ForceLogoutReq) Descriptor() ([]byte, []int) {
	return file_admin_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *ForceLogoutReq) GetPlatformID() int32 {
	if x != nil {
		return x.PlatformID
	}
	return 0
}

func (x *ForceLogoutReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type ForceLogoutResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ForceLogoutResp) Reset() {
	*x = ForceLogoutResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_user_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForceLogoutResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForceLogoutResp) ProtoMessage() {}

func (x *ForceLogoutResp) ProtoReflect() protoreflect.Message {
	mi := &file_admin_user_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForceLogoutResp.ProtoReflect.Descriptor instead.
func (*ForceLogoutResp) Descriptor() ([]byte, []int) {
	return file_admin_user_user_proto_rawDescGZIP(), []int{3}
}

type ParseTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ParseTokenReq) Reset() {
	*x = ParseTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_user_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParseTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseTokenReq) ProtoMessage() {}

func (x *ParseTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_admin_user_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseTokenReq.ProtoReflect.Descriptor instead.
func (*ParseTokenReq) Descriptor() ([]byte, []int) {
	return file_admin_user_user_proto_rawDescGZIP(), []int{4}
}

func (x *ParseTokenReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ParseTokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID            string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Platform          string `protobuf:"bytes,2,opt,name=platform,proto3" json:"platform,omitempty"`
	ExpireTimeSeconds int64  `protobuf:"varint,4,opt,name=expireTimeSeconds,proto3" json:"expireTimeSeconds,omitempty"`
}

func (x *ParseTokenResp) Reset() {
	*x = ParseTokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_user_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParseTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseTokenResp) ProtoMessage() {}

func (x *ParseTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_admin_user_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseTokenResp.ProtoReflect.Descriptor instead.
func (*ParseTokenResp) Descriptor() ([]byte, []int) {
	return file_admin_user_user_proto_rawDescGZIP(), []int{5}
}

func (x *ParseTokenResp) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *ParseTokenResp) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *ParseTokenResp) GetExpireTimeSeconds() int64 {
	if x != nil {
		return x.ExpireTimeSeconds
	}
	return 0
}

type GetUserTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlatformID int32  `protobuf:"varint,1,opt,name=platformID,proto3" json:"platformID,omitempty"`
	UserID     string `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *GetUserTokenReq) Reset() {
	*x = GetUserTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_user_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserTokenReq) ProtoMessage() {}

func (x *GetUserTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_admin_user_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserTokenReq.ProtoReflect.Descriptor instead.
func (*GetUserTokenReq) Descriptor() ([]byte, []int) {
	return file_admin_user_user_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserTokenReq) GetPlatformID() int32 {
	if x != nil {
		return x.PlatformID
	}
	return 0
}

func (x *GetUserTokenReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type GetUserTokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token             string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ExpireTimeSeconds int64  `protobuf:"varint,2,opt,name=expireTimeSeconds,proto3" json:"expireTimeSeconds,omitempty"`
}

func (x *GetUserTokenResp) Reset() {
	*x = GetUserTokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_user_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserTokenResp) ProtoMessage() {}

func (x *GetUserTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_admin_user_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserTokenResp.ProtoReflect.Descriptor instead.
func (*GetUserTokenResp) Descriptor() ([]byte, []int) {
	return file_admin_user_user_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserTokenResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetUserTokenResp) GetExpireTimeSeconds() int64 {
	if x != nil {
		return x.ExpireTimeSeconds
	}
	return 0
}

var File_admin_user_user_proto protoreflect.FileDescriptor

var file_admin_user_user_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x41, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x22, 0x3c, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22,
	0x48, 0x0a, 0x0e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49,
	0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x11, 0x0a, 0x0f, 0x66, 0x6f, 0x72,
	0x63, 0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x25, 0x0a, 0x0d,
	0x70, 0x61, 0x72, 0x73, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x72, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x2c, 0x0a, 0x11, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x49, 0x0a, 0x0f, 0x67, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x22, 0x56, 0x0a, 0x10, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2c, 0x0a, 0x11,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x32, 0x61, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x59, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x18, 0x2e, 0x41, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x41, 0x74,
	0x72, 0x65, 0x75, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01,
	0x2a, 0x22, 0x0a, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x42, 0x93, 0x01,
	0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x41, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x42, 0x09, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x28,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x61, 0x74,
	0x72, 0x65, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0xa2, 0x02, 0x03, 0x41, 0x55, 0x58, 0xaa, 0x02,
	0x0b, 0x41, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0xca, 0x02, 0x0b, 0x41,
	0x74, 0x72, 0x65, 0x75, 0x73, 0x5c, 0x55, 0x73, 0x65, 0x72, 0xe2, 0x02, 0x17, 0x41, 0x74, 0x72,
	0x65, 0x75, 0x73, 0x5c, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x41, 0x74, 0x72, 0x65, 0x75, 0x73, 0x3a, 0x3a, 0x55,
	0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_user_user_proto_rawDescOnce sync.Once
	file_admin_user_user_proto_rawDescData = file_admin_user_user_proto_rawDesc
)

func file_admin_user_user_proto_rawDescGZIP() []byte {
	file_admin_user_user_proto_rawDescOnce.Do(func() {
		file_admin_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_user_user_proto_rawDescData)
	})
	return file_admin_user_user_proto_rawDescData
}

var file_admin_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_admin_user_user_proto_goTypes = []interface{}{
	(*UserInfoReq)(nil),      // 0: Atreus.user.userInfoReq
	(*UserInfoResp)(nil),     // 1: Atreus.user.userInfoResp
	(*ForceLogoutReq)(nil),   // 2: Atreus.user.forceLogoutReq
	(*ForceLogoutResp)(nil),  // 3: Atreus.user.forceLogoutResp
	(*ParseTokenReq)(nil),    // 4: Atreus.user.parseTokenReq
	(*ParseTokenResp)(nil),   // 5: Atreus.user.parseTokenResp
	(*GetUserTokenReq)(nil),  // 6: Atreus.user.getUserTokenReq
	(*GetUserTokenResp)(nil), // 7: Atreus.user.getUserTokenResp
}
var file_admin_user_user_proto_depIdxs = []int32{
	0, // 0: Atreus.user.User.getUserInfo:input_type -> Atreus.user.userInfoReq
	1, // 1: Atreus.user.User.getUserInfo:output_type -> Atreus.user.userInfoResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_admin_user_user_proto_init() }
func file_admin_user_user_proto_init() {
	if File_admin_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admin_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoReq); i {
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
		file_admin_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoResp); i {
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
		file_admin_user_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForceLogoutReq); i {
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
		file_admin_user_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForceLogoutResp); i {
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
		file_admin_user_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParseTokenReq); i {
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
		file_admin_user_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParseTokenResp); i {
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
		file_admin_user_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserTokenReq); i {
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
		file_admin_user_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserTokenResp); i {
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
			RawDescriptor: file_admin_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_user_user_proto_goTypes,
		DependencyIndexes: file_admin_user_user_proto_depIdxs,
		MessageInfos:      file_admin_user_user_proto_msgTypes,
	}.Build()
	File_admin_user_user_proto = out.File
	file_admin_user_user_proto_rawDesc = nil
	file_admin_user_user_proto_goTypes = nil
	file_admin_user_user_proto_depIdxs = nil
}
