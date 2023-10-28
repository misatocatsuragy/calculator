package calc

import (
	"bufio"
	"fmt"
	"strings"
)

// valid operators
const validOperators = "+-/*"

// consts for indexing
const (
	firstOpPos = iota
	operatorPos
	secondOpPos
)

// amout of tokens
const tokensLength = 3

// valid roman operands (subset of roman numbers)
var validRomanOperands = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

// valid arabic operands (subset of arabic numbers)
var validArabicOperands = map[string]int{
	"1":  1,
	"2":  2,
	"3":  3,
	"4":  4,
	"5":  5,
	"6":  6,
	"7":  7,
	"8":  8,
	"9":  9,
	"10": 10,
}

// parse string and form Calc struct
func (c *Calc) Parse(s string) error {
	// Split string into words
	r := strings.NewReader(s)
	b := bufio.NewScanner(r)
	b.Split(bufio.ScanWords)

	var tokens []string
	for b.Scan() {
		tokens = append(tokens, b.Text())
	}

	err := b.Err()
	if err != nil {
		return fmt.Errorf("error while parsing: %v", err)
	}
	if len(tokens) != tokensLength {
		return fmt.Errorf("not correct expression string")
	}

	// Check operator symbol
	operator := tokens[operatorPos]
	if !isOperator(operator) {
		return fmt.Errorf("not valid operator, expect one of this: + - / *")
	}

	// Operator is string with len = 1, get first element
	c.op = rune(operator[0])

	op1 := tokens[firstOpPos]
	op2 := tokens[secondOpPos]

	// Check that args in same numeral system
	rom := isRomanOp(op1) && isRomanOp(op2)
	arab := isArabicOp(op1) && isArabicOp(op2)

	switch {
	case rom:
		c.x = &roman{op1, validRomanOperands[op1]}
		c.y = &roman{op2, validRomanOperands[op2]}
	case arab:
		c.x = &arabic{op1, validArabicOperands[op1]}
		c.y = &arabic{op2, validArabicOperands[op2]}
	default:
		return fmt.Errorf("both args should be arabic or roman from 1 to 10")
	}

	return nil
}

// Evaluate expression
func (c *Calc) Eval() (Num, error) {
	res, err := c.x.operateWith(c.op, c.y)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Check if operand is valid roman operand
func isRomanOp(s string) bool {
	_, ok := validRomanOperands[s]
	return ok
}

// Check if operand is valid arabic operand
func isArabicOp(s string) bool {
	_, ok := validArabicOperands[s]
	return ok
}

// Check if operator is valid
func isOperator(s string) bool {
	return strings.Contains(validOperators, s)
}
