package processing

import (
	"../../protocol"
	"net"
)

func AddConnection(conn net.Conn) {
	Connections[conn]             = conn
	PacketsChannelIO[conn]        = make(chan protocol.Packet)
	PacketsChannelSome1[conn]     = make(chan protocol.Packet)
	PacketsChannelSome2[conn]     = make(chan protocol.Packet)
	PacketsChannelSome3[conn]     = make(chan protocol.Packet)
	PacketsChannelNewRoom[conn]   = make(chan protocol.Packet)
	go Stream(conn)
	go PacketsChannelSystemIO(PacketsChannelIO[conn], conn)
	go PacketsChannelSystemSome1(PacketsChannelSome1[conn], conn)
	go PacketsChannelSystemSome2(PacketsChannelSome2[conn], conn)
	go PacketsChannelSystemSome3(PacketsChannelSome3[conn], conn)
	go PacketNewRoomChannelSystem(PacketsChannelNewRoom[conn], conn)
}

func RemoveConnection(key net.Conn) {
	delete(Connections, key)
}

func AddRoom() {

}
func RemoveRoom() {

}
func AddToRoom() {

}
func RemoveFromRoom() {

}
