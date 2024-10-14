package tbin

import (
	"bufio"
	"fmt"
	"reflect"

	"github.com/bbfh-dev/go-tools/tools/terr"
)

type Reader struct {
	buffer *bufio.Reader
	step   uint64
	errs   []error
}

// Allows you to sequence reads of a buffer
func NewReader(reader *bufio.Reader) *Reader {
	return &Reader{
		buffer: reader,
		step:   0,
		errs:   []error{},
	}
}

func (reader *Reader) Error() error {
	return terr.Join(reader.errs...)
}

func (reader *Reader) storeErr(value any, err error) any {
	reader.step += 1
	if err != nil {
		reader.errs = append(
			reader.errs,
			fmt.Errorf("Reading %q (step %d): %w", reflect.TypeOf(value).Name(), reader.step, err),
		)
	}
	return value
}

func (reader *Reader) TuxleString(value *string) *Reader {
	*value = reader.storeErr(TuxleString(reader.buffer)).(string)
	return reader
}

func (reader *Reader) Uint8(value *uint8) *Reader {
	*value = reader.storeErr(Uint8(reader.buffer)).(uint8)
	return reader
}

func (reader *Reader) Uint16(value *uint16) *Reader {
	*value = reader.storeErr(Uint16(reader.buffer)).(uint16)
	return reader
}

func (reader *Reader) Uint32(value *uint32) *Reader {
	*value = reader.storeErr(Uint32(reader.buffer)).(uint32)
	return reader
}

func (reader *Reader) Uint64(value *uint64) *Reader {
	*value = reader.storeErr(Uint64(reader.buffer)).(uint64)
	return reader
}

func (reader *Reader) Int8(value *int8) *Reader {
	*value = reader.storeErr(Int8(reader.buffer)).(int8)
	return reader
}

func (reader *Reader) Int16(value *int16) *Reader {
	*value = reader.storeErr(Int16(reader.buffer)).(int16)
	return reader
}

func (reader *Reader) Int32(value *int32) *Reader {
	*value = reader.storeErr(Int32(reader.buffer)).(int32)
	return reader
}

func (reader *Reader) Int64(value *int64) *Reader {
	*value = reader.storeErr(Int64(reader.buffer)).(int64)
	return reader
}
