package main

import (
	"fmt"
	"testing"
)

func TestAddUpper(t *testing.T) {
	n := 10
	r := addUpper(n)
	if r != 55 {
		t.Fatalf("addUpper(%v) error; return: %v; expect: %v.\n", n, r, 55)
	} else {
		t.Logf("addUpper(%v) right; return: %v; expect: %v.\n", n, r, 55)
	}
}

func TestHello(t *testing.T) {
	fmt.Println("The function is called.")
}
