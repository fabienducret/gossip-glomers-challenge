package handlers

import (
	"encoding/json"
	"maelstrom-broadcast/ports"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func TopologyMessageFactory(n *maelstrom.Node, tr ports.TopologyRepository) func(msg maelstrom.Message) error {
	return func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		for k, v := range body["topology"].(map[string]any) {
			if k == n.ID() {
				for _, nodeName := range v.([]any) {
					tr.Add(nodeName.(string))
				}
			}
		}

		var reply map[string]string = map[string]string{"type": "topology_ok"}
		return n.Reply(msg, reply)
	}
}
