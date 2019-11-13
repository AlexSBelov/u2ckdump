// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

package msg

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

type SearchIPv4Request struct {
	Query                uint32   `protobuf:"varint,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchIPv4Request) Reset()         { *m = SearchIPv4Request{} }
func (m *SearchIPv4Request) String() string { return proto.CompactTextString(m) }
func (*SearchIPv4Request) ProtoMessage()    {}
func (*SearchIPv4Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

func (m *SearchIPv4Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchIPv4Request.Unmarshal(m, b)
}
func (m *SearchIPv4Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchIPv4Request.Marshal(b, m, deterministic)
}
func (m *SearchIPv4Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchIPv4Request.Merge(m, src)
}
func (m *SearchIPv4Request) XXX_Size() int {
	return xxx_messageInfo_SearchIPv4Request.Size(m)
}
func (m *SearchIPv4Request) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchIPv4Request.DiscardUnknown(m)
}

var xxx_messageInfo_SearchIPv4Request proto.InternalMessageInfo

func (m *SearchIPv4Request) GetQuery() uint32 {
	if m != nil {
		return m.Query
	}
	return 0
}

type SearchIPv6Request struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchIPv6Request) Reset()         { *m = SearchIPv6Request{} }
func (m *SearchIPv6Request) String() string { return proto.CompactTextString(m) }
func (*SearchIPv6Request) ProtoMessage()    {}
func (*SearchIPv6Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{1}
}

func (m *SearchIPv6Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchIPv6Request.Unmarshal(m, b)
}
func (m *SearchIPv6Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchIPv6Request.Marshal(b, m, deterministic)
}
func (m *SearchIPv6Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchIPv6Request.Merge(m, src)
}
func (m *SearchIPv6Request) XXX_Size() int {
	return xxx_messageInfo_SearchIPv6Request.Size(m)
}
func (m *SearchIPv6Request) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchIPv6Request.DiscardUnknown(m)
}

var xxx_messageInfo_SearchIPv6Request proto.InternalMessageInfo

func (m *SearchIPv6Request) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type SearchURLRequest struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchURLRequest) Reset()         { *m = SearchURLRequest{} }
func (m *SearchURLRequest) String() string { return proto.CompactTextString(m) }
func (*SearchURLRequest) ProtoMessage()    {}
func (*SearchURLRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{2}
}

func (m *SearchURLRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchURLRequest.Unmarshal(m, b)
}
func (m *SearchURLRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchURLRequest.Marshal(b, m, deterministic)
}
func (m *SearchURLRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchURLRequest.Merge(m, src)
}
func (m *SearchURLRequest) XXX_Size() int {
	return xxx_messageInfo_SearchURLRequest.Size(m)
}
func (m *SearchURLRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchURLRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchURLRequest proto.InternalMessageInfo

func (m *SearchURLRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type SearchDomainRequest struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchDomainRequest) Reset()         { *m = SearchDomainRequest{} }
func (m *SearchDomainRequest) String() string { return proto.CompactTextString(m) }
func (*SearchDomainRequest) ProtoMessage()    {}
func (*SearchDomainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{3}
}

func (m *SearchDomainRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchDomainRequest.Unmarshal(m, b)
}
func (m *SearchDomainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchDomainRequest.Marshal(b, m, deterministic)
}
func (m *SearchDomainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchDomainRequest.Merge(m, src)
}
func (m *SearchDomainRequest) XXX_Size() int {
	return xxx_messageInfo_SearchDomainRequest.Size(m)
}
func (m *SearchDomainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchDomainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchDomainRequest proto.InternalMessageInfo

func (m *SearchDomainRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type SearchResponse struct {
	Results              []*Content `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{4}
}

