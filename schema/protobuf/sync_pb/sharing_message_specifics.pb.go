// Copyright 2020 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// Sync protocol datatype extension for sharing message.

// If you change or add any fields in this file, update proto_visitors.h and
// potentially proto_enum_conversions.{h, cc}.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.4
// source: sharing_message_specifics.proto

package sync_pb

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

// This enum is used in histograms. Entries should not be renumbered and
// numeric values should never be reused. Also remember to update in
// tools/metrics/histograms/enums.xml SyncSharingMessageCommitErrorCode enum.
type SharingMessageCommitError_ErrorCode int32

const (
	SharingMessageCommitError_NONE               SharingMessageCommitError_ErrorCode = 0
	SharingMessageCommitError_INVALID_ARGUMENT   SharingMessageCommitError_ErrorCode = 1
	SharingMessageCommitError_NOT_FOUND          SharingMessageCommitError_ErrorCode = 2
	SharingMessageCommitError_INTERNAL           SharingMessageCommitError_ErrorCode = 3
	SharingMessageCommitError_UNAVAILABLE        SharingMessageCommitError_ErrorCode = 4
	SharingMessageCommitError_RESOURCE_EXHAUSTED SharingMessageCommitError_ErrorCode = 5
	SharingMessageCommitError_UNAUTHENTICATED    SharingMessageCommitError_ErrorCode = 6
	SharingMessageCommitError_PERMISSION_DENIED  SharingMessageCommitError_ErrorCode = 7
	// Client-specific error codes.
	SharingMessageCommitError_SYNC_TURNED_OFF    SharingMessageCommitError_ErrorCode = 8
	SharingMessageCommitError_SYNC_NETWORK_ERROR SharingMessageCommitError_ErrorCode = 9
	// Deprecated UMA bucket, prior to splitting between SYNC_SERVER_ERROR and
	// SYNC_AUTH_ERROR.
	SharingMessageCommitError_DEPRECATED_SYNC_SERVER_OR_AUTH_ERROR SharingMessageCommitError_ErrorCode = 10
	// Message wasn't committed before timeout.
	SharingMessageCommitError_SYNC_TIMEOUT SharingMessageCommitError_ErrorCode = 11
	// Error code for server error or unparsable server response.
	SharingMessageCommitError_SYNC_SERVER_ERROR SharingMessageCommitError_ErrorCode = 12
	// Auth error when communicating with the server.
	SharingMessageCommitError_SYNC_AUTH_ERROR SharingMessageCommitError_ErrorCode = 13
)

// Enum value maps for SharingMessageCommitError_ErrorCode.
var (
	SharingMessageCommitError_ErrorCode_name = map[int32]string{
		0:  "NONE",
		1:  "INVALID_ARGUMENT",
		2:  "NOT_FOUND",
		3:  "INTERNAL",
		4:  "UNAVAILABLE",
		5:  "RESOURCE_EXHAUSTED",
		6:  "UNAUTHENTICATED",
		7:  "PERMISSION_DENIED",
		8:  "SYNC_TURNED_OFF",
		9:  "SYNC_NETWORK_ERROR",
		10: "DEPRECATED_SYNC_SERVER_OR_AUTH_ERROR",
		11: "SYNC_TIMEOUT",
		12: "SYNC_SERVER_ERROR",
		13: "SYNC_AUTH_ERROR",
	}
	SharingMessageCommitError_ErrorCode_value = map[string]int32{
		"NONE":                                 0,
		"INVALID_ARGUMENT":                     1,
		"NOT_FOUND":                            2,
		"INTERNAL":                             3,
		"UNAVAILABLE":                          4,
		"RESOURCE_EXHAUSTED":                   5,
		"UNAUTHENTICATED":                      6,
		"PERMISSION_DENIED":                    7,
		"SYNC_TURNED_OFF":                      8,
		"SYNC_NETWORK_ERROR":                   9,
		"DEPRECATED_SYNC_SERVER_OR_AUTH_ERROR": 10,
		"SYNC_TIMEOUT":                         11,
		"SYNC_SERVER_ERROR":                    12,
		"SYNC_AUTH_ERROR":                      13,
	}
)

func (x SharingMessageCommitError_ErrorCode) Enum() *SharingMessageCommitError_ErrorCode {
	p := new(SharingMessageCommitError_ErrorCode)
	*p = x
	return p
}

