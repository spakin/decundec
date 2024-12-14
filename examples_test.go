package decundec

import (
	"cmp"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Sort a list of integers by the sum of their base-10 digits.
func ExampleSort() {
	numbers := []uint{
		87138, 33405, 12313, 42691, 33687, 56336, 53757, 89532, 75941, 80650,
	}
	sumDigits := func(n uint) uint {
		var s uint
		for ; n > 0; n /= 10 {
			s += n % 10
		}
		return s
	}
	Sort(numbers, sumDigits)
	fmt.Println(numbers)
	// Output: [12313 33405 80650 42691 56336 75941 87138 33687 53757 89532]
}

// Given a list of strings of the form "⟨name⟩_⟨age⟩", sort the list by
// decreasing age.  Preserve the original order of people who are the same
// age.
func ExampleSortStableFunc() {
	people := []string{
		"James_18",
		"Mary_15",
		"Michael_17",
		"Patricia_16",
		"Robert_15",
		"Jennifer_16",
		"John_17",
		"Linda_18",
	}
	SortStableFunc(people,
		func(a, b uint8) int {
			return cmp.Compare(b, a)
		},
		func(s string) uint8 {
			toks := strings.Split(s, "_")
			if len(toks) != 2 {
				panic("invalid name_age string")
			}
			age, err := strconv.ParseUint(toks[1], 10, 8)
			if err != nil {
				panic("invalid name_age string")
			}
			return uint8(age)
		})
	for _, s := range people {
		fmt.Println(s)
	}
	// Output:
	// James_18
	// Linda_18
	// Michael_17
	// John_17
	// Patricia_16
	// Jennifer_16
	// Mary_15
	// Robert_15
}

// Sort a sequence of float64 values in increasing order of their cosine.
func ExampleSorted() {
	seq := func(yield func(x float64) bool) {
		v := 0.0
		for range 25 {
			if !yield(v) {
				return
			}
			v += 0.75
		}
	}
	sorted := Sorted(seq,
		func(x float64) float64 { return math.Cos(x) })
	fmt.Println(sorted)
	// Output: [15.75 3 9.75 9 3.75 15 16.5 2.25 10.5 8.25 4.5 14.25 17.25 1.5 11.25 7.5 5.25 13.5 18 0.75 12 6.75 6 12.75 0]
}
