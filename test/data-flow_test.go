package test

import (
	"github.com/ashkan90/auto-core/src"
	"github.com/ashkan90/auto-engine/builtins"
	"github.com/ashkan90/auto-engine/engine"
	"log"
	"testing"
)

func TestDataFlowEngine(t *testing.T) {
	var bus = src.NewEventBus()
	var editor = src.NewNodeEditor(bus)
	var dfe = engine.NewDataflowEngine(editor)

	dfe.StartEmitter()

	node1 := NewTextNode(src.NewNode())
	node2 := builtins.NewLogNode(dfe)
	node3 := builtins.NewLogNode(dfe)
	node4 := builtins.NewLogNode(dfe)
	node5 := NewHttpClientNode(dfe, src.NewNode())
	node6 := NewTextNode(src.NewNode())
	node7 := NewExtensionNode()

	valueInputCtrl := src.NewInputControl(src.InputControlText, &src.InputControlOptions{
		Readonly: ToPtr(false),
		Initial:  ToPtr("hello"),
		Change: func(value any) {
			log.Println("value input ctrl data has been set", value)
		},
	})
	valueInputCtrl2 := src.NewInputControl(src.InputControlText, &src.InputControlOptions{
		Readonly: ToPtr(false),
		Initial:  ToPtr("https://reqres.in/api/users"),
		Change: func(value any) {
			log.Println("value input ctrl data has been set", value)
		},
	})
	node1.AddControl("value", valueInputCtrl)
	node1.AddOutput("value", src.NewOutput[src.Socket](src.NewSocket("value"), "value", true))
	node1.AddOutput("exec", src.NewOutput[src.Socket](src.NewSocket("exec"), "exec", true))

	node2.AddInput("exec", src.NewInput[src.Socket](src.NewSocket("exec"), "exec", true))
	node2.AddInput("message", src.NewInput[src.Socket](src.NewSocket("message"), "Message", true))
	node2.AddOutput("exec", src.NewOutput[src.Socket](src.NewSocket("exec"), "exec", true))

	node3.AddInput("exec", src.NewInput[src.Socket](src.NewSocket("exec"), "exec", true))
	node3.AddInput("message", src.NewInput[src.Socket](src.NewSocket("message"), "Message", true))
	node3.AddOutput("exec", src.NewOutput[src.Socket](src.NewSocket("exec"), "exec", true))

	node4.AddInput("exec", src.NewInput[src.Socket](src.NewSocket("exec"), "exec", true))
	node4.AddInput("message", src.NewInput[src.Socket](src.NewSocket("message"), "Message", true))
	node4.AddOutput("exec", src.NewOutput[src.Socket](src.NewSocket("exec"), "exec", true))

	node5.AddInput("exec", src.NewInput[src.Socket](src.NewSocket("exec"), "exec", true))
	node5.AddInput("url", src.NewInput[src.Socket](src.NewSocket("url"), "Url", false))
	node5.AddOutput("body", src.NewOutput[src.Socket](src.NewSocket("body"), "Body", true))
	node5.AddOutput("exec", src.NewOutput[src.Socket](src.NewSocket("exec"), "exec", true))

	node6.AddControl("value", valueInputCtrl2)
	node6.AddOutput("value", src.NewOutput[src.Socket](src.NewSocket("value"), "value", true))
	node6.AddOutput("exec", src.NewOutput[src.Socket](src.NewSocket("exec"), "exec", true))

	node7.AddInput("exec", src.NewInput[src.Socket](src.NewSocket("exec"), "exec", true))
	node7.AddInput("message", src.NewInput[src.Socket](src.NewSocket("message"), "Message", true))
	node7.AddOutput("exec", src.NewOutput[src.Socket](src.NewSocket("exec"), "exec", true))

	// add nodes into editor
	editor.AddNode(node1)
	editor.AddNode(node2)
	editor.AddNode(node3)
	editor.AddNode(node4)
	editor.AddNode(node5)
	editor.AddNode(node6)
	editor.AddNode(node7)

	// add node connections into editor
	editor.AddConnection(src.NewConnection(node1, "exec", node2, "exec"))
	editor.AddConnection(src.NewConnection(node1, "value", node2, "message"))
	editor.AddConnection(src.NewConnection(node1, "exec", node3, "exec"))
	editor.AddConnection(src.NewConnection(node1, "value", node3, "message"))

	editor.AddConnection(src.NewConnection(node3, "exec", node5, "exec"))

	editor.AddConnection(src.NewConnection(node6, "exec", node5, "exec"))
	editor.AddConnection(src.NewConnection(node6, "value", node5, "url"))

	editor.AddConnection(src.NewConnection(node5, "exec", node4, "exec"))
	editor.AddConnection(src.NewConnection(node5, "body", node4, "message"))

	editor.AddConnection(src.NewConnection(node5, "exec", node7, "exec"))
	editor.AddConnection(src.NewConnection(node5, "body", node7, "message"))

	var cfe = engine.NewControlFlowEngine(editor)

	cfe.Add(node1)
	cfe.Add(node2)
	cfe.Add(node3)
	cfe.Add(node4)
	cfe.Add(node5)
	cfe.Add(node6)
	cfe.Add(node7)

	cfe.Execute(node1.Node().E.ID, "")

	log.Println(dfe)

}

/*

 */
