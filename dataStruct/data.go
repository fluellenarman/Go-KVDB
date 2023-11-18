package data

type Person struct {
	Name string
	Age  int
	Sex  string
}

// string key, tuple(string, []byte)
type KVmap struct {
	MemoMap map[string]interface{}
}

type KVpair struct {
	Key   string
	Value interface{}
}
