package regexengine

import (
	"github.com/madhu102938/regex-engine/nfa"
	"github.com/madhu102938/regex-engine/utils"
)

func MatchRegexWithString(regexExpression string, stringToMatch string) (bool, error) {
	expression, err := utils.AddConcatenationAndConvertToPostfix(regexExpression)
	if err != nil {
		return false, err
	}
	automaton, adj := nfa.BuildNFA(expression)
	return nfa.MatchString(adj, automaton, stringToMatch), nil
}