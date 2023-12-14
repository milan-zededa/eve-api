// Copyright(c) 2017-2018 Zededa, Inc.
// All rights reserved.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: config/netinst.proto

package config

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

type ZNetworkInstType int32

const (
	ZNetworkInstType_ZNetInstFirst       ZNetworkInstType = 0
	ZNetworkInstType_ZnetInstSwitch      ZNetworkInstType = 1
	ZNetworkInstType_ZnetInstLocal       ZNetworkInstType = 2
	ZNetworkInstType_ZnetInstCloud       ZNetworkInstType = 3
	ZNetworkInstType_ZnetInstMesh        ZNetworkInstType = 4
	ZNetworkInstType_ZnetInstHoneyPot    ZNetworkInstType = 5
	ZNetworkInstType_ZnetInstTransparent ZNetworkInstType = 6
	ZNetworkInstType_ZNetInstLast        ZNetworkInstType = 255
)

// Enum value maps for ZNetworkInstType.
var (
	ZNetworkInstType_name = map[int32]string{
		0:   "ZNetInstFirst",
		1:   "ZnetInstSwitch",
		2:   "ZnetInstLocal",
		3:   "ZnetInstCloud",
		4:   "ZnetInstMesh",
		5:   "ZnetInstHoneyPot",
		6:   "ZnetInstTransparent",
		255: "ZNetInstLast",
	}
	ZNetworkInstType_value = map[string]int32{
		"ZNetInstFirst":       0,
		"ZnetInstSwitch":      1,
		"ZnetInstLocal":       2,
		"ZnetInstCloud":       3,
		"ZnetInstMesh":        4,
		"ZnetInstHoneyPot":    5,
		"ZnetInstTransparent": 6,
		"ZNetInstLast":        255,
	}
)

func (x ZNetworkInstType) Enum() *ZNetworkInstType {
	p := new(ZNetworkInstType)
	*p = x
	return p
}

func (x ZNetworkInstType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ZNetworkInstType) Descriptor() protoreflect.EnumDescriptor {
	return file_config_netinst_proto_enumTypes[0].Descriptor()
}

func (ZNetworkInstType) Type() protoreflect.EnumType {
	return &file_config_netinst_proto_enumTypes[0]
}

