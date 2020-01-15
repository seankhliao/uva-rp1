// Code generated by protoc-gen-go. DO NOT EDIT.
// source: primary.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Clusters struct {
	Clusters             []*Cluster `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Clusters) Reset()         { *m = Clusters{} }
func (m *Clusters) String() string { return proto.CompactTextString(m) }
func (*Clusters) ProtoMessage()    {}
func (*Clusters) Descriptor() ([]byte, []int) {
	return fileDescriptor_92d653641ecf6225, []int{0}
}

func (m *Clusters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Clusters.Unmarshal(m, b)
}
func (m *Clusters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Clusters.Marshal(b, m, deterministic)
}
func (m *Clusters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Clusters.Merge(m, src)
}
func (m *Clusters) XXX_Size() int {
	return xxx_messageInfo_Clusters.Size(m)
}
func (m *Clusters) XXX_DiscardUnknown() {
	xxx_messageInfo_Clusters.DiscardUnknown(m)
}

var xxx_messageInfo_Clusters proto.InternalMessageInfo

func (m *Clusters) GetClusters() []*Cluster {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type Cluster struct {
	Id                   string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Primary              *Endpoint `protobuf:"bytes,2,opt,name=primary,proto3" json:"primary,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Cluster) Reset()         { *m = Cluster{} }
func (m *Cluster) String() string { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()    {}
func (*Cluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_92d653641ecf6225, []int{1}
}

func (m *Cluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cluster.Unmarshal(m, b)
}
func (m *Cluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cluster.Marshal(b, m, deterministic)
}
func (m *Cluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cluster.Merge(m, src)
}
func (m *Cluster) XXX_Size() int {
	return xxx_messageInfo_Cluster.Size(m)
}
func (m *Cluster) XXX_DiscardUnknown() {
	xxx_messageInfo_Cluster.DiscardUnknown(m)
}

var xxx_messageInfo_Cluster proto.InternalMessageInfo

func (m *Cluster) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Cluster) GetPrimary() *Endpoint {
	if m != nil {
		return m.Primary
	}
	return nil
}

type Endpoint struct {
	Scheme               string   `protobuf:"bytes,1,opt,name=scheme,proto3" json:"scheme,omitempty"`
	Host                 string   `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Port                 string   `protobuf:"bytes,3,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_92d653641ecf6225, []int{2}
}

func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint.Unmarshal(m, b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
}
func (m *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(m, src)
}
func (m *Endpoint) XXX_Size() int {
	return xxx_messageInfo_Endpoint.Size(m)
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *Endpoint) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Endpoint) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

type Routes struct {
	Routes               []*Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Routes) Reset()         { *m = Routes{} }
func (m *Routes) String() string { return proto.CompactTextString(m) }
func (*Routes) ProtoMessage()    {}
func (*Routes) Descriptor() ([]byte, []int) {
	return fileDescriptor_92d653641ecf6225, []int{3}
}

func (m *Routes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Routes.Unmarshal(m, b)
}
func (m *Routes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Routes.Marshal(b, m, deterministic)
}
func (m *Routes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Routes.Merge(m, src)
}
func (m *Routes) XXX_Size() int {
	return xxx_messageInfo_Routes.Size(m)
}
func (m *Routes) XXX_DiscardUnknown() {
	xxx_messageInfo_Routes.DiscardUnknown(m)
}

var xxx_messageInfo_Routes proto.InternalMessageInfo

func (m *Routes) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type Route struct {
	Prefix               string    `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Cost                 int64     `protobuf:"varint,2,opt,name=cost,proto3" json:"cost,omitempty"`
	Endpoint             *Endpoint `protobuf:"bytes,3,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_92d653641ecf6225, []int{4}
}

func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *Route) GetCost() int64 {
	if m != nil {
		return m.Cost
	}
	return 0
}

func (m *Route) GetEndpoint() *Endpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

type ChannelRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChannelRequest) Reset()         { *m = ChannelRequest{} }
func (m *ChannelRequest) String() string { return proto.CompactTextString(m) }
func (*ChannelRequest) ProtoMessage()    {}
func (*ChannelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_92d653641ecf6225, []int{5}
}

