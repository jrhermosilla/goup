package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	directory := flag.String("directory", ".", "directory to list")
	flag.Parse()

	// current working directory
	if cwd, err := os.Getwd(); err == nil {
		fmt.Printf("cwd is: %s\n", cwd)
		fmt.Println("-------------------------")
	}

	fmt.Println("Directory entries:")
	files, err := ioutil.ReadDir(*directory)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
