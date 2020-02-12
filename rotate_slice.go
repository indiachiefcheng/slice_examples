package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(rotate(a, 2))
}
func rotate(s []int, r int) []int {
	lens := len(s)
	arr := make([]int, lens)
	for k := range s {
		index := r + k
		if index >= lens {
			index -= lens
		}
		arr[k] = s[index]
	}
	return arr
}
