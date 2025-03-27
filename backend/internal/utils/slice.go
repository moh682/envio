package utils

import "math"

type CallbackFunc[T any] func(int, T) bool

func SplitByWidthMake(str string, size int) []string {
	strLength := len(str)
	splitedLength := int(math.Ceil(float64(strLength) / float64(size)))
	splited := make([]string, splitedLength)
	var start, stop int
	for i := 0; i < splitedLength; i += 1 {
		start = i * size
		stop = start + size
		if stop > strLength {
			stop = strLength
		}
		splited[i] = str[start:stop]
	}
	return splited
}

func Filter[T any](list []T, filter CallbackFunc[T]) []T {
	var filtered []T
	for index, item := range list {
		if filter(index, item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

func Find[T any](list []T, callback CallbackFunc[T]) (T, bool) {
	var res T
	hasFoundItem := false
	for index, item := range list {
		if callback(index, item) {
			res = item
			hasFoundItem = true
			break
		}
	}
	return res, hasFoundItem
}

func Reduce[T any, U any](list []T, reducer func(U, T, []T) U, initialValue U) U {
	acc := initialValue
	for _, item := range list {
		acc = reducer(acc, item, list)
	}
	return acc
}
