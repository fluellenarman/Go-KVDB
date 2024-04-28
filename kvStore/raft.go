package kv

import (
	// "bufio"

	"fmt"
	"math/rand"
	"net/rpc"
	"os"
	"strconv"

	// "strings"

	"time"
)

// Eventually, RaftState will have to track the status of all nodes in the cluster.
// Meaning that it will have to keep track of the term, votedFor, log, etc. of all nodes.
type RaftState struct {
	appName          string
	status           string
	term             int
	electionTimeout  int
	heartbeatTimeout int
	votedFor         string
	log              []LogEntry
	totalSystemNodes int
	port             int
	totalVotes       int
}

type LogEntry struct {
	term    int
	command string
}

var state RaftState
var currentElectionTimer = (time.Duration(state.electionTimeout) * time.Millisecond)
var ResetTimer = make(chan bool)

func InitRaft() {
	rpcPortInt, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		// handle the error here
		fmt.Println("Error converting PORT to int:", err)
		return
	}

	state = RaftState{
		appName:          os.Getenv("NAME"),
		status:           "follower",
		term:             0,
		electionTimeout:  getRandomNumber(), // 150-300ms
		heartbeatTimeout: 100,
		votedFor:         "",
		log:              []LogEntry{},
		totalSystemNodes: 5,
		port:             rpcPortInt,
		totalVotes:       0,
	}
	fmt.Println("raft.go::initRaft(), state", state)

	fmt.Println("raft.go::initRaft(), electionTimeout: " + strconv.Itoa(state.electionTimeout))
	go RPCServer()
	// sleep to allow all nodes to start
	time.Sleep(5 * time.Second)
	go electionTimer()

	// sendRequestVote()
	// go electionTimer()
}

func electionTimer() {
	duration := (time.Duration(state.electionTimeout) * time.Millisecond)
	timer := time.NewTimer(duration)
	for {
		select {
		case <-timer.C:
			fmt.Println("raft.go::electionTimer(), timed out")
			timer.Reset(duration)
			if state.status == "follower" {
				if os.Getenv("IS_LOCAL") == "true" {
					return
				}
				go startElection()
			} else if state.status == "candidate" {
				// demote self to follower
				state.status = "follower"
			}
		case <-ResetTimer:
			fmt.Println("raft.go::electionTimer(), reset timer")
			timer.Reset(duration)
		}
	}
}

func startElection() {
	state.term++
	state.votedFor = state.appName
	state.status = "candidate"
	sendRequestVote()
}

func sendRequestVote() {
	// send request vote to all other nodes
	portList := []int{8080, 8081, 8082, 8083, 8084}
	counter := 0
	for _, port := range portList {
		counter += 1
		if port == state.port {
			continue
		}
		prefix := "app" + strconv.Itoa(counter)
		dest := prefix + ":8081"
		client, err := rpc.Dial("tcp", dest)
		if err != nil {
			fmt.Println("raft.go::sendRequestVote(), err:", err)
			panic(err)
		}
		defer client.Close()

		var reply RequestVoteReply
		args := RequestVoteArgs{Term: state.term, CandidateId: state.appName}

		err = client.Call("rpcService.RequestVote", args, &reply)
		if err != nil {
			panic(err)
		}
		countVotes(reply)
		if state.status == "leader" {
			fmt.Println("raft.go::sendRequestVote(), elected leader")
			go heartbeatTimer()
			break
		} else if state.status == "follower" {
			state.totalVotes = 0
			return
		}
	}
}

func heartbeatTimer() {
	duration := (time.Duration(state.heartbeatTimeout) * time.Millisecond)
	timer := time.NewTimer(duration)
	for {
		select {
		case <-timer.C:
			fmt.Println("raft.go::heartbeatTimer(), timed out")
			timer.Reset(duration)
			if state.status == "leader" {
				fmt.Println("raft.go::heartbeatTimer(), current state:", state)
				ResetElectionTimer()
				sendHeartbeat()
			} else {
				return
			}
		}
	}
}

func sendHeartbeat() {
	// send heartbeat to all other nodes
	portList := []int{8080, 8081, 8082, 8083, 8084}
	counter := 0
	for _, port := range portList {
		counter += 1
		if port == state.port {
			continue
		}
		prefix := "app" + strconv.Itoa(counter)
		dest := prefix + ":8081"
		client, err := rpc.Dial("tcp", dest)
		if err != nil {
			fmt.Println("raft.go::sendHeartbeat(), err:", err)
			panic(err)
		}
		defer client.Close()

		var reply AppendEntriesReply
		args := AppendEntriesArgs{Term: state.term}

		err = client.Call("rpcService.AppendEntries", args, &reply)
		if err != nil {
			panic(err)
		}
	}

}

func countVotes(reply RequestVoteReply) {
	// count votes, if greater than half, become leader.
	// if status is follower, cancel timer and reset
	if reply.VoteGranted {
		state.totalVotes++
	}
	if state.totalSystemNodes%2 == 0 {
		if state.totalVotes >= state.totalSystemNodes/2 {
			state.status = "leader"
		}
	} else {
		if state.totalVotes > state.totalSystemNodes/2 {
			state.status = "leader"
		}
	}
}

func ResetElectionTimer() {
	ResetTimer <- true
}

func getRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(151) + 150
}

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
