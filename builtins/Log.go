package builtins

import (
	"github.com/ashkan90/auto-core/src"
	"github.com/ashkan90/auto/engine"
	"log"
)

type LogNode struct {
	src.NodeInterface
	*engine.DataflowEngine
}

func NewLogNode(dfe *engine.DataflowEngine) src.NodeInterface {
	return &LogNode{
		NodeInterface:  src.NewNode(),
		DataflowEngine: dfe,
	}
}

// Execute no need to write unless it executes something
func (n *LogNode) Execute(_ string, forward func(_ string)) {
	var inputs = n.DataflowEngine.FetchInputs(n.NodeInterface.Node().E.ID)
	var message = inputs["message"].(*string)

	log.Println("[NODE_LOGGER]", *message)

	forward("exec")
}
