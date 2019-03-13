package protocol

import (
	"../../model"
	"../../protocol"
	"fmt"
	"reflect"
	"testing"
)

func TestEncodeDecodePoint(t *testing.T) {
	point 		:= model.Point{X: 11, Y: 22}
	fields 		:= protocol.EncodePoint(point)
	result 		:= protocol.DecodePoint(fields)
	final 		:= model.Point{X: 11, Y: 22}
	if !reflect.DeepEqual(result, final) {
		fmt.Printf("Wrong converted")
	}
}

func TestEncodeDecodePlayer(t *testing.T) {
	player 		:= model.Player{Name: "John", Place: 1, Points: 10}
	fields 		:= protocol.EncodePlayer(player)
	result 		:= protocol.DecodePlayer(fields)
	final 		:= model.Player{Name: "John", Place: 1, Points: 10}
	if !reflect.DeepEqual(result, final) {
		fmt.Printf("Wrong converted")
	}
}

func TestEncodeDecodeMessage(t *testing.T) {
	message 	:= model.Message{Name: "John", Text: "Hello everyone!"}
	fields 		:= protocol.EncodeMessage(message)
	result 		:= protocol.DecodeMessage(fields)
	final 		:= model.Message{Name: "John", Text: "Hello everyone!"}
	if !reflect.DeepEqual(result, final) {
		fmt.Printf("Wrong converted")
	}
}

func TestEncodeDecodeTime(t *testing.T) {
	time 		:= model.Time{Seconds: 1000}
	fields 		:= protocol.EncodeTime(time)
	result 		:= protocol.DecodeTime(fields)
	final 		:= model.Time{Seconds: 1000}
	if !reflect.DeepEqual(result, final) {
		fmt.Printf("Wrong converted")
	}
}

func TestPackUnpackPoint(t *testing.T) {
	point 		:= model.Point{X: 550150550, Y: 550150550}
	fields 		:= protocol.EncodePoint(point)
	packet 		:= protocol.Packet{
					PacketType: protocol.CANVAS,
					PacketSubtype: protocol.LINE_DOT,
					Fields: fields,
				}
	buff 		:= protocol.Pack(packet)
	result 		:= protocol.Unpack(buff)
	final 		:= protocol.Packet{
					PacketType: protocol.CANVAS,
					PacketSubtype: protocol.LINE_DOT,
					Fields: protocol.EncodePoint(model.Point{X: 550150550, Y: 550150550}),
				}
	if !reflect.DeepEqual(result, final) {
		fmt.Printf("Wrong converted")
	}
}

func TestPackUnpackPlayer(t *testing.T) {
	player 		:= model.Player{Name: "John Smith", Place: 11, Points: 22}
	fields 		:= protocol.EncodePlayer(player)
	packet 		:= protocol.Packet{
					PacketType: protocol.ENTITY,
					PacketSubtype: protocol.PLAYER,
					Fields: fields,
				}
	buff 		:= protocol.Pack(packet)
	result 		:= protocol.Unpack(buff)
	final 		:= protocol.Packet{
					PacketType: protocol.ENTITY,
					PacketSubtype: protocol.PLAYER,
					Fields: protocol.EncodePlayer(model.Player{Name: "John Smith", Place: 11, Points: 22}),
				}
	if !reflect.DeepEqual(result, final) {
		fmt.Printf("Wrong converted")
	}
}