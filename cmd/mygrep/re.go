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
	if len(r.pattern) == 0 {
		return true
	}
	if r.pattern[0] == '^' {
		return r.matchhere(1, 0)
	}

	for i := 0; i < len(r.text); i++ {
		if r.matchhere(0, i) {
			return true
		}
	}

	return false
}

func (r *RE) matchhere(posp, postext int) bool {
	if posp >= len(r.pattern) {
		return true
	}

	if string(r.pattern[posp]) == "\\" {
		if postext >= len(r.text) {
			return false
		}
		return r.matchhere(posp+1, postext)
	} else if posp+1 < len(r.pattern) && r.pattern[posp+1] == '+' {
		return r.matchplus(r.pattern[posp], posp+2, postext)
	} else if posp+1 < len(r.pattern) && r.pattern[posp+1] == '?' {
		return r.matchqmark(r.pattern[posp], posp+2, postext)
	} else if r.pattern[posp] == 'd' && posp >= 1 && string(r.pattern[posp-1]) == "\\" {
		if postext >= len(r.text) {
			return false
		}
		if '0' <= r.text[postext] && r.text[postext] <= '9' {
			return r.matchhere(posp+1, postext+1)
		}

		return false
	} else if r.pattern[posp] == 'w' && posp >= 1 && string(r.pattern[posp-1]) == "\\" {
		if postext >= len(r.text) {
			return false
		}
		if ('0' <= r.text[postext] && r.text[postext] <= '9') ||
			('a' <= r.text[postext] && r.text[postext] <= 'z') ||
			('A' <= r.text[postext] && r.text[postext] <= 'Z') ||
			r.text[postext] == '_' {
			return r.matchhere(posp+1, postext+1)
		}

		return false
	} else if r.pattern[posp] == '$' && posp+1 == len(r.pattern) {
		// fmt.Printf("posp: %d, postext: %d\n", posp, postext)
		return postext == len(r.text)
	} else {
		if postext >= len(r.text) {
			return false
		}

		if r.pattern[posp] == r.text[postext] {
			return r.matchhere(posp+1, postext+1)
		}
	}

	return false
}

func (r *RE) matchqmark(c byte, posp, postext int) bool {
	cmatch := 0
	for i := postext; i < len(r.text); i++ {
		if r.text[i] != c {
			break
		}
		cmatch++
	}

	return r.matchhere(posp, postext+cmatch)
}

func (r *RE) matchplus(c byte, posp, postext int) bool {
	cmatch := 0
	for i := postext; i < len(r.text); i++ {
		if r.text[i] != c {
			break
		}
		cmatch++
	}

	if cmatch == 0 {
		return false
	}

	return r.matchhere(posp, postext+cmatch)
}

func (r *RE) PatternCharAt(pos int) (string, error) {
	if len(r.pattern) > pos {
		return string(r.pattern[pos]), nil
	}

	return "", EOF
}