func (x SharingMessageCommitError_ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SharingMessageCommitError_ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_sharing_message_specifics_proto_enumTypes[0].Descriptor()
}

func (SharingMessageCommitError_ErrorCode) Type() protoreflect.EnumType {
	return &file_sharing_message_specifics_proto_enumTypes[0]
}

func (x SharingMessageCommitError_ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *SharingMessageCommitError_ErrorCode) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = SharingMessageCommitError_ErrorCode(num)
	return nil
}

// Deprecated: Use SharingMessageCommitError_ErrorCode.Descriptor instead.
func (SharingMessageCommitError_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_sharing_message_specifics_proto_rawDescGZIP(), []int{1, 0}
}

type SharingMessageSpecifics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier of message.
	MessageId            *string                                       `protobuf:"bytes,1,opt,name=message_id,json=messageId" json:"message_id,omitempty"`
	ChannelConfiguration *SharingMessageSpecifics_ChannelConfiguration `protobuf:"bytes,2,opt,name=channel_configuration,json=channelConfiguration" json:"channel_configuration,omitempty"`
	// Payload encrypted using the target user keys according to WebPush
	// encryption scheme. The payload has to be a valid
	// chrome/browser/sharing/proto/sharing_message.proto serialized using
	// SerializeToString.
	Payload []byte `protobuf:"bytes,3,opt,name=payload" json:"payload,omitempty"`
}

func (x *SharingMessageSpecifics) Reset() {
	*x = SharingMessageSpecifics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sharing_message_specifics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SharingMessageSpecifics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SharingMessageSpecifics) ProtoMessage() {}

func (x *SharingMessageSpecifics) ProtoReflect() protoreflect.Message {
	mi := &file_sharing_message_specifics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SharingMessageSpecifics.ProtoReflect.Descriptor instead.
func (*SharingMessageSpecifics) Descriptor() ([]byte, []int) {
	return file_sharing_message_specifics_proto_rawDescGZIP(), []int{0}
}

func (x *SharingMessageSpecifics) GetMessageId() string {
	if x != nil && x.MessageId != nil {
		return *x.MessageId
	}
	return ""
}

func (x *SharingMessageSpecifics) GetChannelConfiguration() *SharingMessageSpecifics_ChannelConfiguration {
	if x != nil {
		return x.ChannelConfiguration
	}
	return nil
}

func (x *SharingMessageSpecifics) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

// Used for the server to return fine grained commit errors back to the client.
type SharingMessageCommitError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode *SharingMessageCommitError_ErrorCode `protobuf:"varint,1,opt,name=error_code,json=errorCode,enum=sync_pb.SharingMessageCommitError_ErrorCode" json:"error_code,omitempty"`
}

func (x *SharingMessageCommitError) Reset() {
	*x = SharingMessageCommitError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sharing_message_specifics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SharingMessageCommitError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SharingMessageCommitError) ProtoMessage() {}

func (x *SharingMessageCommitError) ProtoReflect() protoreflect.Message {
	mi := &file_sharing_message_specifics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SharingMessageCommitError.ProtoReflect.Descriptor instead.
func (*SharingMessageCommitError) Descriptor() ([]byte, []int) {
	return file_sharing_message_specifics_proto_rawDescGZIP(), []int{1}
}

func (x *SharingMessageCommitError) GetErrorCode() SharingMessageCommitError_ErrorCode {
	if x != nil && x.ErrorCode != nil {
		return *x.ErrorCode
	}
	return SharingMessageCommitError_NONE
}

type SharingMessageSpecifics_ChannelConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ChannelConfiguration:
	//	*SharingMessageSpecifics_ChannelConfiguration_Fcm
	//	*SharingMessageSpecifics_ChannelConfiguration_Server
	ChannelConfiguration isSharingMessageSpecifics_ChannelConfiguration_ChannelConfiguration `protobuf_oneof:"channel_configuration"`
}

func (x *SharingMessageSpecifics_ChannelConfiguration) Reset() {
	*x = SharingMessageSpecifics_ChannelConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sharing_message_specifics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SharingMessageSpecifics_ChannelConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SharingMessageSpecifics_ChannelConfiguration) ProtoMessage() {}

