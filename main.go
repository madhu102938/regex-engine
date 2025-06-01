package main

import (
	"fmt"

	"github.com/madhu102938/regex-engine/regex"
	"github.com/madhu102938/regex-engine/utils"
)

func main() {
	regexExpression := "a*(b|c)+d*"
	stringToMatch := "acd"

	nfa, adj := regex.BuildNFA(utils.AddConcatenationAndConvertToPostfix(regexExpression))

	fmt.Println("regex expression : ", regexExpression)
	fmt.Println("string to be matched : ", stringToMatch)
	fmt.Println("Match: ",regex.MatchString(adj, nfa, stringToMatch))
}
