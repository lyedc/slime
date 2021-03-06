// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/apis/microservice/v1alpha1/plugin_manager.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Plugin_ListenerType int32

const (
	Plugin_Outbound Plugin_ListenerType = 0
	Plugin_Inbound  Plugin_ListenerType = 1
)

var Plugin_ListenerType_name = map[int32]string{
	0: "Outbound",
	1: "Inbound",
}

var Plugin_ListenerType_value = map[string]int32{
	"Outbound": 0,
	"Inbound":  1,
}

func (x Plugin_ListenerType) String() string {
	return proto.EnumName(Plugin_ListenerType_name, int32(x))
}

func (Plugin_ListenerType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3b828c45d9f521e4, []int{1, 0}
}

type PluginManager struct {
	// Zero or more labels that indicate a specific set of pods/VMs whose
	// proxies should be configured to use these additional filters.  The
	// scope of label search is platform dependent. On Kubernetes, for
	// example, the scope includes pods running in all reachable
	// namespaces. Omitting the selector applies the filter to all proxies in
	// the mesh.
	WorkloadLabels map[string]string `protobuf:"bytes,1,rep,name=workload_labels,json=workloadLabels,proto3" json:"workload_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Plugin         []*Plugin         `protobuf:"bytes,2,rep,name=plugin,proto3" json:"plugin,omitempty"`
	// Names of gateways where the rule should be applied to. Gateway names
	// at the top of the VirtualService (if any) are overridden. The gateway
	// match is independent of sourceLabels.
	Gateways             []string `protobuf:"bytes,3,rep,name=gateways,proto3" json:"gateways,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginManager) Reset()         { *m = PluginManager{} }
func (m *PluginManager) String() string { return proto.CompactTextString(m) }
func (*PluginManager) ProtoMessage()    {}
func (*PluginManager) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b828c45d9f521e4, []int{0}
}
func (m *PluginManager) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginManager.Unmarshal(m, b)
}
func (m *PluginManager) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginManager.Marshal(b, m, deterministic)
}
func (m *PluginManager) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginManager.Merge(m, src)
}
func (m *PluginManager) XXX_Size() int {
	return xxx_messageInfo_PluginManager.Size(m)
}
func (m *PluginManager) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginManager.DiscardUnknown(m)
}

var xxx_messageInfo_PluginManager proto.InternalMessageInfo

func (m *PluginManager) GetWorkloadLabels() map[string]string {
	if m != nil {
		return m.WorkloadLabels
	}
	return nil
}

func (m *PluginManager) GetPlugin() []*Plugin {
	if m != nil {
		return m.Plugin
	}
	return nil
}

func (m *PluginManager) GetGateways() []string {
	if m != nil {
		return m.Gateways
	}
	return nil
}

