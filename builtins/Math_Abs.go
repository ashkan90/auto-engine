
package builtins

import (
	"github.com/ashkan90/auto-core/src"
)

type Math_AbsNode struct {
	src.NodeInterface
}

func NewMath_AbsNode() src.NodeInterface {
	return &Math_AbsNode{
		NodeInterface: src.NewNode(),
	}
}

// Execute no need to write unless it executes something
func (n *Math_AbsNode) Execute(input string, forward func(output string)) {
	// do something here
	forward("exec")
}