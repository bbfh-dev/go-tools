package tools_test

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/bbfh-dev/go-tools/tools"
	"gotest.tools/assert"
)

func TestWriteUint8(t *testing.T) {
	var buf bytes.Buffer
	err := tools.WriteUint8(&buf, 69)
	assert.NilError(t, err)

	result, err := tools.ReadUint8(&buf)
	assert.NilError(t, err)
	assert.Equal(t, result, uint8(69))
}

func TestWriteUint16(t *testing.T) {
	var buf bytes.Buffer
	err := tools.WriteUint16(&buf, 6969)
	assert.NilError(t, err)

	result, err := tools.ReadUint16(&buf)
	assert.NilError(t, err)
	assert.Equal(t, result, uint16(6969))
}

func TestWriteUint32(t *testing.T) {
	var buf bytes.Buffer
	err := tools.WriteUint32(&buf, 69696969)
	assert.NilError(t, err)

	result, err := tools.ReadUint32(&buf)
	assert.NilError(t, err)
	assert.Equal(t, result, uint32(69696969))
}

func TestWriteUint64(t *testing.T) {
	var buf bytes.Buffer
	err := tools.WriteUint64(&buf, 6969696969696969)
	assert.NilError(t, err)

	result, err := tools.ReadUint64(&buf)
	assert.NilError(t, err)
	assert.Equal(t, result, uint64(6969696969696969))
}

func TestWriteTuxleString(t *testing.T) {
	var buf bytes.Buffer
	err := tools.WriteTuxleString(&buf, "Hello World!")
	assert.NilError(t, err)

	reader := bufio.NewReader(&buf)
	result, err := tools.ReadTuxleString(*reader)
	assert.NilError(t, err)
	assert.Equal(t, result, "Hello World!\x00")
}