func (m *SearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse.Unmarshal(m, b)
}
func (m *SearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse.Marshal(b, m, deterministic)
}
func (m *SearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse.Merge(m, src)
}
func (m *SearchResponse) XXX_Size() int {
	return xxx_messageInfo_SearchResponse.Size(m)
}
func (m *SearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse proto.InternalMessageInfo

func (m *SearchResponse) GetResults() []*Content {
	if m != nil {
		return m.Results
	}
	return nil
}

type Content struct {
	Id                   int32          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	EntryType            int32          `protobuf:"varint,2,opt,name=entryType,proto3" json:"entryType,omitempty"`
	HttpsBlock           int32          `protobuf:"varint,3,opt,name=httpsBlock,proto3" json:"httpsBlock,omitempty"`
	RegistryUpdateTime   int64          `protobuf:"varint,4,opt,name=registryUpdateTime,proto3" json:"registryUpdateTime,omitempty"`
	Decision             *Decision      `protobuf:"bytes,5,opt,name=decision,proto3" json:"decision,omitempty"`
	IncludeTime          int64          `protobuf:"varint,6,opt,name=includeTime,proto3" json:"includeTime,omitempty"`
	BlockType            string         `protobuf:"bytes,7,opt,name=blockType,proto3" json:"blockType,omitempty"`
	Hash                 string         `protobuf:"bytes,8,opt,name=hash,proto3" json:"hash,omitempty"`
	Ts                   int64          `protobuf:"varint,9,opt,name=ts,proto3" json:"ts,omitempty"`
	U2Hash               uint32         `protobuf:"varint,10,opt,name=u2Hash,proto3" json:"u2Hash,omitempty"`
	Urls                 []*URL         `protobuf:"bytes,11,rep,name=urls,proto3" json:"urls,omitempty"`
	Ip4S                 []*IPv4Address `protobuf:"bytes,12,rep,name=ip4s,proto3" json:"ip4s,omitempty"`
	Ip6S                 []*IPv6Address `protobuf:"bytes,13,rep,name=ip6s,proto3" json:"ip6s,omitempty"`
	Subnets              []*Subnet      `protobuf:"bytes,14,rep,name=subnets,proto3" json:"subnets,omitempty"`
	Subnet6S             []*Subnet6     `protobuf:"bytes,15,rep,name=subnet6s,proto3" json:"subnet6s,omitempty"`
	Domains              []*Domain      `protobuf:"bytes,16,rep,name=domains,proto3" json:"domains,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Content) Reset()         { *m = Content{} }
func (m *Content) String() string { return proto.CompactTextString(m) }
func (*Content) ProtoMessage()    {}
func (*Content) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{5}
}

func (m *Content) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Content.Unmarshal(m, b)
}
func (m *Content) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Content.Marshal(b, m, deterministic)
}
func (m *Content) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Content.Merge(m, src)
}
func (m *Content) XXX_Size() int {
	return xxx_messageInfo_Content.Size(m)
}
func (m *Content) XXX_DiscardUnknown() {
	xxx_messageInfo_Content.DiscardUnknown(m)
}

var xxx_messageInfo_Content proto.InternalMessageInfo

func (m *Content) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Content) GetEntryType() int32 {
	if m != nil {
		return m.EntryType
	}
	return 0
}

func (m *Content) GetHttpsBlock() int32 {
	if m != nil {
		return m.HttpsBlock
	}
	return 0
}

func (m *Content) GetRegistryUpdateTime() int64 {
	if m != nil {
		return m.RegistryUpdateTime
	}
	return 0
}

func (m *Content) GetDecision() *Decision {
	if m != nil {
		return m.Decision
	}
	return nil
}

func (m *Content) GetIncludeTime() int64 {
	if m != nil {
		return m.IncludeTime
	}
	return 0
}

func (m *Content) GetBlockType() string {
	if m != nil {
		return m.BlockType
	}
	return ""
}

func (m *Content) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Content) GetTs() int64 {
	if m != nil {
		return m.Ts
	}
	return 0
}

func (m *Content) GetU2Hash() uint32 {
	if m != nil {
		return m.U2Hash
	}
	return 0
}

func (m *Content) GetUrls() []*URL {
	if m != nil {
		return m.Urls
	}
	return nil
}

func (m *Content) GetIp4S() []*IPv4Address {
	if m != nil {
		return m.Ip4S
	}
	return nil
}

func (m *Content) GetIp6S() []*IPv6Address {
	if m != nil {
		return m.Ip6S
	}
	return nil
}

func (m *Content) GetSubnets() []*Subnet {
	if m != nil {
		return m.Subnets
	}
	return nil
}

func (m *Content) GetSubnet6S() []*Subnet6 {
	if m != nil {
		return m.Subnet6S
	}
	return nil
}

func (m *Content) GetDomains() []*Domain {
	if m != nil {
		return m.Domains
	}
	return nil
}

type Decision struct {
	Date                 string   `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Number               string   `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	Org                  string   `protobuf:"bytes,3,opt,name=org,proto3" json:"org,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Decision) Reset()         { *m = Decision{} }
func (m *Decision) String() string { return proto.CompactTextString(m) }
func (*Decision) ProtoMessage()    {}
func (*Decision) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{6}
}

func (m *Decision) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Decision.Unmarshal(m, b)
}
func (m *Decision) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Decision.Marshal(b, m, deterministic)
}
func (m *Decision) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Decision.Merge(m, src)
}
func (m *Decision) XXX_Size() int {
	return xxx_messageInfo_Decision.Size(m)
}
func (m *Decision) XXX_DiscardUnknown() {
	xxx_messageInfo_Decision.DiscardUnknown(m)
}

