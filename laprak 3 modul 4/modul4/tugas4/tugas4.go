package main

import "fmt"

func main() {
	var bilangan int
	var digit1, digit2, digit3 int
	var berurutan bool

	fmt.Print("masukkan bilangan tiga digit: ")
	fmt.Scan(&bilangan)

	digit1 = bilangan / 100
	digit2 = (bilangan / 10) % 10
	digit3 = bilangan % 10

	berurutan = (digit1 < digit2 && digit2 < digit3) || (digit1 > digit2 && digit2 > digit3)
	fmt.Println(berurutan)
}