package monster

import (
	"testing"
)

func TestMonster_Store(t *testing.T) {
	m := &monster{
		Name:  "Red babe",
		Age:   10,
		Skill: "Spit fire",
	}
	err := m.Store()
	if err != nil {
		t.Fatalf("Store error:\n%v\n", err)
	} else {
		t.Logf("Store successfully, return %v.\n", err)
	}
}

func TestMonster_ReStore(t *testing.T) {
	m := monster{}
	t.Logf("Before restore, m is %v.\n", m)
	err := m.ReStore()
	t.Logf("After restore, m is %v.\n", m)
	if err != nil {
		t.Fatalf("ReStore error:\n%v\n", err)
	} else {
		t.Logf("ReStore successfully, return %v.\n", err)
	}
}
