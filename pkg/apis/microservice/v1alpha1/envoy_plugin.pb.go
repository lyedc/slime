// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/apis/microservice/v1alpha1/envoy_plugin.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

// `WorkloadSelector` specifies the criteria used to determine if the
// `Gateway`, `Sidecar`, or `EnvoyFilter` or `ServiceEntry`
// configuration can be applied to a proxy. The matching criteria
// includes the metadata associated with a proxy, workload instance
// info such as labels attached to the pod/VM, or any other info that
// the proxy provides to Istio during the initial handshake. If
// multiple conditions are specified, all conditions need to match in
// order for the workload instance to be selected. Currently, only
// label based selection mechanism is supported.
type WorkloadSelector struct {
	// One or more labels that indicate a specific set of pods/VMs
	// on which the configuration should be applied. The scope of
	// label search is restricted to the configuration namespace in which the
	// the resource is present.
	Labels               map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *WorkloadSelector) Reset()         { *m = WorkloadSelector{} }
func (m *WorkloadSelector) String() string { return proto.CompactTextString(m) }
func (*WorkloadSelector) ProtoMessage()    {}
func (*WorkloadSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_35868063e6636962, []int{0}
}
func (m *WorkloadSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkloadSelector.Unmarshal(m, b)
}
func (m *WorkloadSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkloadSelector.Marshal(b, m, deterministic)
}
func (m *WorkloadSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkloadSelector.Merge(m, src)
}
func (m *WorkloadSelector) XXX_Size() int {
	return xxx_messageInfo_WorkloadSelector.Size(m)
}
func (m *WorkloadSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkloadSelector.DiscardUnknown(m)
}

var xxx_messageInfo_WorkloadSelector proto.InternalMessageInfo

