package binary_test

import (
	"../../binary"
	"reflect"
	"testing"
)

func TestEncode (t *testing.T) {
	number		:= 510
	final 		:= []byte{254, 1}
	result 		:= binary.Encode(number)

	if !reflect.DeepEqual(result, final) {
		t.Error("Wrong converted")
	}
}