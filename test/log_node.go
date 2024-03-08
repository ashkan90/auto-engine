package test

import (
	"github.com/ashkan90/auto-core/src"
	"github.com/ashkan90/auto/engine"
	"log"
)

type LogNode struct {
	src.NodeInterface
	DataflowEngine *engine.DataflowEngine
}

func (n *LogNode) Data(inputs func() map[string]any) map[string]any {
	return nil
}

func (n *LogNode) Execute(input string, forward func(output string)) {
	var inputs = n.DataflowEngine.FetchInputs(n.NodeInterface.Node().E.ID)
	var message = inputs["message"].(*string)

	log.Println("[NODE_LOGGER]", *message)

	forward("exec")
}

func NewLogNode(dataflow *engine.DataflowEngine) src.NodeInterface {
	return &LogNode{
		NodeInterface:  src.NewNode(),
		DataflowEngine: dataflow,
	}
}
