# example-uber-zap

```go
$ go run main.go

{"level":"info","ts":1746716227.2062423,"caller":"example-uber-zap/main.go:22","msg":"Application started"}
{"level":"info","ts":1746716227.2062423,"caller":"example-uber-zap/main.go:29","msg":"Starting data processing"}
{"level":"info","ts":1746716229.2075026,"caller":"example-uber-zap/main.go:31","msg":"Data processing in progress"}
{"level":"info","ts":1746716231.208436,"caller":"example-uber-zap/main.go:33","msg":"Data processing finished"}
{"level":"info","ts":1746716233.2087674,"caller":"example-uber-zap/main.go:39","msg":"Cleaning up resources"}
{"level":"info","ts":1746716235.2095482,"caller":"example-uber-zap/main.go:41","msg":"Closing database connections"}
{"level":"info","ts":1746716237.2105865,"caller":"example-uber-zap/main.go:43","msg":"Resources cleaned up"}
{"level":"info","ts":1746716239.2116847,"caller":"example-uber-zap/main.go:25","msg":"Application finished"}
```

```go
go tool pprof heap.out
```