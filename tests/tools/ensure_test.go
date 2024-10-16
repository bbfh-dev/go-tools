package tools_test

import (
	"fmt"
	"testing"

	"github.com/bbfh-dev/go-tools/tools/ensure"
	"gotest.tools/assert"
)

func TestEnsure(test *testing.T) {
	previous := ensure.OnFail
	test.Cleanup(func() { ensure.OnFail = previous })

	var result string
	ensure.OnFail = func(format string, args ...any) {
		result = fmt.Sprintf(format, args...)
	}

	in := map[string]int{"a": 1, "b": 2, "c": 3}
	ensure.MapContainsKey(in, "d")
	assert.Equal(test, result, fmt.Sprintf(ensure.MAP_CONTAINS_KEY_FMT, "d", "a, b, c"))

	result = ""
	ensure.MapContainsKey(in, "a")
	assert.Equal(test, result, "")

	result = ""
	ensure.NotNil(nil, "Testing this!")
	assert.Equal(test, result, fmt.Sprintf(ensure.NOT_NILL_FMT, "Testing this!"))

	result = ""
	ensure.NotNil(0, "Test")
	assert.Equal(test, result, "")
}
