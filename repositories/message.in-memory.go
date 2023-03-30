package repositories

import "sync"

type MessageRepositoryInMemory struct {
	mutex  sync.Mutex
	values []any
}

func NewMessageRepositoryInMemory() *MessageRepositoryInMemory {
	return &MessageRepositoryInMemory{}
}

func (ms *MessageRepositoryInMemory) Add(value any) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()
	ms.values = append(ms.values, value)
}

func (ms *MessageRepositoryInMemory) Contains(value any) bool {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	for _, v := range ms.values {
		if v == value {
			return true
		}
	}

	return false
}

func (ms *MessageRepositoryInMemory) Get() []any {
	return ms.values
}
