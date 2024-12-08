/*
   Package decslices wraps the sorting functions in Go's slices
   package to employ a decorate-sort-undecorate idiom.  This can
   improve sorting performance when the comparison function is
   time-consuming to compute.
*/

package decslices

import (
	"cmp"
	"iter"
	"slices"
)

// An assoc associates a value of any type with a value of any other type.
// The latter is used as a sort key.
type assoc[E, Ekey any] struct {
	e    E
	eKey Ekey
}

// Sort sorts a slice in ascending order given a function that maps each
// element to a sort key.
func Sort[S ~[]E, E any, Ekey cmp.Ordered](x S, key func(E) Ekey) {
	// Decorate each element of x.
	type Pair assoc[E, Ekey]
	pairs := make([]Pair, len(x))
	for i, e := range x {
		pairs[i].e = e
		pairs[i].eKey = key(x[i])
	}

	// Sort the decorated slice
	slices.SortStableFunc(pairs, func(a, b Pair) int {
		return cmp.Compare(a.eKey, b.eKey)
	})

	// Undecorate each element of pairs back into x.
	for i, p := range pairs {
		x[i] = p.e
	}
}

// SortFunc sorts a slice in ascending order as determined by the cmp
// function and given a function that maps each element to a sort key.
func SortFunc[S ~[]E, E, Ekey any](x S, cmp func(a, b Ekey) int, key func(E) Ekey) {
	// Decorate each element of x.
	type Pair assoc[E, Ekey]
	pairs := make([]Pair, len(x))
	for i, e := range x {
		pairs[i].e = e
		pairs[i].eKey = key(x[i])
	}

	// Sort the decorated array.
	slices.SortFunc(pairs, func(a, b Pair) int {
		return cmp(a.eKey, b.eKey)
	})

	// Undecorate each element of xAlt back into x.
	for i, p := range pairs {
		x[i] = p.e
	}
}

// SortStableFunc sorts a slice in ascending order as determined by the cmp
// function and given a function that maps each element to a sort key.
// SortStableFunc preserves the original order of equal elements.
func SortStableFunc[S ~[]E, E, Ekey any](x S, cmp func(a, b Ekey) int, key func(E) Ekey) {
	// Decorate each element of x.
	type Pair assoc[E, Ekey]
	pairs := make([]Pair, len(x))
	for i, e := range x {
		pairs[i].e = e
		pairs[i].eKey = key(x[i])
	}

	// Sort the decorated array.
	slices.SortStableFunc(pairs, func(a, b Pair) int {
		return cmp(a.eKey, b.eKey)
	})

	// Undecorate each element of xAlt back into x.
	for i, p := range pairs {
		x[i] = p.e
	}
}

// Sorted takes a slice and a function that maps each element to a sort key
// and returns a new, sorted slice.
func Sorted[E any, Ekey cmp.Ordered](seq iter.Seq[E], key func(E) Ekey) []E {
	// Decorate each element of seq.
	type Pair assoc[E, Ekey]
	pairs := make([]Pair, 0, 1000) // 1000 is arbitrary
	for e := range seq {
		p := Pair{e: e, eKey: key(e)}
		pairs = append(pairs, p)
	}

	// Sort the decorated array.
	slices.SortStableFunc(pairs, func(a, b Pair) int {
		return cmp.Compare(a.eKey, b.eKey)
	})

	// Undecorate each element of pairs into a new slice.
	xSort := make([]E, len(pairs))
	for i, p := range pairs {
		xSort[i] = p.e
	}
	return xSort
}
