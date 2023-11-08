package kv

import (
	"os"
	"testing"
)

func TestEmptyInit(t *testing.T) {
	t.Log("testing creating init")
	os.Remove("/data/test1.bson")
	test := Init("/data/test1.bson")
	t.Log(test)
	Close()

	if test != "Initialized" {
		t.Errorf("failed")
	}
	// should return "Initialized"
}

func TestInit(t *testing.T) {
	t.Log("testing init")
	test := Init("/data/test2.bson")
	t.Log(test)
	Close()

	if test != "Loaded" {
		t.Errorf("failed")
	}
	// should return "Loaded"
}

func TestGetSet1(t *testing.T) {
	t.Log("testing get set")
	Init("/data/test2.bson")
	Put("key1", "Hello World")
	Put("key2", 123)
	Put("key3", 321)
	Put("key4", "Goodbye World")
	val2, exists2 := Get("key2")
	val1, exists1 := Get("key1")
	val3, exists3 := Get("key3")
	val4, exists4 := Get("key4")
	Close()

	if (val1 != "Hello World") && (exists1 != true) {
		t.Errorf("failed at key1")
		t.Log(exists1, val1)
	}
	if (val2 != 123) && (exists2 != true) {
		t.Errorf("failed at key2")
		t.Log(exists2, val2)
	}
	if (val3 != 321) && (exists3 != true) {
		t.Errorf("failed at key3")
		t.Log(exists3, val3)
	}
	if (val4 != "Goodbye World") && (exists4 != true) {
		t.Errorf("failed at key4")
		t.Log(exists4, val4)
	}
}

func TestDelete(t *testing.T) {
	t.Log("testing delete")
	Init("/data/test2.bson")
	Put("key5", "Delete me")
	DeletePair("key5")
	_, exists := Get("key5")
	Close()

	if exists == true {
		t.Errorf("failed, key5 still exists")
	}

}
