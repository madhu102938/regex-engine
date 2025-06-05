package utils

import (
	"github.com/golang-collections/collections/stack"

	"github.com/madhu102938/regex-engine/constants"
	"github.com/madhu102938/regex-engine/token"
)

func IsOperator(tokenType token.RegexTokenType) bool {
	_, ok := constants.Precedence[tokenType]
	return ok
}

func InfixToPostfix(infix []token.RegexToken) []token.RegexToken {
	s := stack.New()
	postfix := make([]token.RegexToken, 0, len(infix))

	for _, regexToken := range infix {
		if !IsOperator(regexToken.Type) {
			postfix = append(postfix, regexToken)
		} else if regexToken.Type == token.BracketStart || regexToken.Type == token.BracketEnd {
			if regexToken.Type == token.BracketStart {
				s.Push(regexToken)
			} else {
				for (s.Len() > 0) && (s.Peek().(token.RegexToken).Type != token.BracketEnd) {
					postfix = append(postfix, s.Pop().(token.RegexToken))
				}
				s.Pop()
			}
		} else {
			for (s.Len() > 0) && (constants.Precedence[s.Peek().(token.RegexToken).Type] >= constants.Precedence[regexToken.Type]) {
				postfix = append(postfix, s.Pop().(token.RegexToken))
			}
			s.Push(regexToken)
		}
	}

	for s.Len() > 0 {
		postfix = append(postfix, s.Pop().(token.RegexToken))
	}

	return postfix
}
