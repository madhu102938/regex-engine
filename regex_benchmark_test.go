package regexengine

import (
	"testing"
)

func benchmarkRegexMatch(b *testing.B, length int) {
	// a?^n a^n matching against a^n
	stringToMatch := ""
	regexExpression := ""
	for range length {
		regexExpression += "a?"
		stringToMatch += "a"
	}
	for range length {
		regexExpression += "a"
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MatchRegexWithString(regexExpression, stringToMatch)
	}
}

func BenchmarkMatchLen1(b *testing.B) {
	benchmarkRegexMatch(b, 1)
}

func BenchmarkMatchLen30(b *testing.B) {
	benchmarkRegexMatch(b, 30)
}

func BenchmarkMatchLen60(b *testing.B) {
	benchmarkRegexMatch(b, 60)
}

func BenchmarkMatchLen100(b *testing.B) {
	benchmarkRegexMatch(b, 100)
}
