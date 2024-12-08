package decslices_test

import (
	"testing"

	"github.com/spakin/decslices"
)

func TestSortIntsIdentity(t *testing.T) {
	array := []int{-2407, -3170, -4255, 2026, 1976, 0, 4596, 7961, -5923, -3758, -7435, -8423, 9733, 4428, 2829, 2977, -2656, -7921, -2350, -1997, -1505, -4593, -7147, 5994, 2184, 6618}
	expected := []int{-8423, -7921, -7435, -7147, -5923, -4593, -4255, -3758, -3170, -2656, -2407, -2350, -1997, -1505, 0, 1976, 2026, 2184, 2829, 2977, 4428, 4596, 5994, 6618, 7961, 9733}
	decslices.Sort(array, func(v int) int { return v })
	for i, v := range array {
		if expected[i] != v {
			t.Fatalf("erroneous value at index %d of %#v", i, array)
		}
	}
}

func TestSortIntsReversed(t *testing.T) {
	array := []int{-2407, -3170, -4255, 2026, 1976, 0, 4596, 7961, -5923, -3758, -7435, -8423, 9733, 4428, 2829, 2977, -2656, -7921, -2350, -1997, -1505, -4593, -7147, 5994, 2184, 6618}
	expected := []int{9733, 7961, 6618, 5994, 4596, 4428, 2977, 2829, 2184, 2026, 1976, 0, -1505, -1997, -2350, -2407, -2656, -3170, -3758, -4255, -4593, -5923, -7147, -7435, -7921, -8423}
	decslices.Sort(array, func(v int) int { return -v })
	for i, v := range array {
		if expected[i] != v {
			t.Fatalf("erroneous value at index %d of %#v", i, array)
		}
	}
}
