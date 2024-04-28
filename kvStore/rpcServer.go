package kv

import (
	"fmt"
	"net"
	"net/rpc"
	"os"
	"strconv"
)

type rpcService struct{}

type RequestVoteArgs struct {
	Term        int
	CandidateId string
}

type RequestVoteReply struct {
	Term        int
	VoteGranted bool
}

func (r *rpcService) RequestVote(args *RequestVoteArgs, reply *RequestVoteReply) error {
	fmt.Println("rpcServer.go::RequestVote()")
	reply.Term = state.term
	if args.Term < state.term {
		fmt.Println("rpcServer.go::RequestVote(), args.Term < state.term")
		reply.VoteGranted = false
		return nil
	}
	if state.votedFor != "" {
		fmt.Println("rpcServer.go::RequestVote(), state.votedFor != \"\" ")
		reply.VoteGranted = false
		return nil
	}

	state.votedFor = args.CandidateId
	reply.VoteGranted = true
	return nil
}

// NOTE, append entries is not implemented, only here now for heartbeats
type AppendEntriesArgs struct {
	Term int
}

type AppendEntriesReply struct {
	Term    int
	Success bool
}

func (r *rpcService) AppendEntries(args *AppendEntriesArgs, reply *AppendEntriesReply) error {
	fmt.Println("rpcServer.go::AppendEntries()")
	if args.Term > state.term {
		state.term = args.Term
		state.status = "follower"
		state.votedFor = ""
	}
	reply.Term = state.term
	reply.Success = true
	ResetElectionTimer()
	// fmt.Println("rpcServer.go::AppendEntries(), ResetElectionTimer(), resetted timer")
	return nil
}

func RPCServer() {
	rpc.RegisterName("rpcService", new(rpcService))
	rpc.HandleHTTP()
	fmt.Println("rpcServer.go::RPCServer()")
	// rpc.RegisterName("rpcService", new(rpcService))
	fmt.Println("rpcServer.go::RPCServer(), state:", state)

	rpcPort := ""
	if os.Getenv("IS_LOCAL") == "false" {
		fmt.Println("rpcServer.go::RPCServer(), IS_LOCAL is false")
		rpcPort = "8081"
	} else {
		fmt.Println("rpcServer.go::RPCServer(), IS_LOCAL is true")
		rpcPortInt, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			// handle the error here
			fmt.Println("Error converting PORT to int:", err)
			return
		}
		rpcPortInt += state.totalSystemNodes
		rpcPort = strconv.Itoa(rpcPortInt)
	}
	// fmt.Println("rpcServer.go::RPCServer(), rpcPort:", rpcPort)

	listener, err := net.Listen("tcp", ":"+(rpcPort))
	if err != nil {
		panic(err)
	}

	// Start serving incoming connections
	fmt.Println("rpcServer.go::RPCServer(), rpcServer listening at ", rpcPort)
	rpc.Accept(listener)

}
