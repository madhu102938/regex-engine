package utils

import (
	"github.com/madhu102938/regex-engine/constants"
	"github.com/madhu102938/regex-engine/token"
)

const CONCAT_CHARACTER rune = '·'

func convertToToken(expression string) []token.RegexToken {
	tokenSlice := make([]token.RegexToken, 0, len(expression))

	for _, char := range expression {
		switch char {
		case constants.Plus:
			tokenSlice = append(tokenSlice, token.RegexToken{Type: token.Plus})
		case constants.Star:
			tokenSlice = append(tokenSlice, token.RegexToken{Type: token.Star})
		case constants.BracketStart:
			tokenSlice = append(tokenSlice, token.RegexToken{Type: token.BracketStart})
		case constants.BracketEnd:
			tokenSlice = append(tokenSlice, token.RegexToken{Type: token.BracketEnd})
		case constants.Or:
			tokenSlice = append(tokenSlice, token.RegexToken{Type: token.Or})
		case constants.QuestionMark:
			tokenSlice = append(tokenSlice, token.RegexToken{Type: token.QuestionMark})
		case constants.Dot:
			tokenSlice = append(tokenSlice, token.RegexToken{Type: token.Dot})
		default:
			tokenSlice = append(tokenSlice, token.RegexToken{Type: token.Literal, Value: char})
		}
	}

	return tokenSlice
}

func AddConcatenationToExpression(expression string) []token.RegexToken {
	tokenSlice := convertToToken(expression)

	for i := 1; i < len(tokenSlice); i++ {
		if !IsOperator(tokenSlice[i-1].Type) || tokenSlice[i-1].Type == token.BracketEnd || constants.Precedence[tokenSlice[i-1].Type] > constants.Precedence[token.ConcatCharacter] {
			if !IsOperator(tokenSlice[i].Type) || tokenSlice[i].Type == token.BracketStart {
				tokenSlice = append(tokenSlice[:i+1], tokenSlice[i:]...)
				tokenSlice[i] = token.RegexToken{Type: token.ConcatCharacter}
			}
		}
	}
	return tokenSlice
}

func AddConcatenationAndConvertToPostfix(expression string) []token.RegexToken {
	return InfixToPostfix(AddConcatenationToExpression(expression))
}
