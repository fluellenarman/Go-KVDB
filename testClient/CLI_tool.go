package client

import (
	"fmt"
	"strings"
)

// TestInput set(key, value), get(key), delete(key)
// CLI_input is the CLI tool for the client
func CLI_input() string {
	var input string
	// fmt.Print("\n$ ")
	fmt.Scanln(&input)

	// fmt.Println("input:", input)
	extractInfo(input)
	return input
}

func extractInfo(input string) {
	// GET key1, SET key1 val1, DELETE key1
	// extract command - USER NOTES, be aware fof whitespaces
	fmt.Println(input)
	if strings.HasPrefix(input, "get(") && strings.HasSuffix(input, ")") {
		fmt.Println("is get command")
		fmt.Println(input[4 : len(input)-1])
		// key := input[4 : len(input)-1]
		// SendReq(key)
	} else if strings.HasPrefix(input, "set(") && strings.HasSuffix(input, ")") {
		fmt.Println("is set command")
		fmt.Println(input[4 : len(input)-1])
		// key := input[4 : len(input)-1]
		// SendReq(key)
	} else if strings.HasPrefix(input, "delete(") && strings.HasSuffix(input, ")") {
		fmt.Println("is delete command")
		fmt.Println(input[7 : len(input)-1])
		// key := input[4 : len(input)-1]
		// SendReq(key)
	} else {
		fmt.Println("invalid command")
	}
}
