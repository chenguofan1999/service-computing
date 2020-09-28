package main

import "fmt"

func sort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	lo, hi := 0, len(a)-1
	pivot := a[lo]
	i, j := lo+1, hi

	for lo < hi {
		for a[i] < pivot && i != hi {
			i++
		}
		for a[j] > pivot && j != lo {
			j--
		}
		a[i], a[j] = a[j], a[i]
	}
	a[j], a[lo] = a[lo], a[j]

	sort(a[lo:j])
	sort(a[j+1 : hi])
	return a
}

func main() {
	a := []int{5, 4, 3, 2, 1}
	sort(a)
	fmt.Println(a)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