func (x *SharingMessageSpecifics_ChannelConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_sharing_message_specifics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SharingMessageSpecifics_ChannelConfiguration.ProtoReflect.Descriptor instead.
func (*SharingMessageSpecifics_ChannelConfiguration) Descriptor() ([]byte, []int) {
	return file_sharing_message_specifics_proto_rawDescGZIP(), []int{0, 0}
}

func (m *SharingMessageSpecifics_ChannelConfiguration) GetChannelConfiguration() isSharingMessageSpecifics_ChannelConfiguration_ChannelConfiguration {
	if m != nil {
		return m.ChannelConfiguration
	}
	return nil
}

func (x *SharingMessageSpecifics_ChannelConfiguration) GetFcm() *SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration {
	if x, ok := x.GetChannelConfiguration().(*SharingMessageSpecifics_ChannelConfiguration_Fcm); ok {
		return x.Fcm
	}
	return nil
}

func (x *SharingMessageSpecifics_ChannelConfiguration) GetServer() []byte {
	if x, ok := x.GetChannelConfiguration().(*SharingMessageSpecifics_ChannelConfiguration_Server); ok {
		return x.Server
	}
	return nil
}

type isSharingMessageSpecifics_ChannelConfiguration_ChannelConfiguration interface {
	isSharingMessageSpecifics_ChannelConfiguration_ChannelConfiguration()
}

type SharingMessageSpecifics_ChannelConfiguration_Fcm struct {
	// FCM channel configuration. Message will be delivered as a FCM message.
	Fcm *SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration `protobuf:"bytes,1,opt,name=fcm,oneof"`
}

type SharingMessageSpecifics_ChannelConfiguration_Server struct {
	// Opaque server channel configuration. Message will be delivered through
	// server channel.
	Server []byte `protobuf:"bytes,2,opt,name=server,oneof"`
}

func (*SharingMessageSpecifics_ChannelConfiguration_Fcm) isSharingMessageSpecifics_ChannelConfiguration_ChannelConfiguration() {
}

func (*SharingMessageSpecifics_ChannelConfiguration_Server) isSharingMessageSpecifics_ChannelConfiguration_ChannelConfiguration() {
}

type SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FCM registration token of target device.
	Token *string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	// Time to live for a FCM message (in seconds) - if specified, the message
	// will expire based on the TTL.
	Ttl *int32 `protobuf:"varint,2,opt,name=ttl" json:"ttl,omitempty"`
	// Priority level of a FCM message. 5 = normal, 10 = high.
	Priority *int32 `protobuf:"varint,3,opt,name=priority" json:"priority,omitempty"`
}

func (x *SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration) Reset() {
	*x = SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sharing_message_specifics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration) ProtoMessage() {}

