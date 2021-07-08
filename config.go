package main

import (
	"github.com/fadedlamp42/gostatus/model"
	"github.com/fadedlamp42/gostatus/components"
)

const UPDATES_PER_SECOND = 2

func buildTopBar() model.Bar {
	return model.Bar{
		Seperator: '|',
		Edges: [2]string {"[", "]"},
		Components: []model.Component{
			&components.Time{},
			components.NewCPU(),
			&components.RAM{},
			&components.Volume{},
			components.NewBattery("BAT0"),
		},
	}
}

func buildBottomBar() model.Bar {
	return model.Bar{
		Seperator: '|',
		Edges: [2]string{"", ""},
		Components: []model.Component{
			components.NewRSS(170, "  |  "),
		},
	}
}