func (m *ChannelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelRequest.Unmarshal(m, b)
}
func (m *ChannelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelRequest.Marshal(b, m, deterministic)
}
func (m *ChannelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelRequest.Merge(m, src)
}
func (m *ChannelRequest) XXX_Size() int {
	return xxx_messageInfo_ChannelRequest.Size(m)
}
func (m *ChannelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelRequest proto.InternalMessageInfo

type ChannelResponse struct {
	Endpoints            []*Endpoint `protobuf:"bytes,1,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ChannelResponse) Reset()         { *m = ChannelResponse{} }
func (m *ChannelResponse) String() string { return proto.CompactTextString(m) }
func (*ChannelResponse) ProtoMessage()    {}
func (*ChannelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_92d653641ecf6225, []int{6}
}

func (m *ChannelResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelResponse.Unmarshal(m, b)
}
func (m *ChannelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelResponse.Marshal(b, m, deterministic)
}
func (m *ChannelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelResponse.Merge(m, src)
}
func (m *ChannelResponse) XXX_Size() int {
	return xxx_messageInfo_ChannelResponse.Size(m)
}
func (m *ChannelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelResponse proto.InternalMessageInfo

func (m *ChannelResponse) GetEndpoints() []*Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

type HealthRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthRequest) Reset()         { *m = HealthRequest{} }
func (m *HealthRequest) String() string { return proto.CompactTextString(m) }
func (*HealthRequest) ProtoMessage()    {}
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_92d653641ecf6225, []int{7}
}

func (m *HealthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthRequest.Unmarshal(m, b)
}
func (m *HealthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthRequest.Marshal(b, m, deterministic)
}
func (m *HealthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthRequest.Merge(m, src)
}
func (m *HealthRequest) XXX_Size() int {
	return xxx_messageInfo_HealthRequest.Size(m)
}
func (m *HealthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthRequest proto.InternalMessageInfo

type HealthResponse struct {
	Endpoint             *Endpoint `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *HealthResponse) Reset()         { *m = HealthResponse{} }
func (m *HealthResponse) String() string { return proto.CompactTextString(m) }
func (*HealthResponse) ProtoMessage()    {}
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_92d653641ecf6225, []int{8}
}

func (m *HealthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthResponse.Unmarshal(m, b)
}
func (m *HealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthResponse.Marshal(b, m, deterministic)
}
func (m *HealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthResponse.Merge(m, src)
}
func (m *HealthResponse) XXX_Size() int {
	return xxx_messageInfo_HealthResponse.Size(m)
}
func (m *HealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthResponse proto.InternalMessageInfo

func (m *HealthResponse) GetEndpoint() *Endpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

// SecondaryInfo is the current state of the secondary
type SecondaryInfo struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Routes               []*Route `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"`
	CacheCapacity        int64    `protobuf:"varint,3,opt,name=cache_capacity,json=cacheCapacity,proto3" json:"cache_capacity,omitempty"`
	CacheEntries         int64    `protobuf:"varint,4,opt,name=cache_entries,json=cacheEntries,proto3" json:"cache_entries,omitempty"`
	CacheHits            int64    `protobuf:"varint,5,opt,name=cache_hits,json=cacheHits,proto3" json:"cache_hits,omitempty"`
	CacheMisses          int64    `protobuf:"varint,6,opt,name=cache_misses,json=cacheMisses,proto3" json:"cache_misses,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SecondaryInfo) Reset()         { *m = SecondaryInfo{} }
func (m *SecondaryInfo) String() string { return proto.CompactTextString(m) }
func (*SecondaryInfo) ProtoMessage()    {}
func (*SecondaryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_92d653641ecf6225, []int{9}
}

func (m *SecondaryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecondaryInfo.Unmarshal(m, b)
}
func (m *SecondaryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecondaryInfo.Marshal(b, m, deterministic)
}
func (m *SecondaryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecondaryInfo.Merge(m, src)
}
func (m *SecondaryInfo) XXX_Size() int {
	return xxx_messageInfo_SecondaryInfo.Size(m)
}
func (m *SecondaryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SecondaryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SecondaryInfo proto.InternalMessageInfo

func (m *SecondaryInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SecondaryInfo) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *SecondaryInfo) GetCacheCapacity() int64 {
	if m != nil {
		return m.CacheCapacity
	}
	return 0
}

func (m *SecondaryInfo) GetCacheEntries() int64 {
	if m != nil {
		return m.CacheEntries
	}
	return 0
}

func (m *SecondaryInfo) GetCacheHits() int64 {
	if m != nil {
		return m.CacheHits
	}
	return 0
}

func (m *SecondaryInfo) GetCacheMisses() int64 {
	if m != nil {
		return m.CacheMisses
	}
	return 0
}

// SecondaryControl is the desired state of the secondary
type SecondaryControl struct {
	Routes               []*Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	CacheCapacity        int64    `protobuf:"varint,2,opt,name=cache_capacity,json=cacheCapacity,proto3" json:"cache_capacity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SecondaryControl) Reset()         { *m = SecondaryControl{} }
func (m *SecondaryControl) String() string { return proto.CompactTextString(m) }
func (*SecondaryControl) ProtoMessage()    {}
func (*SecondaryControl) Descriptor() ([]byte, []int) {
	return fileDescriptor_92d653641ecf6225, []int{10}
}

func (m *SecondaryControl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecondaryControl.Unmarshal(m, b)
}
func (m *SecondaryControl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecondaryControl.Marshal(b, m, deterministic)
}
func (m *SecondaryControl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecondaryControl.Merge(m, src)
}
func (m *SecondaryControl) XXX_Size() int {
	return xxx_messageInfo_SecondaryControl.Size(m)
}
func (m *SecondaryControl) XXX_DiscardUnknown() {
	xxx_messageInfo_SecondaryControl.DiscardUnknown(m)
}

var xxx_messageInfo_SecondaryControl proto.InternalMessageInfo

func (m *SecondaryControl) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *SecondaryControl) GetCacheCapacity() int64 {
	if m != nil {
		return m.CacheCapacity
	}
	return 0
}

type UnregisterRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnregisterRequest) Reset()         { *m = UnregisterRequest{} }
func (m *UnregisterRequest) String() string { return proto.CompactTextString(m) }
func (*UnregisterRequest) ProtoMessage()    {}
func (*UnregisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_92d653641ecf6225, []int{11}
}

func (m *UnregisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnregisterRequest.Unmarshal(m, b)
}
func (m *UnregisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnregisterRequest.Marshal(b, m, deterministic)
}
func (m *UnregisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnregisterRequest.Merge(m, src)
}
func (m *UnregisterRequest) XXX_Size() int {
	return xxx_messageInfo_UnregisterRequest.Size(m)
}
func (m *UnregisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnregisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnregisterRequest proto.InternalMessageInfo

func (m *UnregisterRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UnregisterResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnregisterResponse) Reset()         { *m = UnregisterResponse{} }
func (m *UnregisterResponse) String() string { return proto.CompactTextString(m) }
func (*UnregisterResponse) ProtoMessage()    {}
func (*UnregisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_92d653641ecf6225, []int{12}
}

func (m *UnregisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnregisterResponse.Unmarshal(m, b)
}
func (m *UnregisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnregisterResponse.Marshal(b, m, deterministic)
}
func (m *UnregisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnregisterResponse.Merge(m, src)
}
func (m *UnregisterResponse) XXX_Size() int {
	return xxx_messageInfo_UnregisterResponse.Size(m)
}
func (m *UnregisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UnregisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UnregisterResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Clusters)(nil), "api.Clusters")
	proto.RegisterType((*Cluster)(nil), "api.Cluster")
	proto.RegisterType((*Endpoint)(nil), "api.Endpoint")
	proto.RegisterType((*Routes)(nil), "api.Routes")
	proto.RegisterType((*Route)(nil), "api.Route")
	proto.RegisterType((*ChannelRequest)(nil), "api.ChannelRequest")
	proto.RegisterType((*ChannelResponse)(nil), "api.ChannelResponse")
	proto.RegisterType((*HealthRequest)(nil), "api.HealthRequest")
	proto.RegisterType((*HealthResponse)(nil), "api.HealthResponse")
	proto.RegisterType((*SecondaryInfo)(nil), "api.SecondaryInfo")
	proto.RegisterType((*SecondaryControl)(nil), "api.SecondaryControl")
	proto.RegisterType((*UnregisterRequest)(nil), "api.UnregisterRequest")
	proto.RegisterType((*UnregisterResponse)(nil), "api.UnregisterResponse")
}

func init() { proto.RegisterFile("primary.proto", fileDescriptor_92d653641ecf6225) }

