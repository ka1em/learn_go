package main

import (
	"bytes"
	"fmt"
)
// intsToString is like fmt.Sprintf(values) but adds commas
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

// main ...
func main()  {
	fmt.Println(intsToString([]int{1, 2, 3}))
}