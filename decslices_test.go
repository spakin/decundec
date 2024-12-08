package decslices_test

import (
	"cmp"
	"testing"

	"github.com/spakin/decslices"
)

// TestSortIntsIdentity uses Sort to sort a list of ints in increasing
// order.  The sort-key function is the identity.
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

// TestSortIntsReversed uses Sort to sort a list of ints in reverse order.
// The sort-key function maps x to -x to achieve this ordering.
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

// TestSortFuncStrings uses SortFunc to sort a list of strings in
// increasing order by length.
func TestSortFuncStrings(t *testing.T) {
	words := []string{
		"interdenominational",
		"unpronounceable",
		"counterrevolutionaries",
		"telecommunication",
		"responsibilities",
		"uncharacteristically",
		"mountaineering",
	}
	expected := []string{
		"mountaineering",
		"unpronounceable",
		"responsibilities",
		"telecommunication",
		"interdenominational",
		"uncharacteristically",
		"counterrevolutionaries",
	}
	decslices.SortFunc(words,
		func(a, b int) int { return cmp.Compare(a, b) },
		func(s string) int { return len(s) })
	for i, w := range words {
		if expected[i] != w {
			t.Fatalf("erroneous value at index %d of %#v", i, words)
		}
	}
}
