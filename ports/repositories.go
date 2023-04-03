package ports

type MessageRepository interface {
	Add(value int)
	Contains(value int) bool
	Get() []int
}

type TopologyRepository interface {
	Add(value string)
	Neighbours() []string
}
