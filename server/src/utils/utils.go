package utils

func Ternary(a any, b any, result bool) any {
	if result {
		return a
	}
	return b
}
