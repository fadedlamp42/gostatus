package components

import (
	"fmt"
	"os/exec"
)

type Song struct {}

func (s *Song) Render() string {
	out, err := exec.Command("playing").Output()
	if err != nil {
		return fmt.Sprintf("ERROR(%v)", err)
	}

	return fmt.Sprintf("SONG: %s", out[:len(out)-1])
}
