package processing

import (
	"../../protocol"
)

func PacketsQueueIOSystem() {
	for {
		for k, v := range PacketsQueueIO {
			conn     := *k
			packets  := v
			roomID   := ConnectionsRooms[&conn]
			for _, c := range Rooms[roomID].Connections {
				conn := *c
				conn.Write([]byte{0xAA})
			}
			for _, p := range packets {
				packet := p
				conn.Write(protocol.Pack(packet))
			}
			// ...
			delete(PacketsQueueIO, k)
		}
	}
	//PacketNewRoomQueueSystem()
	//PacketsSome1QueueSystem()
	//PacketsSome2QueueSystem()
	//PacketSome3System()
}

func PacketNewRoomQueueSystem() {}

func PacketsSome1QueueSystem() {}
func PacketsSome2QueueSystem() {}
func PacketsSome3QueueSystem() {}