package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	projectSfName := flag.String("projectName", "", "Specify the project name")
	tPath := flag.String("path", "C:", "Specify the path")
	flag.Parse()

	projectName := *projectSfName + *projectSfName

	fmt.Println("Creating application V2...")

	cmdMkdir := exec.Command("powershell", "/C", "New-Item", "-Path", ""+*tPath+"/builder/", "-ItemType", "directory")
	cmdMkdir.Stdout = os.Stdout
	cmdMkdir.Stderr = os.Stderr
	errMkdir := cmdMkdir.Run()
	if errMkdir != nil {
		fmt.Println("Error creating directory:", errMkdir)
		return
	}
	fmt.Println("Directory created successfully")
	cmd := exec.Command("cmd", "/C", "cd "+*tPath+"/builder/ && wails init -n "+projectName+" -t vue")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Project created successfully")

	fmt.Println("Cloning project...")
	cmd2 := exec.Command("powershell", "/C", "Copy-Item", "-Path", "./../projects/"+projectName, "-Destination", ""+*tPath+"/builder/"+projectName, "-Recurse", "-Force")
	cmd2.Stdout = os.Stdout
	cmd2.Stderr = os.Stderr
	err2 := cmd2.Run()
	if err2 != nil {
		fmt.Println("Error:", err2)
		return
	}
	fmt.Println("Project cloned successfully")

	fmt.Printf("Building project %s...\n", projectName)
	cmd3 := exec.Command("wails", "build")
	cmd3.Dir = "" + *tPath + "/builder/" + projectName
	cmd3.Stdout = os.Stdout
	cmd3.Stderr = os.Stderr
	err3 := cmd3.Run()
	if err3 != nil {
		fmt.Println("Error:", err3)
		return
	}
	fmt.Println("Project build successfully")

	fmt.Println("Copying .exe file")
	cmd4 := exec.Command("powershell", "/C", "Copy-Item", "-Path", "./"+projectName+".exe", "-Destination", "../../../../", "-Recurse", "-Force")
	cmd4.Dir = "" + *tPath + "/builder/" + projectName + "/build/bin"
	cmd4.Stdout = os.Stdout
	cmd4.Stderr = os.Stderr
	err4 := cmd4.Run()
	if err4 != nil {
		fmt.Println("Error:", err4)
		return
	}
	fmt.Println("Project build successfully.")
}
