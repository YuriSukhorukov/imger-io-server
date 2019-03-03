package protocol

// упаковщик принимает на вход последовательность байт, тип сообщения,
// а возвращает пакет, готовый к отправке

func Pack(packet Packet) []byte {
	packetOffset 	:= 2
	fieldOffset 	:= 2
	fieldsLength 	:= 0

	for i := range packet.Fields {
		fieldsLength += fieldOffset
		fieldsLength += len(packet.Fields[i].Content)
	}

	fields := make([]byte, fieldsLength)
	for i := range packet.Fields {
		fields = append(fields, packet.Fields[i].FieldID)
		fields = append(fields, packet.Fields[i].FieldSize)
		fields = append(fields, packet.Fields[i].Content...)
	}

	buff := make([]byte, fieldsLength+ packetOffset + 1)

	buff[0] 			= packet.PacketType
	buff[1] 			= packet.PacketSubtype
	buff[2:] 			= fields
	buff[len(buff)-1] 	= 0x00

	return buff
}