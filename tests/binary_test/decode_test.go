package binary_test

import (
	"../../binary"
	"testing"
)

func TestDecode (t *testing.T) {
	buff 			:= []byte{4, 1}
	final 			:= 260
	result 			:= binary.Decode(buff)

	if result != final {
		t.Error("Wrong converted")
	}
}