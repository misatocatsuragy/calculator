package calc

// Num represent Number in arbitrary numeral system
type Num interface {
	// Binary operator between two Num's
	operateWith(op rune, n Num) (Num, error)
	// String representation of the Num
	String() string
	// Value of the Num
	Value() int
}

// roman numbers
type roman struct {
	rep string
	val int
}

// arabic numbers
type arabic struct {
	rep string
	val int
}

// calculator struct
type Calc struct {
	// symbol of operator
	op rune
	// operands
	x, y Num
}
