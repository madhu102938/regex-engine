package token

type RegexTokenType int

const (
	Plus RegexTokenType = iota
	Star
	QuestionMark
	BracketStart
	Or
	BracketEnd
	Literal
	Dot
	BackSlash
	ConcatCharacter
	Epsilon
)

type RegexToken struct {
	Type RegexTokenType
	Value rune
}

func (r RegexToken) Match (character rune) bool {
	switch r.Type {
	case Literal:
		return r.Value == character
	case Dot:
		return true
	default:
		return false
	}
}