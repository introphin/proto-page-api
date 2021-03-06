// Code generated by protoc-gen-go. DO NOT EDIT.
// source: page_api.proto

/*
Package page_api is a generated protocol buffer package.

It is generated from these files:
	page_api.proto

It has these top-level messages:
	Template
	ListTemplatesRequest
	ListTemplatesResponse
	Stream
	StreamPageOptions
	GetStreamRequest
	DeleteRowFromTemplateRequest
*/
package page_api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

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

// Templates
type LayoutType int32

const (
	LayoutType_empty    LayoutType = 0
	LayoutType_desktop  LayoutType = 1
	LayoutType_mobile   LayoutType = 2
	LayoutType_adaptive LayoutType = 3
)

var LayoutType_name = map[int32]string{
	0: "empty",
	1: "desktop",
	2: "mobile",
	3: "adaptive",
}
var LayoutType_value = map[string]int32{
	"empty":    0,
	"desktop":  1,
	"mobile":   2,
	"adaptive": 3,
}

func (x LayoutType) String() string {
	return proto.EnumName(LayoutType_name, int32(x))
}
func (LayoutType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Template struct {
	Id           int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	OfferId      int32  `protobuf:"varint,3,opt,name=offer_id,json=offerId" json:"offer_id,omitempty"`
	Type         int32  `protobuf:"varint,4,opt,name=type" json:"type,omitempty"`
	LayoutType   int32  `protobuf:"varint,5,opt,name=layout_type,json=layoutType" json:"layout_type,omitempty"`
	Content      string `protobuf:"bytes,6,opt,name=content" json:"content,omitempty"`
	Uuid         string `protobuf:"bytes,7,opt,name=uuid" json:"uuid,omitempty"`
	IsMyTarget   bool   `protobuf:"varint,8,opt,name=is_my_target,json=isMyTarget" json:"is_my_target,omitempty"`
	TransitedUrl string `protobuf:"bytes,9,opt,name=transited_url,json=transitedUrl" json:"transited_url,omitempty"`
	Status       int32  `protobuf:"varint,10,opt,name=status" json:"status,omitempty"`
	Visibility   int32  `protobuf:"varint,11,opt,name=visibility" json:"visibility,omitempty"`
}

func (m *Template) Reset()                    { *m = Template{} }
func (m *Template) String() string            { return proto.CompactTextString(m) }
func (*Template) ProtoMessage()               {}
func (*Template) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Template) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Template) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Template) GetOfferId() int32 {
	if m != nil {
		return m.OfferId
	}
	return 0
}

func (m *Template) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Template) GetLayoutType() int32 {
	if m != nil {
		return m.LayoutType
	}
	return 0
}

func (m *Template) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Template) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Template) GetIsMyTarget() bool {
	if m != nil {
		return m.IsMyTarget
	}
	return false
}

func (m *Template) GetTransitedUrl() string {
	if m != nil {
		return m.TransitedUrl
	}
	return ""
}

func (m *Template) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Template) GetVisibility() int32 {
	if m != nil {
		return m.Visibility
	}
	return 0
}

type ListTemplatesRequest struct {
	Id          []int32 `protobuf:"varint,1,rep,packed,name=id" json:"id,omitempty"`
	Type        int32   `protobuf:"varint,2,opt,name=type" json:"type,omitempty"`
	Name        string  `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	OfferIds    []int32 `protobuf:"varint,4,rep,packed,name=offer_ids,json=offerIds" json:"offer_ids,omitempty"`
	AffiliateId int32   `protobuf:"varint,5,opt,name=affiliate_id,json=affiliateId" json:"affiliate_id,omitempty"`
	Limit       int32   `protobuf:"varint,6,opt,name=limit" json:"limit,omitempty"`
	Offset      int32   `protobuf:"varint,7,opt,name=offset" json:"offset,omitempty"`
}

func (m *ListTemplatesRequest) Reset()                    { *m = ListTemplatesRequest{} }
func (m *ListTemplatesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListTemplatesRequest) ProtoMessage()               {}
func (*ListTemplatesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListTemplatesRequest) GetId() []int32 {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ListTemplatesRequest) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ListTemplatesRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListTemplatesRequest) GetOfferIds() []int32 {
	if m != nil {
		return m.OfferIds
	}
	return nil
}

func (m *ListTemplatesRequest) GetAffiliateId() int32 {
	if m != nil {
		return m.AffiliateId
	}
	return 0
}

func (m *ListTemplatesRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListTemplatesRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ListTemplatesResponse struct {
	Templates  []*Template `protobuf:"bytes,1,rep,name=templates" json:"templates,omitempty"`
	TotalCount int32       `protobuf:"varint,2,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
}

