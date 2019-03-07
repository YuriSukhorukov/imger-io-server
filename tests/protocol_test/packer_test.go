package protocol_test

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