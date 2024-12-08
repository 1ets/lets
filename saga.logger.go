package lets

import (
	"context"

	"github.com/bongnv/saga"
)

type SagaLogger struct {
	StateDescription map[saga.State]string
}

func (sl *SagaLogger) Log(ctx context.Context, tx saga.Transaction) error {
	LogD("State Result: %v\n", sl.StateDescription[tx.State()])

	return nil
}