var fileDescriptor_92d653641ecf6225 = []byte{
	// 565 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdd, 0x6e, 0xd3, 0x3e,
	0x14, 0xc0, 0x97, 0x76, 0xcb, 0xda, 0xd3, 0xa6, 0xdb, 0xdc, 0xfd, 0xf7, 0x8f, 0x22, 0x21, 0x15,
	0x4f, 0x88, 0x22, 0x50, 0x85, 0xba, 0x49, 0x5c, 0x20, 0x71, 0x41, 0x99, 0x28, 0x48, 0xdc, 0x04,
	0x71, 0x09, 0x55, 0x48, 0x3d, 0x62, 0xa9, 0xb5, 0x83, 0xed, 0x4a, 0xf4, 0x51, 0x78, 0x2c, 0xde,
	0x08, 0xf5, 0xd8, 0x4e, 0x9b, 0x75, 0x20, 0xee, 0x8e, 0x7f, 0xe7, 0xfb, 0x23, 0x81, 0xa8, 0x54,
	0x7c, 0x99, 0xa9, 0xf5, 0xa8, 0x54, 0xd2, 0x48, 0xd2, 0xcc, 0x4a, 0x4e, 0xaf, 0xa1, 0x35, 0x59,
	0xac, 0xb4, 0x61, 0x4a, 0x93, 0x21, 0xb4, 0x72, 0x27, 0xc7, 0xc1, 0xa0, 0x39, 0xec, 0x8c, 0xbb,
	0xa3, 0xac, 0xe4, 0x23, 0x67, 0x90, 0x56, 0x5a, 0xfa, 0x1a, 0x8e, 0x1d, 0x24, 0x3d, 0x68, 0xf0,
	0x79, 0x1c, 0x0c, 0x82, 0x61, 0x3b, 0x6d, 0xf0, 0x39, 0x79, 0x0c, 0xc7, 0x2e, 0x4d, 0xdc, 0x18,
	0x04, 0xc3, 0xce, 0x38, 0xc2, 0x18, 0x37, 0x62, 0x5e, 0x4a, 0x2e, 0x4c, 0xea, 0xb5, 0xf4, 0x3d,
	0xb4, 0x3c, 0x24, 0x17, 0x10, 0xea, 0xbc, 0x60, 0x4b, 0xe6, 0x02, 0xb9, 0x17, 0x21, 0x70, 0x58,
	0x48, 0x6d, 0x30, 0x52, 0x3b, 0x45, 0x79, 0xc3, 0x4a, 0xa9, 0x4c, 0xdc, 0xb4, 0x6c, 0x23, 0xd3,
	0x67, 0x10, 0xa6, 0x72, 0x65, 0x98, 0x26, 0x14, 0x42, 0x85, 0x92, 0xeb, 0x00, 0x30, 0x3b, 0x2a,
	0x53, 0xa7, 0xa1, 0x5f, 0xe0, 0x08, 0xc1, 0x26, 0x6d, 0xa9, 0xd8, 0x2d, 0xff, 0xe1, 0xd3, 0xda,
	0xd7, 0x26, 0x45, 0xee, 0xd3, 0x36, 0x53, 0x94, 0xc9, 0x13, 0x68, 0x31, 0x57, 0x2e, 0xa6, 0xde,
	0x6b, 0xac, 0x52, 0xd3, 0x53, 0xe8, 0x4d, 0x8a, 0x4c, 0x08, 0xb6, 0x48, 0xd9, 0xf7, 0x15, 0xd3,
	0x86, 0xbe, 0x82, 0x93, 0x8a, 0xe8, 0x52, 0x0a, 0xcd, 0xc8, 0x53, 0x68, 0x7b, 0x07, 0x5f, 0xeb,
	0x9d, 0x80, 0x5b, 0x3d, 0x3d, 0x81, 0x68, 0xca, 0xb2, 0x85, 0x29, 0x7c, 0xc0, 0x97, 0xd0, 0xf3,
	0xc0, 0xc5, 0xdb, 0xad, 0x2f, 0xf8, 0x7b, 0x7d, 0xbf, 0x02, 0x88, 0x3e, 0xb2, 0x5c, 0x8a, 0x79,
	0xa6, 0xd6, 0xef, 0xc4, 0xad, 0xdc, 0x5b, 0xe2, 0x76, 0x8a, 0x8d, 0x3f, 0x4d, 0x91, 0x3c, 0x82,
	0x5e, 0x9e, 0xe5, 0x05, 0x9b, 0xe5, 0x59, 0x99, 0xe5, 0xdc, 0xac, 0x71, 0x2c, 0xcd, 0x34, 0x42,
	0x3a, 0x71, 0x90, 0x5c, 0x82, 0x05, 0x33, 0x26, 0x8c, 0xe2, 0x4c, 0xc7, 0x87, 0x68, 0xd5, 0x45,
	0x78, 0x63, 0x19, 0x79, 0x00, 0x60, 0x8d, 0x0a, 0x6e, 0x74, 0x7c, 0x84, 0x16, 0x6d, 0x24, 0x53,
	0x6e, 0x34, 0x79, 0x08, 0xd6, 0x7c, 0xb6, 0xe4, 0x5a, 0x33, 0x1d, 0x87, 0x68, 0xd0, 0x41, 0xf6,
	0x01, 0x11, 0xfd, 0x0c, 0xa7, 0x55, 0x4b, 0x13, 0x29, 0x8c, 0x92, 0x8b, 0x7f, 0xb9, 0x85, 0x7b,
	0xba, 0x68, 0xdc, 0xd3, 0x05, 0xbd, 0x84, 0xb3, 0x4f, 0x42, 0xb1, 0x6f, 0x1c, 0x3f, 0x04, 0xbb,
	0x84, 0xbb, 0x53, 0xa3, 0xe7, 0x40, 0x76, 0x8d, 0xec, 0x62, 0xc6, 0x02, 0xc2, 0xb7, 0x52, 0x6b,
	0x5e, 0x92, 0x6b, 0xe8, 0x59, 0xa9, 0xfa, 0xe2, 0xa2, 0xdd, 0xef, 0x4b, 0x27, 0xf5, 0x27, 0x3d,
	0x18, 0x06, 0xcf, 0x03, 0x32, 0x82, 0xae, 0xf5, 0x72, 0x17, 0xde, 0xd9, 0x76, 0xa1, 0x93, 0xdd,
	0x87, 0xb5, 0x1f, 0x1b, 0x38, 0xc4, 0x9d, 0xbe, 0x80, 0x96, 0xbb, 0x39, 0x4d, 0xfa, 0x36, 0x70,
	0xed, 0x28, 0x93, 0xf3, 0x3a, 0xb4, 0xe5, 0xd2, 0x03, 0x72, 0x05, 0xa1, 0xbd, 0x2d, 0x42, 0xd0,
	0xa2, 0x76, 0x79, 0x49, 0xbf, 0xc6, 0xbc, 0xd3, 0xf8, 0x67, 0x00, 0xc7, 0x7e, 0xee, 0x6f, 0xe0,
	0x2c, 0x75, 0x53, 0xa8, 0x76, 0xe2, 0x62, 0xd5, 0xce, 0x2e, 0xf9, 0xaf, 0xce, 0x9c, 0xbf, 0xeb,
	0x7b, 0x0a, 0xfd, 0xed, 0x34, 0xb7, 0x71, 0x2e, 0xd0, 0x67, 0x6f, 0x19, 0xc9, 0xff, 0x7b, 0xdc,
	0xd7, 0xf6, 0x35, 0xc4, 0xff, 0xdd, 0xd5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x0c, 0xe7,
	0x35, 0x00, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GossipClient is the client API for Gossip service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GossipClient interface {
	// Only gossip about successfully connected ones
	GossipClusters(ctx context.Context, opts ...grpc.CallOption) (Gossip_GossipClustersClient, error)
	// Only gossip about reachable routes
	GossipRoutes(ctx context.Context, opts ...grpc.CallOption) (Gossip_GossipRoutesClient, error)
}

type gossipClient struct {
	cc *grpc.ClientConn
}

func NewGossipClient(cc *grpc.ClientConn) GossipClient {
	return &gossipClient{cc}
}

func (c *gossipClient) GossipClusters(ctx context.Context, opts ...grpc.CallOption) (Gossip_GossipClustersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Gossip_serviceDesc.Streams[0], "/api.Gossip/GossipClusters", opts...)
	if err != nil {
		return nil, err
	}
	x := &gossipGossipClustersClient{stream}
	return x, nil
}

