// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: v1/apb/rpc_update_anime_image.proto

package apbv1

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

type UpdateAnimeImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageID    int64   `protobuf:"varint,1,opt,name=imageID,proto3" json:"imageID,omitempty"`
	Host       *string `protobuf:"bytes,2,opt,name=host,proto3,oneof" json:"host,omitempty"`
	Url        *string `protobuf:"bytes,3,opt,name=url,proto3,oneof" json:"url,omitempty"`
	Thumbnails *string `protobuf:"bytes,4,opt,name=thumbnails,proto3,oneof" json:"thumbnails,omitempty"`
	BlurHash   *string `protobuf:"bytes,5,opt,name=blurHash,proto3,oneof" json:"blurHash,omitempty"`
	Height     *uint32 `protobuf:"varint,6,opt,name=height,proto3,oneof" json:"height,omitempty"`
	Width      *uint32 `protobuf:"varint,7,opt,name=width,proto3,oneof" json:"width,omitempty"`
}

func (x *UpdateAnimeImageRequest) Reset() {
	*x = UpdateAnimeImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_apb_rpc_update_anime_image_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAnimeImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAnimeImageRequest) ProtoMessage() {}

func (x *UpdateAnimeImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_apb_rpc_update_anime_image_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAnimeImageRequest.ProtoReflect.Descriptor instead.
func (*UpdateAnimeImageRequest) Descriptor() ([]byte, []int) {
	return file_v1_apb_rpc_update_anime_image_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateAnimeImageRequest) GetImageID() int64 {
	if x != nil {
		return x.ImageID
	}
	return 0
}

func (x *UpdateAnimeImageRequest) GetHost() string {
	if x != nil && x.Host != nil {
		return *x.Host
	}
	return ""
}

func (x *UpdateAnimeImageRequest) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *UpdateAnimeImageRequest) GetThumbnails() string {
	if x != nil && x.Thumbnails != nil {
		return *x.Thumbnails
	}
	return ""
}

func (x *UpdateAnimeImageRequest) GetBlurHash() string {
	if x != nil && x.BlurHash != nil {
		return *x.BlurHash
	}
	return ""
}

func (x *UpdateAnimeImageRequest) GetHeight() uint32 {
	if x != nil && x.Height != nil {
		return *x.Height
	}
	return 0
}

func (x *UpdateAnimeImageRequest) GetWidth() uint32 {
	if x != nil && x.Width != nil {
		return *x.Width
	}
	return 0
}

type UpdateAnimeImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimeImage *ImageResponse `protobuf:"bytes,1,opt,name=animeImage,proto3" json:"animeImage,omitempty"`
}

func (x *UpdateAnimeImageResponse) Reset() {
	*x = UpdateAnimeImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_apb_rpc_update_anime_image_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAnimeImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAnimeImageResponse) ProtoMessage() {}

func (x *UpdateAnimeImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_apb_rpc_update_anime_image_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAnimeImageResponse.ProtoReflect.Descriptor instead.
func (*UpdateAnimeImageResponse) Descriptor() ([]byte, []int) {
	return file_v1_apb_rpc_update_anime_image_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateAnimeImageResponse) GetAnimeImage() *ImageResponse {
	if x != nil {
		return x.AnimeImage
	}
	return nil
}

var File_v1_apb_rpc_update_anime_image_proto protoreflect.FileDescriptor

var file_v1_apb_rpc_update_anime_image_proto_rawDesc = []byte{
	0x0a, 0x23, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x62, 0x76, 0x31, 0x1a,
	0x18, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x62, 0x2f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x02, 0x0a, 0x17, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x44, 0x12,
	0x17, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12,
	0x23, 0x0a, 0x0a, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0a, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c,
	0x73, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x62, 0x6c, 0x75, 0x72, 0x48, 0x61, 0x73, 0x68,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x08, 0x62, 0x6c, 0x75, 0x72, 0x48, 0x61,
	0x73, 0x68, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0d, 0x48, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x75, 0x72, 0x6c, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x0b, 0x0a,
	0x09, 0x5f, 0x62, 0x6c, 0x75, 0x72, 0x48, 0x61, 0x73, 0x68, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x77, 0x69, 0x64, 0x74, 0x68, 0x22,
	0x53, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x61,
	0x6e, 0x69, 0x6d, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69, 0x6e, 0x65, 0x2d, 0x66, 0x6c, 0x75,
	0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x70, 0x62, 0x3b, 0x61, 0x70, 0x62, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_v1_apb_rpc_update_anime_image_proto_rawDescOnce sync.Once
	file_v1_apb_rpc_update_anime_image_proto_rawDescData = file_v1_apb_rpc_update_anime_image_proto_rawDesc
)

func file_v1_apb_rpc_update_anime_image_proto_rawDescGZIP() []byte {
	file_v1_apb_rpc_update_anime_image_proto_rawDescOnce.Do(func() {
		file_v1_apb_rpc_update_anime_image_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_apb_rpc_update_anime_image_proto_rawDescData)
	})
	return file_v1_apb_rpc_update_anime_image_proto_rawDescData
}

var file_v1_apb_rpc_update_anime_image_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_apb_rpc_update_anime_image_proto_goTypes = []interface{}{
	(*UpdateAnimeImageRequest)(nil),  // 0: v1.apbv1.UpdateAnimeImageRequest
	(*UpdateAnimeImageResponse)(nil), // 1: v1.apbv1.UpdateAnimeImageResponse
	(*ImageResponse)(nil),            // 2: v1.apbv1.ImageResponse
}
var file_v1_apb_rpc_update_anime_image_proto_depIdxs = []int32{
	2, // 0: v1.apbv1.UpdateAnimeImageResponse.animeImage:type_name -> v1.apbv1.ImageResponse
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_apb_rpc_update_anime_image_proto_init() }
func file_v1_apb_rpc_update_anime_image_proto_init() {
	if File_v1_apb_rpc_update_anime_image_proto != nil {
		return
	}
	file_v1_apb_anime_image_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_apb_rpc_update_anime_image_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAnimeImageRequest); i {
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
		file_v1_apb_rpc_update_anime_image_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAnimeImageResponse); i {
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
	file_v1_apb_rpc_update_anime_image_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_apb_rpc_update_anime_image_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_apb_rpc_update_anime_image_proto_goTypes,
		DependencyIndexes: file_v1_apb_rpc_update_anime_image_proto_depIdxs,
		MessageInfos:      file_v1_apb_rpc_update_anime_image_proto_msgTypes,
	}.Build()
	File_v1_apb_rpc_update_anime_image_proto = out.File
	file_v1_apb_rpc_update_anime_image_proto_rawDesc = nil
	file_v1_apb_rpc_update_anime_image_proto_goTypes = nil
	file_v1_apb_rpc_update_anime_image_proto_depIdxs = nil
}
