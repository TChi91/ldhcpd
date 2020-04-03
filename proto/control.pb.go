// Code generated by protoc-gen-go. DO NOT EDIT.
// source: control.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type MACAddress struct {
	Address              string   `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MACAddress) Reset()         { *m = MACAddress{} }
func (m *MACAddress) String() string { return proto.CompactTextString(m) }
func (*MACAddress) ProtoMessage()    {}
func (*MACAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c5120591600887d, []int{0}
}

func (m *MACAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MACAddress.Unmarshal(m, b)
}
func (m *MACAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MACAddress.Marshal(b, m, deterministic)
}
func (m *MACAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MACAddress.Merge(m, src)
}
func (m *MACAddress) XXX_Size() int {
	return xxx_messageInfo_MACAddress.Size(m)
}
func (m *MACAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_MACAddress.DiscardUnknown(m)
}

var xxx_messageInfo_MACAddress proto.InternalMessageInfo

func (m *MACAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type Lease struct {
	MACAddress           string               `protobuf:"bytes,1,opt,name=MACAddress,proto3" json:"MACAddress,omitempty"`
	IPAddress            string               `protobuf:"bytes,2,opt,name=IPAddress,proto3" json:"IPAddress,omitempty"`
	LeaseEnd             *timestamp.Timestamp `protobuf:"bytes,3,opt,name=LeaseEnd,proto3" json:"LeaseEnd,omitempty"`
	Dynamic              bool                 `protobuf:"varint,4,opt,name=Dynamic,proto3" json:"Dynamic,omitempty"`
	Persistent           bool                 `protobuf:"varint,5,opt,name=Persistent,proto3" json:"Persistent,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Lease) Reset()         { *m = Lease{} }
func (m *Lease) String() string { return proto.CompactTextString(m) }
func (*Lease) ProtoMessage()    {}
func (*Lease) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c5120591600887d, []int{1}
}

func (m *Lease) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Lease.Unmarshal(m, b)
}
func (m *Lease) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Lease.Marshal(b, m, deterministic)
}
func (m *Lease) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lease.Merge(m, src)
}
func (m *Lease) XXX_Size() int {
	return xxx_messageInfo_Lease.Size(m)
}
func (m *Lease) XXX_DiscardUnknown() {
	xxx_messageInfo_Lease.DiscardUnknown(m)
}

var xxx_messageInfo_Lease proto.InternalMessageInfo

func (m *Lease) GetMACAddress() string {
	if m != nil {
		return m.MACAddress
	}
	return ""
}

func (m *Lease) GetIPAddress() string {
	if m != nil {
		return m.IPAddress
	}
	return ""
}

func (m *Lease) GetLeaseEnd() *timestamp.Timestamp {
	if m != nil {
		return m.LeaseEnd
	}
	return nil
}

func (m *Lease) GetDynamic() bool {
	if m != nil {
		return m.Dynamic
	}
	return false
}

func (m *Lease) GetPersistent() bool {
	if m != nil {
		return m.Persistent
	}
	return false
}

