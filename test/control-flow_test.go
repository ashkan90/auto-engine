package test

import (
	"github.com/ashkan90/auto-core/src"
	"github.com/ashkan90/auto-engine/engine"
	"log"
	"testing"
)

func TestNewControlFlow(t *testing.T) {
	var editor = src.NewNodeEditor(src.NewEventBus())
	//var cf = NewControlFlow(editor)

	node1, _ := editor.AddNode(NewLogNode(nil))
	node2, _ := editor.AddNode(src.NewNode())

	editor.AddConnection(src.NewConnection(node1, node1.Node().E.ID, node2, node2.Node().E.ID))

	var conns = editor.GetConnections()
	for _, conn := range conns {
		log.Printf("connection source output: %s / target input: %s", conn.SourceOutput, conn.TargetInput)
	}

	log.Println(node1, node2)
}

func TestNewControlFlow2(t *testing.T) {
	var editor = src.NewNodeEditor(src.NewEventBus())

	node1, _ := editor.AddNode(src.NewNode())
	node2, _ := editor.AddNode(src.NewNode())

	node1.AddInput("name", src.NewPort[src.Socket](src.Socket{
		Name: "Socket Name",
	}, "Name", false))
	node1.AddInput("surname", src.NewPort[src.Socket](src.Socket{
		Name: "Socket Name",
	}, "Surname", false))
	node1.AddOutput("name", src.NewPort[src.Socket](src.Socket{
		Name: "Socket Name",
	}, "name", false))
	node1.AddOutput("exec", src.NewPort[src.Socket](src.Socket{
		Name: "Socket Name",
	}, "exec", false))

	node2.AddInput("name", src.NewPort[src.Socket](src.Socket{
		Name: "Socket Name",
	}, "Name", false))
	node2.AddInput("exec", src.NewPort[src.Socket](src.Socket{
		Name: "Socket Name",
	}, "exec", false))
	node2.AddOutput("exec", src.NewPort[src.Socket](src.Socket{
		Name: "Socket Name",
	}, "exec", false))

	editor.AddConnection(src.NewConnection(node1, "exec", node2, "exec"))
	editor.AddConnection(src.NewConnection(node1, "name", node2, "name"))

	var conns = editor.GetConnections()
	for _, conn := range conns {
		log.Printf("connection source output: %s / target input: %s", conn.SourceOutput, conn.TargetInput)
	}
	var cf = engine.NewControlFlowEngine(editor)
	cf.Add(node1)
	cf.Add(node2)

	cf.Execute(node1.Node().E.ID, "123")
}
