package main

import (
	"github.com/fadedlamp42/gostatus/components"
	"github.com/fadedlamp42/gostatus/model"
)

const UPDATES_PER_SECOND = 1

func buildTopBar() model.Bar {
	return model.Bar{
		Seperator: '|',
		Edges:     [2]string{"[", "]"},
		Components: []model.Component{
			&components.Time{},
			components.NewCPU(),
			&components.RAM{},
			components.NewNetwork("wlo1"),
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
			components.NewBatteryBar("BAT0", 20, false),
			&components.Command{Text: "hp-battery.sh"},
		},
	}
}
