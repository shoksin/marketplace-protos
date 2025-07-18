// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: proto/pbproduct/product.proto

package pbproduct

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price         float64                `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
	Stock         int64                  `protobuf:"varint,3,opt,name=stock,proto3" json:"stock,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateProductRequest) Reset() {
	*x = CreateProductRequest{}
	mi := &file_proto_pbproduct_product_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductRequest) ProtoMessage() {}

func (x *CreateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pbproduct_product_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductRequest.ProtoReflect.Descriptor instead.
func (*CreateProductRequest) Descriptor() ([]byte, []int) {
	return file_proto_pbproduct_product_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProductRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateProductRequest) GetStock() int64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

type CreateProductResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        int64                  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error         string                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Id            string                 `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateProductResponse) Reset() {
	*x = CreateProductResponse{}
	mi := &file_proto_pbproduct_product_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductResponse) ProtoMessage() {}

func (x *CreateProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pbproduct_product_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductResponse.ProtoReflect.Descriptor instead.
func (*CreateProductResponse) Descriptor() ([]byte, []int) {
	return file_proto_pbproduct_product_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProductResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateProductResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CreateProductResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FindOneData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price         float64                `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	Stock         int64                  `protobuf:"varint,4,opt,name=stock,proto3" json:"stock,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindOneData) Reset() {
	*x = FindOneData{}
	mi := &file_proto_pbproduct_product_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindOneData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneData) ProtoMessage() {}

func (x *FindOneData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pbproduct_product_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneData.ProtoReflect.Descriptor instead.
func (*FindOneData) Descriptor() ([]byte, []int) {
	return file_proto_pbproduct_product_proto_rawDescGZIP(), []int{2}
}

func (x *FindOneData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FindOneData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FindOneData) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *FindOneData) GetStock() int64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

type FindOneRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindOneRequest) Reset() {
	*x = FindOneRequest{}
	mi := &file_proto_pbproduct_product_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneRequest) ProtoMessage() {}

func (x *FindOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pbproduct_product_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneRequest.ProtoReflect.Descriptor instead.
func (*FindOneRequest) Descriptor() ([]byte, []int) {
	return file_proto_pbproduct_product_proto_rawDescGZIP(), []int{3}
}

func (x *FindOneRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FindOneResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        int64                  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error         string                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data          *FindOneData           `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindOneResponse) Reset() {
	*x = FindOneResponse{}
	mi := &file_proto_pbproduct_product_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindOneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneResponse) ProtoMessage() {}

func (x *FindOneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pbproduct_product_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneResponse.ProtoReflect.Descriptor instead.
func (*FindOneResponse) Descriptor() ([]byte, []int) {
	return file_proto_pbproduct_product_proto_rawDescGZIP(), []int{4}
}

func (x *FindOneResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *FindOneResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *FindOneResponse) GetData() *FindOneData {
	if x != nil {
		return x.Data
	}
	return nil
}

type FindAllRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindAllRequest) Reset() {
	*x = FindAllRequest{}
	mi := &file_proto_pbproduct_product_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllRequest) ProtoMessage() {}

