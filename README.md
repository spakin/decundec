decundec
========

Introduction
------------

`decundec` is a package for the [Go programming language](http://go.dev/) that sorts data using a decorate-sort-undecorate idiom to improve performance when the comparison function is time-consuming to compute.

By "decorate-sort-undecorate" we mean that
1. a sort key is associated with each data element to be sorted,
2. the data are sorted based on those keys, carrying along the original data, and finally,
3. the sort keys that were added are stripped, leaving only the sorted elements.

Motivation
----------

Consider sorting a slice of five host names by IP address:
```Go
var hosts = []string{
    "github.com",
    "go.dev",
    "google.com",
    "sourceforge.net",
    "stackoverflow.com",
}

func getIP(host string) net.IP {
	ips, err := net.LookupIP(host)
	if err != nil {
		panic(err)
	}
	return ips[0]
}
```
Assuming the IP addresses for those host names are, respectively, 140.82.116.3, 216.239.34.21, 142.250.72.142, 104.18.13.149, and 104.18.32.7, sorting with Go's [`slices.SortFunc`](https://pkg.go.dev/slices#SortFunc) performs a total of 20 (slow) IP-address lookups: two for each of ten invocations of the comparison function:

```Go
slices.SortFunc(hosts,
    func(a, b string) int { return slices.Compare(getIP(a), getIP(b)) })
```

1. Compare go.dev and github.com (2 lookups)
2. Compare google.com and go.dev (2 lookups)
3. Compare google.com and github.com (2 lookups)
4. Compare sourceforge.net and go.dev (2 lookups)
5. Compare sourceforge.net and google.com (2 lookups)
6. Compare sourceforge.net and github.com (2 lookups)
7. Compare stackoverflow.com and go.dev (2 lookups)
8. Compare stackoverflow.com and google.com (2 lookups)
9. Compare stackoverflow.com and github.com (2 lookups)
10. Compare stackoverflow.com and sourceforge.net (2 lookups)

In contrast, sorting with `decundec.SortFunc` _first_ looks up all five IP addresses then sorts by those:

```Go
decundec.SortFunc(hosts,
    func(a, b net.IP) int { return slices.Compare(a, b) },
    getIP)
```

1. Compare 216.239.34.21 and 140.82.116.3 (0 lookups)
2. Compare 142.250.72.142 and 216.239.34.21 (0 lookups)
3. Compare 142.250.72.142 and 140.82.116.3 (0 lookups)
4. Compare 104.18.13.149 and 216.239.34.21 (0 lookups)
5. Compare 104.18.13.149 and 142.250.72.142 (0 lookups)
6. Compare 104.18.13.149 and 140.82.116.3 (0 lookups)
7. Compare 104.18.32.7 and 216.239.34.21 (0 lookups)
8. Compare 104.18.32.7 and 142.250.72.142 (0 lookups)
9. Compare 104.18.32.7 and 140.82.116.3 (0 lookups)
10. Compare 104.18.32.7 and 104.18.13.149 (0 lookups)

Although the number of comparisons has not changed, each comparison is significantly faster because it does not require calling out to the operating system to map host names to IP addresses.

Usage
-----

`decundec` replacements are provided for all of the sorting functions in Go's [`slices`](https://pkg.go.dev/slices) package.  In all cases, the `decundec` version takes an extra argument: a function that maps an element to a sort key.  User-provided comparison functions are defined in terms of the sort key's type rather than the original elements' type.


Author
------

[Scott Pakin](http://www.pakin.org/~scott/), *scott-dsu@pakin.org*
