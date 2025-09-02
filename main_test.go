package main

import (
	"math/rand"
	"testing"
)

const ARR_SIZE = 1000
const ARRS_SIZE = 110000 // a bit hacky, but this way there is more arrays than iterations

var arrs [][]int
var arrsrand [][]int
var arrsnone [][]int

func init() {
	arrs = make([][]int, ARRS_SIZE)
	arrsrand = make([][]int, ARRS_SIZE)
	arrsnone = make([][]int, ARRS_SIZE)
	for i := range arrs {
		arr := make([]int, ARR_SIZE)
		for y := range ARR_SIZE {
			arr[y] = rand.Intn(1000)
		}
		arrs[i] = arr
		arrsrand[i] = append([]int{}, arr...)
		arrsnone[i] = append([]int{}, arr...)
	}
}

func BenchmarkSort__median(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sort(arrs[i%ARRS_SIZE], 0, len(arrs[i%ARRS_SIZE])-1, PivotSetters["median"])
	}
}

func BenchmarkSort__random(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sort(arrsrand[i%ARRS_SIZE], 0, len(arrsrand[i%ARRS_SIZE])-1, PivotSetters["random"])
	}
}

func BenchmarkSort__fixed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sort(arrsnone[i%ARRS_SIZE], 0, len(arrsnone[i%ARRS_SIZE])-1, PivotSetters["fixed"])
	}
}