func (x *FindAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pbproduct_product_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllRequest.ProtoReflect.Descriptor instead.
func (*FindAllRequest) Descriptor() ([]byte, []int) {
	return file_proto_pbproduct_product_proto_rawDescGZIP(), []int{5}
}

type FindAllResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        int64                  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error         string                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Products      []*FindOneData         `protobuf:"bytes,3,rep,name=products,proto3" json:"products,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindAllResponse) Reset() {
	*x = FindAllResponse{}
	mi := &file_proto_pbproduct_product_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllResponse) ProtoMessage() {}

func (x *FindAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pbproduct_product_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllResponse.ProtoReflect.Descriptor instead.
func (*FindAllResponse) Descriptor() ([]byte, []int) {
	return file_proto_pbproduct_product_proto_rawDescGZIP(), []int{6}
}

func (x *FindAllResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *FindAllResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *FindAllResponse) GetProducts() []*FindOneData {
	if x != nil {
		return x.Products
	}
	return nil
}

type DecreaseStockRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderID       string                 `protobuf:"bytes,2,opt,name=orderID,proto3" json:"orderID,omitempty"`
	Quantity      int64                  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DecreaseStockRequest) Reset() {
	*x = DecreaseStockRequest{}
	mi := &file_proto_pbproduct_product_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DecreaseStockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecreaseStockRequest) ProtoMessage() {}

func (x *DecreaseStockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pbproduct_product_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecreaseStockRequest.ProtoReflect.Descriptor instead.
func (*DecreaseStockRequest) Descriptor() ([]byte, []int) {
	return file_proto_pbproduct_product_proto_rawDescGZIP(), []int{7}
}

func (x *DecreaseStockRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DecreaseStockRequest) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

func (x *DecreaseStockRequest) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type DecreaseStockResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        int64                  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error         string                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DecreaseStockResponse) Reset() {
	*x = DecreaseStockResponse{}
	mi := &file_proto_pbproduct_product_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DecreaseStockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecreaseStockResponse) ProtoMessage() {}

func (x *DecreaseStockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pbproduct_product_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecreaseStockResponse.ProtoReflect.Descriptor instead.
func (*DecreaseStockResponse) Descriptor() ([]byte, []int) {
	return file_proto_pbproduct_product_proto_rawDescGZIP(), []int{8}
}

func (x *DecreaseStockResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *DecreaseStockResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_proto_pbproduct_product_proto protoreflect.FileDescriptor

const file_proto_pbproduct_product_proto_rawDesc = "" +
	"\n" +
	"\x1dproto/pbproduct/product.proto\x12\tpbproduct\"V\n" +
	"\x14CreateProductRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x14\n" +
	"\x05price\x18\x02 \x01(\x01R\x05price\x12\x14\n" +
	"\x05stock\x18\x03 \x01(\x03R\x05stock\"U\n" +
	"\x15CreateProductResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\x03R\x06status\x12\x14\n" +
	"\x05error\x18\x02 \x01(\tR\x05error\x12\x0e\n" +
	"\x02id\x18\x03 \x01(\tR\x02id\"]\n" +
	"\vFindOneData\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05price\x18\x03 \x01(\x01R\x05price\x12\x14\n" +
	"\x05stock\x18\x04 \x01(\x03R\x05stock\" \n" +
	"\x0eFindOneRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"k\n" +
	"\x0fFindOneResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\x03R\x06status\x12\x14\n" +
	"\x05error\x18\x02 \x01(\tR\x05error\x12*\n" +
	"\x04data\x18\x03 \x01(\v2\x16.pbproduct.FindOneDataR\x04data\"\x10\n" +
	"\x0eFindAllRequest\"s\n" +
	"\x0fFindAllResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\x03R\x06status\x12\x14\n" +
	"\x05error\x18\x02 \x01(\tR\x05error\x122\n" +
	"\bproducts\x18\x03 \x03(\v2\x16.pbproduct.FindOneDataR\bproducts\"\\\n" +
	"\x14DecreaseStockRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x18\n" +
	"\aorderID\x18\x02 \x01(\tR\aorderID\x12\x1a\n" +
	"\bquantity\x18\x03 \x01(\x03R\bquantity\"E\n" +
	"\x15DecreaseStockResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\x03R\x06status\x12\x14\n" +
	"\x05error\x18\x02 \x01(\tR\x05error2\xbc\x02\n" +
	"\x0eProductService\x12R\n" +
	"\rCreateProduct\x12\x1f.pbproduct.CreateProductRequest\x1a .pbproduct.CreateProductResponse\x12@\n" +
	"\aFindOne\x12\x19.pbproduct.FindOneRequest\x1a\x1a.pbproduct.FindOneResponse\x12@\n" +
	"\aFindAll\x12\x19.pbproduct.FindAllRequest\x1a\x1a.pbproduct.FindAllResponse\x12R\n" +
	"\rDecreaseStock\x12\x1f.pbproduct.DecreaseStockRequest\x1a .pbproduct.DecreaseStockResponseBAZ?github.com/shoksin/marketplace-protos/proto/pbproduct;pbproductb\x06proto3"

var (
	file_proto_pbproduct_product_proto_rawDescOnce sync.Once
	file_proto_pbproduct_product_proto_rawDescData []byte
)

func file_proto_pbproduct_product_proto_rawDescGZIP() []byte {
	file_proto_pbproduct_product_proto_rawDescOnce.Do(func() {
		file_proto_pbproduct_product_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_pbproduct_product_proto_rawDesc), len(file_proto_pbproduct_product_proto_rawDesc)))
	})
	return file_proto_pbproduct_product_proto_rawDescData
}

var file_proto_pbproduct_product_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_pbproduct_product_proto_goTypes = []any{
	(*CreateProductRequest)(nil),  // 0: pbproduct.CreateProductRequest
	(*CreateProductResponse)(nil), // 1: pbproduct.CreateProductResponse
	(*FindOneData)(nil),           // 2: pbproduct.FindOneData
	(*FindOneRequest)(nil),        // 3: pbproduct.FindOneRequest
	(*FindOneResponse)(nil),       // 4: pbproduct.FindOneResponse
	(*FindAllRequest)(nil),        // 5: pbproduct.FindAllRequest
	(*FindAllResponse)(nil),       // 6: pbproduct.FindAllResponse
	(*DecreaseStockRequest)(nil),  // 7: pbproduct.DecreaseStockRequest
	(*DecreaseStockResponse)(nil), // 8: pbproduct.DecreaseStockResponse
}
var file_proto_pbproduct_product_proto_depIdxs = []int32{
	2, // 0: pbproduct.FindOneResponse.data:type_name -> pbproduct.FindOneData
	2, // 1: pbproduct.FindAllResponse.products:type_name -> pbproduct.FindOneData
	0, // 2: pbproduct.ProductService.CreateProduct:input_type -> pbproduct.CreateProductRequest
	3, // 3: pbproduct.ProductService.FindOne:input_type -> pbproduct.FindOneRequest
	5, // 4: pbproduct.ProductService.FindAll:input_type -> pbproduct.FindAllRequest
	7, // 5: pbproduct.ProductService.DecreaseStock:input_type -> pbproduct.DecreaseStockRequest
	1, // 6: pbproduct.ProductService.CreateProduct:output_type -> pbproduct.CreateProductResponse
	4, // 7: pbproduct.ProductService.FindOne:output_type -> pbproduct.FindOneResponse
	6, // 8: pbproduct.ProductService.FindAll:output_type -> pbproduct.FindAllResponse
	8, // 9: pbproduct.ProductService.DecreaseStock:output_type -> pbproduct.DecreaseStockResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_pbproduct_product_proto_init() }
func file_proto_pbproduct_product_proto_init() {
	if File_proto_pbproduct_product_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_pbproduct_product_proto_rawDesc), len(file_proto_pbproduct_product_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_pbproduct_product_proto_goTypes,
		DependencyIndexes: file_proto_pbproduct_product_proto_depIdxs,
		MessageInfos:      file_proto_pbproduct_product_proto_msgTypes,
	}.Build()
	File_proto_pbproduct_product_proto = out.File
	file_proto_pbproduct_product_proto_goTypes = nil
	file_proto_pbproduct_product_proto_depIdxs = nil
}
