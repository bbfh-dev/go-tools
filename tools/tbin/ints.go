package tbin

import (
	"encoding/binary"
	"io"
)

func WriteUint8(writer io.Writer, in uint8) error {
	_, err := writer.Write([]byte{in})
	return err
}

func WriteUint16(writer io.Writer, in uint16) error {
	var data = make([]byte, 2)
	binary.BigEndian.PutUint16(data, in)
	_, err := writer.Write(data)
	return err
}

func WriteUint32(writer io.Writer, in uint32) error {
	var data = make([]byte, 4)
	binary.BigEndian.PutUint32(data, in)
	_, err := writer.Write(data)
	return err
}

func WriteUint64(writer io.Writer, in uint64) error {
	var data = make([]byte, 8)
	binary.BigEndian.PutUint64(data, in)
	_, err := writer.Write(data)
	return err
}

func Uint8(reader io.Reader) (uint8, error) {
	var data = make([]byte, 1)
	_, err := reader.Read(data)
	return data[0], err
}

func Uint16(reader io.Reader) (uint16, error) {
	var data = make([]byte, 2)
	_, err := reader.Read(data)
	return binary.BigEndian.Uint16(data), err
}

func Uint32(reader io.Reader) (uint32, error) {
	var data = make([]byte, 4)
	_, err := reader.Read(data)
	return binary.BigEndian.Uint32(data), err
}

func Uint64(reader io.Reader) (uint64, error) {
	var data = make([]byte, 8)
	_, err := reader.Read(data)
	return binary.BigEndian.Uint64(data), err
}

func CombineUint(high, low uint32) uint64 {
	return (uint64(high) << 32) | uint64(low)
}

func SeparateUint(combined uint64) (uint32, uint32) {
	high := uint32(combined >> 32)
	low := uint32(combined & 0xFFFFFFFF) // Mask the lower 32 bits
	return high, low
}
