package components

import (
	"os/exec"
)

type Command struct {
	Text string
}

func (c *Command) Render() string {
	output, err := exec.Command(c.Text).Output()
	if err != nil {
		return err.Error()
	}

	if(len(output) == 0) {
		return ""
	} else {
		return string(output[:len(output)-1])
	}
}
