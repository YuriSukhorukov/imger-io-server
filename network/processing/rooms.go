package processing

import (
	"net"
)

var Rooms = make(map[string][]net.Conn)

// Rooms - карта коллекций соединений
// Rooms = [roomID:[connID, connID, ...], roomID:[connID, connID, ...], ...]