// Copyright (c) 2012 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// Sync protocol datatype extension for extensions.

// If you change or add any fields in this file, update proto_visitors.h and
// potentially proto_enum_conversions.{h, cc}.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.4
// source: extension_specifics.proto

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

// Properties of extension sync objects.
//
// Merge policy: the settings for the higher version number win; in
// the case of a tie, server wins.
type ExtensionSpecifics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Globally unique id for this extension.
	Id *string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// The known installed version.
	Version *string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	// Auto-update URL to use for this extension.  May be blank, in
	// which case the default one (i.e., the one for the Chrome
	// Extensions Gallery) is used.
	UpdateUrl *string `protobuf:"bytes,3,opt,name=update_url,json=updateUrl" json:"update_url,omitempty"`
	// Whether or not this extension is enabled.
	Enabled *bool `protobuf:"varint,4,opt,name=enabled" json:"enabled,omitempty"`
	// Whether or not this extension is enabled in incognito mode.
	IncognitoEnabled *bool `protobuf:"varint,5,opt,name=incognito_enabled,json=incognitoEnabled" json:"incognito_enabled,omitempty"`
	// The name of the extension. Used for bookmark apps.
	Name *string `protobuf:"bytes,6,opt,name=name" json:"name,omitempty"`
	// Whether this extension was installed remotely, and hasn't been approved by
	// a user in chrome yet.
	RemoteInstall *bool `protobuf:"varint,7,opt,name=remote_install,json=remoteInstall" json:"remote_install,omitempty"`
	// DEPRECATED. See https://crbug.com/1014183.
	//
	// Deprecated: Do not use.
	InstalledByCustodian *bool `protobuf:"varint,8,opt,name=installed_by_custodian,json=installedByCustodian" json:"installed_by_custodian,omitempty"`
	// DEPRECATED. See https://crbug.com/839681.
	//
	// Deprecated: Do not use.
	AllUrlsEnabled *bool `protobuf:"varint,9,opt,name=all_urls_enabled,json=allUrlsEnabled" json:"all_urls_enabled,omitempty"`
	// Bitmask of the set of reasons why the extension is disabled (see
	// extensions::disable_reason::DisableReason). Only relevant when enabled ==
	// false. Note that old clients (<M45) won't set this, even when enabled is
	// false.
	DisableReasons *int32 `protobuf:"varint,10,opt,name=disable_reasons,json=disableReasons" json:"disable_reasons,omitempty"`
}

func (x *ExtensionSpecifics) Reset() {
	*x = ExtensionSpecifics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extension_specifics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtensionSpecifics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtensionSpecifics) ProtoMessage() {}

func (x *ExtensionSpecifics) ProtoReflect() protoreflect.Message {
	mi := &file_extension_specifics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtensionSpecifics.ProtoReflect.Descriptor instead.
func (*ExtensionSpecifics) Descriptor() ([]byte, []int) {
	return file_extension_specifics_proto_rawDescGZIP(), []int{0}
}

func (x *ExtensionSpecifics) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *ExtensionSpecifics) GetVersion() string {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return ""
}

func (x *ExtensionSpecifics) GetUpdateUrl() string {
	if x != nil && x.UpdateUrl != nil {
		return *x.UpdateUrl
	}
	return ""
}

func (x *ExtensionSpecifics) GetEnabled() bool {
	if x != nil && x.Enabled != nil {
		return *x.Enabled
	}
	return false
}

func (x *ExtensionSpecifics) GetIncognitoEnabled() bool {
	if x != nil && x.IncognitoEnabled != nil {
		return *x.IncognitoEnabled
	}
	return false
}

func (x *ExtensionSpecifics) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *ExtensionSpecifics) GetRemoteInstall() bool {
	if x != nil && x.RemoteInstall != nil {
		return *x.RemoteInstall
	}
	return false
}

// Deprecated: Do not use.
func (x *ExtensionSpecifics) GetInstalledByCustodian() bool {
	if x != nil && x.InstalledByCustodian != nil {
		return *x.InstalledByCustodian
	}
	return false
}

// Deprecated: Do not use.
func (x *ExtensionSpecifics) GetAllUrlsEnabled() bool {
	if x != nil && x.AllUrlsEnabled != nil {
		return *x.AllUrlsEnabled
	}
	return false
}

func (x *ExtensionSpecifics) GetDisableReasons() int32 {
	if x != nil && x.DisableReasons != nil {
		return *x.DisableReasons
	}
	return 0
}

var File_extension_specifics_proto protoreflect.FileDescriptor

var file_extension_specifics_proto_rawDesc = []byte{
	0x0a, 0x19, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x66, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x79, 0x6e,
	0x63, 0x5f, 0x70, 0x62, 0x22, 0xf0, 0x02, 0x0a, 0x12, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x2b,
	0x0a, 0x11, 0x69, 0x6e, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x6f, 0x5f, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x6e, 0x63, 0x6f, 0x67,
	0x6e, 0x69, 0x74, 0x6f, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x25, 0x0a, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x12, 0x38, 0x0a, 0x16, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x69, 0x61, 0x6e,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x42, 0x02, 0x18, 0x01, 0x52, 0x14, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6c, 0x6c, 0x65, 0x64, 0x42, 0x79, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x69, 0x61, 0x6e,
	0x12, 0x2c, 0x0a, 0x10, 0x61, 0x6c, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x5f, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0e,
	0x61, 0x6c, 0x6c, 0x55, 0x72, 0x6c, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x27,
	0x0a, 0x0f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x42, 0x2b, 0x0a, 0x25, 0x6f, 0x72, 0x67, 0x2e, 0x63,
	0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x48, 0x03, 0x50, 0x01,
}

var (
	file_extension_specifics_proto_rawDescOnce sync.Once
	file_extension_specifics_proto_rawDescData = file_extension_specifics_proto_rawDesc
)

func file_extension_specifics_proto_rawDescGZIP() []byte {
	file_extension_specifics_proto_rawDescOnce.Do(func() {
		file_extension_specifics_proto_rawDescData = protoimpl.X.CompressGZIP(file_extension_specifics_proto_rawDescData)
	})
	return file_extension_specifics_proto_rawDescData
}

var file_extension_specifics_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_extension_specifics_proto_goTypes = []interface{}{
	(*ExtensionSpecifics)(nil), // 0: sync_pb.ExtensionSpecifics
}
var file_extension_specifics_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_extension_specifics_proto_init() }
func file_extension_specifics_proto_init() {
	if File_extension_specifics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_extension_specifics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtensionSpecifics); i {
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
			RawDescriptor: file_extension_specifics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_extension_specifics_proto_goTypes,
		DependencyIndexes: file_extension_specifics_proto_depIdxs,
		MessageInfos:      file_extension_specifics_proto_msgTypes,
	}.Build()
	File_extension_specifics_proto = out.File
	file_extension_specifics_proto_rawDesc = nil
	file_extension_specifics_proto_goTypes = nil
	file_extension_specifics_proto_depIdxs = nil
}
