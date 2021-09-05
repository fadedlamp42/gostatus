package components

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
    "strings"
)

type BatteryBar struct {
    identifier string
    width int
}

func (b* BatteryBar) Render() string {
  lhs, rhs := b.amounts()
	return fmt.Sprintf("[%s%s%s]", strings.Repeat("=", lhs-1), b.status(), strings.Repeat("_", rhs))
}

func NewBatteryBar(id string, width int) *BatteryBar {
	return &BatteryBar{
		identifier: id,
        width: width
	}
}

func (b* BatteryBar) amounts() (int, int) {
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

	percent := n / 100

	return ((percent*b.width), (b.width - (percent*b.width)))
}

func (b* BatteryBar) status() byte {
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
