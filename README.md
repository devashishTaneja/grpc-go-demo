## Compiling the Proto

```shell
protoc proto/*.proto --go_out=pb --go-grpc_out=pb

# Or by using Makefile
make gen-proto
```