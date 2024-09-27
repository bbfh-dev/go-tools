package tools_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/bbfh-dev/go-tools/tools"
	"gotest.tools/assert"
)

func TestAnyErr(test *testing.T) {
	var testErr = errors.New("Test")

	assert.Equal(test, tools.AnyErr(nil, nil, nil), nil)
	assert.Equal(test, tools.AnyErr(nil, testErr), testErr)
	assert.Equal(test, tools.AnyErr(testErr), testErr)
	assert.Equal(test, tools.AnyErr(), nil)
}

func TestPrefixErr(test *testing.T) {
	var prefix = "Example"
	var testErr = errors.New("Test")

	assert.Equal(test, tools.PrefixErr(prefix, nil), nil)
	if !strings.HasPrefix(tools.PrefixErr(prefix, testErr).Error(), prefix) {
		test.Fatal("Must have prefix")
	}
}

func TestAssert(test *testing.T) {
	tools.Assert(true, "Should not assert!")
}
