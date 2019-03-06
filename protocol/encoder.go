package protocol

import (
	"../model"
	"fmt"
	"unsafe"
)

func FieldsFromPlayer(p model.Player) []Field {
	a := []byte(p.Name)
	b := []byte{byte(p.Place)}
	c := []byte{byte(p.Points)}

	field1 := Field{FieldID: 0xAA, FieldSize: 1, Content: a}
	field2 := Field{FieldID: 0xAA, FieldSize: 2, Content: b}
	field3 := Field{FieldID: 0xAA, FieldSize: 3, Content: c}

	fields := []Field{field1, field2, field3}

	return fields
}

func FieldsFromPoint(p model.Point) []Field {
	a := make([]byte, unsafe.Sizeof(p.X))
	b := make([]byte, unsafe.Sizeof(p.Y))

	field1 := Field{FieldID: 0xAA, FieldSize: byte(len(a)), Content: a}
	field2 := Field{FieldID: 0xAA, FieldSize: byte(len(a)), Content: b}
	fmt.Printf("%d\n", len(a))
	fields := []Field{field1, field2}

	return fields
}