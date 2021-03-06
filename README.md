# mapping

A mapping library for Go.

![Go](https://github.com/theantichris/mapping/workflows/Go/badge.svg)

## Install

```sh
go get github.com/theantichris/mapping
```

## Use

### ConcurrentMap

Maps an array to another array concurrently using goroutines.

```go
// result: []int{2,3,4}
result, err := ConcurrentMap([]int{1,2,3}, func(num int) int { return num+1 })
```

### Run tests

```sh
go test
```

### Run benchmarks

```sh
go test -bench=.
```
