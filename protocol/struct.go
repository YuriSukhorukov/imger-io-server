package protocol

type Packet struct {
	PacketType      byte
	PacketSubtype   byte
	Fields          []Field
}

type Field struct {
	FieldID         byte
	FieldSize       byte
	Content         []byte
}

type Queue struct {
	Packets         []Packet
}

// тип пакета
// подтип пакета
// ...
// id свойства
// размер свойства
// последовательность байт определяющих значение свойства
// конец свойства
// конец пакета

// для экономии траффика лучше агрегировать пакеты и отправлять четыре раза в секунду
