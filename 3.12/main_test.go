package main

import "testing"

func TestAnagram(t *testing.T) {
	if !anagram("heart", "earth") {
		t.Error("Expected heart anagram earth, got false")
	}
	if !anagram("python", "typhon") {
		t.Error("Expected python anagram typhon, got false")
	}
	if anagram("abcde", "qwerty") {
		t.Error("Expected abcde not anagram qwerty, got true")
	}
}
