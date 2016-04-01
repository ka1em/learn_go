package main

import (
	"fmt"
	"strings"
)

type BitFlag int
const (
	Active BitFlag = 1 << iota	// 1<<0
	Send						// 1<<1
	Receive						// 1<<2
)

func (flag BitFlag) String() string {
	var flags []string
	if flag & Active == Active {
		flags = append(flags, "Active")
	}
	if flag & Send == Send {
		flags = append(flags, "Send")
	}
	if flag & Receive == Receive {
		flags = append(flags, "Receive")
	}
	if len(flags) > 0 {
		return fmt.Sprintf("%d(%s)", int(flag), strings.Join(flags, "|"))
	}
  	return "0()"
}

func main() {
	a := Active | Send
	s := a.String()
	fmt.Printf("%s\n", s)
}