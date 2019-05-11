package main

import "fmt"

func main(){
	a := []int {1,2,3,4,5,6,7}
	b := make([]int,len(a),(cap(a)+1)*2)
	for i := range a{
		b[i] = a[i]
	}
	for i := range b{
		fmt.Print(i," ")
	}
	}
