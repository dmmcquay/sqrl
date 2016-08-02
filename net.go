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

func newInterfaces() Interfaces {
	return Interfaces{
		in: []Interface{},
	}
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

func GetInterfaces() (Interfaces, error) {
	n, err := net.Interfaces()
	if err != nil {
		return Interfaces{}, err
	}

	interfaces := newInterfaces()
	for _, i := range n {
		interfaces.in = append(
			interfaces.in,
			Interface{
				Name:  i.Name,
				Addrs: i.Addrs,
				Flags: i.Flags,
				Speed: getInterfaceSpeed(i.Name),
			},
		)
	}
	//fmt.Printf("%+v\n", interfaces)
	return interfaces, nil
}

func (i Interfaces) String() string {
	final := ""
	for _, e := range i.in {
		j, err := json.Marshal(e)
		if err != nil {
			return ""
		}
		final += string(j)
	}
	return final
}

func (i *Interfaces) LogWriter() {
	l, e := syslog.New(syslog.LOG_NOTICE, "sqrl-InterfaceInfo")
	if e == nil {
		log.SetOutput(l)
	}

	log.Print(i)
}
