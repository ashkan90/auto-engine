
package builtins

import (
	"github.com/ashkan90/auto-core/src"
)

type TextNode struct {
	src.NodeInterface
}

func NewTextNode() src.NodeInterface {
	return &TextNode{
		NodeInterface: src.NewNode(),
	}
}

// Execute no need to write unless it executes something
func (n *TextNode) Execute(input string, forward func(output string)) {
	// do something here
	forward("exec")
}