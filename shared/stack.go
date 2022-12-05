package shared

type Stack[T interface{}] struct {
	stack []T
}

// Create a new stack containing elements of type T
func NewStack[T interface{}]() Stack[T] {
	s := Stack[T]{}
	s.stack = make([]T, 0)
	return s
}

func (s *Stack[T]) Pop() T {
	return s.PopMany(1)[0]
}

// Pops the num topmost elements, returning a slice with the topmost element last
func (s *Stack[T]) PopMany(num int) []T {
	n := len(s.stack) - num
	elements := s.stack[n:]
	s.stack = s.stack[:n]
	return elements
}

// Put an element at the top of the stack
func (s *Stack[T]) Put(element T) {
	s.stack = append(s.stack, element)
}

// Puts elements onto the stack, leaving the last element of elements at the top
func (s *Stack[T]) PutMany(elements []T) {
	s.stack = append(s.stack, elements...)
}

// Look at the topmost element without popping
func (s *Stack[T]) Peek() T {
	n := len(s.stack) - 1
	return s.stack[n]
}
