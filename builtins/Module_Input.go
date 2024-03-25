package builtins

import (
	"github.com/ashkan90/auto-core/src"
	"github.com/ashkan90/auto-engine/engine"
)

type ModuleInputNode struct {
	InputValue any

	src.NodeInterface
	dfe *engine.DataflowEngine
}

func NewModuleInputNode(dfe *engine.DataflowEngine) src.NodeInterface {
	return &ModuleInputNode{
		NodeInterface: src.NewNode(),
		dfe:           dfe,
	}
}

func (n *ModuleInputNode) Data(_ func() map[string]any) map[string]any {
	return n.Node().Controls.ToMap()
}
