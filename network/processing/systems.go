package processing

import (
	"../../protocol"
	"../../util"
	"fmt"
	"net"
)

func PacketsChannelSystemIO(channel chan protocol.Packet, conn net.Conn) {
	for {
		packet := <-channel
		util.PrintMemUsage()
		fmt.Printf("packet PacketsChannelSystemIO: %v\n", packet)
		if packet.PacketType == protocol.PLAYER {
			PacketsChannelSome1[conn] <- packet
		} else if packet.PacketType == protocol.CANVAS {
			PacketsChannelSome2[conn] <- packet
		} else if packet.PacketType == protocol.ENTITY {
			PacketsChannelSome3[conn] <- packet
		}
		//roomID := ConnectionsRooms[conn]
		//for _, c := range Rooms[roomID] {
		//	c.Write(protocol.Pack(packet))
		//}
	}
}

func PacketsChannelSystemSome1(channel chan protocol.Packet, conn net.Conn) {
	for {
		packet := <-channel
		fmt.Printf("packet PacketsChannelSystemSome1: %v\n", packet)
		conn.Write(protocol.Pack(packet))
	}
}

func PacketsChannelSystemSome2(channel chan protocol.Packet, conn net.Conn) {
	for {
		packet := <-channel
		fmt.Printf("packet PacketsChannelSystemSome2: %v\n", packet)
		conn.Write(protocol.Pack(packet))
	}
}

func PacketsChannelSystemSome3(channel chan protocol.Packet, conn net.Conn) {
	for {
		packet := <-channel
		fmt.Printf("packet PacketsChannelSystemSome3: %v\n", packet)
		conn.Write(protocol.Pack(packet))
	}
}

func PacketNewRoomChannelSystem(channel chan protocol.Packet, conn net.Conn) {
	for {
		packet := <-channel
		fmt.Printf("packet PacketNewRoomChannelSystem: %v\n", packet)
		conn.Write(protocol.Pack(packet))
	}
}