package utils

import (
	"bytes"
	"log"
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
	loc, err:= time.LoadLocation("Asia/Chongqing")
	if err!=nil {
		log.Println(err.Error())
	}
	// "2006-01-02 15:04:05 is the birth day of golang, fixed format"
	date:=time.Now().In(loc).Format("2006-01-02T15:04:05")
	return date
}