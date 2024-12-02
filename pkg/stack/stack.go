package stack

type ElementValue struct {
	Key   string
	Value interface{}
}

type Element struct {
	Value []ElementValue
	next  *Element
}

type Stack struct {
	topElement Element
	length     int
}

func NewStack() Stack {
	return Stack{
		topElement: Element{},
		length:     0,
	}
}

func newElement(value []ElementValue) Element {
	return Element{
		Value: value,
		next:  nil,
	}
}
