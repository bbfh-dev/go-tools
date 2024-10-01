package tools_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/bbfh-dev/go-tools/tools/terr"
	"gotest.tools/assert"
)

func TestAnyErr(test *testing.T) {
	var testErr = errors.New("Test")

	assert.Equal(test, terr.Join(nil, nil, nil), nil)
	assert.ErrorContains(test, terr.Join(nil, testErr), "Errors occured")
	assert.ErrorContains(test, terr.Join(testErr), "Errors occured")
	assert.Equal(test, terr.Join(), nil)
}

func TestPrefixErr(test *testing.T) {
	var prefix = "Example"
	var testErr = errors.New("Test")

	assert.Equal(test, terr.Prefix(prefix, nil), nil)
	if !strings.HasPrefix(terr.Prefix(prefix, testErr).Error(), prefix) {
		test.Fatal("Must have prefix")
	}
}

func TestAssert(test *testing.T) {
	terr.Assert(true, "Should not assert!")
}
