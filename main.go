package main

import (
	"fmt"

	"github.com/madhu102938/regex-engine/regex"
	"github.com/madhu102938/regex-engine/utils"
)

func main() {
	regexExpression := "ab+|cd*"
	fmt.Println(utils.AddConcatenationToExpression(regexExpression))
	fmt.Println(utils.AddConcatenationAndConvertToPostfix(regexExpression))

	nfa, adj := regex.BuildNFA(utils.AddConcatenationAndConvertToPostfix(regexExpression))

	fmt.Println(nfa)

	for from, edges := range adj {
		for _, edge := range edges {
			fmt.Printf("%d --%c--> %d\n", from, edge.Char, edge.To)
		}
	}

}
