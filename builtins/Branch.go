package builtins

import (
	"github.com/ashkan90/auto-core/src"
)

type BranchNode struct {
	src.NodeInterface
}

func NewBranchNode() src.NodeInterface {
	node := &BranchNode{
		NodeInterface: src.NewNode(),
	}
	node.AddInput("exec", src.NewInput(src.NewSocket("empty"), "Exec Label", false))
	node.AddInput("condition", src.NewInput(src.NewSocket("empty"), "Condition", false))

	node.AddOutput("exec", src.NewOutput(src.NewSocket("empty"), "Exec Label", false))
	node.AddOutput("true", src.NewOutput(src.NewSocket("empty"), "True Label", false))
	node.AddOutput("false", src.NewOutput(src.NewSocket("empty"), "False Label", false))
	return node
}

// Execute no need to write unless it executes something
func (n *BranchNode) Execute(input string, forward func(output string)) {
	// do something here
	forward("exec")
}
