package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "wkm"

	fmt.Printf("%d\n", len(str))
	fmt.Printf("%d\n", []rune(str))
	fmt.Printf("%d\n", len([]rune(str)))

	//////////////
	names := "wang-kai-min"
	fmt.Print("|")
	for _, name := range strings.Split(names, "-") {
		fmt.Printf("%s|", name)
	}

	////////////////
}