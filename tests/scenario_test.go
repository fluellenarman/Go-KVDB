package main_test

import (
	kv "here/kvStore"
	client "here/testClient"
	"os"
	"testing"
)

func TestEmptyInit(t *testing.T) {
	t.Log("testing creating init")
	os.Remove("../kvStore/data/test1.bson")
	test := kv.Init("../kvStore/data/test1.bson")
	t.Log(test)
	kv.Close()
	if test != "Initialized" {
		t.Errorf("failed")
	}
	// should return "Initialized"
}

func TestInit(t *testing.T) {
	t.Log("testing init")
	test := kv.Init("../kvStore/data/test2.bson")
	t.Log(test)
	kv.Close()

	if test != "Loaded" {
		t.Errorf("failed")
	}
	// should return "Loaded"
}

func TestGetSet1(t *testing.T) {
	t.Log("testing get set")
	kv.Init("../kvStore/data/test2.bson")
	client.Set("key1", "Hello World")
	client.Set("key2", 123)
	client.Set("key3", 321)
	client.Set("key4", "Goodbye World")
	val1, exists1 := client.Get("key1")
	val2, exists2 := client.Get("key2")
	val3, exists3 := client.Get("key3")
	val4, exists4 := client.Get("key4")
	kv.Close()

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

func TestDelete1(t *testing.T) {
	t.Log("testing delete")
	kv.Init("../kvStore/data/test2.bson")
	client.Set("key5", "I am to be deleted")
	client.Delete("key5")
	_, exists := client.Get("key5")
	kv.Close()

	if exists != false {
		t.Errorf("failed at deleting key5")
	}
}
