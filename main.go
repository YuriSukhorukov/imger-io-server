package main

import (
	"bufio"
	"net"
)
import "os"
import "fmt"
import "./protocol"

// требуется только ниже для обработки примера

const CONN_HOST = "localhost"
const CONN_PORT = ":8081"
const CONN_TYPE = "tcp"
const CONN_URL  = CONN_HOST + ":" + CONN_PORT

func main() {
	fmt.Println("Launching server...")
	// Устанавливаем прослушивание порта
	ln, err := net.Listen(CONN_TYPE, CONN_PORT)

	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	// Close the listener when this application closes
	defer ln.Close()

	// Открываем порт
	conn, _ := ln.Accept()
	defer conn.Close()
	// Запускаем цикл
	for {
		buff, err := bufio.NewReader(conn).ReadBytes(0x00)
		if err != nil {
			return
		}
	}
}
