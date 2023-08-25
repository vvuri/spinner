package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

const dataDir = "./data"

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func printDir() {
	entries, err := os.ReadDir(dataDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		if path.Ext(e.Name()) == ".md" {
			fmt.Println(e.Name())
			dat, err := os.ReadFile(dataDir + "/" + e.Name())
			check(err)
			fmt.Print(string(dat))
		}
	}
}
