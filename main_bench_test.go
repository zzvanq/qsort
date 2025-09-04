package main

import (
	"math/rand"
	"sort"
	"testing"
)

const ARR_SIZE = 100
const ARRS_SIZE = 1000
const SORTED_SIZE = int(ARR_SIZE * 0.5)

var unsrtdNums [][]int
var partSrtdNums [][]int

func init() {
	lrand := rand.New(rand.NewSource(42))

	unsrtdNums = make([][]int, ARRS_SIZE)
	for i := range unsrtdNums {
		arr := make([]int, ARR_SIZE)
		for y := range ARR_SIZE {
			arr[y] = lrand.Intn(1000)
		}
		unsrtdNums[i] = arr
	}

	partSrtdNums = duplicateNums(unsrtdNums, ARRS_SIZE)
	for i := range partSrtdNums {
		k := lrand.Intn(ARR_SIZE - SORTED_SIZE)
		sort.Ints(partSrtdNums[i][k:])

		// The only case where Hoare is faster is when there is a plenty of duplicates
		// for k := range SORTED_SIZE {
		// 	partSrtdNums[i][k] = 42
		// }
	}
}

func duplicateNums(src [][]int, n int) [][]int {
	nums := make([][]int, n)
	for i := range nums {
		nums[i] = append([]int{}, src[i%ARRS_SIZE]...)
	}
	return nums
}

func BenchmarkSort__builtin(b *testing.B) {
	sorted := map[string][][]int{
		"sorted":   partSrtdNums,
		"unsorted": unsrtdNums,
	}

	for sk, src := range sorted {
		b.Run(sk, func(b *testing.B) {
			nums := duplicateNums(src, b.N)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				sort.Ints(nums[i])
			}
		})
	}
}

func BenchmarkSort(b *testing.B) {
	sorted := map[string][][]int{
		"sorted":   partSrtdNums,
		"unsorted": unsrtdNums,
	}
	partitioners := []string{"hoare", "lomuto"}
	pivotGetter := []string{"random", "median", "fixed"}

	for isSorted, src := range sorted {
		for _, pivotGetter := range pivotGetter {
			for _, partitioner := range partitioners {
				b.Run(isSorted+"_"+partitioner+"_"+pivotGetter, func(b *testing.B) {
					nums := duplicateNums(src, b.N)
					b.ResetTimer()

					for i := 0; i < b.N; i++ {
						Sort(nums[i], 0, len(nums[i])-1, pivotGetter, partitioner)
					}
				})
			}
		}
	}
}
