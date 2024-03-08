
package builtins

import (
	"github.com/ashkan90/auto-core/src"
)

type Math_MultiplyNode struct {
	src.NodeInterface
}

func NewMath_MultiplyNode() src.NodeInterface {
	return &Math_MultiplyNode{
		NodeInterface: src.NewNode(),
	}
}

// Execute no need to write unless it executes something
func (n *Math_MultiplyNode) Execute(input string, forward func(output string)) {
	// do something here
	forward("exec")
}