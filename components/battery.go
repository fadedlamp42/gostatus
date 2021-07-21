package components

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Battery struct {
	identifier string
}

func (b* Battery) Render() string {
	return fmt.Sprintf("BAT: %3d%%{%c}", b.amount(), b.status())
}

func NewBattery(id string) *Battery {
	return &Battery{
		identifier: id,
	}
}

func (b* Battery) amount() int64 {
	file, err := os.Open(fmt.Sprintf("/sys/class/power_supply/%s/capacity", b.identifier))
	if err != nil {
		fmt.Println(fmt.Errorf("ERROR(%v)", err))
		return -1
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	n, err := strconv.ParseInt(scanner.Text(), 10, 32)
	if err != nil {
		fmt.Println(fmt.Errorf("ERROR(%v)", err))
		return -1
	}

	if n < 0 {
		n = 0
	} else if n > 100 {
		n = 100
	}

	return n
}

func (b* Battery) status() byte {
	file, err := os.Open(fmt.Sprintf("/sys/class/power_supply/%s/status", b.identifier))
	if err != nil {
		fmt.Println(fmt.Errorf("ERROR(%v)", err))
		return 'E'
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	switch(scanner.Text()) {
	case "Charging": return '+'
	case "Discharging": return '-'
	case "Full": return '='
	}
	return '?'
}
