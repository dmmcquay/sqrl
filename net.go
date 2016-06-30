package sqrl

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/shirou/gopsutil/net"
)

func netInfo() {
	a, _ := net.Interfaces()

	fmt.Println()
	fmt.Println("\t\t\tNetwork Info")

	netSpeed, err := exec.Command("ethtool", "eno16777984").Output()
	if err != nil {
		log.Fatal(err)
	}
	output1 := string(netSpeed)
	result := strings.Split(output1, "\n")

	fmt.Println(a[1].Name, a[1].Addrs, a[1].Flags, result[9])

}
