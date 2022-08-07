package main

// Standard packages do not have import prefixes
import (
	"fmt"
	"github.com/brigaldies/go-hello-world/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
