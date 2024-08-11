package gotl

type Set[T comparable] struct {
	impl map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		impl: map[T]struct{}{},
	}
}

func (s *Set[T]) Has(elem T) bool {
	_, hasKey := s.impl[elem]
	return hasKey
}

func (s *Set[T]) Add(elem T) {
	s.impl[elem] = struct{}{}
}

func (s *Set[T]) Remove(elem T) {
	delete(s.impl, elem)
}

func (s Set[T]) Len() int {
	return len(s.impl)
}

func (s *Set[T]) Items() []T {
	output := make([]T, 0, len(s.impl))
	for k, _ := range s.impl {
		output = append(output, k)
	}

	return output
}
