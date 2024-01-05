package kv

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	loadLogToMemory()
	fmt.Println("loaded log to memory\n", state.log)
}

func appendLogEntry(command string) {
	var entry = LogEntry{
		term:    state.term,
		command: command,
	}
	state.log = append(state.log, entry)
}

func writeLogToStorage() {
	fmt.Println("writeLogToStorage() state.log", state.log)
	file, err := os.OpenFile("kvStore/data/logfile.txt", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file")
	}
	defer file.Close()

	lastEntry := state.log[len(state.log)-1]
	logString := fmt.Sprintf("%d %s\n", lastEntry.term, lastEntry.command)

	_, err = file.WriteString(logString)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func loadLogToMemory() {
	file, err := os.OpenFile("kvStore/data/logfile.txt", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("Error opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)
		entryTerm, _ := strconv.Atoi(parts[0])
		entryCommand := strings.Join(parts[1:len(parts)], " ")
		var entry = LogEntry{
			term:    entryTerm,
			command: entryCommand,
		}
		state.log = append(state.log, entry)
	}
}
