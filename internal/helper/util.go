package helper

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func IsInRune(a rune, bs []rune) bool {
	for _, v := range bs {
		if a == v {
			return true
		}
	}
	return false
}
