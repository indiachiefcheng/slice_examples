package main

import (
	"fmt"
)

func main(){
	s1 := []string{"one","","three"}
	s2 := []string{"one","","three"}
	fmt.Println("the original s1 is ",s1)
	fmt.Println("the original s2 is ",s2)
	fmt.Println("the deleted s1 is ",noneempty(s1))
	fmt.Println("the deleted s2 is ",noneempty(s2))
}

func noneempty(strings []string) []string{
	i := 0
	for _,s := range strings{
		if s!= ""{
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func noneempty2(strings []string) []string{
	out := strings[:0] //give a zero length slice of original
    for _,s := range strings{
        if s!= ""{
			out = append(out,s)
        }
    }
    return out
}

