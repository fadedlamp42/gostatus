package components

import (
	"fmt"
	"time"
)

type Time struct{}

func (t *Time) Render() string {
	return time.Now().Format("Mon 01/02 15:04:05")
}

type ZonedTime struct {
	location       time.Location
	representation string
}

func (z *ZonedTime) Render() string {
	return fmt.Sprintf("%s: %s",
		z.representation,
		time.Now().In(&z.location).Format("15:04:05"))
}

func NewZonedTime(full string, short string) *ZonedTime {
	location, err := time.LoadLocation(full)
	if err != nil {
		fmt.Printf("WARNING: could not parse \"%s\" as location\n", full)
		location = time.Local
	}

	return &ZonedTime{
		location:       *location,
		representation: short,
	}
}
