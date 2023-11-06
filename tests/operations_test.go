package main_test

import (
	kv "here/kvStore"
	"testing"
)

func TestInit(t *testing.T) {
	t.Log("testing init")
	test := kv.Init("../kvStore/data/test1.bson")
	t.Log(test)
	if test != "Initialized" {
		t.Errorf("failed")
	} else {
		t.Log("passed")
	}
}

// func TestGetPut1(t *testing.T) {
// 	t.Log("testing put")
// 	t.Errorf("failed")
// }
