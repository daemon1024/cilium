// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/matcher/v3/string.proto

package envoy_type_matcher_v3

import (
	fmt "fmt"
	_ "github.com/cilium/proxy/go/envoy/annotations"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

// Specifies the way to match a string.
// [#next-free-field: 7]
type StringMatcher struct {
	// Types that are valid to be assigned to MatchPattern:
	//	*StringMatcher_Exact
	//	*StringMatcher_Prefix
	//	*StringMatcher_Suffix
	//	*StringMatcher_SafeRegex
	MatchPattern isStringMatcher_MatchPattern `protobuf_oneof:"match_pattern"`
	// If true, indicates the exact/prefix/suffix matching should be case insensitive. This has no
	// effect for the safe_regex match.
	// For example, the matcher *data* will match both input string *Data* and *data* if set to true.
	IgnoreCase           bool     `protobuf:"varint,6,opt,name=ignore_case,json=ignoreCase,proto3" json:"ignore_case,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringMatcher) Reset()         { *m = StringMatcher{} }
func (m *StringMatcher) String() string { return proto.CompactTextString(m) }
func (*StringMatcher) ProtoMessage()    {}
func (*StringMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_e33cffa01bf36e0e, []int{0}
}

func (m *StringMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringMatcher.Unmarshal(m, b)
}
func (m *StringMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringMatcher.Marshal(b, m, deterministic)
}
func (m *StringMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringMatcher.Merge(m, src)
}
func (m *StringMatcher) XXX_Size() int {
	return xxx_messageInfo_StringMatcher.Size(m)
}
func (m *StringMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_StringMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_StringMatcher proto.InternalMessageInfo

type isStringMatcher_MatchPattern interface {
	isStringMatcher_MatchPattern()
}

type StringMatcher_Exact struct {
	Exact string `protobuf:"bytes,1,opt,name=exact,proto3,oneof"`
}

type StringMatcher_Prefix struct {
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3,oneof"`
}

type StringMatcher_Suffix struct {
	Suffix string `protobuf:"bytes,3,opt,name=suffix,proto3,oneof"`
}

type StringMatcher_SafeRegex struct {
	SafeRegex *RegexMatcher `protobuf:"bytes,5,opt,name=safe_regex,json=safeRegex,proto3,oneof"`
}

func (*StringMatcher_Exact) isStringMatcher_MatchPattern() {}

func (*StringMatcher_Prefix) isStringMatcher_MatchPattern() {}

func (*StringMatcher_Suffix) isStringMatcher_MatchPattern() {}

func (*StringMatcher_SafeRegex) isStringMatcher_MatchPattern() {}

func (m *StringMatcher) GetMatchPattern() isStringMatcher_MatchPattern {
	if m != nil {
		return m.MatchPattern
	}
	return nil
}

func (m *StringMatcher) GetExact() string {
	if x, ok := m.GetMatchPattern().(*StringMatcher_Exact); ok {
		return x.Exact
	}
	return ""
}

func (m *StringMatcher) GetPrefix() string {
	if x, ok := m.GetMatchPattern().(*StringMatcher_Prefix); ok {
		return x.Prefix
	}
	return ""
}

func (m *StringMatcher) GetSuffix() string {
	if x, ok := m.GetMatchPattern().(*StringMatcher_Suffix); ok {
		return x.Suffix
	}
	return ""
}

func (m *StringMatcher) GetSafeRegex() *RegexMatcher {
	if x, ok := m.GetMatchPattern().(*StringMatcher_SafeRegex); ok {
		return x.SafeRegex
	}
	return nil
}

func (m *StringMatcher) GetIgnoreCase() bool {
	if m != nil {
		return m.IgnoreCase
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StringMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StringMatcher_Exact)(nil),
		(*StringMatcher_Prefix)(nil),
		(*StringMatcher_Suffix)(nil),
		(*StringMatcher_SafeRegex)(nil),
	}
}