type Gossip_GossipClustersClient interface {
	Send(*Clusters) error
	Recv() (*Clusters, error)
	grpc.ClientStream
}

type gossipGossipClustersClient struct {
	grpc.ClientStream
}

func (x *gossipGossipClustersClient) Send(m *Clusters) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gossipGossipClustersClient) Recv() (*Clusters, error) {
	m := new(Clusters)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gossipClient) GossipRoutes(ctx context.Context, opts ...grpc.CallOption) (Gossip_GossipRoutesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Gossip_serviceDesc.Streams[1], "/api.Gossip/GossipRoutes", opts...)
	if err != nil {
		return nil, err
	}
	x := &gossipGossipRoutesClient{stream}
	return x, nil
}

type Gossip_GossipRoutesClient interface {
	Send(*Routes) error
	Recv() (*Routes, error)
	grpc.ClientStream
}

type gossipGossipRoutesClient struct {
	grpc.ClientStream
}

func (x *gossipGossipRoutesClient) Send(m *Routes) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gossipGossipRoutesClient) Recv() (*Routes, error) {
	m := new(Routes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GossipServer is the server API for Gossip service.
type GossipServer interface {
	// Only gossip about successfully connected ones
	GossipClusters(Gossip_GossipClustersServer) error
	// Only gossip about reachable routes
	GossipRoutes(Gossip_GossipRoutesServer) error
}

// UnimplementedGossipServer can be embedded to have forward compatible implementations.
type UnimplementedGossipServer struct {
}

func (*UnimplementedGossipServer) GossipClusters(srv Gossip_GossipClustersServer) error {
	return status.Errorf(codes.Unimplemented, "method GossipClusters not implemented")
}
func (*UnimplementedGossipServer) GossipRoutes(srv Gossip_GossipRoutesServer) error {
	return status.Errorf(codes.Unimplemented, "method GossipRoutes not implemented")
}

func RegisterGossipServer(s *grpc.Server, srv GossipServer) {
	s.RegisterService(&_Gossip_serviceDesc, srv)
}

func _Gossip_GossipClusters_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GossipServer).GossipClusters(&gossipGossipClustersServer{stream})
}

