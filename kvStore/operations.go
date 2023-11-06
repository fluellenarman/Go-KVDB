package kv

import (
	"fmt"
	data "here/dataStruct"
	"log"
	"os"

	"gopkg.in/mgo.v2/bson"
)

var fileName string

var file *os.File
var err error
var memory data.KVmap

// Open file and load data into memory
func Init(fileNameInput string) string {
	fmt.Println("operations.go : Init()")
	fileName = fileNameInput
	var state string

	file, err = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	errorCatch(err)

	buffer := make([]byte, 1024)
	n, err := file.Read(buffer)
	// errorCatch(err)
	// fmt.Println(n)
	dataByte := buffer[:n]
	fmt.Println("operations.go : Init() dataBytes\n", dataByte)

	if (dataByte != nil) && (n != 0) {
		// load data into memory
		fmt.Println("loading data into memory")
		err = bson.Unmarshal(dataByte, &memory)
		errorCatch(err)
		fmt.Println("memory", memory)
		state = "Loaded"
	} else {
		// initialize memory
		fmt.Println("initializing data into memory")
		memory = data.KVmap{
			MemoMap: make(map[string]interface{}),
		}
		state = "Initialized"
	}
	return state
}

func writeToStorage(memory data.KVmap) {
	file.Close()
	file, err = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	fmt.Println("operations.go : writeToStorage() memory\n", memory)
	//serialize memory before writing to storage
	bsonData, err := bson.Marshal(memory)
	errorCatch(err)

	fmt.Println("writeToStorage() bsonData\n", bsonData)
	n, _ := file.Write(bsonData)
	fmt.Println("writeToStorage() wrote ", n, " bytes to storage")

}

// close file and load memory into storage
func Close() {
	writeToStorage(memory)
	fmt.Println("operations.go : close()")
	file.Close()
}

// insert and update
func Put(key string, value interface{}) {
	fmt.Println("operations.go: Put()", key, value)
	memory.MemoMap[key] = value
}

func Get(key string) (interface{}, bool) {
	fmt.Println("operations.go: Get()", key)
	value, exists := memory.MemoMap[key]
	if exists == true {
		fmt.Println("Get() successful", key, value)
	} else {
		fmt.Println("Get() failed", key)
	}
	return value, exists
}

func DeletePair(key string) {
	fmt.Println("operations.go : delete()")
	fmt.Println("delete() deleting", key)
	delete(memory.MemoMap, key)
}

func errorCatch(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
