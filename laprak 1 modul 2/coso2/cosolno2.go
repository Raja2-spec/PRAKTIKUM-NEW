package main

import "fmt"

func main() {
	
	var x, fx float32
	fmt.Scan(&x)
	fx = 2/(x+5) + 5
	fmt.Print(fx)
}