func (m *WorkloadSelector) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type EnvoyPlugin struct {
	WorkloadSelector *WorkloadSelector `protobuf:"bytes,9,opt,name=workload_selector,json=workloadSelector,proto3" json:"workload_selector,omitempty"`
	// route level plugin
	Route []string `protobuf:"bytes,1,rep,name=route,proto3" json:"route,omitempty"`
	// host level plugin
	Host []string `protobuf:"bytes,2,rep,name=host,proto3" json:"host,omitempty"`
	// service level plugin
	Service []string  `protobuf:"bytes,3,rep,name=service,proto3" json:"service,omitempty"`
	Plugins []*Plugin `protobuf:"bytes,4,rep,name=plugins,proto3" json:"plugins,omitempty"`
	// which gateway should use this plugin setting
	Gateway []string `protobuf:"bytes,5,rep,name=gateway,proto3" json:"gateway,omitempty"`
	// which user should use this plugin setting
	User []string `protobuf:"bytes,6,rep,name=user,proto3" json:"user,omitempty"`
	// group setting 用于路由组级别的配置设置，其优先级低于路由级别的配置
	IsGroupSetting bool `protobuf:"varint,7,opt,name=isGroupSetting,proto3" json:"isGroupSetting,omitempty"`
	// listener level
	Listener             []*EnvoyPlugin_Listener `protobuf:"bytes,8,rep,name=listener,proto3" json:"listener,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *EnvoyPlugin) Reset()         { *m = EnvoyPlugin{} }
func (m *EnvoyPlugin) String() string { return proto.CompactTextString(m) }
func (*EnvoyPlugin) ProtoMessage()    {}
func (*EnvoyPlugin) Descriptor() ([]byte, []int) {
	return fileDescriptor_35868063e6636962, []int{1}
}
func (m *EnvoyPlugin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnvoyPlugin.Unmarshal(m, b)
}
func (m *EnvoyPlugin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnvoyPlugin.Marshal(b, m, deterministic)
}
func (m *EnvoyPlugin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnvoyPlugin.Merge(m, src)
}
func (m *EnvoyPlugin) XXX_Size() int {
	return xxx_messageInfo_EnvoyPlugin.Size(m)
}
func (m *EnvoyPlugin) XXX_DiscardUnknown() {
	xxx_messageInfo_EnvoyPlugin.DiscardUnknown(m)
}

var xxx_messageInfo_EnvoyPlugin proto.InternalMessageInfo

func (m *EnvoyPlugin) GetWorkloadSelector() *WorkloadSelector {
	if m != nil {
		return m.WorkloadSelector
	}
	return nil
}

func (m *EnvoyPlugin) GetRoute() []string {
	if m != nil {
		return m.Route
	}
	return nil
}

func (m *EnvoyPlugin) GetHost() []string {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *EnvoyPlugin) GetService() []string {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *EnvoyPlugin) GetPlugins() []*Plugin {
	if m != nil {
		return m.Plugins
	}
	return nil
}

func (m *EnvoyPlugin) GetGateway() []string {
	if m != nil {
		return m.Gateway
	}
	return nil
}

func (m *EnvoyPlugin) GetUser() []string {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *EnvoyPlugin) GetIsGroupSetting() bool {
	if m != nil {
		return m.IsGroupSetting
	}
	return false
}

func (m *EnvoyPlugin) GetListener() []*EnvoyPlugin_Listener {
	if m != nil {
		return m.Listener
	}
	return nil
}

type EnvoyPlugin_Listener struct {
	Port                 uint32   `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	Outbound             bool     `protobuf:"varint,2,opt,name=outbound,proto3" json:"outbound,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnvoyPlugin_Listener) Reset()         { *m = EnvoyPlugin_Listener{} }
func (m *EnvoyPlugin_Listener) String() string { return proto.CompactTextString(m) }
func (*EnvoyPlugin_Listener) ProtoMessage()    {}
func (*EnvoyPlugin_Listener) Descriptor() ([]byte, []int) {
	return fileDescriptor_35868063e6636962, []int{1, 0}
}
func (m *EnvoyPlugin_Listener) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnvoyPlugin_Listener.Unmarshal(m, b)
}
func (m *EnvoyPlugin_Listener) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnvoyPlugin_Listener.Marshal(b, m, deterministic)
}
func (m *EnvoyPlugin_Listener) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnvoyPlugin_Listener.Merge(m, src)
}
func (m *EnvoyPlugin_Listener) XXX_Size() int {
	return xxx_messageInfo_EnvoyPlugin_Listener.Size(m)
}
func (m *EnvoyPlugin_Listener) XXX_DiscardUnknown() {
	xxx_messageInfo_EnvoyPlugin_Listener.DiscardUnknown(m)
}

var xxx_messageInfo_EnvoyPlugin_Listener proto.InternalMessageInfo

func (m *EnvoyPlugin_Listener) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *EnvoyPlugin_Listener) GetOutbound() bool {
	if m != nil {
		return m.Outbound
	}
	return false
}

func init() {
	proto.RegisterType((*WorkloadSelector)(nil), "netease.microservice.v1alpha1.WorkloadSelector")
	proto.RegisterMapType((map[string]string)(nil), "netease.microservice.v1alpha1.WorkloadSelector.LabelsEntry")
	proto.RegisterType((*EnvoyPlugin)(nil), "netease.microservice.v1alpha1.EnvoyPlugin")
	proto.RegisterType((*EnvoyPlugin_Listener)(nil), "netease.microservice.v1alpha1.EnvoyPlugin.Listener")
}

func init() {
	proto.RegisterFile("pkg/apis/microservice/v1alpha1/envoy_plugin.proto", fileDescriptor_35868063e6636962)
}

var fileDescriptor_35868063e6636962 = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xdd, 0x8a, 0x1a, 0x31,
	0x14, 0x66, 0x74, 0xd5, 0xf1, 0x48, 0x8b, 0x0d, 0x7b, 0x11, 0x84, 0x82, 0x2c, 0xb4, 0x78, 0x35,
	0x83, 0x6b, 0x29, 0xed, 0xf6, 0xa2, 0x50, 0x58, 0x7a, 0xb3, 0xd0, 0x12, 0x2f, 0x0a, 0xa5, 0x20,
	0xd1, 0x3d, 0xcc, 0x06, 0x63, 0x32, 0xe4, 0x47, 0x99, 0xd7, 0xe9, 0x2b, 0xf5, 0x85, 0xca, 0x24,
	0x33, 0x22, 0x5e, 0xac, 0xec, 0xdd, 0xf9, 0x3e, 0xf8, 0x7e, 0xce, 0x49, 0x60, 0x5e, 0x6e, 0x8b,
	0x9c, 0x97, 0xc2, 0xe6, 0x3b, 0xb1, 0x31, 0xda, 0xa2, 0xd9, 0x8b, 0x0d, 0xe6, 0xfb, 0x39, 0x97,
	0xe5, 0x13, 0x9f, 0xe7, 0xa8, 0xf6, 0xba, 0x5a, 0x95, 0xd2, 0x17, 0x42, 0x65, 0xa5, 0xd1, 0x4e,
	0x93, 0xb7, 0x0a, 0x1d, 0x72, 0x8b, 0xd9, 0xa9, 0x22, 0x6b, 0x15, 0x93, 0xc5, 0x05, 0xc7, 0xe8,
	0xb5, 0xda, 0x71, 0xc5, 0x0b, 0x34, 0xd1, 0xf3, 0xe6, 0x6f, 0x02, 0xe3, 0x5f, 0xda, 0x6c, 0xa5,
	0xe6, 0x8f, 0x4b, 0x94, 0xb8, 0x71, 0xda, 0x90, 0x25, 0xf4, 0x25, 0x5f, 0xa3, 0xb4, 0x34, 0x99,
	0x76, 0x67, 0xa3, 0xdb, 0x2f, 0xd9, 0xb3, 0xc9, 0xd9, 0xb9, 0x41, 0xf6, 0x10, 0xd4, 0xf7, 0xca,
	0x99, 0x8a, 0x35, 0x56, 0x93, 0xcf, 0x30, 0x3a, 0xa1, 0xc9, 0x18, 0xba, 0x5b, 0xac, 0x68, 0x32,
	0x4d, 0x66, 0x43, 0x56, 0x8f, 0xe4, 0x1a, 0x7a, 0x7b, 0x2e, 0x3d, 0xd2, 0x4e, 0xe0, 0x22, 0xb8,
	0xeb, 0x7c, 0x4a, 0x6e, 0xfe, 0x75, 0x61, 0x74, 0x5f, 0xdf, 0xe3, 0x67, 0x58, 0x81, 0xfc, 0x81,
	0x37, 0x87, 0x26, 0x72, 0x65, 0x9b, 0x4c, 0x3a, 0x9c, 0x26, 0xb3, 0xd1, 0x6d, 0xfe, 0xc2, 0xaa,
	0x6c, 0x7c, 0x38, 0xdf, 0xfe, 0x1a, 0x7a, 0x46, 0x7b, 0x87, 0x61, 0xf9, 0x21, 0x8b, 0x80, 0x10,
	0xb8, 0x7a, 0xd2, 0xd6, 0xd1, 0x4e, 0x20, 0xc3, 0x4c, 0x28, 0x0c, 0x9a, 0x00, 0xda, 0x0d, 0x74,
	0x0b, 0xc9, 0x57, 0x18, 0xc4, 0x73, 0x5b, 0x7a, 0x15, 0x4e, 0xf8, 0xee, 0x42, 0xaf, 0xb8, 0x19,
	0x6b, 0x55, 0xb5, 0x75, 0xc1, 0x1d, 0x1e, 0x78, 0x45, 0x7b, 0xd1, 0xba, 0x81, 0x75, 0x11, 0x6f,
	0xd1, 0xd0, 0x7e, 0x2c, 0x52, 0xcf, 0xe4, 0x3d, 0xbc, 0x16, 0xf6, 0xbb, 0xd1, 0xbe, 0x5c, 0xa2,
	0x73, 0x42, 0x15, 0x74, 0x30, 0x4d, 0x66, 0x29, 0x3b, 0x63, 0xc9, 0x0f, 0x48, 0xa5, 0xb0, 0x0e,
	0x15, 0x1a, 0x9a, 0x86, 0x5e, 0x8b, 0x0b, 0xbd, 0x4e, 0xce, 0x9e, 0x3d, 0x34, 0x52, 0x76, 0x34,
	0x99, 0xdc, 0x41, 0xda, 0xb2, 0x75, 0xb1, 0x52, 0x1b, 0x17, 0x9e, 0xf4, 0x15, 0x0b, 0x33, 0x99,
	0x40, 0xaa, 0xbd, 0x5b, 0x6b, 0xaf, 0x1e, 0xc3, 0xb3, 0xa6, 0xec, 0x88, 0xbf, 0x7d, 0xfc, 0xfd,
	0xa1, 0xf2, 0xea, 0x98, 0xbf, 0xd1, 0xbb, 0xdc, 0x4a, 0xb1, 0xc3, 0xfc, 0xf9, 0x7f, 0xbc, 0xee,
	0x87, 0x9f, 0xbb, 0xf8, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x64, 0xc5, 0x3b, 0xf7, 0x42, 0x03, 0x00,
	0x00,
}