package module01

func GCD(a, b int) int {
	//result := 0
	//divider := b / 2

	// if a%b == 0 {
	// 	return b
	// }

	// for i := b / 2; i >= 1; i-- {

	// 	if a%divider == 0 && b%divider == 0 {
	// 		result = divider
	// 		break
	// 	}
	// 	divider--

	// }

	for b >= 0 {
		if b == 0 {
			return a
		}
		a, b = b, a%b
	}
	return 0
}
