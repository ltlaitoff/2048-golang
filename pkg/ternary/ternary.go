package ternary

import (
// "fmt"
// "log/slog"
)

func Ternary[T any](condition bool, a T, b T) T {
	// slog.Debug(fmt.Sprintf("condition = %t, a = %t, b = %t", condition, a, b))

	if condition {
		return a
	}

	return b
}
