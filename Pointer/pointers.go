package main

import "fmt"

func main()  {
	i,j := 42,2701
	p := &i
	fmt.Println("%T",p)
	fmt.Println(*p)
	fmt.Println(i)
	p = &j
	*p = *p / 37
	fmt.Println(j)
}
