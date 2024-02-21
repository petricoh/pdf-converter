package list

func ForEach[T any](slice []T, fn func(T)) {
	for _, elem := range slice {
		fn(elem)
	}
}

func Filter[T any](slice []T, fn func(T) bool) []T {
	var result []T
	for _, elem := range slice {
		ok := fn(elem)
		if ok {
			result = append(result, elem)
		}
	}
	return result
}

func Map[T any, R any](slice []T, fn func(T) R) []R {
	var result []R
	for _, elem := range slice {
		result = append(result, fn(elem))
	}
	return result
}

func Concat(origin []string, list []string) []string {
	for _, elem := range list {
		origin = append(origin, elem)
	}
	return origin
}
