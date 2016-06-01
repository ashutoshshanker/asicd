//
//Copyright [2016] [SnapRoute Inc]
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//	 Unless required by applicable law or agreed to in writing, software
//	 distributed under the License is distributed on an "AS IS" BASIS,
//	 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//	 See the License for the specific language governing permissions and
//	 limitations under the License.
//
// _______  __       __________   ___      _______.____    __    ____  __  .___________.  ______  __    __  
// |   ____||  |     |   ____\  \ /  /     /       |\   \  /  \  /   / |  | |           | /      ||  |  |  | 
// |  |__   |  |     |  |__   \  V  /     |   (----` \   \/    \/   /  |  | `---|  |----`|  ,----'|  |__|  | 
// |   __|  |  |     |   __|   >   <       \   \      \            /   |  |     |  |     |  |     |   __   | 
// |  |     |  `----.|  |____ /  .  \  .----)   |      \    /\    /    |  |     |  |     |  `----.|  |  |  | 
// |__|     |_______||_______/__/ \__\ |_______/        \__/  \__/     |__|     |__|      \______||__|  |__| 
//                                                                                                           

package pluginCommon

import (
	"net"
)

//#include "pluginCommon.h"
import "C"

var SwitchMacAddr net.HardwareAddr

//Consolidated list of constants used by underlying plugins
const (
	//Generic consts
	DEFAULT_SWITCH_MAC_ADDR = "00:11:22:33:44:55"
	MAC_ADDR_LEN            = C.MAC_ADDR_LEN
	INVALID_OBJECT_ID       = 0xFFFFFFFFFFFFFFFF
	INVALID_IFINDEX         = -1

	//System related consts
	BOOT_MODE_COLDBOOT = C.BOOT_MODE_COLDBOOT
	BOOT_MODE_WARMBOOT = C.BOOT_MODE_WARMBOOT
	INTF_STATE_UP      = C.INTF_STATE_UP
	INTF_STATE_DOWN    = C.INTF_STATE_DOWN
	MIN_SYS_PORTS      = 0
	MAX_SYS_PORTS      = 256

	//FDB relate consts
	MAC_ENTRY_LEARNED = C.MAC_ENTRY_LEARNED
	MAC_ENTRY_AGED    = C.MAC_ENTRY_AGED

	//Vlan related consts
	SVI_PREFIX        = "vlan"
	SYS_RSVD_VLAN     = -1
	MAX_VLAN_ID       = C.MAX_VLAN_ID
	DEFAULT_VLAN_ID   = C.DEFAULT_VLAN_ID
	SYS_RSVD_VLAN_MIN = 3835
	SYS_RSVD_VLAN_MAX = 4090

	//Port related consts
	MAX_PORT_STAT_TYPES            = C.portStatTypesMax
	PORT_BREAKOUT_MODE_UNSUPPORTED = C.PORT_BREAKOUT_MODE_UNSUPPORTED
	PORT_BREAKOUT_MODE_1x40        = C.PORT_BREAKOUT_MODE_1x40
	PORT_BREAKOUT_MODE_4x10        = C.PORT_BREAKOUT_MODE_4x10
	PORT_ATTR_PHY_INTF_TYPE        = C.PORT_ATTR_PHY_INTF_TYPE
	PORT_ATTR_ADMIN_STATE          = C.PORT_ATTR_ADMIN_STATE
	PORT_ATTR_MAC_ADDR             = C.PORT_ATTR_MAC_ADDR
	PORT_ATTR_SPEED                = C.PORT_ATTR_SPEED
	PORT_ATTR_DUPLEX               = C.PORT_ATTR_DUPLEX
	PORT_ATTR_AUTONEG              = C.PORT_ATTR_AUTONEG
	PORT_ATTR_MEDIA_TYPE           = C.PORT_ATTR_MEDIA_TYPE
	PORT_ATTR_MTU                  = C.PORT_ATTR_MTU
	PORT_ATTR_BREAKOUT_MODE        = C.PORT_ATTR_BREAKOUT_MODE

	//STP related consts
	STP_PORT_STATE_BLOCKING   = C.StpPortStateBlocking
	STP_PORT_STATE_LEARNING   = C.StpPortStateLearning
	STP_PORT_STATE_FORWARDING = C.StpPortStateForwarding

	//Lag related consts
	LAG_PREFIX             = "LAG"
	HASHTYPE_SRCMAC_DSTMAC = C.HASHTYPE_SRCMAC_DSTMAC
	HASHTYPE_SRCIP_DSTIP   = C.HASHTYPE_SRCIP_DSTIP

	//Next hop related consts
	NEIGHBOR_TYPE_COPY_TO_CPU       = C.NEIGHBOR_TYPE_COPY_TO_CPU
	NEIGHBOR_TYPE_BLACKHOLE         = C.NEIGHBOR_TYPE_BLACKHOLE
	NEIGHBOR_TYPE_FULL_SPEC_NEXTHOP = C.NEIGHBOR_TYPE_FULL_SPEC_NEXTHOP
	NEIGHBOR_L2_ACCESS_TYPE_PORT    = C.NEIGHBOR_L2_ACCESS_TYPE_PORT
	NEIGHBOR_L2_ACCESS_TYPE_LAG     = C.NEIGHBOR_L2_ACCESS_TYPE_LAG

	//Next hop related consts
	NEXTHOP_TYPE_COPY_TO_CPU       = C.NEXTHOP_TYPE_COPY_TO_CPU
	NEXTHOP_TYPE_BLACKHOLE         = C.NEXTHOP_TYPE_BLACKHOLE
	NEXTHOP_TYPE_FULL_SPEC_NEXTHOP = C.NEXTHOP_TYPE_FULL_SPEC_NEXTHOP
	NEXTHOP_L2_ACCESS_TYPE_PORT    = C.NEXTHOP_L2_ACCESS_TYPE_PORT
	NEXTHOP_L2_ACCESS_TYPE_LAG     = C.NEXTHOP_L2_ACCESS_TYPE_LAG
	INVALID_NEXTHOP_ID             = 0xFFFFFFFFFFFFFFFF

	//Route related consts
	MAX_NEXTHOPS_PER_GROUP      = C.MAX_NEXTHOPS_PER_GROUP
	ROUTE_OPERATION_TYPE_UPDATE = C.ROUTE_OPERATION_TYPE_UPDATE
	ROUTE_TYPE_CONNECTED        = C.ROUTE_TYPE_CONNECTED
	ROUTE_TYPE_MULTIPATH        = C.ROUTE_TYPE_MULTIPATH
	ROUTE_TYPE_SINGLEPATH       = C.ROUTE_TYPE_SINGLEPATH
)

