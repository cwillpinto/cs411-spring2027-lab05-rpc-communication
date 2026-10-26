# Lab 5: RPC Communication

**CS 411: Computing Paradigms - Spring 2027**  
**Week 6: February 15-21, 2027**  
**Due:** February 21, 2027  
**Value:** 100 points

**Starter repository:** [Open the Spring 2027 starter](https://github.com/cwillpinto/cs411-spring2027-lab05-rpc-communication)

## Purpose

Implement a small gRPC service whose network boundary remains visible. Trace one request, inject a delayed response, and use request identifiers to make duplicate retries observable and safe.

## Learning objectives

- Design a Protocol Buffers request and response
- Separate transport failure from application failure
- Implement an idempotent retry path

## Before you begin

1. Review the Week 6 slide deck and assigned VoiceThread or reading.
2. Create your own repository from the Spring 2027 template linked in Blackboard. It includes `proto/service.proto`, generated Go bindings, public tests, and TODOs in the client and server packages.
3. Run `go test ./...` before editing and save the failing output in `Lab5_Terminal.txt`. The generated client stub and server handler are connected through an in-memory gRPC transport in the tests.

## Part A Define the contract (25 points)

Make the wire boundary explicit.

- Inspect the supplied `proto/service.proto`. In `Lab5_Analysis.md` or PDF, explain why `request_id`, `operation`, `amount`, `result`, and `status` are on the wire and identify which fields a retry must preserve.
- Regenerate the Go client and server bindings using the command in the repository README. Confirm that `go test ./...` still compiles; commit the generated files in your repository.
- Draw `Lab5_Sequence.pdf` or `.png` for one normal request and one delayed response. Label client call, generated stub, transport, server handler, state change, deadline, retry, and reply.

## Part B Implement the RPC (45 points)

Preserve failure semantics.

- Complete `client.ApplyWithRetry`: give each attempt its own deadline, retry only a deadline failure, use the *same* request ID, and make no more than two attempts.
- Complete `server.Ledger.Apply`: validate the operation and positive amount, return explicit `applied`, `duplicate`, or `invalid` application status, and preserve the server's cumulative value.
- Under the mutex, check the request ID before any mutation and remember the result of a newly applied request. A duplicate must return the recorded result without applying another side effect.

## Part C Inject failure (30 points)

Test uncertainty rather than only success.

- Use the supplied delayed-response test, which commits the first operation, delays its reply past the client deadline, then retries the same request ID. Do not hold the ledger mutex while delaying the response.
- Run `go test ./...` and `go test -race ./...`; save complete commands and output in `Lab5_Terminal.txt`.
- In the analysis, state what the client knows and does not know immediately after the timeout. Explain why one observed side effect, rather than two, supports the idempotency claim.

## Verification

- `go test ./...` and `go test -race ./...` pass.
- A successful request crosses the generated stub and handler and returns `applied`.
- A duplicate returns `duplicate` with the original result; the delayed retry produces one side effect.

## What to submit

Submit your own repository URL and:

1. Your own repository URL (created from the Spring 2027 template).
2. `proto/service.proto`, generated Go bindings, `client/client.go`, and `server/server.go` in that repository.
3. `Lab5_Sequence.pdf` or `Lab5_Sequence.png`.
4. `Lab5_Analysis.pdf` or `Lab5_Analysis.md`.
5. `Lab5_Terminal.txt`.

## Rubric

| Criterion | Points |
| --- | ---: |
| Protocol contract | 20 |
| Request response trace | 15 |
| Client and server implementation | 30 |
| Deadline and retry behavior | 25 |
| Failure-semantics explanation | 10 |
| **Total** | **100** |

## AI usage and academic integrity

Add an **AI Usage** section to the required analysis file. Name each tool and list the prompts or interactions that affected the work. If you did not use AI, write `AI Usage: None.` You must be able to explain and defend every submitted design and line of code.

## Late work

Late work may be accepted for up to three calendar days with a 25 percent deduction per day, following the course syllabus.
