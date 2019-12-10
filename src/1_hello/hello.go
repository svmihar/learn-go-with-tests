package main

import "fmt"

const sapaan = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return sapaan + name + "!"
}

func main() {
	fmt.Println(Hello("Tian"))
}
