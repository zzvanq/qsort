package main

import (
	"math/rand"
)

type PivotSetter func(arr []int, l, h int)

var PivotSetters = map[string]PivotSetter{
	"random": setRandomPivot,
	"median": setMedianPivot,
	"fixed":  func(arr []int, l, h int) {},
}

func setRandomPivot(arr []int, l, h int) {
	p := rand.Intn(h-l) + l
	arr[h], arr[p] = arr[p], arr[h]
}

func setMedianPivot(arr []int, l, h int) {
	m := (l + h) / 2

	if arr[m] < arr[l] {
		arr[m], arr[l] = arr[l], arr[m]
	}
	if arr[h] < arr[l] {
		arr[l], arr[h] = arr[h], arr[l]
	}
	if arr[m] < arr[h] {
		arr[m], arr[h] = arr[h], arr[m]
	}
}

func Sort(arr []int, l, h int, pivotSetter PivotSetter) {
	if l >= h {
		return
	}

	p := partition(arr, l, h, pivotSetter)
	Sort(arr, l, p-1, pivotSetter)
	Sort(arr, p+1, h, pivotSetter)
}

func partition(arr []int, l, h int, pivotSetter PivotSetter) int {
	pivotSetter(arr, l, h)
	p := h

	i := l
	for j := l; j < h; j++ {
		if arr[j] <= arr[p] {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[p] = arr[p], arr[i]
	return i
}
