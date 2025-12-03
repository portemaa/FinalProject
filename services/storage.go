package services

import "sync"

// in memory password storage

type StorageService interface {
    SavePassword(p string)
    GetAll() []string
}

type MemStorage struct {
    mu        sync.Mutex
    passwords []string
}

func NewMemStorage() *MemStorage {
    return &MemStorage{
        passwords: []string{},
    }
}

func (s *MemStorage) SavePassword(p string) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.passwords = append(s.passwords, p)
}

func (s *MemStorage) GetAll() []string {
    s.mu.Lock()
    defer s.mu.Unlock()

    // return a copy
    rcp := make([]string, len(s.passwords))
    copy(rcp, s.passwords)
    return rcp
}
