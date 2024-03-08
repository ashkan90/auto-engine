
package builtins

import (
	"github.com/ashkan90/auto-core/src"
)

type Math_MaxNode struct {
	src.NodeInterface
}

func NewMath_MaxNode() src.NodeInterface {
	return &Math_MaxNode{
		NodeInterface: src.NewNode(),
	}
}

// Execute no need to write unless it executes something
func (n *Math_MaxNode) Execute(input string, forward func(output string)) {
	// do something here
	forward("exec")
}