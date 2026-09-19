package main

import "fmt"

func showType(t interface{}) {
	fmt.Printf("Type: %T, value: %v\n", t, t)
}

func main() {
	var x interface{} = 10
	var y interface{} = "Hello, World!"
	showType(x)
	showType(y)
}
