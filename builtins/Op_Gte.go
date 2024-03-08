
package builtins

import (
	"github.com/ashkan90/auto-core/src"
)

type Op_GteNode struct {
	src.NodeInterface
}

func NewOp_GteNode() src.NodeInterface {
	return &Op_GteNode{
		NodeInterface: src.NewNode(),
	}
}

// Execute no need to write unless it executes something
func (n *Op_GteNode) Execute(input string, forward func(output string)) {
	// do something here
	forward("exec")
}