package stack

import "fmt"

func (stack *Stack) Empty() bool {
	return stack.length == 0
}

func (stack *Stack) Size() int {
	return stack.length
}

func (stack *Stack) Add(value []ElementValue) {
	element := newElement(value)

	if stack.Empty() {
		stack.topElement = element
	} else {
		currentElement := stack.topElement
		element.next = &currentElement
		stack.topElement = element
	}

	stack.length++
}

func (stack *Stack) Top() []ElementValue {
	value := []ElementValue{}
	if !stack.Empty() {
		return stack.topElement.Value
	}
	return value
}

func (stack *Stack) Pop() []ElementValue {
	if !stack.Empty() {
		stack.length--
		currentElement := stack.topElement
		if stack.topElement.next != nil {
			stack.topElement = *stack.topElement.next
		}
		if stack.length == 0 {
			stack.topElement = Element{}
		}
		return currentElement.Value
	}
	fmt.Println("lista vazia")
	return []ElementValue{}
}

func (stack *Stack) Run() []map[string][]interface{} {
	var finalResult []map[string][]interface{}

	for !stack.Empty() {
		combinedElements := make(map[string][]interface{})
		elementValues := stack.Pop()

		for _, element := range elementValues {
			combinedElements[element.Key] = append(combinedElements[element.Key], element.Value)
		}

		finalResult = append(finalResult, combinedElements)
	}

	return finalResult
}
