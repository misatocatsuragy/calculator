package calc

import (
	"fmt"
)

// for conversion from int to string representing roman numbers
var romanLiterals = []roman{
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

func (r *roman) String() string {
	return r.rep
}

func (r *roman) Value() int {
	return r.val
}

func (r *roman) operateWith(op rune, n Num) (Num, error) {
	var res int

	nr, ok := n.(*roman)
	if !ok {
		panic("Impossible case, both args should be in same num sys")
	}

	switch op {
	case '+':
		res = r.val + nr.val
	case '-':
		res = r.val - nr.val
		if res <= 0 {
			return nil, fmt.Errorf("result can't be negative in roman numeral system")
		}
	case '*':
		res = r.val * nr.val
	case '/':
		res = r.val / nr.val

	}

	return &roman{decimalToRoman(res), res}, nil
}

// conversion from int to string representing roman numbers
func decimalToRoman(n int) string {
	var res string
	for n > 0 {
		for _, l := range romanLiterals {
			if n >= l.val {
				res += l.rep
				n -= l.val
				break
			}
		}
	}
	return res
}
