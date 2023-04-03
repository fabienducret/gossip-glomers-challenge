package repositories

import "sync"

type MessageRepositoryInMemorySync struct {
	mutex    sync.Mutex
	original *MessageRepositoryInMemory
}

func NewMessageRepositoryInMemorySync(original *MessageRepositoryInMemory) *MessageRepositoryInMemorySync {
	return &MessageRepositoryInMemorySync{original: original}
}

func (repo *MessageRepositoryInMemorySync) Add(value int) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	repo.original.Add(value)
}

func (repo *MessageRepositoryInMemorySync) Contains(value int) bool {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	return repo.original.Contains(value)
}

func (repo *MessageRepositoryInMemorySync) Get() []int {
	return repo.original.Get()
}
