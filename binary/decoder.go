package binary

func DecodeInt(buff []byte) int {
	number := 0
	for i := len(buff) - 1; i >= 0; i-- {
		number |= int(buff[i]) << uint(i * 8)
	}
	return number
}

func DecodeString(buff []byte) string {
	return string(buff)
}
