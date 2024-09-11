package main

import (
	"bytes"
	"testing"
)

func TestLogger(t *testing.T) {
	var buf bytes.Buffer
	logger(&buf, "Test item")

	expected := "Test item\n"

	if buf.String() != expected {
		t.Errorf("got %v want %v", buf.String(), expected)
	}
}
