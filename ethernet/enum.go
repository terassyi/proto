package ethernet

const (
	ETHER_TYPE_IP   EtherType = 0x0800
	ETHER_TYPE_ARP  EtherType = 0x0806
	ETHER_TYPE_IPV6 EtherType = 0x86dd
)

const EthernetHeaderLength int = 14

var (
	EmptyAddress     = HardwareAddress{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	InvalidAddress   = HardwareAddress{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	BroadcastAddress = HardwareAddress{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
)
