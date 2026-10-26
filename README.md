# CS411 Spring 2027 · Lab 5: RPC Communication

Starter template: https://github.com/cwillpinto/cs411-spring2027-lab05-rpc-communication

Read the included handout. The Protocol Buffers contract and generated Go bindings are supplied so the starter builds. Explain the fields and wire trace in your analysis, then complete the TODOs in `server/server.go` and `client/client.go`.

```sh
go test ./...
go test -race ./...
```

The public test uses an in-memory gRPC connection; it still crosses the generated client stub and server handler. To regenerate bindings after a protocol change, run `protoc --go_out=. --go_opt=module=cs411-lab05-rpc --go-grpc_out=. --go-grpc_opt=module=cs411-lab05-rpc proto/service.proto` from the repository root.

## Work in your own GitHub repository

The instructor starter is an upstream source. Keep your completed work in a repository under your own GitHub account, then submit the repository URL and the evidence listed in the assignment handout through Blackboard.

If Blackboard links a template repository, choose **Use this template**, create the new repository under your account, and clone that new repository. Check `git remote -v` before pushing; `origin` must be your repository.

For a direct clone, first create an empty repository under your account. Replace the placeholders below with the starter link supplied in Blackboard and your own empty repository URL:

```sh
git clone https://github.com/cwillpinto/cs411-spring2027-lab05-rpc-communication
cd <cloned-directory>
git remote rename origin upstream
git remote add origin <your-empty-repository-url>
git remote -v
git push -u origin HEAD
```

Commit and push your completed work to your own `origin`. Do not push to the instructor's upstream repository. Use the exact submission files and verification commands in the assignment handout.
