package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	exercise()
	secs := time.Since(start).Seconds()
	fmt.Printf("Total execution time: %.2fs\n", secs)
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
