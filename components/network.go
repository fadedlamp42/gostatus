
package components

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Network struct {
	interfaceName string
	previous *stats
}

func (n *Network) Render() string {
	current := readStats(n.interfaceName)
	down := current.down - n.previous.down
	up := current.up - n.previous.up

	return fmt.Sprintf("D/U: %s/%s", humanBytes(down), humanBytes(up))
}

func NewNetwork(interfaceName string) *Network {
	return &Network{
		interfaceName: interfaceName,
		previous: readStats(interfaceName),
	}
}

type stats struct {
	down, up int64
}

func readStats(interfaceName string) *stats {
	downFile, err := os.Open(fmt.Sprintf("/sys/class/net/%s/statistics/rx_bytes", interfaceName))
	if err != nil {
		fmt.Println(fmt.Errorf("ERROR(%v)", err))
		return &stats{-1, -1}
	}
	defer downFile.Close()

	upFile, err := os.Open(fmt.Sprintf("/sys/class/net/%s/statistics/tx_bytes", interfaceName))
	if err != nil {
		fmt.Println(fmt.Errorf("ERROR(%v)", err))
		return &stats{-1, -1}
	}
	defer upFile.Close()

	scanner := bufio.NewScanner(downFile)
	scanner.Scan()
	down, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		fmt.Println(fmt.Errorf("ERROR(%v)", err))
		return &stats{-1, -1}
	}


	scanner = bufio.NewScanner(upFile)
	scanner.Scan()
	up, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		fmt.Println(fmt.Errorf("ERROR(%v)", err))
		return &stats{-1, -1}
	}

	return &stats{
		up: up,
		down: down,
	}
}

// shamelessly stolen from https://programming.guide/go/formatting-byte-size-to-human-readable-format.html
func humanBytes(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}
