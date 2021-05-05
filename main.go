package main

import (
	"os"
	"flag"
	"log"
)

func main() {
	var pArg = flag.String("p", "patterns", "patterns file")
	flag.Parse()

	patternsFile, err := os.Open(*pArg)
	if err != nil {
		log.Fatalf("Error occurred while opening patterns_file: %v", err)
	}
	defer patternsFile.Close()

	var dataFile = os.Stdin
	if len(os.Args) == 4 {
		dataFile, err = os.Open(flag.Arg(0))
		if err != nil {
			log.Fatalf("Error occurred while opening data_file: %v", err)
		}
		defer dataFile.Close()
	}

	status := OrderedIntersection(dataFile, patternsFile, os.Stdout)
	os.Exit(status)
}
