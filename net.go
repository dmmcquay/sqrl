package sqrl

import (
	"fmt"

	"github.com/shirou/gopsutil/net"
)

type Interfaces struct {
	Name  string
	Addrs []net.InterfaceAddr
	Flags []string
}

func NetworkInfo() {
	a, _ := net.Interfaces()

	fmt.Println()
	fmt.Println("\t\t\tNetwork Info")

	for i := 0; i < 5; i++ {
		info := Interfaces{a[i].Name, a[i].Addrs, a[i].Flags}
		fmt.Println(info)
	}
	fmt.Println()

}
