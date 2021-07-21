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
			components.NewZonedTime("America/Los_Angeles", "PST"),
			components.NewCPU(),
			&components.RAM{},
			components.NewNetwork("enp0s31f6"),
			&components.Song{},
			&components.Volume{},
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
