package repositories

type MessageRepositoryInMemory struct {
	values []any
}

func NewMessageRepositoryInMemory() *MessageRepositoryInMemory {
	return &MessageRepositoryInMemory{}
}

func (repo *MessageRepositoryInMemory) Add(value any) {
	repo.values = append(repo.values, value)
}

func (repo *MessageRepositoryInMemory) Contains(value any) bool {
	for _, v := range repo.values {
		if v == value {
			return true
		}
	}

	return false
}

func (repo *MessageRepositoryInMemory) Get() []any {
	return repo.values
}
