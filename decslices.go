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

// Sort sorts a slice in ascending order given a function that maps each
// element to a sort key.
func Sort[S ~[]E, E any, Ealt cmp.Ordered](x S, key func(E) Ealt) {
	// Decorate each element of x.
	type Wrapper struct {
		e    E
		eAlt Ealt
	}
	xAlt := make([]Wrapper, len(x))
	for i, e := range x {
		xAlt[i].e = e
		xAlt[i].eAlt = key(x[i])
	}

	// Sort the decorated array.
	slices.SortStableFunc(xAlt, func(a, b Wrapper) int {
		return cmp.Compare(a.eAlt, b.eAlt)
	})

	// Undecorate each element of xAlt back into x.
	for i, w := range xAlt {
		x[i] = w.e
	}
}

// SortFunc sorts a slice in ascending order as determined by the cmp
// function and given a function that maps each element to a sort key.
func SortFunc[S ~[]E, E, Ealt any](x S, cmp func(a, b Ealt) int, key func(E) Ealt) {
	// Decorate each element of x.
	type Wrapper struct {
		e    E
		eAlt Ealt
	}
	xAlt := make([]Wrapper, len(x))
	for i, e := range x {
		xAlt[i].e = e
		xAlt[i].eAlt = key(x[i])
	}

	// Sort the decorated array.
	slices.SortFunc(xAlt, func(a, b Wrapper) int {
		return cmp(a.eAlt, b.eAlt)
	})

	// Undecorate each element of xAlt back into x.
	for i, w := range xAlt {
		x[i] = w.e
	}
}

// SortStableFunc sorts a slice in ascending order as determined by the cmp
// function and given a function that maps each element to a sort key.
// SortStableFunc preserves the original order of equal elements,
func SortStableFunc[S ~[]E, E, Ealt any](x S, cmp func(a, b Ealt) int, key func(E) Ealt) {
	// Decorate each element of x.
	type Wrapper struct {
		e    E
		eAlt Ealt
	}
	xAlt := make([]Wrapper, len(x))
	for i, e := range x {
		xAlt[i].e = e
		xAlt[i].eAlt = key(x[i])
	}

	// Sort the decorated array.
	slices.SortStableFunc(xAlt, func(a, b Wrapper) int {
		return cmp(a.eAlt, b.eAlt)
	})

	// Undecorate each element of xAlt back into x.
	for i, w := range xAlt {
		x[i] = w.e
	}
}

// Sorted takes a slice and a function that maps each element to a sort key
// and returns a new, sorted slice.
func Sorted[E any, Ealt cmp.Ordered](seq iter.Seq[E], key func(E) Ealt) []E {
	// Decorate each element of seq.
	type Wrapper struct {
		e    E
		eAlt Ealt
	}
	xAlt := make([]Wrapper, 0, 64) // 64 is arbitrary.
	for e := range seq {
		w := Wrapper{e: e, eAlt: key(e)}
		xAlt = append(xAlt, w)
	}

	// Sort the decorated array.
	slices.SortStableFunc(xAlt, func(a, b Wrapper) int {
		return cmp.Compare(a.eAlt, b.eAlt)
	})

	// Undecorate each element of xAlt into a new slice.
	xSort := make([]E, len(xAlt))
	for i, w := range xAlt {
		xSort[i] = w.e
	}
	return xSort
}
