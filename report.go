package sqrl

import (
	"github.com/dmmcquay/sqrl"
	"github.com/shirou/gopsutil/net"
)

type Report struct {
	Ram     `json:"Ram"`
	Swap    `json:"Swap"`
	OS      `json:"OS"`
	CPUInfo `json:"CPUInfo"`
	inter   []Interface `json:"Interface"`
}

type Ram struct {
	Total string `json:"total"`
	Free  string `json:"free"`
	Used  string `json:"used"`
}
type Swap struct {
	Total string `json:"total"`
	Free  string `json:"free"`
	Used  string `json:"used"`
}
type OS struct {
	Hostname string `json:"hostname"`
	Kernel   string `json:"kernel"`
	Version  string `json:"version"`
}
type CPUInfo struct {
	CPU   int    `json:"cpu"`
	Cores int32  `json:"cores"`
	Speed string `json:"speed"`
	Model string `json:"model"`
}
type Interface struct {
	Name  string              `json:"name"`
	Addrs []net.InterfaceAddr `json:"addrs"`
	Flags []string            `json:"flags"`
	Speed string              `json:"speed"`
}

func MakeReport() (Report, err) {
	ram, os, swap, err := sqrl.RamOSInfo()
	if err != nil {
		return Report{}, err
	}
	interfaces, err := sqrl.GetInterfaces()
	if err != nil {
		return Report{}, err
	}
	cpu, err := sqrl.GetCPUInfo()
	if err != nil {
		return Report{}, err
	}
	report := Report{ram, swap, os, cpu, interfaces}
	return report, err
}
