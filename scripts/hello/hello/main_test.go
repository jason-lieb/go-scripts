package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
	result := HelloWorld()

	if result != expected {
		t.Errorf("MyFunction() = %v; want %v", result, expected)
	}
}
