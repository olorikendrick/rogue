package main

import (
	"fmt"
	"log"
	"github.com/olorikendrick/rogue/internal/store"
)

func main() {
	fmt.Printf("Host %+v\n", store.GetHost())

	s := store.NewStore(".store")

	deps := []store.Dep{
		{
			Name:    "react",
			Version: [2]int{1,4},
		},
	}

	s.Add(deps)

	if err := s.Save(); err != nil {
		log.Fatalf("failed to save store: %v", err)
	}
}
