package processing

import (
	"../../protocol"
	"net"
)

func AddConnection(conn net.Conn) {
	Connections[conn]           = conn
	PacketsQueueIO[conn]        = make(chan protocol.Packet)
	PacketsQueueSome1[conn]     = make(chan protocol.Packet)
	PacketsQueueSome2[conn]     = make(chan protocol.Packet)
	PacketsQueueSome3[conn]     = make(chan protocol.Packet)
	PacketsQueueNewRoom[conn]   = make(chan protocol.Packet)
	go Stream(conn)
	go PacketsQueueSystemIO(PacketsQueueIO[conn], conn)
	go PacketsQueueSystemSome1(PacketsQueueSome1[conn], conn)
	go PacketsQueueSystemSome2(PacketsQueueSome2[conn], conn)
	go PacketsQueueSystemSome3(PacketsQueueSome3[conn], conn)
	go PacketNewRoomQueueSystem(PacketsQueueNewRoom[conn], conn)
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