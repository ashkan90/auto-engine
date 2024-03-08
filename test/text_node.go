package test

import (
	"github.com/ashkan90/auto-core/src"
	"log"
)

type TextNode struct {
	Base *src.Node[src.NodeBase]
}

func (n *TextNode) Data(inputs func() map[string]any) map[string]any {
	var m = make(map[string]any)
	n.Base.Controls.Range(func(key, value any) bool {
		m[key.(string)] = value.(src.ControlInterface).GetValue()
		return true
	})
	log.Println("[TextNode.Data]", m)
	return m
}

func (n *TextNode) Node() *src.Node[src.NodeBase] {
	return n.Base
}

func (n *TextNode) HasInput(k string) bool {
	return n.Base.HasInput(k)
}

func (n *TextNode) AddInput(k string, input src.InputInterface) {
	n.Base.AddInput(k, input)
}

func (n *TextNode) RemoveInput(k string) {
	n.Base.RemoveInput(k)
}

func (n *TextNode) HasOutput(k string) bool {
	return n.HasOutput(k)
}

func (n *TextNode) AddOutput(k string, output src.PortInterface) {
	n.Base.AddOutput(k, output)
}

func (n *TextNode) RemoveOutput(k string) {
	n.Base.RemoveOutput(k)
}

func (n *TextNode) HasControl(k string) bool {
	return n.Base.HasControl(k)
}

func (n *TextNode) AddControl(k string, control src.ControlInterface) {
	n.Base.AddControl(k, control)
}

func (n *TextNode) RemoveControl(k string) {
	n.Base.RemoveControl(k)
}

func (n *TextNode) Execute(input string, forward func(output string)) {
	forward("exec")
}

func NewTextNode(base *src.Node[src.NodeBase]) src.NodeInterface {
	return &TextNode{
		Base: base,
	}
}
