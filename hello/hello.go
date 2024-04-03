package main

import "fmt"

const HELLO_PREFIX = "Hello "

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	return HELLO_PREFIX + name
}

func main() {
	fmt.Println(Hello("Mike", ""))
}
