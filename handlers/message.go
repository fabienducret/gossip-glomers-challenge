package handlers

import (
	"maelstrom-broadcast/ports"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func ReadMessageFactory(n *maelstrom.Node, mr ports.MessageRepository) func(msg maelstrom.Message) error {
	return func(msg maelstrom.Message) error {
		var reply map[string]any = map[string]any{
			"type":     "read_ok",
			"messages": mr.Get(),
		}

		return n.Reply(msg, reply)
	}
}
