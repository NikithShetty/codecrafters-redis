package utils

import (
	"fmt"
	"os"
)

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func OkRespByte(s string) []byte {
	return []byte("+" + s + "\r\n")
}

func ErrRespByte(err string) []byte {
	return []byte("-Error " + err + "\r\n")
}
