package tools_test

import (
	"sort"
	"testing"

	"github.com/bbfh-dev/go-tools/tools/tslice"
	"gotest.tools/assert"
)

func sortedDiff(a, b []int) []int {
	diff := tslice.Diff(a, b)
	sort.Ints(diff)
	return diff
}

func TestSlicesDiff(test *testing.T) {
	var a = []int{1, 2, 3, 4}
	var b = []int{1, 2, 3, 4, 5}
	assert.DeepEqual(test, sortedDiff(a, b), []int{5})

	var c = []int{2, 3, 4, 5}
	assert.DeepEqual(test, sortedDiff(c, a), []int{1, 5})
	assert.DeepEqual(test, sortedDiff(a, c), []int{1, 5})
}
