package main

import (
	"math/rand"
)

type PivotGetter func([]int, int, int) int
type Partitioner func([]int, int, int, int) int

var Partitioners = map[string]Partitioner{
	"hoare":  partitionHoare,
	"lomuto": partitionLomuto,
}

var PivotGetters = map[string]PivotGetter{
	"random": getRandomPivot,
	"median": getMedianPivot,
	"fixed":  func(nums []int, l, h int) int { return h },
}

func Sort(nums []int, l, h int, pivotGetter, partitioner string) {
	if l >= h {
		return
	}

	p := PivotGetters[pivotGetter](nums, l, h)
	p = Partitioners[partitioner](nums, l, h, p)

	if partitioner == "hoare" {
		Sort(nums, l, p, pivotGetter, partitioner)
	} else {
		Sort(nums, l, p-1, pivotGetter, partitioner)
	}
	Sort(nums, p+1, h, pivotGetter, partitioner)
}

func partitionHoare(nums []int, l, h, p int) int {
	nums[l], nums[p] = nums[p], nums[l]
	piv := nums[l]
	i, j := l-1, h+1

	for {
		i++
		for nums[i] < piv {
			i++
		}

		j--
		for nums[j] > piv {
			j--
		}

		if i >= j {
			return j
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func partitionLomuto(nums []int, l, h, p int) int {
	nums[p], nums[h] = nums[h], nums[p]
	p = h

	i := l
	for j := l; j < h; j++ {
		if nums[j] <= nums[p] {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[i], nums[p] = nums[p], nums[i]
	return i
}

func getRandomPivot(nums []int, l, h int) int {
	return rand.Intn(h-l) + l
}

func getMedianPivot(nums []int, l, h int) int {
	m := (l + h) / 2

	if (nums[l] < nums[m]) != (nums[l] < nums[h]) {
		return l
	}
	if (nums[m] < nums[l]) != (nums[m] < nums[h]) {
		return m
	}
	return h
}
