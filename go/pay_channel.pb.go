// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: pay_channel.proto

package pay

import (
	context "context"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

type HTTPRequest_HttpMethod int32

const (
	HTTPRequest_GET  HTTPRequest_HttpMethod = 0
	HTTPRequest_POST HTTPRequest_HttpMethod = 1
	HTTPRequest_PUT  HTTPRequest_HttpMethod = 2
)

// Enum value maps for HTTPRequest_HttpMethod.
var (
	HTTPRequest_HttpMethod_name = map[int32]string{
		0: "GET",
		1: "POST",
		2: "PUT",
	}
	HTTPRequest_HttpMethod_value = map[string]int32{
		"GET":  0,
		"POST": 1,
		"PUT":  2,
	}
)

func (x HTTPRequest_HttpMethod) Enum() *HTTPRequest_HttpMethod {
	p := new(HTTPRequest_HttpMethod)
	*p = x
	return p
}

func (x HTTPRequest_HttpMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HTTPRequest_HttpMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_pay_channel_proto_enumTypes[0].Descriptor()
}

func (HTTPRequest_HttpMethod) Type() protoreflect.EnumType {
	return &file_pay_channel_proto_enumTypes[0]
}

func (x HTTPRequest_HttpMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HTTPRequest_HttpMethod.Descriptor instead.
func (HTTPRequest_HttpMethod) EnumDescriptor() ([]byte, []int) {
	return file_pay_channel_proto_rawDescGZIP(), []int{4, 0}
}

type ChannelPayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GatewayOrderId string            `protobuf:"bytes,1,opt,name=gateway_order_id,json=gatewayOrderId,proto3" json:"gateway_order_id,omitempty"`
	ChannelAccount string            `protobuf:"bytes,2,opt,name=channel_account,json=channelAccount,proto3" json:"channel_account,omitempty"`
	PayAmount      uint64            `protobuf:"varint,3,opt,name=pay_amount,json=payAmount,proto3" json:"pay_amount,omitempty"`
	Product        *Product          `protobuf:"bytes,4,opt,name=product,proto3" json:"product,omitempty"`
	NotifyUrl      string            `protobuf:"bytes,5,opt,name=notify_url,json=notifyUrl,proto3" json:"notify_url,omitempty"`
	ReturnUrl      string            `protobuf:"bytes,6,opt,name=return_url,json=returnUrl,proto3" json:"return_url,omitempty"`
	UserIp         string            `protobuf:"bytes,7,opt,name=user_ip,json=userIp,proto3" json:"user_ip,omitempty"`
	Method         Method            `protobuf:"varint,98,opt,name=method,proto3,enum=pay.Method" json:"method,omitempty"`
	Meta           map[string]string `protobuf:"bytes,99,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ChannelPayRequest) Reset() {
	*x = ChannelPayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pay_channel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelPayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelPayRequest) ProtoMessage() {}

func (x *ChannelPayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pay_channel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelPayRequest.ProtoReflect.Descriptor instead.
func (*ChannelPayRequest) Descriptor() ([]byte, []int) {
	return file_pay_channel_proto_rawDescGZIP(), []int{0}
}

func (x *ChannelPayRequest) GetGatewayOrderId() string {
	if x != nil {
		return x.GatewayOrderId
	}
	return ""
}

func (x *ChannelPayRequest) GetChannelAccount() string {
	if x != nil {
		return x.ChannelAccount
	}
	return ""
}

func (x *ChannelPayRequest) GetPayAmount() uint64 {
	if x != nil {
		return x.PayAmount
	}
	return 0
}

func (x *ChannelPayRequest) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

func (x *ChannelPayRequest) GetNotifyUrl() string {
	if x != nil {
		return x.NotifyUrl
	}
	return ""
}

func (x *ChannelPayRequest) GetReturnUrl() string {
	if x != nil {
		return x.ReturnUrl
	}
	return ""
}

func (x *ChannelPayRequest) GetUserIp() string {
	if x != nil {
		return x.UserIp
	}
	return ""
}

func (x *ChannelPayRequest) GetMethod() Method {
	if x != nil {
		return x.Method
	}
	return Method_WEB
}

func (x *ChannelPayRequest) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

type ChannelPayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelOrderId string            `protobuf:"bytes,4,opt,name=channel_order_id,json=channelOrderId,proto3" json:"channel_order_id,omitempty"`
	Data           map[string]string `protobuf:"bytes,5,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ChannelPayResponse) Reset() {
	*x = ChannelPayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pay_channel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelPayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelPayResponse) ProtoMessage() {}

func (x *ChannelPayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pay_channel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelPayResponse.ProtoReflect.Descriptor instead.
func (*ChannelPayResponse) Descriptor() ([]byte, []int) {
	return file_pay_channel_proto_rawDescGZIP(), []int{1}
}

func (x *ChannelPayResponse) GetChannelOrderId() string {
	if x != nil {
		return x.ChannelOrderId
	}
	return ""
}

func (x *ChannelPayResponse) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

//来自第三方的请求的Notify
type ChannelNotifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//支付账户（可选）, 允许用户配置多个支付宝账户, 如果为空，则默认选择第一个配置文件
	PaymentAccount string `protobuf:"bytes,1,opt,name=payment_account,json=paymentAccount,proto3" json:"payment_account,omitempty"`
	//支付类型：  pay（付款相关通知）,sign_pay（签约）, 根据具体通知配置而定。
	Type PayType `protobuf:"varint,2,opt,name=type,proto3,enum=pay.PayType" json:"type,omitempty"`
	//请求的HTTP详情
	Request *HTTPRequest `protobuf:"bytes,3,opt,name=request,proto3" json:"request,omitempty"`
	//可选，支付方式， 可用于统一网关层
	Method string `protobuf:"bytes,99,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *ChannelNotifyRequest) Reset() {
	*x = ChannelNotifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pay_channel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelNotifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelNotifyRequest) ProtoMessage() {}

func (x *ChannelNotifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pay_channel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelNotifyRequest.ProtoReflect.Descriptor instead.
func (*ChannelNotifyRequest) Descriptor() ([]byte, []int) {
	return file_pay_channel_proto_rawDescGZIP(), []int{2}
}

func (x *ChannelNotifyRequest) GetPaymentAccount() string {
	if x != nil {
		return x.PaymentAccount
	}
	return ""
}

func (x *ChannelNotifyRequest) GetType() PayType {
	if x != nil {
		return x.Type
	}
	return PayType_PAY
}

func (x *ChannelNotifyRequest) GetRequest() *HTTPRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *ChannelNotifyRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

//来自第三方的请求的Notify
type ChannelNotifyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//付款状态
	Status PayStatus `protobuf:"varint,1,opt,name=status,proto3,enum=pay.PayStatus" json:"status,omitempty"`
	//同步返回的数据
	Response *HTTPResponse `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *ChannelNotifyResponse) Reset() {
	*x = ChannelNotifyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pay_channel_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelNotifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelNotifyResponse) ProtoMessage() {}

func (x *ChannelNotifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pay_channel_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelNotifyResponse.ProtoReflect.Descriptor instead.
func (*ChannelNotifyResponse) Descriptor() ([]byte, []int) {
	return file_pay_channel_proto_rawDescGZIP(), []int{3}
}

func (x *ChannelNotifyResponse) GetStatus() PayStatus {
	if x != nil {
		return x.Status
	}
	return PayStatus_SUCCESS
}

func (x *ChannelNotifyResponse) GetResponse() *HTTPResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

//来自第三方的请求body
type HTTPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//请求方法： GET or POST
	Method HTTPRequest_HttpMethod `protobuf:"varint,1,opt,name=method,proto3,enum=pay.HTTPRequest_HttpMethod" json:"method,omitempty"`
	//其他附加信息，可选返回商户返回的meta信息
	Header map[string]string `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	//请求的第三方的url，包含 /some?a=b&c=d, 也可以简化为：?a=b&c=d,不填写具体URL PATH
	Url string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	//请求二进制，可能是：xml或json
	Body []byte `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *HTTPRequest) Reset() {
	*x = HTTPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pay_channel_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPRequest) ProtoMessage() {}

func (x *HTTPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pay_channel_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPRequest.ProtoReflect.Descriptor instead.
func (*HTTPRequest) Descriptor() ([]byte, []int) {
	return file_pay_channel_proto_rawDescGZIP(), []int{4}
}

func (x *HTTPRequest) GetMethod() HTTPRequest_HttpMethod {
	if x != nil {
		return x.Method
	}
	return HTTPRequest_GET
}

func (x *HTTPRequest) GetHeader() map[string]string {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *HTTPRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *HTTPRequest) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

//向第三方响应的Response
type HTTPResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//其他附加信息，可选返回商户返回的meta信息
	Header map[string]string `protobuf:"bytes,1,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	//请求二进制，可能是：xml或json
	Body []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	//请求响应状态码，默认200
	Status int32 `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *HTTPResponse) Reset() {
	*x = HTTPResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pay_channel_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPResponse) ProtoMessage() {}

func (x *HTTPResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pay_channel_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPResponse.ProtoReflect.Descriptor instead.
func (*HTTPResponse) Descriptor() ([]byte, []int) {
	return file_pay_channel_proto_rawDescGZIP(), []int{5}
}

func (x *HTTPResponse) GetHeader() map[string]string {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *HTTPResponse) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *HTTPResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_pay_channel_proto protoreflect.FileDescriptor

var file_pay_channel_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x61, 0x79, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x70, 0x61, 0x79, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x03, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x50, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x61, 0x79, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x09, 0x70, 0x61, 0x79, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a,
	0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x70, 0x61, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x55, 0x72, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x70, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x70, 0x12, 0x23, 0x0a, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x62, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70,
	0x61, 0x79, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x34, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x63, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x70, 0x61, 0x79, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x61, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0xae, 0x01, 0x0a, 0x12, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x61, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x35, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x70, 0x61, 0x79, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x61, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0xa5, 0x01, 0x0a, 0x14, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0c, 0x2e, 0x70, 0x61, 0x79, 0x2e, 0x50, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x61, 0x79, 0x2e, 0x48, 0x54, 0x54,
	0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x63, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x6e, 0x0a, 0x15, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x70, 0x61, 0x79, 0x2e, 0x50, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2d, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70,
	0x61, 0x79, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x83, 0x02, 0x0a, 0x0b, 0x48, 0x54,
	0x54, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x70, 0x61, 0x79, 0x2e,
	0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x74, 0x74, 0x70,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x34,
	0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x70, 0x61, 0x79, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x1a, 0x39, 0x0a, 0x0b, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x28, 0x0a, 0x0a, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x45, 0x54, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x50, 0x4f, 0x53, 0x54, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x55, 0x54, 0x10, 0x02, 0x22,
	0xac, 0x01, 0x0a, 0x0c, 0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x35, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x70, 0x61, 0x79, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xd2,
	0x01, 0x0a, 0x0a, 0x50, 0x61, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x53, 0x0a,
	0x03, 0x50, 0x61, 0x79, 0x12, 0x16, 0x2e, 0x70, 0x61, 0x79, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x50, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70,
	0x61, 0x79, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x61, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x7d, 0x3a,
	0x01, 0x2a, 0x12, 0x6f, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x12, 0x19, 0x2e, 0x70, 0x61, 0x79, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x70, 0x61, 0x79, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x21, 0x1a, 0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2f, 0x7b,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x7d,
	0x3a, 0x01, 0x2a, 0x42, 0x1b, 0x0a, 0x14, 0x70, 0x75, 0x62, 0x2e, 0x70, 0x6a, 0x6f, 0x63, 0x2e,
	0x70, 0x61, 0x79, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5a, 0x03, 0x70, 0x61, 0x79,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pay_channel_proto_rawDescOnce sync.Once
	file_pay_channel_proto_rawDescData = file_pay_channel_proto_rawDesc
)

func file_pay_channel_proto_rawDescGZIP() []byte {
	file_pay_channel_proto_rawDescOnce.Do(func() {
		file_pay_channel_proto_rawDescData = protoimpl.X.CompressGZIP(file_pay_channel_proto_rawDescData)
	})
	return file_pay_channel_proto_rawDescData
}

var file_pay_channel_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pay_channel_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pay_channel_proto_goTypes = []interface{}{
	(HTTPRequest_HttpMethod)(0),   // 0: pay.HTTPRequest.HttpMethod
	(*ChannelPayRequest)(nil),     // 1: pay.ChannelPayRequest
	(*ChannelPayResponse)(nil),    // 2: pay.ChannelPayResponse
	(*ChannelNotifyRequest)(nil),  // 3: pay.ChannelNotifyRequest
	(*ChannelNotifyResponse)(nil), // 4: pay.ChannelNotifyResponse
	(*HTTPRequest)(nil),           // 5: pay.HTTPRequest
	(*HTTPResponse)(nil),          // 6: pay.HTTPResponse
	nil,                           // 7: pay.ChannelPayRequest.MetaEntry
	nil,                           // 8: pay.ChannelPayResponse.DataEntry
	nil,                           // 9: pay.HTTPRequest.HeaderEntry
	nil,                           // 10: pay.HTTPResponse.HeaderEntry
	(*Product)(nil),               // 11: pay.Product
	(Method)(0),                   // 12: pay.Method
	(PayType)(0),                  // 13: pay.PayType
	(PayStatus)(0),                // 14: pay.PayStatus
}
var file_pay_channel_proto_depIdxs = []int32{
	11, // 0: pay.ChannelPayRequest.product:type_name -> pay.Product
	12, // 1: pay.ChannelPayRequest.method:type_name -> pay.Method
	7,  // 2: pay.ChannelPayRequest.meta:type_name -> pay.ChannelPayRequest.MetaEntry
	8,  // 3: pay.ChannelPayResponse.data:type_name -> pay.ChannelPayResponse.DataEntry
	13, // 4: pay.ChannelNotifyRequest.type:type_name -> pay.PayType
	5,  // 5: pay.ChannelNotifyRequest.request:type_name -> pay.HTTPRequest
	14, // 6: pay.ChannelNotifyResponse.status:type_name -> pay.PayStatus
	6,  // 7: pay.ChannelNotifyResponse.response:type_name -> pay.HTTPResponse
	0,  // 8: pay.HTTPRequest.method:type_name -> pay.HTTPRequest.HttpMethod
	9,  // 9: pay.HTTPRequest.header:type_name -> pay.HTTPRequest.HeaderEntry
	10, // 10: pay.HTTPResponse.header:type_name -> pay.HTTPResponse.HeaderEntry
	1,  // 11: pay.PayChannel.Pay:input_type -> pay.ChannelPayRequest
	3,  // 12: pay.PayChannel.ChannelNotify:input_type -> pay.ChannelNotifyRequest
	2,  // 13: pay.PayChannel.Pay:output_type -> pay.ChannelPayResponse
	4,  // 14: pay.PayChannel.ChannelNotify:output_type -> pay.ChannelNotifyResponse
	13, // [13:15] is the sub-list for method output_type
	11, // [11:13] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_pay_channel_proto_init() }
func file_pay_channel_proto_init() {
	if File_pay_channel_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pay_channel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelPayRequest); i {
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
		file_pay_channel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelPayResponse); i {
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
		file_pay_channel_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelNotifyRequest); i {
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
		file_pay_channel_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelNotifyResponse); i {
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
		file_pay_channel_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPRequest); i {
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
		file_pay_channel_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPResponse); i {
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
			RawDescriptor: file_pay_channel_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pay_channel_proto_goTypes,
		DependencyIndexes: file_pay_channel_proto_depIdxs,
		EnumInfos:         file_pay_channel_proto_enumTypes,
		MessageInfos:      file_pay_channel_proto_msgTypes,
	}.Build()
	File_pay_channel_proto = out.File
	file_pay_channel_proto_rawDesc = nil
	file_pay_channel_proto_goTypes = nil
	file_pay_channel_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PayChannelClient is the client API for PayChannel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PayChannelClient interface {
	Pay(ctx context.Context, in *ChannelPayRequest, opts ...grpc.CallOption) (*ChannelPayResponse, error)
	ChannelNotify(ctx context.Context, in *ChannelNotifyRequest, opts ...grpc.CallOption) (*ChannelNotifyResponse, error)
}

type payChannelClient struct {
	cc grpc.ClientConnInterface
}

func NewPayChannelClient(cc grpc.ClientConnInterface) PayChannelClient {
	return &payChannelClient{cc}
}

func (c *payChannelClient) Pay(ctx context.Context, in *ChannelPayRequest, opts ...grpc.CallOption) (*ChannelPayResponse, error) {
	out := new(ChannelPayResponse)
	err := c.cc.Invoke(ctx, "/pay.PayChannel/Pay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payChannelClient) ChannelNotify(ctx context.Context, in *ChannelNotifyRequest, opts ...grpc.CallOption) (*ChannelNotifyResponse, error) {
	out := new(ChannelNotifyResponse)
	err := c.cc.Invoke(ctx, "/pay.PayChannel/ChannelNotify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PayChannelServer is the server API for PayChannel service.
type PayChannelServer interface {
	Pay(context.Context, *ChannelPayRequest) (*ChannelPayResponse, error)
	ChannelNotify(context.Context, *ChannelNotifyRequest) (*ChannelNotifyResponse, error)
}

// UnimplementedPayChannelServer can be embedded to have forward compatible implementations.
type UnimplementedPayChannelServer struct {
}

func (*UnimplementedPayChannelServer) Pay(context.Context, *ChannelPayRequest) (*ChannelPayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pay not implemented")
}
func (*UnimplementedPayChannelServer) ChannelNotify(context.Context, *ChannelNotifyRequest) (*ChannelNotifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChannelNotify not implemented")
}

func RegisterPayChannelServer(s *grpc.Server, srv PayChannelServer) {
	s.RegisterService(&_PayChannel_serviceDesc, srv)
}

func _PayChannel_Pay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChannelPayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayChannelServer).Pay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pay.PayChannel/Pay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayChannelServer).Pay(ctx, req.(*ChannelPayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayChannel_ChannelNotify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChannelNotifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayChannelServer).ChannelNotify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pay.PayChannel/ChannelNotify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayChannelServer).ChannelNotify(ctx, req.(*ChannelNotifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PayChannel_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pay.PayChannel",
	HandlerType: (*PayChannelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pay",
			Handler:    _PayChannel_Pay_Handler,
		},
		{
			MethodName: "ChannelNotify",
			Handler:    _PayChannel_ChannelNotify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pay_channel.proto",
}
