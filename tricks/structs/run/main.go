package main

import (
	"fmt"

	"github.com/buddhamagnet/gbelt/tricks/structs"
)

func main() {
	// Refer to package level struct.
	fmt.Println("name:", structs.T.Name, "number", structs.T.Number)
	// Spawn an anonymous struct and populate it inline.
	golem := struct {
		name    string
		species string
	}{"frankenstein", "golem"}
	fmt.Println("name:", golem.name, "species:", golem.species)
}
