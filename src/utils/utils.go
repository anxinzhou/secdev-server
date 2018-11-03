package utils

import "bytes"

func FillZero(src []byte, length int) []byte{
	prefix:= make([]byte,length,length+len(src))
	var buffer bytes.Buffer
	buffer.Write(prefix)
	buffer.Write(src)
	dst:=buffer.Bytes()
	return dst[len(dst)-length:]
}