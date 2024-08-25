/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/peyton-spencer/b2-folder/cmd"
	"os/exec"
)

func main() {
	_, err := exec.Command("b2", "version").Output()
	if err != nil {
		panic("The b2 executable is not available. Please install it.")
	}
	cmd.Execute()
}
