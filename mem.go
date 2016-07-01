package sqrl

import (
	"fmt"
	"math/big"
	"os"

	"github.com/dustin/go-humanize"
	"github.com/likexian/host-stat-go"
	"github.com/shirou/gopsutil/mem"
)

type Ram struct {
	total, free, swapSpace string
}

type OS struct {
	hostname, kernel, version string
}

func ramOSInfo() (Ram, OS) {
	v, e := mem.VirtualMemory()
	if e != nil {
		fmt.Println(e)
	}
	s, err := mem.SwapMemory()
	if err != nil {
		fmt.Println(err)
	}
	i, er := host_stat.GetHostInfo()
	if er != nil {
		fmt.Println(er)
	}
	total := v.Total
	free := v.Free
	swap := s.Free
	totalint := int64(total)
	totalbig := big.NewInt(totalint)
	totalmb := humanize.BigBytes(totalbig)
	freeint := int64(free)
	freebig := big.NewInt(freeint)
	freemb := humanize.BigBytes(freebig)
	swapint := int64(swap)
	swapbig := big.NewInt(swapint)
	swapmb := humanize.BigBytes(swapbig)
	hostname, e := os.Hostname()
	if e != nil {
		fmt.Println(e)
	}
	kernel := i.OSType
	version := i.OSRelease
	r := Ram{totalmb, freemb, swapmb}
	o := OS{hostname, version, kernel}
	return r, o
}

func printInfo() {
	r, o := ramOSInfo()
	fmt.Println(r)
	fmt.Println(o)
}
