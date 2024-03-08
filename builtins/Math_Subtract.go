
package builtins

import (
	"github.com/ashkan90/auto-core/src"
)

type Math_SubtractNode struct {
	src.NodeInterface
}

func NewMath_SubtractNode() src.NodeInterface {
	return &Math_SubtractNode{
		NodeInterface: src.NewNode(),
	}
}

// Execute no need to write unless it executes something
func (n *Math_SubtractNode) Execute(input string, forward func(output string)) {
	// do something here
	forward("exec")
}