package protocol

import (
	"../binary"
	"../model"
)

func DecodePlayer(f []Field) model.Player {
	name 		:= binary.DecodeString(f[0].Content)
	place 		:= binary.DecodeInt(f[1].Content)
	points 		:= binary.DecodeInt(f[2].Content)
	return 		model.Player{Name: name, Place: uint8(place), Points: uint8(points)}
}

func DecodePoint(f []Field) model.Point {
	x           := binary.DecodeInt(f[0].Content)
	y           := binary.DecodeInt(f[1].Content)
	return      model.Point{X: uint(x), Y: uint(y)}
}

func DecodeMessage(f []Field) model.Message {
	name 		:= binary.DecodeString(f[0].Content)
	text 		:= binary.DecodeString(f[1].Content)
	return 		model.Message{Name: name, Text: text}
}

func DecodeTime(f []Field) model.Time {
	seconds 	:= binary.DecodeInt(f[0].Content)
	return 		model.Time{Seconds: uint(seconds)}
}
