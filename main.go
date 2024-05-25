// IN the name of god
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var quickSortMovement int
var quickSortComparison int

func quickSort(arr []int, start, end int) {
	if end <= start {
		return //base case
	}
	pivot := partition(arr, start, end)
	quickSort(arr, start, pivot-1)
	quickSort(arr, pivot+1, end)

}

func partition(arr []int, start, end int) int {
	pivot := arr[end]
	i := start - 1

	for j := start; j <= end-1; j++ {
		if arr[j] < pivot {
			i++
			//swaping the numbers
			arr[i], arr[j] = arr[j], arr[i]
			quickSortMovement++
		}
		quickSortComparison++
	}
	i++
	arr[i], arr[end] = arr[end], arr[i]
	quickSortMovement++
	return i //+1
}

var mergeSortMovement int
var mergeSortComparison int

func mergeSort(arr []int) {
	if len(arr) <= 1 {
		return // Base case: already sorted or empty
	}

	mid := len(arr) / 2
	left := make([]int, mid)
	right := make([]int, len(arr)-mid)

	copy(left, arr[:mid])
	copy(right, arr[mid:])

	mergeSort(left)  // Recursively sort left half
	mergeSort(right) // Recursively sort right half

	merge(arr, left, right) // Merge the sorted halves
}

func merge(arr, left, right []int) {
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
			mergeSortMovement++
		} else {
			arr[k] = right[j]
			j++
			mergeSortMovement++
		}
		k++
		mergeSortComparison++
	}

	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
		mergeSortMovement++
	}

	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
		mergeSortMovement++
	}
}

// The below code is a merge sort function that i wrote before the up code,
// without using Copy() & make() functions wich it has bugs and not working correctly.
// -------------------------------------------------------
// func mergSort(array []int) []int {
// 	inputLength := len(array)
// 	if inputLength < 2 {
// 		return array
// 	}
// 	midIndex := inputLength / 2
// 	leftHalf := array[:midIndex]
// 	rigthHalf := array[midIndex:]
// 	for i := 0; i < midIndex; i++ {
// 		leftHalf[i] = array[i]
// 	}
// 	for i := midIndex; i < inputLength; i++ {
// 		rigthHalf[i-midIndex] = array[i]
// 	}
// 	mergSort(leftHalf)
// 	mergSort(rigthHalf)
// 	//Merge
// 	return merge(array, leftHalf, rigthHalf)
// }

// func merge(inputArray []int, leftHalf []int, rigthHalf []int) []int {
// 	leftSize := len(leftHalf)
// 	rigthSize := len(rigthHalf)
// 	i := 0 //for left half
// 	j := 0 //for right half
// 	k := 0 //for our merged array
// 	for i < leftSize && j < rigthSize {
// 		if leftHalf[i] <= rigthHalf[j] { //if it was less, then it will add to the merged array
// 			inputArray[k] = leftHalf[i]
// 			i++ //increment i to go and see the next element of the leftHalf array
// 		} else {
// 			inputArray[k] = rigthHalf[j]
// 			j++
// 		}
// 		k++
// 	} //adding other remaining datas in the merged array
// 	for i < leftSize { //by pass every remaining elemnt from the left array
// 		inputArray[k] = leftHalf[i]
// 		i++
// 		k++
// 	}
// 	for j < rigthSize { //by pass every remaining elemnt from the right array
// 		inputArray[k] = rigthHalf[j]
// 		j++
// 		k++
// 	}
// 	return inputArray
// }

func main() {
	var flag bool = false
	//creating a loop for repeating the options
	for !flag {
		fmt.Println("Sorting algorithms programs:")
		fmt.Println("1) Quick sort")
		fmt.Println("2) Merge sort")
		fmt.Println("3) Exit")
		fmt.Print("choose your option: ")
		var inputOption int
		fmt.Scanln(&inputOption)
		switch inputOption {
		case 1:
			fmt.Print("Creating a random list")
			time.Sleep(1 * time.Second)
			fmt.Print(".")
			time.Sleep(1 * time.Second)
			fmt.Print(".")
			time.Sleep(1 * time.Second)
			fmt.Print(".")
			time.Sleep(1 * time.Second)
			fmt.Println()
			//creating random list:
			var qlist = [1000]int{}
			for i := range qlist {
				qlist[i] = rand.Intn(1000)
			}
			for j := range qlist {
				p := qlist[j]
				fmt.Print(p, " - ")
			}
			fmt.Println()
			//arr := []int{10, 7, 8, 9, 1, 5}
			n := len(qlist)
			quickSort(qlist[:], 0, n-1)
			fmt.Println("Sorted array:", qlist)
			fmt.Println()
			fmt.Println("Number of comparisons: ", quickSortComparison)
			fmt.Println("Number of movements: ", quickSortMovement)
		case 2:
			fmt.Print("Creating a random list")
			time.Sleep(1 * time.Second)
			fmt.Print(".")
			time.Sleep(1 * time.Second)
			fmt.Print(".")
			time.Sleep(1 * time.Second)
			fmt.Print(".")
			time.Sleep(1 * time.Second)
			fmt.Println()
			//creating random list:
			qlist := [1000]int{}
			for i := range qlist {
				qlist[i] = rand.Intn(1000)
			}
			for j := range qlist {
				p := qlist[j]
				fmt.Print(p, " - ")
			}
			fmt.Println()
			// sorting
			mergeSort(qlist[:])
			//print the sorted list
			fmt.Print("Sorted list: ")
			for j := range qlist {
				p := qlist[j]
				fmt.Print(p, " - ")
			}
			fmt.Println()
			fmt.Println("Number of comparisons: ", mergeSortComparison)
			fmt.Println("Number of movements: ", mergeSortMovement)
			//tests:
			//arr := []int{9, 7, 2, 11, 1, 39, 54}
			// fmt.Println("Original array:", qlist)
			// mergSort(arr)
			// fmt.Println("Sorted array:", qlist)
		case 3:
			flag = true
			fmt.Println("you are logged out.")
		default:
			fmt.Print("Error: ")
			fmt.Println("Entered value is not valid.")
		}
	}
}