type Plugin struct {
	Enable       bool                `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	Name         string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Settings     *types.Struct       `protobuf:"bytes,3,opt,name=settings,proto3" json:"settings,omitempty"`
	ListenerType Plugin_ListenerType `protobuf:"varint,4,opt,name=listenerType,proto3,enum=netease.microservice.v1alpha1.Plugin_ListenerType" json:"listenerType,omitempty"`
	TypeUrl      string              `protobuf:"bytes,5,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// Types that are valid to be assigned to PluginSettings:
	//	*Plugin_Wasm
	//	*Plugin_Inline
	PluginSettings       isPlugin_PluginSettings `protobuf_oneof:"plugin_settings"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Plugin) Reset()         { *m = Plugin{} }
func (m *Plugin) String() string { return proto.CompactTextString(m) }
func (*Plugin) ProtoMessage()    {}
func (*Plugin) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b828c45d9f521e4, []int{1}
}
func (m *Plugin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Plugin.Unmarshal(m, b)
}
func (m *Plugin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Plugin.Marshal(b, m, deterministic)
}
func (m *Plugin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Plugin.Merge(m, src)
}
func (m *Plugin) XXX_Size() int {
	return xxx_messageInfo_Plugin.Size(m)
}
func (m *Plugin) XXX_DiscardUnknown() {
	xxx_messageInfo_Plugin.DiscardUnknown(m)
}

var xxx_messageInfo_Plugin proto.InternalMessageInfo

type isPlugin_PluginSettings interface {
	isPlugin_PluginSettings()
}

type Plugin_Wasm struct {
	Wasm *Wasm `protobuf:"bytes,6,opt,name=wasm,proto3,oneof"`
}
type Plugin_Inline struct {
	Inline *Inline `protobuf:"bytes,7,opt,name=inline,proto3,oneof"`
}

func (*Plugin_Wasm) isPlugin_PluginSettings()   {}
func (*Plugin_Inline) isPlugin_PluginSettings() {}

func (m *Plugin) GetPluginSettings() isPlugin_PluginSettings {
	if m != nil {
		return m.PluginSettings
	}
	return nil
}

func (m *Plugin) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *Plugin) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Plugin) GetSettings() *types.Struct {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *Plugin) GetListenerType() Plugin_ListenerType {
	if m != nil {
		return m.ListenerType
	}
	return Plugin_Outbound
}

func (m *Plugin) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *Plugin) GetWasm() *Wasm {
	if x, ok := m.GetPluginSettings().(*Plugin_Wasm); ok {
		return x.Wasm
	}
	return nil
}

func (m *Plugin) GetInline() *Inline {
	if x, ok := m.GetPluginSettings().(*Plugin_Inline); ok {
		return x.Inline
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Plugin) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Plugin_OneofMarshaler, _Plugin_OneofUnmarshaler, _Plugin_OneofSizer, []interface{}{
		(*Plugin_Wasm)(nil),
		(*Plugin_Inline)(nil),
	}
}

func _Plugin_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Plugin)
	// plugin_settings
	switch x := m.PluginSettings.(type) {
	case *Plugin_Wasm:
		_ = b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Wasm); err != nil {
			return err
		}
	case *Plugin_Inline:
		_ = b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Inline); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Plugin.PluginSettings has unexpected type %T", x)
	}
	return nil
}

func _Plugin_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Plugin)
	switch tag {
	case 6: // plugin_settings.wasm
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Wasm)
		err := b.DecodeMessage(msg)
		m.PluginSettings = &Plugin_Wasm{msg}
		return true, err
	case 7: // plugin_settings.inline
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Inline)
		err := b.DecodeMessage(msg)
		m.PluginSettings = &Plugin_Inline{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Plugin_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Plugin)
	// plugin_settings
	switch x := m.PluginSettings.(type) {
	case *Plugin_Wasm:
		s := proto.Size(x.Wasm)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Plugin_Inline:
		s := proto.Size(x.Inline)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Wasm struct {
	RootID               string        `protobuf:"bytes,1,opt,name=rootID,proto3" json:"rootID,omitempty"`
	FileName             string        `protobuf:"bytes,2,opt,name=fileName,proto3" json:"fileName,omitempty"`
	Settings             *types.Struct `protobuf:"bytes,3,opt,name=settings,proto3" json:"settings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Wasm) Reset()         { *m = Wasm{} }
func (m *Wasm) String() string { return proto.CompactTextString(m) }
func (*Wasm) ProtoMessage()    {}
func (*Wasm) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b828c45d9f521e4, []int{2}
}
func (m *Wasm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Wasm.Unmarshal(m, b)
}
func (m *Wasm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Wasm.Marshal(b, m, deterministic)
}
func (m *Wasm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Wasm.Merge(m, src)
}
func (m *Wasm) XXX_Size() int {
	return xxx_messageInfo_Wasm.Size(m)
}
func (m *Wasm) XXX_DiscardUnknown() {
	xxx_messageInfo_Wasm.DiscardUnknown(m)
}

var xxx_messageInfo_Wasm proto.InternalMessageInfo

func (m *Wasm) GetRootID() string {
	if m != nil {
		return m.RootID
	}
	return ""
}

func (m *Wasm) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *Wasm) GetSettings() *types.Struct {
	if m != nil {
		return m.Settings
	}
	return nil
}

