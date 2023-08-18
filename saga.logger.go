package lets

import (
	"context"
	"fmt"

	"github.com/bongnv/saga"
)

type SagaLogger struct {
	StateDescription map[saga.State]string
}

func (sl *SagaLogger) Log(ctx context.Context, tx saga.Transaction) error {
	LogD("State Result: %v", sl.StateDescription[tx.State()])
	fmt.Println()

	return nil
}
