package main

import "testing"

func TestContains(t *testing.T) {
	rs := NewRandomizedSet()
	rs.Insert(1)
	if rs.Contains(1) != true {
		t.Error("Expected true, got false")
	}
	if rs.Contains(2) == true {
		t.Error("Expected false, got true")
	}
}

func TestRemove(t *testing.T) {
	rs := NewRandomizedSet()
	rs.Insert(1)
	if rs.Remove(1) != true {
		t.Error("Expected true, got false")
	}
	if rs.Remove(2) == true {
		t.Error("Expected false, got true")
	}
}

// func TestInsert(t *testing.T) {
// 	rs := RandomizedSet{}
// 	rs.Insert(1)
// 	if 1 != len(RandomizedSet.hashTable) {
// 		t.Errorf("Expected len of 1, got %d", len(RandomizedSet.hashTable))
// 	}
// }
