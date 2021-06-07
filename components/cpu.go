package components

import (
	"fmt"
	"github.com/mackerelio/go-osstat/cpu"
)

type CPU struct {
	previous *cpu.Stats
}

func (c* CPU) Render() string {
	current, err := cpu.Get()
	if err != nil {
		return fmt.Sprintf("ERROR(%v)", err)
	}

	total := float64(current.Total - c.previous.Total)
	usage := (current.System + current.User + current.Nice) - (c.previous.System + c.previous.User + c.previous.Nice)

	c.previous = current
	return fmt.Sprintf("CPU: %4.1f%%", float64(usage)/total * 100)
}

func NewCPU() *CPU {
	stats, err := cpu.Get()
	if err != nil {
		stats = &cpu.Stats{
			Total: 0,
			System: 0,
			User: 0,
			Nice: 0,
		}
	}

	return &CPU{
		previous: stats,
	}
}
