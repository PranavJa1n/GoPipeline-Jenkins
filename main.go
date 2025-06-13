package main

import "fmt"

func hello(num int) string {
	if num == 7 {
		return ""
	}
	return "Hello"
}

func main() {
	fmt.Println(hello(10))
}