type Gossip_GossipClustersServer interface {
	Send(*Clusters) error
	Recv() (*Clusters, error)
	grpc.ServerStream
}

type gossipGossipClustersServer struct {
	grpc.ServerStream
}

func (x *gossipGossipClustersServer) Send(m *Clusters) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gossipGossipClustersServer) Recv() (*Clusters, error) {
	m := new(Clusters)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Gossip_GossipRoutes_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GossipServer).GossipRoutes(&gossipGossipRoutesServer{stream})
}

type Gossip_GossipRoutesServer interface {
	Send(*Routes) error
	Recv() (*Routes, error)
	grpc.ServerStream
}

type gossipGossipRoutesServer struct {
	grpc.ServerStream
}

func (x *gossipGossipRoutesServer) Send(m *Routes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gossipGossipRoutesServer) Recv() (*Routes, error) {
	m := new(Routes)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Gossip_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Gossip",
	HandlerType: (*GossipServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GossipClusters",
			Handler:       _Gossip_GossipClusters_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GossipRoutes",
			Handler:       _Gossip_GossipRoutes_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "primary.proto",
}

// InfoClient is the client API for Info service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InfoClient interface {
	// discover channels
	Channels(ctx context.Context, in *ChannelRequest, opts ...grpc.CallOption) (*ChannelResponse, error)
	// discover the health endpoint for prometheus
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
}

type infoClient struct {
	cc *grpc.ClientConn
}

func NewInfoClient(cc *grpc.ClientConn) InfoClient {
	return &infoClient{cc}
}

