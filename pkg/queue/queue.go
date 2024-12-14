package queue

/*
Default maxSize = 10
*/
func NewClassicQueue(maxSize ...int) ClassicQueue {
	size := 10
	var queue ClassicQueue

	if len(maxSize) > 0 {
		size = maxSize[0]
	}

	queue = ClassicQueue{
		Queue{
			Type:   "classic",
			Head:   Element{},
			Tail:   Element{},
			Length: 0,
			Full:   size,
		},
	}

	return queue
}

/*
Default maxSize = 10
*/
func NewCircleQueue(maxSize ...int) CircleQueue {
	size := 10
	var queue CircleQueue

	if len(maxSize) > 0 {
		size = maxSize[0]
	}

	queue = CircleQueue{
		Queue{
			Type:   "circle",
			Head:   Element{},
			Tail:   Element{},
			Length: 0,
			Full:   size,
		},
	}

	return queue
}

func NewElement(value []EContent) Element {
	return Element{
		Value: value,
		next:  nil,
	}
}
