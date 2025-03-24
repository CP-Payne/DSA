package main

// ===========================================================================================
// ================================== Recommended Sorting Algorithms =========================
// ===========================================================================================

import "slices"

// Quick Sort
// One of the fastest sorting Algorithms
// Best case: (n log n)
// Average Case: (n log n)
// Worst case: O(n^2) --> if the array is already sorted, or when the smallest or largest elements are always chosen for the pivot

func QuickSort(arr []int, low, high int) {
	if high == -1 {
		high = len(arr) - 1
	}

	if low < high {
		pivotIndex := Partition(arr, low, high)
		QuickSort(arr, low, pivotIndex-1)
		QuickSort(arr, pivotIndex+1, high)
	}
}

func Partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i+1], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// Merge Sort
// O(log n)
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	// leftHalf := arr[:mid]
	// rightHalf := arr[mid:]

	// Ensure full independence. Prevent unexpected bugs if modified due to referencing same underlying array.
	leftHalf := slices.Clone(arr[:mid])
	rightHalf := slices.Clone(arr[mid:])

	sortedLeft := MergeSort(leftHalf)
	sortedRight := MergeSort(rightHalf)

	return Merge(sortedLeft, sortedRight)
}

func Merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// ===========================================================================================
// ================================== Sorting Not Recommended ================================
// ===========================================================================================

// Selection Sort
// O(n^2)
// No real use, only for academics
func SelectionSort(arr []int) {
	n := len(arr)

	// for i := 0; i < n-1; i++ {
	for i := range n - 1 {
		min_idx := i

		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min_idx] {
				min_idx = j
			}
		}

		arr[i], arr[min_idx] = arr[min_idx], arr[i]
	}
}

// BubbleSort
// O(n^2)
// If already sorted, then O(n)
// No real use, only for academics
func BubbleSort(arr []int) {
	n := len(arr)

	// for i := 0; i < n-1; i++ {
	for i := range n - 1 {
		swapped := false
		for j := range n - 1 - i {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}

// InsertionSort
// O(n^2) --> Worst case and average case
// O(n) --> Best case, if it already sorted.
// Stable sorting algorithm, efficient for small or nearly sorted lists
func InsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}
