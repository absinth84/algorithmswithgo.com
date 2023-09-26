package module01

import "fmt"

// BaseToBase takes in a number, the base it is currently
// in, and the base you want it to be converted to. It then
// returns a string representing the number in the new base.
//
// Eg:
//
//	BaseToBase("E", 16, 2) => "1110"
func BaseToBase(value string, base, newBase int) string {

	result := ""
	nRes := 0
	multiplier := 1
	remains := 0

	for i := len(value) - 1; i >= 0; i-- {
		var val int
		fmt.Sscanf(string(value[i]), "%X", &val)
		fmt.Printf("Val:%d\n", val)
		nRes += val * multiplier
		multiplier *= base
	}
	fmt.Printf("base10:%d", nRes)
	for nRes > 0 {
		remains = nRes % newBase
		result = fmt.Sprintf("%X", remains) + result
		nRes = nRes / newBase
	}

	return result
}
