package main

import (
	"strings"
)

type LinkedList struct {
	first *LinkedListElem
}

type LinkedListElem struct {
	value string
	next  *LinkedListElem
}

func createLinkedList(firstValue string) *LinkedList {
	return &LinkedList{
		first: &LinkedListElem{value: firstValue, next: nil},
	}
}

func (l *LinkedList) add(newValue string) bool {
	if l.first == nil {
		l.first = &LinkedListElem{value: newValue, next: nil}
	}

	last := l.first
	for ; last.next != nil; last = last.next {
	}
	last.next = &LinkedListElem{value: newValue, next: nil}

	return true
}

func (l *LinkedList) insert(newValue string, index int) bool {
	if index == 1 {
		l.first = &LinkedListElem{value: newValue, next: l.first}
		return true
	}

	current := l.first
	previous := l.first
	i := 1
	for ; i != index && current != nil; current = current.next {
		previous = current
		i++
	}
	if i == index {
		newElement := &LinkedListElem{value: newValue, next: current}
		previous.next = newElement
		return true
	}

	return false
}

func (l *LinkedList) remove(index int) bool {
	current := l.first
	previous := current
	for i := 1; current != nil; current = current.next {
		if i == index {
			if previous == current {
				l.first = current.next
				return true
			} else {
				previous.next = current.next
				return true
			}
		}
		previous = current
		i++
	}

	return false
}

func (l *LinkedList) toString() string {
	if l.first == nil {
		return ""
	}
	elem := l.first
	values := ""

	for ; elem.next != nil; elem = elem.next {
		values = values + " " + elem.value
	}

	values = values + " " + elem.value

	return strings.Trim(values, " ")
}
