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