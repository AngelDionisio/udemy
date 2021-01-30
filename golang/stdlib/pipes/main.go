package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	pr, pw := io.Pipe()

	go func() {
		// having the writer close the channel, the reader will get an EOF, signaling nothing is left to send
		defer pw.Close()
		_, err := fmt.Fprintln(pw, "hello")
		if err != nil {
			panic(err)
		}
	}()

	// io.Copy is ranging, reading  until it gets an EOF
	_, err := io.Copy(os.Stdout, pr)
	if err != nil {
		panic(err)
	}
}
