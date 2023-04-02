package main

import (
	"log"
	"maelstrom-broadcast/handlers"
	"maelstrom-broadcast/repositories"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()

	mr := repositories.NewMessageRepositoryInMemoryMutex(repositories.NewMessageRepositoryInMemory())
	tr := repositories.NewTopologyRepositoryInMemory()

	n.Handle("topology", handlers.TopologyMessageFactory(n, tr))
	n.Handle("broadcast", handlers.BroadcastMessageFactory(n, mr, tr))
	n.Handle("read", handlers.ReadMessageFactory(n, mr))

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
