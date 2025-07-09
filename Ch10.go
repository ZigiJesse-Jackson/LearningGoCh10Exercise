package Add

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add performs addition on two numbers a and b and returns the result which is also a number of the same type. [https://www.mathsisfun.com/numbers/addition.html]
func Add[T Number](a T, b T) T {
	return a + b
}
