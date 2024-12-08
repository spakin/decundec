/*
   Package decslices wraps the sorting functions in Go's slices
   package to employ a decorate-sort-undecorate idiom.  This can
   improve sorting performance when the comparison function is
   time-consuming to compute.
*/

package decslices

import (
	"cmp"
	"slices"
)

func Sort[S ~[]E, E cmp.Ordered](x S) {
	slices.Sort(x)
}
