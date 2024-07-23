## Test

### Create coverage profile
```go
go test -coverprofile=coverage.out
```
### Open the coverage profile
```go
go tool cover -html=coverage.out
```
## Benchmark

### Run benchmark with tests
```go
go test -v -bench=.
```
### Run just the benchmark
```go
go test -bench=. -run=^#
```
### Run just the benchmark with count 10 with 3 seconds in each operation
```go
go test -bench=. -run=^# -count=10 -benchtime=3s
```
### Docs
```go
go help test
```