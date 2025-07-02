package main

import (
	"fmt"
	"os"
)

func main() {
	exercise()
}

func example() {
	s, sep := "", " "
	for _, arg := range os.Args[1:] {
		s += sep + arg
	}
	fmt.Println(s)
}

func exercise() {
	for index, arg := range os.Args[1:] {
		fmt.Println(index, arg)
	}
}
