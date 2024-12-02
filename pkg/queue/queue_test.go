package queue

import (
	"testing"
)

func TestClassicQueue(t *testing.T) {
	myQueue := NewSimpleQueue(4)
	myQueue.ClassicEnqueue([]ElementValue{
		{Key: "teste", Value: "teste"},
	})
	myQueue.ClassicEnqueue([]ElementValue{
		{Key: "teste2", Value: "teste2"},
	})

	myQueue.ClassicDequeue()
	myQueue.ClassicDequeue()
	myQueue.ClassicDequeue()

	myQueue.Head()

	myQueue.Size()
	myQueue.Full()
}
