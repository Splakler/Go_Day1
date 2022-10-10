package main

import "fmt"

// This is a copy of pointers.go
func main() {
	var arr = []int{1, 2, 3}
	var pt = &arr

	fmt.Println(&pt)
	fmt.Println(arr)
	fmt.Println(cap(arr))
	pointer(pt)
	fmt.Println(&pt)
	fmt.Println(arr)
	fmt.Println(cap(arr))

}
