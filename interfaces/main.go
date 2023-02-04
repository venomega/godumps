package main

import "fmt"

func test[T []byte | string](a T) {
	fmt.Printf("The string is %T %X", a[0], a[0])
}

func main() {
	println("hello world")
	test[[]byte]([]byte("asd"))
}
