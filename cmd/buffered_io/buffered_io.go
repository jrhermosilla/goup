package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// current working directory
	if cwd, err := os.Getwd(); err == nil {
		fmt.Printf("cwd is: %s\n", cwd)
		fmt.Println("-------------------------")
	}

	filename := flag.String("filename", "", "")
	flag.Parse()

	file, err := os.OpenFile(*filename, os.O_RDONLY, 755)
	if err != nil {
		log.Print(err)
	}

	defer file.Close()

	buffer := make([]byte, 16)
	reader := bufio.NewReader(file)

	// reading bytes
	fmt.Println("buffered io: reading bytes")
	n, err := reader.Read(buffer)
	if err != nil {
		log.Println(err)
	}

	if n != len(buffer) {
		log.Printf("read only %d bytes of %d", n, len(buffer))
	}

	fmt.Println(string(buffer))
	fmt.Println("-------------------------")

	// reading all
	fmt.Println("buffered io: reading all")

	// needed because was read before
	file.Seek(0, io.SeekStart)
	reader.Reset(file)

	content, err := ioutil.ReadAll(reader)
	fmt.Println(string(content))
	fmt.Println("-------------------------")
}
