package main

import "fmt"

func main() {
	fmt.Println(hello())
}

func hello() string {
	return "Hello Go" // Исправили: было "Hello go", стало "Hello Go"
}
