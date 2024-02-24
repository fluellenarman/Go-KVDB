package kv

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type RaftState struct {
	nodeName         string
	status           string
	term             int
	electionTimeout  int
	heartbeatTimeout int
	votedFor         string
	log              []LogEntry
}

type LogEntry struct {
	term    int
	command string
}

var state RaftState
var currentElectionTimer = (time.Duration(state.electionTimeout) * time.Millisecond)

func InitRaft() {
	state = RaftState{
		nodeName:         "node1",
		status:           "follower",
		term:             0,
		electionTimeout:  getRandomNumber(), // 150-300ms
		heartbeatTimeout: 100,
		votedFor:         "",
		log:              []LogEntry{},
	}
	fmt.Println("initRaft() state", state)
	// loadLogToMemory()
	// fmt.Println("loaded log to memory\n", state.log)
	// heartBeatTimer()
	fmt.Println("electionTimeout: " + strconv.Itoa(state.electionTimeout))
	go electionTimer()
}

func electionTimer() {
	for {
		select {
		case <-time.After(currentElectionTimer):
			fmt.Println("Currently in state:", state.status)
			if state.status == "follower" {
				startElection()
			} else if state.status == "candidate" {
				countVotes()
			} else if state.status == "leader" {
				// sendHeartbeat()
			}
		}
	}
}

func startElection() {
	state.term++
	state.votedFor = state.nodeName
	state.status = "candidate"
	// sendRequestVote()
	resetElectionTimer()
}

func countVotes() {
	// count votes from other nodes
	// if majority votes, become leader

	// else, transition back into follower
	state.status = "follower"
	// fmt.Println("lost election")
	resetElectionTimer()
}

func resetElectionTimer() {
	currentElectionTimer = (time.Duration(state.electionTimeout) * time.Millisecond)
	// fmt.Println("resetElectionTimer() currentElectionTimer", currentElectionTimer)
}

func getRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(151) + 150
}

// func heartBeatTimer() {
// 	for {
// 		select {
// 		case <-time.After(time.Duration(state.heartbeatTimeout) * time.Millisecond):
// 			fmt.Println("Sending heartbeat")
// 			// sendHeartbeat()
// 		}
// 	}
// }

// Ignore Below for now //

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
