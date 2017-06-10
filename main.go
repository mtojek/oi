package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s data_file pattterns_file", os.Args[0])
	}

	dataFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Error occurred while opening data_file: %v", err)
	}
	defer dataFile.Close()

	patternsFile, err := os.Open(os.Args[2])
	if err != nil {
		log.Fatalf("Error occurred while opening patterns_file: %v", err)
	}
	defer patternsFile.Close()

	status := OrderedIntersect(dataFile, patternsFile, os.Stdout)
	os.Exit(status)
}