package sqrl

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/shirou/gopsutil/cpu"
)

type CPUInfo struct {
	CPU   int    `json:"cpu"`
	Cores int32  `json:"cores"`
	Speed string `json:"speed"`
	Model string `json:"model"`
}

func GetCPUInfo() {
	info, err := cpu.Info()
	if err != nil {
		log.Fatal(err)
	}
	mInfo := strings.Split(info[0].ModelName, "@") // mInfo contains speed and model
	c := CPUInfo{
		CPU:   runtime.NumCPU(),
		Cores: info[0].Cores,
		Speed: mInfo[1],
		Model: mInfo[0],
	}

	j, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(j))
}
