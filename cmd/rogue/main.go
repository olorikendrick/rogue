package main

import (
	"fmt"
	"github.com/olorikendrick/rogue/internal/store"
)

func main() {
	fmt.Printf("Host %+v\n", store.GetHost())
}
