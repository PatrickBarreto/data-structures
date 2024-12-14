package queue

// Cada tipo de fila tem um indicador de cabeça
func (queue *Queue) getHead() *Element {
	return &queue.Head
}

// Cada tipo de fila tem um indicador de última posição?
func (queue *Queue) getTail() Element {
	return queue.Tail
}

// Cada tipo lista pode saber se é vazia
func (queue *Queue) Empty() bool {
	return queue.Length == 0
}

// Cada tipo lista pode saber seu tamanho
func (queue *Queue) Size() int {
	return queue.Length
}

// Cada tipo lista pode saber se está cheia
func (queue *Queue) isFull() bool {
	return queue.Length >= queue.Full
}
