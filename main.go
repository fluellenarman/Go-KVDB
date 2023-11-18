// GOAL: Create a simple DBMS that can be used to store key-value pairs
// GO will implemenent the kv store
// TKINTER will implement the client GUI
// SPECIFICATIONS: Persistence, Concurrency, Transactions
package main

import (
	"fmt"
	client "here/testClient"
	"os"
	// client "here/testClient"
)

func main() {
	fmt.Println("main.go")
	fmt.Println(os.Args[1:])
	fmt.Println(len(os.Args))
	if len(os.Args) != 2 {
		printError()
	} else if os.Args[1] == "client" {
		fmt.Println("Running client mode")
		clientMode()
	} else if os.Args[1] == "server" {
		fmt.Println("Running server mode")
		serverMode()
	} else {
		printError()
	}
}

func clientMode() {
	// go kv.Main()
	input := client.CLI_input()
	for input != "exit" {
		fmt.Println("client()", input)
		input = client.CLI_input()
	}
	// client.SendReq()
}

func serverMode() {
	// kv.Main()
}

func testClientMode() {
	// kv.Init(fileName)
	// client.Set("key1", "Hello World")
	// client.Set("key2", 123)
	// client.Set("key3", 321)
	// client.Set("key4", "Goodbye World")
	// client.Get("key1")
	// client.Get("key2")
	// kv.Close()
}

func printError() {
	fmt.Println("Enter arguments: client or server")
	fmt.Println("client: run CLI tool for KVstore's client")
	fmt.Println("server: run KVstore server")
}
