package protocol

// упаковщик принимает на вход последовательность байт, тип сообщения,
// а возвращает пакет, готовый к отправке

func Pack(packet Packet) []byte {
	fields := make([]byte, 0)
	for i := range packet.Fields {
		fields = append(fields, packet.Fields[i].FieldID)
		fields = append(fields, packet.Fields[i].FieldSize)
		fields = append(fields, packet.Fields[i].Content...)
	}

	buff := make([]byte, 0)
	buff = append(buff, packet.PacketType)
	buff = append(buff, packet.PacketSubtype)
	buff = append(buff, fields...)
	buff = append(buff, 0x00)

	return buff
}