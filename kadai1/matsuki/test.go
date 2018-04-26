package main

import "fmt"

// imports:  import "fmt"

// golint error
func Greet() {
	fmt.Println("Hello, World!")
}

func main() {
	Greet()
	// vet error
	return
	fmt.Println("Hello, World!") // 到達しないコード
}
