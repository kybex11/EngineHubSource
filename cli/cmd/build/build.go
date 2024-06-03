package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	projectName := flag.String("projectName", "", "Specify the project name")
	flag.Parse()

	cmd := exec.Command("cmd", "/C", "cd ./ && cd ../builder/ && wails init -n "+*projectName+" -t vue")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("Building project %s...\n", *projectName)
}

//wails build
