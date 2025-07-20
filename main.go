package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	// Binary Search Examples
	fmt.Println("=== Binary Search Examples ===")

	// Receive the value of `name` from user input from keyboard
	fmt.Print("What's your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello, %s\n", name)

	// Example 1: Basic binary search
	nums := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	target := 7

	fmt.Printf("Array: %v\n", nums)
	fmt.Printf("Searching for: %d\n", target)

	index := binarySearch(nums, target)
	if index != -1 {
		fmt.Printf("Found %d at index %d\n", target, index)
	} else {
		fmt.Printf("%d not found in array\n", target)
	}
	fmt.Println()

	// Example 2: Search for non-existent element
	target2 := 8
	fmt.Printf("Searching for: %d\n", target2)
	index2 := binarySearch(nums, target2)
	if index2 != -1 {
		fmt.Printf("Found %d at index %d\n", target2, index2)
	} else {
		fmt.Printf("%d not found in array\n", target2)
	}
	fmt.Println()

	// Example 3: Search in unsorted array (need to sort first)
	unsorted := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Unsorted array: %v\n", unsorted)
	sort.Ints(unsorted)
	fmt.Printf("Sorted array: %v\n", unsorted)

	target3 := 25
	index3 := binarySearch(unsorted, target3)
	fmt.Printf("Searching for %d: ", target3)
	if index3 != -1 {
		fmt.Printf("Found at index %d\n", index3)
	} else {
		fmt.Printf("Not found\n")
	}
	fmt.Println()

	// Example 4: Edge cases
	fmt.Println("=== Edge Cases ===")

	// Empty array
	empty := []int{}
	fmt.Printf("Empty array search for 5: %d\n", binarySearch(empty, 5))

	// Single element
	single := []int{42}
	fmt.Printf("Single element [42] search for 42: %d\n", binarySearch(single, 42))
	fmt.Printf("Single element [42] search for 10: %d\n", binarySearch(single, 10))

	// First and last elements
	fmt.Printf("Search for first element (1): %d\n", binarySearch(nums, 1))
	fmt.Printf("Search for last element (19): %d\n", binarySearch(nums, 19))
	fmt.Println()

	// Example 5: Recursive vs Iterative comparison
	fmt.Println("=== Recursive vs Iterative Binary Search ===")
	testArray := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	searchValue := 12

	fmt.Printf("Array: %v\n", testArray)
	fmt.Printf("Searching for: %d\n", searchValue)

	// Iterative version
	iterativeResult := binarySearch(testArray, searchValue)
	fmt.Printf("Iterative binary search result: %d\n", iterativeResult)

	// Recursive version
	recursiveResult := binarySearchRecursive(testArray, searchValue, 0, len(testArray)-1)
	fmt.Printf("Recursive binary search result: %d\n", recursiveResult)
}

// binarySearch performs binary search on a sorted slice of integers
// Returns the index of the target element, or -1 if not found
// Time complexity: O(log n), Space complexity: O(1)
func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		// Calculate middle index (avoids integer overflow)
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid // Found the target
		}

		if arr[mid] < target {
			left = mid + 1 // Target is in the right half
		} else {
			right = mid - 1 // Target is in the left half
		}
	}

	return -1 // Target not found
}

// binarySearchRecursive is a recursive implementation of binary search
func binarySearchRecursive(arr []int, target, left, right int) int {
	// Base case: element not found
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	if arr[mid] == target {
		return mid
	}

	if arr[mid] < target {
		return binarySearchRecursive(arr, target, mid+1, right)
	} else {
		return binarySearchRecursive(arr, target, left, mid-1)
	}
}
