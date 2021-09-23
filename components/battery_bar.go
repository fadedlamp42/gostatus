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
	width      int
	showStatus bool
}

func (b *BatteryBar) Render() string {
	lhs, rhs := b.amounts()
	if b.showStatus {
		return fmt.Sprintf("%s%c%s", strings.Repeat("█", lhs-1), b.status(), strings.Repeat("_", rhs))
	} else {
		return fmt.Sprintf("%s%s", strings.Repeat("█", lhs), strings.Repeat("_", rhs))
	}

}

func NewBatteryBar(id string, width int, showStatus bool) *BatteryBar {
	return &BatteryBar{
		identifier: id,
		width:      width,
		showStatus: showStatus,
	}
}

func (b *BatteryBar) amounts() (int, int) {
	file, err := os.Open(fmt.Sprintf("/sys/class/power_supply/%s/capacity", b.identifier))
	if err != nil {
		fmt.Println(fmt.Errorf("ERROR(%v)", err))
		return -1, -1
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	n, err := strconv.ParseInt(scanner.Text(), 10, 32)
	if err != nil {
		fmt.Println(fmt.Errorf("ERROR(%v)", err))
		return -1, -1
	}

	if n < 0 {
		n = 0
	} else if n > 100 {
		n = 100
	}

	percent := float64(n) / float64(100)

	return int(percent * float64(b.width)), int(b.width - int(percent*float64(b.width)))
}

func (b *BatteryBar) status() byte {
	file, err := os.Open(fmt.Sprintf("/sys/class/power_supply/%s/status", b.identifier))
	if err != nil {
		fmt.Println(fmt.Errorf("ERROR(%v)", err))
		return 'E'
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	switch scanner.Text() {
	case "Charging":
		return '+'
	case "Discharging":
		return '-'
	case "Full":
		return '='
	}
	return '?'
}
