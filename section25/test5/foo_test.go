package main

import "testing"

func TestFoo(t *testing.T) {

	v := Foo(10, 15)
	if v != 25 {
		t.Error("Expected 25, got ", v)
	}
}
