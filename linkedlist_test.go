package main

import (
	"testing"
)

func TestLinkedListAdd(t *testing.T) {
	list := createLinkedList("simon")
	list.add("a")
	if list.toString() != "simon a" {
		t.Errorf("error on add")
	}
	list.add("b")
	if list.toString() != "simon a b" {
		t.Errorf("error on add")
	}
}

func TestLinkedListInsert(t *testing.T) {
	list := createLinkedList("a")
	if list.toString() != "a" {
		t.Errorf("error on insert")
	}
	list.insert("d", 2)
	if list.toString() != "a d" {
		t.Errorf("error on insert")
	}
	list.insert("c", 2)
	if list.toString() != "a c d" {
		t.Errorf("error on insert")
	}
	list.insert("b", 2)
	if list.toString() != "a b c d" {
		t.Errorf("error on insert")
	}
	result := list.insert("e", 6)
	if list.toString() != "a b c d" && result != false {
		t.Errorf("error on insert")
	}
	list.insert("e", 5)
	if list.toString() != "a b c d e" {
		t.Errorf("error on insert")
	}
}

func TestLinkedListDelete(t *testing.T) {
	list := createLinkedList("simon")
	list.add("a")
	list.add("b")
	success := list.remove(3)
	if !success || list.toString() != "simon a" {
		t.Errorf("error on remove")
	}
	success = list.remove(2)
	if !success || list.toString() != "simon" {
		t.Errorf("error on remove")
	}
	success = list.remove(1)
	if !success || list.toString() != "" {
		t.Errorf("error on remove")
	}

	success = list.remove(1)
	if success || list.toString() != "" {
		t.Errorf("error on remove")
	}

	success = list.remove(3)
	if success || list.toString() != "" {
		t.Errorf("error on remove")
	}
}
