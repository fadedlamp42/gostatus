package model

import (
	"fmt"
)

// Bar consists of a list of Components, seperator, and edge runes to be specified
// and renders a string usable as a status bar
type Bar struct {
	Components []Component
	Seperator  rune
	Edges      [2]string
}

func (b *Bar) Render() string {
	result := fmt.Sprintf("%s ", b.Edges[0])
	for i := range b.Components {
		renderOutput := fmt.Sprintf("%v ", b.Components[i].Render())

		if len(renderOutput) == 1 {
			continue
		}

		if i != 0 {
			result += fmt.Sprintf("%c ", b.Seperator)
		}

		result += renderOutput
	}
	result += b.Edges[1]

	return result
}
