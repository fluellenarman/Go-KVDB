package client

import (
	"fmt"
	kv "here/kvStore"
)

func Set(key string, value interface{}) {
	fmt.Println("client.go : Set()", key, value)
	kv.Put(key, value)
}

func Get(key string) (interface{}, bool) {
	fmt.Println("client.go : Get()", key)
	value, exists := kv.Get(key)
	// fmt.Println("client.go : Get() received -> ", value)
	return value, exists
}

func Delete(key string) {
	fmt.Println("client.go : Delete()")
	fmt.Println(key)
	kv.DeletePair(key)
	fmt.Println("client.go : Delete() completed")
}

func Init() {
	fmt.Println("client.go : Init()")
}
