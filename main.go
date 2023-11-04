// GOAL: Create a simple DBMS that can be used to store key-value pairs
// GO will implemenent the kv store
// TKINTER will implement the client GUI
// SPECIFICATIONS: Persistence, Concurrency, Transactions
package main

import (
	"encoding/json"
	"fmt"
	data "here/dataStruct"
	kv "here/kvStore"
	client "here/testClient"
	"os"

	// client "here/testClient"
	"reflect"

	"gopkg.in/mgo.v2/bson"
)

func main() {
	fmt.Println("main.go")
	// kv.Main()

	// kv.SerializeValue([]int{1, 2, 3})
	kv.Init()
	client.Set("key1", "Hello World")
	client.Set("key2", 123)
	client.Get("key1")
	client.Get("key2")
	kv.Close()
}

func testingKVmap() {
	fmt.Println("\nmain.go : testingByteArrayMap() ")
	fmt.Println("KVmap writing and loading")

	testData1 := data.ValueTuple{
		DataType: "string",
		Value:    []byte("Hello world"),
	}
	fmt.Println("data being loaded onto KVmap in memory", testData1)

	testKVmap := data.KVmap{
		MemoMap: make(map[string]data.ValueTuple),
	}

	testKVmap.MemoMap["key1"] = testData1
	fmt.Println(testKVmap.MemoMap["key1"])
	fmt.Println(string(testKVmap.MemoMap["key1"].Value))
}

func testingJson() {
	fmt.Println("\nmain.go : testingJson() ")
	// var myMap map[string]string
	alice := data.Person{
		Name: "Alice",
		Age:  21,
		Sex:  "Female",
	}
	fmt.Println(alice)

	jsonData, err := json.Marshal(alice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reflect.TypeOf(jsonData), jsonData)
	var deserialized data.Person
	err = json.Unmarshal(jsonData, &deserialized)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(deserialized)
	fmt.Println("\n next test")

	jsonData2, err := json.Marshal("Test string 2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jsonData2)
	var deserialized2 string
	err = json.Unmarshal(jsonData2, &deserialized2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(deserialized2)
}

func testingBson() {
	fmt.Println("\nmain.go : testingBson() ")
	alice := data.Person{
		Name: "Alice",
		Age:  21,
		Sex:  "Female",
	}
	fmt.Println(alice)

	bsonData, err := bson.Marshal(alice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reflect.TypeOf(bsonData), bsonData)

	testData := data.ValueTuple{
		DataType: "Alice",
		Value:    bsonData,
	}
	fmt.Println(testData)

	bsonData2, err := bson.Marshal(testData)
	file, err := os.OpenFile("kvStore/data/data.db", os.O_RDWR|os.O_CREATE, 0644)
	file.Write(bsonData2)
	fmt.Println(reflect.TypeOf(bsonData2), bsonData2)

	var deserialized2 data.ValueTuple
	err = bson.Unmarshal(bsonData2, &deserialized2)
	fmt.Println(deserialized2)

	var deserialized data.Person
	err = bson.Unmarshal(bsonData, &deserialized)
	if err != nil {
		fmt.Println(err)
	}
}

func testingMap() {
	fmt.Println("\nmain.go : testingMap() ")
	// declaration and initialization
	myMap := make(map[string]string)

	// adding elements
	myMap["name"] = "Alice"
	myMap["age"] = "21"

	fmt.Println(myMap)
	// deleting elements
	delete(myMap, "age")
	fmt.Println(myMap)
	fmt.Println(myMap["age"] == "")
}
