# go-lib

A collection of reusable Go utility libraries.

## Packages

### arrays

Generic array/slice utility functions for Go.

```go
import "github.com/THBE81/go-lib/arrays"

// Map over a slice
result := arrays.Map([]int{1, 2, 3}, func(x int) int { return x * 2 })

// Create typed arrays from variadic arguments
arr := arrays.TypedArray(1, 2, 3)
```

## Installation

```bash
go get github.com/THBE81/go-lib
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
