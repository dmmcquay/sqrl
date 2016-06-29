package sqrl

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/cpu"
)

func CPUInfo() error {
	info, err := cpu.Info()
	if err != nil {
		log.Printf("error getting cpu: %v", err)
		return err
	}
	fmt.Println(info)
	return nil
}
