package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/samchandra/pups/pupstemplate"
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

	fmt.Printf("Copying template .....\n")

	// create project domain file
	createProjectDomainFile(name)

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

func createProjectDomainFile(name string) error {
	t, err := template.New("pups").Parse(pupstemplate.ProjectDomain())
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(filepath.Join(".", name, fmt.Sprintf("%s.go", name)))
	if err != nil {
		log.Println("create file: ", err)
		return err
	}
	defer f.Close()

	config := struct {
		Name string
	}{
		name,
	}

	err = t.Execute(f, config)
	if err != nil {
		log.Print("execute: ", err)
		return err
	}

	return err
}
