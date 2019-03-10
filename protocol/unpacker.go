package protocol

import "fmt"

// распаковщик принимает на вход последодвательность байт, коллекцию типов сообщений
// а возвращает пакет, готовый к использованию

func Unpack(buff []byte) Packet {
	packetType 		:= buff[0]
	packetSubtype 	:= buff[1]

	fields 			:= make([]Field, 0)

	var fieldID 	byte
	var fieldSize 	byte
	fmt.Printf("buff size: %d\n", len(buff))

	for i := 2; i + 2 + int(fieldSize) < len(buff); i += 2 + int(fieldSize) {
		fmt.Printf("index: %d\n", i)
		fieldID 	= buff[i]
		fieldSize 	= buff[i + 1]
		fmt.Printf("field id: %d \n", int(fieldID))
		fmt.Printf("field size: %d \n", int(fieldSize))
		fmt.Printf("slice from: %d to %d\n", i+2, i+2+int(fieldSize))
		fmt.Printf("content: %v \n", buff[i+2:i+2+int(fieldSize)])
		content		:= buff[i+2:i+2+int(fieldSize)]
		fields 		= append(fields, Field{FieldID: fieldID, FieldSize: fieldSize, Content: content})
	}

	packet := Packet{
		PacketType: 	packetType,
		PacketSubtype: 	packetSubtype,
		Fields: 		fields,
	}

	fmt.Printf("%v: \n", packet)

	return packet
}
