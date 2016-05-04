package main

import (
	"fmt"
)

// nonempty ...
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

// non ...
func nonempty2(strings []string) []string  {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	data := []string{"one", "", "3", "4"}
	fmt.Printf("data     = %q\n", data)
	fmt.Printf("non1data = %q\n", nonempty(data))
	fmt.Printf("non2data = %q\n", nonempty2(data))
}