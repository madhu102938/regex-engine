package regexengine

import (
	"github.com/madhu102938/regex-engine/regex"
	"github.com/madhu102938/regex-engine/utils"
)

func MatchRegexWithString(regexExpression string, stringToMatch string) bool {
	nfa, adj := regex.BuildNFA(utils.AddConcatenationAndConvertToPostfix(regexExpression))
	return regex.MatchString(adj, nfa, stringToMatch)
}