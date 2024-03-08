
package builtins

import (
	"github.com/ashkan90/auto-core/src"
)

type Op_LtNode struct {
	src.NodeInterface
}

func NewOp_LtNode() src.NodeInterface {
	return &Op_LtNode{
		NodeInterface: src.NewNode(),
	}
}

// Execute no need to write unless it executes something
func (n *Op_LtNode) Execute(input string, forward func(output string)) {
	// do something here
	forward("exec")
}