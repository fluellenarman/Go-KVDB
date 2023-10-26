package data

type Person struct {
	Name string
	Age  int
	Sex  string
}

type Data struct {
	Key   string
	Value []byte
}

type ByteArrayMap struct {
	data map[string][]byte
}
