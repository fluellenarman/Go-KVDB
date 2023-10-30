package kv

import (
	"fmt"
	data "here/dataStruct"
	"log"
	"os"

	"gopkg.in/mgo.v2/bson"
)

const fileName string = "kvStore/data/data.db"

var file *os.File
var err error
var memory data.KVmap

func SerializeValue(key string, value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("SerializeValue() : int")
		fmt.Println("SerializeValue() -> ", key, value)
		put(key, IntToBytes(value.(int)), "int")
	case string:
		fmt.Println("SerializeValue() : string")
		fmt.Println("SerializeValue() -> ", key, value)
		put(key, StringToBytes(value.(string)), "string")
	default:
		fmt.Println("SerializeValue() : unknown type")
	}
}

func DeserializeValue(key string) interface{} {
	fmt.Println("DeserializeValue() -> ", key)
	fmt.Println(
		"DeserializeValue() -> memory.MemoMap", key, "->",
		memory.MemoMap[key])
	switch memory.MemoMap[key].DataType {
	case "int":
		fmt.Println("DeserializeValue() : int")
		return bytesToInt(memory.MemoMap[key].Value)
	case "string":
		fmt.Println("DeserializeValue() : string")
		return BytesToString(memory.MemoMap[key].Value)
	default:
		fmt.Println("DeserializeValue() : unknown type")
	}
	return nil
}

// Open file and load data into memory
func Init() {
	fmt.Println("kvStore/methods.go : Init()")

	file, err = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	errorCatch(err)

	buffer := make([]byte, 1024)
	n, _ := file.Read(buffer)
	// errorCatch(err)
	// fmt.Println(n)
	dataByte := buffer[:n]
	fmt.Println("methods.go : Init() dataBytes\n", dataByte)

	if (dataByte != nil) && (n != 0) {
		// load data into memory
		fmt.Println("loading data into memory")
		err = bson.Unmarshal(dataByte, &memory)
		errorCatch(err)
		fmt.Println("memory", memory)
	} else {
		// initialize memory
		fmt.Println("initializing data into memory")
		memory = data.KVmap{
			MemoMap: make(map[string]data.ValueTuple),
		}
	}
}

func writeToStorage(memory data.KVmap) {
	fmt.Println("kvStore/methods.go : writeToStorage()")
	fmt.Println("writeToStorage() memory\n", memory)
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
	fmt.Println("kvStore/methods.go : close()")
	file.Close()
}

// insert and update
func put(key string, value []byte, dataType string) {
	fmt.Println("kvStore/methods.go : put()")
	fmt.Println("put() putting", key, value)
	valueTuple := data.ValueTuple{
		DataType: dataType,
		Value:    value,
	}
	memory.MemoMap[key] = valueTuple
	fmt.Println("put() memory", memory)
}

func get() {

}

func delete() {

}

func errorCatch(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
