package main

import (
	"io"
	"bufio"
	"strings"
)

const (
	ALL_PATTERNS_FOUND = 0
	NO_PATTERNS = 2
	NOT_ALL_PATTERNS_FOUND = 3
)

func OrderedIntersection(data io.Reader, patterns io.Reader, filtered io.Writer) int {
	dataScanner := bufio.NewScanner(data)
	patternsScanner := bufio.NewScanner(patterns)

	var pattern string
	if patternsScanner.Scan() {
		pattern = patternsScanner.Text()
	} else {
		return NO_PATTERNS
	}

	for dataScanner.Scan() {
		scanned := dataScanner.Text()

		if strings.Contains(scanned, pattern) {
			filtered.Write([]byte(scanned + "\n"))

			if patternsScanner.Scan() {
				pattern = patternsScanner.Text()
			} else {
				return ALL_PATTERNS_FOUND
			}
		}
	}
	return NOT_ALL_PATTERNS_FOUND
}
