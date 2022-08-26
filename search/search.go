package search

import "fmt"

func LinearSearch(a []int, n int, x int) int {
	vt := 0
	for vt < n && a[vt] != x {
		vt++
	}
	if vt == n {
		return -1
	}
	return vt
}

func BinarySearch(a []int, n int, x int) int {
	l := 0
	r := n - 1
	vt := -1
	for l <= r {
		m := (l + r) / 2
		if a[m] == x {
			fmt.Println("Vi tri cua x: ", m)
			return m
		} else if a[m] > x {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return vt
}
