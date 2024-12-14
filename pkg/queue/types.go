package queue

type IQueue interface {
	Enqueue() error
	Dequeue() error
	NewElement() error
	Head() error
	Tail() error
	Empty() error
	Size() error
	Full() error
}
type Queue struct {
	Type   string
	Head   Element
	Tail   Element
	Length int
	Full   int
}

type QueueStruct struct {
	Type   string
	Head   Element
	Tail   Element
	Length int
	Full   int
	IQueue
}

type ClassicQueue struct {
	Queue
}

type CircleQueue struct {
	Queue
}

type Element struct {
	Value interface{}
	next  *Element
}

type EContent struct {
	Key   string
	Value interface{}
}
