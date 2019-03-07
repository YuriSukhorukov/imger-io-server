package binary_test

import (
	"../../binary"
	"testing"
)

func TestEncodeDecodeInt (t *testing.T) {
	number 		:= 123
	buff 		:= binary.EncodeInt(number)
	result 		:= binary.DecodeInt(buff)
	final   	:= 123

	if result != int(final) {
		t.Error("Wrong converted")
	}
}

func TestEncodeDecodeString (t *testing.T) {
	str 		:= "Hi, there!"
	buff 		:= binary.EncodeString(str)
	result 		:= binary.DecodeString(buff)
	final   	:= "Hi, there!"

	if result != final {
		t.Error("Wrong converted")
	}
}