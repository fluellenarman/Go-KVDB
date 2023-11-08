// GOAL: Create a simple DBMS that can be used to store key-value pairs
// GO will implemenent the kv store
// TKINTER will implement the client GUI
// SPECIFICATIONS: Persistence, Concurrency, Transactions
package main

import (
	"fmt"
	kv "here/kvStore"
	client "here/testClient"
	// client "here/testClient"
)

const fileName string = "kvStore/data/test2.bson"

func main() {
	fmt.Println("main.go")
	// kv.Main()

	kv.Init(fileName)
	client.Set("key1", "Hello World")
	client.Set("key2", 123)
	client.Set("key3", 321)
	client.Set("key4", "Goodbye World")
	client.Get("key1")
	client.Get("key2")
	kv.Close()
}
