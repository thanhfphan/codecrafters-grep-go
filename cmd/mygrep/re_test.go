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
			pattern: `\d apple`,
			text:    "1 apple",
			expect:  true,
		},
		{
			pattern: `\d apple`,
			text:    "1 orange",
			expect:  false,
		},
		{
			pattern: `\d\d\d apple`,
			text:    "100 apple",
			expect:  true,
		},
		{
			pattern: `\d\d\d apple`,
			text:    "1 orange",
			expect:  false,
		},
		{
			pattern: `\d \w\w\ws`,
			text:    "3 dogs",
			expect:  true,
		},
		{
			pattern: `\d \w\w\ws`,
			text:    "4 cats",
			expect:  true,
		},
		{
			pattern: `\d \w\w\ws`,
			text:    "1 dog",
			expect:  false,
		},
		{
			pattern: `\d apple`,
			text:    "sally has 3 apples",
			expect:  true,
		},
		{
			pattern: "^log",
			text:    "log",
			expect:  true,
		},
		{
			pattern: "^log",
			text:    "dlog",
			expect:  false,
		},
		{
			pattern: "dog$",
			text:    "dog",
			expect:  true,
		},
		{
			pattern: "dog$",
			text:    "dogs",
			expect:  false,
		},
		{
			pattern: "ca+ts",
			text:    "caats",
			expect:  true,
		},
		{
			pattern: "ca+t",
			text:    "caaats",
			expect:  true,
		},
		{
			pattern: "dogs?",
			text:    "dogs",
			expect:  true,
		},
		{
			pattern: "d.g",
			text:    "dog",
			expect:  true,
		},
		{
			pattern: "d.g",
			text:    "cog",
			expect:  false,
		},
		{
			pattern: "(cat|dog)",
			text:    "cat",
			expect:  true,
		},
		{
			pattern: "(cat|dog)",
			text:    "dog",
			expect:  true,
		},
		{
			pattern: "a (cat|dog)",
			text:    "a cat",
			expect:  true,
		},
		{
			pattern: `(cat) and \1`,
			text:    "cat and cat",
			expect:  true,
		},
		{
			pattern: `(cat) and \1`,
			text:    "cat and dog",
			expect:  false,
		},
		{
			pattern: `(\w\w\w\w \d\d\d) is doing \1 times`,
			text:    "grep 101 is doing grep 101 times",
			expect:  true,
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
