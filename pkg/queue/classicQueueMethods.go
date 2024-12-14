package queue

import (
	"fmt"
)

func (queue *Queue) persistQueue(currentElement *Element, newElement Element) bool {
	currentElement.next = &newElement
	queue.Tail = newElement
	queue.Length++
	return true
}

func (queue *Queue) Enqueue(value []EContent) bool {
	newElement := NewElement(value)

	if queue.isFull() {
		fmt.Println("pilha Cheia")
		return false
	}

	if queue.Empty() {
		queue.Head = newElement
		queue.Tail = Element{}
		queue.Length++
		return true
	}

	if !queue.Empty() {
		currentElement := queue.getHead()
		for currentElement.next != nil {
			currentElement = currentElement.next
		}
		return queue.persistQueue(currentElement, newElement)
	}

	return true
}

func (queue *Queue) Dequeue() Element {
	if queue.Empty() {
		return Element{}
	}

	toReturn := queue.Head

	if queue.Head.next != nil {
		queue.Head = *queue.Head.next
	}

	queue.Length--

	return toReturn
}
