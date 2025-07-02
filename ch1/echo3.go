package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	s := strings.Join(os.Args[1:], " ")
	secs := time.Since(start).Seconds()
	fmt.Printf("Total execution time: %.2fs Inputs: %s \n", secs, s)
}
