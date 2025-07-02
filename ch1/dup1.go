package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {

		line := input.Text()
		fmt.Println("line: ", line)
		fmt.Println("counts: ", counts[line])
		counts[input.Text()]++
		fmt.Println("counts: ", counts)
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
