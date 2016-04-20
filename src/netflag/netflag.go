package main

import (
	"fmt"
)

type Flags uint

const (
	Flagup Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPint
	FlagMuticast
)

// IsUp ...
func IsUp(v Flags) bool {
	return v&Flagup == Flagup
}

// TurnDown ...
func TurnDown(v *Flags)  {
	*v &^= Flagup
}

// SetBroadcast ...
func SetBroadcast(v *Flags)  {
	*v |= FlagBroadcast
}

// IsCast ...
func IsCast(v Flags) bool {
	return v&(FlagBroadcast|FlagMuticast) != 0
}

func main() {
	var v Flags = FlagMuticast | Flagup

	fmt.Printf("%b %t\n", v, IsUp(v))

	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))

	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	fmt.Printf("%b %t\n", v, IsCast(v))
}