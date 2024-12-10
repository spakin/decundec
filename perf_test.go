package decslices

import (
	"cmp"
	"iter"
	"math"
	"math/rand"
	"slices"
	"testing"
)

const (
	rngSeed = 0x5350 // Seed for the random-number generator
	numElts = 10000  // Number of elements to sort
)

// BenchmarkSlicesSortInts sorts a slice using slices.Sort to acquire a
// baseline performance number.
func BenchmarkSlicesSortInts(b *testing.B) {
	for range b.N {
		b.StopTimer()
		rng := rand.New(rand.NewSource(rngSeed))
		slice := make([]int, numElts)
		for i := range slice {
			slice[i] = rng.Int()
		}
		b.StartTimer()
		slices.Sort(slice)
	}
}

// BenchmarkSortInts sorts a slice using Sort to measure the overhead of
// decorating and undecorating a slice.
func BenchmarkSortInts(b *testing.B) {
	for range b.N {
		b.StopTimer()
		rng := rand.New(rand.NewSource(rngSeed))
		slice := make([]int, numElts)
		for i := range slice {
			slice[i] = rng.Int()
		}
		b.StartTimer()
		Sort(slice, func(v int) int { return v })
	}
}

// countVowels is a helper function for BenchmarkSlicesSortFuncVowels and
// BenchmarkSortFuncVowels that counts the number of vowels in a string.
func countVowels(s string) int {
	var nv int
	for _, c := range s {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			nv++
		}
	}
	return nv
}

// BenchmarkSlicesSortFuncVowels sorts a slice of strings by number of
// vowels using slices.SortFunc.
func BenchmarkSlicesSortFuncVowels(b *testing.B) {
	for range b.N {
		b.StopTimer()
		rng := rand.New(rand.NewSource(rngSeed))
		slice := make([]string, numElts)
		for i := range slice {
			bs := make([]rune, rng.Intn(50)+50)
			for j := range bs {
				bs[j] = rune(rng.Intn(26) + 'a')
			}
			slice[i] = string(bs)
		}
		b.StartTimer()
		slices.SortFunc(slice,
			func(a, b string) int {
				return cmp.Compare(countVowels(a),
					countVowels(b))
			},
		)
	}
}

// BenchmarkSortFuncVowels sorts a slice of strings by length using
// SortFunc.
func BenchmarkSortFuncVowels(b *testing.B) {
	for range b.N {
		b.StopTimer()
		rng := rand.New(rand.NewSource(rngSeed))
		slice := make([]string, numElts)
		for i := range slice {
			bs := make([]rune, rng.Intn(50)+50)
			for j := range bs {
				bs[j] = rune(rng.Intn(26) + 'a')
			}
			slice[i] = string(bs)
		}
		b.StartTimer()
		SortFunc(slice,
			func(a, b int) int {
				return cmp.Compare(a, b)
			},
			func(s string) int {
				return countVowels(s)
			},
		)
	}
}

// revDigits reverses a uint16's digits.  It accepts larger unsigned types
// but acts as if they are uint16s.
func revDigits[T uint16 | uint32 | uint64](x T) T {
	var rx T
	for range 5 {
		d := x % 10
		rx = rx*10 + d
		x /= 10
	}
	return rx
}

// genReversable yields uint16s that remain uint16s when their digits are
// reversed.
func genReversable(rng *rand.Rand, n int) iter.Seq[uint16] {
	return func(yield func(uint16) bool) {
		for range n {
			x16 := uint32(rng.Intn(math.MaxUint16 + 1))
			rx16 := revDigits(x16)
			if rx16 > math.MaxUint16 {
				continue
			}
			if !yield(uint16(rx16)) {
				return
			}
		}
	}
}

// BenchmarkSlicesSortedFuncDigits sorts numbers by reverse order of their
// digits using slices.SortedFunc.
func BenchmarkSlicesSortedFuncDigits(b *testing.B) {
	for range b.N {
		b.StopTimer()
		rng := rand.New(rand.NewSource(rngSeed))
		seq := genReversable(rng, numElts)
		b.StartTimer()
		_ = slices.SortedFunc(seq,
			func(a, b uint16) int {
				return cmp.Compare(revDigits(a), revDigits(b))
			})
	}
}

// BenchmarkSortedFuncDigits sorts numbers by reverse order of their digits
// using SortedFunc.
func BenchmarkSortedFuncDigits(b *testing.B) {
	for range b.N {
		b.StopTimer()
		rng := rand.New(rand.NewSource(rngSeed))
		seq := genReversable(rng, numElts)
		b.StartTimer()
		_ = SortedFunc(seq,
			func(a, b uint16) int {
				return cmp.Compare(a, b)
			},
			func(x uint16) uint16 {
				return revDigits(x)
			})
	}
}
