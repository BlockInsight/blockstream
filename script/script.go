package script

// WorkflowBuilder .
type WorkflowBuilder interface {
	Load(data []byte) (Workflow, error)
}

// Workflow .
type Workflow interface{}