// Interface id/type mgmt
const (
	INTF_TYPE_MASK  = C.INTF_TYPE_MASK
	INTF_TYPE_SHIFT = C.INTF_TYPE_SHIFT
	INTF_ID_MASK    = C.INTF_ID_MASK
	INTF_ID_SHIFT   = C.INTF_ID_SHIFT
)

// Func types for intf id mgmt
type GetId func(int32) int
type GetType func(int32) int
type GetIfIndex func(int, int) int32

func GetTypeFromIfIndex(ifIndex int32) int {
	return int((ifIndex & INTF_TYPE_MASK) >> INTF_TYPE_SHIFT)
}
func GetIdFromIfIndex(ifIndex int32) int {
	return int((ifIndex & INTF_ID_MASK) >> INTF_ID_SHIFT)
}
func GetIfIndexFromIdType(ifId, ifType int) int32 {
	return int32((ifId & INTF_ID_MASK) | ((ifType << INTF_TYPE_SHIFT) & INTF_TYPE_MASK))
}

var OnOffState map[int]string = map[int]string{0: "OFF", 1: "ON"}
var UpDownState map[int]string = map[int]string{0: "DOWN", 1: "UP"}
var IfType map[int]string = map[int]string{
	int(C.PortIfTypeMII):    "MII",
	int(C.PortIfTypeGMII):   "GMII",
	int(C.PortIfTypeSGMII):  "SGMII",
	int(C.PortIfTypeQSGMII): "QSGMII",
	int(C.PortIfTypeSFI):    "SFI",
	int(C.PortIfTypeXFI):    "XFI",
	int(C.PortIfTypeXAUI):   "XAUI",
	int(C.PortIfTypeXLAUI):  "XLAUI",
	int(C.PortIfTypeRXAUI):  "RXAUI",
	int(C.PortIfTypeCR):     "CR",
	int(C.PortIfTypeCR2):    "CR2",
	int(C.PortIfTypeCR4):    "CR4",
	int(C.PortIfTypeKR):     "KR",
	int(C.PortIfTypeKR2):    "KR2",
	int(C.PortIfTypeKR4):    "KR4",
	int(C.PortIfTypeSR):     "SR",
	int(C.PortIfTypeSR2):    "SR2",
	int(C.PortIfTypeSR4):    "SR4",
	int(C.PortIfTypeSR10):   "SR10",
	int(C.PortIfTypeLR):     "LR",
	int(C.PortIfTypeLR4):    "LR4",
}
var DuplexType map[int]string = map[int]string{
	int(C.HalfDuplex): "Half Duplex",
	int(C.FullDuplex): "Full Duplex",
}

