package calc

import "strconv"

func (a *arabic) String() string {
	return a.rep
}

func (a *arabic) Value() int {
	return a.val
}

func (a *arabic) operateWith(op rune, n Num) (Num, error) {
	var res int

	na, ok := n.(*arabic)
	if !ok {
		panic("Impossible case, both args should be in same num sys")
	}

	switch op {
	case '+':
		res = a.val + na.val
	case '-':
		res = a.val - na.val
	case '*':
		res = a.val * na.val
	case '/':
		res = a.val / na.val

	}

	return &arabic{strconv.Itoa(res), res}, nil
}
