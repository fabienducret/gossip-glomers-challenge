package handlers

import (
	"encoding/json"
	"maelstrom-broadcast/ports"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func sendToNeighbours(n *maelstrom.Node, body map[string]any, tr ports.TopologyRepository) {
	neighbours := tr.Neighbours()

	for _, neighbour := range neighbours {
		n.RPC(neighbour, body, nil)
	}
}

func BroadcastMessageFactory(n *maelstrom.Node, mr ports.MessageRepository, tr ports.TopologyRepository) func(msg maelstrom.Message) error {
	return func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		message := body["message"]

		if mr.Contains(message) {
			return nil
		}

		mr.Add(message)
		sendToNeighbours(n, body, tr)

		var reply map[string]string = map[string]string{"type": "broadcast_ok"}
		return n.Reply(msg, reply)
	}
}
