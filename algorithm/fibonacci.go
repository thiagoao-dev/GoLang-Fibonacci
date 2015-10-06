package algorithm

// fibonacci is a function that returns
// a function that returns an int.
func Fibonacci() func() int {
	sum, before, next := 0, 0, 1
	return func() int {
		sum    = before
		before = next
		next   = before + sum
		return sum
	}
}