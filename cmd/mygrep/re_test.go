package main

import (
	"testing"
)

func Test_RE(t *testing.T) {
	tcs := []struct {
		pattern string
		text    string
		expect  bool
	}{
		{
			pattern: "\\d apple",
			text:    "1 apple",
			expect:  true,
		},
		{
			pattern: "\\d apple",
			text:    "1 orange",
			expect:  false,
		},
		{
			pattern: "\\d\\d\\d apple",
			text:    "100 apple",
			expect:  true,
		},
		{
			pattern: "\\d\\d\\d apple",
			text:    "1 orange",
			expect:  false,
		},
		{
			pattern: "\\d \\w\\w\\ws",
			text:    "3 dogs",
			expect:  true,
		},
		{
			pattern: "\\d \\w\\w\\ws",
			text:    "4 cats",
			expect:  true,
		},
		{
			pattern: "\\d \\w\\w\\ws",
			text:    "1 dog",
			expect:  false,
		},
	}

	for _, tc := range tcs {
		re := NewRE(tc.pattern, []byte(tc.text))
		result := re.IsMatch()
		if result != tc.expect {
			t.Errorf("pattern: %s, text: %s, actual: %v, expected: %v", tc.pattern, string(tc.text), result, tc.expect)
		}
	}
}