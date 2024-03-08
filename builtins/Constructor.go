
package builtins

import (
	"github.com/ashkan90/auto-core/src"
)

type ConstructorNode struct {
	src.NodeInterface
}

func NewConstructorNode() src.NodeInterface {
	return &ConstructorNode{
		NodeInterface: src.NewNode(),
	}
}

// Execute no need to write unless it executes something
func (n *ConstructorNode) Execute(input string, forward func(output string)) {
	// do something here
	forward("exec")
}