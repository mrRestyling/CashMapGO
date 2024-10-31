package storage

import (
	"sync"
	"time"
)

type Storage struct {
	List map[string]interface{}
	mu   sync.Mutex
	l    string
}

type Item struct {
	Value interface{}
	Ttl   time.Time
}

func New(m map[string]any) Storage {
	return Storage{
		List: m,
	}
}

func (s *Storage) Set(key string, value any, ttl time.Duration) {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.List[key] = Item{
		Value: value,
		Ttl:   time.Now().Add(ttl),
	}

}

func (s *Storage) Get(key string) (interface{}, bool) {

	s.mu.Lock()
	defer s.mu.Unlock()

	item, ok := s.List[key]
	if !ok {
		return nil, false
	}

	result := item.(Item)

	if time.Now().After(result.Ttl) {
		delete(s.List, key)
		return nil, false
	}

	return result.Value, true

}
