package main

import (
	"fmt"
	"testing"
)

type mockReader struct {
}

func (r mockReader) ReadFile(string string) ([]byte, error) {
	return []byte("TEST_DATA"), nil
}

func TestReadFile(t *testing.T) {
	var mReader mockReader

	data := GetFile(mReader, "fakeFile")

	fmt.Println(data)
}
