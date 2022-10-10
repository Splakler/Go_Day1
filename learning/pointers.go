package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3}
	var pt = &arr

	fmt.Println(&pt)
	fmt.Println(arr)
	fmt.Println(cap(arr))
	pointers(pt)
	fmt.Println(&pt)
	fmt.Println(arr)
	fmt.Println(cap(arr))

}

func pointers(x_pt *[]int) {
	*x_pt = append(*x_pt, 30000)
}
