package module01

import "fmt"

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//	BaseToDec("E", 16) => 14
//	BaseToDec("1110", 2) => 14
func BaseToDec(value string, base int) int {
	const charset = "0123456789ABCDEF"
	result := 0
	multiplier := 1
	lenValue := len(value)
	fmt.Println(value)
	fmt.Println(lenValue)
	for i := len(value) - 1; i >= 0; i-- {
		index := -1
		for j, char := range charset {
			if char == rune(value[i]) {
				index = j
				break
			}
		}
		if index < 0 {
			panic("Someting went wrong")
		}
		result = result + index*multiplier
		fmt.Printf("Parz:%d, Index:%d, multiplier:%d\n", result, index, multiplier)
		multiplier = multiplier * base
	}
	fmt.Printf("Res=%d", result)
	return result
}