var xxx_messageInfo_Decision proto.InternalMessageInfo

func (m *Decision) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *Decision) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Decision) GetOrg() string {
	if m != nil {
		return m.Org
	}
	return ""
}

type IPv4Address struct {
	Ip4                  uint32   `protobuf:"varint,1,opt,name=ip4,proto3" json:"ip4,omitempty"`
	Ts                   int64    `protobuf:"varint,2,opt,name=ts,proto3" json:"ts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IPv4Address) Reset()         { *m = IPv4Address{} }
func (m *IPv4Address) String() string { return proto.CompactTextString(m) }
func (*IPv4Address) ProtoMessage()    {}
func (*IPv4Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{7}
}

func (m *IPv4Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPv4Address.Unmarshal(m, b)
}
func (m *IPv4Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPv4Address.Marshal(b, m, deterministic)
}
func (m *IPv4Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPv4Address.Merge(m, src)
}
func (m *IPv4Address) XXX_Size() int {
	return xxx_messageInfo_IPv4Address.Size(m)
}
func (m *IPv4Address) XXX_DiscardUnknown() {
	xxx_messageInfo_IPv4Address.DiscardUnknown(m)
}

var xxx_messageInfo_IPv4Address proto.InternalMessageInfo

func (m *IPv4Address) GetIp4() uint32 {
	if m != nil {
		return m.Ip4
	}
	return 0
}

func (m *IPv4Address) GetTs() int64 {
	if m != nil {
		return m.Ts
	}
	return 0
}

type IPv6Address struct {
	Ip6                  string   `protobuf:"bytes,1,opt,name=ip6,proto3" json:"ip6,omitempty"`
	Ts                   int64    `protobuf:"varint,2,opt,name=ts,proto3" json:"ts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IPv6Address) Reset()         { *m = IPv6Address{} }
func (m *IPv6Address) String() string { return proto.CompactTextString(m) }
func (*IPv6Address) ProtoMessage()    {}
func (*IPv6Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{8}
}

func (m *IPv6Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPv6Address.Unmarshal(m, b)
}
func (m *IPv6Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPv6Address.Marshal(b, m, deterministic)
}
func (m *IPv6Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPv6Address.Merge(m, src)
}
func (m *IPv6Address) XXX_Size() int {
	return xxx_messageInfo_IPv6Address.Size(m)
}
func (m *IPv6Address) XXX_DiscardUnknown() {
	xxx_messageInfo_IPv6Address.DiscardUnknown(m)
}

var xxx_messageInfo_IPv6Address proto.InternalMessageInfo

func (m *IPv6Address) GetIp6() string {
	if m != nil {
		return m.Ip6
	}
	return ""
}

func (m *IPv6Address) GetTs() int64 {
	if m != nil {
		return m.Ts
	}
	return 0
}

type Subnet struct {
	Subnet               string   `protobuf:"bytes,1,opt,name=subnet,proto3" json:"subnet,omitempty"`
	Ts                   int64    `protobuf:"varint,2,opt,name=ts,proto3" json:"ts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Subnet) Reset()         { *m = Subnet{} }
func (m *Subnet) String() string { return proto.CompactTextString(m) }
func (*Subnet) ProtoMessage()    {}
func (*Subnet) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{9}
}

func (m *Subnet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subnet.Unmarshal(m, b)
}
func (m *Subnet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subnet.Marshal(b, m, deterministic)
}
func (m *Subnet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subnet.Merge(m, src)
}
func (m *Subnet) XXX_Size() int {
	return xxx_messageInfo_Subnet.Size(m)
}
func (m *Subnet) XXX_DiscardUnknown() {
	xxx_messageInfo_Subnet.DiscardUnknown(m)
}

var xxx_messageInfo_Subnet proto.InternalMessageInfo

func (m *Subnet) GetSubnet() string {
	if m != nil {
		return m.Subnet
	}
	return ""
}

func (m *Subnet) GetTs() int64 {
	if m != nil {
		return m.Ts
	}
	return 0
}

type Subnet6 struct {
	Subnet6              string   `protobuf:"bytes,1,opt,name=subnet6,proto3" json:"subnet6,omitempty"`
	Ts                   int64    `protobuf:"varint,2,opt,name=ts,proto3" json:"ts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Subnet6) Reset()         { *m = Subnet6{} }
func (m *Subnet6) String() string { return proto.CompactTextString(m) }
func (*Subnet6) ProtoMessage()    {}
func (*Subnet6) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{10}
}

func (m *Subnet6) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subnet6.Unmarshal(m, b)
}
func (m *Subnet6) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subnet6.Marshal(b, m, deterministic)
}
func (m *Subnet6) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subnet6.Merge(m, src)
}
func (m *Subnet6) XXX_Size() int {
	return xxx_messageInfo_Subnet6.Size(m)
}
func (m *Subnet6) XXX_DiscardUnknown() {
	xxx_messageInfo_Subnet6.DiscardUnknown(m)
}

