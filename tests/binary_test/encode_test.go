package binary_test

import (
	"../../binary"
	"reflect"
	"testing"
)

func TestEncode (t *testing.T) {
	number		:= 510
	result 		:= binary.Encode(number)
	final 		:= []byte{254, 1}

	if !reflect.DeepEqual(result, final) {
		t.Error("Wrong converted")
	}
}