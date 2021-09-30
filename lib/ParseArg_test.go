package lib

import "testing"

func TestParseArgs(t *testing.T) {
	type inout struct {
		input  string
		expect []string
	}

	patterns := []inout{
		{
			input:  "ls -al",
			expect: {"ls", "-al"},
		},
		{
			input:  "ls -a -l",
			expect: {"ls", "-a", "-l"},
		},
	}
	for i, pattern := range patterns {
		actual := ParseArgs(pattern.input)
		if pattern.expected != actual {
			t.Errorf("pattern %d: want %v, actual %v", i, pattern.expected, actual)
		}
	}
}
