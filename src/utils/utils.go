package utils

import (
	"bytes"
	"strings"
	"time"
)

func FillZero(src []byte, length int) []byte{
	prefix:= make([]byte,length,length+len(src))
	var buffer bytes.Buffer
	buffer.Write(prefix)
	buffer.Write(src)
	dst:=buffer.Bytes()
	return dst[len(dst)-length:]
}

func GetCurrentTime() string {
	// "2006-01-02 15:04:05 is the birth day of golang, fixed format"
	date:=time.Now().Format("2006-01-02 15:04:05")
	tmp:=strings.Split(date," ")
	date = strings.Join(tmp,"T")
	return date
}