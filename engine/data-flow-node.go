package engine

type DataflowNode struct {
	inputs  []string
	outputs []string
	data    func(inputs func() map[string]any) map[string]any
}

func NewDataflowNode(i []string, o []string, d func(inputs func() map[string]any) map[string]any) *DataflowNode {
	return &DataflowNode{
		inputs:  i,
		outputs: o,
		data:    d,
	}
}

func (cn *DataflowNode) Inputs() []string {
	return cn.inputs
}

func (cn *DataflowNode) Outputs() []string {
	return cn.outputs
}

func (cn *DataflowNode) Data(inputs func() map[string]any) map[string]any {
	return cn.data(inputs)
}
