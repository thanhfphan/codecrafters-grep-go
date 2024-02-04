package main

import (
	"errors"
)

var (
	EOF = errors.New("EOF")
)

type RE struct {
	pattern string
	text    []byte
}

func NewRE(pattern string, text []byte) *RE {
	return &RE{
		pattern: pattern,
		text:    text,
	}

}

func (r *RE) IsMatch() bool {
	if r.matchhere(0, 0) {
		return true
	}

	return false
}

func (r *RE) matchhere(posp, postext int) bool {
	if posp >= len(r.pattern) {
		return true
	}
	if postext >= len(r.text) {
		return false
	}

	// fmt.Printf("posp: %d, postext: %d\n", posp, postext)

	if string(r.pattern[posp]) == "\\" {
		return r.matchhere(posp+1, postext)
	} else if r.pattern[posp] == 'd' && posp >= 1 && string(r.pattern[posp-1]) == "\\" {
		if '0' <= r.text[postext] && r.text[postext] <= '9' {
			return r.matchhere(posp+1, postext+1)
		}

		return false
	} else if r.pattern[posp] == 'w' && posp >= 1 && string(r.pattern[posp-1]) == "\\" {
		if ('0' <= r.text[postext] && r.text[postext] <= '9') ||
			('a' <= r.text[postext] && r.text[postext] <= 'z') ||
			('A' <= r.text[postext] && r.text[postext] <= 'Z') ||
			r.text[postext] == '_' {
			return r.matchhere(posp+1, postext+1)
		}

		return false
	} else {
		if r.pattern[posp] == r.text[postext] {
			// fmt.Println("match char: ", string(r.pattern[posp]), " - ", string(r.text[postext]))
			return r.matchhere(posp+1, postext+1)
		}
	}

	return false
}

func (r *RE) PatternCharAt(pos int) (string, error) {
	if len(r.pattern) > pos {
		return string(r.pattern[pos]), nil
	}

	return "", EOF
}
