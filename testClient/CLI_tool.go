package client

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TestInput set(key, value), get(key), delete(key)
// CLI_input is the CLI tool for the client
func CLI_input() string {
	var input string
	// fmt.Print("\n$ ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
		extractInfo(input)
	}

	// fmt.Println("input:", input)
	// extractInfo(input)
	return input
}

func extractInfo(input string) {
	// GET key1, SET key1 dataType val1, DELETE key1
	// extract command - USER NOTES, be aware of whitespaces
	command := strings.Fields(input)
	fmt.Println("command->", command)
	fmt.Println(command[0])
	switch command[0] {
	case "get":
		fmt.Println("in get, 1 param key->", command[1])
		SendReq("GET", command[1], "")
	case "set":
		// CLI set can only set strings for now
		fmt.Println("in set, 2 params")
		fmt.Println(command[2])
		SendReq("POST", command[1], command[2])
	case "delete":
		fmt.Println("in delete, 1 param")
		SendReq("DELETE", command[1], "")
	}
}
