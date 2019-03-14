package protocol

import (
	"../model"
)
import "../binary"

func EncodePlayer(p model.Player) []Field {
	b1 		:= binary.EncodeString(p.Name)
	b2 		:= binary.EncodeInt(int(p.Place))
	b3		:= binary.EncodeInt(int(p.Points))
	f1 		:= Field{FieldID: 0xAA, FieldSize: byte(len(b1)), Content: b1}
	f2 		:= Field{FieldID: 0xAA, FieldSize: byte(len(b2)), Content: b2}
	f3 		:= Field{FieldID: 0xAA, FieldSize: byte(len(b3)), Content: b3}
	return 		[]Field{f1, f2, f3}
}

func EncodePoint(p model.Point) []Field {
	b1 		:= binary.EncodeInt(int(p.X))
	b2 		:= binary.EncodeInt(int(p.Y))
	f1 		:= Field{FieldID: 0xAA, FieldSize: byte(len(b1)), Content: b1}
	f2 		:= Field{FieldID: 0xAA, FieldSize: byte(len(b2)), Content: b2}
	return 	    	[]Field{f1, f2}
}

func EncodeMessage(m model.Message) []Field {
	b1 		:= binary.EncodeString(m.Name)
	b2 		:= binary.EncodeString(m.Text)
	f1 		:= Field{FieldID: 0xAA, FieldSize: byte(len(b1)), Content: b1}
	f2 		:= Field{FieldID: 0xAA, FieldSize: byte(len(b2)), Content: b2}
	return 	    	[]Field{f1, f2}
}

func EncodeTime(t model.Time) []Field {
	b1 		:= binary.EncodeInt(int(t.Seconds))
	f1 		:= Field{FieldID: 0xAA, FieldSize: byte(len(b1)), Content: b1}
	return 	    	[]Field{f1}
}
