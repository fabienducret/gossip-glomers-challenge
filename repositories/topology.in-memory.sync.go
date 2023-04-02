package repositories

import "sync"

type TopologyRepositoryInMemorySync struct {
	mutex    sync.Mutex
	original *TopologyRepositoryInMemory
}

func NewTopologyRepositoryInMemorySync(original *TopologyRepositoryInMemory) *TopologyRepositoryInMemorySync {
	return &TopologyRepositoryInMemorySync{original: original}
}

func (repo *TopologyRepositoryInMemorySync) Add(value string) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	repo.original.Add(value)
}

func (repo *TopologyRepositoryInMemorySync) Get() []string {
	return repo.original.Get()
}
