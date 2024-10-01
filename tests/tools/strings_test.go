package tools_test

import (
	"sort"
	"testing"

	"github.com/bbfh-dev/go-tools/tools"
	"gotest.tools/assert"
)

func TestFormatMap(test *testing.T) {
	var in = map[string]string{
		"keyA": "valueA",
		"keyB": "valueB",
	}
	assert.DeepEqual(
		test,
		tools.FormatMap(in, tools.DefaultFormat("'%s': %s")),
		[]string{`'keyA': valueA`, `'keyB': valueB`},
	)
}

func TestSeparateMap(test *testing.T) {
	var in = map[string]string{
		"keyA": "valueA",
		"keyB": "valueB",
	}

	keys, values := tools.SeparateMap(in)
	sort.Strings(keys)
	sort.Strings(values)
	assert.DeepEqual(test, keys, []string{"keyA", "keyB"})
	assert.DeepEqual(test, values, []string{"valueA", "valueB"})
}
