package tools_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/bbfh-dev/go-tools/tools/tmap"
	"gotest.tools/assert"
)

func TestSplit(test *testing.T) {
	in := map[string]int{"a": 1, "b": 2, "c": 3}
	keys, values := tmap.Split(in)
	sort.Strings(keys)
	sort.Ints(values)
	assert.DeepEqual(test, keys, []string{"a", "b", "c"})
	assert.DeepEqual(test, values, []int{1, 2, 3})
}

func TestFlatten(test *testing.T) {
	in := map[string]int{"a": 1, "b": 2, "c": 3}
	out := tmap.Flatten(in, func(k string, v int) string { return fmt.Sprintf("%q %v", k, v) })
	sort.Strings(out)
	assert.DeepEqual(test, out, []string{`"a" 1`, `"b" 2`, `"c" 3`})
}
