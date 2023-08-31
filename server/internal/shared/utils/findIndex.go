package utils

func Find[T any](slice []T, f func(T) bool) (int, T) {
	for i, v := range slice {
		if f(v) {
			return i, v
		}
	}

	var defValue T
	return -1, defValue
}