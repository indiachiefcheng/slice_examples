package main

import (
	"fmt"
)

func main(){
	s1 := []int{1,2,3,4,5,6}
	s2 := []int{1,2,3,4,5,6}
	fmt.Println("the original s1 result is",s1)
	fmt.Println("the original s2 result is",s2)
    fmt.Println("the deleted s1 result is",remove(s1,2))
    fmt.Println("the deleted s2 result is",remove2(s2,2))
}

//remove the i element of a slice,it will keep the sequence of the slice
func remove(slice []int, i int) []int{
	copy(slice[i:],slice[i+1:])
	return slice[:len(slice)-1]
}

//remove the i element of a slice,it wont keep the sequence of the slice
func remove2(slice []int, i int) []int{
	slice[i] = slice[len(slice)-1]
    return slice[:len(slice)-1]
}


