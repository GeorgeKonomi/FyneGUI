package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

type Command interface {
	Execute() (string, error)
}

type runModemPanelTest struct{}

func (p *runModemPanelTest) Execute() (string, error) {
	npx := "npx"
	playwright := "playwright"
	arg0 := "test"
	arg1 := "/modemPanel"
	arg2 := "--headed"
	log.Println("Executing command: " + npx + " " + playwright + " " + arg0 + " " + arg1 + " " + arg2)

	cmd := exec.Command(npx, playwright, arg0, arg1, arg2)
	cmd.Dir = "../../trafficGenerator/playwrightTests/tests"
	// By runnning ./testgui in the work/trafficGenerator/Qt/testGui directory to open the project, that is our default cmd.Dir
	// momemPanel tests should be executed through work/trafficGenerator/playwrightTest/tests directory
	// .. changes the current directory to the parent directory
	// so the ../../trafficGenerator/playwrightTests/tests from the current directory has absolute path
	// work/trafficGenerator/

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	}

	fmt.Println("Result: " + out.String())

	return out.String(), nil
}

func ExecuteCmd(pCmdName string) (string, error) {
	commands := map[string]Command{
		"modemPanelTest": &runModemPanelTest{},
	}

	if command := commands[pCmdName]; command == nil {
		return "", fmt.Errorf("fatal error")
	} else {
		return command.Execute()
	}
}

// func ExecuteCmd() (string, error) {
// 	npx := "npx"
// 	playwright := "playwright"
// 	arg0 := "test"
// 	arg1 := "/modemPanel"
// 	arg2 := "--headed"
// 	log.Println("Executing command: " + npx + " " + playwright + " " + arg0 + " " + arg1 + " " + arg2)

// 	cmd := exec.Command(npx, playwright, arg0, arg1, arg2)
// 	cmd.Dir = "../../trafficGenerator/playwrightTests/tests"
// 	// By runnning ./testgui in the work/trafficGenerator/Qt/testGui directory to open the project, that is our default cmd.Dir
// 	// momemPanel tests should be executed through work/trafficGenerator/playwrightTest/tests directory
// 	// .. changes the current directory to the parent directory
// 	// so the ../../trafficGenerator/playwrightTests/tests from the current directory has absolute path
// 	// work/trafficGenerator/

// 	var out bytes.Buffer
// 	var stderr bytes.Buffer
// 	cmd.Stdout = &out
// 	cmd.Stderr = &stderr

// 	err := cmd.Run()

// 	if err != nil {
// 		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
// 	}

// 	fmt.Println("Result: " + out.String())

// 	return out.String(), nil
// }
