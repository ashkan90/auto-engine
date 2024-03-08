
package builtins

import (
	"github.com/ashkan90/auto-core/src"
)

type Op_LteNode struct {
	src.NodeInterface
}

func NewOp_LteNode() src.NodeInterface {
	return &Op_LteNode{
		NodeInterface: src.NewNode(),
	}
}

// Execute no need to write unless it executes something
func (n *Op_LteNode) Execute(input string, forward func(output string)) {
	// do something here
	forward("exec")
}