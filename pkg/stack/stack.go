package stack

type EContent struct {
	Key   string
	Value interface{}
}

type Element struct {
	Value []EContent
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

func newElement(value []EContent) Element {
	return Element{
		Value: value,
		next:  nil,
	}
}
