package processing

import (
	"net"
	"../../protocol"
)

var PacketsChannelIO       = make(map[net.Conn]chan protocol.Packet)
var PacketsChannelSome1    = make(map[net.Conn]chan protocol.Packet)
var PacketsChannelSome2    = make(map[net.Conn]chan protocol.Packet)
var PacketsChannelSome3    = make(map[net.Conn]chan protocol.Packet)
var PacketsChannelNewRoom  = make(map[net.Conn]chan protocol.Packet)