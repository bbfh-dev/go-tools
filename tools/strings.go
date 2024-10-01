package tools

import (
	"fmt"
)

func Log(fn func(msg string, args ...any), message string, args ...any) {
	fn(fmt.Sprintf(message, args...))
}

type formatFunc func(string, string) string

// Format using fmt.Sprintf by providing key and value as strings.
//
// Example format: '%s' %s -> '<key>' <value>
func DefaultFormat(format string) formatFunc {
	return func(key string, value string) string {
		return fmt.Sprintf(format, key, value)
	}
}

// Format map into a slice with specific key+value transformation function.
func FormatMap(in map[string]string, format formatFunc) []string {
	var out = make([]string, len(in))

	var i = 0
	for key, value := range in {
		out[i] = format(key, value)
		i++
	}

	return out
}

// Separates map into its keys and values
func SeparateMap(in map[string]string) (keys []string, values []string) {
	for key, value := range in {
		keys = append(keys, key)
		values = append(values, value)
	}
	return keys, values
}
