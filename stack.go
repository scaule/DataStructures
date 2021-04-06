package main

type Stack struct {
	list *LinkedList
}

func createStack(firstValue string) *Stack {
	return &Stack{
		list: createLinkedList(firstValue),
	}
}

func (s *Stack) push(value string) {
	s.list.insert(value, 1)
}

func (s *Stack) pop() string {
	value := s.list.first.value
	s.list.first = s.list.first.next

	return value
}
