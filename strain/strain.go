package strain

// Keep returns a new slice containing only the elements from list
// for which predicate returns true.
func Keep[T any](list []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, item := range list {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// Keep returns a new slice containing only the elements from list
// for which predicate returns true.
func Discard[T any](list []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, item := range list {
		if !predicate(item) {
			result = append(result, item)
		}
	}
	return result
}
