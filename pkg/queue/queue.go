package queue

type ElementValue struct {
	Key   string
	Value interface{}
}

type Element struct {
	Value []ElementValue
	next  *Element
}

type Queue struct {
	head   Element
	tail   Element
	length int
	full   int
}

func newElement(value []ElementValue) Element {
	return Element{
		Value: value,
		next:  nil,
	}
}

/*
Default maxSize = 10
*/
func NewSimpleQueue(maxSize ...int) Queue {
	size := 10

	if len(maxSize) > 0 {
		size = maxSize[0]
	}

	return Queue{
		head:   Element{},
		length: 0,
		full:   size,
	}
}
