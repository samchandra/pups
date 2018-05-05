package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	// get options from command line arguments
	name, apex := parseFlag()

	fmt.Println("Creating a new Go project with Pups")

	fmt.Printf("name: %s\n", name)
	fmt.Printf("apex: %v\n", apex)
	fmt.Println("")

	// create project folder based
	createProjectFolder(name)

	fmt.Printf("Copying template .....")

	// create project domain file
	// create http module
	// create mysql module
	// create up config
	// create docker config
}

func parseFlag() (string, bool) {
	name := flag.String(
		"name",
		fmt.Sprintf("Pups-Project#%d", time.Now().Unix()),
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

func createProjectFolder(name string) error {
	err := os.Mkdir(name, os.ModeDir)
	if err != nil {
		log.Fatal(err)
	}

	return err
}