const (
	//Asicd notification msgs
	NOTIFY_L2INTF_STATE_CHANGE = iota
	NOTIFY_L3INTF_STATE_CHANGE
	NOTIFY_VLAN_CREATE
	NOTIFY_VLAN_DELETE
	NOTIFY_VLAN_UPDATE
	NOTIFY_LOGICAL_INTF_CREATE
	NOTIFY_LOGICAL_INTF_DELETE
	NOTIFY_LOGICAL_INTF_UPDATE
	NOTIFY_IPV4INTF_CREATE
	NOTIFY_IPV4INTF_DELETE
	NOTIFY_LAG_CREATE
	NOTIFY_LAG_DELETE
	NOTIFY_LAG_UPDATE
	NOTIFY_IPV4NBR_MAC_MOVE
	NOTIFY_IPV4_ROUTE_CREATE_FAILURE
	NOTIFY_IPV4_ROUTE_DELETE_FAILURE
)

// Format of asicd's published messages
type AsicdNotification struct {
	MsgType uint8
	Msg     []byte
}

// Following sections contains formats for individual message types
type L2IntfStateNotifyMsg struct {
	IfIndex int32
	IfState uint8
}
type L3IntfStateNotifyMsg struct {
	IpAddr  string
	IfIndex int32
	IfState uint8
}
type VlanNotifyMsg struct {
	VlanId     uint16
	VlanName   string
	TagPorts   []int32
	UntagPorts []int32
}
type LogicalIntfNotifyMsg struct {
	IfIndex         int32
	LogicalIntfName string
}
type LagNotifyMsg struct {
	LagName     string
	IfIndex     int32
	IfIndexList []int32
}
type IPv4IntfNotifyMsg struct {
	IpAddr  string
	IfIndex int32
}
type IPv4NbrMacMoveNotifyMsg struct {
	IpAddr  string
	IfIndex int32
	VlanId  int32
}

// Struct containing info required for mapping asic ports to linux interfaces
type IfMapInfo struct {
	IfName string
	Port   int
}

// Struct used for configuring sub interface on all plugins
type SubIntfPluginObj struct {
	IfIndex        int32
	IpAddr         uint32
	MaskLen        int
	VlanId         int
	StateUp        bool
	SubIntfIfIndex int32
	MacAddr        net.HardwareAddr
}

// Struct used for configuring ports
type PortConfig struct {
	PortNum           int32
	Description       string
	PhyIntfType       string
	AdminState        string
	MacAddr           string
	Speed             int32
	Duplex            string
	Autoneg           string
	MediaType         string
	Mtu               int32
	LogicalPortInfo   bool
	MappedToHw        bool
	BreakOutMode      int32
	BreakOutSupported bool
}

// Struct used for retrieving port state
type PortState struct {
	PortNum           int32
	IfIndex           int32
	Name              string
	OperState         string
	NumUpEvents       int32
	LastUpEventTime   string
	NumDownEvents     int32
	LastDownEventTime string
	Pvid              int32
	IfInOctets        int64
	IfInUcastPkts     int64
	IfInDiscards      int64
	IfInErrors        int64
	IfInUnknownProtos int64
	IfOutOctets       int64
	IfOutUcastPkts    int64
	IfOutDiscards     int64
	IfOutErrors       int64
	ErrDisableReason  string
}

// Struct used for ECMP/WCMP group creation
type NextHopGroupMemberInfo struct {
	IpAddr    uint32
	NextHopId uint64
	Weight    int
	RifId     int32
}

//Format of callback functions for updating DBs in individual resource managers
type ProcessLinkStateChangeCB func(int32, int32, string, string)
type InitPortConfigDBCB func(*PortConfig)
type InitPortStateDBCB func(int32, string)
type UpdatePortStateDBCB func(*PortState)
type UpdateLagDBCB func(int, int, []int32)
type UpdateIPv4NeighborDBCB func(uint32, net.HardwareAddr, int32, uint64)
type UpdateIPv4RouteDBCB func(uint32, uint32, uint32)
type UpdateIPv4NextHopDBCB func()
type UpdateIPv4NextHopGroupDBCB func()
type MacEntryTableCB func(int, int32, int32, net.HardwareAddr)

func ComputeSetDifference(a, b []int32) (aDiffb []int32) {
	var bMap map[int32]bool = make(map[int32]bool, 0)
	if len(a) == 0 {
		return a
	}
	if len(b) == 0 {
		return a
	}
	for _, elem := range b {
		bMap[elem] = true
	}
	for _, elem := range a {
		if _, ok := bMap[elem]; !ok {
			aDiffb = append(aDiffb, elem)
		}
	}
	return aDiffb
}
