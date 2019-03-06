package binary_test

import (
	"../../binary"
	"testing"
)

func TestEncodeDecode (t *testing.T) {
	number 		:= 123456789
	buff 		:= binary.Encode(number)
	result 		:= binary.Decode(buff)
	final   	:= 123456789

	if result != final {
		t.Error("Wrong converted")
	}
}