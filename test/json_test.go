package test

import (
	"github.com/ashkan90/auto-core/src"
	"log"
	"testing"
)

func TestJson(t *testing.T) {
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
	//var cf = engine.NewControlFlowEngine(editor)
	//cf.Add(node1)
	//cf.Add(node2)

	editor.Deserialize()

	log.Println(editor)
}
