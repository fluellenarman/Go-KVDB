package kv

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
)

// int
func intToBytes(num int) []byte {
	byteArray := make([]byte, 4)
	binary.BigEndian.PutUint32(byteArray, uint32(num))
	return byteArray
}

func bytesToInt(byteArray []byte) int {
	return int(binary.BigEndian.Uint32(byteArray))
}

//

// int array
func intArrToBytes(arr []int) []byte {
	byteCount := len(arr) * 4 // Assuming 32-bit integers, hence 4 bytes each
	byteSlice := make([]byte, byteCount)
	for i := 0; i < len(arr); i++ {
		offset := i * 4 // Each integer takes 4 bytes
		binary.BigEndian.PutUint32(byteSlice[offset:offset+4], uint32(arr[i]))
	}

	return byteSlice
}

func bytesToIntArr(byteSlice []byte) []int {
	intCount := len(byteSlice) / 4 // Assuming 32-bit integers, hence 4 bytes each
	intArr := make([]int, intCount)
	for i := 0; i < intCount; i++ {
		offset := i * 4 // Each integer takes 4 bytes
		intArr[i] = int(binary.BigEndian.Uint32(byteSlice[offset : offset+4]))
	}

	return intArr
}

//

// string
func stringToBytes(str string) []byte {
	return []byte(str)
}

func bytesToString(byteSlice []byte) string {
	return string(byteSlice)
}

//

// string array
func stringArrToBytes(strs []string) []byte {
	//using json because it's simpler, but it's not the most efficient
	jsonStr, err := json.Marshal(strs)
	if err != nil {
		fmt.Println(err)
	}
	return jsonStr
}

func bytesToStringArr(byteSlice []byte) []string {
	// using json because it's simpler, but it's not the most efficient
	var strs []string
	err := json.Unmarshal(byteSlice, &strs)
	if err != nil {
		fmt.Println(err)
	}
	return strs
}

//

func Main() {
	fmt.Println("helperMethods.go : main()")
	// testingIntBytes()
	// testingIntArrBytes()
	// testingStringBytes()
	// testingStringArrBytes()
}

func testingStringBytes() {
	byteArr := stringToBytes("My Name is John 123 !@#")
	fmt.Println(byteArr)
	testStr := bytesToString(byteArr)
	fmt.Println(testStr)
}

func testingStringArrBytes() {
	strArr := []string{"My", "Name", "is", "John", "123", "!@#"}
	byteArr := stringArrToBytes(strArr)
	fmt.Println(byteArr)
	testStrArr := bytesToStringArr(byteArr)
	fmt.Println(testStrArr)
}

func testingIntBytes() {
	byteArr := intToBytes(2500)
	fmt.Println(byteArr)
	testInt := bytesToInt(byteArr)
	fmt.Println(testInt)
}

func testingIntArrBytes() {
	intArr := []int{1, 2, 3, 4, 5}
	byteArr := intArrToBytes(intArr)
	fmt.Println(byteArr)
	testIntArr := bytesToIntArr(byteArr)
	fmt.Println(testIntArr)
}
