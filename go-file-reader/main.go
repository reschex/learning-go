package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// get file pointer from program argument
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	//since *File implements reader, use io.Copy again
	io.Copy(os.Stdout, f)

	// close file or log error if closing returned error
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}
