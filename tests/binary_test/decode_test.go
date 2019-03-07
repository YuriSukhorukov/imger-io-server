package binary_test

import (
	"../../binary"
	"testing"
)

func TestDecode (t *testing.T) {
	buff 			:= []byte{4, 1}
	result 			:= binary.DecodeInt(buff)
	final 			:= 260

	if result != final {
		t.Error("Wrong converted")
	}
}