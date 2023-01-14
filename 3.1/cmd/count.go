package main

import (
	"counter"
	"fmt"
)

func main() {
	count := counter.Lines()
	fmt.Println(count)
}
