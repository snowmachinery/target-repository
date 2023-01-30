package main

import "fmt"

func greet() string {
	return "hello world"
}

func main() {
	s := greet()
	fmt.Println(s)
}
