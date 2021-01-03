package main

import (
	"fmt"
	"github.com/hamelsmu/test_go_module/hello" // this is where the hello module package is imported.
)

func main() {
	fmt.Println(hello.Proverb())
}
