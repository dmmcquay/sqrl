package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"runtime"
)

func main() {
	info, err := cpu.Info()
	serial, erro := disk.IOCounters()
	fmt.Println("The serial number is:", serial["serialNumber"])
	fmt.Println("The model/speed of CPU is:", info[0].ModelName)
	fmt.Println("The number of cores per CPU is:", info[0].Cores)
	fmt.Println("The total number of CPUs is:", runtime.NumCPU())
	fmt.Println(err)
	fmt.Println(erro)
}
