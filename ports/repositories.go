package ports

type MessageRepository interface {
	Add(value any)
	Contains(value any) bool
	Get() []any
}

type TopologyRepository interface {
	Add(value string)
	Neighbours() []string
}
