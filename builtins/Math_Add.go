
package builtins

import (
	"github.com/ashkan90/auto-core/src"
)

type Math_AddNode struct {
	src.NodeInterface
}

func NewMath_AddNode() src.NodeInterface {
	return &Math_AddNode{
		NodeInterface: src.NewNode(),
	}
}

// Execute no need to write unless it executes something
func (n *Math_AddNode) Execute(input string, forward func(output string)) {
	// do something here
	forward("exec")
}