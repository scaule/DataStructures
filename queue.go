package main

type Queue struct {
	list *LinkedList
}

func createQueue(firstValue string) *Queue {
	return &Queue{
		list: createLinkedList(firstValue),
	}
}
func (s *Queue) enqueue(value string) {
	s.list.add(value)
}

func (s *Queue) dequeue() string {
	value := s.list.first.value
	s.list.remove(1)

	return value
}
