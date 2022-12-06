package utils

import (
	"golang.org/x/exp/maps"
)

// Union returns the union of the set of elements of the two provided slices
// Note that ordering is not preserved, and as this is a set, duplicates are removed.
func Union[T comparable](s1 []T, s2 []T) []T {
	union := make(map[T]struct{}, 0)
	for _, t := range s1 {
		union[t] = struct{}{}
	}
	for _, t := range s2 {
		union[t] = struct{}{}
	}

	return maps.Keys(union)
}

// Intersection returns the intersection of the set of elements of the two provided slices
// Note that ordering is not preserved, and as this is a set, duplicates are removed
func Intersection[T comparable](s1 []T, s2 []T) []T {
	intersection := make(map[T]struct{})
	for _, t1 := range s1 {
		for i, t2 := range s2 {
			if t1 == t2 {
				// remove t2 from s2 so that we don't have to keep iterating over it
				s2 = append(s2[:i], s2[i+1:]...)
				intersection[t1] = struct{}{}
				break
			}
		}
	}

	return maps.Keys(intersection)
}

func SetDifference[T comparable](s1 []T, s2 []T) []T {
	difference := make(map[T]struct{})
	for _, t1 := range s1 {
		difference[t1] = struct{}{}
	}
	for _, t2 := range s2 {
		if _, ok := difference[t2]; !ok {
			delete(difference, t2)
		}
	}
	return maps.Keys(difference)
}
