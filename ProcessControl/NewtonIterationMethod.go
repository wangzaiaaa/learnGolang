package main

import (

	"fmt"
	"math"
)

/*
牛顿迭代法
x^m = a
f(x) = x^m -a
f(x) = 0
x^m - a = 0
f'(x) = m*x^(m-1)
初次迭代 x(n) = a / 2
x(n+1) = x(n) * (1-1/m) + a*x(n)/m*x^m
m = 2时 x(n+1) = (x(n) + a/x(n)) / 2
*/

//牛顿迭代法
func newtonIterationMethod(num float64) {
	x := math.Sqrt(num)
	y := float64(num / 2)
	for i := 0;math.Abs(y-x)>0.000001;i++{
		y = (y*1.0 + (1.0*num)/y) / 2.0
		fmt.Println("第",i,"次迭代","  ",y,x)
	}

	}
//二分法
func dichotomy(num float64){
	low,hi := 0.0,num
	x := math.Sqrt(num)
	y := low + (hi - low) / 2
	for i := 0;math.Abs(x-y) > 1e-6;i++{
		if y*y > num{
			hi = y
		}else{
			low = y
		}
		fmt.Println("第",i,"次迭代","   ",y,x)
		y = low + (hi - low) / 2
	}
}
func main(){
	fmt.Println("牛顿迭代法：")
	newtonIterationMethod(100)
	fmt.Println("二分法：")
	dichotomy(100)
}
