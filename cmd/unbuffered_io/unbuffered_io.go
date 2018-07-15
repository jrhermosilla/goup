package main

import (
	"flag"
	"fmt"
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

	n, err := file.Read(buffer)
	if err != nil {
		log.Println(err)
	}

	if n != len(buffer) {
		log.Printf("read only %d bytes of %d", n, len(buffer))
	}

	fmt.Println(string(buffer))
}
