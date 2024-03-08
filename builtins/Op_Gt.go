
package builtins

import (
	"github.com/ashkan90/auto-core/src"
)

type Op_GtNode struct {
	src.NodeInterface
}

func NewOp_GtNode() src.NodeInterface {
	return &Op_GtNode{
		NodeInterface: src.NewNode(),
	}
}

// Execute no need to write unless it executes something
func (n *Op_GtNode) Execute(input string, forward func(output string)) {
	// do something here
	forward("exec")
}