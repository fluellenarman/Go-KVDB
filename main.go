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
	"reflect"
)

func main() {
	fmt.Println("main.go")
	client.Init()
	kv.Init()
	// testingJson()
	testingMap()
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
}

func testingMap() {
	fmt.Println("\nmain.go : testingMap() ")
	myMap := make(map[string]string)
	myMap["name"] = "Alice"
	myMap["age"] = "21"
	fmt.Println(myMap)
}
