package tools

import (
	"bufio"
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

// Write a string that ends with '\x00' (null character)
func WriteTuxleString(writer io.Writer, in string) error {
	_, err := writer.Write(append([]byte(in), '\x00'))
	return err
}

func ReadUint8(reader io.Reader) (uint8, error) {
	var data = make([]byte, 1)
	_, err := reader.Read(data)
	return data[0], err
}

func ReadUint16(reader io.Reader) (uint16, error) {
	var data = make([]byte, 2)
	_, err := reader.Read(data)
	return binary.BigEndian.Uint16(data), err
}

func ReadUint32(reader io.Reader) (uint32, error) {
	var data = make([]byte, 4)
	_, err := reader.Read(data)
	return binary.BigEndian.Uint32(data), err
}

func ReadUint64(reader io.Reader) (uint64, error) {
	var data = make([]byte, 8)
	_, err := reader.Read(data)
	return binary.BigEndian.Uint64(data), err
}

// Read a string that ends with '\x00' (null character)
func ReadTuxleString(reader bufio.Reader) (string, error) {
	return reader.ReadString('\x00')
}
