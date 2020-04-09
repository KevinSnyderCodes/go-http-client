package ptr

func ByteSlice(value []byte) *[]byte {
	return &value
}

func MapStringString(value map[string]string) *map[string]string {
	return &value
}
