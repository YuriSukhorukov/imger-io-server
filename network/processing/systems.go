package processing

import (
	"../../protocol"
	"../../util"
	"fmt"
	"net"
)

func PacketsQueueSystemIO(chanel chan protocol.Packet, conn net.Conn) {
	for {
		packet := <-chanel
		util.PrintMemUsage()
		fmt.Printf("packet PacketsQueueSystemIO: %v\n", packet)
		if packet.PacketType == protocol.PLAYER {
			PacketsQueueSome1[conn] <- packet
		} else if packet.PacketType == protocol.CANVAS {
			PacketsQueueSome2[conn] <- packet
		} else if packet.PacketType == protocol.ENTITY {
			PacketsQueueSome3[conn] <- packet
		}
		//roomID := ConnectionsRooms[conn]
		//for _, c := range Rooms[roomID] {
		//	c.Write(protocol.Pack(packet))
		//}
	}
}

func PacketsQueueSystemSome1(chanel chan protocol.Packet, conn net.Conn) {
	for {
		packet := <- chanel
		fmt.Printf("packet PacketsQueueSystemSome1: %v\n", packet)
		conn.Write(protocol.Pack(packet))
	}
}

func PacketsQueueSystemSome2(chanel chan protocol.Packet, conn net.Conn) {
	for {
		packet := <- chanel
		fmt.Printf("packet PacketsQueueSystemSome2: %v\n", packet)
		conn.Write(protocol.Pack(packet))
	}
}

func PacketsQueueSystemSome3(chanel chan protocol.Packet, conn net.Conn) {
	for {
		packet := <- chanel
		fmt.Printf("packet PacketsQueueSystemSome3: %v\n", packet)
		conn.Write(protocol.Pack(packet))
	}
}

func PacketNewRoomQueueSystem(chanel chan protocol.Packet, conn net.Conn) {
	for {
		packet := <- chanel
		fmt.Printf("packet PacketNewRoomQueueSystem: %v\n", packet)
		conn.Write(protocol.Pack(packet))
	}
}