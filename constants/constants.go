package constants

import (
	"github.com/madhu102938/regex-engine/token"
)

const (
	Plus rune = '+'
	Star rune = '*'
	QuestionMark rune = '?'
	BracketStart rune = '('
	Or rune = '|'
	BracketEnd rune = ')'
	Dot rune = '.'
	BackSlash rune = '\\'
)

var Precedence = map[token.RegexTokenType]int{
	token.BracketStart:    1,
	token.BracketEnd:      1,
	token.Star:            4,
	token.Plus:            4,
	token.QuestionMark:    4,
	token.ConcatCharacter: 3,
	token.Or:              2,
}