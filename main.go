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

	"gopkg.in/mgo.v2/bson"
)

func main() {
	fmt.Println("main.go")
	client.Init()
	kv.Init()

	testingBson()
	// testingJson()
	// testingMap()
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

	testData := data.Data{
		Key:   "Alice",
		Value: bsonData,
	}
	fmt.Println(testData)

	bsonData2, err := bson.Marshal(testData)

	fmt.Println(reflect.TypeOf(bsonData2), bsonData2)

	var deserialized2 data.Data
	err = bson.Unmarshal(bsonData2, &deserialized2)
	fmt.Println(deserialized2)

	var deserialized data.Person
	err = bson.Unmarshal(bsonData, &deserialized)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(deserialized)

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
