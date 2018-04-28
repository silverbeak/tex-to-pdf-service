// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tex-to-pdf-service.proto

/*
Package com_sootsafe_latex is a generated protocol buffer package.

It is generated from these files:
	tex-to-pdf-service.proto

It has these top-level messages:
	TexMessage
	ErrorMessage
	PdfResponse
*/
package com_sootsafe_latex

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TexMessage struct {
	Tex []byte `protobuf:"bytes,1,opt,name=tex,proto3" json:"tex,omitempty"`
}

func (m *TexMessage) Reset()                    { *m = TexMessage{} }
func (m *TexMessage) String() string            { return proto.CompactTextString(m) }
func (*TexMessage) ProtoMessage()               {}
func (*TexMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TexMessage) GetTex() []byte {
	if m != nil {
		return m.Tex
	}
	return nil
}

type ErrorMessage struct {
	ErrorCode    int32  `protobuf:"varint,1,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage" json:"error_message,omitempty"`
}

func (m *ErrorMessage) Reset()                    { *m = ErrorMessage{} }
func (m *ErrorMessage) String() string            { return proto.CompactTextString(m) }
func (*ErrorMessage) ProtoMessage()               {}
func (*ErrorMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ErrorMessage) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *ErrorMessage) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

type PdfResponse struct {
	Error *ErrorMessage `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Pdf   []byte        `protobuf:"bytes,2,opt,name=pdf,proto3" json:"pdf,omitempty"`
}

func (m *PdfResponse) Reset()                    { *m = PdfResponse{} }
func (m *PdfResponse) String() string            { return proto.CompactTextString(m) }
func (*PdfResponse) ProtoMessage()               {}
func (*PdfResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PdfResponse) GetError() *ErrorMessage {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *PdfResponse) GetPdf() []byte {
	if m != nil {
		return m.Pdf
	}
	return nil
}

func init() {
	proto.RegisterType((*TexMessage)(nil), "com.sootsafe.latex.TexMessage")
	proto.RegisterType((*ErrorMessage)(nil), "com.sootsafe.latex.ErrorMessage")
	proto.RegisterType((*PdfResponse)(nil), "com.sootsafe.latex.PdfResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TexToPdf service

type TexToPdfClient interface {
	Convert(ctx context.Context, in *TexMessage, opts ...grpc.CallOption) (*PdfResponse, error)
}

type texToPdfClient struct {
	cc *grpc.ClientConn
}

func NewTexToPdfClient(cc *grpc.ClientConn) TexToPdfClient {
	return &texToPdfClient{cc}
}

func (c *texToPdfClient) Convert(ctx context.Context, in *TexMessage, opts ...grpc.CallOption) (*PdfResponse, error) {
	out := new(PdfResponse)
	err := grpc.Invoke(ctx, "/com.sootsafe.latex.TexToPdf/Convert", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TexToPdf service

type TexToPdfServer interface {
	Convert(context.Context, *TexMessage) (*PdfResponse, error)
}

func RegisterTexToPdfServer(s *grpc.Server, srv TexToPdfServer) {
	s.RegisterService(&_TexToPdf_serviceDesc, srv)
}

func _TexToPdf_Convert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TexMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TexToPdfServer).Convert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.sootsafe.latex.TexToPdf/Convert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TexToPdfServer).Convert(ctx, req.(*TexMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _TexToPdf_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.sootsafe.latex.TexToPdf",
	HandlerType: (*TexToPdfServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Convert",
			Handler:    _TexToPdf_Convert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tex-to-pdf-service.proto",
}

func init() { proto.RegisterFile("tex-to-pdf-service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0xad, 0x32, 0x75, 0x6f, 0x15, 0x24, 0x78, 0x28, 0x82, 0xb3, 0xd4, 0xcb, 0x2e, 0xcd,
	0x61, 0x82, 0x7f, 0x80, 0xc3, 0x9b, 0xc2, 0x08, 0x05, 0xbd, 0x49, 0x6d, 0xbe, 0x88, 0xe0, 0xf6,
	0x4a, 0x12, 0x46, 0xfe, 0x7c, 0x49, 0x8a, 0x58, 0xb0, 0xb7, 0x97, 0x7c, 0x3f, 0xbe, 0xfc, 0xf2,
	0xa8, 0xf0, 0x08, 0xb5, 0xe7, 0xba, 0xd7, 0xa6, 0x76, 0xb0, 0x87, 0xaf, 0x0e, 0xb2, 0xb7, 0xec,
	0x59, 0x88, 0x8e, 0x77, 0xd2, 0x31, 0x7b, 0xd7, 0x1a, 0xc8, 0xef, 0xd6, 0x23, 0x54, 0x4b, 0xa2,
	0x06, 0xe1, 0x05, 0xce, 0xb5, 0x9f, 0x10, 0x97, 0x74, 0xe2, 0x11, 0x8a, 0xac, 0xcc, 0x56, 0xb9,
	0x8a, 0x63, 0xa5, 0x28, 0x7f, 0xb2, 0x96, 0xed, 0x2f, 0x71, 0x43, 0x84, 0x78, 0x7e, 0xef, 0x58,
	0x23, 0x81, 0x33, 0x35, 0x4f, 0x37, 0x1b, 0xd6, 0x10, 0x77, 0x74, 0x31, 0xc4, 0xbb, 0x81, 0x2f,
	0x8e, 0xcb, 0x6c, 0x35, 0x57, 0x39, 0x46, 0x1d, 0xd5, 0x2b, 0x2d, 0xb6, 0xda, 0x28, 0xb8, 0x9e,
	0xf7, 0x0e, 0xe2, 0x81, 0x66, 0x29, 0x4e, 0x6d, 0x8b, 0x75, 0x29, 0xff, 0x6b, 0xca, 0xb1, 0x83,
	0x1a, 0xf0, 0x28, 0xdb, 0x6b, 0x93, 0x5e, 0xc8, 0x55, 0x1c, 0xd7, 0x6f, 0x74, 0xde, 0x20, 0x34,
	0xbc, 0xd5, 0x46, 0x3c, 0xd3, 0xd9, 0x86, 0xf7, 0x07, 0x58, 0x2f, 0x96, 0x53, 0x8d, 0x7f, 0xbf,
	0xbe, 0xbe, 0x9d, 0xca, 0x47, 0x86, 0xd5, 0xd1, 0xe3, 0x15, 0x4d, 0x2c, 0xef, 0xe3, 0x34, 0xed,
	0xf5, 0xfe, 0x27, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x1a, 0x5f, 0x74, 0x73, 0x01, 0x00, 0x00,
}
