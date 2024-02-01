package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// Usage: echo <input_text> | your_grep.sh -E <pattern>
func main() {
	if len(os.Args) < 3 || os.Args[1] != "-E" {
		fmt.Fprintf(os.Stderr, "usage: mygrep -E <pattern>\n")
		os.Exit(2) // 1 means no lines were selected, >1 means error
	}

	pattern := os.Args[2]

	line, err := io.ReadAll(os.Stdin) // assume we're only dealing with a single line
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: read input text: %v\n", err)
		os.Exit(2)
	}

	ok, err := matchLine(line, pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}

	if !ok {
		os.Exit(1)
	}

	// default exit code is 0 which means success
}

func matchLine(line []byte, pattern string) (bool, error) {
	// if utf8.RuneCountInString(pattern) != 1 {
	// 	return false, fmt.Errorf("unsupported pattern: %q", pattern)
	// }

	var ok bool
	switch pattern {
	case "\\d":
		ok = isHasDigit(line)
	case "\\w":
		ok = isHasDigit(line) || isHasCharacter(line) || containChar(line, '_')
	default:
		if strings.HasPrefix(pattern, "[") && strings.HasSuffix(pattern, "]") {
			newPattern := strings.TrimPrefix(pattern, "[")
			newPattern = strings.TrimSuffix(pattern, "]")
			for _, c := range []byte(newPattern) {
				if has := containChar(line, c); has {
					return true, nil
				}
			}

			return false, nil
		} else {
			ok = bytes.ContainsAny(line, pattern)
		}

	}

	return ok, nil
}

func containChar(line []byte, char byte) bool {
	for _, l := range line {
		if l == char {
			return true
		}
	}

	return false
}

func isHasCharacter(line []byte) bool {
	for _, l := range line {
		if 'a' <= l && l <= 'z' {
			return true
		}

		if 'A' <= l && l <= 'Z' {
			return true
		}
	}

	return false
}

func isHasDigit(line []byte) bool {
	for _, c := range line {
		if '0' <= c && c <= '9' {
			return true
		}
	}
	return false
}
