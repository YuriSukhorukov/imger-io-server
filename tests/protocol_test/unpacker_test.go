package protocol_test

import (
	"../../model"
	"../../protocol"
	"fmt"
	"reflect"
	"testing"
)

func TestUnpackPoint(t *testing.T) {
	buff 	:= []byte{1, 4, 170, 1, 11, 170, 1, 22, 0}

	result 	:= protocol.Unpack(buff)
	final	:= protocol.Packet{
		PacketType: protocol.CANVAS,
		PacketSubtype: protocol.LINE_DOT,
		Fields: protocol.EncodePoint(model.Point{X: 11, Y: 22}),
	}

	point := protocol.DecodePoint(result.Fields)
	_ = point

	if !reflect.DeepEqual(result, final) {
		fmt.Printf("Wrong converted")
	}
}

func TestUnpackPlayer(t *testing.T) {
	buff 	:= []byte{7, 8, 170, 4, 74, 111, 104, 110, 170, 1, 19, 170, 1, 21, 0}

	result 	:= protocol.Unpack(buff)
	final	:= protocol.Packet{
		PacketType: protocol.ENTITY,
		PacketSubtype: protocol.PLAYER,
		Fields: protocol.EncodePlayer(model.Player{Name: "John", Place: 19, Points: 21}),
	}

	player := protocol.DecodePlayer(result.Fields)
	_ = player

	if !reflect.DeepEqual(result, final) {
		fmt.Printf("Wrong converted")
	}
}