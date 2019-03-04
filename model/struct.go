package model

type Player struct {
	Name string
	Place uint8
	Points uint8
}

type Point struct {
	X uint16
	Y uint16
}

type Message struct {
	Name string
	Text string
}

type Time struct {
	Seconds uint16
}