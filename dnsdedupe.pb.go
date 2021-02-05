// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        (unknown)
// source: dnsdedupe.proto

package nmsg_sie

import (
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

type DnsDedupeType int32

const (
	DnsDedupeType_INSERTION            DnsDedupeType = 0
	DnsDedupeType_EXPIRATION           DnsDedupeType = 1
	DnsDedupeType_CHAFF                DnsDedupeType = 2
	DnsDedupeType_AUTHORITATIVE        DnsDedupeType = 3
	DnsDedupeType_MERGED               DnsDedupeType = 4
	DnsDedupeType_MERGED_AUTHORITATIVE DnsDedupeType = 5
	DnsDedupeType_MERGED_INSERTION     DnsDedupeType = 6
)

// Enum value maps for DnsDedupeType.
var (
	DnsDedupeType_name = map[int32]string{
		0: "INSERTION",
		1: "EXPIRATION",
		2: "CHAFF",
		3: "AUTHORITATIVE",
		4: "MERGED",
		5: "MERGED_AUTHORITATIVE",
		6: "MERGED_INSERTION",
	}
	DnsDedupeType_value = map[string]int32{
		"INSERTION":            0,
		"EXPIRATION":           1,
		"CHAFF":                2,
		"AUTHORITATIVE":        3,
		"MERGED":               4,
		"MERGED_AUTHORITATIVE": 5,
		"MERGED_INSERTION":     6,
	}
)

func (x DnsDedupeType) Enum() *DnsDedupeType {
	p := new(DnsDedupeType)
	*p = x
	return p
}

func (x DnsDedupeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DnsDedupeType) Descriptor() protoreflect.EnumDescriptor {
	return file_dnsdedupe_proto_enumTypes[0].Descriptor()
}

func (DnsDedupeType) Type() protoreflect.EnumType {
	return &file_dnsdedupe_proto_enumTypes[0]
}

func (x DnsDedupeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *DnsDedupeType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = DnsDedupeType(num)
	return nil
}

// Deprecated: Use DnsDedupeType.Descriptor instead.
func (DnsDedupeType) EnumDescriptor() ([]byte, []int) {
	return file_dnsdedupe_proto_rawDescGZIP(), []int{0}
}

type DnsDedupe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type          *DnsDedupeType `protobuf:"varint,13,opt,name=type,enum=nmsg.sie.DnsDedupeType" json:"type,omitempty"`
	Count         *uint32        `protobuf:"varint,10,opt,name=count" json:"count,omitempty"`
	TimeFirst     *uint32        `protobuf:"varint,11,opt,name=time_first,json=timeFirst" json:"time_first,omitempty"`
	TimeLast      *uint32        `protobuf:"varint,12,opt,name=time_last,json=timeLast" json:"time_last,omitempty"`
	ZoneTimeFirst *uint32        `protobuf:"varint,17,opt,name=zone_time_first,json=zoneTimeFirst" json:"zone_time_first,omitempty"`
	ZoneTimeLast  *uint32        `protobuf:"varint,18,opt,name=zone_time_last,json=zoneTimeLast" json:"zone_time_last,omitempty"`
	ResponseIp    []byte         `protobuf:"bytes,14,opt,name=response_ip,json=responseIp" json:"response_ip,omitempty"`
	Rrname        []byte         `protobuf:"bytes,1,opt,name=rrname" json:"rrname,omitempty"`
	Rrtype        *uint32        `protobuf:"varint,2,opt,name=rrtype" json:"rrtype,omitempty"`
	Rrclass       *uint32        `protobuf:"varint,3,opt,name=rrclass" json:"rrclass,omitempty"`
	Rrttl         *uint32        `protobuf:"varint,4,opt,name=rrttl" json:"rrttl,omitempty"`
	Rdata         [][]byte       `protobuf:"bytes,5,rep,name=rdata" json:"rdata,omitempty"`
	Response      []byte         `protobuf:"bytes,15,opt,name=response" json:"response,omitempty"`
	Bailiwick     []byte         `protobuf:"bytes,16,opt,name=bailiwick" json:"bailiwick,omitempty"`
}

func (x *DnsDedupe) Reset() {
	*x = DnsDedupe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dnsdedupe_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsDedupe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsDedupe) ProtoMessage() {}

func (x *DnsDedupe) ProtoReflect() protoreflect.Message {
	mi := &file_dnsdedupe_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsDedupe.ProtoReflect.Descriptor instead.
func (*DnsDedupe) Descriptor() ([]byte, []int) {
	return file_dnsdedupe_proto_rawDescGZIP(), []int{0}
}

func (x *DnsDedupe) GetType() DnsDedupeType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return DnsDedupeType_INSERTION
}

func (x *DnsDedupe) GetCount() uint32 {
	if x != nil && x.Count != nil {
		return *x.Count
	}
	return 0
}

func (x *DnsDedupe) GetTimeFirst() uint32 {
	if x != nil && x.TimeFirst != nil {
		return *x.TimeFirst
	}
	return 0
}

func (x *DnsDedupe) GetTimeLast() uint32 {
	if x != nil && x.TimeLast != nil {
		return *x.TimeLast
	}
	return 0
}

func (x *DnsDedupe) GetZoneTimeFirst() uint32 {
	if x != nil && x.ZoneTimeFirst != nil {
		return *x.ZoneTimeFirst
	}
	return 0
}

func (x *DnsDedupe) GetZoneTimeLast() uint32 {
	if x != nil && x.ZoneTimeLast != nil {
		return *x.ZoneTimeLast
	}
	return 0
}

