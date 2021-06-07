package main

import (
	"fmt"
	"time"
)

func main() {
	top := buildTopBar()
	bottom := buildBottomBar()
	for true {
		err := apply(top.Render(), bottom.Render())
		if err != nil {
			fmt.Printf("WARNING: %v\n", err)
		}

		time.Sleep(time.Millisecond * (1000 / UPDATES_PER_SECOND))
	}
}
