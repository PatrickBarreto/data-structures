package queue

import "fmt"

// Cada tipo de fila tem um tipo de enqueue
func (queue *Queue) ClassicEnqueue(value []ElementValue) bool {
	newElement := newElement(value)

	if queue.Full() {
		fmt.Println("pilha Cheia")
		return false
	}

	if queue.Empty() {
		queue.head = newElement
	} else {
		currentElement := queue.head

		for currentElement.next != nil {
			currentElement = *currentElement.next
		}
		currentElement.next = &newElement
		queue.head = currentElement
	}

	queue.length++

	return true
}

// Cada tipo de fila tem um tipo de enqueue
func (queue *Queue) ClassicDequeue() Element {
	var toReturn Element
	if !queue.Empty() {
		toReturn = queue.head

		if queue.head.next != nil {
			queue.head = *queue.head.next
		} else {
			queue.head = Element{}
		}

		queue.length--
		return toReturn
	}
	return toReturn
}

/*






 */
// Cada tipo de fila tem um indicador de cabeça
func (queue *Queue) Head() Element {
	return queue.head
}

// Cada tipo de fila tem um indicador de última posição?
func (queue *Queue) Tail() Element {
	return queue.tail
}

// Cada tipo lista pode saber se é vazia
func (queue *Queue) Empty() bool {
	return queue.length == 0
}

// Cada tipo lista pode saber seu tamanho
func (queue *Queue) Size() int {
	return queue.length
}

// Cada tipo lista pode saber se está cheia
func (queue *Queue) Full() bool {
	return queue.length >= queue.full
}
