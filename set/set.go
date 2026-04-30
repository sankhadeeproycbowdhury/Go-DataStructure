package set

import "sync"

type Set[T comparable] struct {
	mu   sync.RWMutex
	data map[T]struct{}
}

func New[T comparable]() *Set[T] {
	return &Set[T]{
        data: make(map[T]struct{}),
    }

}

func (s *Set[T]) Add(value T) {
	s.mu.Lock()
    s.data[value] = struct{}{}
    s.mu.Unlock()

}

func (s *Set[T]) Remove(value T) {
	s.mu.Lock()
    delete(s.data, value)
    s.mu.Unlock()

}

func (s *Set[T]) Contains(value T) bool {
	s.mu.RLock()
    _, exists := s.data[value]
    s.mu.RUnlock()
    return exists
}

func (s *Set[T]) Size() int {
	s.mu.RLock()
	size := len(s.data)
	s.mu.RUnlock()
	return size
}

func (s *Set[T]) Clear() {
	s.mu.Lock()
	s.data = make(map[T]struct{})
	s.mu.Unlock()		
}

func (s *Set[T]) Values() []T {
	s.mu.RLock()
	values := make([]T, 0, len(s.data))			
	for key := range s.data {
		values = append(values, key)
	}
	s.mu.RUnlock()
	return values
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := New[T]()
	s.mu.RLock()
	for key := range s.data {
		result.Add(key)
	}
	s.mu.RUnlock()

	other.mu.RLock()
	for key := range other.data {
		result.Add(key)
	}
	other.mu.RUnlock()

	return result
}

func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := New[T]()
	s.mu.RLock()
	for key := range s.data {
		if other.Contains(key) {
			result.Add(key)
		}
	}
	s.mu.RUnlock()
	return result
}

func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := New[T]()
	s.mu.RLock()				
	for key := range s.data {
		if !other.Contains(key) {
			result.Add(key)
		}
	}
	s.mu.RUnlock()
	return result
}	


