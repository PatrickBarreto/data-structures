package queue

import "testing"

func TestNewQueueMethod(t *testing.T) {
	newQueue := NewClassicQueue(4)
	t.Log(newQueue)
}

func TestNewElementQueueMethod(t *testing.T) {
	if 4/2 == 2 {
		t.Error()
	}
}
