package main

import (
	"fmt"

	"github.com/madhu102938/regex-engine/utils"
)

func main() {
	regexExpression := "ab+|cd*"
	fmt.Println(utils.AddConcatenationToExpression(regexExpression))
	fmt.Println(utils.AddConcatenationAndConvertToPostfix(regexExpression))


	
}