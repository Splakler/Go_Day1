package main

import "fmt"

func main() {
	var i = 2
	var j = 2
	for i = 1; i <= 50; i++ {
		if i%j == 0 {
			fmt.Println(i)
			j = i / 2
		} else {
			j++
		}
	}

}
