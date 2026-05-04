package main

import "errors"

var ErrNotFound = errors.New("Key Not Found")

// ! Struct for Store
type Store struct {
	data map[string][]byte
}

// ! Create New Store
func NewStore() *Store {
	return &Store{
		data: make(map[string][]byte),
	}
}

// ! Insert the Key or Overwrite the Key
func (s *Store) Put(key string, value []byte) {
	s.data[key] = value
}

// ! Get the Value for key
func (s *Store) Get(key string) ([]byte, error) {
	val, ok := s.data[key]
	if !ok {
		return nil, ErrNotFound
	}
	return val, nil
}

// ! Delete the key
func (s *Store) Delete(key string) {
	delete(s.data, key)
}
