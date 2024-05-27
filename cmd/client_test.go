package main

import "testing"

func InputFromUser_test(t *testing.T) {
	got := InputFromUser()
	if got != nil {
		t.Error("test failed")
	}
}
