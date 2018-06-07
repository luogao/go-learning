package main

import "fmt"

func add(x , y int) int{
	return x + y
}

func swap(a ,b string) (string, string) {
	return b, a
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main () {
	fmt.Println(add(43,22))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(55))
}

