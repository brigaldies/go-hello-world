package main

// Standard packages do not have import prefixes
import (
	"fmt"
	"github.com/brigaldies/go-hello-world/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
}
