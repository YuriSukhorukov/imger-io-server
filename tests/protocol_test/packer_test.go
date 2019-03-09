package protocol_test

import (
"../../model"
"../../protocol"
"fmt"
"reflect"
"testing"
)

func TestPackPoint(t *testing.T) {
	point 	:= model.Point{X: 11, Y: 22}
	fields 	:= protocol.EncodePoint(point)
	packet  := protocol.Packet{
		PacketType: protocol.CANVAS,
		PacketSubtype: protocol.LINE_DOT,
		Fields: fields,
	}
	result 	:= protocol.Pack(packet)
	final	:= []byte{1, 4, 170, 1, 11, 170, 1, 22, 0}

	if !reflect.DeepEqual(result, final) {
		fmt.Printf("Wrong converted")
	}
}

func TestPackPlayer(t *testing.T) {
	player 	:= model.Player{Name: "John", Place: 19, Points: 21}
	fields 	:= protocol.EncodePlayer(player)
	packet  := protocol.Packet{
		PacketType: protocol.ENTITY,
		PacketSubtype: protocol.PLAYER,
		Fields: fields,
	}
	result 	:= protocol.Pack(packet)
	final	:= []byte{7, 8, 170, 4, 74, 111, 104, 110, 170, 1, 19, 170, 1, 21, 0}

	if !reflect.DeepEqual(result, final) {
		fmt.Printf("Wrong converted")
	}
}