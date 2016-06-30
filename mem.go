package sqrl

import (
	"fmt"
	"math/big"

	"github.com/dmmcquay/likexian/host-stat-go"
	"github.com/dmmcquay/sqrl/dustin/go-humanize"
)

func meminfo() (uint64, uint64) {
	v, _ := mem.VirtualMemory()
	var total uint64
	total = v.Total
	free := v.Free
	return total, free
}

func swapinfo() (uint64, uint64) {
	s, _ := mem.SwapMemory()
	total := s.Total
	free := s.Free
	return total, free
}

func versioninfo() (string, string, string) {
	i, _ := host_stat.GetHostInfo()
	osVersion := i.OSType
	hostname := i.HostName
	version := i.OSRelease
	return osVersion, hostname, version
}
func mem() {
	total, free := meminfo()
	totalint := int64(total)
	totalbig := big.NewInt(totalint)
	swaptot, swapfree := swapinfo()
	totalmb := humanize.BigBytes(totalbig)
	freeint := int64(free)
	freebig := big.NewInt(freeint)
	freemb := humanize.BigBytes(freebig)
	swaptotint := int64(swaptot)
	swaptotbig := big.NewInt(swaptotint)
	swaptotmb := humanize.BigBytes(swaptotbig)
	swapfreeint := int64(swapfree)
	swapfreebig := big.NewInt(swapfreeint)
	swapfreemb := humanize.BigBytes(swapfreebig)
	osVers, host, kern := versioninfo()
	fmt.Println()
	fmt.Println("RAM Info")
	fmt.Println("Total RAM: ", totalmb)
	fmt.Println("Free RAM: ", freemb)
	fmt.Println()
	fmt.Println("SwapMemory Info")
	fmt.Println("Total SwapMemory: ", swaptotmb)
	fmt.Println("Free SwapMemory: ", swapfreemb)
	fmt.Println()
	fmt.Println("Version info")
	fmt.Println("OS Version: ", osVers)
	fmt.Println("Hostname: ", host)
	fmt.Println("Kernel Version: ", kern)
}
