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

func TestReader(test *testing.T) {
	var buffer bytes.Buffer

	err := tbin.WriteTuxleString(&buffer, "Hello World!")
	assert.NilError(test, err)

	err = tbin.WriteUint8(&buffer, 69)
	assert.NilError(test, err)

	err = tbin.WriteUint16(&buffer, 420)
	assert.NilError(test, err)

	err = tbin.WriteInt32(&buffer, 69)
	assert.NilError(test, err)

	err = tbin.WriteInt64(&buffer, 420)
	assert.NilError(test, err)

	var reader = tbin.NewReader(bufio.NewReader(&buffer))
	var a string
	var b uint8
	var c uint16
	var d int32
	var e int64

	assert.NilError(test, reader.TuxleString(&a).Uint8(&b).Uint16(&c).Int32(&d).Int64(&e).Error())
	assert.DeepEqual(test, a, "Hello World!")
	assert.Equal(test, b, uint8(69))
	assert.Equal(test, c, uint16(420))
	assert.Equal(test, d, int32(69))
	assert.Equal(test, e, int64(420))
	assert.Equal(test, len(buffer.Bytes()), 0)
}

func TestCombineInt(test *testing.T) {
	assert.DeepEqual(
		test,
		tbin.CombineUint(uint32(0x12345678), uint32(0x9ABCDEF0)),
		uint64(0x123456789ABCDEF0),
	)
	assert.DeepEqual(test, tbin.CombineUint(69, 420), uint64(296352743844))
}

func TestSeparateInt(test *testing.T) {
	h, l := tbin.SeparateUint(296352743844)
	assert.Equal(test, h, uint32(69))
	assert.Equal(test, l, uint32(420))

	h, l = tbin.SeparateUint(0x123456789ABCDEF0)
	assert.Equal(test, h, uint32(0x12345678))
	assert.Equal(test, l, uint32(0x9ABCDEF0))
}
