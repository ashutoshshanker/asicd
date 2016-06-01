namespace go asicdInt
typedef i32 int

// Struct for Configuring Reserved Mac Addr on Chip
// so that packets are punted to CPU when enabled
struct RsvdProtocolMacConfig {
        1: string MacAddr
        2: string MacAddrMask
        3: i32 VlanId
}
struct Lag {
    1:i32 LagIfIndex
    2:i32 HashType
    3:list<i32> IfIndexList
}
struct LagGetInfo {
    1: int StartIdx
    2: int EndIdx
    3: int Count
    4: bool More
    5: list<Lag> LagList
}
struct IPv4NextHop {
    1: string NextHopIp
    2: i32 Weight
    3: i32 NextHopIfType
}
struct IPv4Route {
    1: string destinationNw
    2: string networkMask
    3: list<IPv4NextHop> NextHopList
}
struct Vlan {
	1 : i32 VlanId
	2 : list<i32> IfIndexList
	3 : list<i32> UntagIfIndexList
}
struct VlanGetInfo {
	1: int StartIdx
	2: int EndIdx
	3: int Count
	4: bool More
	5: list<Vlan> VlanList
}
struct Intf {
    1: string IfName
    2: i32 IfIndex
}
struct IntfGetInfo {
	1: int StartIdx
	2: int EndIdx
	3: int Count
	4: bool More
	5: list<Intf> IntfList
}
service ASICDINTServices {
    // All services listed here are utilities provided to other daemons. These are hand-written and not auto genereated.
    //Vlan
    VlanGetInfo GetBulkVlan(1: int fromIndex, 2: int count);

    //Intf
    IntfGetInfo GetBulkIntf(1: int fromIndex, 2: int count);

    //STP
    i32 CreateStg(1:list<i32> vlanList);
    bool DeleteStg(1:i32 stgId);
    bool SetPortStpState(1:i32 stgId, 2:i32 port, 3:i32 stpState);
    i32 GetPortStpState(1:i32 stgId, 2:i32 port);
    bool UpdateStgVlanList(1:i32 stgId, 2:list<i32> vlanList);
    i32 FlushFdbStgGroup(1:i32 stgId);

    //LAG
    i32 CreateLag(1:i32 hashType, 2:string ifIndexList);
    i32 DeleteLag(1:i32 lagId);
    i32 UpdateLag(1:i32 lagId, 2:i32 hashType, 3:string ifIndexList);
    LagGetInfo GetBulkLag(1:int fromIndex, 2:int count);

    //IPv4 neighbors
    i32 CreateIPv4Neighbor(1:string ipAddr, 2:string macAddr, 3:i32 vlanId, 4:i32 ifIndex);
    i32 UpdateIPv4Neighbor(1:string ipAddr, 2:string macAddr, 3:i32 vlanId, 4:i32 ifIndex);
    i32 DeleteIPv4Neighbor(1:string ipAddr, 2:string macAddr, 3:i32 vlanId, 4:i32 ifIndex);

    //IPv4 routes
    oneway void OnewayCreateIPv4Route(1:list<IPv4Route> ipv4RouteList);
    oneway void OnewayDeleteIPv4Route(1:list<IPv4Route> ipv4RouteList);

    //Protocol Mac Addr
    bool EnablePacketReception(1:RsvdProtocolMacConfig config);
    bool DisablePacketReception(1:RsvdProtocolMacConfig config);

    //Err-disable
    bool ErrorDisablePort(1: i32 ifIndex, 2:string AdminState, 3:string ErrDisableReason)
}