var xxx_messageInfo_Subnet6 proto.InternalMessageInfo

func (m *Subnet6) GetSubnet6() string {
	if m != nil {
		return m.Subnet6
	}
	return ""
}

func (m *Subnet6) GetTs() int64 {
	if m != nil {
		return m.Ts
	}
	return 0
}

type URL struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Ts                   int64    `protobuf:"varint,2,opt,name=ts,proto3" json:"ts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *URL) Reset()         { *m = URL{} }
func (m *URL) String() string { return proto.CompactTextString(m) }
func (*URL) ProtoMessage()    {}
func (*URL) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{11}
}

func (m *URL) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_URL.Unmarshal(m, b)
}
func (m *URL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_URL.Marshal(b, m, deterministic)
}
func (m *URL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_URL.Merge(m, src)
}
func (m *URL) XXX_Size() int {
	return xxx_messageInfo_URL.Size(m)
}
func (m *URL) XXX_DiscardUnknown() {
	xxx_messageInfo_URL.DiscardUnknown(m)
}

var xxx_messageInfo_URL proto.InternalMessageInfo

func (m *URL) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *URL) GetTs() int64 {
	if m != nil {
		return m.Ts
	}
	return 0
}

type Domain struct {
	Domain               string   `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Ts                   int64    `protobuf:"varint,2,opt,name=ts,proto3" json:"ts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Domain) Reset()         { *m = Domain{} }
func (m *Domain) String() string { return proto.CompactTextString(m) }
func (*Domain) ProtoMessage()    {}
func (*Domain) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{12}
}

func (m *Domain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Domain.Unmarshal(m, b)
}
func (m *Domain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Domain.Marshal(b, m, deterministic)
}
func (m *Domain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Domain.Merge(m, src)
}
func (m *Domain) XXX_Size() int {
	return xxx_messageInfo_Domain.Size(m)
}
func (m *Domain) XXX_DiscardUnknown() {
	xxx_messageInfo_Domain.DiscardUnknown(m)
}

var xxx_messageInfo_Domain proto.InternalMessageInfo

func (m *Domain) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Domain) GetTs() int64 {
	if m != nil {
		return m.Ts
	}
	return 0
}

func init() {
	proto.RegisterType((*SearchIPv4Request)(nil), "SearchIPv4Request")
	proto.RegisterType((*SearchIPv6Request)(nil), "SearchIPv6Request")
	proto.RegisterType((*SearchURLRequest)(nil), "SearchURLRequest")
	proto.RegisterType((*SearchDomainRequest)(nil), "SearchDomainRequest")
	proto.RegisterType((*SearchResponse)(nil), "SearchResponse")
	proto.RegisterType((*Content)(nil), "Content")
	proto.RegisterType((*Decision)(nil), "Decision")
	proto.RegisterType((*IPv4Address)(nil), "IPv4Address")
	proto.RegisterType((*IPv6Address)(nil), "IPv6Address")
	proto.RegisterType((*Subnet)(nil), "Subnet")
	proto.RegisterType((*Subnet6)(nil), "Subnet6")
	proto.RegisterType((*URL)(nil), "URL")
	proto.RegisterType((*Domain)(nil), "Domain")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899) }

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 589 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0x15, 0x98, 0x60, 0x3c, 0x24, 0x84, 0x6c, 0xab, 0x6a, 0x55, 0x55, 0x95, 0x8b, 0x5a, 0x95,
	0xaa, 0x92, 0xd3, 0x12, 0xe4, 0x7b, 0x9b, 0xa8, 0x4a, 0xa5, 0x1c, 0xaa, 0x25, 0xfc, 0x00, 0xc0,
	0x2b, 0xb0, 0x6a, 0x6c, 0x67, 0x67, 0x1d, 0x89, 0xff, 0xdd, 0x1f, 0x50, 0xed, 0x17, 0x98, 0x40,
	0x73, 0xe8, 0x6d, 0xe7, 0xcd, 0x9b, 0xe7, 0xd9, 0x37, 0xeb, 0x81, 0x60, 0x8d, 0xcb, 0xa8, 0x14,
	0x85, 0x2c, 0x06, 0x9f, 0xe0, 0x62, 0xc2, 0x67, 0x62, 0xb1, 0xfa, 0xf9, 0xeb, 0x71, 0xcc, 0xf8,
	0x43, 0xc5, 0x51, 0x92, 0x97, 0x70, 0xf2, 0x50, 0x71, 0xb1, 0xa1, 0x8d, 0xb0, 0x31, 0x3c, 0x63,
	0x26, 0xd8, 0xa3, 0xc6, 0x47, 0xa9, 0x81, 0xa3, 0x0e, 0xa1, 0x6f, 0xa8, 0x53, 0x76, 0xf7, 0x3c,
	0xf3, 0x33, 0xbc, 0x30, 0xcc, 0x9b, 0x62, 0x3d, 0x4b, 0xf3, 0xe7, 0xc9, 0x63, 0xe8, 0x19, 0x32,
	0xe3, 0x58, 0x16, 0x39, 0x72, 0x32, 0x00, 0x5f, 0x70, 0xac, 0x32, 0x89, 0xb4, 0x11, 0x7a, 0xc3,
	0xee, 0xa8, 0x13, 0x5d, 0x17, 0xb9, 0xe4, 0xb9, 0x64, 0x2e, 0x31, 0xf8, 0xe3, 0x81, 0x6f, 0x41,
	0xd2, 0x83, 0x66, 0x9a, 0x68, 0xd1, 0x13, 0xd6, 0x4c, 0x13, 0xf2, 0x06, 0x02, 0x9e, 0x4b, 0xb1,
	0xb9, 0xdf, 0x94, 0x9c, 0x36, 0x35, 0xbc, 0x03, 0xc8, 0x5b, 0x80, 0x95, 0x94, 0x25, 0x7e, 0xcf,
	0x8a, 0xc5, 0x6f, 0xea, 0xe9, 0x74, 0x0d, 0x21, 0x11, 0x10, 0xc1, 0x97, 0x29, 0x4a, 0xb1, 0x99,
	0x96, 0xc9, 0x4c, 0xf2, 0xfb, 0x74, 0xcd, 0x69, 0x2b, 0x6c, 0x0c, 0x3d, 0x76, 0x24, 0x43, 0x3e,
	0x40, 0x27, 0xe1, 0x8b, 0x14, 0xd3, 0x22, 0xa7, 0x27, 0x61, 0x63, 0xd8, 0x1d, 0x05, 0xd1, 0x8d,
	0x05, 0xd8, 0x36, 0x45, 0x42, 0xe8, 0xa6, 0xf9, 0x22, 0xab, 0x12, 0xa3, 0xd7, 0xd6, 0x7a, 0x75,
	0x48, 0xb5, 0x3d, 0x57, 0x1d, 0xe8, 0xb6, 0x7d, 0x6d, 0xd1, 0x0e, 0x20, 0x04, 0x5a, 0xab, 0x19,
	0xae, 0x68, 0x47, 0x27, 0xf4, 0x59, 0x5d, 0x5c, 0x22, 0x0d, 0xb4, 0x54, 0x53, 0x22, 0x79, 0x05,
	0xed, 0x6a, 0x74, 0xab, 0x58, 0xa0, 0x67, 0x6c, 0x23, 0x42, 0xa1, 0x55, 0x89, 0x0c, 0x69, 0x57,
	0xbb, 0xd9, 0x8a, 0xd4, 0x00, 0x35, 0x42, 0x42, 0x68, 0xa5, 0xe5, 0x18, 0xe9, 0xa9, 0xce, 0x9c,
	0x46, 0xea, 0xc1, 0x7c, 0x4b, 0x12, 0xc1, 0x11, 0x99, 0xce, 0x18, 0x46, 0x8c, 0xf4, 0x6c, 0xc7,
	0x88, 0x6b, 0x8c, 0x18, 0xc9, 0x3b, 0xf0, 0xb1, 0x9a, 0xe7, 0x5c, 0x22, 0xed, 0x69, 0x92, 0x1f,
	0x4d, 0x74, 0xcc, 0x1c, 0x4e, 0xde, 0x43, 0xc7, 0x1c, 0x63, 0xa4, 0xe7, 0x76, 0xa4, 0x86, 0x13,
	0xb3, 0x6d, 0x46, 0x09, 0x25, 0xfa, 0xc1, 0x20, 0xed, 0x5b, 0x21, 0xfb, 0x80, 0x1c, 0x3e, 0xb8,
	0x85, 0x8e, 0xf3, 0x56, 0x39, 0xa2, 0x86, 0x60, 0x5f, 0x93, 0x3e, 0x2b, 0x07, 0xf2, 0x6a, 0x3d,
	0xe7, 0x42, 0xcf, 0x3d, 0x60, 0x36, 0x22, 0x7d, 0xf0, 0x0a, 0xb1, 0xd4, 0xd3, 0x0e, 0x98, 0x3a,
	0x0e, 0x2e, 0xa1, 0x5b, 0xbb, 0xac, 0x22, 0xa4, 0xe5, 0xd8, 0xfe, 0x1b, 0xea, 0x68, 0xcd, 0x6d,
	0x3a, 0x73, 0x6d, 0x41, 0xbc, 0x57, 0x10, 0xdb, 0x8f, 0xab, 0xe3, 0x41, 0xc1, 0x17, 0x68, 0x9b,
	0x3b, 0xaa, 0xae, 0xcc, 0x25, 0x2d, 0xdd, 0x46, 0x07, 0x15, 0x57, 0xe0, 0x5b, 0x57, 0x08, 0x75,
	0xa6, 0xba, 0x4f, 0xb8, 0xf0, 0xa0, 0xe8, 0x23, 0x78, 0x53, 0x76, 0xa7, 0xfa, 0xa9, 0x44, 0xe6,
	0xfa, 0xa9, 0x44, 0x76, 0xac, 0x1f, 0x63, 0xa7, 0xea, 0xc7, 0x18, 0xea, 0xfa, 0x31, 0xd1, 0xd3,
	0x8a, 0xd1, 0x8f, 0xfa, 0x1e, 0x99, 0x70, 0xf1, 0x98, 0x2e, 0x38, 0xf9, 0x0a, 0xb0, 0x03, 0x09,
	0x89, 0x0e, 0x36, 0xcd, 0xeb, 0xf3, 0x68, 0xff, 0x87, 0xde, 0xd3, 0x89, 0x8f, 0xe9, 0xc4, 0x75,
	0x9d, 0xf8, 0x9f, 0x3a, 0xd7, 0xb5, 0x0d, 0xe4, 0x64, 0x2e, 0x21, 0xd8, 0x62, 0xe4, 0x22, 0x7a,
	0xba, 0xa1, 0x8e, 0x35, 0xb3, 0xb7, 0x9c, 0xfe, 0x57, 0x67, 0xde, 0xd6, 0xbb, 0xf6, 0xea, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x25, 0x59, 0x23, 0x33, 0x78, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SearchIPv4ServiceClient is the client API for SearchIPv4Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SearchIPv4ServiceClient interface {
	SearchIPv4(ctx context.Context, in *SearchIPv4Request, opts ...grpc.CallOption) (*SearchResponse, error)
}

type searchIPv4ServiceClient struct {
	cc *grpc.ClientConn
}

func NewSearchIPv4ServiceClient(cc *grpc.ClientConn) SearchIPv4ServiceClient {
	return &searchIPv4ServiceClient{cc}
}

func (c *searchIPv4ServiceClient) SearchIPv4(ctx context.Context, in *SearchIPv4Request, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/SearchIPv4Service/SearchIPv4", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchIPv4ServiceServer is the server API for SearchIPv4Service service.
type SearchIPv4ServiceServer interface {
	SearchIPv4(context.Context, *SearchIPv4Request) (*SearchResponse, error)
}

// UnimplementedSearchIPv4ServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSearchIPv4ServiceServer struct {
}

func (*UnimplementedSearchIPv4ServiceServer) SearchIPv4(ctx context.Context, req *SearchIPv4Request) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchIPv4 not implemented")
}

func RegisterSearchIPv4ServiceServer(s *grpc.Server, srv SearchIPv4ServiceServer) {
	s.RegisterService(&_SearchIPv4Service_serviceDesc, srv)
}

func _SearchIPv4Service_SearchIPv4_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchIPv4Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchIPv4ServiceServer).SearchIPv4(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SearchIPv4Service/SearchIPv4",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchIPv4ServiceServer).SearchIPv4(ctx, req.(*SearchIPv4Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _SearchIPv4Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SearchIPv4Service",
	HandlerType: (*SearchIPv4ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchIPv4",
			Handler:    _SearchIPv4Service_SearchIPv4_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "msg.proto",
}

// SearchIPv6ServiceClient is the client API for SearchIPv6Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SearchIPv6ServiceClient interface {
	SearchIPv6(ctx context.Context, in *SearchIPv6Request, opts ...grpc.CallOption) (*SearchResponse, error)
}

type searchIPv6ServiceClient struct {
	cc *grpc.ClientConn
}

func NewSearchIPv6ServiceClient(cc *grpc.ClientConn) SearchIPv6ServiceClient {
	return &searchIPv6ServiceClient{cc}
}

func (c *searchIPv6ServiceClient) SearchIPv6(ctx context.Context, in *SearchIPv6Request, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/SearchIPv6Service/SearchIPv6", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchIPv6ServiceServer is the server API for SearchIPv6Service service.
type SearchIPv6ServiceServer interface {
	SearchIPv6(context.Context, *SearchIPv6Request) (*SearchResponse, error)
}

// UnimplementedSearchIPv6ServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSearchIPv6ServiceServer struct {
}

func (*UnimplementedSearchIPv6ServiceServer) SearchIPv6(ctx context.Context, req *SearchIPv6Request) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchIPv6 not implemented")
}

func RegisterSearchIPv6ServiceServer(s *grpc.Server, srv SearchIPv6ServiceServer) {
	s.RegisterService(&_SearchIPv6Service_serviceDesc, srv)
}

func _SearchIPv6Service_SearchIPv6_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchIPv6Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchIPv6ServiceServer).SearchIPv6(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SearchIPv6Service/SearchIPv6",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchIPv6ServiceServer).SearchIPv6(ctx, req.(*SearchIPv6Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _SearchIPv6Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SearchIPv6Service",
	HandlerType: (*SearchIPv6ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchIPv6",
			Handler:    _SearchIPv6Service_SearchIPv6_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "msg.proto",
}

// SearchURLServiceClient is the client API for SearchURLService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SearchURLServiceClient interface {
	SearchURL(ctx context.Context, in *SearchURLRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type searchURLServiceClient struct {
	cc *grpc.ClientConn
}

func NewSearchURLServiceClient(cc *grpc.ClientConn) SearchURLServiceClient {
	return &searchURLServiceClient{cc}
}

func (c *searchURLServiceClient) SearchURL(ctx context.Context, in *SearchURLRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/SearchURLService/SearchURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchURLServiceServer is the server API for SearchURLService service.
type SearchURLServiceServer interface {
	SearchURL(context.Context, *SearchURLRequest) (*SearchResponse, error)
}

// UnimplementedSearchURLServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSearchURLServiceServer struct {
}

func (*UnimplementedSearchURLServiceServer) SearchURL(ctx context.Context, req *SearchURLRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchURL not implemented")
}

func RegisterSearchURLServiceServer(s *grpc.Server, srv SearchURLServiceServer) {
	s.RegisterService(&_SearchURLService_serviceDesc, srv)
}

func _SearchURLService_SearchURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchURLServiceServer).SearchURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SearchURLService/SearchURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchURLServiceServer).SearchURL(ctx, req.(*SearchURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SearchURLService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SearchURLService",
	HandlerType: (*SearchURLServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchURL",
			Handler:    _SearchURLService_SearchURL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "msg.proto",
}

// SearchDomainServiceClient is the client API for SearchDomainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SearchDomainServiceClient interface {
	SearchURL(ctx context.Context, in *SearchURLRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type searchDomainServiceClient struct {
	cc *grpc.ClientConn
}

func NewSearchDomainServiceClient(cc *grpc.ClientConn) SearchDomainServiceClient {
	return &searchDomainServiceClient{cc}
}

func (c *searchDomainServiceClient) SearchURL(ctx context.Context, in *SearchURLRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/SearchDomainService/SearchURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchDomainServiceServer is the server API for SearchDomainService service.
type SearchDomainServiceServer interface {
	SearchURL(context.Context, *SearchURLRequest) (*SearchResponse, error)
}

// UnimplementedSearchDomainServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSearchDomainServiceServer struct {
}

func (*UnimplementedSearchDomainServiceServer) SearchURL(ctx context.Context, req *SearchURLRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchURL not implemented")
}

func RegisterSearchDomainServiceServer(s *grpc.Server, srv SearchDomainServiceServer) {
	s.RegisterService(&_SearchDomainService_serviceDesc, srv)
}

func _SearchDomainService_SearchURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchDomainServiceServer).SearchURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SearchDomainService/SearchURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchDomainServiceServer).SearchURL(ctx, req.(*SearchURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SearchDomainService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SearchDomainService",
	HandlerType: (*SearchDomainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchURL",
			Handler:    _SearchDomainService_SearchURL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "msg.proto",
}
