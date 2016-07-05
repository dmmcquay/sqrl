package sqrl

import (
	"math/big"
	"os"

	"github.com/dustin/go-humanize"
	"github.com/likexian/host-stat-go"
	"github.com/shirou/gopsutil/mem"
)

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

func RamOSInfo() (Ram, OS, Swap, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return Ram{}, OS{}, Swap{}, err
	}
	p, err := mem.SwapMemory()
	if err != nil {
		return Ram{}, OS{}, Swap{}, err
	}
	i, err := host_stat.GetHostInfo()
	if err != nil {
		return Ram{}, OS{}, Swap{}, err
	}

	r := Ram{
		Total: humanize.BigBytes(big.NewInt(int64(v.Total))),
		Free:  humanize.BigBytes(big.NewInt(int64(v.Free))),
		Used:  humanize.BigBytes(big.NewInt(int64(v.Total))),
	}

	s := Swap{
		Total: humanize.BigBytes(big.NewInt(int64(p.Total))),
		Free:  humanize.BigBytes(big.NewInt(int64(p.Free))),
		Used:  humanize.BigBytes(big.NewInt(int64(p.Used))),
	}
	hostname, err := os.Hostname()
	if err != nil {
		return Ram{}, OS{}, Swap{}, err
	}
	o := OS{
		Hostname: hostname,
		Kernel:   i.OSType,
		Version:  i.OSRelease,
	}

	return r, o, s, err
}
