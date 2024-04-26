// IN the name of god
package main

import (
	"fmt"
	"time"
)

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
		case 2:
		case 3:
			flag = true
			fmt.Println("you are logged out.")
		default:
			fmt.Print("Error: ")
			fmt.Println("Entered value is not valid.")
		}
	}
}
