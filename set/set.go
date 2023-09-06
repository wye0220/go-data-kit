package set

type Set[T comparable] interface {
	Add(key T)
	Remove(key T)
	Contains(key T) bool
	Len() int
	Clear()
	Keys() []T
}

type HashSet[T comparable] struct {
	m map[T]struct{}
}

func NewHashSet[T comparable](size int) *HashSet[T] {
	return &HashSet[T]{
		m: make(map[T]struct{}, size),
	}
}

func (s *HashSet[T]) Add(key T) {
	s.m[key] = struct{}{}
}

func (s *HashSet[T]) Remove(key T) {
	delete(s.m, key)
}

func (s *HashSet[T]) Contains(key T) bool {
	_, ok := s.m[key]
	return ok
}

func (s *HashSet[T]) Len() int {
	return len(s.m)
}

func (s *HashSet[T]) Clear() {
	s.m = make(map[T]struct{})
}

func (s *HashSet[T]) Keys() []T {
	res := make([]T, 0, len(s.m))
	for k := range s.m {
		res = append(res, k)
	}
	return res
}
