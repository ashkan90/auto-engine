
package builtins

import (
	"github.com/ashkan90/auto-core/src"
)

type Op_EqNode struct {
	src.NodeInterface
}

func NewOp_EqNode() src.NodeInterface {
	return &Op_EqNode{
		NodeInterface: src.NewNode(),
	}
}

// Execute no need to write unless it executes something
func (n *Op_EqNode) Execute(input string, forward func(output string)) {
	// do something here
	forward("exec")
}