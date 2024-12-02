package stack

import (
	"testing"
)

func TestEmptyMethod(t *testing.T) {
	myStack := NewStack()

	// Testa o caso em que a pilha está vazia
	if !myStack.Empty() {
		t.Error()
	}

	t.Log("Test passed successfully!")
}

func TestStack(t *testing.T) {

	myStack := NewStack()
	myStack.Add([]ElementValue{
		{Key: "name", Value: "João"},
		{Key: "age", Value: 32},
		{Key: "profile", Value: "admin"},
	})
	myStack.Add([]ElementValue{
		{Key: "name", Value: "Maria"},
		{Key: "age", Value: 24},
		{Key: "profile", Value: "admin"},
	})

	results := myStack.Run()

	if len(results) == 0 {
		t.Fatal("Expected results, but got none")
	}

	t.Log("Test passed successfully!")
}
