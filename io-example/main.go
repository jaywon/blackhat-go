package main

import (
	"fmt"
	"log"
	"os"
)

type FooReader struct{}

func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in> ")
	return os.Stdin.Read(b)
}

type FooWriter struct{}

func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out> ")
	return os.Stdout.Write(b)
}

func main() {
	var (
		reader FooReader
		writer FooWriter
	)

	input := make([]byte, 4096)
	s, err := reader.Read(input)

	if err != nil {
		log.Fatalln("Cannot read data from stdin")
	}
	fmt.Printf("Read %d bytes from stdin\n", s)

	s, err = writer.Write(input)

	if err != nil {
		log.Fatalln("Unable to write data to stdout")
	}
	fmt.Printf("Wrote %d bytes to stdout", s)
}
