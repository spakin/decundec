decundec
========

[![GitHub Actions](https://github.com/spakin/decundec/actions/workflows/go.yml/badge.svg)](https://github.com/spakin/decundec/actions) [![Go Reference](https://pkg.go.dev/badge/github.com/spakin/decundec.svg)](https://pkg.go.dev/github.com/spakin/decundec) [![Go Report Card](https://goreportcard.com/badge/spakin/decundec)](https://goreportcard.com/report/spakin/decundec)

Introduction
------------

`decundec` is a package for the [Go programming language](http://go.dev/) that sorts data using a decorate-sort-undecorate idiom to improve performance when the comparison function is time-consuming to compute.

By "decorate-sort-undecorate" we mean that
1. a sort key is associated with each data element to be sorted,
2. the data are sorted based on those keys, carrying along the original data, and finally,
3. the sort keys that were added are stripped, leaving only the sorted elements.

The advantage of decorate-sort-undecorate is that sort keys are computed only once per element, not each time an element is used in a comparison, which typically will be many more than the number of elements.  The disadvantages are that additional memory must be allocated to store the decorated slice and that more data must be moved on each element swap.

Usage
-----

`decundec` replacements are provided for all of the sorting functions in Go's [`slices`](https://pkg.go.dev/slices) package.  In all cases, the `decundec` version takes an extra argument: a function that maps an element to a sort key.  User-provided comparison functions are defined in terms of the sort key's type rather than the original element's type.

Performance
-----------

As a test of `decundec`'s performance, consider sorting a list of `uint16` values in reverse order of their base-10 digits.  For example, the list

* 46792
* 25213
* 27803
* 26265
* 33681

should be sorted as

* 33681
* 46792
* 27803
* 25213
* 26265

The following graph plots the speedup of [`decundec.SortFunc`](https://pkg.go.dev/github.com/spakin/decundec#SortFunc) and [`decundec.SortedFunc`](https://pkg.go.dev/github.com/spakin/decundec#SortedFunc) over their [`slices.SortFunc`](https://pkg.go.dev/slices#SortFunc) and [`slices.SortedFunc`](https://pkg.go.dev/slices#SortedFunc) counterparts as a function of element count:

![speedup](https://github.com/user-attachments/assets/d0a24235-39f9-4565-9440-6fefa4f79299)

Speedup is defined as
```math
100\% \cdot \left( \frac{T_\text{slices}}{T_\text{decundec}} - 1 \right)
```
Hence, bars less than 0% indicate that `slices` is faster while bars greater than 0% indicate that `decundec` is faster.  Data were gathered by running
```bash
go test -bench . -count=11 -timeout 0
```
and processing the output with [`benchstat`](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat).

The graph illustrates that for this digit-reversal test, `slices` is faster up to a list length of 32 elements and `decundec` is faster for all longer lists.  `decundec`'s `SortFunc` runs twice as fast as `slices`'s (100% speedup) while `decundec`'s `SortedFunc` runs 75% faster than `slices`'s.


Author
------

[Scott Pakin](http://www.pakin.org/~scott/), *scott-dsu@pakin.org*
