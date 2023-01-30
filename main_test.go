package main

import "testing"

func Test_greet(t *testing.T) {
	if got := greet(); got != "hello world" {
		t.Errorf("greet() = %v, want %v", got, "hello world")
	}
}