func (x *SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_sharing_message_specifics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration.ProtoReflect.Descriptor instead.
func (*SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration) Descriptor() ([]byte, []int) {
	return file_sharing_message_specifics_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

func (x *SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration) GetTtl() int32 {
	if x != nil && x.Ttl != nil {
		return *x.Ttl
	}
	return 0
}

func (x *SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration) GetPriority() int32 {
	if x != nil && x.Priority != nil {
		return *x.Priority
	}
	return 0
}

var File_sharing_message_specifics_proto protoreflect.FileDescriptor

var file_sharing_message_specifics_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x22, 0xcc, 0x03, 0x0a, 0x17, 0x53,
	0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x6a, 0x0a, 0x15, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x53,
	0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x14, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x8b, 0x02, 0x0a, 0x14,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x61, 0x0a, 0x03, 0x66, 0x63, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x4d, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x53, 0x68, 0x61, 0x72,
	0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66,
	0x69, 0x63, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x43, 0x4d, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x00, 0x52, 0x03, 0x66, 0x63, 0x6d, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x1a, 0x5d, 0x0a, 0x17, 0x46, 0x43, 0x4d, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x74, 0x74, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x42, 0x17, 0x0a, 0x15, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x9d, 0x03, 0x0a, 0x19, 0x53, 0x68,
	0x61, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x4b, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x73, 0x79,
	0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x22, 0xb2, 0x02, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54,
	0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10,
	0x02, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x03, 0x12,
	0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x04,
	0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x45, 0x58, 0x48,
	0x41, 0x55, 0x53, 0x54, 0x45, 0x44, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x4e, 0x41, 0x55,
	0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x45, 0x44, 0x10, 0x06, 0x12, 0x15, 0x0a,
	0x11, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x4e, 0x49,
	0x45, 0x44, 0x10, 0x07, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x54, 0x55, 0x52,
	0x4e, 0x45, 0x44, 0x5f, 0x4f, 0x46, 0x46, 0x10, 0x08, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x59, 0x4e,
	0x43, 0x5f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x09, 0x12, 0x28, 0x0a, 0x24, 0x44, 0x45, 0x50, 0x52, 0x45, 0x43, 0x41, 0x54, 0x45, 0x44, 0x5f,
	0x53, 0x59, 0x4e, 0x43, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x4f, 0x52, 0x5f, 0x41,
	0x55, 0x54, 0x48, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x0a, 0x12, 0x10, 0x0a, 0x0c, 0x53,
	0x59, 0x4e, 0x43, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x0b, 0x12, 0x15, 0x0a,
	0x11, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x0c, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x41, 0x55, 0x54,
	0x48, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x0d, 0x42, 0x2b, 0x0a, 0x25, 0x6f, 0x72, 0x67,
	0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x48, 0x03, 0x50, 0x01,
}

var (
	file_sharing_message_specifics_proto_rawDescOnce sync.Once
	file_sharing_message_specifics_proto_rawDescData = file_sharing_message_specifics_proto_rawDesc
)

func file_sharing_message_specifics_proto_rawDescGZIP() []byte {
	file_sharing_message_specifics_proto_rawDescOnce.Do(func() {
		file_sharing_message_specifics_proto_rawDescData = protoimpl.X.CompressGZIP(file_sharing_message_specifics_proto_rawDescData)
	})
	return file_sharing_message_specifics_proto_rawDescData
}

var file_sharing_message_specifics_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sharing_message_specifics_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sharing_message_specifics_proto_goTypes = []interface{}{
	(SharingMessageCommitError_ErrorCode)(0),                                     // 0: sync_pb.SharingMessageCommitError.ErrorCode
	(*SharingMessageSpecifics)(nil),                                              // 1: sync_pb.SharingMessageSpecifics
	(*SharingMessageCommitError)(nil),                                            // 2: sync_pb.SharingMessageCommitError
	(*SharingMessageSpecifics_ChannelConfiguration)(nil),                         // 3: sync_pb.SharingMessageSpecifics.ChannelConfiguration
	(*SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration)(nil), // 4: sync_pb.SharingMessageSpecifics.ChannelConfiguration.FCMChannelConfiguration
}
var file_sharing_message_specifics_proto_depIdxs = []int32{
	3, // 0: sync_pb.SharingMessageSpecifics.channel_configuration:type_name -> sync_pb.SharingMessageSpecifics.ChannelConfiguration
	0, // 1: sync_pb.SharingMessageCommitError.error_code:type_name -> sync_pb.SharingMessageCommitError.ErrorCode
	4, // 2: sync_pb.SharingMessageSpecifics.ChannelConfiguration.fcm:type_name -> sync_pb.SharingMessageSpecifics.ChannelConfiguration.FCMChannelConfiguration
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_sharing_message_specifics_proto_init() }
func file_sharing_message_specifics_proto_init() {
	if File_sharing_message_specifics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sharing_message_specifics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SharingMessageSpecifics); i {
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
		file_sharing_message_specifics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SharingMessageCommitError); i {
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
		file_sharing_message_specifics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SharingMessageSpecifics_ChannelConfiguration); i {
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
		file_sharing_message_specifics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration); i {
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
	file_sharing_message_specifics_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*SharingMessageSpecifics_ChannelConfiguration_Fcm)(nil),
		(*SharingMessageSpecifics_ChannelConfiguration_Server)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sharing_message_specifics_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sharing_message_specifics_proto_goTypes,
		DependencyIndexes: file_sharing_message_specifics_proto_depIdxs,
		EnumInfos:         file_sharing_message_specifics_proto_enumTypes,
		MessageInfos:      file_sharing_message_specifics_proto_msgTypes,
	}.Build()
	File_sharing_message_specifics_proto = out.File
	file_sharing_message_specifics_proto_rawDesc = nil
	file_sharing_message_specifics_proto_goTypes = nil
	file_sharing_message_specifics_proto_depIdxs = nil
}
