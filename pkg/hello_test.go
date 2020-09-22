package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	if "Hello World" != Hello() {
		t.Fail()
	}
}
