package fp

import (
	"bytes"
	"encoding/gob"
	"reflect"
)

// IsEmptySlice reports whether the slice is empty or nil.
func IsEmptySlice[T any](slice []T) bool {
	return len(slice) <= 0
}

// Head returns the first element of the slice or the zero value of T if the slice is empty.
func Head[T any](slice []T) T {
	var v T

	if len(slice) > 0 {
		return slice[0]
	}

	return v
}

func Tail[T any](slice []T) []T {

	if len(slice) == 1 {
		return []T{}
	}

	if len(slice) > 1 {
		return slice[1:]
	}

	return slice
}

// Last returns the last element of the slice or the zero value of T if the slice is empty.
func Last[T any](slice []T) T {
	var v T

	if len(slice) > 0 {
		return slice[len(slice)-1]

	}

	return v
}

// Pop returns the last element of the slice, and a slice without the last element.
//
// The third return value is false if the slice is empty.
func Pop[T any](slice []T) (T, []T, bool) {
	var v T

	if len(slice) > 0 {
		return slice[len(slice)-1], slice[:len(slice)-1], true
	}

	return v, []T{}, false
}

// Shift returns the first element of the slice, and a slice without the first element.
//
// The third return value is false if the slice is empty.
func Shift[T any](slice []T) (T, []T, bool) {
	var v T

	if len(slice) > 0 {
		return slice[0], slice[1:], true
	}

	return v, []T{}, false
}

// Map applys fn to each element returns a new slice with the results
func Map[T any, R any](slice []T, fn func(T) R) []R {
	result := make([]R, len(slice))
	for i := range slice {
		result[i] = fn(slice[i])
	}
	return result
}

// ForEach calls fn for each element of slice
func ForEach[T any](slice []T, fn func(T)) {
	for i := range slice {
		fn(slice[i])
	}
}

// FindIndices returns the indices of all elements that match the predicate
func FindIndices[T any](slice []T, predicate func(T) bool) []int {

	var indices []int

	for i := range slice {
		if predicate(slice[i]) {
			indices = append(indices, i)
		}
	}

	return indices
}

// Find returns a new slice with all elements that match the predicate
//
// Identical to [Filter]
func Find[T any](slice []T, predicate func(T) bool) []T {
	var s []T

	for i := range slice {
		if predicate(slice[i]) {
			s = append(s, slice[i])
		}
	}

	return s
}

// FindFirst returns the first element in the slice that matches the predicate.
//
// See also [slices.IndexFunc]
func FindFirst[T any](slice []T, predicate func(T) bool) (T, bool) {
	head := Head(Find(slice, predicate))
	return head, !IsZeroValue(head)
}

// Filter returns a slice with all elements that match the predicate.
//
// Identical to [Find]
func Filter[T any](slice []T, predicate func(T) bool) []T {

	filtered := make([]T, 0, len(slice))

	for i := range slice {
		if predicate(slice[i]) {
			filtered = append(filtered, slice[i])
		}
	}

	return filtered
}

// Replace replaces all elements that match the predicate with the new value, returns a new slice and the number of replacements
func Replace[T any](slice []T, v T, predicate func(T) bool) ([]T, int) {
	cnt := 0

	s := Map(slice, func(x T) T {
		if predicate(x) {
			cnt++
			return v
		} else {
			return x
		}
	})

	return s, cnt
}

// Contains reports whether the slice T contains the value v.
// The values are compared using reflect.DeepEqual.
func Contains[T any](slice []T, v T) bool {

	for i := range slice {
		if reflect.DeepEqual(slice[i], v) {
			return true
		}
	}

	return false
}

// Replace returns a new slice with all distinct elements
func Distinct[T any](slice []T) []T {

	m := map[string]struct{}{}
	distinct := make([]T, 0, len(slice))
	for i := range slice {
		var b bytes.Buffer
		gob.NewEncoder(&b).Encode(slice[i])
		key := b.String()

		if _, exists := m[key]; !exists {
			m[key] = struct{}{}
			distinct = append(distinct, slice[i])
		}
	}

	return distinct
}

func DistinctBy[T any, R any](slice []T, fn func(T) R) []T {

	m := map[string]struct{}{}
	distinct := make([]T, 0, len(slice))
	for i := range slice {
		var b bytes.Buffer
		gob.NewEncoder(&b).Encode(fn(slice[i]))
		key := b.String()

		if _, exists := m[key]; !exists {
			m[key] = struct{}{}
			distinct = append(distinct, slice[i])
		}
	}

	return distinct
}

// Reduce reduces the slice to a single value of type R.
func Reduce[T any, R any](slice []T, initial R, fn func(R, T) R) R {
	total := initial

	for i := range slice {
		total = fn(total, slice[i])
	}

	return total
}

// Flatten flattens a slice of slices.
//
// See [Concat]
func Flatten[T any](slice [][]T) []T {
	return Reduce(slice, []T{}, func(a, b []T) []T {
		return append(a, b...)
	})
}

// FlatMap maps the array, then flattens all the elements into a single slice.
func FlatMap[T any, R any](slice []T, fn func(T) []R) []R {

	result := make([]R, 0, len(slice))

	for i := range slice {
		result = append(result, fn(slice[i])...)
	}

	return result
}

func CopySlice[T any](slice []T) []T {
	target := make([]T, len(slice))
	copy(target, slice)
	return target
}

// Concat concatenates all slices into a single slice.
// If you only concat two slices, use the builtin append function instead:
//
//	slice1 = append(slice1, slice2...)
//
// See [Flatten]
func Concat[T any](slices ...[]T) []T {

	sliceLen := Reduce(slices, 0, func(initial int, slice []T) int {
		return initial + len(slice)
	})

	return Reduce(slices, make([]T, 0, sliceLen), func(a, b []T) []T {

		if IsEmptySlice(b) {
			return a
		}

		return append(a, b...)
	})
}

// GroupBy groups the element of a slice into categories. The categorie is determined by the given function.
func GroupBy[K comparable, T any](slice []T, fn func(T) K) map[K][]T {

	return Reduce(slice, make(map[K][]T), func(initial map[K][]T, value T) map[K][]T {
		key := fn(value)
		mapSlice := initial[key]
		mapSlice = append(mapSlice, value)
		initial[key] = mapSlice
		return initial
	})

}
