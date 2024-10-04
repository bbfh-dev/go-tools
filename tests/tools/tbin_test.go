package tools_test

import (
	"bufio"
	"bytes"
	"io"
	"testing"

	"github.com/bbfh-dev/go-tools/tools/tbin"
	"gotest.tools/assert"
)

func TestWriteUint8(t *testing.T) {
	var buf bytes.Buffer
	err := tbin.WriteUint8(&buf, 69)
	assert.NilError(t, err)

	result, err := tbin.Uint8(&buf)
	assert.NilError(t, err)
	assert.Equal(t, result, uint8(69))
	_, err = buf.ReadByte()
	assert.ErrorType(t, err, io.EOF)
}

func TestWriteUint16(t *testing.T) {
	var buf bytes.Buffer
	err := tbin.WriteUint16(&buf, 6969)
	assert.NilError(t, err)

	result, err := tbin.Uint16(&buf)
	assert.NilError(t, err)
	assert.Equal(t, result, uint16(6969))
	_, err = buf.ReadByte()
	assert.ErrorType(t, err, io.EOF)
}

func TestWriteUint32(t *testing.T) {
	var buf bytes.Buffer
	err := tbin.WriteUint32(&buf, 69696969)
	assert.NilError(t, err)

	result, err := tbin.Uint32(&buf)
	assert.NilError(t, err)
	assert.Equal(t, result, uint32(69696969))
	_, err = buf.ReadByte()
	assert.ErrorType(t, err, io.EOF)
}

func TestWriteUint64(t *testing.T) {
	var buf bytes.Buffer
	err := tbin.WriteUint64(&buf, 6969696969696969)
	assert.NilError(t, err)

	result, err := tbin.Uint64(&buf)
	assert.NilError(t, err)
	assert.Equal(t, result, uint64(6969696969696969))
	_, err = buf.ReadByte()
	assert.ErrorType(t, err, io.EOF)
}

func TestWriteTuxleString(t *testing.T) {
	var buf bytes.Buffer
	err := tbin.WriteTuxleString(&buf, "Hello World!")
	assert.NilError(t, err)

	reader := bufio.NewReader(&buf)
	result, err := tbin.TuxleString(reader)
	assert.NilError(t, err)
	assert.Equal(t, result, "Hello World!")
	_, err = reader.ReadByte()
	assert.ErrorType(t, err, io.EOF)
}
