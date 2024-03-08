package engine

type ControlFlowNode struct {
	inputs  []string
	outputs []string
	execute func(input string, forward func(output string))
}

func NewControlFlowNode(i []string, o []string, e func(input string, forward func(output string))) *ControlFlowNode {
	return &ControlFlowNode{
		inputs:  i,
		outputs: o,
		execute: e,
	}
}

func (cn *ControlFlowNode) Inputs() []string {
	return cn.inputs
}

func (cn *ControlFlowNode) Outputs() []string {
	return cn.outputs
}

func (cn *ControlFlowNode) Execute(input string, forward func(output string)) {
	cn.execute(input, forward)
}
