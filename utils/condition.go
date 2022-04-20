package utils

// IFF if yes return a else b
func IFF[T any](yes bool, a, b T) T {
	if yes {
		return a
	}
	return b
}

// IFN if yes return func, a() else b().
func IFN[T any](yes bool, a, b func() T) T {
	if yes {
		return a()
	}
	return b()
}
