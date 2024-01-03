package client

import (
	"bytes"
	"fmt"
	data "here/dataStruct"
	kv "here/kvStore"
	"io"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

func encodeBSON(input data.KVpair) []byte {
	bsonData, err := bson.Marshal(input)
	if err != nil {
		fmt.Println("Error marshalling data")
	}

	return bsonData
}

// method, key, data
func SendReq(method string, param1 string, param2 string) {
	url := "http://localhost:8080/"

	var myKVpair data.KVpair
	// fmt.Println("param1", param1, "param2", param2)
	myKVpair.Key = param1
	myKVpair.Value = param2

	// fmt.Println("myKVpair", myKVpair)

	bsonData := encodeBSON(myKVpair)
	// fmt.Println("bsonData->", bsonData)

	var test data.KVpair
	err := bson.Unmarshal(bsonData, &test)
	if err != nil {
		fmt.Println("Error unmarshalling data")
	}
	// fmt.Println("test", test)

	fmt.Println("Sending request to server...")
	fmt.Println("method:", method, "url:", url)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(bsonData))
	errorHandling(err)

	resp, err := http.DefaultClient.Do(req)
	errorHandling(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	errorHandling(err)

	fmt.Println("Response:", string(body))
}

func sendMiscReq(method string) {
	url := "http://localhost:8080/"
	fmt.Println("In sendMiscReq :", method)
	req, err := http.NewRequest(method, url, nil)
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
