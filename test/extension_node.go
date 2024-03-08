package test

import (
	"github.com/ashkan90/auto-core/src"
	"log"
)

type ExtensionNode struct {
	src.NodeInterface
}

func NewExtensionNode() src.NodeInterface {
	return &ExtensionNode{
		NodeInterface: src.NewNode(),
	}
}

// Execute no need to write unless it executes something
func (n *ExtensionNode) Execute(input string, forward func(output string)) {
	log.Println("[EXTENSION_NODE]")
	// do something here
	forward("exec")
}

// no need to write unless it returns data
//func (n *ExtensionNode) Data(inputs func() map[string]any) map[string]any {
//	return nil
//}
