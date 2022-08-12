package main

import (
	"fmt"
	"testing"
)

type mockReaderStub struct {
}

func (r mockReaderStub) ReadFile(string string) ([]byte, error) {
	return []byte("TEST_DATA"), nil
}

func TestReadFile(t *testing.T) {
	var mReader mockReaderStub

	data := GetFile(mReader, "fakeFile")

	fmt.Println(data)
}
