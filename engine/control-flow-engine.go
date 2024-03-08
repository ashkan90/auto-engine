package engine

import (
	"github.com/ashkan90/auto-core/src"
)

// Önceden tanımlanmış Node, Connection ve ClassicScheme yapılarını kullanıyoruz.

type ControlFlowEngine struct {
	controlFlow *ControlFlow
}

type ControlFlowEngineExecutor interface {
	Execute(nodeID src.NodeId, input string)
}

func NewControlFlowEngine(editor *src.NodeEditor) *ControlFlowEngine {
	return &ControlFlowEngine{
		controlFlow: NewControlFlow(editor),
	}
}

// Configure, ControlFlow için node setup'ını yapılandırmayı sağlar.
func (cfe *ControlFlowEngine) Add(node src.NodeInterface) {
	n := node.Node()
	inputs := func() []string {
		keys := make([]string, 0, n.Inputs.Len())
		n.Inputs.Range(func(key, _ any) bool {
			keys = append(keys, key.(string))
			return true
		})
		return keys
	}()
	outputs := func() []string {
		keys := make([]string, 0, n.Outputs.Len())
		n.Outputs.Range(func(key, _ any) bool {
			keys = append(keys, key.(string))
			return true
		})
		return keys
	}()

	cfe.controlFlow.Add(node, NewControlFlowNode(inputs, outputs, func(input string, forward func(output string)) {
		node.Execute(input, forward)
	}))
}

// Execute, belirli bir node'dan başlayarak akışı tetikler.
func (cfe *ControlFlowEngine) Execute(nodeID src.NodeId, input string) {
	cfe.controlFlow.Execute(nodeID, input)
}
