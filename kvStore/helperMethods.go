package kv

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
)

// IntToBytes converts an integer to a byte array.
func IntToBytes(num int) []byte {
	byteArray := make([]byte, 4)
	binary.BigEndian.PutUint32(byteArray, uint32(num))
	return byteArray
}

// bytesToInt converts a byte array to an integer.
func bytesToInt(byteArray []byte) int {
	return int(binary.BigEndian.Uint32(byteArray))
}

// intArrToBytes converts an integer array to a byte array.
func intArrToBytes(arr []int) []byte {
	byteCount := len(arr) * 4 // Assuming 32-bit integers, hence 4 bytes each
	byteSlice := make([]byte, byteCount)
	for i := 0; i < len(arr); i++ {
		offset := i * 4 // Each integer takes 4 bytes
		binary.BigEndian.PutUint32(byteSlice[offset:offset+4], uint32(arr[i]))
	}

	return byteSlice
}

// bytesToIntArr converts a byte array to an integer array.
func bytesToIntArr(byteSlice []byte) []int {
	intCount := len(byteSlice) / 4 // Assuming 32-bit integers, hence 4 bytes each
	intArr := make([]int, intCount)
	for i := 0; i < intCount; i++ {
		offset := i * 4 // Each integer takes 4 bytes
		intArr[i] = int(binary.BigEndian.Uint32(byteSlice[offset : offset+4]))
	}

	return intArr
}

// StringToBytes converts a string to a byte array.
func StringToBytes(str string) []byte {
	fmt.Println("StringToBytes- ", str)
	return []byte(str)
}

// BytesToString converts a byte array to a string.
func BytesToString(byteSlice []byte) string {
	return string(byteSlice)
}

// stringArrToBytes converts a string array to a byte array.
func stringArrToBytes(strs []string) []byte {
	//using json because it's simpler, but it's not the most efficient
	jsonStr, err := json.Marshal(strs)
	if err != nil {
		fmt.Println(err)
	}
	return jsonStr
}

// bytesToStringArr converts a byte array to a string array.
func bytesToStringArr(byteSlice []byte) []string {
	// using json because it's simpler, but it's not the most efficient
	var strs []string
	err := json.Unmarshal(byteSlice, &strs)
	if err != nil {
		fmt.Println(err)
	}
	return strs
}

// Main is the main function for helperMethods.go.
func Main() {
	fmt.Println("helperMethods.go : main()")
	// testingIntBytes()
	// testingIntArrBytes()
	// testingStringBytes()
	// testingStringArrBytes()
}

// testingStringBytes is a helper function to test stringToBytes and bytesToString functions.
func testingStringBytes() {
	byteArr := StringToBytes("My Name is John 123 !@#")
	fmt.Println(byteArr)
	testStr := BytesToString(byteArr)
	fmt.Println(testStr)
}

// testingStringArrBytes is a helper function to test stringArrToBytes and bytesToStringArr functions.
func testingStringArrBytes() {
	strArr := []string{"My", "Name", "is", "John", "123", "!@#"}
	byteArr := stringArrToBytes(strArr)
	fmt.Println(byteArr)
	testStrArr := bytesToStringArr(byteArr)
	fmt.Println(testStrArr)
}

// testingIntBytes is a helper function to test IntToBytes and bytesToInt functions.
func testingIntBytes() {
	byteArr := IntToBytes(2500)
	fmt.Println(byteArr)
	testInt := bytesToInt(byteArr)
	fmt.Println(testInt)
}

// testingIntArrBytes is a helper function to test intArrToBytes and bytesToIntArr functions.
func testingIntArrBytes() {
	intArr := []int{1, 2, 3, 4, 5}
	byteArr := intArrToBytes(intArr)
	fmt.Println(byteArr)
	testIntArr := bytesToIntArr(byteArr)
	fmt.Println(testIntArr)
}
