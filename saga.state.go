package lets

import "github.com/bongnv/saga"

type SagaState struct {
	ProcessState saga.State
}

func (t *SagaState) State() saga.State {
	return t.ProcessState
}
