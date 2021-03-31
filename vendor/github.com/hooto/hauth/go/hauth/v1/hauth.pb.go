// Copyright 2020 Eryx <evorui аt gmail dοt com>, All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.3
// source: hauth.proto

package hauth

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AccessKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty"`
	Secret      string         `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty" toml:"secret,omitempty"`
	User        string         `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty" toml:"user,omitempty"`
	Status      uint64         `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty" toml:"status,omitempty"`
	Roles       []string       `protobuf:"bytes,6,rep,name=roles,proto3" json:"roles,omitempty" toml:"roles,omitempty"`
	Scopes      []*ScopeFilter `protobuf:"bytes,11,rep,name=scopes,proto3" json:"scopes,omitempty" toml:"scopes,omitempty"`
	Description string         `protobuf:"bytes,13,opt,name=description,proto3" json:"description,omitempty" toml:"description,omitempty"`
}

func (x *AccessKey) Reset() {
	*x = AccessKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hauth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessKey) ProtoMessage() {}

func (x *AccessKey) ProtoReflect() protoreflect.Message {
	mi := &file_hauth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessKey.ProtoReflect.Descriptor instead.
func (*AccessKey) Descriptor() ([]byte, []int) {
	return file_hauth_proto_rawDescGZIP(), []int{0}
}

func (x *AccessKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AccessKey) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *AccessKey) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *AccessKey) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AccessKey) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *AccessKey) GetScopes() []*ScopeFilter {
	if x != nil {
		return x.Scopes
	}
	return nil
}

func (x *AccessKey) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type UserPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty"` // struct:object_slice_key
	Name    string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty"`
	Roles   []uint32 `protobuf:"varint,4,rep,packed,name=roles,proto3" json:"roles,omitempty" toml:"roles,omitempty"`
	Groups  []string `protobuf:"bytes,5,rep,name=groups,proto3" json:"groups,omitempty" toml:"groups,omitempty"`
	Expired int64    `protobuf:"varint,9,opt,name=expired,proto3" json:"expired,omitempty" toml:"expired,omitempty"` // unix time in seconds
}

func (x *UserPayload) Reset() {
	*x = UserPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hauth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPayload) ProtoMessage() {}

func (x *UserPayload) ProtoReflect() protoreflect.Message {
	mi := &file_hauth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPayload.ProtoReflect.Descriptor instead.
func (*UserPayload) Descriptor() ([]byte, []int) {
	return file_hauth_proto_rawDescGZIP(), []int{1}
}

func (x *UserPayload) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserPayload) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserPayload) GetRoles() []uint32 {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *UserPayload) GetGroups() []string {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *UserPayload) GetExpired() int64 {
	if x != nil {
		return x.Expired
	}
	return 0
}

type AppPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty"`
	User      string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty" toml:"user,omitempty"`
	AccessKey string `protobuf:"bytes,3,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty" toml:"access_key,omitempty"`
	Created   int64  `protobuf:"varint,9,opt,name=created,proto3" json:"created,omitempty" toml:"created,omitempty"` // unix time in milliseconds
}

func (x *AppPayload) Reset() {
	*x = AppPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hauth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppPayload) ProtoMessage() {}

func (x *AppPayload) ProtoReflect() protoreflect.Message {
	mi := &file_hauth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppPayload.ProtoReflect.Descriptor instead.
func (*AppPayload) Descriptor() ([]byte, []int) {
	return file_hauth_proto_rawDescGZIP(), []int{2}
}

func (x *AppPayload) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AppPayload) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *AppPayload) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *AppPayload) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

