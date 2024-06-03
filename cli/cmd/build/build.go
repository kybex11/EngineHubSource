package main

import (
	"flag"
	"fmt"
)

func main() {
	projectName := flag.String("projectName", "", "Specify the project name")
	flag.Parse()

	fmt.Printf("Building project %s...\n", *projectName)
}
