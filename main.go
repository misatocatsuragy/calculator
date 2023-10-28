// console-calc reads expression string from stdin
// end eval result
// expression string should have format = op  {+|-|*|/} op
// where op is a number [1..10] in roman or arabic numeral system
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"console-calc/calc"
)

const errorFormatString = "console-calc: %v\n"

func main() {
	log.SetFlags(0)
	// Read string from Stdin
	b := bufio.NewReader(os.Stdin)
	s, err := b.ReadString('\n')
	if err != nil {
		log.Fatalf(errorFormatString, err)
	}

	// Calc variable
	var c calc.Calc

	// Parse string
	if err := c.Parse(s); err != nil {
		log.Fatalf(errorFormatString, err)
	}

	// Eval expression parsed from string
	res, err := c.Eval()
	if err != nil {
		log.Fatalf(errorFormatString, err)
	}

	// Print result
	fmt.Println(res)
}
