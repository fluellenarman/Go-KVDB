package client

import (
	"bytes"
	"fmt"
	data "here/dataStruct"
	kv "here/kvStore"
	"io"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

func encodeBSON(input data.KVpair) []byte {
	bsonData, err := bson.Marshal(input)
	if err != nil {
		fmt.Println("Error marshalling data")
	}

	return bsonData
}

func extractParams(input string) (string, string) {
	// Extract command(key, valueType, value)
	// extract command - USER NOTES, be aware fof whitespaces
	var command, params string
	fmt.Println(input)
	if strings.HasPrefix(input, "get(") && strings.HasSuffix(input, ")") {
		fmt.Println("is get command")
		// fmt.Println(input[4 : len(input)-1])
		command = "get"
		params = input[4 : len(input)-1]
	}

	return command, params
}

func SendReq() {
	url := "http://localhost:8080/"

	var data data.KVpair
	data.Key = "key1"
	data.Value = "Hello World From client.go!"

	bsonData := encodeBSON(data)
	// fmt.Println(bsonData)

	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(bsonData))
	errorHandling(err)

	resp, err := http.DefaultClient.Do(req)
	errorHandling(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	errorHandling(err)

	fmt.Println("Response:", string(body))
}

func Set(key string, value interface{}) {
	fmt.Println("client.go : Set()", key, value)
	kv.Put(key, value)
}

func Get(key string) (interface{}, bool) {
	fmt.Println("client.go : Get()", key)
	value, exists := kv.Get(key)
	// fmt.Println("client.go : Get() received -> ", value)
	return value, exists
}

func Delete(key string) {
	fmt.Println("client.go : Delete()")
	fmt.Println(key)
	kv.DeletePair(key)
	fmt.Println("client.go : Delete() completed")
}

func Init() {
	fmt.Println("client.go : Init()")
}

func errorHandling(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}
