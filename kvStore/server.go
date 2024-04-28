package kv

import (
	"fmt"
	data "here/dataStruct"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

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