func (c *infoClient) Channels(ctx context.Context, in *ChannelRequest, opts ...grpc.CallOption) (*ChannelResponse, error) {
	out := new(ChannelResponse)
	err := c.cc.Invoke(ctx, "/api.Info/Channels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/api.Info/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InfoServer is the server API for Info service.
type InfoServer interface {
	// discover channels
	Channels(context.Context, *ChannelRequest) (*ChannelResponse, error)
	// discover the health endpoint for prometheus
	Health(context.Context, *HealthRequest) (*HealthResponse, error)
}

// UnimplementedInfoServer can be embedded to have forward compatible implementations.
type UnimplementedInfoServer struct {
}

func (*UnimplementedInfoServer) Channels(ctx context.Context, req *ChannelRequest) (*ChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Channels not implemented")
}
func (*UnimplementedInfoServer) Health(ctx context.Context, req *HealthRequest) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}

func RegisterInfoServer(s *grpc.Server, srv InfoServer) {
	s.RegisterService(&_Info_serviceDesc, srv)
}

func _Info_Channels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServer).Channels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Info/Channels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServer).Channels(ctx, req.(*ChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Info_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Info/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Info_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Info",
	HandlerType: (*InfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Channels",
			Handler:    _Info_Channels_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _Info_Health_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "primary.proto",
}

// ControlClient is the client API for Control service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ControlClient interface {
	RegisterSecondary(ctx context.Context, opts ...grpc.CallOption) (Control_RegisterSecondaryClient, error)
	UnregisterSecondary(ctx context.Context, in *UnregisterRequest, opts ...grpc.CallOption) (*UnregisterResponse, error)
}

type controlClient struct {
	cc *grpc.ClientConn
}

func NewControlClient(cc *grpc.ClientConn) ControlClient {
	return &controlClient{cc}
}

func (c *controlClient) RegisterSecondary(ctx context.Context, opts ...grpc.CallOption) (Control_RegisterSecondaryClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Control_serviceDesc.Streams[0], "/api.Control/RegisterSecondary", opts...)
	if err != nil {
		return nil, err
	}
	x := &controlRegisterSecondaryClient{stream}
	return x, nil
}

type Control_RegisterSecondaryClient interface {
	Send(*SecondaryInfo) error
	Recv() (*SecondaryControl, error)
	grpc.ClientStream
}

type controlRegisterSecondaryClient struct {
	grpc.ClientStream
}

func (x *controlRegisterSecondaryClient) Send(m *SecondaryInfo) error {
	return x.ClientStream.SendMsg(m)
}

func (x *controlRegisterSecondaryClient) Recv() (*SecondaryControl, error) {
	m := new(SecondaryControl)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *controlClient) UnregisterSecondary(ctx context.Context, in *UnregisterRequest, opts ...grpc.CallOption) (*UnregisterResponse, error) {
	out := new(UnregisterResponse)
	err := c.cc.Invoke(ctx, "/api.Control/UnregisterSecondary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControlServer is the server API for Control service.
type ControlServer interface {
	RegisterSecondary(Control_RegisterSecondaryServer) error
	UnregisterSecondary(context.Context, *UnregisterRequest) (*UnregisterResponse, error)
}

// UnimplementedControlServer can be embedded to have forward compatible implementations.
type UnimplementedControlServer struct {
}

func (*UnimplementedControlServer) RegisterSecondary(srv Control_RegisterSecondaryServer) error {
	return status.Errorf(codes.Unimplemented, "method RegisterSecondary not implemented")
}
func (*UnimplementedControlServer) UnregisterSecondary(ctx context.Context, req *UnregisterRequest) (*UnregisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterSecondary not implemented")
}

func RegisterControlServer(s *grpc.Server, srv ControlServer) {
	s.RegisterService(&_Control_serviceDesc, srv)
}

func _Control_RegisterSecondary_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ControlServer).RegisterSecondary(&controlRegisterSecondaryServer{stream})
}

type Control_RegisterSecondaryServer interface {
	Send(*SecondaryControl) error
	Recv() (*SecondaryInfo, error)
	grpc.ServerStream
}

type controlRegisterSecondaryServer struct {
	grpc.ServerStream
}

func (x *controlRegisterSecondaryServer) Send(m *SecondaryControl) error {
	return x.ServerStream.SendMsg(m)
}

func (x *controlRegisterSecondaryServer) Recv() (*SecondaryInfo, error) {
	m := new(SecondaryInfo)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Control_UnregisterSecondary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnregisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).UnregisterSecondary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Control/UnregisterSecondary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServer).UnregisterSecondary(ctx, req.(*UnregisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Control_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Control",
	HandlerType: (*ControlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnregisterSecondary",
			Handler:    _Control_UnregisterSecondary_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RegisterSecondary",
			Handler:       _Control_RegisterSecondary_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "primary.proto",
}
