package main

import (
	"fmt"

	"ratwater.xyz/ratwater"
)

func main() {
	fmt.Printf("main() response from ratwater_mod: %s\n", ratwater.Hello("world"))
}
