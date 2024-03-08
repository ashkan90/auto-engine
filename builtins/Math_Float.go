
package builtins

import (
	"github.com/ashkan90/auto-core/src"
)

type Math_FloatNode struct {
	src.NodeInterface
}

func NewMath_FloatNode() src.NodeInterface {
	return &Math_FloatNode{
		NodeInterface: src.NewNode(),
	}
}

// Execute no need to write unless it executes something
func (n *Math_FloatNode) Execute(input string, forward func(output string)) {
	// do something here
	forward("exec")
}