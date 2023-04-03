package repositories

type MessageRepositoryInMemory struct {
	values []int
}

func NewMessageRepositoryInMemory() *MessageRepositoryInMemory {
	return &MessageRepositoryInMemory{}
}

func (repo *MessageRepositoryInMemory) Add(value int) {
	repo.values = append(repo.values, value)
}

func (repo *MessageRepositoryInMemory) Contains(value int) bool {
	for _, v := range repo.values {
		if v == value {
			return true
		}
	}

	return false
}

func (repo *MessageRepositoryInMemory) Get() []int {
	return repo.values
}
