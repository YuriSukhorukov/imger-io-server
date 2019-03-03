package model

type Player struct {
	name string
	place uint
	points uint8
}

type Point struct {
	x uint16
	y uint16
}

type Message struct {
	name string
	text string
}

type Time struct {
	seconds uint16
}