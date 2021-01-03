package main

import (
	"fmt"
	"rsc.io/quote/v3"
	"example.com/user/hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("Hello, world.")
	fmt.Println(quote.HelloV3())
	fmt.Println(morestrings.ReverseRunes("!oG, olleh"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
