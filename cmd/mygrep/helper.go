package main

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
