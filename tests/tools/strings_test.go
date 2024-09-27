package tools_test

import (
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
