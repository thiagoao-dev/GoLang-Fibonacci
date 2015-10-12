package algorithm

// fibonacci is a function that returns
// a function that returns an int.
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
