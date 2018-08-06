package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	actual := getHello()
	expected := "Hello World"
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}
