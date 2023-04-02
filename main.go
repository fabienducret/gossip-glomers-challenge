package main

import (
	"log"
	"maelstrom-broadcast/handlers"
	"maelstrom-broadcast/repositories"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()

	mr := repositories.NewMessageRepositoryInMemorySync(repositories.NewMessageRepositoryInMemory())
	tr := repositories.NewTopologyRepositoryInMemorySync(repositories.NewTopologyRepositoryInMemory())

	n.Handle("topology", handlers.TopologyMessageFactory(n, tr))
	n.Handle("broadcast", handlers.BroadcastMessageFactory(n, mr, tr))
	n.Handle("read", handlers.ReadMessageFactory(n, mr))

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
