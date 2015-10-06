package algorithm

// fibonacci is a function that returns
// a function that returns an int.
func Fibonacci() func(int) int {
	sum := 0
	return func (x int) int {
     sum += x
     return sum
	}
}