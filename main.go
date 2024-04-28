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
		// fmt.Println(os.Args[2])
		os.Setenv("IS_LOCAL", "false")
		serverMode()
	} else if os.Args[1] == "localServer" {
		fmt.Println("Running local server mode")
		os.Setenv("PORT", "8080")
		os.Setenv("NAME", "app1")
		os.Setenv("IS_LOCAL", "true")
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
	go kv.InitRaft()
	kv.Main()
}

func printError() {
	fmt.Println("Enter arguments: client or server")
	fmt.Println("client: run CLI tool for KVstore's client")
	fmt.Println("server: run KVstore server")
}
