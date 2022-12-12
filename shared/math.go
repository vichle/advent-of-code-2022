package shared

func IAbs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func ISign(n int) int {
	if n > 0 {
		return 1
	}
	if n == 0 {
		return 0
	}
	return -1
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(integers ...int) int {
	a := integers[0]
	b := integers[1]
	integers = integers[2:]
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
