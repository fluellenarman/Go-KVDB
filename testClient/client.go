package client

import (
	"fmt"
	kv "here/kvStore"
)

func Set(key string, value interface{}) {
	fmt.Println("client.go : Set()")
	fmt.Println(key, value)
	kv.SerializeValue(key, value)
}

func Get(key string) {
	fmt.Println("client.go : Get()")
	fmt.Println(key)
	value := kv.DeserializeValue(key)
	fmt.Println("client.go : Get() received -> ", value)
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
