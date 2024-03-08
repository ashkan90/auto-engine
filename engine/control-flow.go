package engine

import (
	"fmt"
	"github.com/ashkan90/auto-core/src"
	"log"
	"sync"
)

type ControlFlowNodeSetup interface {
	Inputs() []string
	Outputs() []string
	src.NodeExecutor
}

type ControlFlow struct {
	Setups map[src.NodeId]ControlFlowNodeSetup
	Editor *src.NodeEditor
	mu     sync.Mutex
}

func NewControlFlow(editor *src.NodeEditor) *ControlFlow {
	return &ControlFlow{
		Setups: make(map[src.NodeId]ControlFlowNodeSetup),
		Editor: editor,
	}
}

func (cf *ControlFlow) Add(node src.NodeInterface, setup ControlFlowNodeSetup) {
	cf.mu.Lock()
	defer cf.mu.Unlock()

	n := node.Node()
	if _, exists := cf.Setups[n.E.ID]; exists {
		fmt.Println("Node already processed")
		return
	}
	cf.Setups[n.E.ID] = setup
}

func (cf *ControlFlow) Remove(nodeID src.NodeId) {
	cf.mu.Lock()
	defer cf.mu.Unlock()

	delete(cf.Setups, nodeID)
}

func (cf *ControlFlow) Execute(nodeID src.NodeId, input string) {
	cf.mu.Lock()
	setup, exists := cf.Setups[nodeID]
	cf.mu.Unlock()

	if !exists {
		log.Panic("[ControlFlow.Execute] Node is not initialized")
	}

	setup.Execute(input, func(output string) {
		var outputs = setup.Outputs()
		var found = false
		for _, o := range outputs {
			if o == output {
				found = !found
			}

			if found {
				break
			}
		}

		if !found {
			log.Panicf("outputs don't have a key. expected: '%s', in: '%v'", output, outputs)
		}

		var cons []*src.Connection[src.ConnectionBase]
		for _, con := range cf.Editor.GetConnections() {
			if con.Source == nodeID && con.SourceOutput == src.NodeId(output) {
				cons = append(cons, con)
			}
		}

		for _, con := range cons {
			log.Println(con.Target, con.TargetInput)
			cf.Execute(con.Target, string(con.TargetInput))
		}
	})
}
