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
			Version: "1.2",
		},
	}

	s.Add(deps)

	if err := s.Save(); err != nil {
		log.Fatalf("failed to save store: %v", err)
	}
}
