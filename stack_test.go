package main

import (
	"testing"
)

func TestStackPushAndPop(t *testing.T) {
	list := createStack("simon")
	list.push("a")
	list.push("b")
	list.push("c")
	value := list.pop()
	if value != "c" {
		t.Errorf("error on stack")
	}
	value = list.pop()
	if value != "b" {
		t.Errorf("error on stack")
	}
	value = list.pop()
	if value != "a" {
		t.Errorf("error on stack")
	}
	list.push("a")
	if value != "a" {
		t.Errorf("error on stack")
	}
}
