package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type NodeContainer struct {
	Node     *maelstrom.Node
	Messages []int
}

type BroadcastMessage struct {
	Type    string
	Message int
}

type ReadMessage struct {
	Type  string
	MsgID int
}

type TypologyMessage struct {
	Type      string         `json:type`
	InReplyTo string         `json:in_reply_to`
	MsgID     int            `json:msg_id`
	Topology  map[string]any `json:topology`
}

type BroadcastResponse struct {
	Type string `json:"type,omitempty"`
}

type ReadResponse struct {
	Type      string `json:"type,omitempty"`
	Messages  []int  `json:"messages"`
	MessageId int    `json:"k,omitempty"`
}

type TopologyResponse struct {
	Type      string `json:"type,omitempty"`
	MessageId int    `json:"k,omitempty"`
	InReplyTo string `json:"in_reply_to,omitempty"`
}

func (nc *NodeContainer) broadcast(messageId int) *BroadcastResponse {
	nc.Messages = append(nc.Messages, messageId)
	return &BroadcastResponse{Type: "broadcast_ok"}
}

func (nc *NodeContainer) read(messageId int) *ReadResponse {
	return &ReadResponse{"read_ok", nc.Messages, messageId}
}

func (nc *NodeContainer) topology(messageId int, in_reply_to string) *TopologyResponse {
	return &TopologyResponse{"topology_ok", messageId, in_reply_to}
}

func main() {
	n := maelstrom.NewNode()
	node_container := NodeContainer{Node: n}

	node_container.Node.Handle("broadcast", func(msg maelstrom.Message) error {
		var bm BroadcastMessage
		if err := json.Unmarshal(msg.Body, &bm); err != nil {
			return err
		}
		response := node_container.broadcast(bm.Message)
		return node_container.Node.Reply(msg, response)
	})

	node_container.Node.Handle("read", func(msg maelstrom.Message) error {
		var rm ReadMessage
		if err := json.Unmarshal(msg.Body, &rm); err != nil {
			return err
		}
		response := node_container.read(rm.MsgID)
		return node_container.Node.Reply(msg, response)
	})

	node_container.Node.Handle("topology", func(msg maelstrom.Message) error {
		var tm TypologyMessage
		if err := json.Unmarshal(msg.Body, &tm); err != nil {
			return err
		}
		response := node_container.topology(tm.MsgID, tm.InReplyTo)
		return node_container.Node.Reply(msg, response)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
