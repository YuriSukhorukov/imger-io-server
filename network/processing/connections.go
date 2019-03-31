package processing

import "net"

var Connections = make(map[*net.Conn]net.Conn)
var ConnectionsRooms = make(map[*net.Conn]string)