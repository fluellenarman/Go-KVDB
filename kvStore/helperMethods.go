package kv

import (
	"encoding/binary"
	"fmt"
)

func intToBytes(num int) []byte {
	byteArray := make([]byte, 4)
	binary.BigEndian.PutUint32(byteArray, uint32(num))
	return byteArray
}

func bytesToInt(byteArray []byte) int {
	return int(binary.BigEndian.Uint32(byteArray))
}

func Main() {
	fmt.Println("helperMethods.go : main()")
	byteArr := intToBytes(42 * 16)
	fmt.Println(byteArr)
	testInt := bytesToInt(byteArr)
	fmt.Println(testInt)
}
