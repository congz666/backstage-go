// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.5.1
// source: otherModels.proto

package services

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

type CarouselModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	ID uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"id"`
	// @inject_tag: json:"product_id"
	ProductID uint32 `protobuf:"varint,3,opt,name=ProductID,proto3" json:"product_id"`
	// @inject_tag: json:"img_path"
	ImgPath string `protobuf:"bytes,4,opt,name=ImgPath,proto3" json:"img_path"`
	// @inject_tag: json:"created_at"
	CreatedAt int64 `protobuf:"varint,5,opt,name=CreatedAt,proto3" json:"created_at"`
	// @inject_tag: json:"updated_at"
	UpdatedAt int64 `protobuf:"varint,6,opt,name=UpdatedAt,proto3" json:"updated_at"`
	// @inject_tag: json:"deleted_at"
	DeletedAt int64 `protobuf:"varint,7,opt,name=DeletedAt,proto3" json:"deleted_at"`
}

func (x *CarouselModel) Reset() {
	*x = CarouselModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_otherModels_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarouselModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarouselModel) ProtoMessage() {}

func (x *CarouselModel) ProtoReflect() protoreflect.Message {
	mi := &file_otherModels_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarouselModel.ProtoReflect.Descriptor instead.
func (*CarouselModel) Descriptor() ([]byte, []int) {
	return file_otherModels_proto_rawDescGZIP(), []int{0}
}

func (x *CarouselModel) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *CarouselModel) GetProductID() uint32 {
	if x != nil {
		return x.ProductID
	}
	return 0
}

func (x *CarouselModel) GetImgPath() string {
	if x != nil {
		return x.ImgPath
	}
	return ""
}

func (x *CarouselModel) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *CarouselModel) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *CarouselModel) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

type NoticeModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	ID uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"id"`
	// @inject_tag: json:"text"
	Text string `protobuf:"bytes,2,opt,name=Text,proto3" json:"text"`
	// @inject_tag: json:"created_at"
	CreatedAt int64 `protobuf:"varint,3,opt,name=CreatedAt,proto3" json:"created_at"`
	// @inject_tag: json:"updated_at"
	UpdatedAt int64 `protobuf:"varint,4,opt,name=UpdatedAt,proto3" json:"updated_at"`
	// @inject_tag: json:"deleted_at"
	DeletedAt int64 `protobuf:"varint,5,opt,name=DeletedAt,proto3" json:"deleted_at"`
}

func (x *NoticeModel) Reset() {
	*x = NoticeModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_otherModels_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoticeModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoticeModel) ProtoMessage() {}

func (x *NoticeModel) ProtoReflect() protoreflect.Message {
	mi := &file_otherModels_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoticeModel.ProtoReflect.Descriptor instead.
func (*NoticeModel) Descriptor() ([]byte, []int) {
	return file_otherModels_proto_rawDescGZIP(), []int{1}
}

func (x *NoticeModel) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *NoticeModel) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *NoticeModel) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *NoticeModel) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *NoticeModel) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

type CategoryModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	ID uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"id"`
	// @inject_tag: json:"category_id"
	CategoryID uint32 `protobuf:"varint,2,opt,name=CategoryID,proto3" json:"category_id"`
	// @inject_tag: json:"category_name"
	CategoryName string `protobuf:"bytes,3,opt,name=CategoryName,proto3" json:"category_name"`
	// @inject_tag: json:"num"
	Num uint32 `protobuf:"varint,4,opt,name=Num,proto3" json:"num"`
	// @inject_tag: json:"created_at"
	CreatedAt int64 `protobuf:"varint,5,opt,name=CreatedAt,proto3" json:"created_at"`
	// @inject_tag: json:"updated_at"
	UpdatedAt int64 `protobuf:"varint,6,opt,name=UpdatedAt,proto3" json:"updated_at"`
	// @inject_tag: json:"deleted_at"
	DeletedAt int64 `protobuf:"varint,7,opt,name=DeletedAt,proto3" json:"deleted_at"`
}

func (x *CategoryModel) Reset() {
	*x = CategoryModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_otherModels_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryModel) ProtoMessage() {}

func (x *CategoryModel) ProtoReflect() protoreflect.Message {
	mi := &file_otherModels_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryModel.ProtoReflect.Descriptor instead.
func (*CategoryModel) Descriptor() ([]byte, []int) {
	return file_otherModels_proto_rawDescGZIP(), []int{2}
}

func (x *CategoryModel) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *CategoryModel) GetCategoryID() uint32 {
	if x != nil {
		return x.CategoryID
	}
	return 0
}

func (x *CategoryModel) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

func (x *CategoryModel) GetNum() uint32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *CategoryModel) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *CategoryModel) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *CategoryModel) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

var File_otherModels_proto protoreflect.FileDescriptor

var file_otherModels_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0xb1, 0x01,
	0x0a, 0x0d, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x18, 0x0a,
	0x07, 0x49, 0x6d, 0x67, 0x50, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x49, 0x6d, 0x67, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x8b, 0x01, 0x0a, 0x0b, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x54, 0x65, 0x78, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0xcf, 0x01, 0x0a, 0x0d, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49,
	0x44, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x4e, 0x75, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_otherModels_proto_rawDescOnce sync.Once
	file_otherModels_proto_rawDescData = file_otherModels_proto_rawDesc
)

func file_otherModels_proto_rawDescGZIP() []byte {
	file_otherModels_proto_rawDescOnce.Do(func() {
		file_otherModels_proto_rawDescData = protoimpl.X.CompressGZIP(file_otherModels_proto_rawDescData)
	})
	return file_otherModels_proto_rawDescData
}

var file_otherModels_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_otherModels_proto_goTypes = []interface{}{
	(*CarouselModel)(nil), // 0: services.CarouselModel
	(*NoticeModel)(nil),   // 1: services.NoticeModel
	(*CategoryModel)(nil), // 2: services.CategoryModel
}
var file_otherModels_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_otherModels_proto_init() }
func file_otherModels_proto_init() {
	if File_otherModels_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_otherModels_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarouselModel); i {
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
		file_otherModels_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoticeModel); i {
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
		file_otherModels_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryModel); i {
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
			RawDescriptor: file_otherModels_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_otherModels_proto_goTypes,
		DependencyIndexes: file_otherModels_proto_depIdxs,
		MessageInfos:      file_otherModels_proto_msgTypes,
	}.Build()
	File_otherModels_proto = out.File
	file_otherModels_proto_rawDesc = nil
	file_otherModels_proto_goTypes = nil
	file_otherModels_proto_depIdxs = nil
}