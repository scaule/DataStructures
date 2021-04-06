package main

import (
	"testing"
)

func TestQueue(t *testing.T) {
	queue := createQueue("a")
	queue.enqueue("b")
	queue.enqueue("c")
	value := queue.dequeue()
	if value != "a" {
		t.Errorf("error on queue")
	}
	value = queue.dequeue()
	if value != "b" {
		t.Errorf("error on queue")
	}
	value = queue.dequeue()
	if value != "c" {
		t.Errorf("error on queue")
	}
	queue.enqueue("end")
	value = queue.dequeue()
	if value != "end" {
		t.Errorf("error on queue")
	}
}
