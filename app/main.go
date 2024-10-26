package main

import (
	"fmt"

	stack "github.com/PatrickBarreto/data-structures/app/stack"
)

func simpleUseStack() {
	myStack := stack.NewStack()
	myStack.Add([]stack.ElementValue{
		{Key: "name", Value: "João"},
		{Key: "age", Value: 32},
		{Key: "profile", Value: "admin"},
	})
	myStack.Add([]stack.ElementValue{
		{Key: "name", Value: "Maria"},
		{Key: "age", Value: 24},
		{Key: "profile", Value: "admin"},
	})

	fmt.Println(myStack.Size())

	results := myStack.Run()

	for _, result := range results {
		fmt.Println(result)
	}
}

func main() {
	simpleUseStack()
}
