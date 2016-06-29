package sqrl

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

func MemInfo() {
	v, _ := mem.VirtualMemory()

	fmt.Println(v)
}
