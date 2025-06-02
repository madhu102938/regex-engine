package regex_test

import (
	"testing"

	"github.com/madhu102938/regex-engine/regex"
	"github.com/madhu102938/regex-engine/utils"
)

func benchmarkRegexMatch(b *testing.B, length int) {
	// a?^n a^n matching again a^n
	stringToMatch := ""
	regexExpression := ""
	for range length {
		regexExpression += "a?"
		stringToMatch += "a"
	}
	for range length {
		regexExpression += "a"
	}

	postfix := utils.AddConcatenationAndConvertToPostfix(regexExpression)
	nfa, adj := regex.BuildNFA(postfix)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		regex.MatchString(adj, nfa, stringToMatch)
	}
}

func BenchmarkMatchLen1(b *testing.B) {
	benchmarkRegexMatch(b, 1)
}

func BenchmarkMatchLen60(b *testing.B) {
	benchmarkRegexMatch(b, 60)
}

func BenchmarkMatchLen100(b *testing.B) {
	benchmarkRegexMatch(b, 100)
}

func BenchmarkMatchLen200(b *testing.B) {
	benchmarkRegexMatch(b, 200)
}
