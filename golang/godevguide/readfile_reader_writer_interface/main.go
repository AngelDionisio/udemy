package main

import (
	"fmt"
	"io"
	"os"
)

// to run, 'go run main.go myfile.txt' on console
func main() {
	// slice of strings representing inputs from terminal
	args := os.Args

	f, err := os.Open(args[1])
	if err != nil {
		fmt.Println("Error trying to open file", err)
		os.Exit(1)
	}
	defer f.Close()

	// we can use the io.Copy which takes a writer, and some content to read that implements reader
	// the os.Open returns a *File which implements the reader interface. So we use copy to write
	// the contents on the writer, in this case os.Stdout (os standard out),
	// which then writes the content of the text file to the console
	// io.Copy(os.Stdout, f)

	fow := fileOpenWriter{}
	io.Copy(fow, f)
}

type fileOpenWriter struct{}

// Write implements the Writer interface, meant to return the # of bytes written
// and an error. Can we used with functions that accept the writer interface
func (fileOpenWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Wrote this many bytes:", len(bs))

	return len(bs), nil
}
