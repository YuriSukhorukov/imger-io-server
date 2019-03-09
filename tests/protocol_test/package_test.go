package protocol

import (
	"../../model"
	"../../protocol"
	"fmt"
	"reflect"
	"testing"
)

func TestEncodeDecodePlayer(t *testing.T) {
	player 	:= model.Player{Name: "John", Place: 1, Points: 10}
	fields 	:= protocol.EncodePlayer(player)
	result 	:= protocol.DecodePlayer(fields)
	final 	:= model.Player{Name: "John", Place: 1, Points: 10}

	if !reflect.DeepEqual(result, final) {
		fmt.Printf("Wrong converted")
	}
}

func TestEncodeDecodePoint(t *testing.T) {
	point := model.Point{X: 11, Y: 22}
	fields := protocol.EncodePoint(point)
	result := protocol.DecodePoint(fields)
	final := model.Point{X: 11, Y: 22}

	if !reflect.DeepEqual(result, final) {
		fmt.Printf("Wrong converted")
	}
}

//func TestPackUnpackPlayer(t *testing.T) {
//	player 	:= model.Player{Name: "John", Place: 1, Points: 10}
//	fields 	:= protocol.EncodePlayer(player)
//	packet  := protocol.Packet{
//		PacketType: protocol.ENTITY,
//		PacketSubtype: protocol.PLAYER,
//		Fields: fields,
//	}
//
//	buff 	:= protocol.Pack(packet)
//
//	fmt.Printf("packet: %v\n", packet)
//	fmt.Printf("packed buff: %v\n", buff)
//
//	if !reflect.DeepEqual(result, final) {
//		fmt.Printf("Wrong converted")
//	}
//}