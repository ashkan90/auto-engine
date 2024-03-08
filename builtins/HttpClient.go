
package builtins

import (
	"github.com/ashkan90/auto-core/src"
)

type HttpClientNode struct {
	src.NodeInterface
}

func NewHttpClientNode() src.NodeInterface {
	return &HttpClientNode{
		NodeInterface: src.NewNode(),
	}
}

// Execute no need to write unless it executes something
func (n *HttpClientNode) Execute(input string, forward func(output string)) {
	// do something here
	forward("exec")
}