func (m *ListTemplatesResponse) Reset()                    { *m = ListTemplatesResponse{} }
func (m *ListTemplatesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListTemplatesResponse) ProtoMessage()               {}
func (*ListTemplatesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListTemplatesResponse) GetTemplates() []*Template {
	if m != nil {
		return m.Templates
	}
	return nil
}

func (m *ListTemplatesResponse) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

// Stream
type Stream struct {
	Id                int32              `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Type              string             `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	UserId            int32              `protobuf:"varint,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	OfferId           int32              `protobuf:"varint,4,opt,name=offer_id,json=offerId" json:"offer_id,omitempty"`
	Name              string             `protobuf:"bytes,5,opt,name=name" json:"name,omitempty"`
	Code              string             `protobuf:"bytes,6,opt,name=code" json:"code,omitempty"`
	PrelandingOptions *StreamPageOptions `protobuf:"bytes,7,opt,name=prelanding_options,json=prelandingOptions" json:"prelanding_options,omitempty"`
	LandingOptions    *StreamPageOptions `protobuf:"bytes,8,opt,name=landing_options,json=landingOptions" json:"landing_options,omitempty"`
	ThanksOptions     *StreamPageOptions `protobuf:"bytes,9,opt,name=thanks_options,json=thanksOptions" json:"thanks_options,omitempty"`
	TrafficBackUrl    string             `protobuf:"bytes,10,opt,name=traffic_back_url,json=trafficBackUrl" json:"traffic_back_url,omitempty"`
	IsMyTarget        bool               `protobuf:"varint,11,opt,name=is_my_target,json=isMyTarget" json:"is_my_target,omitempty"`
	Sub1              string             `protobuf:"bytes,12,opt,name=sub1" json:"sub1,omitempty"`
	Sub2              string             `protobuf:"bytes,13,opt,name=sub2" json:"sub2,omitempty"`
	Sub3              string             `protobuf:"bytes,14,opt,name=sub3" json:"sub3,omitempty"`
	Sub4              string             `protobuf:"bytes,15,opt,name=sub4" json:"sub4,omitempty"`
	Sub5              string             `protobuf:"bytes,16,opt,name=sub5" json:"sub5,omitempty"`
	GeoCode           []string           `protobuf:"bytes,17,rep,name=geo_code,json=geoCode" json:"geo_code,omitempty"`
	Currency          string             `protobuf:"bytes,18,opt,name=currency" json:"currency,omitempty"`
}

func (m *Stream) Reset()                    { *m = Stream{} }
func (m *Stream) String() string            { return proto.CompactTextString(m) }
func (*Stream) ProtoMessage()               {}
func (*Stream) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Stream) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Stream) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Stream) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Stream) GetOfferId() int32 {
	if m != nil {
		return m.OfferId
	}
	return 0
}

func (m *Stream) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Stream) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Stream) GetPrelandingOptions() *StreamPageOptions {
	if m != nil {
		return m.PrelandingOptions
	}
	return nil
}

func (m *Stream) GetLandingOptions() *StreamPageOptions {
	if m != nil {
		return m.LandingOptions
	}
	return nil
}

func (m *Stream) GetThanksOptions() *StreamPageOptions {
	if m != nil {
		return m.ThanksOptions
	}
	return nil
}

func (m *Stream) GetTrafficBackUrl() string {
	if m != nil {
		return m.TrafficBackUrl
	}
	return ""
}

func (m *Stream) GetIsMyTarget() bool {
	if m != nil {
		return m.IsMyTarget
	}
	return false
}

func (m *Stream) GetSub1() string {
	if m != nil {
		return m.Sub1
	}
	return ""
}

func (m *Stream) GetSub2() string {
	if m != nil {
		return m.Sub2
	}
	return ""
}

func (m *Stream) GetSub3() string {
	if m != nil {
		return m.Sub3
	}
	return ""
}

func (m *Stream) GetSub4() string {
	if m != nil {
		return m.Sub4
	}
	return ""
}

func (m *Stream) GetSub5() string {
	if m != nil {
		return m.Sub5
	}
	return ""
}

func (m *Stream) GetGeoCode() []string {
	if m != nil {
		return m.GeoCode
	}
	return nil
}

func (m *Stream) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

type StreamPageOptions struct {
	SubDomain string  `protobuf:"bytes,1,opt,name=sub_domain,json=subDomain" json:"sub_domain,omitempty"`
	Domain    string  `protobuf:"bytes,2,opt,name=domain" json:"domain,omitempty"`
	IsParking bool    `protobuf:"varint,3,opt,name=is_parking,json=isParking" json:"is_parking,omitempty"`
	Templates []int32 `protobuf:"varint,4,rep,packed,name=templates" json:"templates,omitempty"`
	Script    string  `protobuf:"bytes,5,opt,name=script" json:"script,omitempty"`
}

func (m *StreamPageOptions) Reset()                    { *m = StreamPageOptions{} }
func (m *StreamPageOptions) String() string            { return proto.CompactTextString(m) }
func (*StreamPageOptions) ProtoMessage()               {}
func (*StreamPageOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StreamPageOptions) GetSubDomain() string {
	if m != nil {
		return m.SubDomain
	}
	return ""
}

func (m *StreamPageOptions) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *StreamPageOptions) GetIsParking() bool {
	if m != nil {
		return m.IsParking
	}
	return false
}

func (m *StreamPageOptions) GetTemplates() []int32 {
	if m != nil {
		return m.Templates
	}
	return nil
}

func (m *StreamPageOptions) GetScript() string {
	if m != nil {
		return m.Script
	}
	return ""
}

type GetStreamRequest struct {
	Code string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
}

func (m *GetStreamRequest) Reset()                    { *m = GetStreamRequest{} }
func (m *GetStreamRequest) String() string            { return proto.CompactTextString(m) }
func (*GetStreamRequest) ProtoMessage()               {}
func (*GetStreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetStreamRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type DeleteRowFromTemplateRequest struct {
	Row string `protobuf:"bytes,1,opt,name=row" json:"row,omitempty"`
}

func (m *DeleteRowFromTemplateRequest) Reset()                    { *m = DeleteRowFromTemplateRequest{} }
func (m *DeleteRowFromTemplateRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRowFromTemplateRequest) ProtoMessage()               {}
func (*DeleteRowFromTemplateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeleteRowFromTemplateRequest) GetRow() string {
	if m != nil {
		return m.Row
	}
	return ""
}

func init() {
	proto.RegisterType((*Template)(nil), "page_api.Template")
	proto.RegisterType((*ListTemplatesRequest)(nil), "page_api.ListTemplatesRequest")
	proto.RegisterType((*ListTemplatesResponse)(nil), "page_api.ListTemplatesResponse")
	proto.RegisterType((*Stream)(nil), "page_api.Stream")
	proto.RegisterType((*StreamPageOptions)(nil), "page_api.StreamPageOptions")
	proto.RegisterType((*GetStreamRequest)(nil), "page_api.GetStreamRequest")
	proto.RegisterType((*DeleteRowFromTemplateRequest)(nil), "page_api.DeleteRowFromTemplateRequest")
	proto.RegisterEnum("page_api.LayoutType", LayoutType_name, LayoutType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PageApi service

type PageApiClient interface {
	ListTemplates(ctx context.Context, in *ListTemplatesRequest, opts ...grpc.CallOption) (*ListTemplatesResponse, error)
	GetStream(ctx context.Context, in *GetStreamRequest, opts ...grpc.CallOption) (*Stream, error)
	DeleteRowFromTemplate(ctx context.Context, in *DeleteRowFromTemplateRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type pageApiClient struct {
	cc *grpc.ClientConn
}

func NewPageApiClient(cc *grpc.ClientConn) PageApiClient {
	return &pageApiClient{cc}
}

func (c *pageApiClient) ListTemplates(ctx context.Context, in *ListTemplatesRequest, opts ...grpc.CallOption) (*ListTemplatesResponse, error) {
	out := new(ListTemplatesResponse)
	err := grpc.Invoke(ctx, "/page_api.PageApi/ListTemplates", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pageApiClient) GetStream(ctx context.Context, in *GetStreamRequest, opts ...grpc.CallOption) (*Stream, error) {
	out := new(Stream)
	err := grpc.Invoke(ctx, "/page_api.PageApi/GetStream", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pageApiClient) DeleteRowFromTemplate(ctx context.Context, in *DeleteRowFromTemplateRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/page_api.PageApi/DeleteRowFromTemplate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PageApi service

type PageApiServer interface {
	ListTemplates(context.Context, *ListTemplatesRequest) (*ListTemplatesResponse, error)
	GetStream(context.Context, *GetStreamRequest) (*Stream, error)
	DeleteRowFromTemplate(context.Context, *DeleteRowFromTemplateRequest) (*google_protobuf.Empty, error)
}

func RegisterPageApiServer(s *grpc.Server, srv PageApiServer) {
	s.RegisterService(&_PageApi_serviceDesc, srv)
}

func _PageApi_ListTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PageApiServer).ListTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/page_api.PageApi/ListTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PageApiServer).ListTemplates(ctx, req.(*ListTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PageApi_GetStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PageApiServer).GetStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/page_api.PageApi/GetStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PageApiServer).GetStream(ctx, req.(*GetStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PageApi_DeleteRowFromTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRowFromTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PageApiServer).DeleteRowFromTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/page_api.PageApi/DeleteRowFromTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PageApiServer).DeleteRowFromTemplate(ctx, req.(*DeleteRowFromTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PageApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "page_api.PageApi",
	HandlerType: (*PageApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTemplates",
			Handler:    _PageApi_ListTemplates_Handler,
		},
		{
			MethodName: "GetStream",
			Handler:    _PageApi_GetStream_Handler,
		},
		{
			MethodName: "DeleteRowFromTemplate",
			Handler:    _PageApi_DeleteRowFromTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "page_api.proto",
}

func init() { proto.RegisterFile("page_api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 849 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xcd, 0x8e, 0xdc, 0x44,
	0x10, 0x5e, 0xcf, 0xaf, 0x5d, 0x33, 0x3b, 0xf1, 0xb6, 0x92, 0xd0, 0xcc, 0x86, 0x64, 0x30, 0x52,
	0x34, 0xe2, 0x30, 0x09, 0xb3, 0xc9, 0x09, 0x09, 0x89, 0x64, 0x01, 0x2d, 0x0a, 0x22, 0x32, 0xcb,
	0x81, 0x93, 0xd5, 0x63, 0xf7, 0x98, 0x66, 0x6c, 0xb7, 0x71, 0xb7, 0x13, 0xcd, 0xd3, 0xf0, 0x00,
	0xbc, 0x02, 0x57, 0x1e, 0x89, 0x3b, 0xea, 0x1f, 0xdb, 0x93, 0xd9, 0x85, 0xbd, 0x55, 0x7d, 0x5d,
	0x55, 0xee, 0xfa, 0xfa, 0xab, 0x32, 0xcc, 0x4a, 0x92, 0xd2, 0x88, 0x94, 0x6c, 0x55, 0x56, 0x5c,
	0x72, 0xe4, 0x36, 0xfe, 0xfc, 0x3c, 0xe5, 0x3c, 0xcd, 0xe8, 0x33, 0x8d, 0x6f, 0xea, 0xed, 0x33,
	0x9a, 0x97, 0x72, 0x6f, 0xc2, 0x82, 0x3f, 0x7b, 0xe0, 0x5e, 0xd3, 0xbc, 0xcc, 0x88, 0xa4, 0x68,
	0x06, 0x3d, 0x96, 0x60, 0x67, 0xe1, 0x2c, 0x87, 0x61, 0x8f, 0x25, 0x08, 0xc1, 0xa0, 0x20, 0x39,
	0xc5, 0xbd, 0x85, 0xb3, 0xf4, 0x42, 0x6d, 0xa3, 0x8f, 0xc1, 0xe5, 0xdb, 0x2d, 0xad, 0x22, 0x96,
	0xe0, 0xbe, 0x8e, 0x1c, 0x6b, 0xff, 0x4a, 0x87, 0xcb, 0x7d, 0x49, 0xf1, 0x40, 0xc3, 0xda, 0x46,
	0x4f, 0x60, 0x92, 0x91, 0x3d, 0xaf, 0x65, 0xa4, 0x8f, 0x86, 0xfa, 0x08, 0x0c, 0x74, 0xad, 0x02,
	0x30, 0x8c, 0x63, 0x5e, 0x48, 0x5a, 0x48, 0x3c, 0xd2, 0x9f, 0x69, 0x5c, 0x55, 0xae, 0xae, 0x59,
	0x82, 0xc7, 0xe6, 0xeb, 0xca, 0x46, 0x0b, 0x98, 0x32, 0x11, 0xe5, 0xfb, 0x48, 0x92, 0x2a, 0xa5,
	0x12, 0xbb, 0x0b, 0x67, 0xe9, 0x86, 0xc0, 0xc4, 0x0f, 0xfb, 0x6b, 0x8d, 0xa0, 0xcf, 0xe0, 0x54,
	0x56, 0xa4, 0x10, 0x4c, 0xd2, 0x24, 0xaa, 0xab, 0x0c, 0x7b, 0x3a, 0x7d, 0xda, 0x82, 0x3f, 0x57,
	0x19, 0x7a, 0x08, 0x23, 0x21, 0x89, 0xac, 0x05, 0x06, 0x7d, 0x21, 0xeb, 0xa1, 0xc7, 0x00, 0xef,
	0x98, 0x60, 0x1b, 0x96, 0x31, 0xb9, 0xc7, 0x13, 0x73, 0xd9, 0x0e, 0x09, 0xfe, 0x72, 0xe0, 0xfe,
	0x1b, 0x26, 0x64, 0xc3, 0x98, 0x08, 0xe9, 0xef, 0x35, 0x15, 0xb2, 0x65, 0xae, 0xdf, 0x31, 0xa7,
	0xfb, 0xed, 0x1d, 0x50, 0xd1, 0xb0, 0xd9, 0x3f, 0x60, 0xf3, 0x1c, 0xbc, 0x86, 0x4d, 0x81, 0x07,
	0x3a, 0xdd, 0xb5, 0x74, 0x0a, 0xf4, 0x29, 0x4c, 0xc9, 0x76, 0xcb, 0x32, 0x46, 0x24, 0x55, 0x74,
	0x1b, 0xf2, 0x26, 0x2d, 0x76, 0x95, 0xa0, 0xfb, 0x30, 0xcc, 0x58, 0xce, 0x0c, 0x77, 0xc3, 0xd0,
	0x38, 0xaa, 0x3d, 0xbe, 0xdd, 0x0a, 0x2a, 0x35, 0x77, 0xc3, 0xd0, 0x7a, 0xc1, 0x6f, 0xf0, 0xe0,
	0xe8, 0xf6, 0xa2, 0xe4, 0x85, 0xa0, 0xe8, 0x39, 0x78, 0xb2, 0x01, 0x75, 0x17, 0x93, 0x35, 0x5a,
	0xb5, 0x82, 0x6a, 0xe2, 0xc3, 0x2e, 0x48, 0xbd, 0xab, 0xe4, 0x92, 0x64, 0x51, 0xcc, 0xeb, 0x42,
	0xda, 0x3e, 0x41, 0x43, 0xaf, 0x15, 0x12, 0xfc, 0x3d, 0x80, 0xd1, 0x4f, 0xb2, 0xa2, 0x24, 0xbf,
	0x4d, 0x56, 0x2d, 0x39, 0x9e, 0x25, 0xe7, 0x23, 0x18, 0xd7, 0xe2, 0x50, 0x55, 0x23, 0xe5, 0x5e,
	0x25, 0x1f, 0xe8, 0x6d, 0x70, 0x43, 0x6f, 0x9a, 0xd0, 0xe1, 0x01, 0xa1, 0x08, 0x06, 0x31, 0x4f,
	0xa8, 0xd5, 0x92, 0xb6, 0xd1, 0xf7, 0x80, 0xca, 0x8a, 0x66, 0xa4, 0x48, 0x58, 0x91, 0x46, 0xbc,
	0x94, 0x8c, 0x17, 0x42, 0x53, 0x33, 0x59, 0x9f, 0x77, 0x6d, 0x9a, 0xdb, 0xbe, 0x25, 0x29, 0xfd,
	0xd1, 0x84, 0x84, 0x67, 0x5d, 0x9a, 0x85, 0xd0, 0x25, 0xdc, 0x3b, 0x2e, 0xe4, 0xde, 0x5d, 0x68,
	0x76, 0x54, 0xe5, 0x15, 0xcc, 0xe4, 0xaf, 0xa4, 0xd8, 0x89, 0xb6, 0x88, 0x77, 0x77, 0x91, 0x53,
	0x93, 0xd2, 0xd4, 0x58, 0x82, 0x2f, 0x2b, 0xa5, 0x85, 0x38, 0xda, 0x90, 0x78, 0xa7, 0xb5, 0x0e,
	0xba, 0xeb, 0x99, 0xc5, 0x5f, 0x91, 0x78, 0xa7, 0xd4, 0x7e, 0x3c, 0x34, 0x93, 0x1b, 0x43, 0x83,
	0x60, 0x20, 0xea, 0xcd, 0x17, 0x78, 0x6a, 0x58, 0x53, 0xb6, 0xc5, 0xd6, 0xf8, 0xb4, 0xc5, 0xd6,
	0x16, 0xbb, 0xc0, 0xb3, 0x16, 0xbb, 0xb0, 0xd8, 0x0b, 0x7c, 0xaf, 0xc5, 0x5e, 0x58, 0xec, 0x25,
	0xf6, 0x5b, 0xec, 0xa5, 0x7a, 0xc8, 0x94, 0xf2, 0x48, 0xbf, 0xce, 0xd9, 0xa2, 0xaf, 0x26, 0x3d,
	0xa5, 0xfc, 0xb5, 0x7a, 0xa0, 0x39, 0xb8, 0x71, 0x5d, 0x55, 0xb4, 0x88, 0xf7, 0x18, 0xe9, 0x94,
	0xd6, 0x0f, 0xfe, 0x70, 0xe0, 0xec, 0x06, 0x17, 0xe8, 0x13, 0x00, 0x51, 0x6f, 0xa2, 0x84, 0xe7,
	0x84, 0x15, 0x5a, 0x5a, 0x5e, 0xe8, 0x89, 0x7a, 0x73, 0xa9, 0x01, 0x35, 0x00, 0xf6, 0xc8, 0x68,
	0xcc, 0x7a, 0x2a, 0x8d, 0x89, 0xa8, 0x24, 0xd5, 0x8e, 0x15, 0xa9, 0x16, 0x9a, 0x1b, 0x7a, 0x4c,
	0xbc, 0x35, 0x00, 0x7a, 0x74, 0x38, 0x06, 0x66, 0x1a, 0x0f, 0x24, 0xaf, 0x96, 0x46, 0x5c, 0xb1,
	0x52, 0x5a, 0xc1, 0x59, 0x2f, 0x78, 0x0a, 0xfe, 0x77, 0x54, 0x9a, 0x3b, 0x36, 0xfb, 0xa0, 0x91,
	0xa1, 0xd3, 0xc9, 0x30, 0x78, 0x0e, 0x8f, 0x2e, 0x69, 0x46, 0x25, 0x0d, 0xf9, 0xfb, 0x6f, 0x2b,
	0x9e, 0xb7, 0x63, 0x65, 0x73, 0x7c, 0xe8, 0x57, 0xfc, 0xbd, 0x4d, 0x51, 0xe6, 0xe7, 0x5f, 0x01,
	0xbc, 0xe9, 0x36, 0xa5, 0x07, 0x43, 0xbd, 0xb9, 0xfd, 0x13, 0x34, 0x81, 0x71, 0x42, 0xc5, 0x4e,
	0xf2, 0xd2, 0x77, 0x10, 0xc0, 0x28, 0xe7, 0x1b, 0x96, 0x51, 0xbf, 0x87, 0xa6, 0xe0, 0x92, 0x84,
	0x94, 0x92, 0xbd, 0xa3, 0x7e, 0x7f, 0xfd, 0x8f, 0x03, 0x63, 0xc5, 0xda, 0xd7, 0x25, 0x43, 0x21,
	0x9c, 0x7e, 0x30, 0xfb, 0xe8, 0x71, 0xa7, 0xb5, 0xdb, 0x56, 0xda, 0xfc, 0xc9, 0x7f, 0x9e, 0x9b,
	0xa5, 0x11, 0x9c, 0xa0, 0x2f, 0xc1, 0x6b, 0x3b, 0x47, 0xf3, 0x2e, 0xfe, 0x98, 0x8e, 0xb9, 0x7f,
	0xac, 0xeb, 0xe0, 0x04, 0xfd, 0x02, 0x0f, 0x6e, 0xa5, 0x03, 0x3d, 0xed, 0x82, 0xff, 0x8f, 0xaf,
	0xf9, 0xc3, 0x95, 0xf9, 0xb1, 0xad, 0x9a, 0x1f, 0xdb, 0xea, 0x1b, 0x45, 0x4f, 0x70, 0xb2, 0x19,
	0x69, 0xe4, 0xe2, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x07, 0x32, 0x3f, 0x14, 0x07, 0x00,
	0x00,
}
