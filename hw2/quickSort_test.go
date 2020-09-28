package main

import "testing"

func TestSort(t *testing.T) {
	array := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	sortedArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	array = sort(array)
	for i := 0; i < 5; i++ {
		if sortedArray[i] != array[i] {
			t.Error("unsorted!")
		}
	}
}