func (x ZNetworkInstType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ZNetworkInstType.Descriptor instead.
func (ZNetworkInstType) EnumDescriptor() ([]byte, []int) {
	return file_config_netinst_proto_rawDescGZIP(), []int{0}
}

type AddressType int32

const (
	AddressType_First      AddressType = 0
	AddressType_IPV4       AddressType = 1
	AddressType_IPV6       AddressType = 2
	AddressType_CryptoIPV4 AddressType = 3
	AddressType_CryptoIPV6 AddressType = 4
	AddressType_Last       AddressType = 255
)

// Enum value maps for AddressType.
var (
	AddressType_name = map[int32]string{
		0:   "First",
		1:   "IPV4",
		2:   "IPV6",
		3:   "CryptoIPV4",
		4:   "CryptoIPV6",
		255: "Last",
	}
	AddressType_value = map[string]int32{
		"First":      0,
		"IPV4":       1,
		"IPV6":       2,
		"CryptoIPV4": 3,
		"CryptoIPV6": 4,
		"Last":       255,
	}
)

func (x AddressType) Enum() *AddressType {
	p := new(AddressType)
	*p = x
	return p
}

func (x AddressType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AddressType) Descriptor() protoreflect.EnumDescriptor {
	return file_config_netinst_proto_enumTypes[1].Descriptor()
}

func (AddressType) Type() protoreflect.EnumType {
	return &file_config_netinst_proto_enumTypes[1]
}

func (x AddressType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AddressType.Descriptor instead.
func (AddressType) EnumDescriptor() ([]byte, []int) {
	return file_config_netinst_proto_rawDescGZIP(), []int{1}
}

type ZNetworkOpaqueConfigType int32

const (
	ZNetworkOpaqueConfigType_ZNetOConfigVPN  ZNetworkOpaqueConfigType = 0
	ZNetworkOpaqueConfigType_ZNetOConfigLisp ZNetworkOpaqueConfigType = 1
)

// Enum value maps for ZNetworkOpaqueConfigType.
var (
	ZNetworkOpaqueConfigType_name = map[int32]string{
		0: "ZNetOConfigVPN",
		1: "ZNetOConfigLisp",
	}
	ZNetworkOpaqueConfigType_value = map[string]int32{
		"ZNetOConfigVPN":  0,
		"ZNetOConfigLisp": 1,
	}
)

func (x ZNetworkOpaqueConfigType) Enum() *ZNetworkOpaqueConfigType {
	p := new(ZNetworkOpaqueConfigType)
	*p = x
	return p
}

func (x ZNetworkOpaqueConfigType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ZNetworkOpaqueConfigType) Descriptor() protoreflect.EnumDescriptor {
	return file_config_netinst_proto_enumTypes[2].Descriptor()
}

func (ZNetworkOpaqueConfigType) Type() protoreflect.EnumType {
	return &file_config_netinst_proto_enumTypes[2]
}

func (x ZNetworkOpaqueConfigType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ZNetworkOpaqueConfigType.Descriptor instead.
func (ZNetworkOpaqueConfigType) EnumDescriptor() ([]byte, []int) {
	return file_config_netinst_proto_rawDescGZIP(), []int{2}
}

type ZcServiceType int32

const (
	ZcServiceType_zcloudInvalidSrv ZcServiceType = 0
	// mapping service for zededa overlay service
	ZcServiceType_mapServer ZcServiceType = 1
	// if device has support feature enabled, this is cloud service from
	// device can be reached
	ZcServiceType_supportServer ZcServiceType = 2
)

// Enum value maps for ZcServiceType.
var (
	ZcServiceType_name = map[int32]string{
		0: "zcloudInvalidSrv",
		1: "mapServer",
		2: "supportServer",
	}
	ZcServiceType_value = map[string]int32{
		"zcloudInvalidSrv": 0,
		"mapServer":        1,
		"supportServer":    2,
	}
)

func (x ZcServiceType) Enum() *ZcServiceType {
	p := new(ZcServiceType)
	*p = x
	return p
}

func (x ZcServiceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ZcServiceType) Descriptor() protoreflect.EnumDescriptor {
	return file_config_netinst_proto_enumTypes[3].Descriptor()
}

func (ZcServiceType) Type() protoreflect.EnumType {
	return &file_config_netinst_proto_enumTypes[3]
}

func (x ZcServiceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ZcServiceType.Descriptor instead.
func (ZcServiceType) EnumDescriptor() ([]byte, []int) {
	return file_config_netinst_proto_rawDescGZIP(), []int{3}
}

// Network Instance Opaque config. In future we might add more fields here
// but idea is here. This is service specific configuration.
type NetworkInstanceOpaqueConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Oconfig    string                     `protobuf:"bytes,1,opt,name=oconfig,proto3" json:"oconfig,omitempty"`
	LispConfig *NetworkInstanceLispConfig `protobuf:"bytes,2,opt,name=lispConfig,proto3" json:"lispConfig,omitempty"`
	Type       ZNetworkOpaqueConfigType   `protobuf:"varint,3,opt,name=type,proto3,enum=org.lfedge.eve.config.ZNetworkOpaqueConfigType" json:"type,omitempty"`
}

func (x *NetworkInstanceOpaqueConfig) Reset() {
	*x = NetworkInstanceOpaqueConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_netinst_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkInstanceOpaqueConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInstanceOpaqueConfig) ProtoMessage() {}

func (x *NetworkInstanceOpaqueConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_netinst_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInstanceOpaqueConfig.ProtoReflect.Descriptor instead.
func (*NetworkInstanceOpaqueConfig) Descriptor() ([]byte, []int) {
	return file_config_netinst_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkInstanceOpaqueConfig) GetOconfig() string {
	if x != nil {
		return x.Oconfig
	}
	return ""
}

func (x *NetworkInstanceOpaqueConfig) GetLispConfig() *NetworkInstanceLispConfig {
	if x != nil {
		return x.LispConfig
	}
	return nil
}

func (x *NetworkInstanceOpaqueConfig) GetType() ZNetworkOpaqueConfigType {
	if x != nil {
		return x.Type
	}
	return ZNetworkOpaqueConfigType_ZNetOConfigVPN
}

// This is way to tell the device if there is service in cloud somewhere,
// what type it is how to access it
type ZcServicePoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ZsType     ZcServiceType `protobuf:"varint,3,opt,name=zsType,proto3,enum=org.lfedge.eve.config.ZcServiceType" json:"zsType,omitempty"`
	NameOrIp   string        `protobuf:"bytes,1,opt,name=NameOrIp,proto3" json:"NameOrIp,omitempty"`
	Credential string        `protobuf:"bytes,2,opt,name=Credential,proto3" json:"Credential,omitempty"`
}

func (x *ZcServicePoint) Reset() {
	*x = ZcServicePoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_netinst_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZcServicePoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZcServicePoint) ProtoMessage() {}

func (x *ZcServicePoint) ProtoReflect() protoreflect.Message {
	mi := &file_config_netinst_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZcServicePoint.ProtoReflect.Descriptor instead.
func (*ZcServicePoint) Descriptor() ([]byte, []int) {
	return file_config_netinst_proto_rawDescGZIP(), []int{1}
}

func (x *ZcServicePoint) GetZsType() ZcServiceType {
	if x != nil {
		return x.ZsType
	}
	return ZcServiceType_zcloudInvalidSrv
}

func (x *ZcServicePoint) GetNameOrIp() string {
	if x != nil {
		return x.NameOrIp
	}
	return ""
}

func (x *ZcServicePoint) GetCredential() string {
	if x != nil {
		return x.Credential
	}
	return ""
}

// Lisp NetworkInstance config
type NetworkInstanceLispConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LispMSs             []*ZcServicePoint `protobuf:"bytes,1,rep,name=LispMSs,proto3" json:"LispMSs,omitempty"`
	LispInstanceId      uint32            `protobuf:"varint,2,opt,name=LispInstanceId,proto3" json:"LispInstanceId,omitempty"`
	Allocate            bool              `protobuf:"varint,3,opt,name=allocate,proto3" json:"allocate,omitempty"`
	Exportprivate       bool              `protobuf:"varint,4,opt,name=exportprivate,proto3" json:"exportprivate,omitempty"`
	Allocationprefix    []byte            `protobuf:"bytes,5,opt,name=allocationprefix,proto3" json:"allocationprefix,omitempty"`
	Allocationprefixlen uint32            `protobuf:"varint,6,opt,name=allocationprefixlen,proto3" json:"allocationprefixlen,omitempty"`
	// various configuration to dataPlane, lispers.net vs Zededa
	Experimental bool `protobuf:"varint,20,opt,name=experimental,proto3" json:"experimental,omitempty"`
}

func (x *NetworkInstanceLispConfig) Reset() {
	*x = NetworkInstanceLispConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_netinst_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkInstanceLispConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInstanceLispConfig) ProtoMessage() {}

func (x *NetworkInstanceLispConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_netinst_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInstanceLispConfig.ProtoReflect.Descriptor instead.
func (*NetworkInstanceLispConfig) Descriptor() ([]byte, []int) {
	return file_config_netinst_proto_rawDescGZIP(), []int{2}
}

func (x *NetworkInstanceLispConfig) GetLispMSs() []*ZcServicePoint {
	if x != nil {
		return x.LispMSs
	}
	return nil
}

func (x *NetworkInstanceLispConfig) GetLispInstanceId() uint32 {
	if x != nil {
		return x.LispInstanceId
	}
	return 0
}

func (x *NetworkInstanceLispConfig) GetAllocate() bool {
	if x != nil {
		return x.Allocate
	}
	return false
}

func (x *NetworkInstanceLispConfig) GetExportprivate() bool {
	if x != nil {
		return x.Exportprivate
	}
	return false
}

func (x *NetworkInstanceLispConfig) GetAllocationprefix() []byte {
	if x != nil {
		return x.Allocationprefix
	}
	return nil
}

func (x *NetworkInstanceLispConfig) GetAllocationprefixlen() uint32 {
	if x != nil {
		return x.Allocationprefixlen
	}
	return 0
}

func (x *NetworkInstanceLispConfig) GetExperimental() bool {
	if x != nil {
		return x.Experimental
	}
	return false
}

type IPRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Destination network address in the CIDR format: <IP-address>/<prefix-length>
	// It is allowed to submit default route with all-zeroes destination network address
	// 0.0.0.0/0 or ::/0.
	DestinationNetwork string `protobuf:"bytes,1,opt,name=destination_network,json=destinationNetwork,proto3" json:"destination_network,omitempty"`
	// Gateway IP address.
	// This must be a valid IP address and can not be all-zeroes.
	Gateway string `protobuf:"bytes,2,opt,name=gateway,proto3" json:"gateway,omitempty"`
}

func (x *IPRoute) Reset() {
	*x = IPRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_netinst_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPRoute) ProtoMessage() {}

func (x *IPRoute) ProtoReflect() protoreflect.Message {
	mi := &file_config_netinst_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPRoute.ProtoReflect.Descriptor instead.
func (*IPRoute) Descriptor() ([]byte, []int) {
	return file_config_netinst_proto_rawDescGZIP(), []int{3}
}

func (x *IPRoute) GetDestinationNetwork() string {
	if x != nil {
		return x.DestinationNetwork
	}
	return ""
}

func (x *IPRoute) GetGateway() string {
	if x != nil {
		return x.Gateway
	}
	return ""
}

type NetworkInstanceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuidandversion *UUIDandVersion `protobuf:"bytes,1,opt,name=uuidandversion,proto3" json:"uuidandversion,omitempty"`
	Displayname    string          `protobuf:"bytes,2,opt,name=displayname,proto3" json:"displayname,omitempty"`
	// instType - Type of network instance ( local, bridge etc )
	InstType ZNetworkInstType `protobuf:"varint,4,opt,name=instType,proto3,enum=org.lfedge.eve.config.ZNetworkInstType" json:"instType,omitempty"`
	// activate
	//   - True by default. If set to false ( deactivate), the network instance
	//     configuration is downloaded to the device, but the network instance
	//     itself is not created on the device.
	Activate bool `protobuf:"varint,5,opt,name=activate,proto3" json:"activate,omitempty"`
	// port - Only a single port is supported.
	//
	//	This is used as the external connection for the network instance.
	//	This can be a physical (eth0 ) or logical port (vlan 0).
	//	The port name comes from DeviceConfig ( When it is supported in future).
	//	If the user needs multiple physical ports, Device config should be
	//	used to create a label for multiple physical ports.
	Port *Adapter `protobuf:"bytes,20,opt,name=port,proto3" json:"port,omitempty"`
	// cfg - Used to pass some feature-specific configuration to the
	//
	//	network instance. For Ex: Lisp, StriongSwan etc
	Cfg *NetworkInstanceOpaqueConfig `protobuf:"bytes,30,opt,name=cfg,proto3" json:"cfg,omitempty"`
	// type of ipSpec
	IpType AddressType `protobuf:"varint,39,opt,name=ipType,proto3,enum=org.lfedge.eve.config.AddressType" json:"ipType,omitempty"`
	// network ip specification
	// If ip.gateway is set to all-zeroes IP, default route will not be propagated
	// to applications for interfaces connected to this network instance.
	// Default route propagation is also automatically suppressed when the network
	// instance is air-gapped or when the uplink is app-shared without default route
	// configured. This behaviour can be further customized using static_routes
	// (see below).
	Ip *Ipspec `protobuf:"bytes,40,opt,name=ip,proto3" json:"ip,omitempty"`
	// static DNS entry, if we are running DNS/DHCP service
	Dns []*ZnetStaticDNSEntry `protobuf:"bytes,41,rep,name=dns,proto3" json:"dns,omitempty"`
	// Enable to use DHCP to automatically propagate routes for uplink subnets
	// into applications connected to them indirectly through local network instances.
	// This option is only valid for local network instances. For other types
	// of network instances, it will be ignored.
	PropagateConnectedRoutes bool `protobuf:"varint,42,opt,name=propagate_connected_routes,json=propagateConnectedRoutes,proto3" json:"propagate_connected_routes,omitempty"`
	// List of IP routes statically added to the network instance routing table.
	// Statically routed subnets are also propagated to connected applications
	// using DHCP, with gateway set to the network instance bridge IP if it is
	// outside of the network instance subnet.
	//
	// IP route gateway may point to an external endpoint (provided that network
	// instance is not air-gapped), or to an IP address of one of the applications
	// connected to the network instance.
	//
	// Static routes are handled independently from connected routes. While connected
	// routes are propagated to applications only if enabled by propagate_connected_routes,
	// static routes are always propagated. Both connected and statically configured
	// routes can be propagated at the same time, there are no restrictions for using both.
	//
	// Note that the default route (with the bridge IP as the gateway) is automatically
	// propagated to connected applications with these exceptions:
	//
	//	a) default route propagation is explicitly disabled by setting
	//	   NetworkInstanceConfig.ip.gateway to an all-zeroes IP
	//	b) network instance is air-gapped (without uplink)
	//	c) the uplink is app-shared (not management) and does not have a default route
	//	   of its own
	//
	// In the b) and c) cases, it is possible to enforce default route propagation
	// by configuring a static default route for the network instance.
	//
	// This option is only valid for local network instances. For other types
	// of network instances, it will be ignored.
	StaticRoutes []*IPRoute `protobuf:"bytes,43,rep,name=static_routes,json=staticRoutes,proto3" json:"static_routes,omitempty"`
}

func (x *NetworkInstanceConfig) Reset() {
	*x = NetworkInstanceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_netinst_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkInstanceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInstanceConfig) ProtoMessage() {}

func (x *NetworkInstanceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_netinst_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInstanceConfig.ProtoReflect.Descriptor instead.
func (*NetworkInstanceConfig) Descriptor() ([]byte, []int) {
	return file_config_netinst_proto_rawDescGZIP(), []int{4}
}

func (x *NetworkInstanceConfig) GetUuidandversion() *UUIDandVersion {
	if x != nil {
		return x.Uuidandversion
	}
	return nil
}

func (x *NetworkInstanceConfig) GetDisplayname() string {
	if x != nil {
		return x.Displayname
	}
	return ""
}

func (x *NetworkInstanceConfig) GetInstType() ZNetworkInstType {
	if x != nil {
		return x.InstType
	}
	return ZNetworkInstType_ZNetInstFirst
}

func (x *NetworkInstanceConfig) GetActivate() bool {
	if x != nil {
		return x.Activate
	}
	return false
}

func (x *NetworkInstanceConfig) GetPort() *Adapter {
	if x != nil {
		return x.Port
	}
	return nil
}

func (x *NetworkInstanceConfig) GetCfg() *NetworkInstanceOpaqueConfig {
	if x != nil {
		return x.Cfg
	}
	return nil
}

func (x *NetworkInstanceConfig) GetIpType() AddressType {
	if x != nil {
		return x.IpType
	}
	return AddressType_First
}

func (x *NetworkInstanceConfig) GetIp() *Ipspec {
	if x != nil {
		return x.Ip
	}
	return nil
}

func (x *NetworkInstanceConfig) GetDns() []*ZnetStaticDNSEntry {
	if x != nil {
		return x.Dns
	}
	return nil
}

func (x *NetworkInstanceConfig) GetPropagateConnectedRoutes() bool {
	if x != nil {
		return x.PropagateConnectedRoutes
	}
	return false
}

func (x *NetworkInstanceConfig) GetStaticRoutes() []*IPRoute {
	if x != nil {
		return x.StaticRoutes
	}
	return nil
}

var File_config_netinst_proto protoreflect.FileDescriptor

var file_config_netinst_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6e, 0x65, 0x74, 0x69, 0x6e, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64,
	0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x16, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x64, 0x65, 0x76, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6e, 0x65,
	0x74, 0x63, 0x6d, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x01, 0x0a, 0x1b, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4f, 0x70,
	0x61, 0x71, 0x75, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x50, 0x0a, 0x0a, 0x6c, 0x69, 0x73, 0x70, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c,
	0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x4c, 0x69, 0x73, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0a, 0x6c, 0x69, 0x73, 0x70,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x43, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67,
	0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x5a, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x8a, 0x01, 0x0a, 0x0e,
	0x5a, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x3c,
	0x0a, 0x06, 0x7a, 0x73, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x5a, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x7a, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x4e, 0x61, 0x6d, 0x65, 0x4f, 0x72, 0x49, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x4e, 0x61, 0x6d, 0x65, 0x4f, 0x72, 0x49, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x22, 0xc8, 0x02, 0x0a, 0x19, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x70,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3f, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x70, 0x4d, 0x53,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66,
	0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x5a, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x07,
	0x4c, 0x69, 0x73, 0x70, 0x4d, 0x53, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x70, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0e, 0x4c, 0x69, 0x73, 0x70, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x65,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x12, 0x2a, 0x0a, 0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x61, 0x6c, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x30, 0x0a,
	0x13, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x6c, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x61, 0x6c, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x6c, 0x65, 0x6e, 0x12,
	0x22, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x6c, 0x22, 0x54, 0x0a, 0x07, 0x49, 0x50, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x2f,
	0x0a, 0x13, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12,
	0x18, 0x0a, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x22, 0x8e, 0x05, 0x0a, 0x15, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x4d, 0x0a, 0x0e, 0x75, 0x75, 0x69, 0x64, 0x61, 0x6e, 0x64, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x61, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x0e, 0x75, 0x75, 0x69, 0x64, 0x61, 0x6e, 0x64, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65,
	0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x5a,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x08, 0x69, 0x6e, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65,
	0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x64, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x44, 0x0a, 0x03, 0x63, 0x66, 0x67,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65,
	0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4f, 0x70,
	0x61, 0x71, 0x75, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x03, 0x63, 0x66, 0x67, 0x12,
	0x3a, 0x0a, 0x06, 0x69, 0x70, 0x54, 0x79, 0x70, 0x65, 0x18, 0x27, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x22, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x06, 0x69, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2d, 0x0a, 0x02, 0x69,
	0x70, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66,
	0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x69, 0x70, 0x73, 0x70, 0x65, 0x63, 0x52, 0x02, 0x69, 0x70, 0x12, 0x3b, 0x0a, 0x03, 0x64, 0x6e,
	0x73, 0x18, 0x29, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66,
	0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x5a, 0x6e, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x44, 0x4e, 0x53, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x03, 0x64, 0x6e, 0x73, 0x12, 0x3c, 0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x70, 0x61,
	0x67, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x70, 0x72, 0x6f,
	0x70, 0x61, 0x67, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x43, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x2b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x49, 0x50, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x0c, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x63, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2a, 0xb3, 0x01, 0x0a, 0x10, 0x5a,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x11, 0x0a, 0x0d, 0x5a, 0x4e, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x46, 0x69, 0x72, 0x73, 0x74,
	0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x5a, 0x6e, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x53, 0x77,
	0x69, 0x74, 0x63, 0x68, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x5a, 0x6e, 0x65, 0x74, 0x49, 0x6e,
	0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x5a, 0x6e, 0x65,
	0x74, 0x49, 0x6e, 0x73, 0x74, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c,
	0x5a, 0x6e, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x68, 0x10, 0x04, 0x12, 0x14,
	0x0a, 0x10, 0x5a, 0x6e, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x48, 0x6f, 0x6e, 0x65, 0x79, 0x50,
	0x6f, 0x74, 0x10, 0x05, 0x12, 0x17, 0x0a, 0x13, 0x5a, 0x6e, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x10, 0x06, 0x12, 0x11, 0x0a,
	0x0c, 0x5a, 0x4e, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x10, 0xff, 0x01,
	0x2a, 0x57, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x09, 0x0a, 0x05, 0x46, 0x69, 0x72, 0x73, 0x74, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x50,
	0x56, 0x34, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x50, 0x56, 0x36, 0x10, 0x02, 0x12, 0x0e,
	0x0a, 0x0a, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x49, 0x50, 0x56, 0x34, 0x10, 0x03, 0x12, 0x0e,
	0x0a, 0x0a, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x49, 0x50, 0x56, 0x36, 0x10, 0x04, 0x12, 0x09,
	0x0a, 0x04, 0x4c, 0x61, 0x73, 0x74, 0x10, 0xff, 0x01, 0x2a, 0x43, 0x0a, 0x18, 0x5a, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x5a, 0x4e, 0x65, 0x74, 0x4f, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x56, 0x50, 0x4e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x5a, 0x4e, 0x65,
	0x74, 0x4f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4c, 0x69, 0x73, 0x70, 0x10, 0x01, 0x2a, 0x47,
	0x0a, 0x0d, 0x5a, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x10, 0x7a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x53, 0x72, 0x76, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x6d, 0x61, 0x70, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x10, 0x02, 0x42, 0x3d, 0x0a, 0x15, 0x6f, 0x72, 0x67, 0x2e, 0x6c,
	0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x66, 0x2d,
	0x65, 0x64, 0x67, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_netinst_proto_rawDescOnce sync.Once
	file_config_netinst_proto_rawDescData = file_config_netinst_proto_rawDesc
)

func file_config_netinst_proto_rawDescGZIP() []byte {
	file_config_netinst_proto_rawDescOnce.Do(func() {
		file_config_netinst_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_netinst_proto_rawDescData)
	})
	return file_config_netinst_proto_rawDescData
}

var file_config_netinst_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_config_netinst_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_config_netinst_proto_goTypes = []interface{}{
	(ZNetworkInstType)(0),               // 0: org.lfedge.eve.config.ZNetworkInstType
	(AddressType)(0),                    // 1: org.lfedge.eve.config.AddressType
	(ZNetworkOpaqueConfigType)(0),       // 2: org.lfedge.eve.config.ZNetworkOpaqueConfigType
	(ZcServiceType)(0),                  // 3: org.lfedge.eve.config.ZcServiceType
	(*NetworkInstanceOpaqueConfig)(nil), // 4: org.lfedge.eve.config.NetworkInstanceOpaqueConfig
	(*ZcServicePoint)(nil),              // 5: org.lfedge.eve.config.ZcServicePoint
	(*NetworkInstanceLispConfig)(nil),   // 6: org.lfedge.eve.config.NetworkInstanceLispConfig
	(*IPRoute)(nil),                     // 7: org.lfedge.eve.config.IPRoute
	(*NetworkInstanceConfig)(nil),       // 8: org.lfedge.eve.config.NetworkInstanceConfig
	(*UUIDandVersion)(nil),              // 9: org.lfedge.eve.config.UUIDandVersion
	(*Adapter)(nil),                     // 10: org.lfedge.eve.config.Adapter
	(*Ipspec)(nil),                      // 11: org.lfedge.eve.config.ipspec
	(*ZnetStaticDNSEntry)(nil),          // 12: org.lfedge.eve.config.ZnetStaticDNSEntry
}
var file_config_netinst_proto_depIdxs = []int32{
	6,  // 0: org.lfedge.eve.config.NetworkInstanceOpaqueConfig.lispConfig:type_name -> org.lfedge.eve.config.NetworkInstanceLispConfig
	2,  // 1: org.lfedge.eve.config.NetworkInstanceOpaqueConfig.type:type_name -> org.lfedge.eve.config.ZNetworkOpaqueConfigType
	3,  // 2: org.lfedge.eve.config.ZcServicePoint.zsType:type_name -> org.lfedge.eve.config.ZcServiceType
	5,  // 3: org.lfedge.eve.config.NetworkInstanceLispConfig.LispMSs:type_name -> org.lfedge.eve.config.ZcServicePoint
	9,  // 4: org.lfedge.eve.config.NetworkInstanceConfig.uuidandversion:type_name -> org.lfedge.eve.config.UUIDandVersion
	0,  // 5: org.lfedge.eve.config.NetworkInstanceConfig.instType:type_name -> org.lfedge.eve.config.ZNetworkInstType
	10, // 6: org.lfedge.eve.config.NetworkInstanceConfig.port:type_name -> org.lfedge.eve.config.Adapter
	4,  // 7: org.lfedge.eve.config.NetworkInstanceConfig.cfg:type_name -> org.lfedge.eve.config.NetworkInstanceOpaqueConfig
	1,  // 8: org.lfedge.eve.config.NetworkInstanceConfig.ipType:type_name -> org.lfedge.eve.config.AddressType
	11, // 9: org.lfedge.eve.config.NetworkInstanceConfig.ip:type_name -> org.lfedge.eve.config.ipspec
	12, // 10: org.lfedge.eve.config.NetworkInstanceConfig.dns:type_name -> org.lfedge.eve.config.ZnetStaticDNSEntry
	7,  // 11: org.lfedge.eve.config.NetworkInstanceConfig.static_routes:type_name -> org.lfedge.eve.config.IPRoute
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_config_netinst_proto_init() }
func file_config_netinst_proto_init() {
	if File_config_netinst_proto != nil {
		return
	}
	file_config_devcommon_proto_init()
	file_config_netcmn_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_config_netinst_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkInstanceOpaqueConfig); i {
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
		file_config_netinst_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZcServicePoint); i {
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
		file_config_netinst_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkInstanceLispConfig); i {
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
		file_config_netinst_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPRoute); i {
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
		file_config_netinst_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkInstanceConfig); i {
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
			RawDescriptor: file_config_netinst_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_netinst_proto_goTypes,
		DependencyIndexes: file_config_netinst_proto_depIdxs,
		EnumInfos:         file_config_netinst_proto_enumTypes,
		MessageInfos:      file_config_netinst_proto_msgTypes,
	}.Build()
	File_config_netinst_proto = out.File
	file_config_netinst_proto_rawDesc = nil
	file_config_netinst_proto_goTypes = nil
	file_config_netinst_proto_depIdxs = nil
}
