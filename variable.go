package main

import "fmt"

var a int
var b int = 5
var c string = "test"

func main(){
	a = 10
	d := 3.14159
	a, b = b, a
	e, f, g := 7, 5.5, "TEST"
	fmt.Println(a, b, c, d, e, f, g)
	fmt.Println("test")
	fmt.Println("OK")
}

