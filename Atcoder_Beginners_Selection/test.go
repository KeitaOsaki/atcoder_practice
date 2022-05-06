package main

import (
	"fmt"
	"strings"
)

func main() {
	s := ""

	fmt.Printf("[%s]\n", strings.TrimPrefix(s, "dream")) // -> "[    123456]"
}
