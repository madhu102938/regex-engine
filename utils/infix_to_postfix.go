package utils

import (
	"github.com/golang-collections/collections/stack"
)

var precedence = map[rune]int{
	'(':              1,
	')':              1,
	'*':              4,
	'+':              4,
	'?':              4,
	CONCAT_CHARACTER: 3,
	'|':              2,
}

func pushToPostfix(postfix string, newCharacter rune) string {
	return postfix + string(newCharacter)
}

func IsOperator(character rune) bool {
	_, ok := precedence[character]
	return ok
}

func InfixToPostfix(infix string) string {
	s := stack.New()
	postfix := ""

	for _, character := range infix {
		if !IsOperator(character) {
			postfix = pushToPostfix(postfix, character)
		} else if character == '(' || character == ')' {
			if character == '(' {
				s.Push(character)
			} else {
				for (s.Len() > 0) && (s.Peek().(rune) != '(') {
					postfix = pushToPostfix(postfix, s.Pop().(rune))
				}
				s.Pop()
			}
		} else {
			for (s.Len() > 0) && (precedence[s.Peek().(rune)] >= precedence[character]) {
				postfix = pushToPostfix(postfix, s.Pop().(rune))
			}
			s.Push(character)
		}
	}

	for s.Len() > 0 {
		postfix = pushToPostfix(postfix, s.Pop().(rune))
	}

	return postfix
}
