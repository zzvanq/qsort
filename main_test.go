package main

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	partitioners := []string{"hoare", "lomuto"}
	pivotGetters := []string{"random", "median", "fixed"}
	src := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, pivotGetter := range pivotGetters {
		for _, partitioner := range partitioners {
			nums := append([]int(nil), src...)
			Sort(nums, 0, len(nums)-1, pivotGetter, partitioner)
			if !reflect.DeepEqual(nums, want) {
				t.Errorf("%s_%s: got %v, want %v", partitioner, pivotGetter, nums, want)
			}
		}
	}
}
