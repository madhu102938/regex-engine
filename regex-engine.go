package regexengine

import (
	"github.com/madhu102938/regex-engine/nfa"
	"github.com/madhu102938/regex-engine/utils"
)

func MatchRegexWithString(regexExpression string, stringToMatch string) bool {
	automaton, adj := nfa.BuildNFA(utils.AddConcatenationAndConvertToPostfix(regexExpression))
	return nfa.MatchString(adj, automaton, stringToMatch)
}