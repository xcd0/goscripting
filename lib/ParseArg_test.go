package lib

import (
	"reflect"
	"testing"
)

func TestParseArgs(t *testing.T) {

	patterns := []struct {
		input  string
		expect []string
	}{
		{
			input:  "ls -al",
			expect: []string{"ls", "-al"},
		},
		{
			input:  "ls -a -l",
			expect: []string{"ls", "-a", "-l"},
		},
		{
			input:  "aaa.exe --aaa bbb",
			expect: []string{"aaa.exe", "--aaa", "bbb"},
		},
		{
			input:  "aaa.exe --ccc \"aaa bbb ccc\"",
			expect: []string{"aaa.exe", "--ccc", "\"aaa bbb ccc\""},
		},
		{
			input:  "aaa.exe --ddd 'aaa bbb ccc'",
			expect: []string{"aaa.exe", "--ddd", "'aaa bbb ccc'"},
		},
		{
			input:  "aaa.exe --eee \"aaa 'bbb' ccc\"",
			expect: []string{"aaa.exe", "--eee", "\"aaa 'bbb' ccc\""},
		},
		{
			input:  "aaa.exe --fff 'aaa \"bbb\" ccc'",
			expect: []string{"aaa.exe", "--fff", "'aaa \"bbb\" ccc'"},
		},
	}
	for i, pattern := range patterns {
		actual := ParseArgs(pattern.input)
		if !reflect.DeepEqual(pattern.expect, actual) {
			t.Errorf("ng : pattern %d: want %v, actual %v", i, pattern.expect, actual)
		} else {
			t.Logf("ok : pattern %d: want %v, actual %v", i, pattern.expect, actual)
		}
	}
}
