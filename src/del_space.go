package main

import (
	"fmt"
 	"bytes"
	"unicode"
)

func del_whitespace(s string) string {
	var buffer bytes.Buffer
	skip := true
	for _, char := range s {
		if unicode.IsSpace(char) {
			if !skip {
				buffer.WriteRune(' ')
				skip = true
			}
		} else {
			buffer.WriteRune(char)
			skip = false
		}
	}
	s = buffer.String()
	if skip && len(s) > 0 {
		s = s[:len(s)-1]
	}
	return s
}

func main() {
	s := "wkm                  wkm"
	s2 := del_whitespace(s)
	fmt.Printf("%s\n", s2)
}