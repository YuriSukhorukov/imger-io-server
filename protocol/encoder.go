package protocol

import (
	"../model"
)
import "../binary"

func EncodePlayer(p model.Player) []Field {
	name 		:= binary.EncodeString(p.Name)
	place 		:= binary.EncodeInt(int(p.Place))
	points		:= binary.EncodeInt(int(p.Points))
	f1 			:= Field{FieldID: 0xAA, FieldSize: byte(len(name)), Content: name}
	f2 			:= Field{FieldID: 0xAA, FieldSize: byte(len(place)), Content: place}
	f3 			:= Field{FieldID: 0xAA, FieldSize: byte(len(points)), Content: points}
	return 		[]Field{f1, f2, f3}
}

func EncodePoint(p model.Point) []Field {
	x 			:= binary.EncodeInt(int(p.X))
	y 			:= binary.EncodeInt(int(p.Y))
	f1 			:= Field{FieldID: 0xAA, FieldSize: byte(len(x)), Content: x}
	f2 			:= Field{FieldID: 0xAA, FieldSize: byte(len(y)), Content: y}
	return 	    []Field{f1, f2}
}

func EncodeMessage(m model.Message) []Field {
	name 		:= binary.EncodeString(m.Name)
	text 		:= binary.EncodeString(m.Text)
	f1 			:= Field{FieldID: 0xAA, FieldSize: byte(len(name)), Content: name}
	f2 			:= Field{FieldID: 0xAA, FieldSize: byte(len(text)), Content: text}
	return 	    []Field{f1, f2}
}

func EncodeTime(t model.Time) []Field {
	seconds 	:= binary.EncodeInt(int(t.Seconds))
	f1 			:= Field{FieldID: 0xAA, FieldSize: byte(len(seconds)), Content: seconds}
	return 	    []Field{f1}
}