package components

import (
	"fmt"
	"github.com/mackerelio/go-osstat/memory"
)

type RAM struct {}

func (r *RAM) Render() string {
	current, err := memory.Get()
	if err != nil {
		return fmt.Sprintf("ERROR(%v)", err)
	}

	total := float64(current.Total)
	usage := float64(current.Used)

	return fmt.Sprintf("RAM: %4.1f%%", usage/total * 100)
}
