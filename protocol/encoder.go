package protocol

import "../model"
import "../binary"

func EncodePlayer(p model.Player) []Field {
	name 	:= p.Name
	place 	:= int(p.Place)
	points 	:= int(p.Points)

	a := binary.EncodeString(name)
	b := binary.EncodeInt(place)
	c := binary.EncodeInt(points)

	f1 := Field{FieldID: 0xAA, FieldSize: byte(len(a)), Content: a}
	f2 := Field{FieldID: 0xAA, FieldSize: byte(len(b)), Content: b}
	f3 := Field{FieldID: 0xAA, FieldSize: byte(len(c)), Content: c}

	fields := []Field{f1, f2, f3}

	return fields
}

func EncodePoint(p model.Point) []Field {
	a := []byte{byte(p.X)}
	b := []byte{byte(p.Y)}

	field1 := Field{FieldID: 0xAA, FieldSize: byte(len(a)), Content: a}
	field2 := Field{FieldID: 0xAA, FieldSize: byte(len(b)), Content: b}

	fields := []Field{field1, field2}

	return fields
}