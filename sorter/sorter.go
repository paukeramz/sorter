package sorter

import "fmt"

func Help() {
	fmt.Println("This is a helper function")
}

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				swap(&arr[i], &arr[j])
			}
		}
	}
	return arr
}

func swap(numA *int, numB *int) {
	temp := *numA
	*numA = *numB
	*numB = temp
}

func QuickSort(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, left int, right int) []int {
	if left >= right {
		return arr
	}
	pivot := right
	newPivot := partition(arr, left, right, pivot)
	quickSort(arr, left, newPivot-1)
	quickSort(arr, newPivot+1, right)
	return arr
}

func partition(arr []int, left int, right int, pivot int) int {
	newPivot := left
	for i := left; i < len(arr); i++ {
		if arr[i] < arr[pivot] {
			swap(&arr[i], &arr[newPivot])
			newPivot++
		}
	}
	swap(&arr[newPivot], &arr[pivot])
	return newPivot
}
