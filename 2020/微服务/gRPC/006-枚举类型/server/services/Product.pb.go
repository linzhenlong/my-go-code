// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: Product.proto

package services

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// 枚举类型,商品所在区域a,b,c 三个区域
// 枚举类型的值必须要影射为数值类型，并且第一个必须为零，代表默认值
type ProductAreas int32

const (
	ProductAreas_BEI_JING  ProductAreas = 0 //
	ProductAreas_HE_BEI    ProductAreas = 1
	ProductAreas_SHANG_HAI ProductAreas = 2
	ProductAreas_SI_CHUANG ProductAreas = 3
	ProductAreas_WU_HANG   ProductAreas = 4
	ProductAreas_SHAN_DONG ProductAreas = 6
)

// Enum value maps for ProductAreas.
var (
	ProductAreas_name = map[int32]string{
		0: "BEI_JING",
		1: "HE_BEI",
		2: "SHANG_HAI",
		3: "SI_CHUANG",
		4: "WU_HANG",
		6: "SHAN_DONG",
	}
	ProductAreas_value = map[string]int32{
		"BEI_JING":  0,
		"HE_BEI":    1,
		"SHANG_HAI": 2,
		"SI_CHUANG": 3,
		"WU_HANG":   4,
		"SHAN_DONG": 6,
	}
)

func (x ProductAreas) Enum() *ProductAreas {
	p := new(ProductAreas)
	*p = x
	return p
}

func (x ProductAreas) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProductAreas) Descriptor() protoreflect.EnumDescriptor {
	return file_Product_proto_enumTypes[0].Descriptor()
}

func (ProductAreas) Type() protoreflect.EnumType {
	return &file_Product_proto_enumTypes[0]
}

func (x ProductAreas) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProductAreas.Descriptor instead.
func (ProductAreas) EnumDescriptor() ([]byte, []int) {
	return file_Product_proto_rawDescGZIP(), []int{0}
}

type ProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProName string       `protobuf:"bytes,1,opt,name=pro_name,json=proName,proto3" json:"pro_name,omitempty"`
	Size    int32        `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	ProId   int32        `protobuf:"varint,3,opt,name=pro_id,json=proId,proto3" json:"pro_id,omitempty"`
	ProArea ProductAreas `protobuf:"varint,4,opt,name=pro_area,json=proArea,proto3,enum=services.ProductAreas" json:"pro_area,omitempty"`
}

func (x *ProductRequest) Reset() {
	*x = ProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductRequest) ProtoMessage() {}

func (x *ProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductRequest.ProtoReflect.Descriptor instead.
func (*ProductRequest) Descriptor() ([]byte, []int) {
	return file_Product_proto_rawDescGZIP(), []int{0}
}

func (x *ProductRequest) GetProName() string {
	if x != nil {
		return x.ProName
	}
	return ""
}

func (x *ProductRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ProductRequest) GetProId() int32 {
	if x != nil {
		return x.ProId
	}
	return 0
}

func (x *ProductRequest) GetProArea() ProductAreas {
	if x != nil {
		return x.ProArea
	}
	return ProductAreas_BEI_JING
}

type ProductInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProName  string       `protobuf:"bytes,1,opt,name=pro_name,json=proName,proto3" json:"pro_name,omitempty"`
	ProId    int32        `protobuf:"varint,2,opt,name=pro_id,json=proId,proto3" json:"pro_id,omitempty"`
	ProTag   []string     `protobuf:"bytes,3,rep,name=pro_tag,json=proTag,proto3" json:"pro_tag,omitempty"`
	ProPrice float32      `protobuf:"fixed32,4,opt,name=pro_price,json=proPrice,proto3" json:"pro_price,omitempty"`
	ProArea  ProductAreas `protobuf:"varint,5,opt,name=pro_area,json=proArea,proto3,enum=services.ProductAreas" json:"pro_area,omitempty"`
}

func (x *ProductInfo) Reset() {
	*x = ProductInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductInfo) ProtoMessage() {}

func (x *ProductInfo) ProtoReflect() protoreflect.Message {
	mi := &file_Product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductInfo.ProtoReflect.Descriptor instead.
func (*ProductInfo) Descriptor() ([]byte, []int) {
	return file_Product_proto_rawDescGZIP(), []int{1}
}

func (x *ProductInfo) GetProName() string {
	if x != nil {
		return x.ProName
	}
	return ""
}

func (x *ProductInfo) GetProId() int32 {
	if x != nil {
		return x.ProId
	}
	return 0
}

func (x *ProductInfo) GetProTag() []string {
	if x != nil {
		return x.ProTag
	}
	return nil
}

func (x *ProductInfo) GetProPrice() float32 {
	if x != nil {
		return x.ProPrice
	}
	return 0
}

func (x *ProductInfo) GetProArea() ProductAreas {
	if x != nil {
		return x.ProArea
	}
	return ProductAreas_BEI_JING
}

type ProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode int32          `protobuf:"varint,1,opt,name=err_code,json=errCode,proto3" json:"err_code,omitempty"`
	ErrMsg  string         `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
	Data    []*ProductInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ProductResponse) Reset() {
	*x = ProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductResponse) ProtoMessage() {}

