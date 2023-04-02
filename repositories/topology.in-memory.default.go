package repositories

type TopologyRepositoryInMemory struct {
	values []string
}

func NewTopologyRepositoryInMemory() *TopologyRepositoryInMemory {
	return &TopologyRepositoryInMemory{}
}

func (tr *TopologyRepositoryInMemory) Add(value string) {
	tr.values = append(tr.values, value)
}

func (tr *TopologyRepositoryInMemory) Get() []string {
	return tr.values
}
