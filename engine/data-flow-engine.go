package engine

import (
	"github.com/ashkan90/auto-core/src"
	"log"
)

type DataflowEngine struct {
	Dataflow *Dataflow
	EventBus *src.EventBus
	Cache    *src.Cache
}

func NewDataflowEngine(editor *src.NodeEditor) *DataflowEngine {
	return &DataflowEngine{
		Dataflow: NewDataflow(editor),
		Cache:    src.NewCache(),
		EventBus: editor.GetBus(),
	}
}

func (e *DataflowEngine) SetParent() {}

func (e *DataflowEngine) StartEmitter() {
	e.EventBus.Subscribe("nodeCreated", func(ev src.Event) {
		e.Add(ev.Data.(src.NodeInterface))
	})
	e.EventBus.Subscribe("nodeRemoved", func(ev src.Event) {
		e.Remove(ev.Data.(src.NodeId))
	})
}

func (e *DataflowEngine) GetDataflow() *Dataflow {
	if e.Dataflow == nil {
		log.Panic("data flow is not initiated")
	}

	return e.Dataflow
}

func (e *DataflowEngine) Add(node src.NodeInterface) {
	_, ok := node.(src.NodeData)
	if !ok {
		log.Panic("node is not a type of NodeData")
	}

	n := node.Node()
	inputs := func() []string {
		keys := make([]string, 0, n.Inputs.Len())
		n.Inputs.Range(func(key, _ any) bool {
			keys = append(keys, key.(string))
			return true
		})
		return keys
	}()
	outputs := func() []string {
		keys := make([]string, 0, n.Outputs.Len())
		n.Outputs.Range(func(key, _ any) bool {
			keys = append(keys, key.(string))
			return true
		})
		return keys
	}()

	err := e.GetDataflow().Add(node, NewDataflowNode(inputs, outputs, func(fetchInputs func() map[string]any) map[string]any {
		val, ok := e.Cache.Get(string(n.E.ID))
		if ok {
			return val.(map[string]any)
		}

		cancellable := src.NewCancellable(e.FetchInputs(n.E.ID), node.Data)

		e.Cache.Set(string(n.E.ID), cancellable)

		return cancellable
	}))

	if err != nil {
		log.Println(err)
	}
}

func (e *DataflowEngine) Reset(nodeID src.NodeId) {
	if nodeID != "" {
		var setup = e.GetDataflow().setups[nodeID]
		if setup == nil {
			log.Panic("setup is not initiated")
		}

		e.Cache.Delete(string(nodeID))

		var outputKeys = setup.Outputs()
		var conns = e.Dataflow.editor.GetConnections()

		for _, conn := range conns {
			if conn.Source == nodeID {
				for _, key := range outputKeys {
					if src.NodeId(key) == conn.SourceOutput {
						e.Reset(conn.Target)
					}
				}
			}
		}
	} else {
		e.Cache = e.Cache.Clone()
	}
	e.GetDataflow().Remove(nodeID)
}

func (e *DataflowEngine) Remove(nodeID src.NodeId) {
	e.GetDataflow().Remove(nodeID)
}

func (e *DataflowEngine) FetchInputs(nodeID src.NodeId) map[string]any {
	return e.GetDataflow().FetchInputs(nodeID)
}

func (e *DataflowEngine) Fetch(nodeID src.NodeId) map[string]any {
	data, err := e.GetDataflow().Fetch(nodeID)
	if err != nil {
		log.Println("Fetch", err)
	}
	return data
}
