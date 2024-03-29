package kv

import (
	"fmt"
	data "here/dataStruct"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"encoding/json"

	"gopkg.in/mgo.v2/bson"
)

// Main is the main function for the kvStore
func Main() {
	fmt.Println("kvStore server runnning...")
	// go sendHeartbeat()
	port := os.Getenv("PORT")
	fmt.Println("port:", port)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading request body:", err)
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var data data.KVpair
	bson.Unmarshal(body, &data)

	fmt.Println("server.go: r->", r.Method)
	switch r.Method {
	case http.MethodGet:
		fmt.Println("GET")
		val, exists := Get(data.Key)
		if !exists {
			fmt.Fprintf(w, "Key does not exist")
		} else {
			fmt.Println("server.go: key exists")
			strVal := convertValToString(val)
			fmt.Fprint(w, strVal)
		}
		// fmt.Fprintf(w, "GET")
	case http.MethodPost:
		Put(data.Key, data.Value)
		dataVal := convertValToString(data.Value)
		appendLogEntry("SET " + data.Key + " " + dataVal)
		writeLogToStorage()
		fmt.Fprintf(w, "POST Sucessful")
	case http.MethodDelete:
		dataVal := convertValToString(data.Value)
		appendLogEntry("DELETE " + data.Key + " " + dataVal)
		DeletePair(data.Key)
		writeLogToStorage()

		fmt.Fprintf(w, "DELETE")
	case "CLOSE_DB":
		fmt.Println("closing server..")
		Close()
		fmt.Fprintf(w, "server's database is closed")
		fmt.Println(state)
	case "HEART_BEAT":
		fmt.Println("Heartbeat received from", r.RemoteAddr)
	case "VOTE_REQUEST":
		// fmt.Println("Request vote received from", string(body))
		ResetElectionTimer()
		var voteRequest RequestVote
		err := json.Unmarshal(body, &voteRequest)
		if err != nil {
			fmt.Println("Error unmarshalling request vote")
		}
		fmt.Println("Request vote received", voteRequest)

		SendVoteResponse(voteRequest)
	case "VOTE_RESPONSE":
		var voteResponse ResponseVote
		err := json.Unmarshal(body, &voteResponse)
		if err != nil {
			fmt.Println("Error unmarshalling vote response")
		}
		fmt.Println("Vote response received from ", voteResponse.VoterId)
	default:
		fmt.Println("default")
		fmt.Fprintf(w, "default")
	}
}

func convertValToString(val interface{}) string {
	switch v := val.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	default:
		return "Error converting value to string"
	}
}

func sendHeartbeat() {
	for {
		time.Sleep(5000 * time.Millisecond)
		fmt.Println("Sending heartbeat")
		sendHeartBeatMessage()
	}
}

func sendHeartBeatMessage() {
	randomNum := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(5)
	appList := []string{"app1", "app2", "app3", "app4", "app5"}
	url := "http://" + appList[randomNum] + ":8080/"

	fmt.Println("Sending heartbeat")

	req, err := http.NewRequest("HEART_BEAT", url, nil)
	if err != nil {
		fmt.Println("Error creating request")
	} else {
		fmt.Println("Heartbeat request created")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending heartbeat", err)
	} else {
		defer resp.Body.Close()
		fmt.Println("Heartbeat sent to " + appList[randomNum])
	}
}
