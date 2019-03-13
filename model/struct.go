package model

type Player struct {
	Name string
	Place uint8
	Points uint8
}

type Point struct {
	X uint
	Y uint
}

type Message struct {
	Name string
	Text string
}

type Time struct {
	Seconds uint
}