package gotl

type Set[T comparable] interface {
	Has(elem T) bool
	Add(elem T)
	Remove(elem T)
	Len() int
	Items() []T
}

func NewSet[T comparable]() Set[T] {
	return &set[T]{
		impl: map[T]struct{}{},
	}
}

type set[T comparable] struct {
	impl map[T]struct{}
}

func (s *set[T]) Has(elem T) bool {
	_, hasKey := s.impl[elem]
	return hasKey
}

func (s *set[T]) Add(elem T) {
	s.impl[elem] = struct{}{}
}

func (s *set[T]) Remove(elem T) {
	delete(s.impl, elem)
}

func (s set[T]) Len() int {
	return len(s.impl)
}

func (s *set[T]) Items() []T {
	output := make([]T, 0, len(s.impl))
	for k, _ := range s.impl {
		output = append(output, k)
	}

	return output
}
