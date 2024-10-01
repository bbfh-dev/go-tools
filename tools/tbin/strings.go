package tbin

import (
	"bufio"
	"io"
)

// Write a string that ends with '\x00' (null character)
func WriteTuxleString(writer io.Writer, in string) error {
	_, err := writer.Write(append([]byte(in), '\x00'))
	return err
}

// Read a string that ends with '\x00' (null character)
func TuxleString(reader *bufio.Reader) (string, error) {
	return reader.ReadString('\x00')
}
