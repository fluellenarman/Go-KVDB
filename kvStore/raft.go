package kv

import (
	"fmt"
)

type RaftState struct {
	nodeName string
	term     int
	timeout  int
	// votedFor string
	log []LogEntry
}

type LogEntry struct {
	term    int
	command string
}

var state RaftState

func InitRaft() {
	state = RaftState{
		nodeName: "node1",
		term:     0,
		timeout:  100,
		// votedFor: "",
		log: []LogEntry{},
	}
	fmt.Println("initRaft() state", state)
}

func appendLogEntry(command string) {
	var entry = LogEntry{
		term:    state.term,
		command: command,
	}
	state.log = append(state.log, entry)
}
