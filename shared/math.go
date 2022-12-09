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
