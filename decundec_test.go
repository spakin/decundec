package decundec_test

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"testing"

	"github.com/spakin/decundec"
)

// TestSortIntsIdentity uses Sort to sort a list of ints in increasing
// order.  The sort-key function is the identity.
func TestSortIntsIdentity(t *testing.T) {
	array := []int{-2407, -3170, -4255, 2026, 1976, 0, 4596, 7961, -5923, -3758, -7435, -8423, 9733, 4428, 2829, 2977, -2656, -7921, -2350, -1997, -1505, -4593, -7147, 5994, 2184, 6618}
	expected := []int{-8423, -7921, -7435, -7147, -5923, -4593, -4255, -3758, -3170, -2656, -2407, -2350, -1997, -1505, 0, 1976, 2026, 2184, 2829, 2977, 4428, 4596, 5994, 6618, 7961, 9733}
	decundec.Sort(array, func(v int) int { return v })
	for i, v := range array {
		if expected[i] != v {
			t.Fatalf("erroneous value at index %d of %v", i, array)
		}
	}
}

// TestSortIntsReversed uses Sort to sort a list of ints in reverse order.
// The sort-key function maps x to -x to achieve this ordering.
func TestSortIntsReversed(t *testing.T) {
	array := []int{-2407, -3170, -4255, 2026, 1976, 0, 4596, 7961, -5923, -3758, -7435, -8423, 9733, 4428, 2829, 2977, -2656, -7921, -2350, -1997, -1505, -4593, -7147, 5994, 2184, 6618}
	expected := []int{9733, 7961, 6618, 5994, 4596, 4428, 2977, 2829, 2184, 2026, 1976, 0, -1505, -1997, -2350, -2407, -2656, -3170, -3758, -4255, -4593, -5923, -7147, -7435, -7921, -8423}
	decundec.Sort(array, func(v int) int { return -v })
	for i, v := range array {
		if expected[i] != v {
			t.Fatalf("erroneous value at index %d of %v", i, array)
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
	decundec.SortFunc(words,
		cmp.Compare,
		func(s string) int { return len(s) })
	for i, w := range words {
		if expected[i] != w {
			t.Fatalf("erroneous value at index %d of %v", i, words)
		}
	}
}

// TestSortStableFuncStrings uses SortStableFunc to sort a list of strings in
// increasing order by length.
func TestSortStableFuncStrings(t *testing.T) {
	words := []string{
		"acknowledgements",
		"oversimplification",
		"misinterpretation",
		"mispronunciations",
		"electroencephalographs",
		"comprehensiveness",
		"multimillionaires",
		"misunderstandings",
		"ophthalmologists",
		"hyperventilating",
		"conceptualizations",
	}
	expected := []string{
		"acknowledgements",
		"ophthalmologists",
		"hyperventilating",
		"misinterpretation",
		"mispronunciations",
		"comprehensiveness",
		"multimillionaires",
		"misunderstandings",
		"oversimplification",
		"conceptualizations",
		"electroencephalographs",
	}
	decundec.SortStableFunc(words,
		cmp.Compare,
		func(s string) int { return len(s) })
	for i, w := range words {
		if expected[i] != w {
			t.Fatalf("erroneous value at index %d of %v", i, words)
		}
	}
}

// TestSortedUint16s uses Sorted to sort a sequence of uint16s based on their
// base-10 digits specified in reverse order.
func TestSortedUint16s(t *testing.T) {
	// The following values fit in a uint16 both as is and with digits
	// reversed.
	seq := slices.Values([]uint16{46792, 25213, 27803, 26265, 33681, 28782, 13034, 64363, 6915, 40721, 33774, 56093, 20411, 56350, 9644, 425, 6693, 62111, 65533, 39440, 17622, 24273, 12475, 52161, 63284})
	revDigits := func(x uint16) uint16 {
		// Reverse a uint16's digits.
		var rx uint16
		for range 5 {
			d := x % 10
			rx = rx*10 + d
			x /= 10
		}
		return rx
	}
	expected := []uint16{39440, 56350, 62111, 20411, 40721, 52161, 33681, 17622, 28782, 46792, 27803, 25213, 65533, 64363, 24273, 56093, 6693, 13034, 9644, 33774, 63284, 6915, 425, 26265, 12475}
	array := decundec.Sorted(seq, revDigits)
	for i, v := range array {
		if expected[i] != v {
			t.Fatalf("erroneous value at index %d of %v", i, array)
		}
	}
}

// TestSortedFuncStrings uses SortedFunc to sort a sequence of string
// representations of numbers.
func TestSortedFuncStrings(t *testing.T) {
	// Define a finite sequence of numbers, represented with strings.
	seq := func(yield func(string) bool) {
		var x uint = 17
		for {
			if !yield(fmt.Sprintf("%04d", x)) || x == 9 {
				return
			}
			x = (x * 7) % 23
		}
	}
	expected := make([]string, 22)
	for i := range expected {
		expected[i] = fmt.Sprintf("%04d", i+1)
	}

	// Define a function that extract numbers from strings.
	num2str := func(s string) int {
		v, err := strconv.Atoi(s)
		if err != nil {
			t.Fatal(err)
		}
		return v
	}

	// Sort and validate the sequence.
	array := decundec.SortedFunc(seq, cmp.Compare, num2str)
	for i, v := range array {
		if expected[i] != v {
			t.Fatalf("erroneous value at index %d of %v", i, array)
		}
	}
}

// TestSortedStableFuncStrings uses SortedFunc to sort a sequence of string
// representations of numbers, each prefixed by a letter.
func TestSortedStableFuncStrings(t *testing.T) {
	// Define a finite sequence of numbers, represented with strings.
	seq := func(yield func(string) bool) {
		var c rune = 'C'
		var x uint = 17
		for {
			if !yield(fmt.Sprintf("%c%d", c, x)) || c == 'A' && x == 9 {
				return
			}
			x = (x * 7) % 23
			if x == 17 {
				c--
			}
		}
	}
	expected := make([]string, 0, 22*3)
	for x := 1; x <= 22; x++ {
		for _, c := range []rune{'C', 'B', 'A'} {
			expected = append(expected, fmt.Sprintf("%c%d", c, x))
		}
	}

	// Define a function that extract numbers from strings.
	num2str := func(s string) int {
		v, err := strconv.Atoi(s[1:])
		if err != nil {
			t.Fatal(err)
		}
		return v
	}

	// Sort and validate the sequence.
	array := decundec.SortedStableFunc(seq, cmp.Compare, num2str)
	for i, v := range array {
		if expected[i] != v {
			t.Fatalf("erroneous value at index %d of %v", i, array)
		}
	}
}
