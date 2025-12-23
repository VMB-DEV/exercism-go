package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

func Keep[T any](list []T, predicate func(T) bool) (result []T) {
	return myFilter(list, predicate, false)
}

func Discard[T any](list []T, predicate func(T) bool) (result []T) {
	return myFilter(list, predicate, true)
}

func myFilter[T any](list []T, predicate func(T) bool, invert bool) (result []T) {
	for _, item := range list {
		if (predicate(item) && !invert) || (!predicate(item) && invert) {
			result = append(result, item)
		}
	}
	return
}
