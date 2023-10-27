package data

type Person struct {
	Name string
	Age  int
	Sex  string
}

type ValueTuple struct {
	DataType string // represents datatype of value
	Value    []byte
}

// string key, tuple(string, []byte)
type KVmap struct {
	MemoMap map[string]ValueTuple
}
