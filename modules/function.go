package modules

import (
	"fmt"
	"github.com/ashkan90/auto-core/src"
	"github.com/ashkan90/auto-engine/builtins"
	"github.com/ashkan90/auto-engine/engine"
	"log"
)

type FunctionInterface interface {
	//Apply()
	Execute(data map[string]any) any
}

type Function[N src.NodeInterface] struct {
	Node N `json:"node"`
}

func NewFunctionNode[N src.NodeInterface]() *Function[N] {
	return &Function[N]{}
}

func (f *Function[N]) FindFunction() FunctionInterface {
	return f
}

func (f *Function[N]) Execute(data map[string]any) any {
	var bus = src.NewEventBus()
	var editor = src.NewNodeEditor(bus)
	var _engine = engine.NewDataflowEngine(editor)

	return f.execute(data, editor, _engine)
}

func (f *Function[N]) execute(
	inputs map[string]any,
	editor *src.NodeEditor,
	_engine *engine.DataflowEngine,
) any {
	var nodes = editor.GetNodes()

	f.injectInputs(nodes, inputs)
	return f.retrieveOutputs(nodes, _engine)
}

func (f *Function[N]) isInputNode(node src.NodeInterface) bool {
	var _, ok = node.(*builtins.ModuleInputNode)
	return node.FromModule() && ok
}

func (f *Function[N]) injectInputs(nodes []src.NodeInterface, inputs map[string]any) {
	var inputNodes []src.NodeInterface

	for _, node := range nodes {
		if f.isInputNode(node) {
			inputNodes = append(inputNodes, node)
		}
	}

	for _, node := range inputNodes {
		var ctrl, ok = node.Node().Controls.Get("name")
		if ok {
			var key = ctrl.(*src.InputControl).Value
			log.Println("[FUNCTION] injectInputs() => key", key)
			node.(*builtins.ModuleInputNode).InputValue = inputs[fmt.Sprintf("%s", key)]
		}
	}
}

func (f *Function[N]) isOutputNode(node src.NodeInterface) bool {
	var _, ok = node.(*builtins.ModuleOutputNode)
	return node.FromModule() && ok
}

func (f *Function[N]) retrieveOutputs(nodes []src.NodeInterface, dfe *engine.DataflowEngine) any {
	var outputNodes []src.NodeInterface

	for _, node := range nodes {
		if f.isOutputNode(node) {
			outputNodes = append(outputNodes, node)
		}
	}

	var outputs map[string]any
	for _, node := range outputNodes {
		var data = dfe.Fetch(node.Node().E.ID)
		if node.HasControl("name") {
			var ctrl, _ = node.Node().Controls.Get("name")
			var key = ctrl.(*src.InputControl).Value

			if key == nil {
				log.Panicln("cannot get output node name")
			}

			outputs[fmt.Sprintf("%s", key)] = data
		}
	}

	return outputs
}
