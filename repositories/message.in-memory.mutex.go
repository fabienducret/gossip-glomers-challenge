package repositories

import "sync"

type MessageRepositoryInMemoryMutex struct {
	mutex    sync.Mutex
	original *MessageRepositoryInMemory
}

func NewMessageRepositoryInMemoryMutex(original *MessageRepositoryInMemory) *MessageRepositoryInMemoryMutex {
	return &MessageRepositoryInMemoryMutex{original: original}
}

func (repo *MessageRepositoryInMemoryMutex) Add(value any) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	repo.original.Add(value)
}

func (repo *MessageRepositoryInMemoryMutex) Contains(value any) bool {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	return repo.original.Contains(value)
}

func (repo *MessageRepositoryInMemoryMutex) Get() []any {
	return repo.original.Get()
}
