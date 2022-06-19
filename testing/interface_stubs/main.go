package main

import (
	"log"
	"os"
)

type FileReader interface {
	ReadFile(string string) ([]byte, error)
}

type MyReader struct {
}

func (r MyReader) ReadFile(string string) ([]byte, error) {
	return os.ReadFile(string)
}

func GetFile(reader FileReader, fileName string) []byte {
	file, err := reader.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	return file
}

func main() {
	var reader MyReader
	fileBytes := GetFile(reader, "testFile")
	_ = fileBytes
}
