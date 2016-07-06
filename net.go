package sqrl

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/shirou/gopsutil/net"
)

type Interface struct {
	Name  string              `json:"name"`
	Addrs []net.InterfaceAddr `json:"addrs"` // IP Address, MAC Address and their respective ports
	Flags []string            `json:"flags"` // Types of links (up, down, multicast, etc.)
	Speed string              `json:"speed"`
}

func getInterfaceSpeed(name string) string {
	path := fmt.Sprintf("/sys/class/net/%s/speed", name)
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return "DNE"
	}
	speed := strings.Trim(string(dat), "\n")
	return speed
}

func GetInterfaces() ([]Interface, error) {
	n, err := net.Interfaces()
	if err != nil {
		return []Interface{}, err
	}

	interfaces := []Interface{}
	for _, i := range n {
		interfaces = append(
			interfaces,
			Interface{
				Name:  i.Name,
				Addrs: i.Addrs,
				Flags: i.Flags,
				Speed: getInterfaceSpeed(i.Name),
			},
		)
	}
	return interfaces, nil
}
