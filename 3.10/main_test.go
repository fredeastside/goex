package main

import (
	"testing"
)

func TestComma(t *testing.T) {
	s := comma("123456789")
	if s != "123,456,789" {
		t.Errorf("Expected 123,456,789 got %s", s)
	}
	s = comma("22")
	if s != "22" {
		t.Errorf("Expected 22 got %s", s)
	}
	s = comma("1234567890")
	if s != "1,234,567,890" {
		t.Errorf("Expected 1,234,567,890 got %s", s)
	}
}

func TestCommaBuff(t *testing.T) {
	s := commaBuff("22")
	if s != "22" {
		t.Errorf("Expected 22 got %s", s)
	}
	s = commaBuff("123456789")
	if s != "123,456,789" {
		t.Errorf("Expected 123,456,789 got %s", s)
	}
	s = commaBuff("123456789.50")
	if s != "123,456,789.50" {
		t.Errorf("Expected 123,456,789.50 got %s", s)
	}
}
