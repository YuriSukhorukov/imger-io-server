package protocol_test

import (
	"fmt"
	"testing"
)
import "../../protocol"
import "../../model"

//func TestFoo(t *testing.T) {
//	result := protocol.Sum(1, 2)
//	fmt.Printf(string(result))
//	if result != 3 {
//		t.Error("wrong test", result)
//	}
//}

func TestByteInt(t *testing.T) {
	i := 5
	iByte := byte(i)
	result := int(iByte)
	if result != 5 {
		t.Error("wrong convert int from byte")
	}

	//jjj := uint(257)
	//lll := binary.Size(jjj)
	//fmt.Printf("%v:--- \n", lll)

	//var j uint64 = 255
	//buff := make([]byte, int(iByte))
	//buff[0] = byte(j)
	//buff[1] = byte(j >> 8)
	//buff[2] = byte(j >> 16)
	//buff[3] = byte(j >> 24)
	//buff[4] = byte(j >> 32)
	//fmt.Printf("%v: \n", buff)
	//fmt.Printf("%v \n", len(buff))
	//k := int(buff[4])<<32 | int(buff[3])<<24 | int(buff[2])<<16 | int(buff[1])<<8 | int(buff[0])
	//fmt.Printf("%v: \n", k)

	// определить размер поля в байтах
	// создать буфер с нужного размера
	// через битовые сдвиги заполнить масив байт
	// произвести обратную операцию через обратный сдвиг

	number 		:= 10000
	bytesCount 	:= 0
	cursor		:= number
	bitsCount 	:= 8

	for {
		if cursor == 0 {
			break
		} else {
			bytesCount++
			fmt.Printf("cursor: %v\n", cursor)
			cursor >>= 8
		}
	}
	fmt.Printf("bytes count: %v\n", bytesCount)

	buff 		:= make([]byte, bytesCount)
	for i := 0; i < len(buff); i++ {
		buff[i] = byte(number)
		number >>= 8
	}

	fmt.Printf("buff: %v\n", buff)


	restoredNumber 		:= 0

	for i := len(buff) - 1; i >= 0; i-- {
		restoredNumber |= int(buff[i]) << uint(i * bitsCount)
	}

	fmt.Printf("restored number: %v\n", restoredNumber)
}

func TestPack(t *testing.T) {
	//player := model.Player{Name: "John", Place: 10, Points: 10}
	point := model.Point{X: 2, Y: 5}

	//a := []byte(player.Name)
	//b := []byte{byte(player.Place)}
	//c := []byte{byte(player.Points)}
	//
	//field1 := protocol.Field{FieldID: 0xAA, FieldSize: 1, Content: a}
	//field2 := protocol.Field{FieldID: 0xAA, FieldSize: 2, Content: b}
	//field3 := protocol.Field{FieldID: 0xAA, FieldSize: 3, Content: c}
	//
	//fields := []protocol.Field{field1, field2, field3}
	fields := protocol.FieldsFromPoint(point)

	packet := protocol.Packet{
		PacketType: protocol.LINE_DOT,
		PacketSubtype: protocol.END_FIELD,
		Fields: fields,
	}

	buff := protocol.Pack(packet)


	var _ = packet
	var _ = buff
}