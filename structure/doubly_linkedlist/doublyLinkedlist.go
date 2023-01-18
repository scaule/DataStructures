package main

import (
	"strings"
)

type DoublyLinkedList struct {
	first *DoublyLinkedListElem
}

type DoublyLinkedListElem struct {
	value    string
	previous *DoublyLinkedListElem
	next     *DoublyLinkedListElem
}

func createDoublyLinkedList(firstValue string) *DoublyLinkedList {
	return &DoublyLinkedList{
		first: &DoublyLinkedListElem{value: firstValue, next: nil, previous: nil},
	}
}

func (l *DoublyLinkedList) add(newValue string) bool {
	last := l.first
	var previous *DoublyLinkedListElem
	for ; last.next != nil; last = last.next {
		previous = last
	}
	last.next = &DoublyLinkedListElem{value: newValue, next: nil, previous: previous}

	return true
}

func (l *DoublyLinkedList) remove(index int) bool {
	current := l.first
	previous := current
	for i := 1; current != nil; current = current.next {
		if i == index {
			if previous == current {
				l.first = current.next
				if l.first != nil {
					l.first.previous = nil
				}

				return true
			} else {
				previous.next = current.next
				if previous.next != nil {
					previous.next.previous = previous
				}

				return true
			}
		}
		previous = current
		i++
	}

	return false
}

func (l *DoublyLinkedList) toString() string {
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

func (l *DoublyLinkedList) insert(newValue string, index int) bool {
	if index == 1 {
		l.first = &DoublyLinkedListElem{value: newValue, previous: nil, next: l.first}
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
		newElement := &DoublyLinkedListElem{value: newValue, previous: previous, next: current}
		if newElement.next != nil {
			newElement.next.previous = newElement
		}
		previous.next = newElement

		return true
	}

	return false
}
