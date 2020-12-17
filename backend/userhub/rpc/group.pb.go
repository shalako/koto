// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: group.proto

package rpc

import (
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

type GroupAddGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	AvatarId    string `protobuf:"bytes,3,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id,omitempty"`
	IsPublic    bool   `protobuf:"varint,4,opt,name=is_public,json=isPublic,proto3" json:"is_public,omitempty"`
}

func (x *GroupAddGroupRequest) Reset() {
	*x = GroupAddGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupAddGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupAddGroupRequest) ProtoMessage() {}

func (x *GroupAddGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupAddGroupRequest.ProtoReflect.Descriptor instead.
func (*GroupAddGroupRequest) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{0}
}

func (x *GroupAddGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GroupAddGroupRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GroupAddGroupRequest) GetAvatarId() string {
	if x != nil {
		return x.AvatarId
	}
	return ""
}

func (x *GroupAddGroupRequest) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

type GroupAddGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group *Group `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *GroupAddGroupResponse) Reset() {
	*x = GroupAddGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupAddGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupAddGroupResponse) ProtoMessage() {}

func (x *GroupAddGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupAddGroupResponse.ProtoReflect.Descriptor instead.
func (*GroupAddGroupResponse) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{1}
}

func (x *GroupAddGroupResponse) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

type GroupEditGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId            string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	DescriptionChanged bool   `protobuf:"varint,2,opt,name=description_changed,json=descriptionChanged,proto3" json:"description_changed,omitempty"`
	Description        string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	AvatarChanged      bool   `protobuf:"varint,4,opt,name=avatar_changed,json=avatarChanged,proto3" json:"avatar_changed,omitempty"`
	AvatarId           string `protobuf:"bytes,5,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id,omitempty"`
	IsPublicChanged    bool   `protobuf:"varint,6,opt,name=is_public_changed,json=isPublicChanged,proto3" json:"is_public_changed,omitempty"`
	IsPublic           bool   `protobuf:"varint,7,opt,name=is_public,json=isPublic,proto3" json:"is_public,omitempty"`
}

func (x *GroupEditGroupRequest) Reset() {
	*x = GroupEditGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupEditGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupEditGroupRequest) ProtoMessage() {}

func (x *GroupEditGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupEditGroupRequest.ProtoReflect.Descriptor instead.
func (*GroupEditGroupRequest) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{2}
}

func (x *GroupEditGroupRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *GroupEditGroupRequest) GetDescriptionChanged() bool {
	if x != nil {
		return x.DescriptionChanged
	}
	return false
}

func (x *GroupEditGroupRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GroupEditGroupRequest) GetAvatarChanged() bool {
	if x != nil {
		return x.AvatarChanged
	}
	return false
}

func (x *GroupEditGroupRequest) GetAvatarId() string {
	if x != nil {
		return x.AvatarId
	}
	return ""
}

func (x *GroupEditGroupRequest) GetIsPublicChanged() bool {
	if x != nil {
		return x.IsPublicChanged
	}
	return false
}

func (x *GroupEditGroupRequest) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

var File_group_proto protoreflect.FileDescriptor

var file_group_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72,
	0x70, 0x63, 0x1a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x86, 0x01, 0x0a, 0x14, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x64, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b,
	0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69,
	0x73, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x22, 0x39, 0x0a, 0x15, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x41, 0x64, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x20, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x22, 0x92, 0x02, 0x0a, 0x15, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x45, 0x64, 0x69,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x13, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x12,
	0x2a, 0x0a, 0x11, 0x69, 0x73, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69,
	0x73, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x32, 0x86, 0x01, 0x0a, 0x0c, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x41, 0x64, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x19, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x41, 0x64, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x64, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x09,
	0x45, 0x64, 0x69, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1a, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x45, 0x64, 0x69, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2e, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_group_proto_rawDescOnce sync.Once
	file_group_proto_rawDescData = file_group_proto_rawDesc
)

func file_group_proto_rawDescGZIP() []byte {
	file_group_proto_rawDescOnce.Do(func() {
		file_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_group_proto_rawDescData)
	})
	return file_group_proto_rawDescData
}

var file_group_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_group_proto_goTypes = []interface{}{
	(*GroupAddGroupRequest)(nil),  // 0: rpc.GroupAddGroupRequest
	(*GroupAddGroupResponse)(nil), // 1: rpc.GroupAddGroupResponse
	(*GroupEditGroupRequest)(nil), // 2: rpc.GroupEditGroupRequest
	(*Group)(nil),                 // 3: rpc.Group
	(*Empty)(nil),                 // 4: rpc.Empty
}
var file_group_proto_depIdxs = []int32{
	3, // 0: rpc.GroupAddGroupResponse.group:type_name -> rpc.Group
	0, // 1: rpc.GroupService.AddGroup:input_type -> rpc.GroupAddGroupRequest
	2, // 2: rpc.GroupService.EditGroup:input_type -> rpc.GroupEditGroupRequest
	1, // 3: rpc.GroupService.AddGroup:output_type -> rpc.GroupAddGroupResponse
	4, // 4: rpc.GroupService.EditGroup:output_type -> rpc.Empty
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_group_proto_init() }
func file_group_proto_init() {
	if File_group_proto != nil {
		return
	}
	file_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupAddGroupRequest); i {
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
		file_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupAddGroupResponse); i {
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
		file_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupEditGroupRequest); i {
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
			RawDescriptor: file_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_group_proto_goTypes,
		DependencyIndexes: file_group_proto_depIdxs,
		MessageInfos:      file_group_proto_msgTypes,
	}.Build()
	File_group_proto = out.File
	file_group_proto_rawDesc = nil
	file_group_proto_goTypes = nil
	file_group_proto_depIdxs = nil
}
