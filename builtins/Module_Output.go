package builtins

import (
	"github.com/ashkan90/auto-core/src"
)

type ModuleOutputNode struct {
	src.NodeInterface
}

func NewModuleOutputNode() src.NodeInterface {
	return &ModuleOutputNode{
		NodeInterface: src.NewNode(),
	}
}

func (n *ModuleOutputNode) Data(inputs func() map[string]any) map[string]any {
	var inputValues = inputs()
	return map[string]any{
		"value": inputValues["input"],
	}
}
