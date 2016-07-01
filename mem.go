package sqrl

import (
	"fmt"
	"math/big"

	"github.com/dustin/go-humanize"
	"github.com/likexian/host-stat-go"
	"github.com/shirou/gopsutil/mem"
)

type Ram struct {
	total, free, swapSpace uint64
}

type OS struct {
	hostname, kernel, version string
}

func ramOSInfo() (Ram, OS) {
	v, _ := mem.VirtualMemory()
	s, _ := mem.SwapMemory()
	i, _ := host_stat.GetHostInfo()
	o := OS{i.HostName, i.OSType, i.OSRelease}
	r := Ram{v.Total, v.Free, s.Free}
	return r, o
}
func printInfo() {
	mem, os := ramOSInfo()
	total := mem.total
	free := mem.free
	swap := mem.swapSpace
	totalint := int64(total)
	totalbig := big.NewInt(totalint)
	totalmb := humanize.BigBytes(totalbig)
	freeint := int64(free)
	freebig := big.NewInt(freeint)
	freemb := humanize.BigBytes(freebig)
	swapint := int64(swap)
	swapbig := big.NewInt(swapint)
	swapmb := humanize.BigBytes(swapbig)
	hostname := os.hostname
	kernel := os.kernel
	version := os.version
	fmt.Println()
	fmt.Println("RAM Info")
	fmt.Println("Total RAM: ", totalmb)
	fmt.Println("Free RAM: ", freemb)
	fmt.Println()
	fmt.Println("SwapSpace Info")
	fmt.Println("SwapSpace: ", swapmb)
	fmt.Println()
	fmt.Println("Version info")
	fmt.Println("OS Version: ", version)
	fmt.Println("Hostname: ", hostname)
	fmt.Println("Kernel Version: ", kernel)
	fmt.Println()
}
