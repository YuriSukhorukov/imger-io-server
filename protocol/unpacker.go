package protocol

// распаковщик принимает на вход последодвательность байт, коллекцию типов сообщений
// а возвращает пакет, готовый к использованию

func Unpack(buff []byte) Packet {
	packetType 		:= buff[0]
	packetSubtype 	:= buff[1]

	fields 			:= make([]Field, 0)

	var fieldID 	byte
	var fieldSize 	byte
	var content		[]byte

	for i := 2; i < len(buff) - 1; i += int(fieldSize) {
		fieldID 	= buff[i]
		fieldSize 	= buff[i+1]
		content		= buff[i + 2:fieldSize]
		fields 		= append(fields, Field{FieldID: fieldID, FieldSize: fieldSize, Content: content})
	}

	packet := Packet{
		PacketType: 	packetType,
		PacketSubtype: 	packetSubtype,
		Fields: 		fields,
	}

	return packet
}
