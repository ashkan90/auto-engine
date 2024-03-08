
package builtins

import (
	"github.com/ashkan90/auto-core/src"
)

type BranchNode struct {
	src.NodeInterface
}

func NewBranchNode() src.NodeInterface {
	return &BranchNode{
		NodeInterface: src.NewNode(),
	}
}

// Execute no need to write unless it executes something
func (n *BranchNode) Execute(input string, forward func(output string)) {
	// do something here
	forward("exec")
}