### Create coverage profile
```go
go test -coverprofile=coverage.out
```
### Open the coverage profile
```go
go tool cover -html=coverage.out