/*
// Sample code to perform I/O:

fmt.Scanf("%s", &myname)            // Reading input from STDIN
fmt.Println("Hello", myname)        // Writing output to STDOUT

// Warning: Printing unwanted or ill-formatted data to output will cause the test cases to fail
*/

// Write your code here

package main

import "fmt"

func main() {
	var n int
	fmt.Printf("enter size of array")
	fmt.Scan(&n)
	array1 := make([]int, n)
	//array1 := [5]int{2, 5, 9, 0, 1}
	fmt.Printf("enter the %d numbers", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&array1[i])
	}
	left := 0
	right := len(array1) - 1
	//fmt.Println(array1)

	for left < right {
		array1[left], array1[right] = array1[right], array1[left]
		right--
		left++
	}
	fmt.Println("output")
	for j := 0; j < n; j++ {
		fmt.Println(array1[j])
	}
}
