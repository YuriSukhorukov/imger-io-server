package processing

import (
	"net"
	"../../protocol"
)

var PacketsQueueIO  = make(map[*net.Conn][]protocol.Packet)
var PacketsNewRoom  = make(map[*net.Conn][]protocol.Packet)
var PacketsSome1    = make(map[*net.Conn][]protocol.Packet)
var PacketsSome2    = make(map[*net.Conn][]protocol.Packet)
var PacketsSome3    = make(map[*net.Conn][]protocol.Packet)