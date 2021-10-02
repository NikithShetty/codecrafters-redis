package utils

import "fmt"

var debugLevel = 0

func CheckError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func OkRespByte(s string) []byte {
	return []byte("+" + s + "\r\n")
}

func ErrRespByte(err string) []byte {
	return []byte("-Error " + err + "\r\n")
}

func LogInfo(any ...interface{}) {
	if debugLevel <= 0 {
		fmt.Println(any...)
	}
}

func LogError(any ...interface{}) {
	if debugLevel <= 0 {
		fmt.Println(any...)
	}
}
