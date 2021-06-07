package model

import (
	"fmt"
)

// Bar consists of a list of Components, seperator, and edge runes to be specified
// and renders a string usable as a status bar
type Bar struct {
	Components []Component
	Seperator rune
	Edges [2]string
}

func (b *Bar) Render() string {
	result := fmt.Sprintf("%s ", b.Edges[0])
	for i := range(b.Components) {
		if i != 0 {
			result += fmt.Sprintf("%c ", b.Seperator)
		}

		result += fmt.Sprintf("%v ", b.Components[i].Render())
	}
	result += b.Edges[1]

	return result
}
