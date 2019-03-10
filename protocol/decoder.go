package protocol

import (
	"../binary"
	"../model"
)

func DecodePlayer(f []Field) model.Player {
	a := f[0].Content
	b := f[1].Content
	c := f[2].Content

	name 	:= binary.DecodeString(a)
	place 	:= binary.DecodeInt(b)
	points 	:= binary.DecodeInt(c)

	player := model.Player{Name: name, Place: uint8(place), Points: uint8(points)}

	return player
}

func DecodePoint(f []Field) model.Point {
	x := f[0].Content
	y := f[1].Content

	point := model.Point{X: uint(x[0]), Y: uint(y[0])}

	return point
}