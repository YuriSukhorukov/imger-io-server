package processing

import (
	"net"
	"../../protocol"
)

//var PacketsQueueIO       = make(map[*net.Conn][]protocol.Packet)
var PacketsQueueIO       = make(map[net.Conn]chan protocol.Packet)
var PacketsQueueSome1    = make(map[net.Conn]chan protocol.Packet)
var PacketsQueueSome2    = make(map[net.Conn]chan protocol.Packet)
var PacketsQueueSome3    = make(map[net.Conn]chan protocol.Packet)
var PacketsQueueNewRoom  = make(map[net.Conn]chan protocol.Packet)

//var chanel = make(chan protocol.Packet)