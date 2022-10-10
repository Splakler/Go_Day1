package main

func pointer(x_pt *[]int) {
	*x_pt = append(*x_pt, 30000)
}
