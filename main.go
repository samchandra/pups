package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	// get options from command line arguments
	name, apex := parseFlag()

	fmt.Printf("name: %s\n", name)
	fmt.Printf("apex: %v", apex)

	// create project folder based
}

func parseFlag() (string, bool) {
	name := flag.String(
		"name",
		fmt.Sprintf("DindiProject-%d", time.Now().Unix()),
		"This is the name that the project directory will be called",
	)

	apex := flag.Bool(
		"apex",
		true,
		"Add Apex Up config up.json",
	)

	flag.Parse()
	return *name, *apex
}
