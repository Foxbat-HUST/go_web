package array

func Any[T any](arr []T, f func(T) bool) bool {
	for _, e := range arr {
		if f(e) {
			return true
		}
	}

	return false
}

func All[T any](arr []T, f func(T) bool) bool {
	for _, e := range arr {
		if !f(e) {
			return false
		}
	}

	return true
}

func Filter[T any](arr []T, f func(T) bool) []T {
	result := []T{}
	for _, e := range arr {
		if f(e) {
			result = append(result, e)
		}
	}

	return result
}

func Transform[S, T any](arr []S, f func(S) T) []T {
	result := []T{}
	for _, e := range arr {
		result = append(result, f(e))
	}

	return result
}

func Consume[S, C any](arr []S, f func(S, C) C, init C) C {
	result := init
	for _, e := range arr {
		result = f(e, result)
	}
	return result
}
