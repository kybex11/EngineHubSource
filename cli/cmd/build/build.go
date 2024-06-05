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

	fmt.Println("Creating application V2...")

	cmd := exec.Command("cmd", "/C", "cd ./ && cd ../builder/ && wails init -n "+*projectName+""+*projectName+" -t vue")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Cloning project %s...\n", *projectName)

	cmd2 := exec.Command("cmd", "/C", "Copy-Item -Path ../../../projects/"+*projectName+""+*projectName+" -Destination ../builder/ -Recurse -Force")
	cmd2.Stdout = os.Stdout
	cmd2.Stderr = os.Stderr
	err2 := cmd2.Run()
	if err2 != nil {
		fmt.Println("Error:", err2)
	}
	fmt.Println("Project cloned successfully")
	fmt.Printf("Building project %s...\n", *projectName)

	cmd3 := exec.Command("cmd", "/C", "cd ./ && cd ../builder/"+*projectName+""+*projectName+" && wails build")
	cmd3.Stdout = os.Stdout
	cmd3.Stderr = os.Stderr
	err3 := cmd3.Run()
	if err3 != nil {
		fmt.Println("Error:", err3)
	}

}
