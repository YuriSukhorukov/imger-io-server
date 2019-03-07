package binary

func EncodeInt(number int) []byte {
	cursor		:= number
	bytesCount 	:= 0
	for cursor != 0 {
		bytesCount++
		cursor >>= 8
	}
	buff := make([]byte, bytesCount)
	for i := 0; i < len(buff); i++ {
		buff[i] = byte(number)
		number >>= 8
	}
	return buff
}

func EncodeString(str string) []byte {
	buff := []byte(str)
	return buff
}