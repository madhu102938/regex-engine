package utils

func AddConcatenationToExpression(expression string) string {
	runesExpression := make([]rune, 0, len(expression))

	for _, char := range expression {
		runesExpression = append(runesExpression, char)
	}

	for i := 1; i < len(runesExpression); i++ {
		if !IsOperator(runesExpression[i-1]) || runesExpression[i-1] == ')' || precedence[runesExpression[i-1]] > precedence['.'] {
			if !IsOperator(runesExpression[i]) || runesExpression[i] == '(' {
				runesExpression = append(runesExpression[:i+1], runesExpression[i:]...)
				runesExpression[i] = '.'
			}
		}
	}

	returnExpression := ""
	for _, character := range runesExpression {
		returnExpression += string(character)
	}

	return returnExpression
}

func AddConcatenationAndConvertToPostfix(expression string) string {
	return InfixToPostfix(AddConcatenationToExpression(expression))
}
