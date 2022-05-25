package mm

import x "golang.org/x/exp/constraints"

// Return the larger value
func Max[T x.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Return the smaller value
func Min[T x.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}
