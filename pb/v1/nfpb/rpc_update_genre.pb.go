// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: v1/nfpb/rpc_update_genre.proto

package nfpbv1

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

type UpdateGenreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GenreID   int32   `protobuf:"varint,1,opt,name=genreID,proto3" json:"genreID,omitempty"`
	GenreName *string `protobuf:"bytes,2,opt,name=genreName,proto3,oneof" json:"genreName,omitempty"`
}

func (x *UpdateGenreRequest) Reset() {
	*x = UpdateGenreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_nfpb_rpc_update_genre_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGenreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGenreRequest) ProtoMessage() {}

func (x *UpdateGenreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_nfpb_rpc_update_genre_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGenreRequest.ProtoReflect.Descriptor instead.
func (*UpdateGenreRequest) Descriptor() ([]byte, []int) {
	return file_v1_nfpb_rpc_update_genre_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateGenreRequest) GetGenreID() int32 {
	if x != nil {
		return x.GenreID
	}
	return 0
}

func (x *UpdateGenreRequest) GetGenreName() string {
	if x != nil && x.GenreName != nil {
		return *x.GenreName
	}
	return ""
}

type UpdateGenreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Genre *Genre `protobuf:"bytes,1,opt,name=genre,proto3" json:"genre,omitempty"`
}

func (x *UpdateGenreResponse) Reset() {
	*x = UpdateGenreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_nfpb_rpc_update_genre_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGenreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGenreResponse) ProtoMessage() {}

func (x *UpdateGenreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_nfpb_rpc_update_genre_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGenreResponse.ProtoReflect.Descriptor instead.
func (*UpdateGenreResponse) Descriptor() ([]byte, []int) {
	return file_v1_nfpb_rpc_update_genre_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateGenreResponse) GetGenre() *Genre {
	if x != nil {
		return x.Genre
	}
	return nil
}

var File_v1_nfpb_rpc_update_genre_proto protoreflect.FileDescriptor

var file_v1_nfpb_rpc_update_genre_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x76, 0x31, 0x2f, 0x6e, 0x66, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x76, 0x31, 0x2e, 0x6e, 0x66, 0x70, 0x62, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x31, 0x2f,
	0x6e, 0x66, 0x70, 0x62, 0x2f, 0x6d, 0x73, 0x67, 0x5f, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x65,
	0x6e, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x65,
	0x6e, 0x72, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67, 0x65, 0x6e,
	0x72, 0x65, 0x49, 0x44, 0x12, 0x21, 0x0a, 0x09, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x67, 0x65, 0x6e, 0x72, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x67, 0x65, 0x6e, 0x72,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3d, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47,
	0x65, 0x6e, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05,
	0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76, 0x31,
	0x2e, 0x6e, 0x66, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x52, 0x05, 0x67,
	0x65, 0x6e, 0x72, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69, 0x6e, 0x65, 0x2d, 0x66, 0x6c, 0x75,
	0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f,
	0x6e, 0x66, 0x70, 0x62, 0x3b, 0x6e, 0x66, 0x70, 0x62, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_v1_nfpb_rpc_update_genre_proto_rawDescOnce sync.Once
	file_v1_nfpb_rpc_update_genre_proto_rawDescData = file_v1_nfpb_rpc_update_genre_proto_rawDesc
)

func file_v1_nfpb_rpc_update_genre_proto_rawDescGZIP() []byte {
	file_v1_nfpb_rpc_update_genre_proto_rawDescOnce.Do(func() {
		file_v1_nfpb_rpc_update_genre_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_nfpb_rpc_update_genre_proto_rawDescData)
	})
	return file_v1_nfpb_rpc_update_genre_proto_rawDescData
}

var file_v1_nfpb_rpc_update_genre_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_nfpb_rpc_update_genre_proto_goTypes = []interface{}{
	(*UpdateGenreRequest)(nil),  // 0: v1.nfpbv1.UpdateGenreRequest
	(*UpdateGenreResponse)(nil), // 1: v1.nfpbv1.UpdateGenreResponse
	(*Genre)(nil),               // 2: v1.nfpbv1.Genre
}
var file_v1_nfpb_rpc_update_genre_proto_depIdxs = []int32{
	2, // 0: v1.nfpbv1.UpdateGenreResponse.genre:type_name -> v1.nfpbv1.Genre
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_nfpb_rpc_update_genre_proto_init() }
func file_v1_nfpb_rpc_update_genre_proto_init() {
	if File_v1_nfpb_rpc_update_genre_proto != nil {
		return
	}
	file_v1_nfpb_msg_genre_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_nfpb_rpc_update_genre_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGenreRequest); i {
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
		file_v1_nfpb_rpc_update_genre_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGenreResponse); i {
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
	file_v1_nfpb_rpc_update_genre_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_nfpb_rpc_update_genre_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_nfpb_rpc_update_genre_proto_goTypes,
		DependencyIndexes: file_v1_nfpb_rpc_update_genre_proto_depIdxs,
		MessageInfos:      file_v1_nfpb_rpc_update_genre_proto_msgTypes,
	}.Build()
	File_v1_nfpb_rpc_update_genre_proto = out.File
	file_v1_nfpb_rpc_update_genre_proto_rawDesc = nil
	file_v1_nfpb_rpc_update_genre_proto_goTypes = nil
	file_v1_nfpb_rpc_update_genre_proto_depIdxs = nil
}
