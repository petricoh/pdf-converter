package utils

func Concat(origin []string, list []string) []string {
	for _, elem := range list {
		origin = append(origin, elem)
	}
	return origin
}
