package processing

import (
	"net"
)

type Room struct {
	Connections []*net.Conn
}

var Rooms = make(map[string]Room)

// Rooms - карта коллекций соединений
// Rooms = [roomID:[connID, connID, ...], roomID:[connID, connID, ...], ...]