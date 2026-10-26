package server

import (
	"context"
	"sync"
	"time"

	ledger "cs411-lab05-rpc/proto"
)

type Ledger struct {
	ledger.UnimplementedLedgerServer
	mu              sync.Mutex
	value           int64
	seen            map[string]int64
	DelayAfterApply time.Duration // test hook: delay the reply after committing
}

func NewLedger() *Ledger { return &Ledger{seen: make(map[string]int64)} }

func (s *Ledger) Value() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.value
}

func (s *Ledger) Apply(ctx context.Context, req *ledger.ApplyRequest) (*ledger.ApplyResponse, error) {
	// TODO: validate request ID, operation and amount; return explicit
	// application status. Under the mutex, look up request ID before any
	// mutation, apply each new request exactly once, and remember its result.
	// Release the mutex before DelayAfterApply. A timeout can occur after
	// the mutation but before the response is delivered.
	_ = ctx
	_ = req
	return &ledger.ApplyResponse{Result: s.Value(), Status: "unimplemented"}, nil
}
