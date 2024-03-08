
package builtins

import (
	"github.com/ashkan90/auto-core/src"
)

type Op_OrNode struct {
	src.NodeInterface
}

func NewOp_OrNode() src.NodeInterface {
	return &Op_OrNode{
		NodeInterface: src.NewNode(),
	}
}

// Execute no need to write unless it executes something
func (n *Op_OrNode) Execute(input string, forward func(output string)) {
	// do something here
	forward("exec")
}