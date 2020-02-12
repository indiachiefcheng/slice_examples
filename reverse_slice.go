package main

import (
	"fmt"
)

func main(){
	s := []string{"one","two","three","four"}
	fmt.Println(s)
	fmt.Println(reverse(s))
}

func reverse(slice []string)[]string{
	for i,j := 0,len(slice)-1;i<j;i,j = i+1,j-1{
		slice[i],slice[j] = slice[j],slice[i]
	}
	return slice
}
