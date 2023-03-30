package repositories

import "sync"

type TopologyRepositoryInMemory struct {
	mutex  sync.Mutex
	values []string
}

func NewTopologyRepositoryInMemory() *TopologyRepositoryInMemory {
	return &TopologyRepositoryInMemory{}
}

func (tr *TopologyRepositoryInMemory) Add(value string) {
	tr.mutex.Lock()
	defer tr.mutex.Unlock()
	tr.values = append(tr.values, value)
}

func (tr *TopologyRepositoryInMemory) Get() []string {
	return tr.values
}