// Specifies a list of ways to match a string.
type ListStringMatcher struct {
	Patterns             []*StringMatcher `protobuf:"bytes,1,rep,name=patterns,proto3" json:"patterns,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ListStringMatcher) Reset()         { *m = ListStringMatcher{} }
func (m *ListStringMatcher) String() string { return proto.CompactTextString(m) }
func (*ListStringMatcher) ProtoMessage()    {}
func (*ListStringMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_e33cffa01bf36e0e, []int{1}
}

func (m *ListStringMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListStringMatcher.Unmarshal(m, b)
}
func (m *ListStringMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListStringMatcher.Marshal(b, m, deterministic)
}
func (m *ListStringMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListStringMatcher.Merge(m, src)
}
func (m *ListStringMatcher) XXX_Size() int {
	return xxx_messageInfo_ListStringMatcher.Size(m)
}
func (m *ListStringMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_ListStringMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_ListStringMatcher proto.InternalMessageInfo

func (m *ListStringMatcher) GetPatterns() []*StringMatcher {
	if m != nil {
		return m.Patterns
	}
	return nil
}

func init() {
	proto.RegisterType((*StringMatcher)(nil), "envoy.type.matcher.v3.StringMatcher")
	proto.RegisterType((*ListStringMatcher)(nil), "envoy.type.matcher.v3.ListStringMatcher")
}

func init() { proto.RegisterFile("envoy/type/matcher/v3/string.proto", fileDescriptor_e33cffa01bf36e0e) }

var fileDescriptor_e33cffa01bf36e0e = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0xc6, 0x77, 0xd2, 0xa6, 0x66, 0x27, 0x2c, 0xd4, 0xe0, 0x9f, 0x50, 0x50, 0xd3, 0x74, 0xd1,
	0x80, 0x90, 0xc0, 0xf6, 0xb6, 0x78, 0x1a, 0x2f, 0xcb, 0xa2, 0xb2, 0xc4, 0x0f, 0x50, 0xc6, 0xf4,
	0x6d, 0x1d, 0xd0, 0x99, 0x30, 0x33, 0x0d, 0xe9, 0xcd, 0xa3, 0x88, 0x27, 0x8f, 0x7e, 0x04, 0x3f,
	0x82, 0x77, 0xc1, 0xab, 0xdf, 0x46, 0xf6, 0x24, 0x33, 0x13, 0x17, 0x42, 0xb3, 0xb7, 0x4c, 0x9e,
	0xdf, 0x33, 0xcf, 0xf3, 0xbe, 0x09, 0x4e, 0x81, 0x37, 0x62, 0x5f, 0xe8, 0x7d, 0x0d, 0xc5, 0x47,
	0xaa, 0xab, 0xf7, 0x20, 0x8b, 0x66, 0x59, 0x28, 0x2d, 0x19, 0xdf, 0xe6, 0xb5, 0x14, 0x5a, 0x44,
	0xf7, 0x2d, 0x93, 0x1b, 0x26, 0xef, 0x98, 0xbc, 0x59, 0xce, 0xe6, 0xc3, 0x56, 0x09, 0x5b, 0x68,
	0x9d, 0x73, 0xb6, 0x70, 0x08, 0xe5, 0x5c, 0x68, 0xaa, 0x99, 0xe0, 0xaa, 0x58, 0x43, 0x2d, 0xa1,
	0xb2, 0x87, 0x0e, 0x7a, 0xb4, 0x5b, 0xd7, 0xb4, 0xc7, 0x28, 0x4d, 0xf5, 0x4e, 0x75, 0xf2, 0xfc,
	0x40, 0x6e, 0x40, 0x2a, 0x26, 0xf8, 0x4d, 0xc1, 0xd9, 0xc3, 0x86, 0x7e, 0x60, 0x6b, 0xaa, 0xa1,
	0xf8, 0xff, 0xe0, 0x84, 0xf4, 0x87, 0x87, 0x4f, 0xde, 0xda, 0x51, 0x5e, 0xbb, 0x82, 0xd1, 0x03,
	0xec, 0x43, 0x4b, 0x2b, 0x1d, 0xa3, 0x04, 0x65, 0xc7, 0x17, 0x47, 0xa5, 0x3b, 0x46, 0x73, 0x3c,
	0xa9, 0x25, 0x6c, 0x58, 0x1b, 0x7b, 0x46, 0x20, 0x77, 0xae, 0xc9, 0x58, 0x7a, 0x09, 0xba, 0x38,
	0x2a, 0x3b, 0xc1, 0x20, 0x6a, 0xb7, 0x31, 0xc8, 0xe8, 0x00, 0x71, 0x42, 0xf4, 0x06, 0x63, 0x45,
	0x37, 0xb0, 0xb2, 0x3b, 0x88, 0xfd, 0x04, 0x65, 0xe1, 0xd9, 0x22, 0x1f, 0x5c, 0x5f, 0x5e, 0x1a,
	0xa6, 0xab, 0x45, 0x82, 0x6b, 0xe2, 0x7f, 0x41, 0xde, 0xd4, 0x5c, 0x76, 0x6c, 0xae, 0xb0, 0x6a,
	0xf4, 0x04, 0x87, 0x6c, 0xcb, 0x85, 0x84, 0x55, 0x45, 0x15, 0xc4, 0x93, 0x04, 0x65, 0x41, 0x89,
	0xdd, 0xab, 0x97, 0x54, 0xc1, 0xf9, 0xb3, 0xef, 0xbf, 0x3e, 0x3f, 0x4e, 0x71, 0x32, 0x10, 0xd1,
	0x9b, 0x9b, 0xdc, 0xc3, 0x27, 0x56, 0x58, 0xd5, 0x54, 0x6b, 0x90, 0x3c, 0x1a, 0xfd, 0x25, 0xe8,
	0x72, 0x1c, 0x8c, 0xa7, 0x7e, 0xe9, 0xdb, 0xba, 0xe9, 0x57, 0x84, 0xef, 0xbe, 0x62, 0x4a, 0xf7,
	0x17, 0x76, 0x89, 0x83, 0xce, 0xa2, 0x62, 0x94, 0x8c, 0xb2, 0xf0, 0xec, 0xf4, 0x96, 0x81, 0xfa,
	0x81, 0x66, 0xa2, 0x6f, 0xc8, 0x0b, 0x50, 0x79, 0xe3, 0x3f, 0x7f, 0x6e, 0xda, 0x3e, 0xc5, 0xa7,
	0x03, 0xfe, 0x83, 0x60, 0xf2, 0xe2, 0xe7, 0xa7, 0xdf, 0x7f, 0x26, 0xde, 0xd4, 0xc3, 0x0b, 0x26,
	0x5c, 0x64, 0x2d, 0x45, 0xbb, 0x1f, 0x4e, 0x27, 0xa1, 0x73, 0x5f, 0x99, 0xef, 0x7e, 0x85, 0xde,
	0x4d, 0xec, 0x0f, 0xb0, 0xfc, 0x17, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x08, 0xa0, 0x31, 0xe0, 0x02,
	0x00, 0x00,
}