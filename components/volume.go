package components

import (
	"fmt"
	"github.com/itchyny/volume-go"
)

type Volume struct{}

func (v *Volume) Render() string {
	vol, err := volume.GetVolume()
	if err != nil {
		return fmt.Sprintf("ERROR(%v)", err)
	}

	muted, err := volume.GetMuted()
	if err != nil {
		return fmt.Sprintf("ERROR(%v)", err)
	}

	if muted {
		return "VOL: MUTED"
	}

	return fmt.Sprintf("VOL: %3d%%", vol)
}
