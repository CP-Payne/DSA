package main

import "fmt"

func testSort(sorting func(arr []int)) {
	testCases := map[string][]int{
		"randomIntegers": {34, 7, 23, 32, 5, 62, 32, 99, 2, 74},
		"alreadySorted":  {1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		"reverseSorted":  {10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		"nearlySorted":   {1, 2, 3, 7, 5, 6, 4, 8, 9, 10},
		"duplicates":     {5, 1, 3, 3, 8, 7, 5, 5, 9, 2},
		"identical":      {4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
	}

	for k, v := range testCases {
		fmt.Printf("%v: %v\n", k, v)
		sorting(v)
		fmt.Printf("Sorted: %v\n", v)
		fmt.Println()
	}
}

func testSortReturn(sorting func(arr []int) []int) {
	testCases := map[string][]int{
		"randomIntegers": {34, 7, 23, 32, 5, 62, 32, 99, 2, 74},
		"alreadySorted":  {1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		"reverseSorted":  {10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		"nearlySorted":   {1, 2, 3, 7, 5, 6, 4, 8, 9, 10},
		"duplicates":     {5, 1, 3, 3, 8, 7, 5, 5, 9, 2},
		"identical":      {4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
	}

	for k, v := range testCases {
		fmt.Printf("%v: %v\n", k, v)
		sortedArr := sorting(v)
		fmt.Printf("Sorted: %v\n", sortedArr)
		fmt.Println()
	}
}

func main() {
	// testSort(SelectionSort)
	// testSort(BubbleSort)
	// testSort(InsertionSort)
	testSortReturn(MergeSort)
}
