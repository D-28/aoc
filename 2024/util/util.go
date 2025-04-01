package util

import "slices"

func Any[T any](slice []T, predicate func(T) bool) bool {
	return slices.ContainsFunc(slice, predicate)
}

func All[T any](slice []T, predicate func(T) bool) bool {
	for _, v := range slice {
		if !predicate(v) {
			return false
		}
	}
	return true
}

func Zip[T, U any](slice1 []T, slice2 []U, callback func(T, U)) {
	length := min(len(slice2), len(slice1))
	for i := range length {
		callback(slice1[i], slice2[i])
	}
}
