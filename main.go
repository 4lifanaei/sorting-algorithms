// IN the name of god
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

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

		case 2:
			arr := []int{10, 7, 8, 9, 1, 5}
			n := len(arr)
			quickSort(arr, 0, n-1)
			fmt.Println("Sorted array:", arr)
		case 3:
			flag = true
			fmt.Println("you are logged out.")
		default:
			fmt.Print("Error: ")
			fmt.Println("Entered value is not valid.")
		}
	}
}
