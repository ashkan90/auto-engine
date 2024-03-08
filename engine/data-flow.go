package engine

import (
	"errors"
	"github.com/ashkan90/auto-core/src"
	"log"
	"sync"
)

// DataflowNodeSetup, bir node'un nasıl işleneceğini tanımlar.
type DataflowNodeSetup interface {
	Inputs() []string
	Outputs() []string
	src.NodeData
}

// Dataflow, node'ları işleme yeteneğini sağlar.
type Dataflow struct {
	editor *src.NodeEditor
	setups map[src.NodeId]DataflowNodeSetup
	mu     sync.RWMutex
}

func NewDataflow(editor *src.NodeEditor) *Dataflow {
	return &Dataflow{
		editor: editor,
		setups: make(map[src.NodeId]DataflowNodeSetup),
	}
}

// Add, dataflow'a bir node ekler.
func (d *Dataflow) Add(node src.NodeInterface, setup DataflowNodeSetup) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	n := node.Node()

	if _, exists := d.setups[n.E.ID]; exists {
		return errors.New("node already processed")
	}
	d.setups[n.E.ID] = setup
	return nil
}

// Remove, dataflow'dan bir node kaldırır.
func (d *Dataflow) Remove(nodeID src.NodeId) {
	d.mu.Lock()
	defer d.mu.Unlock()

	delete(d.setups, nodeID)
}

func (d *Dataflow) FetchInputs(nodeID src.NodeId) map[string]interface{} {
	d.mu.RLock()
	setup, exists := d.setups[nodeID]
	d.mu.RUnlock()

	if !exists {
		log.Panic("[Dataflow.FetchInputs] Node is not initialized")
		return nil
	}

	var inputs = make(map[string]any)
	var connections = d.editor.GetConnectionsTo(nodeID, setup.Inputs())

	for _, conn := range connections {
		sourceData, err := d.Fetch(conn.Source)
		if err != nil {
			log.Panic(err)
		}

		might, ok := inputs[string(conn.TargetInput)]
		if ok {
			log.Println("might boy", might)
		}

		inputs[string(conn.TargetInput)] = sourceData[string(conn.SourceOutput)]
	}

	return inputs
}

func (d *Dataflow) Fetch(nodeID src.NodeId) (map[string]any, error) {
	d.mu.RLock()
	setup, exists := d.setups[nodeID]
	d.mu.RUnlock()

	if !exists {
		return nil, errors.New("[Dataflow.Fetch] node is not initialized")
	}

	outputData := setup.Data(func() map[string]any {
		return d.FetchInputs(nodeID)
	})

	return outputData, nil
}
