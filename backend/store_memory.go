package main

import (
	"context"
	"strings"
	"sync"
)

type MemoryClassStore struct {
	mu      sync.RWMutex
	classes []BoatClass
}

func NewMemoryClassStore() *MemoryClassStore {
	return &MemoryClassStore{classes: []BoatClass{}}
}

func (s *MemoryClassStore) List(ctx context.Context) ([]BoatClass, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]BoatClass(nil), s.classes...), nil
}

func (s *MemoryClassStore) Add(ctx context.Context, bc BoatClass) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, existing := range s.classes {
		if strings.EqualFold(existing.Name, bc.Name) {
			return ErrClassAlreadyExists
		}
	}
	s.classes = append(s.classes, bc)
	return nil
}

func (s *MemoryClassStore) DeleteByName(ctx context.Context, name string) (BoatClass, bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, existing := range s.classes {
		if strings.EqualFold(existing.Name, name) {
			deleted := existing
			s.classes = append(s.classes[:i], s.classes[i+1:]...)
			return deleted, true, nil
		}
	}
	return BoatClass{}, false, nil
}
