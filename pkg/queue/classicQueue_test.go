package queue

import (
	"testing"
)

func TestEnqueueMethod(t *testing.T) {
	myQueue := NewClassicQueue(4)
	myQueue.Enqueue([]EContent{
		{Key: "1", Value: "primeiro elemento inserido"},
	})
	myQueue.Enqueue([]EContent{
		{Key: "2", Value: "segundo elemento inserido"},
	})
}

func TestShoulEnqueueElementOnBottomWithoutNextNil(t *testing.T) {
	testQueue := NewClassicQueue(4)

	testQueue.Enqueue([]EContent{
		{Key: "teste", Value: "teste"},
	})
	testQueue.Enqueue([]EContent{
		{Key: "testeKey", Value: "testeValue"},
	})

	testQueue.Enqueue([]EContent{
		{Key: "testeKey2", Value: "testeValue2"},
	})

	testQueue.Enqueue([]EContent{
		{Key: "testeKey3", Value: "testeValue3"},
	})

	head := testQueue.getHead()
	tail := testQueue.getTail()

	t.Log("HEAD")
	t.Log(head)

	t.Log("TAIL")
	t.Log(tail)
}

func TestShoulDequeueElementOnTop(t *testing.T) {
	testQueue := NewClassicQueue(4)

	testQueue.Enqueue([]EContent{
		{Key: "teste", Value: "teste"},
	})
	testQueue.Enqueue([]EContent{
		{Key: "testeKey", Value: "testeValue"},
	})

	element := testQueue.Dequeue()
	t.Log("Element")
	t.Log(element)

	element = testQueue.Dequeue()

	t.Log("Element 1")
	t.Log(element)

	element = testQueue.Dequeue()

	t.Log("Element 2")
	t.Log(element)

}
