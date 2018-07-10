package main

import "fmt"

func NewtonSqrt(x float64) float64{
	z := x/2
	for tmp:=0.0;(tmp-z>0.000000001)||(tmp-z< -0.000000001);{
		tmp = z
		z -= (z*z-x)/(2*z)
	}
	return z
}

func main(){
	res := NewtonSqrt(float64(3))
	fmt.Println(res)
}
