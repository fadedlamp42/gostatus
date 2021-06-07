package components

import (
	"time"
)

type Time struct{}

func (t *Time) Render() string {
	return time.Now().Format("Mon 01/02 15:04:05")
}
