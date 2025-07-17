package models

import "sync"

type SafeStore struct {
	mu    sync.Mutex
	store map[string][]int
}

// Не до конца понял что эта функция делает, но она нужна. Типа инициализирует пустую мапу
func NewSafeStore() *SafeStore {
	return &SafeStore{
		store: map[string][]int{},
	}
}

// Медот записи в мапу
func (s *SafeStore) Save(key string, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[key] = append(s.store[key], value)
}

// Метод чтения из мапы
func (s *SafeStore) Get(key string) ([]int, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	val, ok := s.store[key]
	return val, ok
}
