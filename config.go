package main

import (
	"github.com/fadedlamp42/gostatus/components"
	"github.com/fadedlamp42/gostatus/model"
)

const UPDATES_PER_SECOND = 0.1

func buildTopBar() model.Bar {
	return model.Bar{
		Seperator: '|',
		Edges:     [2]string{"[", "]"},
		Components: []model.Component{
			&components.Time{},
			//components.NewZonedTime("America/Los_Angeles", "PST"),
			components.NewZonedTime("America/New_York", "EST"),
			components.NewZonedTime("", "UTC"),
			components.NewCPU(),
			&components.RAM{},
			components.NewNetwork("wlp108s0"),
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
			components.NewBattery("BAT0"),
			components.NewBatteryBar("BAT0", 50),
		},
	}
}
