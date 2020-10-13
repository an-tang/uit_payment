// Code generated by protoc-gen-go. DO NOT EDIT.
// source: payment.proto

package payment_api

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreatePaymentRequest struct {
	TransactionId        string   `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	Amount               float32  `protobuf:"fixed32,2,opt,name=amount,proto3" json:"amount,omitempty"`
	PaymentMethod        int32    `protobuf:"varint,3,opt,name=payment_method,json=paymentMethod,proto3" json:"payment_method,omitempty"`
	StoreId              string   `protobuf:"bytes,4,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	Prodduct             string   `protobuf:"bytes,5,opt,name=prodduct,proto3" json:"prodduct,omitempty"`
	PartnerKey           string   `protobuf:"bytes,6,opt,name=partner_key,json=partnerKey,proto3" json:"partner_key,omitempty"`
	Token                string   `protobuf:"bytes,7,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePaymentRequest) Reset()         { *m = CreatePaymentRequest{} }
func (m *CreatePaymentRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePaymentRequest) ProtoMessage()    {}
func (*CreatePaymentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6362648dfa63d410, []int{0}
}

func (m *CreatePaymentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePaymentRequest.Unmarshal(m, b)
}
func (m *CreatePaymentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePaymentRequest.Marshal(b, m, deterministic)
}
func (m *CreatePaymentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePaymentRequest.Merge(m, src)
}
func (m *CreatePaymentRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePaymentRequest.Size(m)
}
func (m *CreatePaymentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePaymentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePaymentRequest proto.InternalMessageInfo

func (m *CreatePaymentRequest) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

func (m *CreatePaymentRequest) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *CreatePaymentRequest) GetPaymentMethod() int32 {
	if m != nil {
		return m.PaymentMethod
	}
	return 0
}

func (m *CreatePaymentRequest) GetStoreId() string {
	if m != nil {
		return m.StoreId
	}
	return ""
}

func (m *CreatePaymentRequest) GetProdduct() string {
	if m != nil {
		return m.Prodduct
	}
	return ""
}

func (m *CreatePaymentRequest) GetPartnerKey() string {
	if m != nil {
		return m.PartnerKey
	}
	return ""
}

func (m *CreatePaymentRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CreatePaymentResponse struct {
	QrText               string   `protobuf:"bytes,1,opt,name=qr_text,json=qrText,proto3" json:"qr_text,omitempty"`
	TransactionId        string   `protobuf:"bytes,2,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	PaymentMethod        int32    `protobuf:"varint,3,opt,name=payment_method,json=paymentMethod,proto3" json:"payment_method,omitempty"`
	Status               string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	StatusValue          int32    `protobuf:"varint,5,opt,name=status_value,json=statusValue,proto3" json:"status_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePaymentResponse) Reset()         { *m = CreatePaymentResponse{} }
func (m *CreatePaymentResponse) String() string { return proto.CompactTextString(m) }
func (*CreatePaymentResponse) ProtoMessage()    {}
func (*CreatePaymentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6362648dfa63d410, []int{1}
}

func (m *CreatePaymentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePaymentResponse.Unmarshal(m, b)
}
func (m *CreatePaymentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePaymentResponse.Marshal(b, m, deterministic)
}
func (m *CreatePaymentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePaymentResponse.Merge(m, src)
}
func (m *CreatePaymentResponse) XXX_Size() int {
	return xxx_messageInfo_CreatePaymentResponse.Size(m)
}
func (m *CreatePaymentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePaymentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePaymentResponse proto.InternalMessageInfo

func (m *CreatePaymentResponse) GetQrText() string {
	if m != nil {
		return m.QrText
	}
	return ""
}

func (m *CreatePaymentResponse) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

func (m *CreatePaymentResponse) GetPaymentMethod() int32 {
	if m != nil {
		return m.PaymentMethod
	}
	return 0
}

func (m *CreatePaymentResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *CreatePaymentResponse) GetStatusValue() int32 {
	if m != nil {
		return m.StatusValue
	}
	return 0
}

func init() {
	proto.RegisterType((*CreatePaymentRequest)(nil), "payment.api.CreatePaymentRequest")
	proto.RegisterType((*CreatePaymentResponse)(nil), "payment.api.CreatePaymentResponse")
}

func init() { proto.RegisterFile("payment.proto", fileDescriptor_6362648dfa63d410) }

var fileDescriptor_6362648dfa63d410 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4e, 0x3a, 0x31,
	0x10, 0xc6, 0x53, 0xfe, 0xff, 0x5d, 0x70, 0x10, 0x0e, 0x0d, 0x62, 0xe5, 0x22, 0x90, 0x98, 0x70,
	0xe2, 0xa0, 0x8f, 0xe0, 0x89, 0x18, 0x13, 0xb3, 0x1a, 0xae, 0x9b, 0x4a, 0x27, 0x61, 0x83, 0xb4,
	0xa5, 0x9d, 0x25, 0xf0, 0x68, 0xbe, 0x92, 0x4f, 0x61, 0xe8, 0x16, 0x83, 0x66, 0x63, 0xbc, 0xed,
	0xf7, 0xcd, 0x76, 0x3a, 0xbf, 0x6f, 0x0a, 0x1d, 0x2b, 0xf7, 0x6b, 0xd4, 0x34, 0xb5, 0xce, 0x90,
	0xe1, 0xed, 0xa3, 0x94, 0xb6, 0x18, 0x7f, 0x30, 0xe8, 0xdd, 0x3b, 0x94, 0x84, 0x4f, 0x95, 0x9b,
	0xe1, 0xa6, 0x44, 0x4f, 0xfc, 0x06, 0xba, 0xe4, 0xa4, 0xf6, 0x72, 0x41, 0x85, 0xd1, 0x79, 0xa1,
	0x04, 0x1b, 0xb2, 0xc9, 0x59, 0xd6, 0x39, 0x71, 0x67, 0x8a, 0xf7, 0x21, 0x95, 0x6b, 0x53, 0x6a,
	0x12, 0x8d, 0x21, 0x9b, 0x34, 0xb2, 0xa8, 0x0e, 0xc7, 0xe3, 0x35, 0xf9, 0x1a, 0x69, 0x69, 0x94,
	0xf8, 0x37, 0x64, 0x93, 0x24, 0x3b, 0xce, 0xf2, 0x18, 0x4c, 0x7e, 0x05, 0x2d, 0x4f, 0xc6, 0xe1,
	0xa1, 0xff, 0xff, 0xd0, 0xbf, 0x19, 0xf4, 0x4c, 0xf1, 0x01, 0xb4, 0xac, 0x33, 0x4a, 0x95, 0x0b,
	0x12, 0x49, 0x28, 0x7d, 0x69, 0x7e, 0x0d, 0x6d, 0x2b, 0x1d, 0x69, 0x74, 0xf9, 0x0a, 0xf7, 0x22,
	0x0d, 0x65, 0x88, 0xd6, 0x03, 0xee, 0x79, 0x0f, 0x12, 0x32, 0x2b, 0xd4, 0xa2, 0x19, 0x4a, 0x95,
	0x18, 0xbf, 0x33, 0xb8, 0xf8, 0x01, 0xeb, 0xad, 0xd1, 0x1e, 0xf9, 0x25, 0x34, 0x37, 0x2e, 0x27,
	0xdc, 0x51, 0xc4, 0x4c, 0x37, 0xee, 0x05, 0x77, 0x75, 0x31, 0x34, 0xea, 0x62, 0xf8, 0x23, 0x6e,
	0x1f, 0x52, 0x4f, 0x92, 0x4a, 0x1f, 0x61, 0xa3, 0xe2, 0x23, 0x38, 0xaf, 0xbe, 0xf2, 0xad, 0x7c,
	0x2b, 0x31, 0xf0, 0x26, 0x59, 0xbb, 0xf2, 0xe6, 0x07, 0xeb, 0x76, 0x09, 0xdd, 0x38, 0xf4, 0x33,
	0xba, 0x6d, 0xb1, 0x40, 0x3e, 0x87, 0xce, 0x37, 0x18, 0x3e, 0x9a, 0x9e, 0x6c, 0x76, 0x5a, 0xb7,
	0xd5, 0xc1, 0xf8, 0xb7, 0x5f, 0xaa, 0x2c, 0x5e, 0xd3, 0xf0, 0x4c, 0xee, 0x3e, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xf6, 0xf2, 0x9f, 0x58, 0x37, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PaymentServiceClient is the client API for PaymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PaymentServiceClient interface {
	// create payment
	CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error)
}

type paymentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentServiceClient(cc grpc.ClientConnInterface) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error) {
	out := new(CreatePaymentResponse)
	err := c.cc.Invoke(ctx, "/payment.api.PaymentService/CreatePayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServiceServer is the server API for PaymentService service.
type PaymentServiceServer interface {
	// create payment
	CreatePayment(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error)
}

// UnimplementedPaymentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPaymentServiceServer struct {
}

func (*UnimplementedPaymentServiceServer) CreatePayment(ctx context.Context, req *CreatePaymentRequest) (*CreatePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePayment not implemented")
}

func RegisterPaymentServiceServer(s *grpc.Server, srv PaymentServiceServer) {
	s.RegisterService(&_PaymentService_serviceDesc, srv)
}

func _PaymentService_CreatePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).CreatePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.api.PaymentService/CreatePayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).CreatePayment(ctx, req.(*CreatePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PaymentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "payment.api.PaymentService",
	HandlerType: (*PaymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePayment",
			Handler:    _PaymentService_CreatePayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payment.proto",
}
