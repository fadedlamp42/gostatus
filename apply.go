package main

import (
	"fmt"
	"os/exec"
)

func apply(top, bottom string) error {
	composite := fmt.Sprintf("%s;%s", top, bottom)
	return exec.Command("xsetroot", "-name", composite).Run()
}
