package protocol

// распаковщик принимает на вход последодвательность байт, коллекцию типов сообщений
// а возвращает пакет, готовый к использованию

func Unpack(buff []byte) Packet {
	packetType 		:= buff[0]
	packetSubtype 	:= buff[1]

	fields 			:= make([]Field, 0)

	var fieldID 	byte
	var fieldSize 	byte

	for i := 2; i < len(buff) - 1; i++ {
		fieldID 	= buff[i]
		fieldSize 	= buff[i+1]
		if fieldID == 0xAA && fieldSize == 0xAF {
			fields = append(fields, Field{FieldID: buff[i], FieldSize: buff[i+1]})
		} else {
			fields[len(fields)-1].Content = append(fields[len(fields)-1].Content, buff[i])
		}
	}

	packet := Packet{
		PacketType: 	packetType,
		PacketSubtype: 	packetSubtype,
		Fields: 		fields,
	}

	return packet
}