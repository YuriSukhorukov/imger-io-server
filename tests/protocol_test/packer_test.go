package protocol_test

import (
	"fmt"
	"testing"
)
import "../../protocol"
import "../../model"

func TestFoo(t *testing.T) {
	result := protocol.Sum(1, 2)
	fmt.Printf(string(result))
	if result != 3 {
		t.Error("wrong test", result)
	} else {
		fmt.Printf("correct")
	}
}

func TestPack(t *testing.T) {
	player := model.Player{Name: "John", Place: 10, Points: 10}

	a := []byte(player.Name)
	b := []byte{byte(player.Place)}
	c := []byte{byte(player.Points)}

	field1 := protocol.Field{FieldID: 0xAA, FieldSize: 1, Content: a}
	field2 := protocol.Field{FieldID: 0xAA, FieldSize: 2, Content: b}
	field3 := protocol.Field{FieldID: 0xAA, FieldSize: 3, Content: c}

	fields := []protocol.Field{field1, field2, field3}

	packet := protocol.Packet{
		PacketType: protocol.LINE_DOT,
		PacketSubtype: protocol.END_FIELD,
		Fields: fields,
	}
}