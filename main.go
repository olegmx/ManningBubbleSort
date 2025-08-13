/*
*****************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

******************************************************************************
*/
package main

import "fmt"
import "math/rand"

func makeRandomSlice(numItems, max int) []int {
	slice := make([]int, numItems)
	for i := 0; i < numItems; i++ {
		slice[i] = rand.Intn(max)
	}
	return slice
}

func printSlice(slice []int, numItems int) {
	if numItems > len(slice) {
		fmt.Println(slice)
	} else {
		fmt.Println(slice[:numItems])
	}
}

func checkSorted(slice []int) {
	for i := 1; i < len(slice); i++ {
		if slice[i] < slice[i-1] {
			fmt.Println("\nSlice is not sorted")
			return
		}
	}
	fmt.Println("\nSlice is sorted")
}

func bubbleSort(slice []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(slice); i++ {
			if slice[i-1] > slice[i] {
				slice[i], slice[i-1] = slice[i-1], slice[i]
				swapped = true
			}
		}
	}
}

func main() {
	var numItems, max, itemsToPrint int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)
	fmt.Printf("Items to print: ")
	fmt.Scanln(&itemsToPrint)

	// Make and display an unsorted slice.
	slice := makeRandomSlice(numItems, max)
	fmt.Println("\nSource slice:")
	printSlice(slice, itemsToPrint)
	fmt.Println()

	// Sort and display the result.
	bubbleSort(slice)
	fmt.Println("\nSorted slice with print limit:")
	printSlice(slice, itemsToPrint)
	fmt.Println("\nWhole sorted slice")
	printSlice(slice, len(slice))

	// Verify that it's sorted.
	checkSorted(slice)
}
