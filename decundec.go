/*
Package decundec wraps the sorting functions in Go's slices package to
employ a decorate-sort-undecorate idiom.  This can improve sorting
performance when the comparison function is time-consuming to compute.

decorate-sort-undecorate works by
① associating a sort key with each data element to be sorted,
② sorting based on those keys, carrying along the original data, and finally,
③ removing the sort keys that were added, leaving only the sorted elements.

The advantage of decorate-sort-undecorate is that sort keys are
computed only once per element, not each time an element is used in a
comparison, which typically will be many more than the number of
elements.  The disadvantages are that additional memory must be
allocated to store the decorated slice and that more data must be
moved on each element swap.
*/
package decundec

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

// sortHelper decorates each element of a slice, sorts the new slice using
// given sort and comparison functions, and overwrites the original slice
// with the undecorated sorted slice.
func sortHelper[S ~[]E, E, Ekey any](x S,
	cmp func(a, b Ekey) int,
	key func(E) Ekey,
	sort func([]assoc[E, Ekey], func(a, b assoc[E, Ekey]) int)) {
	// Decorate each element of x.
	pairs := make([]assoc[E, Ekey], len(x))
	for i, e := range x {
		pairs[i].e = e
		pairs[i].eKey = key(x[i])
	}

	// Sort the decorated slice
	sort(pairs, func(a, b assoc[E, Ekey]) int {
		return cmp(a.eKey, b.eKey)
	})

	// Undecorate each element of pairs back into x.
	for i, p := range pairs {
		x[i] = p.e
	}
}

// Sort sorts a slice in ascending order given a function that maps each
// element to a sort key.
func Sort[S ~[]E, E any, Ekey cmp.Ordered](x S, key func(E) Ekey) {
	sortHelper(x, cmp.Compare, key, slices.SortFunc)
}

// SortFunc sorts a slice in ascending order as determined by the cmp
// function and given a function that maps each element to a sort key.
func SortFunc[S ~[]E, E, Ekey any](x S, cmp func(a, b Ekey) int, key func(E) Ekey) {
	sortHelper(x, cmp, key, slices.SortFunc)
}

// SortStableFunc sorts a slice in ascending order as determined by the cmp
// function and given a function that maps each element to a sort key.
// SortStableFunc preserves the original order of elements with equal keys.
func SortStableFunc[S ~[]E, E, Ekey any](x S, cmp func(a, b Ekey) int, key func(E) Ekey) {
	sortHelper(x, cmp, key, slices.SortStableFunc)
}

// sortedHelper decorates each element of a sequence into a slice, sorts
// the slice using given sort and comparison functions, and returns
// undecorated sorted slice.
func sortedHelper[E, Ekey any](seq iter.Seq[E],
	cmp func(a, b Ekey) int,
	key func(E) Ekey,
	sort func([]assoc[E, Ekey], func(a, b assoc[E, Ekey]) int)) []E {
	// Decorate each element of seq.
	pairs := make([]assoc[E, Ekey], 0, 1000) // 1000 is arbitrary.
	for e := range seq {
		p := assoc[E, Ekey]{e: e, eKey: key(e)}
		pairs = append(pairs, p)
	}

	// Sort the decorated slice
	sort(pairs, func(a, b assoc[E, Ekey]) int {
		return cmp(a.eKey, b.eKey)
	})

	// Undecorate each element of pairs into a new slice.
	xSort := make([]E, len(pairs))
	for i, p := range pairs {
		xSort[i] = p.e
	}
	return xSort
}

// Sorted takes a slice and a function that maps each element to a sort key
// and returns a new, sorted slice.
func Sorted[E any, Ekey cmp.Ordered](seq iter.Seq[E], key func(E) Ekey) []E {
	return sortedHelper(seq, cmp.Compare, key, slices.SortFunc)
}

// SortedFunc takes a slice, a key-comparison function, and a function that
// maps each element to a sort key; it returns a new, sorted slice.
func SortedFunc[E any, Ekey cmp.Ordered](seq iter.Seq[E],
	cmp func(a, b Ekey) int,
	key func(E) Ekey) []E {
	return sortedHelper(seq, cmp, key, slices.SortFunc)
}

// SortedStableFunc takes a slice, a key-comparison function, and a
// function that maps each element to a sort key; it returns a new, sorted
// slice.  SortedStableFunc preserves the original order of elements with
// equal keys.
func SortedStableFunc[E any, Ekey cmp.Ordered](seq iter.Seq[E],
	cmp func(a, b Ekey) int,
	key func(E) Ekey) []E {
	return sortedHelper(seq, cmp, key, slices.SortStableFunc)
}
