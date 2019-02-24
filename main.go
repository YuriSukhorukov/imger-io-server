package main

import (
	"bufio"
	"net"
)
import "os"
import "fmt"
// требуется только ниже для обработки примера

const CONN_HOST = "localhost"
const CONN_PORT = ":8081"
const CONN_TYPE = "tcp"
const CONN_URL  = CONN_HOST + ":" + CONN_PORT

const MESSAGE_0xAA = 0xAA
const MESSAGE_0xAF = 0xAF

// PacketType - первый байт, определяюзий тип пакета
// Content 	  - содержимое пакета
type Packet struct {
	PacketType byte
	Content []byte
}

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
		//// Распечатываем полученое сообщение
		//fmt.Print("Message Received: ", s)
		//fmt.Print("Message Received: ", data)
		//// Процесс выборки для полученной строки
		//// Отправить новую строку обратно клиенту
		//conn.Write([]byte(s))
		//fmt.Printf("\n---<<<")

		buff, err := bufio.NewReader(conn).ReadBytes(0x00)
		if err != nil {
			return
		}
		packet := Packet{buff[0], buff[1:len(buff)-2]}

		if packet.PacketType == MESSAGE_0xAA {
			fmt.Printf("Message type: MESSAGE_0xAA \n")
		} else if packet.PacketType == MESSAGE_0xAF {
			fmt.Printf("Message type: MESSAGE_0xAA \n")
		}

		fmt.Printf(string(packet.Content))
	}
}