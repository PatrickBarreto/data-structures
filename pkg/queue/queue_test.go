package queue

import "testing"

func TestNewQueueMethod(t *testing.T) {
	newQueue := NewClassicQueue(4)
	t.Log(newQueue)
}
