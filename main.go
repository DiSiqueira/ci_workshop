package main

import "log"
import "fmt"

func main() {
	Hello()
}

// Hello my func.
func Hello() {
	log.Println("Hello")
	fmt.Println("World")
}
