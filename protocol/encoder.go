package model

import "./"
import "../binary"

func FieldsFromPlayer(p model.Player) []Field {
	name 	:= p.Name
	place 	:= int(p.Place)
	points 	:= int(p.Points)

	a := binary.EncodeString(name)
	b := binary.EncodeInt(place)
	c := binary.EncodeInt(points)

	f1 := Field{FieldID: 0xAA, FieldSize: 1, Content: a}
	f2 := Field{FieldID: 0xAA, FieldSize: 2, Content: b}
	f3 := Field{FieldID: 0xAA, FieldSize: 3, Content: c}

	fields := []Field{f1, f2, f3}

	return fields
}

func FieldsFromPoint(p model.Point) []Field {
	a := []byte{byte(p.X)}
	b := []byte{byte(p.Y)}

	field1 := Field{FieldID: 0xAA, FieldSize: 1, Content: a}
	field2 := Field{FieldID: 0xAA, FieldSize: 2, Content: b}

	fields := []Field{field1, field2}

	return fields
}