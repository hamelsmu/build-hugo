package main

import (
	"fmt"

	"example.com/user/hello/morestrings"
	"github.com/google/go-cmp/cmp"
	"github.com/build-hugo/go_module"
)

func main() {
	fmt.Println("Hello, world.")
	fmt.Println(morestrings.ReverseRunes("!oG, olleh"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
