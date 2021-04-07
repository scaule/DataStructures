package main

import "testing"

func TestBinaryInsert(t *testing.T) {
	binary := createBinaryTree(1)
	binary.insert(2, nil)
	binary.insert(3, nil)
	binary.insert(19, nil)
	binary.insert(12, nil)
	binary.printBinary(nil)
}


func TestBinarySearch(t *testing.T) {
	binary := createBinaryTree(1)
	binary.insert(2, nil)
	binary.insert(3, nil)
	binary.insert(19, nil)
	binary.insert(12, nil)

	if binary.search(2, nil) == nil {
		t.Errorf("error on search")
	}
	if binary.search(1, nil) == nil {
		t.Errorf("error on search")
	}
	if binary.search(3, nil) == nil {
		t.Errorf("error on search")
	}
	if binary.search(19, nil) == nil {
		t.Errorf("error on search")
	}
	if binary.search(12, nil) == nil {
		t.Errorf("error on search")
	}
	if binary.search(40, nil) != nil {
		t.Errorf("error on search")
	}
}