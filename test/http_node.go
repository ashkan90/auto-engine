package test

import (
	"github.com/ashkan90/auto-core/src"
	"github.com/ashkan90/auto-core/utils"
	"github.com/ashkan90/auto-engine/engine"
	"log"
)

type HttpClientNode struct {
	Base           *src.Node[src.NodeBase]
	DataflowEngine *engine.DataflowEngine
	Local          *utils.SyncMap
}

func (n *HttpClientNode) Data(inputs func() map[string]any) map[string]any {
	var m = make(map[string]any)
	n.Local.Range(func(key, value any) bool {
		m[key.(string)] = value
		return true
	})
	return m
}

func (n *HttpClientNode) Node() *src.Node[src.NodeBase] {
	return n.Base
}

func (n *HttpClientNode) HasInput(k string) bool {
	return n.Base.HasInput(k)
}

func (n *HttpClientNode) AddInput(k string, input src.InputInterface) {
	n.Base.AddInput(k, input)
}

func (n *HttpClientNode) RemoveInput(k string) {
	n.Base.RemoveInput(k)
}

func (n *HttpClientNode) HasOutput(k string) bool {
	return n.HasOutput(k)
}

func (n *HttpClientNode) AddOutput(k string, output src.PortInterface) {
	n.Base.AddOutput(k, output)
}

func (n *HttpClientNode) RemoveOutput(k string) {
	n.Base.RemoveOutput(k)
}

func (n *HttpClientNode) HasControl(k string) bool {
	return n.Base.HasControl(k)
}

func (n *HttpClientNode) AddControl(k string, control src.ControlInterface) {
	n.Base.AddControl(k, control)
}

func (n *HttpClientNode) RemoveControl(k string) {
	n.Base.RemoveControl(k)
}

func (n *HttpClientNode) Execute(input string, forward func(output string)) {
	var inputs = n.DataflowEngine.FetchInputs(n.Base.E.ID)
	var url = inputs["url"].(*string)

	log.Println("httpclient...", *url)

	n.Local.Add("body", ToPtr("{name: emirhan}"))

	forward("exec")
}

//func NewHttpClientNode(dataflow *engine.DataflowEngine, base *src.Node[src.NodeBase]) src.NodeInterface {
//	return &HttpClientNode{
//		Base:           base,
//		DataflowEngine: dataflow,
//		Local:          utils.NewSyncMap(),
//	}
//}
