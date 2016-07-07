package sqrl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"log/syslog"
	"strings"

	"github.com/shirou/gopsutil/net"
)

type Interface struct {
	Name  string              `json:"name"`
	Addrs []net.InterfaceAddr `json:"addrs"` // IP Address, MAC Address and their respective ports
	Flags []string            `json:"flags"` // Types of links (up, down, multicast, etc.)
	Speed string              `json:"speed"`
}

type Interfaces struct {
	in []Interface
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
	fmt.Println(interfaces)
	return interfaces, nil
}

func (i Interface) String() string {
	j, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(j)
}

func (i *Interfaces) LogWriter() {
	l, e := syslog.New(syslog.LOG_NOTICE, "sqrl-InterfaceInfo")
	if e == nil {
		log.SetOutput(l)
	}

	log.Print(i)
}
