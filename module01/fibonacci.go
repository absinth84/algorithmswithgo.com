package module01

// Fibonacci returns the nth fibonacci number.
//
// A Fibonacci number N is defined as:
//
//	Fibonacci(N) = Fibonacci(N-1) + Fibonacci(N-2)
//
// Where the following base cases are used:
//
//	Fibonacci(0) => 0
//	Fibonacci(1) => 1
//
// Examples:
//
//	Fibonacci(0) => 0
//	Fibonacci(1) => 1
//	Fibonacci(2) => 1
//	Fibonacci(3) => 2
//	Fibonacci(4) => 3
//	Fibonacci(5) => 5
//	Fibonacci(6) => 8
//	Fibonacci(7) => 13
//	Fibonacci(14) => 377
func Fibonacci(n int) int {

	series := []int{0, 1}
	result := 0

	if n == 0 {
		result = 0
	} else if n == 1 {
		result = 1
	} else if n >= 2 {
		for i := 2; i <= n; i++ {
			series = append(series, series[i-1]+series[i-2])
		}
		result = series[n]
	}

	return result
}
