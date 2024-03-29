package kv

import (
	// "bufio"
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	// "strings"
	"encoding/json"
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
	votedForId       string
	log              []LogEntry
	totalSystemNodes int
}

type RequestVote struct {
	Term         int
	CandidateId  string
	CandidateUrl string
}

type ResponseVote struct {
	VoterId     string
	VoterTerm   int
	CandidateId string
}

type LogEntry struct {
	term    int
	command string
}

var state RaftState
var currentElectionTimer = (time.Duration(state.electionTimeout) * time.Millisecond)

func InitRaft() {
	state = RaftState{
		appName:          os.Getenv("NAME"),
		status:           "follower",
		term:             0,
		electionTimeout:  getRandomNumber(), // 150-300ms
		heartbeatTimeout: 100,
		votedFor:         "",
		log:              []LogEntry{},
		totalSystemNodes: 5,
	}
	fmt.Println("initRaft() state", state)
	// fmt.Println("initRaft() appName", state.appName)
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
	state.votedFor = state.appName
	state.status = "candidate"
	sendRequestVote()
	ResetElectionTimer()
}

func sendRequestVote() {
	// send request to other nodes
	// if received majority votes, become leader
	// else, transition back into follower
	// fmt.Println("sendRequestVote()")

	appList := []string{"app1", "app2", "app3", "app4", "app5"}
	for _, app := range appList {
		if app == state.appName {
			continue
		}
		fmt.Println("creating VOTE_REQUEST for ", app)
		url := "http://" + app + ":8080/"

		curRequestVote := RequestVote{
			Term:         state.term,
			CandidateId:  state.appName,
			CandidateUrl: "http://" + state.appName + ":8080/",
		}

		requestBody, err := json.Marshal(curRequestVote)
		if err != nil {
			fmt.Println("Error marshalling request vote")
		}

		req, err := http.NewRequest("VOTE_REQUEST", url, bytes.NewBuffer(requestBody))
		if err != nil {
			fmt.Println("Error creating request")
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("Error sending request vote", err)
		} else {
			defer resp.Body.Close()
			fmt.Println("VOTE_REQUEST created and sent to " + app)
		}
	}

	ResetElectionTimer()
}

func SendVoteResponse(voteRequest RequestVote) {
	var voteData ResponseVote
	// if node hasn't voted yet and term received is greater than current term
	if voteRequest.Term > state.term && state.votedFor == "" {
		ResetElectionTimer()

		state.votedFor = voteRequest.CandidateId
		state.term = voteRequest.Term
		state.status = "follower"

		// fmt.Println("creating VOTE_RESPONSE from - ", state.appName, " - to - ", voteRequest.CandidateId)
		voteData = ResponseVote{
			VoterId:     state.appName,
			VoterTerm:   state.term,
			CandidateId: voteRequest.CandidateId,
		}
	} else if voteRequest.Term >= state.term && state.votedFor != "" { // if node has already voted
		ResetElectionTimer()
		voteData = ResponseVote{
			VoterId:     state.appName,
			VoterTerm:   state.term,
			CandidateId: voteRequest.CandidateId,
		}
	}

	// SENDING VOTES
	responseBody, err := json.Marshal(voteData)
	if err != nil {
		fmt.Println("Error marshalling vote response")
	}

	// send message to node that requested the vote
	req, err := http.NewRequest("VOTE_RESPONSE", voteRequest.CandidateUrl, bytes.NewBuffer(responseBody))
	if err != nil {
		fmt.Println("Error creating vote response")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending vote response", err)
	} else {
		defer resp.Body.Close()
		fmt.Println("VOTE_RESPONSE sent to " + voteRequest.CandidateId)
	}

}

func countVotes() {
	// count votes from other nodes
	// if majority votes, become leader

	// else, transition back into follower
	state.status = "follower"
	// fmt.Println("lost election")
	ResetElectionTimer()
}

// func sendMessage()

func ResetElectionTimer() {
	currentElectionTimer = (time.Duration(state.electionTimeout) * time.Millisecond)
	// fmt.Println("ResetElectionTimer() currentElectionTimer", currentElectionTimer)
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

// func loadLogToMemory() {
// 	file, err := os.OpenFile("kvStore/data/logfile.txt", os.O_RDONLY, 0666)
// 	if err != nil {
// 		fmt.Println("Error opening file")
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)

// 	for scanner.Scan() {
// 		line := scanner.Text()

// 		parts := strings.Fields(line)
// 		entryTerm, _ := strconv.Atoi(parts[0])
// 		entryCommand := strings.Join(parts[1:len(parts)], " ")
// 		var entry = LogEntry{
// 			term:    entryTerm,
// 			command: entryCommand,
// 		}
// 		state.log = append(state.log, entry)
// 	}
// }

// helper util functions
