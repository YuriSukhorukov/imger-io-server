package network

import (
	"fmt"
	"net"
	"os"
)

func Open() {
	ln, err := net.Listen(CONN_TYPE, CONN_PORT)
	fmt.Println("Launching server...")
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		os.Exit(1)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		if conn != nil {
			fmt.Printf("New connection\n")
			go Stream(conn)
		}
	}
}