func (x *ProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductResponse.ProtoReflect.Descriptor instead.
func (*ProductResponse) Descriptor() ([]byte, []int) {
	return file_Product_proto_rawDescGZIP(), []int{2}
}

func (x *ProductResponse) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *ProductResponse) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

func (x *ProductResponse) GetData() []*ProductInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_Product_proto protoreflect.FileDescriptor

var file_Product_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x70, 0x72, 0x6f,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x49, 0x64,
	0x12, 0x31, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x72, 0x65, 0x61, 0x73, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x41,
	0x72, 0x65, 0x61, 0x22, 0xa8, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x15,
	0x0a, 0x06, 0x70, 0x72, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x70, 0x72, 0x6f, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x5f, 0x74, 0x61, 0x67,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x54, 0x61, 0x67, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x72, 0x6f, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x41, 0x72, 0x65, 0x61, 0x73, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x41, 0x72, 0x65, 0x61, 0x22, 0x70,
	0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x65, 0x72, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65,
	0x72, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x2a, 0x62, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x72, 0x65, 0x61, 0x73,
	0x12, 0x0c, 0x0a, 0x08, 0x42, 0x45, 0x49, 0x5f, 0x4a, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x48, 0x45, 0x5f, 0x42, 0x45, 0x49, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x48,
	0x41, 0x4e, 0x47, 0x5f, 0x48, 0x41, 0x49, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x49, 0x5f,
	0x43, 0x48, 0x55, 0x41, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x55, 0x5f, 0x48,
	0x41, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x48, 0x41, 0x4e, 0x5f, 0x44, 0x4f,
	0x4e, 0x47, 0x10, 0x06, 0x32, 0xe4, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x5f,
	0x69, 0x64, 0x7d, 0x12, 0x66, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x19, 0x12, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f,
	0x6c, 0x69, 0x73, 0x74, 0x2f, 0x7b, 0x73, 0x69, 0x7a, 0x65, 0x7d, 0x42, 0x0c, 0x5a, 0x0a, 0x2e,
	0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_Product_proto_rawDescOnce sync.Once
	file_Product_proto_rawDescData = file_Product_proto_rawDesc
)

func file_Product_proto_rawDescGZIP() []byte {
	file_Product_proto_rawDescOnce.Do(func() {
		file_Product_proto_rawDescData = protoimpl.X.CompressGZIP(file_Product_proto_rawDescData)
	})
	return file_Product_proto_rawDescData
}

var file_Product_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Product_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_Product_proto_goTypes = []interface{}{
	(ProductAreas)(0),       // 0: services.ProductAreas
	(*ProductRequest)(nil),  // 1: services.ProductRequest
	(*ProductInfo)(nil),     // 2: services.ProductInfo
	(*ProductResponse)(nil), // 3: services.ProductResponse
}
var file_Product_proto_depIdxs = []int32{
	0, // 0: services.ProductRequest.pro_area:type_name -> services.ProductAreas
	0, // 1: services.ProductInfo.pro_area:type_name -> services.ProductAreas
	2, // 2: services.ProductResponse.data:type_name -> services.ProductInfo
	1, // 3: services.ProductService.GetProductInfo:input_type -> services.ProductRequest
	1, // 4: services.ProductService.GetProductList:input_type -> services.ProductRequest
	3, // 5: services.ProductService.GetProductInfo:output_type -> services.ProductResponse
	3, // 6: services.ProductService.GetProductList:output_type -> services.ProductResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_Product_proto_init() }
func file_Product_proto_init() {
	if File_Product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductRequest); i {
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
		file_Product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductInfo); i {
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
		file_Product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductResponse); i {
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
			RawDescriptor: file_Product_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Product_proto_goTypes,
		DependencyIndexes: file_Product_proto_depIdxs,
		EnumInfos:         file_Product_proto_enumTypes,
		MessageInfos:      file_Product_proto_msgTypes,
	}.Build()
	File_Product_proto = out.File
	file_Product_proto_rawDesc = nil
	file_Product_proto_goTypes = nil
	file_Product_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductServiceClient interface {
	GetProductInfo(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error)
	GetProductList(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) GetProductInfo(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/services.ProductService/GetProductInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProductList(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/services.ProductService/GetProductList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
type ProductServiceServer interface {
	GetProductInfo(context.Context, *ProductRequest) (*ProductResponse, error)
	GetProductList(context.Context, *ProductRequest) (*ProductResponse, error)
}

// UnimplementedProductServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (*UnimplementedProductServiceServer) GetProductInfo(context.Context, *ProductRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductInfo not implemented")
}
func (*UnimplementedProductServiceServer) GetProductList(context.Context, *ProductRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductList not implemented")
}

func RegisterProductServiceServer(s *grpc.Server, srv ProductServiceServer) {
	s.RegisterService(&_ProductService_serviceDesc, srv)
}

func _ProductService_GetProductInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProductInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ProductService/GetProductInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProductInfo(ctx, req.(*ProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProductList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProductList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ProductService/GetProductList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProductList(ctx, req.(*ProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProductInfo",
			Handler:    _ProductService_GetProductInfo_Handler,
		},
		{
			MethodName: "GetProductList",
			Handler:    _ProductService_GetProductList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Product.proto",
}
