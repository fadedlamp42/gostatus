package components

import (
	"fmt"
	"github.com/fadedlamp42/gostatus/model"
	"os/exec"
	"strconv"
)

type MinWidth struct {
	width int
	inner model.Component
}

func (m *MinWidth) Render() string {
	out, err := exec.Command("current-width").Output()
	if err != nil {
		return fmt.Sprintf("ERROR(%v)", err)
	}

	width, err := strconv.ParseInt(string(out[:len(out)-1]), 10, 32)
	if err != nil {
		return fmt.Sprintf("ERROR(%v)", err)
	}

	if int(width) >= m.width {
		return m.inner.Render()
	} else {
		return ""
	}
}

func NewMinWidth(width int, inner model.Component) *MinWidth {
	return &MinWidth{
		width: width,
		inner: inner,
	}
}
