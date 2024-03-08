
package builtins

import (
	"github.com/ashkan90/auto-core/src"
)

type Op_NotNode struct {
	src.NodeInterface
}

func NewOp_NotNode() src.NodeInterface {
	return &Op_NotNode{
		NodeInterface: src.NewNode(),
	}
}

// Execute no need to write unless it executes something
func (n *Op_NotNode) Execute(input string, forward func(output string)) {
	// do something here
	forward("exec")
}