/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/peyton-spencer/b2-folder/cmd"
)

func main() {
	out, err := exec.Command("b2", "version").Output()
	if err != nil {
		panic("The b2 executable is not available. Please install it.")
	}
	outStr := strings.TrimPrefix(string(out), "b2 command line tool, version ")
	fmt.Printf("Using b2 version %s\n", outStr)
	cmd.Execute()
}
