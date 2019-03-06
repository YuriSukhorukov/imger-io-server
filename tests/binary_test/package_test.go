package binary_test

import (
	"../../binary"
	"testing"
)

func TestEncodeDecode (t *testing.T) {
	number 		:= 123456789
	buff 		:= binary.Encode(number)
	final   	:= 123456789
	result 		:= binary.Decode(buff)

	if result != final {
		t.Error("Wrong converted")
	}
}