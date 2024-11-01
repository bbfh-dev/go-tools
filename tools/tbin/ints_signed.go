package tbin

import "io"

// Has no tests because it doesn't need them
// It's simply a wrapper for tbin/ints.go

func WriteInt8(writer io.Writer, in int8) error {
	return WriteUint8(writer, uint8(in))
}

func WriteInt16(writer io.Writer, in int16) error {
	return WriteUint16(writer, uint16(in))
}

func WriteInt32(writer io.Writer, in int32) error {
	return WriteUint32(writer, uint32(in))
}

func WriteInt64(writer io.Writer, in int64) error {
	return WriteUint64(writer, uint64(in))
}

func Int8(reader io.Reader) (int8, error) {
	out, err := Uint8(reader)
	return int8(out), err
}

func Int16(reader io.Reader) (int16, error) {
	out, err := Uint16(reader)
	return int16(out), err
}

func Int32(reader io.Reader) (int32, error) {
	out, err := Uint32(reader)
	return int32(out), err
}

func Int64(reader io.Reader) (int64, error) {
	out, err := Uint64(reader)
	return int64(out), err
}

func CombineInt(high, low int32) int64 {
	return int64(CombineUint(uint32(high), uint32(low)))
}

func SeparateInt(combined int64) (int32, int32) {
	high, low := SeparateUint(uint64(combined))
	return int32(high), int32(low)
}
