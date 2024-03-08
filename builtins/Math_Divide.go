
package builtins

import (
	"github.com/ashkan90/auto-core/src"
)

type Math_DivideNode struct {
	src.NodeInterface
}

func NewMath_DivideNode() src.NodeInterface {
	return &Math_DivideNode{
		NodeInterface: src.NewNode(),
	}
}

// Execute no need to write unless it executes something
func (n *Math_DivideNode) Execute(input string, forward func(output string)) {
	// do something here
	forward("exec")
}