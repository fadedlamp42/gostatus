package main

import (
	"github.com/fadedlamp42/gostatus/components"
	"github.com/fadedlamp42/gostatus/model"
)

const UPDATES_PER_SECOND = 2

func buildTopBar() model.Bar {
	return model.Bar{
		Seperator: '|',
		Edges:     [2]string{"[", "]"},
		Components: []model.Component{
			&components.Time{},
			components.NewZonedTime("America/Los_Angeles", "PST"),
			components.NewZonedTime("", "UTC"),
			components.NewCPU(),
			&components.RAM{},
			// TODO optimize current-monitor script, causes CPU spikes currently
			//components.NewMinWidth(1920, components.NewNetwork("enp0s31f6")),
			//components.NewMinWidth(1920, &components.Song{}),
			components.NewNetwork("enp0s31f6"),
			&components.Song{},
			&components.Volume{},
		},
	}
}

func buildBottomBar() model.Bar {
	return model.Bar{
		Seperator: '|',
		Edges:     [2]string{"", ""},
		Components: []model.Component{
			components.NewRSS(233, " | "),
		},
	}
}
