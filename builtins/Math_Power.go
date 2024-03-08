
package builtins

import (
	"github.com/ashkan90/auto-core/src"
)

type Math_PowerNode struct {
	src.NodeInterface
}

func NewMath_PowerNode() src.NodeInterface {
	return &Math_PowerNode{
		NodeInterface: src.NewNode(),
	}
}

// Execute no need to write unless it executes something
func (n *Math_PowerNode) Execute(input string, forward func(output string)) {
	// do something here
	forward("exec")
}