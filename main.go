// GOAL: Create a simple DBMS that can be used to store key-value pairs
// GO will implemenent the kv store
// TKINTER will implement the client GUI
// SPECIFICATIONS: Persistence, Concurrency, Transactions
package main

import (
	"fmt"
	kv "here/kvStore"
	client "here/testClient"
	"os"
	// client "here/testClient"
)

func main() {
	fmt.Println("\nmain.go:", os.Args[1:], len(os.Args))
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
	// kv.Init("kvStore/data/test2.bson")
	input := client.CLI_input()
	for input != "exit" {
		input = client.CLI_input()
	}
}

func serverMode() {
	go kv.Init("kvStore/data/test2.bson")
	kv.Main()
}

func printError() {
	fmt.Println("Enter arguments: client or server")
	fmt.Println("client: run CLI tool for KVstore's client")
	fmt.Println("server: run KVstore server")
}
