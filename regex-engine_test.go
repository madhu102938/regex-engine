package regexengine

import (
	"testing"
)

func TestMatchRegex_HappyCases(t *testing.T) {
	tests := []struct {
		regex, input string
	}{
		{"a?a?aa", "aa"},
		{"a*", ""},
		{"a*", "aaaa"},
		{"a+", "a"},
		{"a+", "aaaa"},
		{"a?", ""},
		{"a?", "a"},
		{"a|b", "a"},
		{"a|b", "b"},
		{"(a|b)+", "abab"},
		{"a(b|c)d", "abd"},
		{"a(b|c)d", "acd"},
		{"a*b*", "aaabbb"},
		{"(ab)*", "ababab"},
		{"a.b", "acb"},
		{"a.b", "aab"},
		{"(a|b)*c", "abababc"},
		{"a*(b|c)+d*", "abcd"},
		{"a*(b|c)+d*", "acbd"},
		{"a*bc*d*", "abccd"},
		{"ε?ε?εε", "εε"},
	}

	for _, tc := range tests {
		if !MatchRegexWithString(tc.regex, tc.input) {
			t.Errorf("Expected match failed\nRegex: %v\nString: %v\n", tc.regex, tc.input)
		}
	}
}

func TestMatchRegex_SadCases(t *testing.T) {
	tests := []struct {
		regex, input string
	}{
		{"a", ""},
		{"a+", ""},
		{"a?", "aa"},
		{"a|b", "c"},
		{"a.b", "ab"},
		{"a(b|c)d", "abcd"},
		{"(a|b)+", "cc"},
		{"a*(b|c)+d*", "add"},
		{"a*bc*d*", "acccd"},
		{"εε", "ε"},
	}

	for _, tc := range tests {
		if MatchRegexWithString(tc.regex, tc.input) {
			t.Errorf("Unexpected match\nRegex: %v\nString: %v\n", tc.regex, tc.input)
		}
	}
}
