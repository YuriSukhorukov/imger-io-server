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
		fmt.Printf("\n--->>>\n")
		// Будем прослушивать все сообщения разделенные \n
		//message, _ := bufio.NewReader(conn).ReadString('\n')
		//data := make([]byte, 1024)
		//message, _ := conn.Read(data)
		//s := string(data[:message])

		buff, err := bufio.NewReader(conn).ReadBytes(0x00)
		if err != nil {
			return
		}

		packet := protocol.Unpack(buff)
		if packet.PacketType == protocol.ENTITY &&
			packet.PacketSubtype == protocol.PLAYER {
			player := protocol.DecodePlayer(packet.Fields)
			fmt.Printf("player from packet:  %v\n", player)
			fmt.Printf("received packet:     %v\n", packet)
		}

		//// Отправить новую строку обратно клиенту
		//conn.Write([]byte(s))
		//fmt.Printf("\n---<<<")
	}
}