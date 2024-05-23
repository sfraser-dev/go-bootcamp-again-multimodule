// Go doesn't need a virtual environment like Python. Instead, the Go module system manages dependencies
// Go versions: https://go.dev/dl/
// Go download: cd Downloads; wget https://go.dev/dl/go1.22.3.linux-amd64.tar.gz
// Go extract: sudo tar -C /usr/local -xzf go1.22.3.linux-amd64.tar.gz
// Go add to PATH: export PATH=$PATH:/usr/local/go/bin

package main

import (
	"fmt"
)

func main() {
	// snippet "fmpl" (or "fp")
	fmt.Println("Hello World with automatic newline!")
	// snippet "fmp"
	fmt.Print("Hello World withoutautomatic newline\n")
}
