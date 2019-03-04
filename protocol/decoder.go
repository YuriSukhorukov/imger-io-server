package protocol

import (
	"../model"
)

func PlayerFromFields(f []Field) model.Player {
	name := f[0].Content
	place := f[1].Content
	points := f[2].Content

	player := model.Player{Name: string(name), Place: place[0], Points: points[0]}

	return player
}

func PointFromFields(f []Field) model.Point {
	x := f[0].Content
	y := f[1].Content

	point := model.Point{X: uint16(x[0]), Y: uint16(y[0])}

	return point
}