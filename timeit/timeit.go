package main

import (
	"os"
	"os/exec"
	"strings"
	"time"
	"fmt"
)

func main() {
	argsWithoutProg := strings.Join(os.Args[1:], " ")
	if argsWithoutProg == "" {
		fmt.Println("no command passed to time")
	}
	cmd := exec.Command("cmd", "/c", argsWithoutProg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	start := time.Now()

	if err := cmd.Run(); err != nil {
		fmt.Println("error running command")
		return
	}

	fmt.Println("\nExecution took ", time.Since(start))
}
