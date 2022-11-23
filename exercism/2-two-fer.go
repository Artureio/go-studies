package main

import "fmt"

var name string
var counter int

func main() {
	fmt.Println((twofer("teste")))
}

func twofer(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
