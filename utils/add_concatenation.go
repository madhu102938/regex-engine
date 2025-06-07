package utils

import (
	"fmt"

	"github.com/madhu102938/regex-engine/constants"
	"github.com/madhu102938/regex-engine/token"
)

func convertToToken(expression string) ([]token.RegexToken, error) {
	n := 0
	expressionSlice := make([]rune, 0, len(expression))
	for _, char := range expression {
		expressionSlice = append(expressionSlice, char)
		n++
	}
	tokenSlice := make([]token.RegexToken, 0, n)

	for i := 0; i < n; i++ {
		switch expressionSlice[i] {
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
		case constants.BackSlash:
			i++
			if i < n {
				tokenSlice = append(tokenSlice, token.RegexToken{Type: token.Literal, Value: expressionSlice[i]})
			} else {
				return nil, fmt.Errorf("tokenizer: pattern invalid (ending with a single backslash)")
			}
		default:
			tokenSlice = append(tokenSlice, token.RegexToken{Type: token.Literal, Value: expressionSlice[i]})
		}
	}

	return tokenSlice, nil
}

func AddConcatenationToExpression(expression string) ([]token.RegexToken, error) {
	tokenSlice, err := convertToToken(expression)
	if err != nil {
		return nil, err
	}

	for i := 1; i < len(tokenSlice); i++ {
		if !IsOperator(tokenSlice[i-1].Type) || tokenSlice[i-1].Type == token.BracketEnd || constants.Precedence[tokenSlice[i-1].Type] > constants.Precedence[token.ConcatCharacter] {
			if !IsOperator(tokenSlice[i].Type) || tokenSlice[i].Type == token.BracketStart {
				tokenSlice = append(tokenSlice[:i+1], tokenSlice[i:]...)
				tokenSlice[i] = token.RegexToken{Type: token.ConcatCharacter}
			}
		}
	}
	return tokenSlice, nil
}

func AddConcatenationAndConvertToPostfix(expression string) ([]token.RegexToken, error) {
	tokenSlice, err := AddConcatenationToExpression(expression)
	if err != nil {
		return nil, err
	}
	return InfixToPostfix(tokenSlice), nil
}
