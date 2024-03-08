
package builtins

import (
	"github.com/ashkan90/auto-core/src"
)

type Math_RoundNode struct {
	src.NodeInterface
}

func NewMath_RoundNode() src.NodeInterface {
	return &Math_RoundNode{
		NodeInterface: src.NewNode(),
	}
}

// Execute no need to write unless it executes something
func (n *Math_RoundNode) Execute(input string, forward func(output string)) {
	// do something here
	forward("exec")
}