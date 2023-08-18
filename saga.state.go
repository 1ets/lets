package lets

import "github.com/bongnv/saga"

type SagaState struct {
	state saga.State
}

func (t *SagaState) State() saga.State {
	return t.state
}
