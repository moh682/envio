package utils

func Map[T, R any](list []T, mapper func(int, T) R) []R {
	var mapped []R
	for index, item := range list {
		mapped = append(mapped, mapper(index, item))
	}
	return mapped
}
