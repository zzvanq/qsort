package main

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

const ARR_SIZE = 1000
const ARRS_SIZE = 150000 // a bit hacky, but this way there is more arrays than iterations

var arrs [][]int
var arrsMedian [][]int
var arrsRand [][]int
var arrsFixed [][]int

func init() {
	lrand := rand.New(rand.NewSource(42))
	arrs = make([][]int, ARRS_SIZE) // this is for the builtin sort
	arrsMedian = make([][]int, ARRS_SIZE)
	arrsRand = make([][]int, ARRS_SIZE)
	arrsFixed = make([][]int, ARRS_SIZE)
	for i := range arrsMedian {
		arr := make([]int, ARR_SIZE)
		for y := range ARR_SIZE {
			arr[y] = lrand.Intn(1000)
		}
		arrs[i] = arr
		arrsMedian[i] = append([]int{}, arr...)
		arrsRand[i] = append([]int{}, arr...)
		arrsFixed[i] = append([]int{}, arr...)
	}
}

func BenchmarkSort__builtin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Ints(arrs[i%ARRS_SIZE])
	}
}

func BenchmarkSort__median(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sort(arrsMedian[i%ARRS_SIZE], 0, len(arrsMedian[i%ARRS_SIZE])-1, PivotSetters["median"])
	}
}

func BenchmarkSort__random(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sort(arrsRand[i%ARRS_SIZE], 0, len(arrsRand[i%ARRS_SIZE])-1, PivotSetters["random"])
	}
}

func BenchmarkSort__fixed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sort(arrsFixed[i%ARRS_SIZE], 0, len(arrsFixed[i%ARRS_SIZE])-1, PivotSetters["fixed"])
	}
}

func TestSort__median(t *testing.T) {
	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	Sort(nums, 0, len(nums)-1, PivotSetters["median"])
	if !reflect.DeepEqual(nums, want) {
		t.Errorf("%v, want %v", nums, want)
	}
}
