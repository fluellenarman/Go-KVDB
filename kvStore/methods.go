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

// Open file for writing
func Init() {
	fmt.Println("kvStore/methods.go : Init()")

	file, err = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	errorCatch(err)

	buffer := make([]byte, 1024)
	n, err := file.Read(buffer)
	errorCatch(err)
	dataByte := buffer[:n]
	// fmt.Println((string(dataByte)))
	err = bson.Unmarshal(dataByte, &memory)
	fmt.Println(memory)
}

// close file when done
func Close() {
	fmt.Println("kvStore/methods.go : close()")
	file.Close()
}

// insert and update
func put() {

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
