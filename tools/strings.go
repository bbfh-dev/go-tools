package tools

import (
	"fmt"
)

func Log(fn func(msg string, args ...any), message string, args ...any) {
	fn(fmt.Sprintf(message, args...))
}
