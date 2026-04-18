// Package arrays provides generic array/slice utility functions.
package arrays

// Array creates a new slice from variadic arguments of any type.
func Array(args ...interface{}) []any {
	arr := make([]interface{}, len(args))
	copy(arr, args)
	return arr
}

// TypedArray creates a new typed slice from variadic arguments.
// The type is inferred from the arguments provided.
func TypedArray[T any](args ...T) []T {
	arr := make([]T, len(args))
	copy(arr, args)
	return arr
}

// Map transforms each element of a slice using the provided function,
// returning a new slice of the transformed elements.
func Map[T any, K any](arr []T, fn func(T) K) []K {
	result := make([]K, len(arr))
	for i, v := range arr {
		result[i] = fn(v)
	}
	return result
}
