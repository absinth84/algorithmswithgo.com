package module01

import "fmt"

// Reverse will return the provided word in reverse
// order. Eg:
//
//	Reverse("cat") => "tac"
//	Reverse("alphabet") => "tebahpla"
func Reverse(word string) string {

	var rev string

	for i := len(word) - 1; i >= 0; i-- {

		rev = rev + string(word[i])
	}
	fmt.Println(rev)

	return rev
}
