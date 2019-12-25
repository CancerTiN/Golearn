package main

import "testing"

func TestGetSub(t *testing.T) {
	n1, n2 := 10, 3
	r := getSub(n1, n2)
	if r != 7 {
		t.Fatalf("getSub(%v, %v) error; return: %v; expect: %v.\n", n1, n2, r, 7)
	} else {
		t.Logf("getSub(%v, %v) right; return: %v; expect: %v.\n", n1, n2, r, 7)
	}
}
