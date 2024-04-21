package cache

import "sync"

type Cache interface {
	Get(key string) item[any]
	Put(key string, value item[any]) error
	Delete(key string) error
	Expire(key string) int64
}

type item[T any] struct {
	key   string
	value T
}

type LocalCache[T item[any]] struct {
	data map[string]T
	mux  sync.RWMutex
}

func (l *LocalCache[T]) Get(key string) T {
	return T{}
}

func (l *LocalCache[T]) Put(key string, value T) error {
	//TODO implement me
	panic("implement me")
}

func (l *LocalCache[T]) Delete(key string) error {
	//TODO implement me
	panic("implement me")
}

func (l *LocalCache[T]) Expire(key string) int64 {
	//TODO implement me
	panic("implement me")
}
