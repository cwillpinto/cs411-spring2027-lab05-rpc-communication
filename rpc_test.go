package rpc_test

import (
	"context"
	"net"
	"testing"
	"time"

	"cs411-lab05-rpc/client"
	ledger "cs411-lab05-rpc/proto"
	"cs411-lab05-rpc/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func setup(t *testing.T) (*server.Ledger, ledger.LedgerClient) {
	t.Helper()
	lis := bufconn.Listen(1024 * 1024)
	s := grpc.NewServer()
	store := server.NewLedger()
	ledger.RegisterLedgerServer(s, store)
	go func() { _ = s.Serve(lis) }()
	t.Cleanup(s.Stop)
	conn, err := grpc.NewClient("passthrough:///bufnet", grpc.WithInsecure(), grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) { return lis.Dial() }))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = conn.Close() })
	return store, ledger.NewLedgerClient(conn)
}

func TestApplyAndDuplicate(t *testing.T) {
	store, stub := setup(t)
	req := &ledger.ApplyRequest{RequestId: "request-1", Operation: "add", Amount: 7}
	first, err := stub.Apply(context.Background(), req)
	if err != nil || first.GetStatus() != "applied" || first.GetResult() != 7 {
		t.Fatalf("first response: %v, %v", first, err)
	}
	second, err := stub.Apply(context.Background(), req)
	if err != nil || second.GetStatus() != "duplicate" || second.GetResult() != 7 || store.Value() != 7 {
		t.Fatalf("duplicate response: %v, %v, value=%d", second, err, store.Value())
	}
}

func TestDeadlineThenRetry(t *testing.T) {
	store, stub := setup(t)
	store.DelayAfterApply = 80 * time.Millisecond
	req := &ledger.ApplyRequest{RequestId: "request-2", Operation: "add", Amount: 5}
	response, err := client.ApplyWithRetry(context.Background(), stub, req, 20*time.Millisecond)
	if err != nil || response.GetResult() != 5 || store.Value() != 5 {
		t.Fatalf("retry response: %v, %v, value=%d", response, err, store.Value())
	}
}

func TestInvalidOperation(t *testing.T) {
	_, stub := setup(t)
	response, err := stub.Apply(context.Background(), &ledger.ApplyRequest{RequestId: "bad", Operation: "multiply", Amount: 4})
	if err != nil || response.GetStatus() != "invalid" {
		t.Fatalf("invalid response: %v, %v", response, err)
	}
}