// A role in the RBAC.
type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the role.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty"`
	// Optional. A human-readable title for the role. Typically this
	// is limited to 100 UTF-8 bytes.
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty" toml:"title,omitempty"`
	// Optional. A human-readable description for the role.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty" toml:"description,omitempty"`
	// The names of the permissions this role grants when bound in an IAM policy.
	Permissions []string `protobuf:"bytes,7,rep,name=permissions,proto3" json:"permissions,omitempty" toml:"permissions,omitempty"`
	// The current launch status of the role.
	Status uint64 `protobuf:"varint,10,opt,name=status,proto3" json:"status,omitempty" toml:"status,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hauth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_hauth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_hauth_proto_rawDescGZIP(), []int{3}
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Role) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Role) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Role) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *Role) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

// A permission which can be included by a role.
type Permission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of this Permission.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty"`
	// The title of this Permission.
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty" toml:"title,omitempty"`
	// A brief description of what this Permission is used for.
	// This permission can ONLY be used in predefined roles.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty" toml:"description,omitempty"`
}

func (x *Permission) Reset() {
	*x = Permission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hauth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_hauth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_hauth_proto_rawDescGZIP(), []int{4}
}

func (x *Permission) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Permission) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Permission) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type ScopeFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty" toml:"value,omitempty"`
}

func (x *ScopeFilter) Reset() {
	*x = ScopeFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hauth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScopeFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScopeFilter) ProtoMessage() {}

func (x *ScopeFilter) ProtoReflect() protoreflect.Message {
	mi := &file_hauth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScopeFilter.ProtoReflect.Descriptor instead.
func (*ScopeFilter) Descriptor() ([]byte, []int) {
	return file_hauth_proto_rawDescGZIP(), []int{5}
}

func (x *ScopeFilter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ScopeFilter) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_hauth_proto protoreflect.FileDescriptor

var file_hauth_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x68, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x68,
	0x6f, 0x6f, 0x74, 0x6f, 0x2e, 0x68, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x22, 0xcc, 0x01,
	0x0a, 0x09, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18,
	0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x68, 0x6f, 0x6f, 0x74, 0x6f, 0x2e, 0x68, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x79, 0x0a, 0x0b,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x05,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x22, 0x69, 0x0a, 0x0a, 0x41, 0x70, 0x70, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x22, 0x8c, 0x01, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x58, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x37, 0x0a, 0x0b, 0x53,
	0x63, 0x6f, 0x70, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x0b, 0x48, 0x03, 0x5a, 0x07, 0x2e, 0x3b, 0x68, 0x61, 0x75, 0x74,
	0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hauth_proto_rawDescOnce sync.Once
	file_hauth_proto_rawDescData = file_hauth_proto_rawDesc
)

func file_hauth_proto_rawDescGZIP() []byte {
	file_hauth_proto_rawDescOnce.Do(func() {
		file_hauth_proto_rawDescData = protoimpl.X.CompressGZIP(file_hauth_proto_rawDescData)
	})
	return file_hauth_proto_rawDescData
}

var file_hauth_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_hauth_proto_goTypes = []interface{}{
	(*AccessKey)(nil),   // 0: hooto.hauth.v1.AccessKey
	(*UserPayload)(nil), // 1: hooto.hauth.v1.UserPayload
	(*AppPayload)(nil),  // 2: hooto.hauth.v1.AppPayload
	(*Role)(nil),        // 3: hooto.hauth.v1.Role
	(*Permission)(nil),  // 4: hooto.hauth.v1.Permission
	(*ScopeFilter)(nil), // 5: hooto.hauth.v1.ScopeFilter
}
var file_hauth_proto_depIdxs = []int32{
	5, // 0: hooto.hauth.v1.AccessKey.scopes:type_name -> hooto.hauth.v1.ScopeFilter
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_hauth_proto_init() }
func file_hauth_proto_init() {
	if File_hauth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hauth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessKey); i {
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
		file_hauth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPayload); i {
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
		file_hauth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppPayload); i {
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
		file_hauth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_hauth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Permission); i {
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
		file_hauth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScopeFilter); i {
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
			RawDescriptor: file_hauth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hauth_proto_goTypes,
		DependencyIndexes: file_hauth_proto_depIdxs,
		MessageInfos:      file_hauth_proto_msgTypes,
	}.Build()
	File_hauth_proto = out.File
	file_hauth_proto_rawDesc = nil
	file_hauth_proto_goTypes = nil
	file_hauth_proto_depIdxs = nil
}
