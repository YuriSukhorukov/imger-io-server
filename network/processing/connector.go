package processing

import "net"

func AddConnection(conn net.Conn) {
	Connections[&conn] = conn
	go Stream(conn)
	go PacketsQueueIOSystem()
}

func RemoveConnection(key *net.Conn) {
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