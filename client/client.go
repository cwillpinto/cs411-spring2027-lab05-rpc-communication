package client

import (
	"context"
	"time"

	ledger "cs411-lab05-rpc/proto"
)

// ApplyWithRetry performs at most two calls using the same request ID.
// TODO: place a fresh deadline on each attempt; retry only a timeout, and
// return the second response or error without changing the request ID.
func ApplyWithRetry(ctx context.Context, stub ledger.LedgerClient, req *ledger.ApplyRequest, deadline time.Duration) (*ledger.ApplyResponse, error) {
	_ = deadline
	return stub.Apply(ctx, req)
}