type Leases struct {
	List                 []*Lease `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Leases) Reset()         { *m = Leases{} }
func (m *Leases) String() string { return proto.CompactTextString(m) }
func (*Leases) ProtoMessage()    {}
func (*Leases) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c5120591600887d, []int{2}
}

func (m *Leases) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Leases.Unmarshal(m, b)
}
func (m *Leases) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Leases.Marshal(b, m, deterministic)
}
func (m *Leases) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Leases.Merge(m, src)
}
func (m *Leases) XXX_Size() int {
	return xxx_messageInfo_Leases.Size(m)
}
func (m *Leases) XXX_DiscardUnknown() {
	xxx_messageInfo_Leases.DiscardUnknown(m)
}

var xxx_messageInfo_Leases proto.InternalMessageInfo

func (m *Leases) GetList() []*Lease {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*MACAddress)(nil), "proto.MACAddress")
	proto.RegisterType((*Lease)(nil), "proto.Lease")
	proto.RegisterType((*Leases)(nil), "proto.Leases")
}

func init() {
	proto.RegisterFile("control.proto", fileDescriptor_0c5120591600887d)
}

var fileDescriptor_0c5120591600887d = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x50, 0xdd, 0x4a, 0xc3, 0x30,
	0x14, 0x5e, 0xdc, 0x8f, 0xdd, 0xe9, 0x76, 0x61, 0x2e, 0x24, 0x44, 0xd1, 0x92, 0x0b, 0x29, 0x82,
	0x1d, 0x54, 0xf4, 0xc2, 0xbb, 0x31, 0x87, 0x08, 0x13, 0x46, 0xf5, 0x05, 0xba, 0x35, 0x8e, 0xc2,
	0xda, 0x94, 0x26, 0x0a, 0x7b, 0x32, 0x5f, 0xc8, 0x07, 0x91, 0x26, 0x8d, 0xeb, 0x94, 0x5e, 0xb5,
	0xdf, 0xf9, 0x7e, 0xce, 0x97, 0x03, 0xe3, 0xb5, 0xc8, 0x55, 0x29, 0xb6, 0x41, 0x51, 0x0a, 0x25,
	0x70, 0x5f, 0x7f, 0xe8, 0xe5, 0x46, 0x88, 0xcd, 0x96, 0x4f, 0x34, 0x5a, 0x7d, 0xbc, 0x4f, 0x54,
	0x9a, 0x71, 0xa9, 0xe2, 0xac, 0x30, 0x3a, 0x7a, 0xf6, 0x57, 0xc0, 0xb3, 0x42, 0xed, 0x0c, 0xc9,
	0xae, 0x00, 0x5e, 0xa6, 0xb3, 0x69, 0x92, 0x94, 0x5c, 0x4a, 0x4c, 0xe0, 0xb8, 0xfe, 0x25, 0xc8,
	0x43, 0xfe, 0x30, 0xb2, 0x90, 0x7d, 0x21, 0xe8, 0x2f, 0x78, 0x2c, 0x39, 0xbe, 0x68, 0x3a, 0x6a,
	0x59, 0x33, 0xe3, 0x1c, 0x86, 0xcf, 0x4b, 0x4b, 0x1f, 0x69, 0x7a, 0x3f, 0xc0, 0xf7, 0xe0, 0xe8,
	0x98, 0x79, 0x9e, 0x90, 0xae, 0x87, 0x7c, 0x37, 0xa4, 0x81, 0xe9, 0x17, 0xd8, 0x7e, 0xc1, 0x9b,
	0x7d, 0x40, 0xf4, 0xab, 0xad, 0x9a, 0x3d, 0xee, 0xf2, 0x38, 0x4b, 0xd7, 0xa4, 0xe7, 0x21, 0xdf,
	0x89, 0x2c, 0xac, 0xfa, 0x2c, 0x79, 0x29, 0x53, 0xa9, 0x78, 0xae, 0x48, 0x5f, 0x93, 0x8d, 0x09,
	0xbb, 0x86, 0x81, 0x4e, 0x91, 0xd8, 0x83, 0xde, 0x22, 0x95, 0x8a, 0x20, 0xaf, 0xeb, 0xbb, 0xe1,
	0xc8, 0x2c, 0x0c, 0x34, 0x19, 0x69, 0x26, 0xfc, 0x46, 0x30, 0xd2, 0x78, 0x66, 0x2e, 0x8d, 0x43,
	0x70, 0x5e, 0xb9, 0x32, 0x0f, 0x3f, 0x30, 0xd0, 0xd3, 0x7f, 0xb5, 0xe7, 0xd5, 0x59, 0x59, 0x07,
	0xdf, 0x80, 0xf3, 0x64, 0x3d, 0x27, 0xb5, 0x67, 0x7f, 0x1f, 0x7a, 0x10, 0xc3, 0x3a, 0xf8, 0x0e,
	0xa0, 0xda, 0x5d, 0x77, 0x6c, 0x89, 0xa5, 0xe3, 0xa6, 0x4b, 0xb2, 0x0e, 0x7e, 0x00, 0x37, 0xe2,
	0x99, 0xf8, 0xe4, 0xad, 0x8b, 0x5a, 0x1b, 0xae, 0x06, 0x7a, 0x72, 0xfb, 0x13, 0x00, 0x00, 0xff,
	0xff, 0xc9, 0x41, 0xbd, 0xc9, 0x51, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LeaseControlClient is the client API for LeaseControl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LeaseControlClient interface {
	SetLease(ctx context.Context, in *Lease, opts ...grpc.CallOption) (*empty.Empty, error)
	GetLease(ctx context.Context, in *MACAddress, opts ...grpc.CallOption) (*Lease, error)
	ListLeases(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Leases, error)
	RemoveLease(ctx context.Context, in *MACAddress, opts ...grpc.CallOption) (*empty.Empty, error)
}

type leaseControlClient struct {
	cc grpc.ClientConnInterface
}

func NewLeaseControlClient(cc grpc.ClientConnInterface) LeaseControlClient {
	return &leaseControlClient{cc}
}

func (c *leaseControlClient) SetLease(ctx context.Context, in *Lease, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.LeaseControl/SetLease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leaseControlClient) GetLease(ctx context.Context, in *MACAddress, opts ...grpc.CallOption) (*Lease, error) {
	out := new(Lease)
	err := c.cc.Invoke(ctx, "/proto.LeaseControl/GetLease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leaseControlClient) ListLeases(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Leases, error) {
	out := new(Leases)
	err := c.cc.Invoke(ctx, "/proto.LeaseControl/ListLeases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leaseControlClient) RemoveLease(ctx context.Context, in *MACAddress, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.LeaseControl/RemoveLease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LeaseControlServer is the server API for LeaseControl service.
type LeaseControlServer interface {
	SetLease(context.Context, *Lease) (*empty.Empty, error)
	GetLease(context.Context, *MACAddress) (*Lease, error)
	ListLeases(context.Context, *empty.Empty) (*Leases, error)
	RemoveLease(context.Context, *MACAddress) (*empty.Empty, error)
}

// UnimplementedLeaseControlServer can be embedded to have forward compatible implementations.
type UnimplementedLeaseControlServer struct {
}

func (*UnimplementedLeaseControlServer) SetLease(ctx context.Context, req *Lease) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLease not implemented")
}
func (*UnimplementedLeaseControlServer) GetLease(ctx context.Context, req *MACAddress) (*Lease, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLease not implemented")
}
func (*UnimplementedLeaseControlServer) ListLeases(ctx context.Context, req *empty.Empty) (*Leases, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLeases not implemented")
}
func (*UnimplementedLeaseControlServer) RemoveLease(ctx context.Context, req *MACAddress) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveLease not implemented")
}

func RegisterLeaseControlServer(s *grpc.Server, srv LeaseControlServer) {
	s.RegisterService(&_LeaseControl_serviceDesc, srv)
}

func _LeaseControl_SetLease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Lease)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaseControlServer).SetLease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LeaseControl/SetLease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaseControlServer).SetLease(ctx, req.(*Lease))
	}
	return interceptor(ctx, in, info, handler)
}

func _LeaseControl_GetLease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MACAddress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaseControlServer).GetLease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LeaseControl/GetLease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaseControlServer).GetLease(ctx, req.(*MACAddress))
	}
	return interceptor(ctx, in, info, handler)
}

func _LeaseControl_ListLeases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaseControlServer).ListLeases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LeaseControl/ListLeases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaseControlServer).ListLeases(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _LeaseControl_RemoveLease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MACAddress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaseControlServer).RemoveLease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LeaseControl/RemoveLease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaseControlServer).RemoveLease(ctx, req.(*MACAddress))
	}
	return interceptor(ctx, in, info, handler)
}

var _LeaseControl_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.LeaseControl",
	HandlerType: (*LeaseControlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetLease",
			Handler:    _LeaseControl_SetLease_Handler,
		},
		{
			MethodName: "GetLease",
			Handler:    _LeaseControl_GetLease_Handler,
		},
		{
			MethodName: "ListLeases",
			Handler:    _LeaseControl_ListLeases_Handler,
		},
		{
			MethodName: "RemoveLease",
			Handler:    _LeaseControl_RemoveLease_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "control.proto",
}