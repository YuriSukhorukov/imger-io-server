package network

import (
	"bufio"
	"fmt"
	"net"
)

func Stream(conn net.Conn) {
	for {
		buff, err := bufio.NewReader(conn).ReadBytes(0x00)
		if err != nil {
			return
		}
		fmt.Printf("%v\n", buff)
		conn.Write([]byte{0xAA})
	}
}