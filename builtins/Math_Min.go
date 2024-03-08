
package builtins

import (
	"github.com/ashkan90/auto-core/src"
)

type Math_MinNode struct {
	src.NodeInterface
}

func NewMath_MinNode() src.NodeInterface {
	return &Math_MinNode{
		NodeInterface: src.NewNode(),
	}
}

// Execute no need to write unless it executes something
func (n *Math_MinNode) Execute(input string, forward func(output string)) {
	// do something here
	forward("exec")
}