type Inline struct {
	Settings             *types.Struct `protobuf:"bytes,1,opt,name=settings,proto3" json:"settings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Inline) Reset()         { *m = Inline{} }
func (m *Inline) String() string { return proto.CompactTextString(m) }
func (*Inline) ProtoMessage()    {}
func (*Inline) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b828c45d9f521e4, []int{3}
}
func (m *Inline) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Inline.Unmarshal(m, b)
}
func (m *Inline) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Inline.Marshal(b, m, deterministic)
}
func (m *Inline) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Inline.Merge(m, src)
}
func (m *Inline) XXX_Size() int {
	return xxx_messageInfo_Inline.Size(m)
}
func (m *Inline) XXX_DiscardUnknown() {
	xxx_messageInfo_Inline.DiscardUnknown(m)
}

var xxx_messageInfo_Inline proto.InternalMessageInfo

func (m *Inline) GetSettings() *types.Struct {
	if m != nil {
		return m.Settings
	}
	return nil
}

func init() {
	proto.RegisterEnum("netease.microservice.v1alpha1.Plugin_ListenerType", Plugin_ListenerType_name, Plugin_ListenerType_value)
	proto.RegisterType((*PluginManager)(nil), "netease.microservice.v1alpha1.PluginManager")
	proto.RegisterMapType((map[string]string)(nil), "netease.microservice.v1alpha1.PluginManager.WorkloadLabelsEntry")
	proto.RegisterType((*Plugin)(nil), "netease.microservice.v1alpha1.Plugin")
	proto.RegisterType((*Wasm)(nil), "netease.microservice.v1alpha1.Wasm")
	proto.RegisterType((*Inline)(nil), "netease.microservice.v1alpha1.Inline")
}

func init() {
	proto.RegisterFile("pkg/apis/microservice/v1alpha1/plugin_manager.proto", fileDescriptor_3b828c45d9f521e4)
}

var fileDescriptor_3b828c45d9f521e4 = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xef, 0x6b, 0xd3, 0x40,
	0x1c, 0xc6, 0x9b, 0xb6, 0x4b, 0xbb, 0x6f, 0xeb, 0x36, 0x4f, 0xd1, 0x58, 0x14, 0x4a, 0x45, 0xa8,
	0x6f, 0x12, 0xd6, 0x8a, 0xa8, 0x30, 0xd4, 0xa1, 0xb0, 0xc2, 0xfc, 0x41, 0xfc, 0x31, 0xf0, 0x4d,
	0xb9, 0x74, 0xdf, 0xc5, 0xa3, 0x97, 0xbb, 0x70, 0x77, 0x69, 0xc9, 0x3f, 0x22, 0xfe, 0xb9, 0x92,
	0x4b, 0x5a, 0x5b, 0x90, 0xad, 0xf8, 0x2e, 0xcf, 0x71, 0x9f, 0xe7, 0xfb, 0xdc, 0x3d, 0x17, 0x18,
	0xa7, 0xf3, 0x38, 0xa0, 0x29, 0xd3, 0x41, 0xc2, 0x66, 0x4a, 0x6a, 0x54, 0x0b, 0x36, 0xc3, 0x60,
	0x71, 0x4c, 0x79, 0xfa, 0x93, 0x1e, 0x07, 0x29, 0xcf, 0x62, 0x26, 0xa6, 0x09, 0x15, 0x34, 0x46,
	0xe5, 0xa7, 0x4a, 0x1a, 0x49, 0x1e, 0x09, 0x34, 0x48, 0x35, 0xfa, 0x9b, 0x8c, 0xbf, 0x62, 0x7a,
	0x0f, 0x63, 0x29, 0x63, 0x8e, 0x81, 0xdd, 0x1c, 0x65, 0x57, 0x81, 0x36, 0x2a, 0x9b, 0x99, 0x12,
	0x1e, 0xfc, 0xae, 0xc3, 0xad, 0xcf, 0xd6, 0xf5, 0x43, 0x69, 0x4a, 0x18, 0x1c, 0x2e, 0xa5, 0x9a,
	0x73, 0x49, 0x2f, 0xa7, 0x9c, 0x46, 0xc8, 0xb5, 0xe7, 0xf4, 0x1b, 0xc3, 0xce, 0xe8, 0x8d, 0x7f,
	0xed, 0x20, 0x7f, 0xcb, 0xc6, 0xbf, 0xa8, 0x3c, 0xce, 0xad, 0xc5, 0x7b, 0x61, 0x54, 0x1e, 0x1e,
	0x2c, 0xb7, 0x16, 0xc9, 0x09, 0xb8, 0xe5, 0x89, 0xbc, 0xba, 0x9d, 0xf0, 0x64, 0xa7, 0x09, 0x61,
	0x05, 0x91, 0x1e, 0xb4, 0x63, 0x6a, 0x70, 0x49, 0x73, 0xed, 0x35, 0xfa, 0x8d, 0xe1, 0x7e, 0xb8,
	0xd6, 0xbd, 0xb7, 0x70, 0xe7, 0x1f, 0x09, 0xc8, 0x11, 0x34, 0xe6, 0x98, 0x7b, 0x4e, 0xdf, 0x19,
	0xee, 0x87, 0xc5, 0x27, 0xb9, 0x0b, 0x7b, 0x0b, 0xca, 0x33, 0xf4, 0xea, 0x76, 0xad, 0x14, 0xaf,
	0xea, 0x2f, 0x9c, 0xc1, 0xaf, 0x06, 0xb8, 0xe5, 0x44, 0x72, 0x0f, 0x5c, 0x14, 0x34, 0xe2, 0x68,
	0xc9, 0x76, 0x58, 0x29, 0x42, 0xa0, 0x29, 0x68, 0xb2, 0x62, 0xed, 0x37, 0x19, 0x43, 0x5b, 0xa3,
	0x31, 0x4c, 0xc4, 0x45, 0x2a, 0x67, 0xd8, 0x19, 0xdd, 0xf7, 0xcb, 0x0a, 0xfc, 0x55, 0x05, 0xfe,
	0x17, 0x5b, 0x41, 0xb8, 0xde, 0x48, 0xbe, 0x43, 0x97, 0x33, 0x6d, 0x50, 0xa0, 0xfa, 0x9a, 0xa7,
	0xe8, 0x35, 0xfb, 0xce, 0xf0, 0x60, 0x34, 0xda, 0xe9, 0x3e, 0xfc, 0xf3, 0x0d, 0x32, 0xdc, 0xf2,
	0x21, 0x0f, 0xa0, 0x6d, 0xf2, 0x14, 0xa7, 0x99, 0xe2, 0xde, 0x9e, 0x0d, 0xd9, 0x2a, 0xf4, 0x37,
	0xc5, 0xc9, 0x4b, 0x68, 0x2e, 0xa9, 0x4e, 0x3c, 0xd7, 0x66, 0x7c, 0x7c, 0xc3, 0xa8, 0x0b, 0xaa,
	0x93, 0xb3, 0x5a, 0x68, 0x11, 0xf2, 0x1a, 0x5c, 0x26, 0x38, 0x13, 0xe8, 0xb5, 0x2c, 0x7c, 0x53,
	0x6f, 0x13, 0xbb, 0xf9, 0xac, 0x16, 0x56, 0xd8, 0xe0, 0x29, 0x74, 0x37, 0x43, 0x93, 0x2e, 0xb4,
	0x3f, 0x65, 0x26, 0x92, 0x99, 0xb8, 0x3c, 0xaa, 0x91, 0x0e, 0xb4, 0x26, 0xa2, 0x14, 0xce, 0xe9,
	0x6d, 0x38, 0xac, 0x5e, 0xfd, 0xea, 0xb2, 0x06, 0x12, 0x9a, 0x45, 0x9c, 0xa2, 0x15, 0x25, 0xa5,
	0x99, 0xbc, 0xab, 0xfa, 0xac, 0x54, 0xf1, 0x2e, 0xae, 0x18, 0xc7, 0x8f, 0x7f, 0x9b, 0x59, 0xeb,
	0xff, 0x6a, 0x67, 0x70, 0x02, 0x6e, 0x79, 0x84, 0x2d, 0xdc, 0xd9, 0x11, 0x3f, 0x7d, 0xfe, 0xe3,
	0x59, 0x9e, 0x89, 0xf5, 0x1d, 0xcd, 0x64, 0x12, 0x68, 0xce, 0x12, 0x0c, 0xae, 0xff, 0xdb, 0x23,
	0xd7, 0x5a, 0x8e, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x68, 0x77, 0xf2, 0xc6, 0x16, 0x04, 0x00,
	0x00,
}
