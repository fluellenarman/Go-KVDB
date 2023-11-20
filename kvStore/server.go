package kv

import (
	"fmt"
	data "here/dataStruct"
	"io/ioutil"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

// Main is the main function for the kvStore
func Main() {
	fmt.Println("kvStore server runnning...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Server->", r)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading request body:", err)
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var data data.KVpair
	bson.Unmarshal(body, &data)

	switch r.Method {
	case http.MethodGet:
		fmt.Println("GET")
		val, exists := Get(data.Key)
		if exists != true {
			fmt.Fprintf(w, "Key does not exist")
		} else {
			fmt.Fprintf(w, val.(string))
		}
		// fmt.Fprintf(w, "GET")
	case http.MethodPost:
		fmt.Fprintf(w, "POST")
		Put(data.Key, data.Value)
	case http.MethodDelete:
		fmt.Fprintf(w, "DELETE")
		DeletePair(data.Key)
	default:
		fmt.Fprintf(w, "default")
	}
}
