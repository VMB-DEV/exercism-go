package strain

func Discard[T any](list []T, predicate func(T) bool) (result []T) {
	return Keep(list, func(item T) bool { return !predicate(item) })
}

func Keep[T any](list []T, predicate func(T) bool) (result []T) {
	for _, item := range list {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return
}