func (x *DnsDedupe) GetResponseIp() []byte {
	if x != nil {
		return x.ResponseIp
	}
	return nil
}

func (x *DnsDedupe) GetRrname() []byte {
	if x != nil {
		return x.Rrname
	}
	return nil
}

func (x *DnsDedupe) GetRrtype() uint32 {
	if x != nil && x.Rrtype != nil {
		return *x.Rrtype
	}
	return 0
}

func (x *DnsDedupe) GetRrclass() uint32 {
	if x != nil && x.Rrclass != nil {
		return *x.Rrclass
	}
	return 0
}

func (x *DnsDedupe) GetRrttl() uint32 {
	if x != nil && x.Rrttl != nil {
		return *x.Rrttl
	}
	return 0
}

func (x *DnsDedupe) GetRdata() [][]byte {
	if x != nil {
		return x.Rdata
	}
	return nil
}

func (x *DnsDedupe) GetResponse() []byte {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *DnsDedupe) GetBailiwick() []byte {
	if x != nil {
		return x.Bailiwick
	}
	return nil
}

var File_dnsdedupe_proto protoreflect.FileDescriptor

var file_dnsdedupe_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x6e, 0x73, 0x64, 0x65, 0x64, 0x75, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x6e, 0x6d, 0x73, 0x67, 0x2e, 0x73, 0x69, 0x65, 0x22, 0xa9, 0x03, 0x0a, 0x09,
	0x44, 0x6e, 0x73, 0x44, 0x65, 0x64, 0x75, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6e, 0x6d, 0x73, 0x67, 0x2e, 0x73,
	0x69, 0x65, 0x2e, 0x44, 0x6e, 0x73, 0x44, 0x65, 0x64, 0x75, 0x70, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x46, 0x69, 0x72, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x74, 0x69, 0x6d, 0x65, 0x4c, 0x61, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x7a, 0x6f, 0x6e, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0d, 0x7a, 0x6f, 0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x69, 0x72, 0x73, 0x74,
	0x12, 0x24, 0x0a, 0x0e, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6c, 0x61,
	0x73, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x7a, 0x6f, 0x6e, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x4c, 0x61, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x5f, 0x69, 0x70, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x72, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x72, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x72, 0x72, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x72, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x72, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x72, 0x74, 0x74, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x72, 0x72, 0x74, 0x74, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x05, 0x72, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x61, 0x69,
	0x6c, 0x69, 0x77, 0x69, 0x63, 0x6b, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x61,
	0x69, 0x6c, 0x69, 0x77, 0x69, 0x63, 0x6b, 0x2a, 0x88, 0x01, 0x0a, 0x0d, 0x44, 0x6e, 0x73, 0x44,
	0x65, 0x64, 0x75, 0x70, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4e, 0x53,
	0x45, 0x52, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x58, 0x50, 0x49,
	0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x48, 0x41, 0x46,
	0x46, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x54, 0x41,
	0x54, 0x49, 0x56, 0x45, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x45, 0x52, 0x47, 0x45, 0x44,
	0x10, 0x04, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x45, 0x52, 0x47, 0x45, 0x44, 0x5f, 0x41, 0x55, 0x54,
	0x48, 0x4f, 0x52, 0x49, 0x54, 0x41, 0x54, 0x49, 0x56, 0x45, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10,
	0x4d, 0x45, 0x52, 0x47, 0x45, 0x44, 0x5f, 0x49, 0x4e, 0x53, 0x45, 0x52, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x06, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x66, 0x61, 0x72, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x65, 0x63, 0x2f, 0x67, 0x6f, 0x2d,
	0x6e, 0x6d, 0x73, 0x67, 0x5f, 0x73, 0x69, 0x65, 0x3b, 0x6e, 0x6d, 0x73, 0x67, 0x5f, 0x73, 0x69,
	0x65,
}

var (
	file_dnsdedupe_proto_rawDescOnce sync.Once
	file_dnsdedupe_proto_rawDescData = file_dnsdedupe_proto_rawDesc
)

func file_dnsdedupe_proto_rawDescGZIP() []byte {
	file_dnsdedupe_proto_rawDescOnce.Do(func() {
		file_dnsdedupe_proto_rawDescData = protoimpl.X.CompressGZIP(file_dnsdedupe_proto_rawDescData)
	})
	return file_dnsdedupe_proto_rawDescData
}

var file_dnsdedupe_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_dnsdedupe_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_dnsdedupe_proto_goTypes = []interface{}{
	(DnsDedupeType)(0), // 0: nmsg.sie.DnsDedupeType
	(*DnsDedupe)(nil),  // 1: nmsg.sie.DnsDedupe
}
var file_dnsdedupe_proto_depIdxs = []int32{
	0, // 0: nmsg.sie.DnsDedupe.type:type_name -> nmsg.sie.DnsDedupeType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_dnsdedupe_proto_init() }
func file_dnsdedupe_proto_init() {
	if File_dnsdedupe_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dnsdedupe_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsDedupe); i {
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
			RawDescriptor: file_dnsdedupe_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dnsdedupe_proto_goTypes,
		DependencyIndexes: file_dnsdedupe_proto_depIdxs,
		EnumInfos:         file_dnsdedupe_proto_enumTypes,
		MessageInfos:      file_dnsdedupe_proto_msgTypes,
	}.Build()
	File_dnsdedupe_proto = out.File
	file_dnsdedupe_proto_rawDesc = nil
	file_dnsdedupe_proto_goTypes = nil
	file_dnsdedupe_proto_depIdxs = nil
}
