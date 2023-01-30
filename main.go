package main

import "fmt"

func greet() string {
	return "hello worl"
}

func main() {
	s := greet()
	fmt.Println(s)
}
